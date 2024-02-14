import { Address, getPublicClient } from '@wagmi/core'

export const getGasEstimate = async (
  chainId: number,
  account: Address,
  to: Address,
  data: any,
  value?: bigint
) => {
  const publicClient = getPublicClient()

  const gasEstimate = await publicClient.estimateGas({
    value,
    to,
    account,
    data,
    chainId,
  })

  return gasEstimate
}
