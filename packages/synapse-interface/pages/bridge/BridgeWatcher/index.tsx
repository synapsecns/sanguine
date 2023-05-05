import { fetchBlockNumber, getProvider } from '@wagmi/core'
import { useSigner } from 'wagmi'
import SYNAPSE_BRIDGE_ABI from '@abis/synapseBridge.json'
import { Contract } from 'ethers'
import { BRIDGE_CONTRACTS, INCLUDED_BRIDGE_EVENTS } from '@constants/bridge'
import { useEffect, useState } from 'react'
import { toHexStr } from '@utils/toHexStr'
import _, { add } from 'lodash'
import * as CHAINS from '@constants/chains/master'
import { id } from '@ethersproject/hash'
import { BRIDGABLE_TOKENS } from '@constants/tokens'
import { BigNumber } from 'ethers'
import Grid from '@tw/Grid'
import Card from '@tw/Card'
import PairedTransactionItem from './PairedTransactionItem'
import {
  DOG,
  LINK,
  GOHM,
  HIGHSTREET,
  JUMP,
  NFD,
  NEWO,
  VSTA,
  GMX,
  SDT,
  UNIDX,
  USDB,
  SFI,
  H2O,
  L2DAO,
  PLS,
  AGEUR,
  SYN,
  NUSD,
  NETH,
  FRAX,
  SYN_FRAX,
  WBTC,
  KLAYTN_WETH,
  NOTE,
} from '@constants/tokens/master'
import { WETH } from '@constants/tokens/swapMaster'
const BridgeWatcher = ({
  fromChainId,
  toChainId,
  address,
  destinationAddress,
}: {
  fromChainId: number
  toChainId: number
  address: string
  destinationAddress: string
}) => {
  const [transactions, setTransactions] = useState([])
  const [synapseContract, setSynapseContract] = useState({
    from: null,
    to: null,
  })
  const { data: toSigner } = useSigner({ chainId: toChainId })
  const { data: fromSigner } = useSigner({ chainId: fromChainId })

  const checkTxIn = (tx) => {
    if (tx.args?.chainId) {
      return true
    } else {
      return false
    }
  }

  const pairTxKappa = (transactions) => {
    let transactionsByHash = {}
    for (const tx of transactions) {
      transactionsByHash[tx.transactionHash] = tx
    }

    const inputTxns = transactions.filter((tx) => tx.kekTxSig).filter(checkTxIn)
    const outputTxns = transactions
      .filter((tx) => tx.kekTxSig)
      .filter((tx) => !checkTxIn(tx))

    let outputTxnsDict = {}

    for (const tx of outputTxns) {
      outputTxnsDict[tx.args?.kappa] = tx
    }

    let pairSetsByChain = []

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
  const getLogs = async (isFrom: boolean, currentBlock: number) => {
    const contract = isFrom ? synapseContract.from : synapseContract.to
    const synapseContractEventHandles = Object.keys(contract.filters).filter(
      (key) => INCLUDED_BRIDGE_EVENTS.includes(key)
    )
    const pastEventsByHandle = []

    for (const eventHandle of synapseContractEventHandles) {
      const accountList =
        destinationAddress && address !== destinationAddress
          ? [address, destinationAddress]
          : [address]
      try {
        const filter = await contract.filters[eventHandle](accountList)
        pastEventsByHandle.push(
          await contract.queryFilter(filter, toHexStr(currentBlock - 2040))
        )
      } catch (e) {
        console.log(e)
        return null
      }
    }

    return pastEventsByHandle.filter((x) => x)
  }

  const getBridgeEvents = async () => {
    const currentFromBlock = await fetchBlockNumber({ chainId: fromChainId })
    const currentToBlock = await fetchBlockNumber({ chainId: toChainId })

    const fromEvents = await getLogs(true, currentFromBlock)
    const toEvents = await getLogs(false, currentToBlock)
    // const pastEvents =
    //   toEvents && toEvents ? [address, destinationAddress] : [address]
    const pastEvents = _.flatten([_.flatten(fromEvents), _.flatten(toEvents)])
    // console.log(synapseContract, eventHandle, accountList)
    const [inputTimestamps, transactionReceipts] = await Promise.all([
      Promise.all(pastEvents.map((eventObj) => eventObj.getBlock())),
      Promise.all(
        pastEvents.map((eventObj) => eventObj.getTransactionReceipt())
      ),
    ])

    const transactionsWithTimestamps = _.zip(
      pastEvents,
      inputTimestamps,
      transactionReceipts
    ).map(([eventObj, timestampObj, txReceipt]) => {
      return mergeTxObj(fromChainId, eventObj, timestampObj, txReceipt)
    })
    return transactionsWithTimestamps
  }
  useEffect(() => {
    if (fromSigner && fromChainId && toChainId && address) {
      const fromSynapseContract = new Contract(
        BRIDGE_CONTRACTS[fromChainId],
        SYNAPSE_BRIDGE_ABI,
        fromSigner
      )

      setSynapseContract({ ...synapseContract, from: fromSynapseContract })

      // const toProvider = getProvider({
      //   chainId: toChainId,
      // })
      // console.log('toProvider', toProvider)
      // const toSynapseContract = new Contract(
      //   BRIDGE_CONTRACTS[toChainId],
      //   SYNAPSE_BRIDGE_ABI,
      //   toProvider
      // )
      // console.log('here44', fromSynapseContract, toSynapseContract)
      // setSynapseContract({ to: toSynapseContract, from: fromSynapseContract })
    }
  }, [fromSigner])
  useEffect(() => {
    if (
      toSigner &&
      synapseContract.from &&
      fromChainId &&
      toChainId &&
      address
    ) {
      const provider = getProvider({
        chainId: toChainId,
      })
      const toSynapseContract = new Contract(
        BRIDGE_CONTRACTS[toChainId],
        SYNAPSE_BRIDGE_ABI,
        provider
      )

      setSynapseContract({ ...synapseContract, to: toSynapseContract })
    }
  }, [toSigner, synapseContract.from])

  useEffect(() => {
    if (toSigner && fromSigner && synapseContract.from && synapseContract.to) {
      getBridgeEvents().then((txs) => {
        setTransactions(txs)
      })
    }
  }, [synapseContract])

  return (
    <div className="space-y-2">
      {transactions?.length > 0 && (
        <Card title="Bridge Watcher" divider={false}>
          <Grid cols={{ xs: 1 }} gap={2}>
            {pairTxKappa(transactions).map(([inputTx, outputTx]) => {
              return (
                <PairedTransactionItem
                  key={`${inputTx?.transactionHash}-${outputTx?.transactionHash}`}
                  inputTx={inputTx}
                  outputTx={outputTx}
                  chainId={fromChainId}
                  synapseContract={synapseContract}
                />
              )
            })}
          </Grid>
        </Card>
      )}
    </div>
  )
}

export default BridgeWatcher

const mergeTxObj = (chainId, eventObj, timestampObj, txReceipt) => {
  let TOKEN_HASH_MAP = {}

  for (const [chainId, tokensOnChain] of _.toPairs(BRIDGABLE_TOKENS)) {
    TOKEN_HASH_MAP[chainId] = {}
    for (const token of tokensOnChain) {
      TOKEN_HASH_MAP[chainId][_.toLower(token.addresses[chainId])] = token
    }
  }

  TOKEN_HASH_MAP[CHAINS.AVALANCHE.id][
    _.toLower(GMX.wrapperAddresses[CHAINS.AVALANCHE.id])
  ] = GMX

  const swapTokenAddr = _.toLower(eventObj.args.token)
  const { timestamp } = timestampObj ?? {}

  let outputTokenAddr

  if (
    [
      SYN,
      LINK,
      HIGHSTREET,
      DOG,
      JUMP,
      FRAX,
      NFD,
      GOHM,
      AGEUR,
      H2O,
      L2DAO,
      PLS,
      NEWO,
      VSTA,
      SFI,
      SDT,
      UNIDX,
      USDB,
      GMX,
      WBTC,
      KLAYTN_WETH,
      NOTE,
    ]
      .map((t) => _.toLower(t.addresses[chainId]))
      .includes(swapTokenAddr)
  ) {
    outputTokenAddr = TOKEN_HASH_MAP[chainId][swapTokenAddr].addresses[chainId]
  } else if (swapTokenAddr == _.toLower(SYN_FRAX.addresses[chainId])) {
    outputTokenAddr = FRAX.addresses[chainId]
  } else if (swapTokenAddr == _.toLower(GMX.wrapperAddresses[chainId])) {
    outputTokenAddr = GMX.addresses[chainId]
  } else if (swapTokenAddr == _.toLower(NETH.addresses[chainId])) {
    outputTokenAddr = txReceipt.logs[txReceipt.logs.length - 2].address
  } else if (swapTokenAddr == _.toLower(WETH.addresses[chainId])) {
    if (chainId == CHAINS.ETH.id) {
      outputTokenAddr = txReceipt.logs[txReceipt.logs.length - 2].address
    } else {
      outputTokenAddr = txReceipt.logs[txReceipt.logs.length - 1].address
    }
  } else if (swapTokenAddr == _.toLower(NUSD.addresses[chainId])) {
    if (chainId == CHAINS.ETH.id) {
      if (eventObj.event == 'TokenWithdraw') {
        outputTokenAddr = txReceipt.logs[txReceipt.logs.length - 1].address
      } else {
        outputTokenAddr = txReceipt.logs[txReceipt.logs.length - 2].address
      }
    } else if (chainId == CHAINS.POLYGON.id) {
      outputTokenAddr = txReceipt.logs[txReceipt.logs.length - 3].address
    } else {
      outputTokenAddr = txReceipt.logs[txReceipt.logs.length - 2].address
    }
  } else {
    outputTokenAddr = txReceipt.logs[txReceipt.logs.length - 2].address
  }

  let inputTokenAmount
  if (
    _.toLower(txReceipt.logs[0]?.address) ==
      _.toLower(GMX.addresses[CHAINS.ARBITRUM.id]) ||
    _.toLower(txReceipt.logs[1]?.address) ==
      _.toLower(GMX.addresses[CHAINS.AVALANCHE.id])
  ) {
    inputTokenAmount = txReceipt.logs[1].data
  } else {
    inputTokenAmount = txReceipt.logs[0].data
  }

  let inputTokenAddr
  if (
    _.toLower(txReceipt.logs[1].address) ==
    _.toLower(GMX.addresses[CHAINS.AVALANCHE.id])
  ) {
    inputTokenAddr = _.toLower(GMX.addresses[CHAINS.AVALANCHE.id])
  } else {
    inputTokenAddr = _.toLower(txReceipt.logs[0].address)
  }
  // console.log({
  //   amt: txReceipt.logs[0].data,
  //   inputTokenAmount,
  //   gmxwrapaddr: GMX.wrapperAddresses[ChainId.AVALANCHE],
  //   swapTokenAddr,

  //   txReceipt
  // })
  // const inputTokenAmount = BigNumber.from(txReceipt.logs[0].data)
  // const inputTokenAddr = _.toLower(txReceipt.logs[0].address)
  inputTokenAmount = BigNumber.from(inputTokenAmount)
  inputTokenAddr = _.toLower(inputTokenAddr)
  outputTokenAddr = _.toLower(outputTokenAddr)
  return {
    // ...txReceipt, // ...tx,
    ...eventObj,
    timestamp,
    chainId,
    kekTxSig: id(eventObj.transactionHash),
    event: eventObj.event,
    args: eventObj.args,
    kappa: eventObj.args.kappa,
    inputTokenAmount: inputTokenAmount,
    inputTokenAddr: inputTokenAddr,
    outputTokenAddr: outputTokenAddr,
    inputToken: TOKEN_HASH_MAP[chainId][inputTokenAddr],
    outputToken: TOKEN_HASH_MAP[chainId][outputTokenAddr],
    txReceipt,
  }
}
