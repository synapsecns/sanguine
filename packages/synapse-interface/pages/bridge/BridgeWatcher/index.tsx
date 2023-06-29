import { fetchBlockNumber } from '@wagmi/core'
import { useSigner } from 'wagmi'
import SYNAPSE_BRIDGE_ABI from '@abis/synapseBridge.json'
import { Contract, Signer } from 'ethers'
import { BRIDGE_CONTRACTS, SYNAPSE_CCTP_CONTRACTS } from '@constants/bridge'
import { useEffect, useState } from 'react'
import { Interface } from '@ethersproject/abi'
import _ from 'lodash'
import Grid from '@tw/Grid'
import Card from '@tw/Card'
import BridgeEvent from './BridgeEvent'
import { BridgeWatcherTx } from '@types'
import { GETLOGS_SIZE, GETLOGS_REQUEST_COUNT } from '@constants/bridgeWatcher'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'
import { useSelector } from 'react-redux';
import { RootState } from '@/store/store'
import SYNAPSE_CCTP_ABI from '@abis/synapseCCTP.json'
import * as CHAINS from '@constants/chains/master'

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
  destinationAddress,
}: {
  fromChainId: number
  toChainId: number
  address: string
  destinationAddress: string
}) => {
  const bridgeTxHashes = useSelector((state: RootState) => state.bridge)
  const [fromTransactions, setFromTransactions] = useState([])
  const [fromSynapseContract, setFromSynapseContract] = useState<Contract>()
  const [fromCCTPContract, setFromCCTPContract] = useState<Contract>()

  const [fromSigner, setFromSigner] = useState<Signer>()
  const { data: fromSignerRaw } = useSigner({ chainId: fromChainId })
  const { providerMap } = useSynapseContext()

  const createContractsAndInterfaces = (chainId, provider) => {
    const bridgeAddress = BRIDGE_CONTRACTS[chainId];
    const synapseCCTPAddress = SYNAPSE_CCTP_CONTRACTS[chainId];

    const validBridgeContract = BRIDGE_CONTRACTS[fromChainId]
    ? BRIDGE_CONTRACTS[fromChainId]
    : BRIDGE_CONTRACTS[CHAINS.ETH.id]
    const bridgeContract = new Contract(
      validBridgeContract,
      SYNAPSE_BRIDGE_ABI,
      fromSigner
    )

    const bridgeInterface = new Interface(SYNAPSE_BRIDGE_ABI);

    const synapseCCTPContract = synapseCCTPAddress
      ? new Contract(synapseCCTPAddress, SYNAPSE_CCTP_ABI, provider)
      : null;

    const synapseCCTPInterface = synapseCCTPAddress
      ? new Interface(SYNAPSE_CCTP_ABI)
      : null;


      return { bridgeContract, bridgeInterface, synapseCCTPContract, synapseCCTPInterface };
    }

  const fetchFromBridgeEvents = async (
    currentFromBlock: number,
    provider: any,
    adjustedAddress: string
  ) => {
    const { bridgeContract, bridgeInterface, synapseCCTPContract, synapseCCTPInterface } = createContractsAndInterfaces(provider.network.chainId, provider);
    let allFromEvents = []
    let retryCount = 0;
    const maxRetries = 5; // Adjust this as needed

    // fetch bridge logs
    for (let i = 0; i < GETLOGS_REQUEST_COUNT; i++) {
      let successful = false;
      while (!successful && retryCount < maxRetries) {
        try {
          const fromEvents = await getLogs(
            currentFromBlock - GETLOGS_SIZE * i,
            provider,
            bridgeContract,
            adjustedAddress
          )
          allFromEvents.push(fromEvents)
          successful = true;
        } catch (error) {
          retryCount++;
          console.log(`getLogs failed, retrying in ${Math.pow(2, retryCount)} seconds...`);
          await new Promise(resolve => setTimeout(resolve, Math.pow(2, retryCount) * 1000));
        }
      }
      if (retryCount === maxRetries) {
        console.error("getLogs failed after maximum retries");
        break;
      }
    }
    // fetch synapseCCTP logs if the contract exists for the chain
    if (synapseCCTPContract) {
      for (let i = 0; i < GETLOGS_REQUEST_COUNT; i++) {
        let successful = false;
        while (!successful && retryCount < maxRetries) {
          try {
            const fromEvents = await getLogs(
              currentFromBlock - GETLOGS_SIZE * i,
              provider,
              synapseCCTPContract,
              adjustedAddress
            )
            allFromEvents.push(fromEvents)
            successful = true;
          } catch (error) {
            retryCount++;
            console.log(`getLogs failed, retrying in ${Math.pow(2, retryCount)} seconds...`);
            await new Promise(resolve => setTimeout(resolve, Math.pow(2, retryCount) * 1000));
          }
        }
        if (retryCount === maxRetries) {
          console.error("getLogs failed after maximum retries");
          break;
        }
      }
    }

    return _.flatten(allFromEvents)
  }


  const parseLogs = (
    fromEvents: any[],
    bridgeInterface: Interface,
    synapseCCTPInterface: Interface,
    bridgeAddress: string,
    synapseCCTPAddress: string
  ) => {
    return fromEvents
      .map((log) => {
        // Select the correct interface based on the contract address
        const iface = log.address.toLowerCase() === bridgeAddress.toLowerCase() ? bridgeInterface : synapseCCTPInterface;

        return {
          ...iface.parseLog(log).args,
          transactionHash: log.transactionHash,
          blockNumber: Number(log.blockNumber),
          contractEmittedFrom: log.address.toLowerCase()
        }
      })
      .filter((log) => checkTxIn(log))
  }


  const fetchTimestampsAndReceipts = (
    parsedLogs: any[],
    provider: any
  ) => {
    return Promise.all([
      Promise.all(parsedLogs.map((log) => getBlock(log.blockNumber, provider))),
      Promise.all(
        parsedLogs.map((log) =>
          getTransactionReceipt(log.transactionHash, provider)
        )
      ),
    ])
  }

  const generateBridgeTransactions = (
    parsedLogs: any[],
    inputTimestamps: any[],
    transactionReceipts: any[],
    address: string,
    fromChainId: number,
    destinationAddress: string
  ) => {
    return _.zip(parsedLogs, inputTimestamps, transactionReceipts).map(
      ([parsedLog, timestampObj, txReceipt]) => {
        return generateBridgeTx(
          true,
          address,
          fromChainId,
          parsedLog,
          timestampObj,
          txReceipt,
          destinationAddress
        )
      }
    )
  }

  const getFromBridgeEvents = async (): Promise<BridgeWatcherTx[]> => {
    const currentFromBlock = await fetchBlockNumber({ chainId: fromChainId })
    const provider = providerMap[fromChainId]
    const iface = new Interface(SYNAPSE_BRIDGE_ABI)
    const adjustedAddress = destinationAddress ? destinationAddress : address

       // Define the contracts and interfaces here
       const {
        bridgeContract,
        bridgeInterface,
        synapseCCTPContract,
        synapseCCTPInterface
      } = createContractsAndInterfaces(provider.network.chainId, provider);

    const fromEvents = await fetchFromBridgeEvents(
      currentFromBlock,
      provider,
      adjustedAddress
    )
    // Use the correct interfaces and addresses when parsing the logs
    const parsedLogs = parseLogs(
      fromEvents,
      bridgeInterface,
      synapseCCTPInterface,
      bridgeContract.address,
      synapseCCTPContract.address
    )

    const [inputTimestamps, transactionReceipts] = await fetchTimestampsAndReceipts(
      parsedLogs,
      provider
    )
    const txObjects = generateBridgeTransactions(
      parsedLogs,
      inputTimestamps,
      transactionReceipts,
      address,
      fromChainId,
      destinationAddress
    )

    return txObjects
  }

  useEffect(() => {
    if (fromSigner && fromChainId && toChainId && address) {
      const validBridgeContract = BRIDGE_CONTRACTS[fromChainId]
        ? BRIDGE_CONTRACTS[fromChainId]
        : BRIDGE_CONTRACTS[1]
      const fromSynapseContract = new Contract(
        validBridgeContract,
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

    return () => setFromTransactions([...fromTransactions])
  }, [fromSynapseContract, bridgeTxHashes])

  useEffect(() => {
    setFromSigner(fromSignerRaw)
  }, [fromSignerRaw])

  return (
    <div className="space-y-2">
      {fromTransactions?.length > 0 && (
        <Card title="Bridge Watcher" divider={false} className="px-6 py-4">
          <Grid cols={{ xs: 1 }} gap={2}>
            {fromTransactions.map((fromEvent, i) => {
              return <BridgeEvent key={i} {...fromEvent} />
            })}
          </Grid>
        </Card>
      )}
    </div>
  )
}

export default BridgeWatcher
