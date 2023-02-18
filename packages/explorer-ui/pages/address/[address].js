import { TRANSACTIONS_PATH } from '@urls'
import { useState, useEffect } from 'react'
import { TableHeader } from '@components/TransactionTable/TableHeader'
import { TOKEN_HASH_MAP } from '@constants/tokens/basic'
import { AssetImage } from '@components/misc/AssetImage'
import { useSearchParams } from 'next/navigation'

import { HorizontalDivider } from '@components/misc/HorizontalDivider'
import { formatUSD } from '@utils/formatUSD'
import { ChainInfo } from "@components/misc/ChainInfo";

import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import { BridgeTransactionTable } from '@components/BridgeTransaction/BridgeTransactionTable'
import { useLazyQuery, useQuery } from '@apollo/client'
import { SynapseLogoSvg } from "@components/layouts/MainLayout/SynapseLogoSvg";
import { CHAIN_ID_NAMES_REVERSE } from '@constants/networks'
import { useRouter } from "next/router";

import {
  GET_BRIDGE_TRANSACTIONS_QUERY,
  DAILY_STATISTICS_BY_CHAIN,
} from '@graphql/queries'
import HolisticStats from '@components/misc/HolisticStats'
import _ from 'lodash'

const titles = {
  VOLUME: 'Volume',
  FEE: 'Fees',
  ADDRESSES: 'Addrs',
  TRANSACTIONS: 'TXs',
}
const platformTitles = {
  BRIDGE: 'Bridge',
  SWAP: 'Swap',
  MESSAGE_BUS: 'Message Bus',
}
const formatCurrency = new Intl.NumberFormat('en-US', {
  style: 'currency',
  currency: 'USD',
})

export default function address() {
  const router = useRouter()
  const { address } = router.query


  const [currentTooltipIndex, setCurrentTooltipIndex] = useState(0)
  const [platform, setPlatform] = useState("ALL");
  const [transactionsArr, setTransactionsArr] = useState([])
  const [tokenChainID, setTokenChainID] = useState([])
  const [variables, setVariables] = useState({ page: 1 })
  const [completed, setCompleted] = useState(false)
  const [walletAddress, setWalletAddress] = useState("")
  const unSelectStyle =
    'transition ease-out border-l-0 border-gray-700 border-opacity-30 text-gray-500 bg-gray-700 bg-opacity-30 hover:bg-opacity-20 '
  const selectStyle = 'text-white border-[#BE78FF] bg-synapse-radial'

  const {
    loading,
    error,
    data: dataTx,
    stopPolling,
    startPolling,
  } = useQuery(GET_BRIDGE_TRANSACTIONS_QUERY, {
    pollInterval: 5000,
    variables: variables,
    fetchPolicy: 'network-only',
    onCompleted: (data) => {
      let bridgeTransactionsTable = data.bridgeTransactions
      bridgeTransactionsTable = _.orderBy(
        bridgeTransactionsTable,
        'fromInfo.time',
        ['desc']
      )
      setTransactionsArr(bridgeTransactionsTable)

    },
  })

  useEffect(() => {
    if (!completed) {
      startPolling(10000)
    } else {
      stopPolling()
    }
    return () => {
      stopPolling()
    }
  }, [stopPolling, startPolling, completed])

  // Get initial data
  useEffect(() => {
    setWalletAddress(address)
    setVariables({ page: 1, addressFrom: address, useMv: true })
  }, [address])

  return (
    <StandardPageContainer title={"Address"}>
       <h3 className="text-white text-2xl">{walletAddress}</h3>


      {walletAddress != ""?
      <HolisticStats
        platform={platform}
        loading={false}
        setPlatform={setPlatform}
        baseVariables={{
          platform: platform,
          duration: "ALL_TIME",
          useCache: false,
          address: walletAddress,
           useMv: true
        }}
      /> : null}
      <br />
      <HorizontalDivider />

      <HorizontalDivider />
      <br /> <br />
      <p className="text-white text-2xl font-bold">Recent Transactions</p>
      <div className="h-full flex items-center mt-4">
        <button
          onClick={() => setVariables({ page: 1, addressFrom: walletAddress, useMv: true })}

          className={
            'font-medium rounded-l-md px-4 py-2 border  h-fit  ' +
            (variables?.addressFrom ? selectStyle : unSelectStyle) +
            (loading ? ' pointer-events-none' : '')
          }
        >
          Outgoing
        </button>
        <button
          onClick={() => setVariables({ page: 1,  addressTo: walletAddress, useMv: true })
          }

          className={
            'font-medium rounded-r-md px-4 py-2 border  h-fit ' +
            (variables?.addressTo ? selectStyle : unSelectStyle) +
            (loading ? ' pointer-events-none' : '')
          }
        >
          Incoming
        </button>
      </div>
      {loading ? <div className="flex justify-center align-center w-full my-[100px] animate-spin"><SynapseLogoSvg /></div> : <BridgeTransactionTable queryResult={transactionsArr} />}


      <br />
      <div className="text-center text-white my-6 ">
        <div className="mt-2 mb-14 ">

          <a
            className="text-white rounded-md px-5 py-3 text-opacity-100 transition-all ease-in hover:bg-synapse-radial border-l-0 border-gray-700 border-opacity-30 bg-gray-700 bg-opacity-30 hover:border-[#BE78FF] cursor-pointer"
            href={TRANSACTIONS_PATH}
          >
            {'Explore all transactions'}
          </a>
        </div>
      </div>
      <HorizontalDivider />
    </StandardPageContainer>
  )
}
