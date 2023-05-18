import _ from 'lodash'

export function formatUSD(totalUsdVolumes) {
  if (totalUsdVolumes == '--') {
    return '--'
  }
  if (totalUsdVolumes > 1000000000) {
    return `${_.round(totalUsdVolumes / 1000000000, 2)}B`
  } else if (totalUsdVolumes > 100000) {
    return `${_.round(totalUsdVolumes / 1000000, 2)}M`
  } else if (totalUsdVolumes > 1000) {
    return `${_.round(totalUsdVolumes / 1000, 1)}K`
  }
  return `${_.round(totalUsdVolumes, 1)}`
}
