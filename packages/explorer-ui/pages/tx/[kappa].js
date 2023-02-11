import {ApolloClient, HttpLink, InMemoryCache} from '@apollo/client'

import {Error} from '@components/Error'
import {StandardPageContainer} from '@components/layouts/StandardPageContainer'
import {useRouter} from 'next/router'
import {useSearchParams} from 'next/navigation'

import {GET_BRIDGE_TRANSACTIONS_QUERY,} from '@graphql/queries'
import {API_URL} from '@graphql'
import {HorizontalDivider} from "@components/misc/HorizontalDivider";
import {UniversalSearch} from "@components/pages/Home/UniversalSearch";
import {timeAgo} from "@utils/timeAgo";
import {IconAndAmount} from "@components/misc/IconAndAmount";
import {BridgeTransactionTable} from "@components/BridgeTransaction/BridgeTransactionTable";

const link = new HttpLink({
  uri: API_URL,
  useGETForQueries: true,
})

const client = new ApolloClient({
  link: link,
  cache: new InMemoryCache(),
  fetchPolicy: 'network-only',
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

  let content

  if (!!transaction) {
    content = <>
      <div className="flex items-center mt-10 mb-10">
        <h3 className="text-white text-4xl font-semibold">{kappa}</h3>
      </div>
      <HorizontalDivider />
      <UniversalSearch placeholder={`txid: ${kappa}`} />
      <BridgeTransactionTable queryResult={queryResult.bridgeTransactions} />

      <HorizontalDivider />
      <div className="pb-6">
        <div className="py-6">
          <h3 className="text-white text-xl font-medium">
            {fromInfo.time
              ? timeAgo({ timestamp: fromInfo.time })
              : timeAgo({ timestamp: toInfo.time })}
          </h3>
        </div>
        <div className="flex gap-y-2 flex-col">
          <div className="flex gap-x-4">
            <p className="text-white text-opacity-60">Requested</p>
            <p className="text-white ">{fromInfo.time}</p>
          </div>
          <div className="flex gap-x-4">
            <p className="text-white text-opacity-60">Confirmed</p>
            <p className="text-white ">{toInfo.time}</p>
          </div>
          <div className="flex gap-x-4">
            <p className="text-white text-opacity-60">Elapsed</p>
            <p className="text-white ">30 seconds</p>
          </div>
          <div className="flex mt-4">
            <div className="flex gap-x-6 w-1/2">
              <h1 className="text-white text-2xl text-opacity-60">Sent</h1>
              <IconAndAmount
                formattedValue={fromInfo.formattedValue}
                tokenAddress={fromInfo.tokenAddress}
                chainId={fromInfo.chainId}
                tokenSymbol={fromInfo.tokenSymbol}
                iconSize="w-6 h-6"
                textSize="text-sm"
                styledCoin={true}
              />
            </div>
            <div className="flex gap-x-6 w-1/2">
              <h1 className="text-white text-2xl text-opacity-60">
                Received
              </h1>
              <IconAndAmount
                formattedValue={toInfo.formattedValue}
                tokenAddress={toInfo.tokenAddress}
                chainId={toInfo.chainId}
                tokenSymbol={toInfo.tokenSymbol}
                iconSize="w-6 h-6"
                textSize="text-sm"
                styledCoin={true}
              />
            </div>
          </div>
        </div>
      </div>
      <HorizontalDivider />
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

  return <StandardPageContainer title="">{content}</StandardPageContainer>
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
