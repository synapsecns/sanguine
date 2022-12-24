import { ChevronDoubleRightIcon } from '@heroicons/react/outline'

import Grid from '@components/tailwind/Grid'

import { TransactionDetails } from '@components/BridgeTransaction/TransactionDetails'
import { Indicator } from '@components/misc/StatusIndicators'
import { ContainerCard } from '@components/ContainerCard'
import { HorizontalDivider } from '@components/misc/HorizontalDivider'
import { UniversalSearch } from '@components/pages/home/UniversalSearch'
import { IconAndAmount } from '@components/misc/IconAndAmount'
import {
  TransactionCard,
  TransactionCardLoader,
} from '@components/TransactionCard'
import { timeAgo } from '@utils/timeAgo'

export function BridgeTransactionPageContent({ txn }) {
  const { pending } = txn

  if (!pending) {
    return <Transaction {...txn} />
  } else {
    return (
      <>
        <NotificationBanner pending={pending} />
        <Transaction {...txn} />
      </>
    )
  }
}

function NotificationBanner({ pending }) {
  return (
    <div className="flex justify-center">
      <ContainerCard
        className="w-1/2 mt-5"
        title={`Transaction ${pending}`}
        titleClassName="text-gray-400 "
        subtitle={<Indicator indicatorType={pending} />}
      />
    </div>
  )
}

function Transaction(txn) {
  const { pending, kappa, fromInfo, toInfo } = txn

  if (!pending) {
    return (
      <>
        <div className="flex items-center mt-10 mb-10">
          <h3 className="text-white text-4xl font-semibold">{kappa}</h3>
        </div>
        <HorizontalDivider />
        <UniversalSearch placeholder={`txid: ${txn.kappa}`} />
        <div className="pb-2 px-4 sm:px-6 lg:px-8">
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
                      <TransactionCard txn={txn} />
                    </tbody>
                  </table>
                </div>
              </div>
            </div>
          </div>
        </div>
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
    )
  } else if (fromInfo.hash && fromInfo.address && fromInfo.chainId) {
    return (
      <div className="flex items-center justify-center mt-10 mb-10">
        <Grid gap={8} cols={{ sm: 1, md: 7, lg: 7 }}>
          <div className="col-span-3 ">
            <TransactionDetails info={fromInfo} subtitle="Origin" />
          </div>
          <div className="flex items-center justify-center lg:col-span-1">
            <ChevronDoubleRightIcon className="w-10 h-10 text-purple-500" />
          </div>
          <div className="col-span-3">
            <TransactionDetails info={toInfo} subtitle="Destination" />
          </div>
        </Grid>
      </div>
    )
  } else {
    return (
      <div className="flex items-center justify-center mt-10 mb-10">
        <TransactionDetails info={toInfo} subtitle="Destination" />
      </div>
    )
  }
}
