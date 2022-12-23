import _ from 'lodash'
import { useQuery } from '@apollo/client'

import {
  TransactionCard,
  TransactionCardLoader,
} from '@components/TransactionCard'
import { GET_LATEST_BRIDGE_TRANSACTIONS_QUERY } from '@graphql/queries'

export function LatestBridgeTransactions({ queryResult }) {
  // const { error, data } = useQuery(GET_LATEST_BRIDGE_TRANSACTIONS_QUERY, {
  //   variables: { includePending: false, page: 1 },
  // })

  let content

  // if (!queryResult) {
  //   content = [...Array(5).keys()].map((i) => (
  //     <TransactionCardLoader key={i} ordinal={i} />
  //   ))
  // } else if (queryResult.error) {
  //   content = 'Error'
  // } else {
  let { latestBridgeTransactions } = queryResult

  latestBridgeTransactions = _.orderBy(
    latestBridgeTransactions,
    'fromInfo.time',
    ['desc']
  ).slice(0, 10)

  content = latestBridgeTransactions.map((txn, i) => (
    <TransactionCard txn={txn} key={i} ordinal={i} />
  ))
  // }

  // return <>{content}</>

  return (
    <div className="px-4 sm:px-6 lg:px-8">
      <div className="mt-8 flex flex-col">
        <div className="-my-2 -mx-4 overflow-x-auto sm:-mx-6 lg:-mx-8">
          <div className="inline-block min-w-full py-2 align-middle">
            <div className="overflow-hidden shadow-sm ring-1 ring-black ring-opacity-5">
              <table className="min-w-full">
                <thead className="">
                  <tr>
                    <th
                      scope="col"
                      className="px-2 py-2 text-left text-md font-bold text-white"
                    >
                      From
                    </th>
                    <th
                      scope="col"
                      className="px-2 py-2 text-left text-md font-bold text-white"
                    >
                      To
                    </th>
                    <th
                      scope="col"
                      className="px-2 py-2 text-left text-md font-bold text-white"
                    >
                      Initial
                    </th>
                    <th
                      scope="col"
                      className="px-2 py-2 text-left text-md font-bold text-white"
                    >
                      Final
                    </th>
                    <th
                      scope="col"
                      className="px-2 py-2 text-left text-md font-bold text-white"
                    >
                      Origin
                    </th>
                    <th
                      scope="col"
                      className="px-2 py-2 text-left text-md font-bold text-white"
                    >
                      Destination
                    </th>
                    <th
                      scope="col"
                      className="px-2 py-2 text-left text-md font-bold text-white"
                    >
                      Date
                    </th>
                    <th
                      scope="col"
                      className="px-2 py-2 text-left text-md font-bold text-white"
                    >
                      Tx ID
                    </th>
                  </tr>
                </thead>
                <tbody>
                  {latestBridgeTransactions.map((txn, i) => (
                    <TransactionCard txn={txn} key={i} ordinal={i} />
                  ))}
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
