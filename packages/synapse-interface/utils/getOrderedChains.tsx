import { CHAIN_ID_DISPLAY_ORDER } from '@constants/networks'

export function getOrderedChains(
  connectedChainId: number,
  chainId: number,
  possibleChains: string[] | undefined
) {
  let filteredChains = CHAIN_ID_DISPLAY_ORDER.filter(
    (id) => id !== connectedChainId && possibleChains?.includes(String(id))
  )
  console.log('filteredChains', filteredChains)

  let index = filteredChains.findIndex((e) => e === chainId)
  index = index === -1 ? 0 : index
  let numberOfChains = filteredChains.length
  let newList: number[] = []

  if (index === 0 || index === 1 || index === 2 || index == 3) {
    newList = filteredChains.slice(0, 6)
  } else if (numberOfChains - (index + 1) > 1) {
    newList = filteredChains.slice(index - 3, index + 3)
  } else if (numberOfChains - (index + 1) === 1) {
    newList = filteredChains.slice(index - 4, index + 2)
  } else if (numberOfChains - (index + 1) < 1) {
    newList = filteredChains.slice(index - 5, index + 1)
  }
  console.log('newListnewList', newList, numberOfChains, index)

  return newList
}
