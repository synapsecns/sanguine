import _ from 'lodash'

import { Table } from '@components/TransactionTable/Table'
import { StyleAddress } from "@components/misc/StyleAddress";
import { IconAndAmount } from "@components/misc/IconAndAmount";
import { ChainInfo } from "@components/misc/ChainInfo";
import { timeAgo } from "@utils/timeAgo";
import { getBridgeTransactionUrl } from "@urls";
import { ellipsizeString } from "@utils/ellipsizeString";

export function BridgeTransactionTable({ queryResult }) {
  const handlePending = (date) => {
    let now = new Date().getTime()
    let timeDiff = now - date *1000
    if (timeDiff > 86400000) {
      // @ts-expect-error TS(2304): Cannot find name 'p'.
      return <p>Indexing</p>
    } else {
      // @ts-expect-error TS(2304): Cannot find name 'p'.
      return <p>Pending</p>
    }

  }
  let headers = [
    'Initial',
    'Final',
    'Origin',
    'Destination',
    'From',
    'To',
    'Age',
    'TXID'
  ]


  let tableRows = []
  queryResult?.map((txn) => {
    const { kappa, pending, fromInfo, toInfo } = txn

      let items = [

        // @ts-expect-error TS(2749): 'IconAndAmount' refers to a value, but is being us... Remove this comment to see the full error message
        <IconAndAmount
          // @ts-expect-error TS(2304): Cannot find name 'formattedValue'.
          formattedValue={fromInfo.formattedValue}
          // @ts-expect-error TS(2304): Cannot find name 'tokenAddress'.
          tokenAddress={fromInfo.tokenAddress}
          // @ts-expect-error TS(2304): Cannot find name 'chainId'.
          chainId={fromInfo.chainID}
          // @ts-expect-error TS(2304): Cannot find name 'tokenSymbol'.
          tokenSymbol={fromInfo.tokenSymbol}
          // @ts-expect-error TS(2304): Cannot find name 'iconSize'.
          iconSize="w-4 h-4"
          // @ts-expect-error TS(2304): Cannot find name 'textSize'.
          textSize="text-sm"
          // @ts-expect-error TS(2304): Cannot find name 'styledCoin'.
          styledCoin={true}
        />,
        pending ? handlePending(fromInfo.time) :
        // @ts-expect-error TS(2749): 'IconAndAmount' refers to a value, but is being us... Remove this comment to see the full error message
        <IconAndAmount
          // @ts-expect-error TS(2304): Cannot find name 'formattedValue'.
          formattedValue={toInfo.formattedValue}
          // @ts-expect-error TS(2304): Cannot find name 'tokenAddress'.
          tokenAddress={toInfo.tokenAddress}
          // @ts-expect-error TS(2304): Cannot find name 'chainId'.
          chainId={toInfo.chainID}
          // @ts-expect-error TS(2304): Cannot find name 'tokenSymbol'.
          tokenSymbol={toInfo.tokenSymbol}
          // @ts-expect-error TS(2304): Cannot find name 'iconSize'.
          iconSize="w-4 h-4"
          // @ts-expect-error TS(2304): Cannot find name 'textSize'.
          textSize="text-sm"
          // @ts-expect-error TS(2304): Cannot find name 'styledCoin'.
          styledCoin={true}
        />,
        // @ts-expect-error TS(2749): 'ChainInfo' refers to a value, but is being used a... Remove this comment to see the full error message
        <ChainInfo
          // @ts-expect-error TS(2304): Cannot find name 'chainId'.
          chainId={fromInfo.chainID}
          // @ts-expect-error TS(2304): Cannot find name 'imgClassName'.
          imgClassName="w-6 h-6 rounded-full"
          // @ts-expect-error TS(2304): Cannot find name 'txHash'.
          txHash={fromInfo.hash}
          // @ts-expect-error TS(2304): Cannot find name 'useExplorerLink'.
          useExplorerLink={false}

        />,
        // @ts-expect-error TS(2749): 'ChainInfo' refers to a value, but is being used a... Remove this comment to see the full error message
        pending ?  <ChainInfo
        // @ts-expect-error TS(2304): Cannot find name 'chainId'.
        chainId={fromInfo.destinationChainID}
        // @ts-expect-error TS(2304): Cannot find name 'imgClassName'.
        imgClassName="w-6 h-6 rounded-full"
        // @ts-expect-error TS(2304): Cannot find name 'txHash'.
        txHash={""}
        // @ts-expect-error TS(2304): Cannot find name 'useExplorerLink'.
        useExplorerLink={false}

      /> :
        // @ts-expect-error TS(2749): 'ChainInfo' refers to a value, but is being used a... Remove this comment to see the full error message
        <ChainInfo
          // @ts-expect-error TS(2304): Cannot find name 'chainId'.
          chainId={toInfo.chainID}
          // @ts-expect-error TS(2304): Cannot find name 'imgClassName'.
          imgClassName="w-6 h-6 rounded-full"
          // @ts-expect-error TS(2304): Cannot find name 'txHash'.
          txHash={toInfo.hash}
          // @ts-expect-error TS(2304): Cannot find name 'useExplorerLink'.
          useExplorerLink={false}

        />,
        // @ts-expect-error TS(2749): 'StyleAddress' refers to a value, but is being use... Remove this comment to see the full error message
        <StyleAddress sourceInfo={fromInfo} />,
        pending ? handlePending(fromInfo.time) :
        // @ts-expect-error TS(2749): 'StyleAddress' refers to a value, but is being use... Remove this comment to see the full error message
        <StyleAddress sourceInfo={toInfo} />,
        fromInfo.time
          ? timeAgo({ timestamp: fromInfo.time })
          : timeAgo({ timestamp: toInfo?.time }),
        // @ts-expect-error TS(2304): Cannot find name 'a'.
        <a
          // @ts-expect-error TS(2304): Cannot find name 'className'.
          className="underline transition ease-out hover:text-[#8FEBFF]"
          // @ts-expect-error TS(2304): Cannot find name 'href'.
          href={getBridgeTransactionUrl({
            hash: txn.kappa,
            chainIdFrom: txn.fromInfo.chainID,
            chainIdTo: txn.toInfo?.chainID,
          })}
        >
          {ellipsizeString({ string: txn.kappa, limiter: 4 })}
        </a>
      ]

    let row = {
      // @ts-expect-error TS(18004): No value exists in scope for the shorthand propert... Remove this comment to see the full error message
      items,
      // @ts-expect-error TS(2304): Cannot find name 'kappa'.
      key: kappa
    }
    // @ts-expect-error TS(2304): Cannot find name 'tableRows'.
    tableRows.push(row);
  })
  return (
    // @ts-expect-error TS(2304): Cannot find name 'body'.
    <Table header={headers} body={tableRows} />
  )
}
