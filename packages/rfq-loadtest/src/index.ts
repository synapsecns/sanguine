import fs from 'fs/promises'

import * as viemChains from 'viem/chains'
import { privateKeyToAccount } from 'viem/accounts'
import {
  createPublicClient,
  createWalletClient,
  formatUnits,
  http,
  parseUnits,
  publicActions,
  PublicClient,
  WalletClient,
} from 'viem'
import yaml from 'js-yaml'
import yargs from 'yargs'
import { hideBin } from 'yargs/helpers'
import axios from 'axios'
import { createNonceManager, jsonRpc } from 'viem/nonce'

import { ABI } from './abi.js'
import { delay, print, getRandomInt } from './utils.js'

const argv = await yargs(hideBin(process.argv))
  .option('configFile', {
    alias: 'c',
    type: 'string',
    description: 'Path to the config file',
    demandOption: true,
  })
  .option('pKeyIndex', {
    alias: 'p',
    type: 'number',
    description: 'Index of the private key',
    demandOption: true,
  })
  .help().argv

const configFilePath = argv.configFile
const configFileContent = await fs.readFile(configFilePath, 'utf-8')

let config: any
try {
  config = yaml.load(configFileContent) as object

  if (typeof config !== 'object' || config === null) {
    throw new Error()
  }

  if (
    typeof config.VOLLEY_MILLISECONDS_BETWEEN !== 'number' ||
    typeof config.VOLLEY_MIN_COUNT !== 'number' ||
    typeof config.VOLLEY_MAX_COUNT !== 'number'
  ) {
    throw new Error('Invalid configuration values for volley settings')
  }

  if (typeof config.CHAINS !== 'object' || config.CHAINS === null) {
    throw new Error('Invalid configuration for CHAINS')
  }

  if (typeof config.TEST_ROUTES !== 'object' || config.TEST_ROUTES === null) {
    throw new Error('Invalid configuration for TEST_ROUTES')
  }

  if (typeof config.TEST_BRIDGE_AMOUNT_UNITS !== 'number') {
    throw new Error('Invalid configuration for TEST_BRIDGE_AMOUNT_UNITS')
  }
  if (typeof config.MINIMUM_GAS_UNITS !== 'number') {
    throw new Error('Invalid configuration for MINIMUM_GAS_UNITS')
  }
  if (typeof config.REBALANCE_TO_UNITS !== 'number') {
    throw new Error('Invalid configuration for REBALANCE_TO_UNITS')
  }

  Object.entries(config.TEST_ROUTES).forEach(
    ([route, details]: [string, any]) => {
      if (typeof details !== 'object' || details === null) {
        throw new Error(`Invalid configuration for route: ${route}`)
      }
      if (
        typeof details.fromChainId !== 'number' ||
        typeof details.toChainId !== 'number' ||
        typeof details.testDistributionPercentage !== 'number'
      ) {
        throw new Error(`Invalid configuration values for route: ${route}`)
      }
    }
  )
} catch (error: any) {
  throw new Error(
    `Failed to parse ${configFilePath}. Check your syntax, structure, data, and for duplicates. \n${error.message}`
  )
}

const privateKeyIndex = argv.pKeyIndex
if (typeof privateKeyIndex !== 'number' || privateKeyIndex < 1) {
  throw new Error('pKeyIndex must be a positive integer')
}

const privateKeyName = `PRIVATE_KEY_${privateKeyIndex}`
const privateKey: `0x${string}` = config[privateKeyName] as `0x${string}`

if (!privateKey) {
  throw new Error(`${privateKeyName} is not defined in the config file`)
}

// construct enriched versions of viem chain objects ("vChains") based on what was supplied in config file
const vChains: any = {}
Object.entries(config.CHAINS).forEach(
  ([chainId, chainConfig]: [string, any]) => {
    const viemChain = Object.values(viemChains).find(
      (chain) => chain.id === parseInt(chainId, 10)
    )
    if (!viemChain) {
      throw new Error(
        `No viem chain config found for chain ID ${chainId}. Bump viem version or add manually`
      )
    }

    vChains[chainId] = {
      ...viemChain,
      ...chainConfig,
      vCliRead: {} as PublicClient,
      vCliSim: {} as WalletClient,
      vCliWrite: {} as WalletClient,
    }
  }
)

const nonceManager = createNonceManager({
  source: jsonRpc(),
})

const walletAccount = privateKeyToAccount(privateKey, { nonceManager })

print(`Using ${privateKeyName}: ${walletAccount.address}`)

Object.keys(vChains).forEach((chainId: string) => {
  const chain = vChains[chainId]

  chain.vCliRead = createPublicClient({
    chain,
    transport: http(chain.rpcUrl_Read),
  }) as PublicClient

  chain.vCliSim = createWalletClient({
    chain,
    transport: http(chain.rpcUrl_Sim),
  }).extend(publicActions)

  chain.vCliWrite = createWalletClient({
    chain,
    transport: http(chain.rpcUrl_Write),
  }).extend(publicActions)

  Promise.all([
    chain.vCliRead.readContract({
      address: chain.FastRouterAddr,
      abi: ABI.fastRouterV2,
      functionName: 'fastBridge',
    }),
    // not just used to report block height. also serves as connectivity test for all three clients.
    chain.vCliRead.getBlockNumber(),
    chain.vCliSim.getBlockNumber(),
    chain.vCliWrite.getBlockNumber(),
  ]).then(
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    ([fastBridgeAddr, blockNumber_Read, blockNumber_Sim, blockNumber_Write]: [
      string,
      bigint,
      bigint,
      bigint
    ]) => {
      print(
        `Connected to chain ID ${chainId
          .toString()
          .padStart(7)}. FastBridge at: ${fastBridgeAddr.slice(
          0,
          6
        )}... Current block height: ${blockNumber_Read}`
      )
    }
  )
})

let cachedResponseRFQ: any

let sendCounter = 0

let lastAction: string = 'none'

const testBridgeAmountUnits = config.TEST_BRIDGE_AMOUNT_UNITS

mainFlow()

async function mainFlow() {
  await delay(1500)

  while (!walletAccount.address) {
    print(`%ts Awaiting Initialization...`)
    await delay(5000)
  }

  checkBals()

  looper_getRequestParams()

  await delay(2500)

  while (!cachedResponseRFQ) {
    print(`%ts Awaiting cached RFQ API response to populate...`)
    await delay(5000)
  }

  bridgeLooper()
}

async function checkBals() {
  for (;;) {
    await Promise.all(
      Object.keys(vChains).map(async (chainId: string) => {
        const chain = vChains[chainId]
        try {
          const balance = await chain.vCliRead.getBalance({
            address: walletAccount.address,
          })
          chain.balanceRaw = balance
          chain.balanceUnits = formatUnits(balance, 18)
        } catch (error: any) {
          print(
            `Error fetching balance for chain ID ${chainId}: ${error.message}`
          )
        }
      })
    )

    await delay(15_000)
  }
}

const minGasUnits = config.MINIMUM_GAS_UNITS
const rebalToUnits = config.REBALANCE_TO_UNITS
async function bridgeLooper() {
  for (;;) {
    // Find the chain with the lowest balance below our minimum gas -- if any
    const rebalToChain: any = Object.values(vChains).find(
      (chain: any) => chain.balanceUnits < minGasUnits
    )

    if (rebalToChain) {
      const rebalFromChain: any = Object.values(vChains).reduce(
        (prev: any, current: any) => {
          return prev.balanceUnits > current.balanceUnits ? prev : current
        }
      )

      const rebalLabel = `%ts Rebal: ${rebalFromChain.id} > ${rebalToChain.id}`

      print(rebalLabel)

      let retryCount = 0
      // avoid repeating rebal actions. just loop until it lands on-chain.
      if (lastAction === `rebal${rebalFromChain.id}>${rebalToChain.id}`) {
        print(
          `${rebalLabel} Last action was identical (${lastAction}). Not repeating. Re-evaluating momentarily...`
        )

        if (retryCount > 5) {
          // abort after X attempts - if running in repeater mode this will effectively re-send the rebal tx if it is still needed
          print(`${rebalLabel} Max retries. Exiting process...`)
          await delay(1500)
          process.exit()
        }

        await delay(7500)
        retryCount++
        continue
      }

      // leave rebalFrom chain with X units
      const rebalAmount = rebalToUnits - rebalToChain.balanceUnits

      if (rebalFromChain.balanceUnits < rebalToUnits * 1.1) {
        // if we hit this point, it indicates the wallet has no funds left to keep playing. hang process.
        print(
          `${rebalLabel} - Insuff Funds on From Chain ${rebalFromChain.balanceUnits}. Ending tests.`
        )
        await delay(60_000)
        return
      }

      await sendBridge(
        rebalFromChain,
        rebalToChain,
        Number(rebalAmount.toFixed(18)),
        false,
        rebalLabel
      )

      lastAction = `rebal${rebalFromChain.id}>${rebalToChain.id}`

      await delay(config.VOLLEY_MILLISECONDS_BETWEEN)
      continue
    }

    let fromChain: any
    let toChain: any

    const totalPercentage: number = Object.values(config.TEST_ROUTES).reduce(
      (acc: number, route: any) => acc + route.testDistributionPercentage,
      0
    )
    const randomizer: number = getRandomInt(1, totalPercentage)
    let cumulative = 0

    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    for (const [route, details] of Object.entries(config.TEST_ROUTES) as [
      string,
      any
    ][]) {
      cumulative += details.testDistributionPercentage
      if (randomizer <= cumulative) {
        fromChain = vChains[`${details.fromChainId}`]
        toChain = vChains[`${details.toChainId}`]
        break
      }
    }

    const countToSend = getRandomInt(
      config.VOLLEY_MIN_COUNT,
      config.VOLLEY_MAX_COUNT
    )
    const printLabel = `%ts Batch${(sendCounter + countToSend)
      .toString()
      .padStart(5, '0')} of ${countToSend} : ${fromChain.id
      .toString()
      .padStart(7)} >> ${toChain.id.toString().padEnd(7)}`

    for (let i = 0; i < countToSend; i++) {
      // sendCounter is applied as a tag on the amount just for sloppy tracking purposes. not actually important.
      sendBridge(
        fromChain,
        toChain,
        Number(
          (testBridgeAmountUnits + sendCounter / 100000000000).toFixed(18)
        ),
        true,
        printLabel
      )
      sendCounter++
      await delay(50)
    }

    lastAction = 'testVolley'
    await delay(config.VOLLEY_MILLISECONDS_BETWEEN)
  }
}

async function getRequestParams(
  fromChain: any,
  toChain: any,
  sendAmountUnits: number
) {
  // 480 is not supported currently on the API - using Opti/Base as proxies. Can be improved later as needed.
  if (toChain.id === 480) {
    toChain = fromChain.id === 8453 ? vChains['10'] : vChains['8453']
  }
  if (fromChain.id === 480) {
    fromChain = toChain.id === 8453 ? vChains['10'] : vChains['8453']
  }

  const requestURL = `https://api.synapseprotocol.com/bridge?fromChain=${
    fromChain.id
  }&toChain=${
    toChain.id
  }&fromToken=${'0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'}&toToken=${'0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE'}&amount=${sendAmountUnits}&destAddress=${
    walletAccount.address
  }`

  let responseRFQ: any
  let response: any
  try {
    response = await axios.request({
      url: requestURL,
    })

    if ((response.data?.length ?? 0) === 0) {
      throw new Error(`No data returned from api`)
    }

    responseRFQ = Array.isArray(response.data)
      ? response.data.find(
          (item: any) => item.bridgeModuleName === 'SynapseRFQ'
        )
      : null
    if (!responseRFQ) {
      throw new Error(`No RFQ response returned from api`)
    }
  } catch (error: any) {
    throw new Error(
      `RFQ Api Fail: ${error.message.substr(0, 50)} -- ${requestURL}`
    )
  }

  return responseRFQ
}

async function looper_getRequestParams() {
  for (;;) {
    // in future iteration, this could be improved to dynamically pull a response for each route that is involved w/ testing.
    // for now, all tests just use a cached route btwn two OP stacks as proxies for Worldchain tests -- because this is close enough.
    cachedResponseRFQ = await getRequestParams(
      vChains['8453'],
      vChains['10'],
      testBridgeAmountUnits
    )

    await delay(10_000)
  }
}

async function sendBridge(
  fromChain: any,
  toChain: any,
  sendAmountUnits: number,
  useCachedRequest: boolean,
  printLabel: string
) {
  const sendAmountRaw = parseUnits(sendAmountUnits.toString(), 18)

  printLabel = printLabel + ` ${sendAmountUnits} ETH`

  const _responseRFQ = useCachedRequest
    ? cachedResponseRFQ
    : await getRequestParams(fromChain, toChain, sendAmountUnits)

  const contractCall: any = {
    address: fromChain.FastRouterAddr as `0x${string}`,
    abi: ABI.fastRouterV2,
    functionName: 'bridge',
    account: walletAccount,
    chain: fromChain,
    args: [
      walletAccount.address, //recipient
      toChain.id, // chainId
      '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE', // token
      sendAmountRaw, // amount
      [
        //originQuery
        '0x0000000000000000000000000000000000000000', //routerAdapter
        _responseRFQ.originQuery.tokenOut, //tokenOut
        sendAmountRaw, //minAmountOut
        BigInt(_responseRFQ.originQuery.deadline.hex), //deadline
        _responseRFQ.originQuery.rawParams, //rawParms
      ],
      [
        //destQuery
        '0x0000000000000000000000000000000000000000', //routerAdapter
        _responseRFQ.destQuery.tokenOut, //tokenOut
        BigInt(_responseRFQ.destQuery.minAmountOut.hex), //minAmountOut
        BigInt(_responseRFQ.destQuery.deadline.hex), //deadline
        _responseRFQ.destQuery.rawParams, //rawParms
      ],
    ],
    value: sendAmountRaw,
  }

  let estGasUnits

  try {
    //@ts-ignore
    estGasUnits = await fromChain.vCliSim.estimateContractGas(contractCall)
  } catch (error: any) {
    throw new Error(`${printLabel} Bridge Sim error: ${error.message}`)
  }

  if (estGasUnits <= 50000n) {
    throw new Error(`${printLabel} estimated gas units too low. possible error`)
  }

  contractCall.gas = Math.floor(Number(estGasUnits) * 1.3)

  let txHash
  try {
    txHash = await fromChain.vCliWrite.writeContract(contractCall)
  } catch (error: any) {
    throw new Error(`${printLabel} Send failed: ${error.message}`)
  }

  print(`${printLabel} Submitted ${txHash}`)
}
