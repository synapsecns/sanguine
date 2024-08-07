import Link from 'next/link'
import { Table } from '@components/TransactionTable/Table'
import { StyleAddress } from '@components/misc/StyleAddress'
import { IconAndAmount } from '@components/misc/IconAndAmount'
import { ChainInfo } from '@components/misc/ChainInfo'
import { addressToSymbol } from '@utils/addressToSymbol'
import Arrow from '@components/icons/Arrow'
import { timeAgo } from '@utils/timeAgo'
import { formatDate } from '@utils/formatDate'
import { getBridgeTransactionUrl } from '@urls'

export const BridgeTransactionTable = ({ queryResult }) => {
  const tableRows = []
  queryResult?.map((txn) => {
    const { kappa, pending, fromInfo, toInfo } = txn

    const items = [
      <Link
        href={getBridgeTransactionUrl({
          hash: txn.kappa,
          chainIdFrom: txn.fromInfo.chainID,
          chainIdTo: txn.toInfo?.chainID,
        })}
        className="block w-full no-underline group"
      >
        <div className="flex flex-row space-x-3">
          <div className="flex flex-col min-w-24">
            <span className="text-gray-400">
              {new Date(fromInfo.time * 1000).toDateString() ===
              new Date().toDateString()
                ? 'Today'
                : formatDate(
                    new Date(fromInfo.time * 1000).toISOString().split('T')[0]
                  )}
            </span>
            <span className="hidden mt-2 text-gray-400 group-hover:text-gray-400 group-hover:block">
              {new Date(fromInfo.time * 1000).toLocaleTimeString([], {
                hour: '2-digit',
                minute: '2-digit',
              })}
            </span>
          </div>
          <div className="flex flex-col space-y-2 min-w-24">
            <StyleAddress sourceInfo={fromInfo} />
            {pending ? (
              <StyleAddress sourceInfo={fromInfo} />
            ) : fromInfo.address === toInfo.address ? (
              <span className="text-gray-400">to Self</span>
            ) : (
              <StyleAddress sourceInfo={toInfo} />
            )}
          </div>
          <div className="flex flex-col space-y-2 min-w-28 w-fit">
            <ChainInfo
              chainId={fromInfo.chainID}
              imgClassName="w-4 h-4 rounded-full"
              txHash={fromInfo.hash}
              useExplorerLink={false}
            />
            <IconAndAmount
              value={fromInfo.value}
              tokenAddress={fromInfo.tokenAddress}
              chainId={fromInfo.chainID}
              tokenSymbol={
                addressToSymbol({
                  tokenAddress: fromInfo.tokenAddress,
                  chainId: fromInfo.chainID,
                }) || fromInfo.tokenSymbol
              }
              iconSize="w-4 h-4 rounded-full"
              // textSize="text-sm"
              // styledCoin={true}
            />
          </div>
          <div className="relative">
            <Arrow color="white" />
          </div>
          <div className="flex flex-col space-y-2 min-w-36">
            <ChainInfo
              chainId={pending ? fromInfo.destinationChainID : toInfo?.chainID}
              imgClassName="w-4 h-4 rounded-full"
              txHash={pending ? '' : toInfo?.hash}
              useExplorerLink={false}
            />
            <IconAndAmount
              value={pending ? fromInfo.value : toInfo?.value}
              tokenAddress={
                pending ? fromInfo.tokenAddress : toInfo?.tokenAddress
              }
              chainId={pending ? fromInfo.chainID : toInfo?.chainID}
              tokenSymbol={
                addressToSymbol({
                  tokenAddress: pending
                    ? fromInfo.tokenAddress
                    : toInfo?.tokenAddress,
                  chainId: pending ? fromInfo.chainID : toInfo?.chainID,
                }) || (pending ? fromInfo.tokenSymbol : toInfo?.tokenSymbol)
              }
              iconSize="w-4 h-4 rounded-full"
              // textSize="text-sm"
              // styledCoin={true}
            />
          </div>
          <div className="text-gray-400">
            {fromInfo.time
              ? timeAgo({ timestamp: fromInfo.time }) + ' ago'
              : timeAgo({ timestamp: toInfo?.time }) + ' ago'}
          </div>
          <div className="hidden group-hover:text-white group-hover:block">
            <span className="text-white">▶</span>
          </div>
        </div>
      </Link>,
    ]
    const row = {
      items,
      key: kappa,
    }
    tableRows.push(row)
  })
  return <Table header={[]} body={tableRows} />
}
