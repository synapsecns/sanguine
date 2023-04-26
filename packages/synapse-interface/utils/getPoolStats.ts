import { commify } from '@ethersproject/units'
import { formatBNToString } from '@bignumber/format'

export const getPoolStats = (poolData) => {
  const { apy, totalLockedUSD } = poolData ?? {}

  let fullCompoundedApyStr
  let totalLockedUSDStr

  if (poolData) {
    try {
      if (
        0 < apy?.fullCompoundedAPY &&
        apy?.fullCompoundedAPY < Number.MAX_SAFE_INTEGER
      ) {
        fullCompoundedApyStr = apy?.fullCompoundedAPY
      }
    } catch (error) {
      console.log({ error })
    }
    try {
      if (totalLockedUSD > 0) {
        totalLockedUSDStr = commify(formatBNToString(totalLockedUSD, 18, 0))
        if (totalLockedUSDStr === '0') {
          totalLockedUSDStr = undefined
        }
      }
    } catch (e) {
      console.log({ e })
    }
  }

  return {
    apy,
    totalLockedUSD,
    fullCompoundedApyStr,
    totalLockedUSDStr,
  }
}
