import { Log, Provider } from '@ethersproject/abstract-provider'
import { Contract } from '@ethersproject/contracts'
import { Interface } from '@ethersproject/abi'

/**
 * Extracts the first log from a transaction receipt that matches
 * the provided contract and any of the provided event names.
 *
 * @param provider - Ethers provider for the network
 * @param txHash - Transaction hash
 * @param contract - Contract that should have emitted the event
 * @param eventNames - Names of the events that could have been emitted
 * @returns The first log that matches the contract and any of the event names
 * @throws If the transaction receipt cannot be retrieved, or if no matching log is found
 */
export const getMatchingTxLog = async (
  provider: Provider,
  txHash: string,
  contract: Contract,
  eventNames: string[]
): Promise<Log> => {
  const txReceipt = await provider.getTransactionReceipt(txHash)
  if (!txReceipt) {
    throw new Error('Failed to get transaction receipt')
  }
  const topics = getEventTopics(contract.interface, eventNames)
  // Find the log with the correct contract address and topic matching any of the provided topics
  const matchingLog = txReceipt.logs.find((log) => {
    return log.address === contract.address && topics.includes(log.topics[0])
  })
  if (!matchingLog) {
    // Throw an error and include the event names in the message
    throw new Error(
      `Contract ${
        contract.address
      } in transaction ${txHash} did not emit any of the expected events: ${eventNames.join(
        ', '
      )}`
    )
  }
  return matchingLog
}

const getEventTopics = (
  contractInterface: Interface,
  eventNames: string[]
): string[] => {
  // Filter events that match the provided event names and map them to their topics
  return Object.values(contractInterface.events)
    .filter((event) => eventNames.includes(event.name))
    .map((event) => contractInterface.getEventTopic(event))
}
