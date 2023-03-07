import _ from 'lodash'
import { gql } from '@apollo/client'

import { SUSHISWAP_FEE_RATE } from '@constants/fees'
import { ChainId } from '@constants/networks'

import { useQuery } from '@graphql'
import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'




const PAIR_DAYS_DATAS_QUERY = gql`
  query pairDayDatasQuery(
    $first: Int = 1000
    $date: Int = 0
    $pairs: [Bytes]!
  ) {
    pairDayDatas(
      first: $first
      orderBy: date
      orderDirection: desc
      where: { pair_in: $pairs, date_gt: $date }
    ) {
      date
      reserveUSD
      volumeToken0
      volumeToken1
      volumeUSD
    }
  }
`

export function useExternalPoolQuery( token ) {
  const skip = token.poolType != "EXTERNAL_LP"
  const { chainId } = useActiveWeb3React()

  const { data, error, loading } = useQuery(PAIR_DAYS_DATAS_QUERY, {
    variables: {
      pairs: [token.addresses[chainId]]
    },
    context: {
      clientName: "exchange",
    },
    skip: skip || chainId != ChainId.ETH
  })

  if (skip || !data) {
    return {
      dailyApr: 0,
      yearlyApr: 0,
      yearlyCompoundedApy: 0,
    }
  } else {
    const dataArr = data?.pairDayDatas?.slice(1, 8) // ignores first day bc incomplete
    const aprObj = calcAprBasedOnWeeklyAvg(dataArr)

    return aprObj
  }
}



function calcAprBasedOnWeeklyAvg(dayData) {
  const dataByDay               = dayData?.map(dayObj => calcIndvDayData(dayObj))
  const dailyFeesFrac           = _.meanBy(dataByDay, i => i.dailyFeesFrac) ?? 0

  return {
    dailyApr:            100 * dailyFeesFrac,
    yearlyApr:           100 * dailyFeesFrac * 365,
    yearlyCompoundedApy: 100 * dailyFeesFrac * 365,
  }
}

function calcIndvDayData(dayInfo) {
  const {
    volumeUSD,
    reserveUSD,
  } = convertDataToNumbers(dayInfo)


  const feesGeneratedUSD = SUSHISWAP_FEE_RATE * volumeUSD
  const turnoverFrac     = volumeUSD / reserveUSD
  const turnoverFeesFrac = feesGeneratedUSD / reserveUSD

  const dailyFeesFrac = turnoverFeesFrac


  const resultObj = {
    feesGeneratedUSD,
    turnoverFrac,
    turnoverFeesFrac,
    dailyFeesFrac,
  }

  return resultObj
}



function convertDataToNumbers(i) {
  return {
    reserveUSD:      Number(i.reserveUSD),
    volumeUSD:       Number(i.volumeUSD),
  }
}
