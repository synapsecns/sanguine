import _ from 'lodash'
import { useState, useEffect } from 'react'
import { useRouter } from 'next/router'
import { useQuery } from '@apollo/client'
import { GET_BRIDGE_TRANSACTIONS_QUERY } from '@graphql/queries'
import { CopyTitle } from '@components/misc/CopyTitle'
import { HolisticStats } from '@components/misc/HolisticStats'
import { HorizontalDivider } from '@components/misc/HorizontalDivider'
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import { BridgeTransactionTable } from '@components/BridgeTransaction/BridgeTransactionTable'
import { SynapseLogoSvg } from '@components/layouts/MainLayout/SynapseLogoSvg'
import { checksumAddress } from '@utils/checksum'
import { TRANSACTIONS_PATH } from '@urls'

const truncateAddress = (addr: string) => {
  return `${addr.slice(0, 6)}...${addr.slice(-4)}`
}

interface variableTypes {
  page: number
  addressFrom?: string
  useMv?: boolean
  addressTo?: string
}

export const AddressPage = () => {
  const router = useRouter()
  const { address } = router.query

  const [platform, setPlatform] = useState('ALL')
  const [transactionsArr, setTransactionsArr] = useState([])
  const [variables, setVariables] = useState<variableTypes>({ page: 1 })
  const [walletAddress, setWalletAddress] = useState('')
  const unSelectStyle =
    'transition ease-out border-l-0 border-gray-700 border-opacity-30 text-gray-500 bg-gray-700 bg-opacity-30 hover:bg-opacity-20 '
  const selectStyle = 'text-white border-[#BE78FF] bg-synapse-radial'

  const { loading, stopPolling, startPolling } = useQuery(
    GET_BRIDGE_TRANSACTIONS_QUERY,
    {
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
    }
  )

  useEffect(() => {
    startPolling(10000)
    return () => {
      stopPolling()
    }
  }, [stopPolling, startPolling])

  // Get initial data
  useEffect(() => {
    if (typeof address === 'string') {
      setWalletAddress(checksumAddress(address))
      setVariables({
        page: 1,
        addressFrom: checksumAddress(address),
        useMv: true,
      })
    }
  }, [address])

  return (
    <StandardPageContainer title={'Address'}>
      <CopyTitle title={walletAddress} className="hidden sm:block" />
      <CopyTitle
        title={truncateAddress(walletAddress)}
        className="block sm:hidden"
      />
      {walletAddress !== '' ? (
        <HolisticStats
          platform={platform}
          loading={false}
          setPlatform={setPlatform}
          baseVariables={{
            platform,
            duration: 'ALL_TIME',
            useCache: false,
            address: walletAddress,
            useMv: true,
          }}
        />
      ) : null}
      <br />
      <HorizontalDivider />
      <HorizontalDivider />
      <br /> <br />
      <p className="text-2xl font-bold text-white">Recent Transactions</p>
      <div className="flex items-center h-full mt-4">
        <button
          onClick={() =>
            setVariables({ page: 1, addressFrom: walletAddress, useMv: true })
          }
          className={
            'font-medium rounded-l-md px-4 py-2 border  h-fit  ' +
            (variables?.addressFrom ? selectStyle : unSelectStyle) +
            (loading ? ' pointer-events-none' : '')
          }
        >
          Outgoing
        </button>
        <button
          onClick={() =>
            setVariables({ page: 1, addressTo: walletAddress, useMv: true })
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
      <div className="my-6 text-center text-white ">
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

const Address = () => {
  return <AddressPage />
}

export default Address
