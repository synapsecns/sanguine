import Card from '@components/tailwind/Card'

import { ChainInfo } from '@components/misc/ChainInfo'
import { IconAndAmount } from '@components/misc/IconAndAmount'
import { Indicator } from '@components/misc/StatusIndicators'
import { StyleAddress } from '@components/misc/StyleAddress'

import { timeAgo } from '@utils/timeAgo'
import { getBridgeTransactionUrl } from '@urls'
import Link from 'next/link'

export function TransactionCard({ txn, ordinal }) {
  const { kappa, pending, fromInfo, toInfo } = txn
  // const navigate = useNavigate()

  const backgroundColor =
    ordinal % 2 === 0 ? 'bg-transparent' : 'bg-[#D9D9D9] bg-opacity-5'

  return (
    <a
      className="overflow-visible cursor-pointer"
      role="link"
      href={getBridgeTransactionUrl({
        hash: kappa,
        chainIdFrom: fromInfo.chainId,
        chainIdTo: toInfo.chainId,
      })}
    >
      <Card
        className={`flex-wrap md:flex lg:flex justify-between mt-2 mb-2 rounded-none ${backgroundColor}`}
      >
        <div className="space-y-2">
          <div className="flex items-center text-sm text-white">
            <div className="mr-2">
              <StyleAddress sourceInfo={fromInfo} />
            </div>
            <span className="mr-2 text-gray-400"> sent</span>{' '}
            <IconAndAmount
              formattedValue={fromInfo.formattedValue}
              tokenAddress={fromInfo.tokenAddress}
              chainId={fromInfo.chainId}
              tokenSymbol={fromInfo.tokenSymbol}
              iconSize="w-4 h-4"
              textSize="text-sm"
              styledCoin={true}
            />
            <span className="ml-2 mr-2">on</span>
            <ChainInfo
              chainId={fromInfo.chainId}
              imgClassName="w-4 h-4"
              textClassName="text-white"
            />
          </div>
          <div className="flex items-center text-sm text-white">
            {pending ? (
              <span className="mr-2 text-gray-400">
                tx pending on destination chain
              </span>
            ) : (
              <>
                <span className="mr-2 text-gray-400">to</span>
                <div className="mr-1">
                  <StyleAddress sourceInfo={toInfo} />
                </div>
                <span className="ml-2 mr-2 text-gray-400">on</span>
                <ChainInfo
                  chainId={toInfo.chainId}
                  imgClassName="w-4 h-4"
                  textClassName="text-white"
                />
              </>
            )}
          </div>
        </div>

        <div className="space-y-4">
          <div className="hidden md:block lg:block">
            <Indicator indicatorType={pending} />
          </div>
          <div className="text-sm text-gray-400">
            {fromInfo.time
              ? timeAgo({ timestamp: fromInfo.time })
              : timeAgo({ timestamp: toInfo.time })}
          </div>
        </div>
      </Card>
    </a>
  )
}

export function AllTransactions({ txns }) {
  return txns.map((txn, i) => <TransactionCard txn={txn} key={i} />)
}

export {
  TransactionCardLoader,
  TransactionsLoader,
} from './TransactionCardLoader'
