import { ApolloClient, HttpLink, InMemoryCache } from '@apollo/client'
import { TRANSACTIONS_PATH, ACCOUNTS_PATH } from '@urls'
import { ChainInfo } from "@components/misc/ChainInfo";
import { Error } from '@components/Error'
import { StandardPageContainer } from '@components/layouts/StandardPageContainer'
import { useRouter } from 'next/router'
import { useSearchParams } from 'next/navigation'
import { CHAIN_EXPLORER_URLS, BRIDGE_CONTRACTS } from '@constants/networks'
import { GET_BRIDGE_TRANSACTIONS_QUERY, } from '@graphql/queries'
import { API_URL } from '@graphql'
import { HorizontalDivider } from "@components/misc/HorizontalDivider";
import { timeAgo } from "@utils/timeAgo";
import { formatDateTime } from "@utils/formatDate";
import CopyTitle from '@components/misc/CopyTitle';

import { IconAndAmount } from "@components/misc/IconAndAmount";
import { BridgeTransactionTable } from "@components/BridgeTransaction/BridgeTransactionTable";

const link = new HttpLink({
  uri: API_URL,
  useGETForQueries: true,
})

const client = new ApolloClient({
  link: link,
  cache: new InMemoryCache(),
  // @ts-expect-error TS(2345): Argument of type '{ link: HttpLink; cache: InMemor... Remove this comment to see the full error message
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
  const handlePending = (date) => {
    let now = new Date().getTime()
    let timeDiff = now - date * 1000
    if (timeDiff > 86400000) {
      return "Indexing"
    } else {
      return "Pending"
    }

  }
  let transaction = queryResult.bridgeTransactions[0]
  const { pending, fromInfo, toInfo } = transaction


  const getTimeDifference = (start, end) => {
    const diff = end - start
    if (0 >= diff) {
      return '1'
    }
    return diff.toString()
  }
  let content
  let pendingContent = handlePending(fromInfo?.time)

  if (!!transaction) {
    content = <>
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className=' mt-5 mb-1'><a className='text-white cursor-pointer hover:underline' href={TRANSACTIONS_PATH}>← Analytics</a></div>

      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className=" mb-2">
        // @ts-expect-error TS(2304): Cannot find name 'h3'.
        <h3 className="text-white text-5xl mb-2 font-semibold">TXID</h3>
        // @ts-expect-error TS(2304): Cannot find name 'title'.
        <CopyTitle title={kappa} />
      </div>
      // @ts-expect-error TS(2304): Cannot find name 'br'.
      <br />
      // @ts-expect-error TS(2749): 'HorizontalDivider' refers to a value, but is bein... Remove this comment to see the full error message
      <HorizontalDivider />
      {/* <UniversalSearch placeholder={`txid: ${kappa}`} /> */}
      // @ts-expect-error TS(2749): 'BridgeTransactionTable' refers to a value, but is... Remove this comment to see the full error message
      <BridgeTransactionTable queryResult={queryResult.bridgeTransactions} />

      // @ts-expect-error TS(2362): The left-hand side of an arithmetic operation must... Remove this comment to see the full error message
      <HorizontalDivider />
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className="pb-6">
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className="py-6">
          // @ts-expect-error TS(2304): Cannot find name 'h3'.
          <h3 className="text-white text-xl font-medium ">
            {fromInfo.time
              ? timeAgo({ timestamp: fromInfo.time })
              // @ts-expect-error TS(2304): Cannot find name 'ago'.
              : timeAgo({ timestamp: toInfo?.time })} ago
          </h3>
        </div>
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className="flex gap-x-4 py-1">
          // @ts-expect-error TS(2304): Cannot find name 'p'.
          <p className="text-white text-opacity-60">Requested</p>
          // @ts-expect-error TS(2304): Cannot find name 'p'.
          <p className="text-white ">{formatDateTime(new Date(fromInfo.time * 1000))}</p>
        </div>
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className="flex gap-x-4 py-1">
          // @ts-expect-error TS(2304): Cannot find name 'p'.
          <p className="text-white text-opacity-60">Confirmed</p>
          // @ts-expect-error TS(2304): Cannot find name 'p'.
          <p className="text-white ">{toInfo ? formatDateTime(new Date(toInfo.time * 1000)) : pendingContent}</p>
        </div>
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className="flex gap-x-[1.1rem] py-1">
          // @ts-expect-error TS(2304): Cannot find name 'p'.
          <p className="text-white text-opacity-60">Total Time</p>
          // @ts-expect-error TS(2304): Cannot find name 'p'.
          <p className="text-white ">{toInfo ? getTimeDifference(fromInfo.time, toInfo.time) + " seconds" : pendingContent} </p>
        </div>
        // @ts-expect-error TS(2304): Cannot find name 'br'.
        <br />

        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className="flex gap-y-2 flex-col">
          // @ts-expect-error TS(2749): 'HorizontalDivider' refers to a value, but is bein... Remove this comment to see the full error message
          <HorizontalDivider />

          // @ts-expect-error TS(2304): Cannot find name 'div'.
          <div className="flex mt-4 flex-col">
            // @ts-expect-error TS(2304): Cannot find name 'div'.
            <div className='flex flex-col'>
              // @ts-expect-error TS(2304): Cannot find name 'div'.
              <div className="flex gap-x-[3rem] py-1 ">
                // @ts-expect-error TS(2304): Cannot find name 'p'.
                <p className="text-white text-opacity-60">Origin</p>
                <ChainInfo
                  // @ts-expect-error TS(2304): Cannot find name 'chainId'.
                  chainId={fromInfo.chainID}
                  // @ts-expect-error TS(2304): Cannot find name 'noLink'.
                  noLink={true}
                  // @ts-expect-error TS(2304): Cannot find name 'imgClassName'.
                  imgClassName="w-6 h-6 rounded-full"
                />
              </div>
              // @ts-expect-error TS(2304): Cannot find name 'div'.
              <div className="flex gap-x-[3.4rem] py-1 ">
                // @ts-expect-error TS(2304): Cannot find name 'p'.
                <p className="text-white text-opacity-60">From</p>
                // @ts-expect-error TS(2304): Cannot find name 'a'.
                <a target="_blank"
                  // @ts-expect-error TS(2304): Cannot find name 'rel'.
                  rel="noreferrer" className="text-white break-all text-sm underline" href={ACCOUNTS_PATH + "/" + fromInfo.address}>{fromInfo.address}
                </a>
              </div>

              // @ts-expect-error TS(2304): Cannot find name 'div'.
              <div className="flex gap-x-[1.8rem] py-1">
                // @ts-expect-error TS(2304): Cannot find name 'p'.
                <p className="text-white text-opacity-60">TX Hash</p>
                // @ts-expect-error TS(2304): Cannot find name 'a'.
                <a target="_blank"
                  // @ts-expect-error TS(2304): Cannot find name 'rel'.
                  rel="noreferrer" className="text-white break-all text-sm underline" href={CHAIN_EXPLORER_URLS[fromInfo.chainID] + "/tx/" + fromInfo.hash}>{fromInfo.hash}
                </a>
              </div>
              // @ts-expect-error TS(2304): Cannot find name 'div'.
              <div className="flex gap-x-[1.7rem] py-1">
                // @ts-expect-error TS(2304): Cannot find name 'p'.
                <p className="text-white text-opacity-60">Contract</p>
                // @ts-expect-error TS(2304): Cannot find name 'a'.
                <a target="_blank"
                  // @ts-expect-error TS(2304): Cannot find name 'rel'.
                  rel="noreferrer" className="text-white break-all text-sm underline" href={CHAIN_EXPLORER_URLS[fromInfo.chainID] + "/address/" + BRIDGE_CONTRACTS[fromInfo.chainID]}>Origin Bridge Contract
                </a>
              </div>
              // @ts-expect-error TS(2304): Cannot find name 'div'.
              <div className="flex gap-x-11 mt-3">
                // @ts-expect-error TS(2304): Cannot find name 'h1'.
                <h1 className="text-white text-2xl text-opacity-60">Sent</h1>
                <IconAndAmount
                  // @ts-expect-error TS(2304): Cannot find name 'formattedValue'.
                  formattedValue={fromInfo.formattedValue}
                  // @ts-expect-error TS(2304): Cannot find name 'tokenAddress'.
                  tokenAddress={fromInfo.tokenAddress}
                  // @ts-expect-error TS(2304): Cannot find name 'chainId'.
                  chainId={fromInfo.chainID}
                  // @ts-expect-error TS(2304): Cannot find name 'tokenSymbol'.
                  tokenSymbol={fromInfo.tokenSymbol}
                  // @ts-expect-error TS(2304): Cannot find name 'iconSize'.
                  iconSize="w-4 h-4"
                  // @ts-expect-error TS(2304): Cannot find name 'textSize'.
                  textSize="text-sm"
                  // @ts-expect-error TS(2304): Cannot find name 'styledCoin'.
                  styledCoin={true}
                /></div>
            </div>
            // @ts-expect-error TS(2304): Cannot find name 'br'.
            <br />

            // @ts-expect-error TS(2749): 'HorizontalDivider' refers to a value, but is bein... Remove this comment to see the full error message
            <HorizontalDivider />

            // @ts-expect-error TS(2304): Cannot find name 'div'.
            <div className='flex  mt-8  flex-col'>
              // @ts-expect-error TS(2304): Cannot find name 'div'.
              <div className="flex gap-x-2 py-1 ">
                // @ts-expect-error TS(2304): Cannot find name 'p'.
                <p className="text-white text-opacity-60">Destination</p>
                // @ts-expect-error TS(2304): Cannot find name 'toInfo'.
                {toInfo ?
                  // @ts-expect-error TS(2749): 'ChainInfo' refers to a value, but is being used a... Remove this comment to see the full error message
                  <ChainInfo
                    // @ts-expect-error TS(2304): Cannot find name 'chainId'.
                    chainId={toInfo.chainID}
                    // @ts-expect-error TS(2304): Cannot find name 'noLink'.
                    noLink={true}
                    // @ts-expect-error TS(2304): Cannot find name 'imgClassName'.
                    imgClassName="w-6 h-6 rounded-full"
                  // @ts-expect-error TS(2749): 'ChainInfo' refers to a value, but is being used a... Remove this comment to see the full error message
                  /> : <ChainInfo
                    // @ts-expect-error TS(2304): Cannot find name 'chainId'.
                    chainId={fromInfo.destinationChainID}
                    // @ts-expect-error TS(2304): Cannot find name 'noLink'.
                    noLink={true}
                    // @ts-expect-error TS(2304): Cannot find name 'imgClassName'.
                    imgClassName="w-6 h-6 rounded-full"
                  />}
              </div>
              // @ts-expect-error TS(2304): Cannot find name 'div'.
              <div className="flex gap-x-[4.5rem] py-1">
                // @ts-expect-error TS(2304): Cannot find name 'p'.
                <p className="text-white text-opacity-60">To</p>
                // @ts-expect-error TS(2304): Cannot find name 'a'.
                <a target="_blank"
                  // @ts-expect-error TS(2304): Cannot find name 'rel'.
                  rel="noreferrer" className="text-white break-all text-sm underline" href={toInfo ? ACCOUNTS_PATH + "/" + toInfo.address : ""}>{toInfo ? toInfo.address : pendingContent}
                </a>
              </div>

              // @ts-expect-error TS(2304): Cannot find name 'div'.
              <div className="flex gap-x-[1.7rem] py-1 ">
                // @ts-expect-error TS(2304): Cannot find name 'p'.
                <p className="text-white text-opacity-60">TX Hash</p>
                // @ts-expect-error TS(2304): Cannot find name 'toInfo'.
                {toInfo ?
                  // @ts-expect-error TS(2304): Cannot find name 'a'.
                  <a target="_blank"
                    // @ts-expect-error TS(2304): Cannot find name 'rel'.
                    rel="noreferrer" className="text-white break-all text-sm underline" href={CHAIN_EXPLORER_URLS[toInfo.chainID] + "/tx/" + toInfo.hash}>{toInfo.hash}</a> : <p className="text-white break-all text-sm ">{pendingContent}</p>}
              </div>
              // @ts-expect-error TS(2304): Cannot find name 'div'.
              <div className="flex gap-x-[1.6rem] py-1">
                // @ts-expect-error TS(2304): Cannot find name 'p'.
                <p className="text-white text-opacity-60">Contract</p>
                // @ts-expect-error TS(2304): Cannot find name 'a'.
                <a target="_blank"
                  // @ts-expect-error TS(2304): Cannot find name 'rel'.
                  rel="noreferrer" className="text-white break-all text-sm underline" href={CHAIN_EXPLORER_URLS[toInfo?.chainID ? toInfo.chainID : fromInfo.destinationChainID] + "/address/" + BRIDGE_CONTRACTS[toInfo?.chainID ? toInfo.chainID : fromInfo.destinationChainID]}>Destination Bridge Contract
                </a>
              </div>
              // @ts-expect-error TS(2304): Cannot find name 'div'.
              <div className="flex gap-x-8 mt-3">
                // @ts-expect-error TS(2304): Cannot find name 'h1'.
                <h1 className="text-white text-2xl text-opacity-60">
                  // @ts-expect-error TS(2304): Cannot find name 'Received'.
                  Received
                </h1>

                // @ts-expect-error TS(2304): Cannot find name 'toInfo'.
                {toInfo ?
                  // @ts-expect-error TS(2749): 'IconAndAmount' refers to a value, but is being us... Remove this comment to see the full error message
                  <IconAndAmount
                    // @ts-expect-error TS(2304): Cannot find name 'formattedValue'.
                    formattedValue={toInfo.formattedValue}
                    // @ts-expect-error TS(2304): Cannot find name 'tokenAddress'.
                    tokenAddress={toInfo.tokenAddress}
                    // @ts-expect-error TS(2304): Cannot find name 'chainId'.
                    chainId={toInfo.chainID}
                    // @ts-expect-error TS(2304): Cannot find name 'tokenSymbol'.
                    tokenSymbol={toInfo.tokenSymbol}
                    // @ts-expect-error TS(2304): Cannot find name 'iconSize'.
                    iconSize="w-7 h-7"
                    // @ts-expect-error TS(2304): Cannot find name 'textSize'.
                    textSize="text-md"
                    // @ts-expect-error TS(2304): Cannot find name 'styledCoin'.
                    styledCoin={true}
                  /> : null}
              // @ts-expect-error TS(2365): Operator '<' cannot be applied to types 'boolean' ... Remove this comment to see the full error message
              </div>
            </div>
          </div>
        </div>
        // @ts-expect-error TS(2304): Cannot find name 'br'.
        <br />
        // @ts-expect-error TS(2749): 'HorizontalDivider' refers to a value, but is bein... Remove this comment to see the full error message
        <HorizontalDivider />

      </div>
    </>
  } else {
    // @ts-expect-error TS(2304): Cannot find name 'content'.
    content = (
      <Error
        // @ts-expect-error TS(2304): Cannot find name 'text'.
        text="Sorry, there was a problem with that transaction hash."
        // @ts-expect-error TS(2304): Cannot find name 'param'.
        param={kappa}
        // @ts-expect-error TS(2304): Cannot find name 'subtitle'.
        subtitle="Unknown"
      />
    )
  }

  return <StandardPageContainer >{content}</StandardPageContainer>
}
export async function getServerSideProps(context) {
  const { data } = await client.query({
    query: GET_BRIDGE_TRANSACTIONS_QUERY,
    variables: {
      chainId: context.params.chainIdFrom,
      kappa: context.params.kappa,
      useMv: true
    },
  })

  return {
    props: {
      queryResult: data
    }, // will be passed to the page component as props
  }
}
