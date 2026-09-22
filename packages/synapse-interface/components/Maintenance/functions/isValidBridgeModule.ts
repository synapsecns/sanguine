export const BRIDGE_MODULE_NAMES = [
  'SynapseBridge',
  'SynapseRFQ',
  'SynapseCCTP',
  'SYN',
] as const

export type BridgeModuleName = (typeof BRIDGE_MODULE_NAMES)[number] | 'ALL'

export const isValidBridgeModule = (module: any): module is BridgeModuleName =>
  module === 'ALL' || BRIDGE_MODULE_NAMES.includes(module)
