const Chains = [
  'Ethereum',
  'Arbitrum',
  'Avalanche',
  'Base',
  'BNB Chain',
  'Optimism',
  'Polygon',
  'DFK Chain',
  'Canto',
  'Fantom',
  'Klaytn',
  'Aurora',
  'Boba Chain',
  'Cronos',
  'Metis',
  'Moonbeam',
  'Moonriver',
]

const Tokens = [
  'USDC',
  'USDT',
  'DAI',
  'crvUSD',
  'FRAX',
  'LUSD',
  'JEWEL',
  'AVAX',
  'BTC.b',
  'AVAX',
  'LINK',
  'ETH',
  'SYN',
  'NOTE',
  'sUSD',
  'GMX',
  'MATIC',
  'MOVR',
  'gOHM',
  'H2O',
  'L2DAO',
  'NEWO',
  'PLS',
  'SDT',
  'SFI',
  'UNIDX',
  'veSOLAR',
  'VSTA',
]

const formatAmount = (amount) => {
  const MAX_DECIMALS = 4

  let [, left, right] = amount.toFixed(MAX_DECIMALS).match(/(\d+)\.?(\d*)/) ?? ['', '', '']

  for (let i = 3; i < left.length; i += 4)
    left = `${left.slice(0, left.length - i)},${left.slice(-i)}`

  return left + left.length < MAX_DECIMALS
    ? '.' + right.slice(0, MAX_DECIMALS - left.length + +(left === '0'))
    : ''
}

const randHex = () => {
  const x = () => Math.round(Math.random() * 16).toString(16)
  return `#${x() + x()}…${x() + x() + x() + x()}`
}

export const generateTx = () => {
  let originAmount =
    Math.random() < 0.5 ? Math.round(Math.random() * 10000) : Math.random()
  let destinationAmount = (originAmount * (100 - Math.random() * 5)) / 100

  const origin = {
    payload: Tokens[Math.round(Math.random() * (Tokens.length - 1))],
    chain: Chains[Math.round(Math.random() * (Chains.length - 1))],
    originAmount,
    formattedAmount: formatAmount(originAmount),
    timestamp: Date.now(),
    hash: randHex(),
  }
  const destination = {
    payload: origin.payload,
    chain: Chains[Math.round(Math.random() * (Chains.length - 1))],
    destinationAmount,
    formattedAmount: formatAmount(destinationAmount),
    timestamp: origin.timestamp + Math.round(Math.random() * 600000),
    hash: randHex(),
  }

  return { origin, destination }
}