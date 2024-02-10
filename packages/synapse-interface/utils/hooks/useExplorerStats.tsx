import numeral from 'numeral'
import _ from 'lodash'

import { useSwr } from '@hooks/useSwr'
import { BRIDGESYN_ANALYTICS_API, LLAMA_API_URL } from '@/constants/urls'

export type ExplorerQueryStatsResponse = number | undefined

export const getTotalBridgeVolume = (): ExplorerQueryStatsResponse => {
  const { data } = useSwr(
    `${BRIDGESYN_ANALYTICS_API}?query=%7B%0A%20%20amountStatistic(type%3A%20TOTAL_VOLUME_USD%2C%20duration%3A%20ALL_TIME%2C%20useMv%3A%20true%2C%20platform%3AALL%2C%20useCache%3Atrue)%7Bvalue%7D%0A%7D%0A`
  )
  let value = data?.data?.amountStatistic?.value
  return value ? numeral(value).format(`$0,0`) : undefined
}

export const getTotalTxCount = (): ExplorerQueryStatsResponse => {
  const { data } = useSwr(
    `${BRIDGESYN_ANALYTICS_API}?query=%7B%0A%20%20amountStatistic(type%3A%20COUNT_TRANSACTIONS%2C%20duration%3A%20ALL_TIME%2C%20useMv%3A%20true%2C%20platform%3AALL%2C%20useCache%3Atrue)%7Bvalue%7D%0A%7D%0A`
  )
  let value = data?.data?.amountStatistic?.value
  return value ? numeral(value).format(`0,0`) : undefined
}

export const getTotalValueLocked = (): ExplorerQueryStatsResponse => {
  const { data: tvlData } = useSwr(LLAMA_API_URL)

  if (tvlData?.currentChainTvls) {
    const currentTvlAmt = _.sum(_.values(tvlData?.currentChainTvls))
    return numeral(currentTvlAmt).format(`$0,0`)
  } else {
    return undefined
  }
}
