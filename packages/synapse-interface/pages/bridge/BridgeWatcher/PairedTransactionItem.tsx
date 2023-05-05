import { BigNumber } from '@ethersproject/bignumber'
import { id } from '@ethersproject/hash'

import * as CHAINS from '@constants/chains/master'

import { formatTimestampToDate } from '@utils/time'

import { SubTransactionItem } from '../../../components/TransactionItems'
import { readContract } from '@wagmi/core'
import SYNAPSE_BRIDGE_ABI from '@abis/synapseBridge.json'
import { useState } from 'react'
import BlockCountdown from './BlockCountdown'
const PairedTransactionItem = ({
  inputTx,
  outputTx,
  chainId,
  synapseContract,
}) => {
  const [outputExists, setOutputExists] = useState(false)
  const targetChainId = outputTx?.args?.chainId ?? CHAINS.ETH.id
  const kekTxSig = id(inputTx?.identifier ?? inputTx?.transactionHash ?? '')

  readContract({
    address: synapseContract?.from?.address,
    abi: SYNAPSE_BRIDGE_ABI,
    functionName: 'kappaExists',
    args: [kekTxSig],
    chainId,
  }).then((kappaExistsResult) => {
    setOutputExists(kappaExistsResult?.[0] ?? false)
  })
  const inAmount = inputTx?.inputTokenAmount
  let outAmount = outputTx?.args?.amount

  if (outputTx?.event == 'TokenWithdraw') {
    // TokenWithdraw event features amount before fees.
    // Other events feature amount after fees.
    outAmount = outAmount.sub(outputTx?.args?.fee)
  }

  let inToken = inputTx?.inputToken
  let outToken = outputTx?.outputToken

  try {
    return (
      <div>
        <div className="flex items-center text-gray-500">
          <div className="flex-1 ">
            <div className="pb-1 text-sm text-gray-500">
              {inputTx && formatTimestampToDate(inputTx?.timestamp)}
            </div>
          </div>
          <div className="flex-shrink-0 p-2 align-middle w-9">
            <div />
          </div>
          <div className="flex-1 ">
            <div className="pb-1 pl-4 text-sm text-gray-500">
              {outputTx &&
                outputTx?.timestamp &&
                formatTimestampToDate(outputTx?.timestamp)}
            </div>
          </div>
        </div>
        <div className="flex items-center text-gray-500">
          <div className="flex-1 ">
            {inputTx && (
              <SubTransactionItem
                {...inputTx}
                token={inToken}
                tokenAmount={inAmount}
              />
            )}
          </div>

          <BlockCountdown
            inputTx={inputTx}
            outputTx={outputTx}
            // inToken={inToken}
            outToken={outToken}
            outputExists={outputExists}
            outAmount={outAmount}
            fromChainId={inputTx?.chainId ?? chainId}
            toChainId={targetChainId}
          />
        </div>
      </div>
    )
  } catch (e) {
    return null
  }
}

export default PairedTransactionItem
