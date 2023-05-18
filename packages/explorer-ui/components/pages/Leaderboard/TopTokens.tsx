// TODO: Is there one spot you can register the Chart elements?
import { useEffect, useState } from 'react'
import { Bar } from 'react-chartjs-2'
import { useQuery } from '@apollo/client'
import { COUNT_BY_TOKEN_ADDRESS } from '@graphql/queries'
import { TOKEN_HASH_MAP } from '@constants/tokens/basic'
import { COIN_COLORS } from '@utils/styles/coins'
import { CHAIN_INFO_MAP } from '@constants/networks'
import {
  BarElement,
  CategoryScale,
  Chart as ChartJS,
  Legend,
  LinearScale,
  Title,
  Tooltip,
} from 'chart.js'

import { chartOptions, directColors } from './constants'

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend)

export function TopTokens() {
  const [countByToken, setCount] = useState([])

  const { loading, error, data } = useQuery(COUNT_BY_TOKEN_ADDRESS)

  useEffect(() => {
    if (data) {
      setCount(data.countByTokenAddress)
    }
  }, [data])

  if (loading) {
    return 'loading'
  } else if (error) {
    return 'error'
  } else {
    const labels = countByToken
      .map(({ tokenAddress, chainId }) => {
        const t = TOKEN_HASH_MAP[chainId][tokenAddress.toLowerCase()]
        const symbol = String(t && t.symbol).trim()
        const network = CHAIN_INFO_MAP[chainId].chainName
        return `${symbol} / ${network}`
      })
      .slice(0, 10)

    const tokenTxnCount = countByToken.map(({ count }) => count).slice(0, 10)

    const backgroundColors = countByToken
      .map(({ chainId, tokenAddress }, i) => {
        const t = TOKEN_HASH_MAP[chainId][tokenAddress.toLowerCase()]
        const symbol = String(t && t.symbol).trim()
        return directColors[COIN_COLORS[symbol]]
      })
      .slice(0, 10)

    const dataset = {
      labels,
      datasets: [
        {
          label: 'Transaction Count by Token',
          data: tokenTxnCount,
          borderColor: backgroundColors,
          backgroundColor: backgroundColors,
        },
      ],
    }

    return <Bar options={chartOptions} data={dataset} />
  }
}
