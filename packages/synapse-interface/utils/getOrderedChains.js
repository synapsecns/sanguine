import { CHAIN_ID_DISPLAY_ORDER } from '@constants/networks'

export function getOrderedChains(chainId) {
  let index = CHAIN_ID_DISPLAY_ORDER.findIndex((e) => e === chainId)
  let numberOfChains = CHAIN_ID_DISPLAY_ORDER.length
  let newList = []

  if (index === 0 || index === 1 || index === 2 || index == 3) {
    newList = CHAIN_ID_DISPLAY_ORDER.slice(0, 6)
  } else if (numberOfChains - (index + 1) > 1) {
    newList = CHAIN_ID_DISPLAY_ORDER.slice(index - 3, index + 3)
  } else if (numberOfChains - (index + 1) === 1) {
    newList = CHAIN_ID_DISPLAY_ORDER.slice(index - 4, index + 2)
  } else if (numberOfChains - (index + 1) < 1) {
    newList = CHAIN_ID_DISPLAY_ORDER.slice(index - 5, index + 1)
  }

  return newList
}
