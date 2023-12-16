//@dev TO-DO: Remove mapping if we create Chains constants file
const ExplorerLinks = {
  1: 'https://etherscan.com',
  42161: 'https://arbiscan.io',
  56: 'https://bscscan.com',
  43114: 'https://avascan.info/blockchain/c/',
  7700: 'https://tuber.build/',
  10: 'https://optimistic.etherscan.io',
  137: 'https://polygonscan.com',
  53935: 'https://subnets.avax.network/defi-kingdoms',
  8217: 'https://scope.klaytn.com',
  250: 'https://ftmscan.com',
  25: 'https://cronoscan.com',
  288: 'https://bobascan.com',
  1088: 'https://andromeda-explorer.metis.io',
  1313161554: 'https://explorer.mainnet.aurora.dev',
  1666600000: 'https://explorer.harmony.one',
  1284: 'https://moonbeam.moonscan.io',
  1285: 'https://moonriver.moonscan.io',
  2000: 'https://explorer.dogechain.dog',
  8453: 'https://basescan.org',
}

export const getTxBlockExplorerLink = (chainId: number, txHash: string) => {
  const blockExplorer = ExplorerLinks[chainId]

  if (blockExplorer && txHash) {
    return `${blockExplorer}/tx/${txHash}`
  }

  console.error('ChainID or Transaction Hash missing')
}
