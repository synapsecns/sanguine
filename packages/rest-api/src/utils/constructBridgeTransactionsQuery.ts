export const constructBridgeTransactionsQuery = ({
  useMv,
  originChainId,
  txnHash,
  kappa,
}: {
  useMv: boolean
  originChainId?: string | number | null
  txnHash?: string | null
  kappa?: string | null
}) => {
  return `
    {
      bridgeTransactions(
        useMv: ${useMv}
        ${originChainId ? `chainIDFrom: ${originChainId}` : ''}
        ${txnHash ? `txnHash: "${txnHash}"` : ''}
        ${kappa ? `kappa: "${kappa}"` : ''}
      ) {
        kappa
        fromInfo {
          chainID
          address
          txnHash
          value
          USDValue
          tokenSymbol
          tokenAddress
          blockNumber
          formattedTime
        }
        toInfo {
          chainID
          address
          txnHash
          value
          USDValue
          tokenSymbol
          tokenAddress
          blockNumber
          formattedTime
        }
      }
    }
  `
}
