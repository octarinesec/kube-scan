package risk

import (
	"fmt"
	"math"
)

type GetStatusFunc func(conf *Config, workload IWorloadRisk, score int) RiskStatus

func DefaultStateGetter(conf *Config, workload IWorloadRisk, score int) RiskStatus {
	if float64(score) == conf.RiskCategory.MinScore {
		// todo: change this to none
		return RiskStatusOpen
	}

	return RiskStatusOpen
}

type Formula struct {
	config *Config
}

func NewFormula(config *Config) *Formula {
	formula := &Formula{
		config: config,
	}

	return formula
}

func (formula *Formula) calcBasicScore(s RiskScope, e RiskFactorCategory, av RiskAttackVector, c, i, a RiskCIACategory) float64 {
	confidentiality := formula.config.GetCIAScore(c)
	integrity := formula.config.GetCIAScore(i)
	availability := formula.config.GetCIAScore(a)
	scope := formula.config.GetScopeScore(s)
	attackVector := formula.config.GetAtackVectorScore(av)
	exploitability := formula.config.GetExploitabilityScore(e)

	baseImpactScore := formula.config.ImpactConst * (1 - ((1 - confidentiality) * (1 - integrity) * (1 - availability)))
	impactScore := scope * baseImpactScore

	expScore := formula.config.ExpConst * attackVector * exploitability

	score := impactScore + expScore
	score = math.Round(score*10) / 10
	return score
}

func (formula *Formula) score2category(score float64) RiskCategory {
	if score <= formula.config.RiskCategory.LowScore {
		return Low
	}

	if score <= formula.config.RiskCategory.MediumScore {
		return Medium
	}

	return High
}

func (formula *Formula) individualScore2category(score float64) RiskCategory {
	if score <= formula.config.IndividualRiskCategory.LowScore {
		return Low
	}

	if score <= formula.config.IndividualRiskCategory.MediumScore {
		return Medium
	}

	return High
}

func (formula *Formula) trimScore(score float64) float64 {
	if score < formula.config.RiskCategory.MinScore {
		return formula.config.RiskCategory.MinScore
	}

	if score > formula.config.RiskCategory.MaxScore {
		return formula.config.RiskCategory.MaxScore
	}

	return score
}

func (formula *Formula) trimIndividualScore(score float64) float64 {
	if score < formula.config.IndividualRiskCategory.MinScore {
		return formula.config.IndividualRiskCategory.MinScore
	}

	if score > formula.config.IndividualRiskCategory.MaxScore {
		return formula.config.IndividualRiskCategory.MaxScore
	}

	return score
}

func (formula *Formula) CalculateRiskWithStatusGetter(workload IWorloadRisk, statusFunc GetStatusFunc) (*Risk, error) {
	items := make([]RiskItem, 0)

	avs2score := make(map[string]float64)
	for _, basic := range formula.config.Basic {
		if handler, err := GetHandler(basic.Handler); err != nil {
			return nil, fmt.Errorf("error calculating the risk formula: %v", err.Error())
		} else if !handler(workload) {
			continue
		}

		remC, remI, remA := CIANone, CIANone, CIANone
		for _, rem := range formula.config.Remediation {
			if handler, err := GetHandler(rem.Handler); err != nil {
				return nil, fmt.Errorf("error calculating the risk formula: %v", err.Error())
			} else if handler(workload) && rem.Scope == basic.Scope {
				if rem.Confidentiality.GreaterThan(remC) {
					remC = rem.Confidentiality
				}
				if rem.Integrity.GreaterThan(remI) {
					remI = rem.Integrity
				}
				if rem.Availability.GreaterThan(remA) {
					remA = rem.Availability
				}
			}
		}

		c := basic.Confidentiality.Minus(remC)
		i := basic.Integrity.Minus(remI)
		a := basic.Availability.Minus(remA)

		curScore := formula.calcBasicScore(basic.Scope, basic.Exploitability, basic.AttackVector, c, i, a)
		curScore = formula.trimIndividualScore(curScore)

		items = append(items, NewRiskItem(Basic, basic, &curScore, formula.individualScore2category(curScore)))
		switch basic.Scope {
		case ScopeNone:
			avExists := false
			for _, scope := range NotNoneScopes {
				key := fmt.Sprintf("%v-%v", basic.AttackVector, scope)
				if s, ok := avs2score[key]; ok {
					avExists = true
					if curScore > s {
						avs2score[key] = curScore
					}
				}
			}

			if !avExists {
				key := fmt.Sprintf("%v-%v", basic.AttackVector, basic.Scope)
				if s, ok := avs2score[key]; !ok || curScore > s {
					avs2score[key] = curScore
				}
			}
		default:
			key := fmt.Sprintf("%v-%v", basic.AttackVector, basic.Scope)
			if s, ok := avs2score[key]; !ok || curScore > s {
				avs2score[key] = curScore
			}
		}
	}

	for _, rem := range formula.config.Remediation {
		if handler, err := GetHandler(rem.Handler); err != nil {
			return nil, fmt.Errorf("error calculating the risk formula: %v", err.Error())
		} else if handler(workload) {
			items = append(items, NewRiskItem(Remediation, rem, nil, None))
		}
	}

	score := 0.0
	for _, individualScore := range avs2score {
		score += math.Pow(individualScore, 2)
	}
	score = math.Sqrt(score)

	score = formula.trimScore(score)
	score = math.Round(score)

	return &Risk{
		RiskScore:    int(score),
		RiskCategory: formula.score2category(score),
		RiskItems:    items,
		RiskStatus:   statusFunc(formula.config, workload, int(score)),
	}, nil
}

func (formula *Formula) CalculateRisk(workload IWorloadRisk) (*Risk, error) {
	return formula.CalculateRiskWithStatusGetter(workload, DefaultStateGetter)
}
