import _ from 'lodash'

import { Table } from '@components/TransactionTable/Table'
import { StyleAddress } from "@components/misc/StyleAddress";
import { IconAndAmount } from "@components/misc/IconAndAmount";
import { ChainInfo } from "@components/misc/ChainInfo";
import { timeAgo } from "@utils/timeAgo";
import { getBridgeTransactionUrl } from "@urls";
import { ellipsizeString } from "@utils/ellipsizeString";

export function BridgeTransactionTable({ queryResult }) {
  console.log("UUU", queryResult)
  let headers = [
    'From',
    'To',
    'Initial',
    'Final',
    'Origin',
    'Destination',
    'Date',
    'Tx ID'
  ]


  let tableRows = []
  let pendingContent = (<p>Pending</p>)
  queryResult.map((txn) => {
    const { kappa, pending, fromInfo, toInfo } = txn

      let items = [
        <StyleAddress sourceInfo={fromInfo} />,
        pending ? pendingContent :
        <StyleAddress sourceInfo={toInfo} />,
        <IconAndAmount
          formattedValue={fromInfo.formattedValue}
          tokenAddress={fromInfo.tokenAddress}
          chainId={fromInfo.chainID}
          tokenSymbol={fromInfo.tokenSymbol}
          iconSize="w-6 h-6"
          textSize="text-sm"
          styledCoin={true}
        />,
        pending ? pendingContent :
        <IconAndAmount
          formattedValue={toInfo.formattedValue}
          tokenAddress={toInfo.tokenAddress}
          chainId={toInfo.chainID}
          tokenSymbol={toInfo.tokenSymbol}
          iconSize="w-6 h-6"
          textSize="text-sm"
          styledCoin={true}
        />,
        <ChainInfo
          chainId={fromInfo.chainID}
          imgClassName="w-6 h-6"
          textClassName="text-white"
          txHash={fromInfo.hash}
        />,
        pending ? pendingContent :
        <ChainInfo
          chainId={toInfo.chainID}
          imgClassName="w-6 h-6"
          textClassName="text-white"
          txHash={toInfo.hash}
        />,
        fromInfo.time
          ? timeAgo({ timestamp: fromInfo.time })
          : timeAgo({ timestamp: toInfo?.time }),
        <a
          className="underline"
          href={getBridgeTransactionUrl({
            hash: txn.kappa,
            chainIdFrom: txn.fromInfo.chainID,
            chainIdTo: txn.toInfo?.chainID,
          })}
        >
          {ellipsizeString({ string: txn.kappa, limiter: 6 })}
        </a>
      ]

    let row = {
      items,
      key: kappa
    }
    tableRows.push(row);
  })
  return (
    <Table header={headers} body={tableRows} />
  )
}
