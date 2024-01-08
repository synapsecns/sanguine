import { CHAINS_BY_ID } from '@/constants/chains'
import { getFromTokens } from '@/utils/routeMaker/getFromTokens'
import { useState } from 'react'

const downloadData = (data) => {
  const jsonString = JSON.stringify(data, null, 2)
  const blob = new Blob([jsonString], { type: 'application/json' })
  const link = document.createElement('a')

  link.download = 'data.json'

  link.href = window.URL.createObjectURL(blob)

  document.body.appendChild(link)

  link.click()

  document.body.removeChild(link)
}

const Helper = () => {
  const [fromChainId, setFromChainId] = useState<number | null>(null)
  const [toChainId, setToChainId] = useState<number>(42161)

  const chainIds = Object.keys(CHAINS_BY_ID)

  console.log(`chianIds`, chainIds)

  const handleDownload = async () => {
    if (toChainId) {
      const routes = getFromTokens({
        fromChainId,
        toChainId,
        fromTokenRouteSymbol: null,
        toTokenRouteSymbol: null,
      })

      downloadData(routes)
    }
  }

  return (
    <div className="flex flex-col items-center justify-center h-screen space-y-1">
      <h1 className="text-white">
        tokens that can be sent from chain1 to chain2
      </h1>
      <select
        value={fromChainId || ''}
        onChange={(e) =>
          setFromChainId(e.target.value ? Number(e.target.value) : null)
        }
      >
        <option value="">Select chain...</option>
        {chainIds.map((id) => (
          <option key={id} value={id}>
            {id}
          </option>
        ))}
      </select>

      <select
        value={toChainId}
        onChange={(e) => setToChainId(Number(e.target.value))}
      >
        {chainIds.map((id) => (
          <option key={id} value={id}>
            {id}
          </option>
        ))}
      </select>
      <button className="bg-white" onClick={handleDownload}>
        Download Data
      </button>
    </div>
  )
}

export default Helper
