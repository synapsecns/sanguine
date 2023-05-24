import { BridgeWatcherTx } from '@types'
import { useEffect, useState, memo } from 'react'
import { Contract, Signer } from 'ethers'
import { useSigner } from 'wagmi'
import SYNAPSE_BRIDGE_ABI from '@abis/synapseBridge.json'
import { BRIDGE_CONTRACTS } from '@constants/bridge'
import { PendingCreditTransactionItem } from '@components/TransactionItems'
import { fetchBlockNumber } from '@wagmi/core'
import { Interface } from '@ethersproject/abi'
import { GETLOGS_SIZE } from '@constants/bridgeWatcher'
import _ from 'lodash'
import EventCard from './EventCard'
import BlockCountdown from './BlockCountdown'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import {
  getLogs,
  getBlock,
  getTransactionReceipt,
  generateBridgeTx,
  checkTxIn,
} from '@utils/bridgeWatcher'
import { remove0xPrefix } from '@/utils/remove0xPrefix'

const DestinationTx = (fromEvent: BridgeWatcherTx) => {
  const [toEvent, setToEvent] = useState<BridgeWatcherTx>(undefined)
  const [toSynapseContract, setToSynapseContract] =
    useState<Contract>(undefined)
  const [toSigner, setToSigner] = useState<Signer>()
  const { data: toSignerRaw } = useSigner({ chainId: fromEvent.toChainId })
  const [completedConf, setCompletedConf] = useState(false)
  const [attempted, setAttempted] = useState(false)
  const { providerMap } = useSynapseContext()

  const getToBridgeEvent = async (): Promise<BridgeWatcherTx> => {
    const currentFromBlock = await fetchBlockNumber({
      chainId: fromEvent.toChainId,
    })
    const provider = providerMap[fromEvent.toChainId]

    const iface = new Interface(SYNAPSE_BRIDGE_ABI)
    let allToEvents = []
    let i = 0
    let afterOrginTx = true
    while (afterOrginTx) {
      const startBlock = currentFromBlock - GETLOGS_SIZE * i
      const blockTimestamp = await (
        await getBlock(currentFromBlock - GETLOGS_SIZE * i, provider)
      ).timestamp
      if (blockTimestamp < fromEvent.timestamp) {
        afterOrginTx = false
      }
      const fromEvents = await getLogs(
        startBlock,
        provider,
        toSynapseContract,
        fromEvent.toAddress
      )
      allToEvents.push(fromEvents)
      i++

      // Break if cannot find tx
      if (i > 30) {
        afterOrginTx = false
      }
    }
    const flattendEvents = _.flatten(allToEvents)
    const parsedLog = flattendEvents
      .map((log) => {
        return {
          ...iface.parseLog(log).args,
          transactionHash: log.transactionHash,
          blockNumber: Number(log.blockNumber),
        }
      })
      .filter((log: any) => {
        const convertedKappa = remove0xPrefix(log.kappa)
        return !checkTxIn(log) && convertedKappa === fromEvent.kappa
      })?.[0]
    if (parsedLog) {
      const [inputTimestamp, transactionReceipt] = await Promise.all([
        getBlock(parsedLog.blockNumber, provider),
        getTransactionReceipt(parsedLog.transactionHash, provider),
      ])

      return generateBridgeTx(
        false,
        fromEvent.toAddress,
        fromEvent.toChainId,
        parsedLog,
        inputTimestamp,
        transactionReceipt,
        fromEvent.toAddress
      )
    }
    setAttempted(true)
    return null
  }

  useEffect(() => {
    if (toSigner && fromEvent) {
      const toSynapseContract = new Contract(
        BRIDGE_CONTRACTS[fromEvent.toChainId],
        SYNAPSE_BRIDGE_ABI,
        toSigner
      )
      setToSynapseContract(toSynapseContract)
    }
  }, [fromEvent, toSigner])

  useEffect(() => {
    if (completedConf && toSynapseContract && !toEvent && attempted) {
      getToBridgeEvent().then((tx) => {
        setToEvent(tx)
      })
    }
  }, [completedConf, toEvent, attempted, toSynapseContract])

  useEffect(() => {
    if (toSynapseContract && !toEvent && !completedConf) {
      getToBridgeEvent().then((tx) => {
        setToEvent(tx)
        return
      })
    }
    return
  }, [completedConf, toEvent, attempted, toSynapseContract])

  useEffect(() => {
    setToSigner(toSignerRaw)
  }, [toSignerRaw])

  return (
    <div className="flex items-center">
      <div className="flex-initial w-auto h-full align-middle mt-[22px]">
        <BlockCountdown
          fromEvent={fromEvent}
          toEvent={toEvent ?? null}
          setCompletedConf={setCompletedConf}
        />
      </div>
      <div className="flex-initial w-full">
        {toEvent ? <EventCard {...toEvent} /> : null}
      </div>
    </div>
  )
}
export default DestinationTx
