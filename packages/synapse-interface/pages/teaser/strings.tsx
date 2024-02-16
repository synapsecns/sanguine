const Chains = ['Ethereum', 'Arbitrum', 'Avalanche', 'Base', 'BNB Chain', 'Optimism', 'Polygon', 'DFK Chain', 'Canto', 'Fantom', 'Klaytn', 'Aurora', 'Boba Chain', 'Cronos', 'Metis', 'Dogechain', 'Moonbeam', 'Moonriver']

const Tokens = ['USDC', 'USDT', 'DAI', 'nUSD', 'crvUSD', 'FRAX', 'LUSD', 'JEWEL', 'AVAX', 'BTC.b', 'AVAX', 'WJEWEL', 'synJEWEL', 'LINK', 'ETH', 'nETH', 'SYN', 'NOTE', 'sUSD', 'GMX', 'MATIC', 'MOVR', 'WMOVR', 'WKLAY', 'WMATIC', 'xJEWEL', 'synFRAX', 'DOG', 'gOHM', 'H2O', 'JUMP', 'L2DAI', 'NEWO', 'PLS', 'SDT', 'SFI', 'UNIDX', 'veSOLAR', 'VSTA']

export const generateTx = () => {
  let amount = Math.round(Math.random() * 10000)
  let amountStr = amount.toString()

  if (Math.random() < .5) {
    amountStr = `${amountStr.slice(0, 1)}.${amountStr.slice(2)}`
  } else for (let i = 3; i < amountStr.length; i += 4)
    amountStr = `${amountStr.slice(0, amountStr.length - i)},${amountStr.slice(-i)}`

  if (amountStr.match(/\./)) {
    if (Math.random() < .5)
      amountStr = amountStr.replace(/\d/, '0')

    while (amountStr.length < (+amountStr < 1 ? 6 : 5))
      amountStr += '0'
  }

  const token = Tokens[Math.round(Math.random() * (Tokens.length - 1))]
  const chain = Chains[Math.round(Math.random() * (Chains.length - 1))]


  const origin = {
    payload: Tokens[Math.round(Math.random() * (Tokens.length - 1))],
    chain: Chains[Math.round(Math.random() * (Chains.length - 1))],
    formattedAmount: amountStr,
    timestamp: Date.now()
  }
  const destination = {
    payload: origin.payload,
    chain: Chains[Math.round(Math.random() * (Chains.length - 1))],
    formattedAmount: amountStr,
    timestamp: origin.timestamp + Math.round(Math.random() * 600000)
  }

  const description = `${origin.formattedAmount} ${origin.payload} to ${destination.chain}`

  return { origin, destination }
}
