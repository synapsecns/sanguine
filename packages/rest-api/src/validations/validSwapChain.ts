import { SUPPORTED_SWAP_CHAIN_IDS } from '../constants'

export const validSwapChain = (chain: number | string) => {
  return SUPPORTED_SWAP_CHAIN_IDS.includes(Number(chain))
}
