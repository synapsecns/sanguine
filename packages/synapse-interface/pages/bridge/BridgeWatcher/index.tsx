import { fetchBlockNumber } from '@wagmi/core'
import { useSigner } from 'wagmi'
import SYNAPSE_BRIDGE_ABI from '@abis/synapseBridge.json'
import { Contract, Signer } from 'ethers'
import { BRIDGE_CONTRACTS } from '@constants/bridge'
import { useEffect, useState } from 'react'
import { Interface } from '@ethersproject/abi'
import _ from 'lodash'
import Grid from '@tw/Grid'
import Card from '@tw/Card'
import BridgeEvent from './BridgeEvent'
import { BridgeWatcherTx } from '@types'
import { getProvider } from '@utils/getProvider'
import { GETLOGS_SIZE, GETLOGS_REQUEST_COUNT } from '@constants/bridgeWatcher'
import {
  getLogs,
  getBlock,
  getTransactionReceipt,
  generateBridgeTx,
  checkTxIn,
} from '@utils/bridgeWatcher'
const BridgeWatcher = ({
  fromChainId,
  toChainId,
  address,
}: {
  fromChainId: number
  toChainId: number
  address: string
  destinationAddress: string
}) => {
  const [fromTransactions, setFromTransactions] = useState([])
  const [fromSynapseContract, setFromSynapseContract] = useState<Contract>()
  const [fromSigner, setFromSigner] = useState<Signer>()
  const { data: fromSignerRaw } = useSigner({ chainId: fromChainId })

  const getFromBridgeEvents = async (): Promise<BridgeWatcherTx[]> => {
    const currentFromBlock = await fetchBlockNumber({ chainId: fromChainId })
    const provider = getProvider(fromChainId)
    const iface = new Interface(SYNAPSE_BRIDGE_ABI)
    let allFromEvents = []
    for (let i = 0; i < GETLOGS_REQUEST_COUNT; i++) {
      const fromEvents = await getLogs(
        currentFromBlock - GETLOGS_SIZE * i,
        provider,
        fromSynapseContract,
        address
      )
      allFromEvents.push(fromEvents)
    }
    const flattendEvents = _.flatten(allFromEvents)
    const parsedLogs = flattendEvents
      .map((log) => {
        return {
          ...iface.parseLog(log).args,
          transactionHash: log.transactionHash,
          blockNumber: Number(log.blockNumber),
        }
      })
      .filter((log) => checkTxIn(log))

    const [inputTimestamps, transactionReceipts] = await Promise.all([
      Promise.all(parsedLogs.map((log) => getBlock(log.blockNumber, provider))),
      Promise.all(
        parsedLogs.map((log) =>
          getTransactionReceipt(log.transactionHash, provider)
        )
      ),
    ])
    const txObjects = _.zip(
      parsedLogs,
      inputTimestamps,
      transactionReceipts
    ).map(([parsedLog, timestampObj, txReceipt]) => {
      return generateBridgeTx(
        true,
        address,
        fromChainId,
        parsedLog,
        timestampObj,
        txReceipt
      )
    })
    return txObjects
  }

  useEffect(() => {
    if (fromSigner && fromChainId && toChainId && address) {
      const fromSynapseContract = new Contract(
        BRIDGE_CONTRACTS[fromChainId],
        SYNAPSE_BRIDGE_ABI,
        fromSigner
      )
      setFromSynapseContract(fromSynapseContract)
    }
  }, [fromChainId, fromSigner])
  useEffect(() => {
    if (fromSynapseContract) {
      getFromBridgeEvents().then((txs) => {
        setFromTransactions(txs)
      })
    }
  }, [fromSynapseContract])
  useEffect(() => {
    setFromSigner(fromSignerRaw)
  }, [fromSignerRaw])

  return (
    <div className="space-y-2">
      {fromTransactions?.length > 0 && (
        <Card title="Bridge Watcher" divider={false}>
          <Grid cols={{ xs: 1 }} gap={2}>
            {fromTransactions.map((fromEvent) => {
              return (
                <BridgeEvent
                  key={`${fromEvent?.transactionHash}-${fromEvent?.transactionHash}`}
                  {...fromEvent}
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
