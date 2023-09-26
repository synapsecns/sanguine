export const getSymbol = (tokenAndChainId: string): string => {
  return tokenAndChainId.split('-')[0]
}
