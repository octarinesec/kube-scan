package risk

import (
	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sort"
)

type ScoreConfig struct {
	Name                       string             `yaml:"name"`
	Title                      string             `yaml:"title"`
	ShortDescription           string             `yaml:"shortDescription"`
	Description                string             `yaml:"description"`
	Confidentiality            RiskCIACategory    `yaml:"confidentiality"`
	ConfidentialityDescription string             `yaml:"confidentialityDescription"`
	Integrity                  RiskCIACategory    `yaml:"integrity"`
	IntegrityDescription       string             `yaml:"integrityDescription"`
	Availability               RiskCIACategory    `yaml:"availability"`
	AvailabilityDescription    string             `yaml:"availabilityDescription"`
	Exploitability             RiskFactorCategory `yaml:"exploitability"`
	AttackVector               RiskAttackVector   `yaml:"attackVector"`
	Scope                      RiskScope          `yaml:"scope"`
	Handler                    string             `yaml:"handler"`
}

type ScoreRange struct {
	MinScore    float64 `yaml:"min"`
	LowScore    float64 `yaml:"low"`
	MediumScore float64 `yaml:"medium"`
	MaxScore    float64 `yaml:"max"`
}

type AttackVectorConfig struct {
	Remote float64 `yaml:"remote"`
	Local  float64 `yaml:"local"`
}

type ExploitabilityConfig struct {
	High     float64 `yaml:"high"`
	Moderate float64 `yaml:"moderate"`
	Low      float64 `yaml:"low"`
	VeryLow  float64 `yaml:"veryLow"`
}

type CIAConfig struct {
	High float64 `yaml:"high"`
	Low  float64 `yaml:"low"`
	None float64 `yaml:"none"`
}

type ScopeFactorConfig struct {
	None    float64 `yaml:"none"`
	Host    float64 `yaml:"host"`
	Cluster float64 `yaml:"cluster"`
}

type Config struct {
	ExpConst               float64              `yaml:"expConst"`
	ImpactConst            float64              `yaml:"impactConst"`
	AttackVector           AttackVectorConfig   `yaml:"attackVector"`
	Exploitability         ExploitabilityConfig `yaml:"exploitability"`
	ScopeFactor            ScopeFactorConfig    `yaml:"scopeFactor"`
	CIAScore               CIAConfig            `yaml:"ciaScore"`
	RiskCategory           ScoreRange           `yaml:"riskCategory"`
	IndividualRiskCategory ScoreRange           `yaml:"individualRiskCategory"`
	Basic                  []ScoreConfig        `yaml:"basic"`
	Remediation            []ScoreConfig        `yaml:"remediation"`
}

func (config *Config) Validate() bool {
	for _, b := range config.Basic {
		if !validateScoreConfig(b) {
			return false
		}
	}

	for _, r := range config.Remediation {
		if !validateScoreConfig(r) {
			return false
		}
	}

	return true
}

func (config *Config) GetCIAScore(factor RiskCIACategory) float64 {
	switch factor {
	case CIAHigh:
		return config.CIAScore.High
	case CIALow:
		return config.CIAScore.Low
	default:
		return config.CIAScore.None
	}
}

func (config *Config) GetScopeScore(scope RiskScope) float64 {
	switch scope {
	case ScopeHost:
		return config.ScopeFactor.Host
	case ScopeCluster:
		return config.ScopeFactor.Cluster
	default:
		return config.ScopeFactor.None
	}
}

func (config *Config) GetAtackVectorScore(av RiskAttackVector) float64 {
	switch av {
	case AttackVectorRemote:
		return config.AttackVector.Remote
	case AttackVectorLocal:
		return config.AttackVector.Local
	default:
		return 0
	}
}

func (config *Config) GetExploitabilityScore(exp RiskFactorCategory) float64 {
	switch exp {
	case FactorHigh:
		return config.Exploitability.High
	case FactorModerate:
		return config.Exploitability.Moderate
	case FactorLow:
		return config.Exploitability.Low
	case FactorVeryLow:
		return config.Exploitability.VeryLow
	default:
		return 0
	}
}

func validateScoreConfig(sc ScoreConfig) bool {
	return ValidateCIACategory(sc.Confidentiality) &&
		ValidateCIACategory(sc.Integrity) &&
		ValidateCIACategory(sc.Availability) &&
		ValidateFactorCategory(sc.Exploitability) &&
		ValidateAttackVector(sc.AttackVector) &&
		ValidateScope(sc.Scope)
}

func NewConfigFromFile(configFile string) *Config {
	configStr, err := ioutil.ReadFile(configFile)
	if err != nil {
		glog.Fatalf("Failed to read the risk config from file %v: %v", configFile, err)
	}

	var result *Config
	if err := yaml.Unmarshal(configStr, &result); err != nil {
		glog.Fatalf("Failed unmarshaling insights risk with err: %v (conf string = %s)", err, configStr)
	}

	if !result.Validate() {
		glog.Fatalf("Failed parsing insights risk config: %s", configStr)
	}

	sort.Slice(result.Basic, func(i, j int) bool {
		return result.Basic[i].Scope != ScopeNone
	})

	return result
}
