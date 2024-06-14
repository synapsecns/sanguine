import _ from 'lodash'
import { Table } from '@components/TransactionTable/Table'
import { StyleAddress } from '@components/misc/StyleAddress'
import { IconAndAmount } from '@components/misc/IconAndAmount'
import { ChainInfo } from '@components/misc/ChainInfo'
import { timeAgo } from '@utils/timeAgo'
import { getBridgeTransactionUrl } from '@urls'
import { ellipsizeString } from '@utils/ellipsizeString'
import { addressToSymbol } from '@utils/addressToSymbol'
import { formatDateTime } from '@utils/formatDate'
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

    // ]
    const items = [
      <Link
        href={getBridgeTransactionUrl({
        hash: txn.kappa,
        chainIdFrom: txn.fromInfo.chainID,
        chainIdTo: txn.toInfo?.chainID,
      })}
      className="no-underline block w-full"
    >
      <div className="flex flex-col">
        <div className="flex items-center space-x-2">
          <StyleAddress sourceInfo={fromInfo} />
          <span className='text-gray-400'>sent</span>
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
          />
          <span className='text-gray-400'>on  </span>
          <ChainInfo
            chainId={fromInfo.chainID}
            imgClassName="w-6 h-6 rounded-full"
            txHash={fromInfo.hash}
            useExplorerLink={false}
          />
        </div>
        <div className="flex items-center space-x-2 mt-2">
          <span className='text-gray-400'>to </span>
          {pending ? (
            <StyleAddress sourceInfo={fromInfo} />
          ) : (
            <StyleAddress sourceInfo={toInfo} />
          )}
          <span className='text-gray-400'>on  </span>
            {pending ? (
              <ChainInfo
                chainId={fromInfo.destinationChainID}
              imgClassName="w-6 h-6 rounded-full"
              txHash={''}
              useExplorerLink={false}
            />
            ) : (
              <ChainInfo
                chainId={toInfo?.chainID}
                imgClassName="w-6 h-6 rounded-full"
                txHash={toInfo?.hash}
                useExplorerLink={false}
            />
          )}
        </div>
        <div className="ml-auto">
          {fromInfo.time
            ? timeAgo({ timestamp: fromInfo.time  }) + ' ago'
            : timeAgo({ timestamp: toInfo?.time  }) + ' ago'}
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
