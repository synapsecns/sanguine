export const getTransactionReceipt = async (
  txHash: string,
  chainId: number,
  provider: any
) => {
  try {
    const receipt = await provider.getTransactionReceipt(txHash)
    return receipt
  } catch (error) {
    console.error('[Synapse Widget] Transaction receipt error: ', error)
    return null
  }
}
