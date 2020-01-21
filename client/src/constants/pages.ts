export enum PageName {
  AnomalyAlerts = 'ANOMALY_ALERTS',
  BlockedMessages = 'BLOCKED_MESSAGES',
  NetworkPolicy = 'NETWORK_POLICY',
  NetworkActivity = 'NETWORK_ACTIVITY',
  NetworkAlerts = 'NETWORK_ALERTS',
  KafkaActivity = 'KAFKA_ACTIVITY',
  ReportCards = 'REPORT_CARDS',
  K8sSummary = 'KUBERNETES_SUMMARY',
  K8sAlerts = 'KUBERNETES_ALERTS',
  K8sPolicies = 'KUBERNETES_POLICIES',
  K8sWorkloads = 'KUBERNETES_INVENTORY',
  K8sPolicyCreate = 'KUBERNETES_POLICY_CREATE',
  K8sPolicyDuplicate = 'KUBERNETES_POLICY_DUPLICATE',
  K8sPolicyItemView = 'KUBERNETES_POLICY_ITEM_VIEW',
  K8sPolicyItemEdit = 'KUBERNETES_POLICY_ITEM_EDIT',
  EgressPolicy = 'EGRESS_POLICY',
  ThreatAlerts = 'THREAT_ALERTS',
  RuleRecommendationsExternal = 'RuleRecommendationsExternal',
  RuleRecommendationsInternal = 'RuleRecommendationsInternal',
}

export const validPageNames = new Set(Object.values(PageName));

export type RuleRecommendationsPage = PageName.RuleRecommendationsInternal|PageName.RuleRecommendationsExternal;

export type PageWithWeeklyOffset = PageName.NetworkAlerts|PageName.K8sAlerts|PageName.ThreatAlerts|PageName.AnomalyAlerts|PageName.BlockedMessages;
