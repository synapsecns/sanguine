import { useEffect, useState } from 'react'
import { useLazyQuery } from '@apollo/client'
import { SaveIcon } from '@heroicons/react/outline'
import { GET_CSV } from '@graphql/queries'
import Tooltip from '@components/tailwind/Tooltip'

export function GetCsvButton({ address }) {
  const [getCsv, { data }] = useLazyQuery(GET_CSV, {
    fetchPolicy: 'no-cache',
  })
  const [click, setClick] = useState(1)
  const [ipfsGatewayUrl, setUrl] = useState('')

  useEffect(() => {
    if (data) {
      setUrl(data.getCsv.ipfsGatewayUrl)
      const { getCsv } = data
      window.open(getCsv.ipfsGatewayUrl)
    }
  }, [data, ipfsGatewayUrl, click])

  const onClick = () => {
    setClick(click + 1)
    getCsv({ variables: { address } })
  }

  return (
    <Tooltip
      content="Download CSV of Transactions"
      tooltipClassName="!-mt-16 !-ml-m16"
    >
      <button onClick={() => onClick()}>
        <SaveIcon
          className="w-5 h-5 text-slate-600 hover:text-slate-300"
          strokeWidth={1}
        />
      </button>
    </Tooltip>
  )
}
