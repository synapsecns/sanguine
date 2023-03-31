import _ from 'lodash'

function checkTxIn(tx) {
  if (tx.args?.chainId) {
    // if (tx.args?.chainId ?? false) { //
    return true
  } else {
    return false
  }
}

/**
 * @param {Transaction[]} transactions
 * @return {Transaction[][]}
 */
export function pairTxKappa(transactions) {
  const transactionsByHash = {}
  for (const tx of transactions) {
    transactionsByHash[tx.transactionHash] = tx
  }

  const inputTxns = transactions.filter((tx) => tx.kekTxSig).filter(checkTxIn)
  const outputTxns = transactions
    .filter((tx) => tx.kekTxSig)
    .filter((tx) => !checkTxIn(tx))

  const outputTxnsDict = {}

  for (const tx of outputTxns) {
    outputTxnsDict[tx.args?.kappa] = tx
  }

  const pairSetsByChain = []

  for (const inTx of inputTxns) {
    const outTx = outputTxnsDict[inTx.kekTxSig]
    if (outTx) {
      pairSetsByChain.push([inTx, outTx])
    } else {
      pairSetsByChain.push([inTx, undefined])
    }
  }
  const outTxnKeys = pairSetsByChain.map(
    ([inTx, outTx]) => outTx?.transactionHash
  )
  const remainingOuts = outputTxns.filter(
    (tx) => !outTxnKeys.includes(tx.transactionHash)
  )

  for (const outTx of remainingOuts) {
    pairSetsByChain.push([undefined, outTx])
  }

  return _.sortBy(pairSetsByChain, ([inTx, outTx]) => {
    return -(outTx?.timestamp ?? inTx?.timestamp)
  })
}
