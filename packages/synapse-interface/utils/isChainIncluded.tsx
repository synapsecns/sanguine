/**
 * Checks if any of the chain IDs in `hasChains` are found within the `chainList` array.
 *
 * @param {number[]} chainList - The array of chain IDs to check against.
 * @param {number[]} hasChains - The array of chain IDs to find within `checkChains`.
 * @returns {boolean} - True if any chain ID from `hasChains` is found in `checkChains`, otherwise false.
 */

export const isChainIncluded = (chainList: number[], hasChains: number[]) => {
  return hasChains.some((chainId) => chainList.includes(chainId))
}
