import _ from 'lodash'

import { Table } from '@components/TransactionTable/Table'
import {StyleAddress} from "@components/misc/StyleAddress";
import {IconAndAmount} from "@components/misc/IconAndAmount";
import {ChainInfo} from "@components/misc/ChainInfo";
import {timeAgo} from "@utils/timeAgo";
import {getBridgeTransactionUrl} from "@urls";
import {ellipsizeString} from "@utils/ellipsizeString";

export function BridgeTransactionTable({queryResult}) {
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

  queryResult.map((txn) => {
    const { kappa, pending, fromInfo, toInfo } = txn

      let items = [
        <StyleAddress sourceInfo={fromInfo} />,
        <StyleAddress sourceInfo={toInfo} />,
        <IconAndAmount
          formattedValue={fromInfo.formattedValue}
          tokenAddress={fromInfo.tokenAddress}
          chainId={fromInfo.chainId}
          tokenSymbol={fromInfo.tokenSymbol}
          iconSize="w-6 h-6"
          textSize="text-sm"
          styledCoin={true}
        />,
        <IconAndAmount
          formattedValue={toInfo.formattedValue}
          tokenAddress={toInfo.tokenAddress}
          chainId={toInfo.chainId}
          tokenSymbol={toInfo.tokenSymbol}
          iconSize="w-6 h-6"
          textSize="text-sm"
          styledCoin={true}
        />,
        <ChainInfo
          chainId={fromInfo.chainId}
          imgClassName="w-6 h-6"
          textClassName="text-white"
          txHash={fromInfo.hash}
        />,
        <ChainInfo
          chainId={toInfo.chainId}
          imgClassName="w-6 h-6"
          textClassName="text-white"
          txHash={toInfo.hash}
        />,
        fromInfo.time
            ? timeAgo({ timestamp: fromInfo.time })
            : timeAgo({ timestamp: toInfo.time }),
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
