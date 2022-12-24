import {ApolloClient, HttpLink, InMemoryCache} from '@apollo/client'

import {BridgeTransactionPageContent} from '@components/BridgeTransaction/BridgeTransactionPageContent'
import {Error} from '@components/Error'
import {StandardPageContainer} from '@components/layouts/StandardPageContainer'
import {useRouter} from 'next/router'
import {useSearchParams} from 'next/navigation'

import {GET_BRIDGE_TRANSACTIONS_QUERY,} from '@graphql/queries'
import {API_URL} from '@graphql'

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

export default function BridgeTransaction({ transaction }) {
  const router = useRouter()
  const search = useSearchParams()
  const { kappa } = router.query
  const chainId = Number(search.get('chainIdFrom'))

  let content

  if (!!transaction) {
    content = <BridgeTransactionPageContent txn={transaction} />
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
  const { data: bridgeTransaction } = await client.query({
    query: GET_BRIDGE_TRANSACTIONS_QUERY,
    variables: {
      chainId: context.params.chainIdFrom,
      kappa: context.params.kappa,
    },
  })

  return {
    props: {
      transaction: bridgeTransaction.bridgeTransactions[0],
    }, // will be passed to the page component as props
  }
}
