import _ from 'lodash'
import { useEffect, useState, memo, useMemo } from 'react'
import { useSigner } from 'wagmi'
import { fetchBlockNumber } from '@wagmi/core'
import { Contract, Signer } from 'ethers'
import { Interface } from '@ethersproject/abi'
import {
  ChevronRightIcon,
  ChevronDoubleRightIcon,
} from '@heroicons/react/outline'
import { BridgeWatcherTx } from '@types'
import { getNetworkTextColor } from '@/styles/chains'
import SYNAPSE_BRIDGE_ABI from '@abis/synapseBridge.json'
import { BRIDGE_CONTRACTS } from '@constants/bridge'
import { CHAINS_BY_ID } from '@/constants/chains'
import { GETLOGS_SIZE } from '@constants/bridgeWatcher'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import {
  getLogs,
  getBlock,
  getTransactionReceipt,
  generateBridgeTx,
  checkTxIn,
} from '@utils/bridgeWatcher'
import { remove0xPrefix } from '@/utils/remove0xPrefix'
import EventCard from './EventCard'
import BlockCountdown from './BlockCountdown'
import { CreditedTransactionItem } from '@components/TransactionItems'
import SYNAPSE_CCTP_ABI from '@abis/synapseCCTP.json'
import { SYNAPSE_CCTP_CONTRACTS } from '@constants/bridge'

const DestinationTx = (fromEvent: BridgeWatcherTx) => {
  const [toEvent, setToEvent] = useState<BridgeWatcherTx>(undefined)
  const [toSynapseContract, setToSynapseContract] =
    useState<Contract>(undefined)
  const [toCCTPContract, setToCCTPContract] = useState<Contract>(undefined)
  const [toSigner, setToSigner] = useState<Signer>()
  const { data: toSignerRaw } = useSigner({ chainId: fromEvent.toChainId })
  const [completedConf, setCompletedConf] = useState(false)
  const [attempted, setAttempted] = useState(false)
  const { providerMap } = useSynapseContext()

  const networkTextColorClass: string = useMemo(() => {
    const networkChainById = CHAINS_BY_ID[fromEvent.chainId]
    return getNetworkTextColor(networkChainById?.color)
  }, [fromEvent.toChainId])

  const getToBridgeEvent = async (): Promise<BridgeWatcherTx> => {
    const headOnDestination = await fetchBlockNumber({
      chainId: fromEvent.toChainId,
    })
    const provider = providerMap[fromEvent.toChainId]

    const isCCTP = fromEvent.contractEmittedFrom.toLowerCase() === SYNAPSE_CCTP_CONTRACTS[fromEvent.chainId].toLowerCase() ? true : false

    const iface = new Interface(isCCTP ? SYNASE_CCTP_ABI : SYNAPSE_BRIDGE_ABI)

    let allToEvents = []
    let i = 0
    let afterOrginTx = true
    while (afterOrginTx) {
      const startBlock = headOnDestination - GETLOGS_SIZE * i

      // get timestamp from from block
      const blockRaw = await getBlock(startBlock - (GETLOGS_SIZE + 1), provider)
      const blockTimestamp = blockRaw?.timestamp

      // Exit loop if destination block was mined before the block for the origin tx
      if (blockTimestamp < fromEvent.timestamp) {
        afterOrginTx = false
      }

      const fromEventsBridge = await getLogs(
        startBlock,
        provider,
        toSynapseContract,
        fromEvent.toAddress
      )

      const fromEventsCCTP = await getLogs(
        startBlock,
        provider,
        toCCTPContract,
        fromEvent.toAddress
      )

      allToEvents.push(fromEventsBridge, fromEventsCCTP)
      i++

      // Break if cannot find tx
      if (i > 30) {
        afterOrginTx = false
      }
    }
    const flattendEvents = _.flatten(allToEvents)
    let parsedLogs
    if (!isCCTP) {
      parsedLogs = flattendEvents
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
      })}
      else {
        parsedLogs = flattendEvents
        .map((log) => {
          return {
            ...iface.parseLog(log).args,
            transactionHash: log.transactionHash,
            blockNumber: Number(log.blockNumber),
          }
        }).filter((log: any) => {
          return log.requestID === fromEvent.kappa
        })


      }

    const parsedLog = parsedLogs?.[0]
    if (parsedLog) {
      const [inputTimestamp, transactionReceipt] = await Promise.all([
        getBlock(parsedLog.blockNumber, provider),
        getTransactionReceipt(parsedLog.transactionHash, provider),
      ])

      const destBridgeTx = generateBridgeTx(
        false,
        fromEvent.toAddress,
        fromEvent.toChainId,
        parsedLog,
        inputTimestamp,
        transactionReceipt,
        fromEvent.toAddress
      )
      setAttempted(true)
      return destBridgeTx
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

      // Initialize CCTP Contract when signer and fromEvent are available
      if (SYNAPSE_CCTP_CONTRACTS[fromEvent.toChainId]) {
        const toCCTPContract = new Contract(
          SYNAPSE_CCTP_CONTRACTS[fromEvent.toChainId],
          SYNAPSE_CCTP_ABI,
          toSigner
        )
        setToCCTPContract(toCCTPContract)
      }
    }
  }, [fromEvent, toSigner])

  // Listens for confirmations to complete and if so, recheck destination chain for logs
  useEffect(() => {
    if (completedConf && (toSynapseContract || toCCTPContract) && attempted) {
      getToBridgeEvent().then((tx) => {
        setToEvent(tx)
      })
    }
  }, [completedConf, toEvent, fromEvent, toSynapseContract, toCCTPContract])

  // Listens for SynapseContract to be set and if so, will check destination chain for logs if there is no toEvent
  useEffect(() => {
    if ((toSynapseContract || toCCTPContract) && !toEvent) {
      getToBridgeEvent().then((tx) => {
        setToEvent(tx)
        return
      })
    }
    return
  }, [toSynapseContract, toCCTPContract])

  useEffect(() => {
    setToSigner(toSignerRaw)
  }, [toSignerRaw])

  return (
    <div className="flex items-center">
      <div className="flex items-center px-3 pt-6 pb-1 align-middle">
        {toEvent ? (
          <ChevronDoubleRightIcon
            className={`
              w-5 h-5
              place-self-center
              ${networkTextColorClass}
            `}
          />
        ) : (
          <ChevronRightIcon
            className={`
            w-5 h-5
            place-self-center
            ${networkTextColorClass}
            animate-pulse
            `}
          />
        )}
      </div>

      {toEvent ? (
        <div className="flex-initial w-full">
          <EventCard {...toEvent} />
        </div>
      ) : (
        <div className="flex-initial w-auto h-full align-middle mt-[22px]">
          <BlockCountdown
            fromEvent={fromEvent}
            toEvent={toEvent ?? null}
            setCompletedConf={setCompletedConf}
          />
        </div>
      )}
    </div>
  )
}
export default DestinationTx
