import { BigNumber } from '@ethersproject/bignumber'
import { id } from '@ethersproject/hash'
import { ChainId } from '@constants/networks'
import { formatTimestampToDate } from '@utils/datetime'
import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'
import { useGenericSynapseContract } from '@hooks/contracts/useSynapseContract'
import { useSingleCallResult } from '@hooks/multicall'
import { useTerraKappaCheck } from '@hooks/terra/useTerraKappaCheck'

import { SubTransactionItem } from './TransactionItems'
import BlockCountdown from './BlockCountdown'

export default function PairedTransactionItem({ inputTx, outputTx }) {
  const { chainId } = useActiveWeb3React()
  const targetChainId = inputTx?.args?.chainId ?? ChainId.ETH
  const synapseContract = useGenericSynapseContract(targetChainId)
  const kekTxSig = id(inputTx?.identifier ?? inputTx?.transactionHash ?? '')
  const kappaExistsResult = useSingleCallResult(
    targetChainId,
    synapseContract,
    'kappaExists',
    [kekTxSig],
    { resultOnly: true }
  )

  const terraKappaExists = useTerraKappaCheck({
    kekTxSig,
    isTerra: targetChainId == ChainId.TERRA,
  })

  const outputExists = kappaExistsResult?.[0] ?? terraKappaExists ?? false

  const inAmount = inputTx?.inputTokenAmount
  let outAmount = outputTx?.args?.amount

  if (outputTx?.event == 'TokenWithdraw') {
    // TokenWithdraw event features amount before fees.
    // Other events feature amount after fees.
    outAmount = outAmount.sub(outputTx?.args?.fee)
  }

  const inToken = inputTx?.inputToken
  const outToken = outputTx?.outputToken

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
            inToken={inToken}
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
    return ''
  }
}
