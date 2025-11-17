export const formatSlippage = (slippage: number): string => {
  return `${slippage >= 0 ? '+' : '−'}${Math.abs(slippage).toFixed(2)}%`
}
