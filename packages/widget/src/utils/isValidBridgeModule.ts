export const isValidBridgeModule = (
  module: any
): module is 'SynapseBridge' | 'SynapseRFQ' | 'SynapseCCTP' | 'ALL' => {
  return ['SynapseBridge', 'SynapseRFQ', 'SynapseCCTP', 'ALL'].includes(module)
}
