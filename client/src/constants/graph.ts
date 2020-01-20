
export enum EdgeDirection {
  To = '1-EDGE-DIR-TO',
  Both = '2-EDGE-DIR-BOTH',
}

export const EDGE_DIRECTION_TO = EdgeDirection.To;
export const EDGE_DIRECTION_BOTH = EdgeDirection.Both;

export enum MapNodeType {
  External = 'EXTERNAL',
  Internal = 'INTERNAL',
  IpGroup = 'IP_GROUP',
}

export enum IPGroupDirection {
  Inbound = 'inbound',
  Outbound = 'outbound',
}

export const NODE_TYPE_INTERNAL = MapNodeType.Internal; // Workload
export const NODE_TYPE_EXTERNAL = MapNodeType.External; // External IP (ingress) / Host (egress)
export const NODE_TYPE_IP_GROUP = MapNodeType.IpGroup; // Group ip list of traffic with a single workload

export const NO_GROUP = 'NO_GROUP';


export const EDGE_KEY_SEPARATOR = '-->';

export const MAX_WORKLOADS_IN_WORKLOAD_LAYOUT = 20;

