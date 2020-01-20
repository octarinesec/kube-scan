
export enum NotificationProvider {
  Webhook = 'WEBHOOK',
  Slack = 'SLACK',
}

export enum NotificationFilterType {
  All = 'All',
  Network = 'NETWORK',
  Threats = 'THREATS',
  K8sChanges = 'K8S_CHANGES',
  Anomaly = 'ANOMALY',
}

export enum NotificationType {
  Alert = 'Alert',
}

export type NotificationFilterParams = {
  sender: string,
  receiver: string,
  type: NotificationFilterType,
}

export type NotificationFilter = {
  filterParams: NotificationFilterParams,
  type: NotificationType,
}

export type NotificationProviderParams = {
  url: string,
}

export type NotificationItem = {
  filters: NotificationFilter[]
  id: string,
  provider: NotificationProvider,
  providerParams: NotificationProviderParams,
}
