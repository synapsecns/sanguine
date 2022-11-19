export const formatTotalUsdVolumes = (totalUsdVolumes) => {
  if (totalUsdVolumes > 1000000000) {
    return `${_.round(totalUsdVolumes / 1000000000, 3)}B`
  } else if (totalUsdVolumes > 1000000) {
    return `${_.round(totalUsdVolumes / 1000000, 2)}M`
  }

  return `${_.round(totalUsdVolumes / 1000, 1)}K`
}
