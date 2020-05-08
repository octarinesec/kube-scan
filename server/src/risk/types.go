package risk

import (
	"github.com/go-bongo/bongo"
	"github.com/golang/glog"
	"github.com/toolkits/slice"
	"kube-scan/common"
)

type RiskCategory string
type RiskType string
type RiskFactorCategory string
type RiskCIACategory string
type RiskStatus string
type RiskAttackVector string
type RiskScope string

const (
	None   = RiskCategory("None")
	Low    = RiskCategory("Low")
	Medium = RiskCategory("Medium")
	High   = RiskCategory("High")

	FactorNone     = RiskFactorCategory("None")
	FactorVeryLow  = RiskFactorCategory("VeryLow")
	FactorLow      = RiskFactorCategory("Low")
	FactorModerate = RiskFactorCategory("Moderate")
	FactorHigh     = RiskFactorCategory("High")

	CIANone = RiskCIACategory("None")
	CIALow  = RiskCIACategory("Low")
	CIAHigh = RiskCIACategory("High")

	AttackVectorLocal  = RiskAttackVector("Local")
	AttackVectorRemote = RiskAttackVector("Remote")

	ScopeCluster = RiskScope("Cluster")
	ScopeHost    = RiskScope("Host")
	ScopeNone    = RiskScope("None")

	Basic       = RiskType("Basic")
	Remediation = RiskType("Remediation")

	//RiskStatusNone       = RiskStatus("None")
	RiskStatusOpen       = RiskStatus("Open")
	RiskStatusInProgress = RiskStatus("InProgress")
	RiskStatusAccepted   = RiskStatus("Accepted")
)

var (
	NotNoneScopes = []RiskScope{ScopeHost, ScopeCluster}
)

func (cia RiskCIACategory) getOrder() int {
	switch cia {
	case CIANone:
		return 0
	case CIALow:
		return 1
	case CIAHigh:
		return 2
	default:
		glog.Errorf("unknown cia category %v", cia)
		return -1
	}
}

func (cia RiskCIACategory) GreaterThan(other RiskCIACategory) bool {
	return cia.getOrder() > other.getOrder()
}

func (cia RiskCIACategory) Minus(other RiskCIACategory) RiskCIACategory {
	switch other {
	case CIAHigh:
		return CIANone
	case CIALow:
		switch cia {
		case CIAHigh:
			return CIALow
		default:
			return CIANone
		}
	default:
		return cia
	}
}

func validate(t interface{}, all ...interface{}) bool {
	return slice.Contains(all, t)
}

func ValidateStatus(status RiskStatus) bool {
	return validate(status, RiskStatusOpen, RiskStatusInProgress, RiskStatusAccepted)
}

func ValidateScope(scope RiskScope) bool {
	return validate(scope, ScopeCluster, ScopeHost, ScopeNone)
}

func ValidateFactorCategory(rfc RiskFactorCategory) bool {
	return validate(rfc, FactorNone, FactorVeryLow, FactorLow, FactorModerate, FactorHigh)
}

func ValidateCIACategory(cia RiskCIACategory) bool {
	return validate(cia, CIANone, CIALow, CIAHigh)
}

func ValidateAttackVector(av RiskAttackVector) bool {
	return validate(av, AttackVectorLocal, AttackVectorRemote)
}

type RiskItem struct {
	Name                       string             `json:"name" bson:"name"`
	RiskCategory               RiskCategory       `json:"riskCategory" bson:"riskCategory"`
	RiskType                   RiskType           `json:"type" bson:"type"`
	Title                      string             `json:"title" bson:"title"`
	ShortDescription           string             `json:"shortDescription" bson:"shortDescription"`
	Description                string             `json:"description" bson:"description"`
	Confidentiality            RiskCIACategory    `json:"confidentiality" bson:"confidentiality"`
	ConfidentialityDescription string             `json:"confidentialityDescription" bson:"confidentialityDescription"`
	Integrity                  RiskCIACategory    `json:"integrity" bson:"integrity"`
	IntegrityDescription       string             `json:"integrityDescription" bson:"integrityDescription"`
	Availability               RiskCIACategory    `json:"availability" bson:"availability"`
	AvailabilityDescription    string             `json:"availabilityDescription" bson:"availabilityDescription"`
	Exploitability             RiskFactorCategory `json:"exploitability" bson:"exploitability"`
	AttackVector               RiskAttackVector   `json:"attackVector" bson:"attackVector"`
	Scope                      RiskScope          `json:"scope" bson:"scope"`
	Score                      *float64           `json:"score" bson:"score"`
}

func NewRiskItem(riskType RiskType, riskConfig ScoreConfig, score *float64, category RiskCategory) RiskItem {
	return RiskItem{
		Name:                       riskConfig.Name,
		RiskType:                   riskType,
		RiskCategory:               category,
		Title:                      riskConfig.Title,
		ShortDescription:           riskConfig.ShortDescription,
		Description:                riskConfig.Description,
		Confidentiality:            riskConfig.Confidentiality,
		ConfidentialityDescription: riskConfig.ConfidentialityDescription,
		Integrity:                  riskConfig.Integrity,
		IntegrityDescription:       riskConfig.IntegrityDescription,
		Availability:               riskConfig.Availability,
		AvailabilityDescription:    riskConfig.AvailabilityDescription,
		Exploitability:             riskConfig.Exploitability,
		AttackVector:               riskConfig.AttackVector,
		Scope:                      riskConfig.Scope,
		Score:                      score,
	}
}

type Risk struct {
	RiskScore    int          `json:"riskScore" bson:"riskScore"`
	RiskCategory RiskCategory `json:"riskCategory" bson:"riskCategory"`
	RiskItems    []RiskItem   `json:"riskItems" bson:"riskItems"`
	RiskStatus   RiskStatus   `json:"riskStatus" bson:"riskStatus"`
}

func (risk *Risk) Clone() *Risk {
	items := make([]RiskItem, len(risk.RiskItems))
	copy(items, risk.RiskItems)

	return &Risk{
		RiskScore:    risk.RiskScore,
		RiskCategory: risk.RiskCategory,
		RiskItems:    items,
		RiskStatus:   risk.RiskStatus,
	}
}

type IWorloadRisk interface {
	common.Workload
	SetRisk(r *Risk)
	GetRisk() *Risk
}

type WorkloadRisk struct {
	Risk *Risk `json:"risk" bson:"risk"`
}

func (w *WorkloadRisk) GetRisk() *Risk {
	return w.Risk.Clone()
}

func (w *WorkloadRisk) SetRisk(r *Risk) {
	w.Risk = r
}

type WorkloadRiskData struct {
	Kind             string `json:"kind"`
	Name             string `json:"name"`
	Namespace        string `json:"namespace"`
	Domain           string `json:"domain"`
	IsSystemWorkload bool   `json:"isSystemWorkload"`
	Risk             *Risk  `json:"risk"`
}

func ToWorkloadRiskData(workload IWorloadRisk, isSystemWorkload bool) *WorkloadRiskData {
	return &WorkloadRiskData{
		Kind:             workload.GetKind(),
		Name:             workload.GetName(),
		Namespace:        workload.GetNamespace(),
		Domain:           workload.GetDomain(),
		IsSystemWorkload: isSystemWorkload,
		Risk:             workload.GetRisk(),
	}
}

type WorkloadRiskStatus struct {
	bongo.DocumentBase `bson:",inline"`
	Account            string     `json:"account" bson:"account"`
	Domain             string     `json:"domain" bson:"domain"`
	Namespace          string     `json:"namespace" bson:"namespace"`
	Kind               string     `json:"kind" bson:"kind"`
	Name               string     `json:"name" bson:"name"`
	RiskStatus         RiskStatus `json:"riskStatus" bson:"riskStatus"`
}

type RiskStatusesHolder struct {
	Account string
	//domain->namespace->kind->name->status
	cache map[string]map[string]map[string]map[string]RiskStatus
}

func NewRiskStatusesHolder(account string) *RiskStatusesHolder {
	cache := make(map[string]map[string]map[string]map[string]RiskStatus)

	return &RiskStatusesHolder{
		Account: account,
		cache:   cache,
	}
}

func (holder *RiskStatusesHolder) SetStatus(status WorkloadRiskStatus) {
	domain, ok := holder.cache[status.Domain]
	if !ok {
		holder.cache[status.Domain] = make(map[string]map[string]map[string]RiskStatus)
		domain = holder.cache[status.Domain]
	}

	namespace, ok := domain[status.Namespace]
	if !ok {
		domain[status.Namespace] = make(map[string]map[string]RiskStatus)
		namespace = domain[status.Namespace]
	}

	kind, ok := namespace[status.Kind]
	if !ok {
		namespace[status.Kind] = make(map[string]RiskStatus)
		kind = namespace[status.Kind]
	}

	kind[status.Name] = status.RiskStatus
}

func (holder *RiskStatusesHolder) GetStatus(domain string, namespace string, kind string, name string) (RiskStatus, bool) {
	domainMap, ok := holder.cache[domain]
	if !ok {
		return RiskStatusOpen, false
	}

	namespaceMap, ok := domainMap[namespace]
	if !ok {
		return RiskStatusOpen, false
	}

	kindMap, ok := namespaceMap[kind]
	if !ok {
		return RiskStatusOpen, false
	}

	result, ok := kindMap[name]
	if !ok {
		return RiskStatusOpen, false
	}

	return result, true
}

type WorkloadRiskDataList []*WorkloadRiskData

func (w WorkloadRiskDataList) Sanitized() interface{} {
	return w
}
