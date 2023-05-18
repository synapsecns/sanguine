import { ApolloClient, HttpLink, InMemoryCache } from '@apollo/client'
import { TRANSACTIONS_PATH, ACCOUNTS_PATH } from '@urls'
import { ChainInfo } from '@components/misc/ChainInfo'
import { Error } from '@components/Error'
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import { useRouter } from 'next/router'
import { useSearchParams } from 'next/navigation'
import { CHAIN_EXPLORER_URLS, BRIDGE_CONTRACTS } from '@constants/networks'
import { GET_BRIDGE_TRANSACTIONS_QUERY } from '@graphql/queries'
import { API_URL } from '@graphql'
import { HorizontalDivider } from '@components/misc/HorizontalDivider'
import { timeAgo } from '@utils/timeAgo'
import { formatDateTime } from '@utils/formatDate'
import CopyTitle from '@components/misc/CopyTitle'
import { IconAndAmount } from '@components/misc/IconAndAmount'
import { BridgeTransactionTable } from '@components/BridgeTransaction/BridgeTransactionTable'

const link = new HttpLink({
  uri: API_URL,
  useGETForQueries: true,
})

const client = new ApolloClient({
  link,
  cache: new InMemoryCache()
})

export default function BridgeTransaction({ queryResult }) {
  const router = useRouter()
  const search = useSearchParams()
  const { kappa } = router.query
  const chainId = Number(search.get('chainIdFrom'))
  const handlePending = (date) => {
    const now = new Date().getTime()
    const timeDiff = now - date * 1000
    if (timeDiff > 86400000) {
      return 'Indexing'
    } else {
      return 'Pending'
    }
  }
  const transaction = queryResult.bridgeTransactions[0]
  const { pending, fromInfo, toInfo } = transaction

  const getTimeDifference = (start, end) => {
    const diff = end - start
    if (0 >= diff) {
      return '1'
    }
    return diff.toString()
  }
  let content
  const pendingContent = handlePending(fromInfo?.time)

  if (!!transaction) {
    content = (
      <>
        <div className=" mt-5 mb-1">
          <a
            className="text-white cursor-pointer hover:underline"
            href={TRANSACTIONS_PATH}
          >
            ← Analytics
          </a>
        </div>

        <div className=" mb-2">
          <h3 className="text-white text-5xl mb-2 font-semibold">TXID</h3>
          <CopyTitle title={kappa} />
        </div>
        <br />
        <HorizontalDivider />
        {/* <UniversalSearch placeholder={`txid: ${kappa}`} /> */}
        <BridgeTransactionTable queryResult={queryResult.bridgeTransactions} />

        <HorizontalDivider />
        <div className="pb-6">
          <div className="py-6">
            <h3 className="text-white text-xl font-medium ">
              {fromInfo.time
                ? timeAgo({ timestamp: fromInfo.time })
                : timeAgo({ timestamp: toInfo?.time })}{' '}
              ago
            </h3>
          </div>
          <div className="flex gap-x-4 py-1">
            <p className="text-white text-opacity-60">Requested</p>
            <p className="text-white ">
              {formatDateTime(new Date(fromInfo.time * 1000))}
            </p>
          </div>
          <div className="flex gap-x-4 py-1">
            <p className="text-white text-opacity-60">Confirmed</p>
            <p className="text-white ">
              {toInfo
                ? formatDateTime(new Date(toInfo.time * 1000))
                : pendingContent}
            </p>
          </div>
          <div className="flex gap-x-[1.1rem] py-1">
            <p className="text-white text-opacity-60">Total Time</p>
            <p className="text-white ">
              {toInfo
                ? getTimeDifference(fromInfo.time, toInfo.time) + ' seconds'
                : pendingContent}{' '}
            </p>
          </div>
          <br />

          <div className="flex gap-y-2 flex-col">
            <HorizontalDivider />

            <div className="flex mt-4 flex-col">
              <div className="flex flex-col">
                <div className="flex gap-x-[3rem] py-1 ">
                  <p className="text-white text-opacity-60">Origin</p>
                  <ChainInfo
                    chainId={fromInfo.chainID}
                    noLink={true}
                    imgClassName="w-6 h-6 rounded-full"
                  />
                </div>
                <div className="flex gap-x-[3.4rem] py-1 ">
                  <p className="text-white text-opacity-60">From</p>
                  <a
                    target="_blank"
                    rel="noreferrer"
                    className="text-white break-all text-sm underline"
                    href={ACCOUNTS_PATH + '/' + fromInfo.address}
                  >
                    {fromInfo.address}
                  </a>
                </div>

                <div className="flex gap-x-[1.8rem] py-1">
                  <p className="text-white text-opacity-60">TX Hash</p>
                  <a
                    target="_blank"
                    rel="noreferrer"
                    className="text-white break-all text-sm underline"
                    href={
                      CHAIN_EXPLORER_URLS[fromInfo.chainID] +
                      '/tx/' +
                      fromInfo.hash
                    }
                  >
                    {fromInfo.hash}
                  </a>
                </div>
                <div className="flex gap-x-[1.7rem] py-1">
                  <p className="text-white text-opacity-60">Contract</p>
                  <a
                    target="_blank"
                    rel="noreferrer"
                    className="text-white break-all text-sm underline"
                    href={
                      CHAIN_EXPLORER_URLS[fromInfo.chainID] +
                      '/address/' +
                      BRIDGE_CONTRACTS[fromInfo.chainID]
                    }
                  >
                    Origin Bridge Contract
                  </a>
                </div>
                <div className="flex gap-x-11 mt-3">
                  <h1 className="text-white text-2xl text-opacity-60">Sent</h1>
                  <IconAndAmount
                    formattedValue={fromInfo.formattedValue}
                    tokenAddress={fromInfo.tokenAddress}
                    chainId={fromInfo.chainID}
                    tokenSymbol={fromInfo.tokenSymbol}
                    iconSize="w-4 h-4"
                    textSize="text-sm"
                    styledCoin={true}
                  />
                </div>
              </div>
              <br />

              <HorizontalDivider />

              <div className="flex  mt-8  flex-col">
                <div className="flex gap-x-2 py-1 ">
                  <p className="text-white text-opacity-60">Destination</p>
                  {toInfo ? (
                    <ChainInfo
                      chainId={toInfo.chainID}
                      noLink={true}
                      imgClassName="w-6 h-6 rounded-full"
                    />
                  ) : (
                    <ChainInfo
                      chainId={fromInfo.destinationChainID}
                      noLink={true}
                      imgClassName="w-6 h-6 rounded-full"
                    />
                  )}
                </div>
                <div className="flex gap-x-[4.5rem] py-1">
                  <p className="text-white text-opacity-60">To</p>
                  <a
                    target="_blank"
                    rel="noreferrer"
                    className="text-white break-all text-sm underline"
                    href={toInfo ? ACCOUNTS_PATH + '/' + toInfo.address : ''}
                  >
                    {toInfo ? toInfo.address : pendingContent}
                  </a>
                </div>

                <div className="flex gap-x-[1.7rem] py-1 ">
                  <p className="text-white text-opacity-60">TX Hash</p>
                  {toInfo ? (
                    <a
                      target="_blank"
                      rel="noreferrer"
                      className="text-white break-all text-sm underline"
                      href={
                        CHAIN_EXPLORER_URLS[toInfo.chainID] +
                        '/tx/' +
                        toInfo.hash
                      }
                    >
                      {toInfo.hash}
                    </a>
                  ) : (
                    <p className="text-white break-all text-sm ">
                      {pendingContent}
                    </p>
                  )}
                </div>
                <div className="flex gap-x-[1.6rem] py-1">
                  <p className="text-white text-opacity-60">Contract</p>
                  <a
                    target="_blank"
                    rel="noreferrer"
                    className="text-white break-all text-sm underline"
                    href={
                      CHAIN_EXPLORER_URLS[
                        toInfo?.chainID
                          ? toInfo.chainID
                          : fromInfo.destinationChainID
                      ] +
                      '/address/' +
                      BRIDGE_CONTRACTS[
                        toInfo?.chainID
                          ? toInfo.chainID
                          : fromInfo.destinationChainID
                      ]
                    }
                  >
                    Destination Bridge Contract
                  </a>
                </div>
                <div className="flex gap-x-8 mt-3">
                  <h1 className="text-white text-2xl text-opacity-60">
                    Received
                  </h1>

                  {toInfo ? (
                    <IconAndAmount
                      formattedValue={toInfo.formattedValue}
                      tokenAddress={toInfo.tokenAddress}
                      chainId={toInfo.chainID}
                      tokenSymbol={toInfo.tokenSymbol}
                      iconSize="w-7 h-7"
                      textSize="text-md"
                      styledCoin={true}
                    />
                  ) : null}
                </div>
              </div>
            </div>
          </div>
          <br />
          <HorizontalDivider />
        </div>
      </>
    )
  } else {
    content = (
      <Error
        text="Sorry, there was a problem with that transaction hash."
        param={kappa}
        subtitle="Unknown"
      />
    )
  }

  return <StandardPageContainer>{content}</StandardPageContainer>
}
export async function getServerSideProps(context) {
  const { data } = await client.query({
    query: GET_BRIDGE_TRANSACTIONS_QUERY,
    variables: {
      chainId: context.params.chainIdFrom,
      kappa: context.params.kappa,
      useMv: true,
    },
  })

  return {
    props: {
      queryResult: data,
    }, // will be passed to the page component as props
  }
}
