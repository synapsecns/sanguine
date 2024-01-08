import { ExplorerLinks, ExplorerNames } from './constants'

export const getTxBlockExplorerLink = (chainId: number, txHash: string) => {
  const blockExplorer = ExplorerLinks[chainId]

  if (blockExplorer && txHash) {
    const explorerUrl = `${blockExplorer}/tx/${txHash}`
    const explorerName = ExplorerNames[chainId]

    return [explorerUrl, explorerName]
  }

  console.error('getTxBlockExplorerLink: ChainID or Transaction Hash missing')
  return [null, null]
}
