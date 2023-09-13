export const getTokenAndChainId = (tokenAndChainId: string) => {
  const [symbol, chainId] = tokenAndChainId.split('-')

  return { symbol, chainId: Number(chainId) }
}
