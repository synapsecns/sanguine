import { useQuery } from '@apollo/client'
import numeral from 'numeral'

import { BRIDGE_AMOUNT_STATISTIC } from '@graphql/queries'

export function GetStatistic({ address, type, prefix, duration = 'ALL_TIME' }) {
  const { data } = useQuery(BRIDGE_AMOUNT_STATISTIC, {
    variables: {
      type,
      duration,
      address,
    },
  })

  let value = data?.bridgeAmountStatistic?.USDValue

  if (value) {
    return formatValue(value, prefix)
  } else {
    return <div className="h-8 bg-slate-500 animate-pulse" />
  }
}

function formatValue(value, prefix = '') {
  let formattedValue = numeral(value).format(`${prefix}0,0`)

  return Number(value) !== 0 ? formattedValue : '--'
}
