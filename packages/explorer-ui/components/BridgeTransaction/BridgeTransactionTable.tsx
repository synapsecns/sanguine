import _ from 'lodash'
import { Table } from '@components/TransactionTable/Table'
import { StyleAddress } from '@components/misc/StyleAddress'
import { IconAndAmount } from '@components/misc/IconAndAmount'
import { ChainInfo } from '@components/misc/ChainInfo'
import { timeAgo } from '@utils/timeAgo'
import { addressToSymbol } from '@utils/addressToSymbol'
import { formatDateTime, formatDate } from '@utils/formatDate'
import { getBridgeTransactionUrl } from '@urls'
import Link from 'next/link'

export function BridgeTransactionTable({ queryResult }) {
  const handlePending = (date) => {
    const now = new Date().getTime()
    const timeDiff = now - date * 1000
    if (timeDiff > 86400000) {
      return <p>Indexing</p>
    } else {
      return <p>Pending</p>
    }
  }

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
        className="no-underline block w-full group"
      >
      <div className='flex flex-row'>
        <div className="flex flex-col w-[12.5%] px-1 py-2">
          <span className="text-gray-400">
            {new Date(fromInfo.time * 1000).toDateString() === new Date().toDateString()
              ? 'Today'
              : formatDate(new Date(fromInfo.time * 1000).toISOString().split('T')[0])}
          </span>
          <span className="text-gray-400 group-hover:text-gray-400 hidden group-hover:block mt-2">
            {new Date(fromInfo.time * 1000).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
          </span>
        </div>
      <div className="flex flex-col w-[12.5%] py-2">
        <StyleAddress sourceInfo={fromInfo} />
        {pending ? (
          <StyleAddress sourceInfo={fromInfo} />
        ) : (
          fromInfo.address === toInfo.address ? (
            <span className="text-gray-400">to Self</span>
          ) : (
            <StyleAddress sourceInfo={toInfo} />
          )
        )}
      </div>
      <div className="flex flex-col w-[12.5%]">
        <ChainInfo
          chainId={fromInfo.chainID}
          imgClassName="w-6 h-6 rounded-full"
          txHash={fromInfo.hash}
          useExplorerLink={false}
          className='py-2'
        />
        <IconAndAmount
          formattedValue={fromInfo.formattedValue}
          tokenAddress={fromInfo.tokenAddress}
          chainId={fromInfo.chainID}
          tokenSymbol={addressToSymbol({
            tokenAddress: fromInfo.tokenAddress,
            chainId: fromInfo.chainID,
          }) || fromInfo.tokenSymbol}
          iconSize="w-6 h-6 rounded-full"
          textSize="text-sm"
          styledCoin={true}
          className='py-2'
        />
      </div>
      <div className="flex flex-col w-[12.5%]">
        <ChainInfo
          chainId={pending ? fromInfo.destinationChainID : toInfo?.chainID}
          imgClassName="w-6 h-6 rounded-full"
          txHash={pending ? '' : toInfo?.hash}
          useExplorerLink={false}
          className='py-2'
        />
        <IconAndAmount
          formattedValue={pending ? fromInfo.formattedValue : toInfo?.formattedValue}
          tokenAddress={pending ? fromInfo.tokenAddress : toInfo?.tokenAddress}
          chainId={pending ? fromInfo.chainID : toInfo?.chainID}
          tokenSymbol={addressToSymbol({
            tokenAddress: pending ? fromInfo.tokenAddress : toInfo?.tokenAddress,
            chainId: pending ? fromInfo.chainID : toInfo?.chainID,
          }) || (pending ? fromInfo.tokenSymbol : toInfo?.tokenSymbol)}
          iconSize="w-6 h-6 rounded-full"
          textSize="text-sm"
          styledCoin={true}
          className='py-2'
        />
        </div>
        <div className="text-gray-400 px-1 py-2">
          {fromInfo.time
          ? timeAgo({ timestamp: fromInfo.time }) + ' ago'
            : timeAgo({ timestamp: toInfo?.time }) + ' ago'}
        </div>
        <div className="flex flex-col w-[12.5%] group-hover:text-white hidden group-hover:block px-1 py-2">
          <span className="text-white">▶</span>
        </div>
      </div>

      </Link>
    ]
    const row = {
      items,
      key: kappa,
    }
    tableRows.push(row)
  })
  return <Table header={[]} body={tableRows} />
}
