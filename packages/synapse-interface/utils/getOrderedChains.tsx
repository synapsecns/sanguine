import { ORDERED_CHAINS_BY_ID } from '@constants/chains'

export const getOrderedChains = (
  connectedChainId: number,
  chainId: number,
  possibleChains: string[] | undefined
) => {
  const filteredChains = ORDERED_CHAINS_BY_ID.filter(
    (id) => Number(id) !== connectedChainId && possibleChains?.includes(id)
  )

  let index = filteredChains.findIndex((e) => Number(e) === chainId)
  index = index === -1 ? 0 : index
  const numberOfChains = filteredChains.length
  let newList: number[] = []

  if (index >= 0 && index < 4) {
    newList = filteredChains.slice(0, 6).map((e) => Number(e))
  } else if (numberOfChains - (index + 1) > 1) {
    newList = filteredChains.slice(index - 3, index + 3).map((e) => Number(e))
  } else if (numberOfChains - (index + 1) === 1) {
    newList = filteredChains.slice(index - 4, index + 2).map((e) => Number(e))
  } else if (numberOfChains - (index + 1) < 1) {
    newList = filteredChains.slice(index - 5, index + 1).map((e) => Number(e))
  }

  return newList
}
