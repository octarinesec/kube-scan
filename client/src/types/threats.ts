import { Entity } from 'parsing/common';
import { AlertAction, FetchStatus, SocketInfo } from 'types/common';
import { EntityType } from 'types/activity';
import { ItemsGroup } from 'utils/grouping';


export interface AlertNetworkNode {
  entity: Entity,
  socketInfo: SocketInfo | null,
}

export enum RuleGroup {
  CommandInjection = 'COMMAND_INJECTION',
  DataExfiltration = 'DATA_EXFILTRATION',
  MaliciousTools = 'MALICIOUS_TOOLS',
  ApplicationVulnerabilities = 'APPLICATION_VULNERABILITIES',
  DenialOfService = 'DENIAL_OF_SERVICE',
  Other = 'Other',
}

export enum RuleCategory {
  LinuxCommands = 'Linux commands',
  ShellCode = 'Shellcode',
  SqlInjection = 'SQL Injection',
  InformationDisclosure = 'Information disclosure',
  ApplicationExploits = 'Application exploits',
  ExploitationTools = 'Exploitation tools',
  LocalFileAccess = 'Local file access',
  DenialOfService = 'Denial of Service',
  Scanners = 'Scanners',
  Botnets = 'Botnets',
  P2P = 'P2P',
  ProtocolTunneling = 'Protocol tunneling',
}

export type ThreatRuleItem = {
  id: string;
  account: string;
  action: AlertAction;
  category: string;
  ruleGroup: RuleGroup;
  sender: Entity;
  receiver: Entity;
}

export type RuleInfo = {
  description: string;
  id: number;
  category: string;
  entry: string;
  severity: number;
}

export interface ThreatAlertItem {
  id?: string;
  ruleInfo: RuleInfo;
  sender: AlertNetworkNode;
  receiver: AlertNetworkNode;
  firstSeen: string;
  lastSeen: string;
  count: number;
  matchData: string|null,
  matchStart: number|null,
  matchEnd: number|null,
}

export type ThreatAlertGroup = ItemsGroup<ThreatAlertItem, 'subItems'>;

export interface ThreatAlertGroupTitle {
  isCollapsed: boolean,
  isGroup: true,
  groupKey: string,
  groupSender?: Entity,
  groupCategory?: string,
  groupSeverity?: number,
}

export type ThreatAlertCategoryGroup = ItemsGroup<ThreatAlertGroup, 'subItems'>;

export type FlatListItem = ThreatAlertGroup|ThreatAlertGroupTitle;

export type IPDetails = object;
export type DomainDetails = object;

export type ThreatAlertNodeDetails = {
  type: EntityType.Ip | EntityType.Hostname,
  details: IPDetails | DomainDetails,
}

export type ThreatAlertDetails = {
  threatEntryID: string,
  threatDetailsStatus: FetchStatus,
  threatDetails?: any,
  senderDetailsStatus: FetchStatus,
  senderDetails?: ThreatAlertNodeDetails,
  receiverDetailsStatus: FetchStatus,
  receiverDetails?: ThreatAlertNodeDetails,
}


export enum ThreatGroupBy {
  None = 'NONE',
  Sender = 'SENDER',
  Severity = 'SEVERITY',
  Category = 'CATEGORY',
}

export const DEFAULT_THREAT_ALERTS_GROUP_BY = ThreatGroupBy.None;


