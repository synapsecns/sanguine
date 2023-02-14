import { ApolloClient, HttpLink, InMemoryCache } from '@apollo/client'
import { TRANSACTIONS_PATH } from '@urls'

import { Error } from '@components/Error'
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import { useRouter } from 'next/router'
import { useSearchParams } from 'next/navigation'
import { CHAIN_INFO_MAP, CHAIN_EXPLORER_URLS } from '@constants/networks'

import { GET_BRIDGE_TRANSACTIONS_QUERY, } from '@graphql/queries'
import { API_URL } from '@graphql'
import { HorizontalDivider } from "@components/misc/HorizontalDivider";
import { UniversalSearch } from "@components/pages/Home/UniversalSearch";
import { timeAgo } from "@utils/timeAgo";
import { IconAndAmount } from "@components/misc/IconAndAmount";
import { BridgeTransactionTable } from "@components/BridgeTransaction/BridgeTransactionTable";
import { ellipsizeString } from "@utils/ellipsizeString";

const link = new HttpLink({
  uri: API_URL,
  useGETForQueries: true,
})

const client = new ApolloClient({
  link: link,
  cache: new InMemoryCache(),
  fetchPolicy: 'no-cache',
  fetchOptions: {
    mode: 'no-cors',
  },
})

export default function BridgeTransaction({ queryResult }) {
  const router = useRouter()
  const search = useSearchParams()
  const { kappa } = router.query
  const chainId = Number(search.get('chainIdFrom'))

  let transaction = queryResult.bridgeTransactions[0]
  const { pending, fromInfo, toInfo } = transaction
  const { chainName: oChainName } = CHAIN_INFO_MAP[fromInfo.chainID] ?? {}
  const { chainName: dChainName } = CHAIN_INFO_MAP[toInfo.chainID] ?? {}

  const getTimeDifference = (start, end) => {
    const diff = end - start
    if (0 >= diff) {
      return '1'
    }
    return diff.toString()
  }
  let content

  if (!!transaction) {
    content = <>
      <div className="flex items-center mt-10 mb-2">
        <h3 className="text-white text-2xl font-semibold">Bridge TXID: {kappa}</h3>
      </div>
      <div className='mb-3'><a className='text-white cursor-pointer hover:underline' href={TRANSACTIONS_PATH}>← Back to search</a></div>
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
              : timeAgo({ timestamp: toInfo.time })} ago
          </h3>
        </div>
        <div className="flex gap-x-4 py-1">
          <p className="text-white text-opacity-60">Requested</p>
          <p className="text-white ">{new Date(fromInfo.time * 1000).toISOString()}</p>
        </div>
        <div className="flex gap-x-4 py-1">
          <p className="text-white text-opacity-60">Confirmed</p>
          <p className="text-white ">{new Date(toInfo.time * 1000).toISOString()}</p>
        </div>
        <div className="flex gap-x-8 py-1">
          <p className="text-white text-opacity-60">Elapsed</p>
          <p className="text-white ">≈{getTimeDifference(fromInfo.time, toInfo.time)} seconds</p>
        </div>
        <div className="flex gap-y-2 flex-col">

          <div className="flex mt-4">
            <div className='flex flex-col  w-1/2'>
              <div className="flex gap-x-10">
                <h1 className="text-white text-2xl text-opacity-60">Sent</h1>
                <IconAndAmount
                  formattedValue={fromInfo.formattedValue}
                  tokenAddress={fromInfo.tokenAddress}
                  chainId={fromInfo.chainID}
                  tokenSymbol={fromInfo.tokenSymbol}
                  iconSize="w-4 h-4"
                  textSize="text-sm"
                  styledCoin={true}
                /></div>
              <div className="flex gap-x-8 py-1 ">
                <p className="text-white text-opacity-60">Address</p>
                <p className="text-white break-all text-sm">{fromInfo.address}</p>
              </div>
              <div className="flex gap-x-12 py-1 ">
                <p className="text-white text-opacity-60">Chain</p>
                <p className="text-white text-sm">{oChainName}</p>
              </div>
              <div className="flex gap-x-7 py-1 mb-2">
                <p className="text-white text-opacity-60">TX Hash</p>
                <a target="_blank"
                  rel="noreferrer" className="text-white break-all text-sm underline" href={CHAIN_EXPLORER_URLS[fromInfo.chainID] + "/tx/" + fromInfo.hash}>{fromInfo.hash}</a>
              </div>
            </div>
            <div className='flex flex-col  w-1/2'>
              <div className="flex gap-x-8">

                <h1 className="text-white text-2xl text-opacity-60">
                  Received
                </h1>
                <IconAndAmount
                  formattedValue={toInfo.formattedValue}
                  tokenAddress={toInfo.tokenAddress}
                  chainId={toInfo.chainID}
                  tokenSymbol={toInfo.tokenSymbol}
                  iconSize="w-4 h-4"
                  textSize="text-sm"
                  styledCoin={true}
                />
              </div>
              <div className="flex gap-x-[72px] py-1">
                <p className="text-white text-opacity-60">Sent To</p>
                <p className="text-white break-all text-sm">{toInfo.address}</p>
              </div>
              <div className="flex gap-x-[86px] py-1 ">
                <p className="text-white text-opacity-60">Chain</p>
                <p className="text-white break-all text-sm">{dChainName}</p>
              </div>
              <div className="flex gap-x-[66px] py-1 ">
                <p className="text-white text-opacity-60">TX Hash</p>
                <a target="_blank"
                  rel="noreferrer" className="text-white break-all text-sm underline" href={CHAIN_EXPLORER_URLS[toInfo.chainID] + "/tx/" + toInfo.hash}>{toInfo.hash}</a>
              </div>
            </div>
          </div>
        </div>
        <br />
        <HorizontalDivider />

      </div>
    </>
  } else {
    content = (
      <Error
        text="Sorry, there was a problem with that transaction hash."
        param={kappa}
        subtitle="Unknown"
      />
    )
  }

  return <StandardPageContainer title="Synapse Analytics">{content}</StandardPageContainer>
}
export async function getServerSideProps(context) {
  const { data } = await client.query({
    query: GET_BRIDGE_TRANSACTIONS_QUERY,
    variables: {
      chainId: context.params.chainIdFrom,
      kappa: context.params.kappa,
    },
  })
  return {
    props: {
      queryResult: data
    }, // will be passed to the page component as props
  }
}
