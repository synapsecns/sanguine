import {useEffect, useState} from 'react'
import {useLazyQuery} from '@apollo/client'

import {GET_BRIDGE_TRANSACTIONS_QUERY} from '@graphql/queries'

import {BridgeTransactionPageContent} from '@components/BridgeTransaction/BridgeTransactionPageContent'
import {BridgeTransactionLoader} from '@components/BridgeTransaction/BridgeTransactionLoader'
import {Error} from '@components/Error'
import {StandardPageContainer} from '@components/layouts/StandardPageContainer'
import {useRouter} from 'next/router'
import {useSearchParams} from 'next/navigation'

export function BridgeTransaction() {
  const router = useRouter()
  const [transaction, setTransaction] = useState()
  const search = useSearchParams()
  const { kappa } = router.query
  const chainId = Number(search.get('chainIdFrom'))

  const [getBridgeTransaction, { loading, error, data }] = useLazyQuery(
    GET_BRIDGE_TRANSACTIONS_QUERY,
    {
      variables: {
        chainId,
        kappa,
      },
    }
  )

  useEffect(() => {
    if (data) {
      setTransaction(data.bridgeTransactions[0])
    }

    getBridgeTransaction()
  }, [data])

  let content

  if (loading) {
    content = <BridgeTransactionLoader />
  } else if (error) {
    content = (
      <Error
        text="Sorry, there was a problem with that transaction hash."
        param={kappa}
        subtitle="Unknown"
      />
    )
  } else if (!!transaction) {
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

  return (
    <StandardPageContainer title="Bridge Transaction">
      {content}
    </StandardPageContainer>
  )
}
