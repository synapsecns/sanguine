import {ChainInfo} from '@components/misc/ChainInfo'
import {IconAndAmount} from '@components/misc/IconAndAmount'
import {StyleAddress} from '@components/misc/StyleAddress'

import {timeAgo} from '@utils/timeAgo'
import {getBridgeTransactionUrl} from '@urls'
import {ellipsizeString} from '@utils/ellipsizeString'

export function TransactionCard({ txn, ordinal }) {
  const { kappa, pending, fromInfo, toInfo } = txn
  // const navigate = useNavigate()

  const backgroundColor =
    ordinal % 2 === 0 ? 'bg-transparent' : 'bg-[#D9D9D9] bg-opacity-5'

  return (
    <tr key={txn.kappa}>
      {/* <a
                        className="overflow-visible cursor-pointer"
                        role="link"
                        // href={getBridgeTransactionUrl({
                        //   hash: txn.kappa,
                        //   chainIdFrom: txn.fromInfo.chainId,
                        //   chainIdTo: txn.toInfo.chainId,
                        // })}
                      > */}
      <td className="whitespace-nowrap px-2 py-2 text-sm  text-white">
        <StyleAddress sourceInfo={fromInfo} />
      </td>
      <td className="whitespace-nowrap px-2 py-2 text-sm text-white">
        <StyleAddress sourceInfo={toInfo} />
      </td>
      <td className="whitespace-nowrap px-2 py-2 text-sm text-white">
        <IconAndAmount
          formattedValue={fromInfo.formattedValue}
          tokenAddress={fromInfo.tokenAddress}
          chainId={fromInfo.chainId}
          tokenSymbol={fromInfo.tokenSymbol}
          iconSize="w-6 h-6"
          textSize="text-sm"
          styledCoin={true}
        />
      </td>
      <td className="whitespace-nowrap px-2 py-2 text-sm text-white">
        <IconAndAmount
          formattedValue={toInfo.formattedValue}
          tokenAddress={toInfo.tokenAddress}
          chainId={toInfo.chainId}
          tokenSymbol={toInfo.tokenSymbol}
          iconSize="w-6 h-6"
          textSize="text-sm"
          styledCoin={true}
        />
      </td>
      <td className="whitespace-nowrap px-2 py-2 text-sm text-white">
        <ChainInfo
          chainId={fromInfo.chainId}
          imgClassName="w-6 h-6"
          textClassName="text-white"
          txHash={fromInfo.hash}
        />
      </td>
      <td className="whitespace-nowrap px-2 py-2 text-sm text-white">
        <ChainInfo
          chainId={toInfo.chainId}
          imgClassName="w-6 h-6"
          textClassName="text-white"
          txHash={toInfo.hash}
        />
      </td>
      <td className="whitespace-nowrap px-2 py-2 text-sm text-white">
        {fromInfo.time
          ? timeAgo({ timestamp: fromInfo.time })
          : timeAgo({ timestamp: toInfo.time })}
      </td>
      <td className="whitespace-nowrap px-2 py-2 text-sm text-white">
        <a
          className="underline"
          href={getBridgeTransactionUrl({
            hash: txn.kappa,
            chainIdFrom: txn.fromInfo.chainId,
            chainIdTo: txn.toInfo.chainId,
          })}
        >
          {ellipsizeString({ string: txn.kappa, limiter: 6 })}
        </a>
      </td>
      {/* </a> */}
    </tr>
  )
}

export function AllTransactions({ txns }) {
  return txns.map((txn, i) => <TransactionCard txn={txn} key={i} />)
}

export {
  TransactionCardLoader,
  TransactionsLoader,
} from './TransactionCardLoader'
