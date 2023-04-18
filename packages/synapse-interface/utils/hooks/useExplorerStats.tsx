import useSWR from 'swr'
import numeral from 'numeral'
import _ from 'lodash'

import { useSwr } from '@hooks/useSwr'
import { BRIDGESYN_ANALYTICS_API, LLAMA_API_URL } from '@/constants/urls'

export function getTotalBridgeVolume(): number | undefined {
  const { data } = useSwr(`${BRIDGESYN_ANALYTICS_API}/volume/total/in`)

  let totalAmt
  if (data?.totals) {
    totalAmt = _.sum(_.values(data?.totals))
    return numeral(totalAmt).format(`$0,0`)
  } else {
    return undefined
  }
}

export function getTotalPoolVolume(): number | undefined {
  const { data } = useSwr(`${BRIDGESYN_ANALYTICS_API}/pools/volume/total`)

  let totalAmt
  if (data?.totals) {
    totalAmt = _.sum(_.values(data?.totals))
    return numeral(totalAmt).format(`$0,0`)
  } else {
    return undefined
  }
}

export function getTotalValueLocked(): number | undefined {
  const { data: tvlData } = useSwr(LLAMA_API_URL)

  if (tvlData?.currentChainTvls) {
    const currentTvlAmt = _.sum(_.values(tvlData?.currentChainTvls))
    return numeral(currentTvlAmt).format(`$0,0`)
  } else {
    return undefined
  }
}
