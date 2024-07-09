import { toHexStr } from '@/utils/toHexStr'
import { CHAINS_BY_ID } from '@/constants/chains'

export const switchNetwork = async (chainId: number, provider: any) => {
  try {
    if (typeof chainId !== 'number') {
      throw new Error('Require ChainId')
    }
    if (!provider) {
      throw new Error('Require Provider')
    }
    const hexChainId: string = toHexStr(chainId)

    await provider.send('wallet_addEthereumChain', [
      { chainId: hexChainId, chainName: '' },
    ])
    await provider.send('wallet_switchEthereumChain', [{ chainId: hexChainId }])
  } catch (error) {
    console.error('[Synapse Widget] Error switching networks: ', error)
  }
}
