// TODO: Is there one spot you can register the Chart elements?
import { useState, useEffect } from 'react'
import { Bar } from 'react-chartjs-2'
import { useQuery } from '@apollo/client'

import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js'

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend)

import { COUNT_BY_CHAIN_ID } from '@graphql/queries'
import { CHAIN_INFO_MAP } from '@constants/networks'
import { NETWORK_COLORS } from '@utils/styles/networks'
import { directColors, chartOptions } from './constants'

export function TopChains() {
  const [countByChainId, setCount] = useState([])

  const { loading, error, data } = useQuery(COUNT_BY_CHAIN_ID)

  useEffect(() => {
    if (data) {
      setCount(data.countByChainId)
    }
  }, [data])

  if (loading) {
    return 'loading'
  } else if (error) {
    return 'error'
  } else {
    let labels = countByChainId
      .map(({ chainId }) => CHAIN_INFO_MAP[chainId].chainName)
      .slice(0, 10)

    let txnCount = countByChainId.map(({ count }) => count).slice(0, 10)

    let backgroundColors = countByChainId.map(
      ({ chainId }) => directColors[NETWORK_COLORS[chainId]]
    )

    const dataset = {
      labels,
      datasets: [
        {
          label: 'Transactions',
          data: txnCount,
          borderColor: backgroundColors,
          backgroundColor: backgroundColors,
        },
      ],
    }

    return <Bar options={chartOptions} data={dataset} />
  }
}
