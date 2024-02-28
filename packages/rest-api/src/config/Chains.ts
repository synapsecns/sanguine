interface Chain {
  id: number
}

const CHAINS: Record<string, Chain> = {
  ETH: { id: 1 },
  OPTIMISM: { id: 10 },
  BNB: { id: 56 },
  POLYGON: { id: 137 },
  FANTOM: { id: 250 },
  ARBITRUM: { id: 42161 },
  AVALANCHE: { id: 43114 },
  MOONRIVER: { id: 1285 },
  BOBA: { id: 288 },
  HARMONY: { id: 1666600000 },
  MOONBEAM: { id: 1284 },
  CRONOS: { id: 25 },
  METIS: { id: 1088 },
  DOGE: { id: 2000 },
  CANTO: { id: 7700 },
  KLAYTN: { id: 8217 },
  DFK: { id: 53935 },
  BASE: { id: 8453 },
  AURORA: { id: 1313161554 },
}

export default CHAINS
