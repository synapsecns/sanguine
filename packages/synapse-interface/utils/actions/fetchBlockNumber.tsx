import { fetchBlockNumber } from '@wagmi/core'

const getCurrentBlockNumber = async (chainId: number) => {
  try {
    const blockNumber = await fetchBlockNumber({
      chainId: chainId,
    })
    return blockNumber
  } catch (error) {
    console.error('getCurrentBlockNumber: ', error)
  }
}
