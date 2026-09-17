import { BRIDGE_MODULE_NAMES } from '@/components/Maintenance/functions/isValidBridgeModule'

export const getBridgeModuleNames = (module) => {
  if (module.bridgeModuleName === 'ALL') {
    return [...BRIDGE_MODULE_NAMES]
  }
  return [module.bridgeModuleName]
}
