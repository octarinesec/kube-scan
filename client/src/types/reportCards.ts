export enum ReportCardScoreType {
  count = 'COUNT',
  percentage = 'PERCENTAGE',
}

export interface ReportCardScore {
  scoreType: ReportCardScoreType,
  value: number,
  value2: number|null,
  change: number| null,
}

export interface InsightsReportCardInfo {
  cardType: ReportCardType,
  cardCategory: ReportCardCategory,
  severity: ReportCardSeverity,
  title: string,
  subTitle: string,
  score: ReportCardScore|null,
  summaryText: string|null,
  summaryTextParts: string[]|null,
  hasActionLink: boolean,
  actionLinkText: string|null,
}

export interface ReportCardBase {
  cardType: ReportCardType,
  cardCategory: ReportCardCategory,
  severity: ReportCardSeverity,
}


export interface InsightsReportCardAPIData extends ReportCardBase {
  percentage?: number,
  percentageChange?: number,
  percentageByType?: {[propName:string]: number},
  count?: number,
  countOutOf?: number,
  countChange?: number,
  countByType?: {[propName:string]: number},
}

export enum ReportCardCategory {
  Compliance = 'COMPLIANCE',
  ThreatPrevention = 'THREAT_PREVENTION',
  Governance = 'GOVERNANCE',
}

export enum ReportCardSeverity {
  None,
  Normal,
  Low,
  High,
}

export enum ReportCardType {
  Encryption = 'ENCRYPTION',
  EncryptionStrength = 'ENCRYPTION_STRENGTH',
  PkiStrength = 'PKI_STRENGTH',
  ContainerHardening = 'CONTAINER_HARDENING',
  PolicyTightness = 'POLICY_TIGHTNESS',
  Pci = 'PCI',
  ThreatsDetected = 'THREATS_DETECTED',
  PreventLateralMovement = 'PREVENT_LATERAL_MOVEMENT',
  NetworkPolicyViolations = 'NETWORK_POLICY_VIOLATIONS',
  Egress = 'EGRESS',
  Ingress = 'INGRESS',
  ProtectedPods = 'PROTECTED_PODS',
  InformationAccess = 'INFORMATION_ACCESS',
  TrafficVolume = 'TRAFFIC_VOLUME',
  NumberOfReplicas = 'NUMBER_OF_REPLICAS',
  K8sChanges = 'K8S_CHANGES',
}

export const cardToCategoryMapping: { [propName: string]: ReportCardCategory } = {
  [ReportCardType.Encryption]: ReportCardCategory.Compliance,
  [ReportCardType.EncryptionStrength]: ReportCardCategory.Compliance,
  [ReportCardType.PkiStrength]: ReportCardCategory.Compliance,
  [ReportCardType.ContainerHardening]: ReportCardCategory.Compliance,
  [ReportCardType.PolicyTightness]: ReportCardCategory.Compliance,
  [ReportCardType.Pci]: ReportCardCategory.Compliance,

  [ReportCardType.ThreatsDetected]: ReportCardCategory.ThreatPrevention,
  [ReportCardType.PreventLateralMovement]: ReportCardCategory.ThreatPrevention,
  [ReportCardType.NetworkPolicyViolations]: ReportCardCategory.ThreatPrevention,
  [ReportCardType.K8sChanges]: ReportCardCategory.ThreatPrevention,
  [ReportCardType.Egress]: ReportCardCategory.ThreatPrevention,
  [ReportCardType.Ingress]: ReportCardCategory.ThreatPrevention,

  [ReportCardType.ProtectedPods]: ReportCardCategory.Governance,
  [ReportCardType.InformationAccess]: ReportCardCategory.Governance,
  [ReportCardType.TrafficVolume]: ReportCardCategory.Governance,
  [ReportCardType.NumberOfReplicas]: ReportCardCategory.Governance,
};
