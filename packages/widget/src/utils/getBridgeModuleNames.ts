export const getBridgeModuleNames = (module) => {
  if (module.bridgeModuleName === 'ALL') {
    return ['SynapseRFQ', 'SynapseCCTP', 'SynapseBridge']
  }
  return [module.bridgeModuleName]
}
