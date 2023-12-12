import { toHexStr } from '../toHexStr'

export const switchNetwork = async (chainId: number, provider: any) => {
  try {
    if (typeof chainId !== 'number') {
      throw new Error('Require valid chainId')
    }
    if (!provider) {
      throw new Error('Require valid provider')
    }
    const hexChainId: string = toHexStr(chainId)
    await provider.send('wallet_switchEthereumChain', [{ chainId: hexChainId }])
  } catch (error) {
    console.error('switchNetwork ', error)
  }
}
