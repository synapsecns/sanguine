import { TRANSACTIONS_PATH, getChainUrl } from '@urls'
import { useState, useEffect } from 'react'
import { TOKEN_HASH_MAP } from '@constants/tokens/index'
import { AssetImage } from '@components/misc/AssetImage'
import { useSearchParams } from 'next/navigation'
import { HorizontalDivider } from '@components/misc/HorizontalDivider'
import { ChainInfo } from '@components/misc/ChainInfo'
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import { BridgeTransactionTable } from '@components/BridgeTransaction/BridgeTransactionTable'
import { useQuery } from '@apollo/client'
import { SynapseLogoSvg } from '@components/layouts/MainLayout/SynapseLogoSvg'
import { useRouter } from 'next/router'
import CopyTitle from '@components/misc/CopyTitle'
import { checksumAddress } from '@utils/checksum'
import { GET_BRIDGE_TRANSACTIONS_QUERY } from '@graphql/queries'
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

interface variableTypes {
  page: number
  addressFrom?: string,
  useMv?: boolean,
  addressTo?: string,
  tokenAddressFrom?: string | string[],
  tokenAddressTo?: string[]
  chainIDFrom?: any,
  chainIDTo?: any,
  chainId?: any
}

export default function chainId() {
  const router = useRouter()
  const { tokenAddress } = router.query
  const search = useSearchParams()
  const chainId = Number(search.get('chainId')) || 1

  const [currentTooltipIndex, setCurrentTooltipIndex] = useState(0)
  const [platform, setPlatform] = useState('ALL')
  const [transactionsArr, setTransactionsArr] = useState([])
  const [tokenChainID, setTokenChainID] = useState<any>([])
  const [variables, setVariables] = useState<variableTypes>({ page: 1 })
  const [completed, setCompleted] = useState(false)
  const [address, setAddress] = useState('')

  const [dailyStatisticDuration, SetDailyStatisticDuration] =
    useState('PAST_MONTH')
  const [dailyStatisticCumulative, SetDailyStatisticCumulative] =
    useState(false)

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
    variables,
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
    setAddress(checksumAddress(tokenAddress))
    setTokenChainID(chainId)
    setVariables({
      page: 1,
      tokenAddressFrom: [checksumAddress(tokenAddress)],
      chainIDFrom: chainId,
      useMv: true,
    })
  }, [chainId, tokenAddress])

  return (
    <StandardPageContainer title={''}>
      <a href={getChainUrl({ chainId: tokenChainID })}>
        <div className="rounded-md py-1 px-2 bg-gray-800/50 mb-2 w-fit hover:bg-gray-500/50">
          <ChainInfo
            chainId={tokenChainID}
            imgClassName="w-6 h-6 rounded-full"
            noLink={true}
          />
        </div>
      </a>
      <div className="flex items-center mb-2">
        <AssetImage
          tokenAddress={address}
          chainId={tokenChainID}
          className={`w-9 h-9 inline mr-3 rounded-lg`}
        />
        <h3 className="text-white text-5xl font-semibold">
          {TOKEN_HASH_MAP[tokenChainID]?.[address]?.symbol}{' '}
        </h3>
      </div>
      <CopyTitle title={address} />
      <HolisticStats
        noMessaging={true}
        platform={platform}
        chainID={variables?.chainId}
        // tokenAddress={address}
        loading={false}
        setPlatform={setPlatform}
        baseVariables={{
          platform,
          duration: 'ALL_TIME',
          useCache: false,
          chainID: tokenChainID,
          tokenAddress: address,
          useMv: true,
        }}
      />
      <br />
      <HorizontalDivider />
      <HorizontalDivider />
      <br /> <br />
      <p className="text-white text-2xl font-bold">Recent Transactions</p>
      <div className="h-full flex items-center mt-4">
        <button
          onClick={() =>
            setVariables({
              page: 1,
              tokenAddressFrom: [address],
              chainIDFrom: tokenChainID,
              useMv: true,
            })
          }
          className={
            'font-medium rounded-l-md px-4 py-2 border  h-fit  ' +
            (variables?.tokenAddressFrom ? selectStyle : unSelectStyle) +
            (loading ? ' pointer-events-none' : '')
          }
        >
          Outgoing
        </button>
        <button
          onClick={() =>
            setVariables({
              page: 1,
              tokenAddressTo: [address],
              chainIDTo: tokenChainID,
              useMv: true,
            })
          }
          className={
            'font-medium rounded-r-md px-4 py-2 border  h-fit ' +
            (variables?.tokenAddressTo ? selectStyle : unSelectStyle) +
            (loading ? ' pointer-events-none' : '')
          }
        >
          Incoming
        </button>
      </div>
      {loading ? (
        <div className="flex justify-center align-center w-full my-[100px]">
          <div className="w-[39px] animate-spin">
            <SynapseLogoSvg />
          </div>
        </div>
      ) : (
        <BridgeTransactionTable queryResult={transactionsArr} />
      )}
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
