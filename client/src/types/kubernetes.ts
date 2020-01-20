import { AlertAction, EnumValueMap } from './common';
import { Optional, Omit } from 'utility-types';
import { Condition } from './rules';

export enum K8sChangeCategory {
  Network = 'NETWORK',
  SecurityContext = 'SECURITY_CONTEXT',
  RBAC = 'RBAC',
  Custom = 'CUSTOM',
}

export enum K8sChangeType {
  // Creation / Deletion
  ResourceCreated = 'RESOURCE_CREATED',
  ResourceDeleted = 'RESOURE_DELETED',

  // RBAC
  RoleBindingSubjects = 'ROLE_BINDING_SUBJECTS',
  RoleRef = 'ROLE_REF',

  // Network
  ExposedByService = 'EXPOSED_BY_SERVICE',
  ExposedByNetworkPolicy = 'EXPOSED_BY_NETWORK_POLICY',
  ExposedByNetworkIngressController = 'EXPOSED_BY_INGRESS_CONTROLLER',
  EgressPolicy = 'EGRESS_POLICY',
  IngressPolicy = 'INGRESS_POLICY',
  IngressRules = 'INGRESS_RULES',
  IngressDefaultBackend = 'INGRESS_DEFAULT_BACKEND',
  Labels = 'LABELS',
  MatchExpressions = 'MATCH_EXPRESSIONS',
  PolicyTypes = 'POLICY_TYPES',
  ServiceType = 'SERVICE_TYPE',
  Selectors = 'SELECTORS',
  ServicePorts = 'SERVICE_PORTS',
  SeccompUnconfined = 'SECCOMP_UNCONFINED',
  NoNetworkPolicy = 'NO_NETWORK_POLICY',
  NoNetworkPolicyIngress = 'NO_NETWORK_POLICY_INGRESS',
  NoNetworkPolicyEgress = 'NO_NETWORK_POLICY_EGRESS',


  // Security context
  NoSeLinuxOptions = 'NO_SE_LINUX_OPTIONS',
  NoSecurityContextContainer = 'NO_SECURITY_CONTEXT_CONTAINER',
  PrivilegedContainer = 'PRIVILEGED_CONTAINER',
  PrivilegeEscalationContainer = 'PRIVILEGE_ESCALATION_CONTAINER',
  UnmaskedProcMountContainer = 'PROC_MOUNT_CONTAINER',
  RunAsGroup0 = 'RUN_AS_GROUP_0',
  RunAsUser0 = 'RUN_AS_USER_0',
  ReadOnlyFileSystemContainer = 'READ_ONLY_FILE_SYSTEM_CONTAINER',
  RunAsRootContainer = 'RUN_AS_ROOT_CONTAINER',
  ShareHostPIDContainer = 'SHARE_HOST_PID_CONTAINER',
  ShareHostIPCContainer = 'SHARE_HOST_IPC_CONTAINER',
  ShareHostNetworkContainer = 'SHARE_HOST_NETWORK_CONTAINER',
  ContainerNetRawCapAdded = 'CONTAINER_NET_RAW_CAP_ADDED',
  ContainerSysAdminCapAdded = 'CONTAINER_SYS_ADMIN_CAP_ADDED',
  ContainerNetAdminCapAdded = 'CONTAINER_NET_ADMIN_CAP_ADDED',
  SysctlUnsafeParameterModified = 'SYSCTL_UNSAFE_PARAMETER_MODIFIED',
  NoAppArmor = 'NO_APP_ARMOR',
  ExecCommand = 'EXEC_COMMAND',
  PortForwardCommand = 'PORT_FORWARD_COMMAND',

  Custom = 'CUSTOM',
}

export interface K8sChangeDetail {
  detailName: string,
  detailValue: string | object,
  cis: string | null
}

export interface K8sChangeItem {
  id: string,
  cleared: boolean,
  alertAction: AlertAction|null,
  domain: string,
  changeCategory: K8sChangeCategory,
  changeType: K8sChangeType,
  changeString: string,
  changeDetails?: K8sChangeDetail[],
  cis: string | null
  k8sResourceKind: string,
  k8sResourceName: string,
  k8sNamespace: string,
  octarineName: string | null,
  octarineAccount: string,
  originalChangeString: string,
  oldValue: string,
  newValue: string,
  time: Date,
  user: string,
  policyID?: string,
  customRuleName?: string,
}


export type K8sPolicyViolation = {
  changeCategory: K8sChangeCategory,
  changeType: K8sChangeType,
  cis: string|null
  originalChangeString: string,
  newValue: string|null,
};

export interface HasContainerLookup {
  getContainerElement: () => HTMLElement|null,
}

// export type SelectionByFilter = {
//   [FilterType.DomainGroup]: string[],
//   [FilterType.Domain]: string[],
//   [FilterType.K8sViolations]: K8sViolationsFilterType,
// };

export type K8sChangeCategoryCounts = EnumValueMap<K8sChangeCategory, number>;
export type K8sChangeCategoryActionCounts = EnumValueMap<K8sChangeCategory, EnumValueMap<AlertAction, number>>;

export type K8sAlertCounts = {
  byCategory: K8sChangeCategoryCounts,
  byCategoryAction: K8sChangeCategoryActionCounts,
  totalCount: number,
  alertBlockCount: number,
  blockCount: number,
}


export type K8sPolicyScope = {
  k8sNamespace: Set<string>,
  octarineDomain: Set<string>
  resourceKind: Set<string>
  resourceName: Set<string>
  label: { [propName:string]: string },
};

export type K8sPolicyRule = {
  id: string,
  action: AlertAction,
};

export type K8sCustomRuleMapl = {
  conditions: Condition[],
  metadata: {
    name: string,
    description: string,
  },
}

export type K8sPolicyCustomRule = {
  maplRule: K8sCustomRuleMapl,
  origin: string|null,
  enabled: boolean,
  account?: string,
  id?: string,
};


export type K8sPolicyCustomRuleMap = { [propName: string]: K8sPolicyCustomRule };

export type K8sPolicyCustomRuleUsage = {
  ruleName: string,
  action: AlertAction,
  parameters?: { [propName: string]: string },
}

export enum RuleMethod {
  EQ = 'EQ',
  NEQ = 'NEQ',
  EX = 'EX',
  NEX = 'NEX',
  IN = 'IN',
  NIN = 'NIN',
  RE = 'RE',
  NRE = 'NRE',
  GE = 'GE',
  GT = 'GT',
  LE = 'LE',
  LT = 'LT',
}

export type K8sPolicyTemplate = {
  name: string,
  description: string,
  actionForRule: { [propName: string]: AlertAction }
};

export type K8sPolicyRulesTable = {
  rulesOrder: K8sPolicyRuleInfo[],
  rulesByID: { [propName: string]: K8sPolicyRuleInfo },
  rulesByCategory: { [propName: string]: K8sPolicyRuleInfo[] },
};

export type K8sPolicyOrderChange = {
  policyID: string,
  orderID: number,
};

export type K8sPolicyRuleActionMapping = {[propName: string]: AlertAction};

export type K8sPolicyRuleInfo = {
  id: string,
  name: string,
  description: string,
  cis: string|null,
  resourceKinds: string,
  configuration: string,
  category: string,
};

export type k8sRisksWorkload = {
  kind: K8sResourceKind,
  name: string,
  domain: string,
  namespace: string,
  risk: RiskDescription,
}

export enum RiskType {
  Basic,
  Aggravation,
  Remediation,
}

export enum RiskCategory {
  Low,
  Medium,
  High,
}

export type RiskDescription = {
  riskScore: number,
  riskCategory: RiskCategory,
}

export type K8sPolicy = {
  id: string,
  account: string,
  createdBy: string,
  createDate: string,
  enabled: boolean,
  name: string,
  orderID: number,
  rules: K8sPolicyRule[],
  customRules: K8sPolicyCustomRuleUsage[],
  scope: K8sPolicyScope,
  fromTemplate: string|null,
  includeInitContainers: boolean,
};

export type K8sPolicyDraft = Omit<Optional<K8sPolicy, 'id'|'orderID'>, 'account'|'createdBy'|'createDate'> & {
  customRulesAdd: K8sPolicyCustomRuleMap,
};


export enum K8sPolicyEditType {
  Delete = 'DELETE',
  Update = 'UPDATE',
  ToggleEnabled = 'TOGGLE_ENABLED',
}

export enum K8sPolicyEditAction {
  Start = 'START',
  End = 'END',
}

export enum K8sPolicyEditMode {
  Edit = 'EDIT',
  Create = 'CREATE',
  Duplicate = 'DUPLICATE',
}


export type CIS_1_6Info = {
  compliant: boolean|null,
  securityContextSet: boolean|null,
  seccomp: boolean|null,
  networkPolicySet: boolean|null,
  ingressPolicySet: boolean|null,
  egressPolicySet: boolean|null,
}

export type WorkloadCIS_1_6Info = Pick<CIS_1_6Info, 'networkPolicySet'|'ingressPolicySet'|'egressPolicySet'>;

export type CIS_1_7Info = {
  compliant: boolean|null,
  privileged: boolean|null,
  shareHostPID: boolean|null,
  shareHostIPC: boolean|null,
  shareHostNetwork: boolean|null,
  allowPrivilegeEscalation: boolean|null,
  runAsRoot: boolean|null,
  user: number|null,
  group: number|null,
  netRawCap: boolean|null,
  sysAdminCap: boolean|null
}

export type K8sWorkloadsContainer = {
  imageName: string,
  containerName: string,
  cisInfo_1_6: CIS_1_6Info,
  cisInfo_1_7: CIS_1_7Info,
  envVar: object|null,
  securityContextSet: boolean|null,
  isInitContainer: boolean,
}

export interface K8sWorkloadsWorkload {
  key: string,
  resourceKind: K8sResourceKind,
  resourceName: string,
  domainGroup: string,
  domainMember: string,
  domain: string,
  namespace: string,
  pods: K8sWorkloadsPod[]|null,
  policyID: string|null,
  policyViolations: K8sPolicyViolation[],
  containers: K8sWorkloadsContainer[]|null,
}

export enum K8sResourceKind {
  Service = 'Service',
  NetworkPolicy = 'NetworkPolicy',
  Ingress = 'Ingress',
  ReplicaSet = 'ReplicaSet',
  Pod = 'Pod',
  Deployment = 'Deployment',
  ReplicationController = 'ReplicationController',
  DaemonSet = 'DaemonSet',
  StatefulSet = 'StatefulSet',
  Job = 'Job',
  CronJob = 'CronJob',
  RoleBinding = 'RoleBinding',
}


export interface K8sWorkloadsPod extends K8sWorkloadsWorkload {
  resourceKind: K8sResourceKind.Pod,
  pods: null,
  octarineName: string|null,
}

export type K8sWorkloadsDomainGroup = {
  name: string,
  members: K8sWorkloadsDomainMember[],
}

export type K8sWorkloadsDomainMember = {
  name: string,
  groupName: string,
  domain: string,
  workloads: K8sWorkloadsWorkload[],
}

export enum K8sWorkloadsColumnName {
  resourceKind = 'resourceKind',
  resourceName = 'resourceName',
  namespace = 'namespace',
  policyName = 'policyName',
  policyViolations = 'policyViolations',
  containerName = 'containerName',
  container = 'container',
  workload = 'workload',
  cis_1_6 = 'cis_1_6',
  cisPrivileged = 'cisPrivileged',
  cisShareHostPID = 'cisShareHostPID',
  cisShareHostIPC = 'cisShareHostIPC',
  cisShareHostNetwork =  'cisShareHostNetwork',
  cisAllowPrivilegeEscalation = 'cisAllowPrivilegeEscalation',
  cisUserGroup = 'cisUserGroup',
  cisNetRawCap = 'cisNetRawCap',
  cisRunAsRoot = 'cisRunAsRoot',
  cisSysAdminCap = 'cisSysAdminCap',
  cisSecurityContextSet = 'cisSecurityContextSet',
  cisK8sPolicySet = 'cisK8sPolicySet',
  cisK8sPolicyIngressSet = 'cisK8sPolicyIngressSet',
  cisK8sPolicyEgressSet = 'cisK8sPolicyEgressSet',
  cisSeccomp = 'cisSeccomp',
  cis_1_7 = 'cis_1_7',
  envVar = 'envVar',
  initContainer = 'initContainer',
}

export type K8sColumnGroup = K8sWorkloadsColumnName.cis_1_6|K8sWorkloadsColumnName.cis_1_7|K8sWorkloadsColumnName.workload|K8sWorkloadsColumnName.container;

export type K8sWorkloadsColumnSizes = {
  [key in K8sWorkloadsColumnName]: number
}

export type K8sWorkloadsColumnSizesBase = Omit<K8sWorkloadsColumnSizes, K8sColumnGroup>;

export type K8sWorkloadsVisibleColumns = {
  [key in K8sWorkloadsColumnName]: boolean
}
