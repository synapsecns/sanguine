import { CHAIN_INFO_MAP } from '@constants/networks'

const nameToChainIds = {}
Object.keys(CHAIN_INFO_MAP).forEach((chainId) => {
  const name = CHAIN_INFO_MAP[chainId].chainName
  return (nameToChainIds[name] = Number(chainId))
})

const suggestions = Object.keys(CHAIN_INFO_MAP)
  .map((chainId) => CHAIN_INFO_MAP[chainId].chainName)
  .filter(Boolean)
  .filter((chainName) => chainName !== 'Terra')
  .sort()

export { nameToChainIds, suggestions }
