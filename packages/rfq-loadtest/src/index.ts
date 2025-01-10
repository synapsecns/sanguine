import fs from 'fs/promises'

import * as viemChains from 'viem/chains'
import { privateKeyToAccount } from 'viem/accounts'
import {
  Address,
  ContractFunctionExecutionError,
  createPublicClient,
  createWalletClient,
  decodeErrorResult,
  decodeFunctionData,
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

  Object.entries(config.ASSETS).forEach(([asset, chains]: [string, any]) => {
    if (typeof chains !== 'object' || chains === null) {
      throw new Error(`Invalid ASSETS configuration for asset: ${asset}`)
    }
    Object.entries(chains).forEach(([chainId, chainDetails]: [string, any]) => {
      if (typeof chainDetails !== 'object' || chainDetails === null) {
        throw new Error(`Invalid ASSETS>CHAINS configuration under ${asset}>${chainId}`)
      }
      if (typeof chainDetails.TEST_BRIDGE_AMOUNT_UNITS !== 'number') {
        throw new Error(`Invalid configuration for TEST_BRIDGE_AMOUNT_UNITS under ${asset}>${chainId}`)
      }
    })
  })

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
const privateKey_bridger: Address = config[privateKeyName] as Address

const privateKey_relayer: Address = config['RELAYER_PRIVATE_KEY'] as Address

if (!privateKey_bridger) {
  throw new Error(`${privateKeyName} is not defined in the config file`)
}


if (!privateKey_relayer && config.RELAYER_TEST) {
  throw new Error(`RELAYER_TEST = ${config.RELAYER_TEST}, but no RELAYER_PRIVATE_KEY supplied.`)
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
      assets: Object.keys(config.ASSETS).reduce((acc: any, asset: string) => {
        if (config.ASSETS[asset][chainId]) {
          acc[asset] = { ...config.ASSETS[asset][chainId], label: asset }
        }
        return acc
      }, {})
    }
  }
)


const account_bridger = privateKeyToAccount(privateKey_bridger, { nonceManager: createNonceManager({ source: jsonRpc() })})

const account_relayer = privateKeyToAccount(privateKey_relayer, { nonceManager: createNonceManager({ source: jsonRpc() })})

if(account_relayer) print(`Using Relayer ${account_relayer.address}`)

print(`Using Bridger ${account_bridger.address} (Key #${privateKeyIndex})`)

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
      // Write fastBridgeAddr back to the chain object
      chain.fastBridgeAddr = fastBridgeAddr;

      print(
        `Connected to chain ID ${chainId.toString().padStart(7)}... Current block height: ${blockNumber_Read}`
      )
    }
  )
})

const testRoutes:any = []

//construct all possible routes from all available chains & assets
for (const [fromAsset, fromChains] of Object.entries(config.ASSETS) as [any, any]) {
  for (const [toAsset, toChains] of Object.entries(config.ASSETS) as [any, any]) {
    for (const [fromChainId, fromChainDetails] of Object.entries(fromChains)) {
      for (const [toChainId, toChainDetails] of Object.entries(toChains)) {
        if (fromChainId !== toChainId) {
          testRoutes.push({
            fromChain: vChains[fromChainId],
            toChain: vChains[toChainId],
            fromAsset: vChains[fromChainId].assets[fromAsset],
            toAsset: vChains[toChainId].assets[toAsset],
          })
        }
      }
    }
    
  }
}

let sendCounter = 0

mainFlow()

async function mainFlow() {
  await delay(1500)

  while (!account_bridger.address) {
    print(`%ts Awaiting Initialization...`)
    await delay(5000)
  }

  loopBridges()
}

async function erc20_allowance_check_and_set(chain: any, asset: any, approveSpenderAddress: Address)
{
    // Check allowance of target contract to spend the token for bridgeAccount
    const allowance = await chain.vCliRead.readContract({
      address: asset.TOKEN_ADDRESS,
      abi: ABI.erc20min,
      functionName: 'allowance',
      args: [account_bridger.address, approveSpenderAddress]
    });

    if (allowance < BigInt('0xffffffffffffffffffffffffffffffffffffffffffff')) {
      try {
        const txHash = await chain.vCliWrite.writeContract({
          address: asset.TOKEN_ADDRESS,
          abi: ABI.erc20min,
          functionName: 'approve',
          args: [approveSpenderAddress, BigInt('0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff')],
          account: account_bridger
        });

        const printLabel = `%ts Allowing ${asset.label} spender ${approveSpenderAddress.slice(0,8)}...${approveSpenderAddress.slice(-4)} on Chain ${chain.id} Tx: ${txHash}`;
        print(`${printLabel} ...`);

        await chain.vCliRead.waitForTransactionReceipt({ hash: txHash });
        print(`${printLabel} SUCCESS`);

      } catch (error: any) {
        throw new Error(`Error Allowing ${asset.label} spender ${approveSpenderAddress.slice(0,8)}...${approveSpenderAddress.slice(-4)} on Chain ${chain.id} Err: ${error.message}`);
      }
    }
}

async function checkBals() {
    await Promise.all(
      Object.keys(config.ASSETS).map(async (asset) => {
        const assetChains = config.ASSETS[asset];
        await Promise.all(
          Object.keys(assetChains).map(async (chainId) => {
            const chain = vChains[chainId];
            const mappedAsset:any= chain.assets[asset]
            try {
              let balance;
              if (assetChains[chainId].TOKEN_ADDRESS === '0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE') {
                // Native asset, use getBalance
                balance = await chain.vCliRead.getBalance({
                  address: account_bridger.address,
                });
              } else {
                // ERC20 token, use balanceOf
                balance = await chain.vCliRead.readContract({
                  address: assetChains[chainId].TOKEN_ADDRESS,
                  abi: ABI.erc20min,
                  functionName: 'balanceOf',
                  args: [account_bridger.address]
                });

                if(mappedAsset.balanceRaw==undefined) //not initialize = we can perform our single start-up check & set for allowances. will not run on each invoke.
                {
                    await erc20_allowance_check_and_set(chain, mappedAsset, config.CHAINS[chainId].FastRouterAddr)
                    await erc20_allowance_check_and_set(chain, mappedAsset, config.CHAINS[chainId].SinRouterAddr)
                }
              }
              mappedAsset.balanceRaw = balance;
              mappedAsset.balanceUnits = formatUnits(balance, mappedAsset.DECIMALS);
            } catch (error: any) {
              print(
                `checkBals error for asset ${asset} on chain ID ${chainId}: ${error.message}`
              );
              return
            }
          })
        );
      })
    );
}


async function falseProve(printLabel: string, txHash_bridgePromise: Promise<Address>, fromChain: any)
{
      const txHash_bridge = await txHash_bridgePromise

      printLabel += ' falseProve'

        var receipt_bridge

        try {
          receipt_bridge = await fromChain.vCliRead.waitForTransactionReceipt({hash: txHash_bridge})
        } catch (error:any) {
          print(`${printLabel} Error waiting for bridge transaction receipt: ${error.message}`)
          return
        }

        // bridge tx confirmed. prepare & send a false proof for it.

        // 2nd topic of 1st log = synapse transaction Id
        const transaction_id = receipt_bridge.logs[0].topics[1]

        const contractCall = {
          address: fromChain.fastBridgeAddr as Address,
          abi: ABI.fastBridgeV2,
          functionName: 'proveV2',
          args: [
            transaction_id, // transaction_id
            '0xabcdef123456abcdef123456abcdef123456abcdef123456abcdef123456abcd', //destTxHash
            account_relayer.address // assert self as relayer
          ],
          account: account_relayer
        }

        await delay(3_000) // brief delay for bridge to settle before sending proof

        try {
          const txHash_falseProve = await fromChain.vCliWrite.writeContract(contractCall)
          print(`${printLabel} Sent: ${txHash_falseProve}`)
        } catch (error: any) {
          print(`${printLabel} Error sending false prove transaction: ${error.message}`)
        }
        
}

async function loopBridges() {
  
  await delay(2_000)

  for (;;) {

    await checkBals()

    // only consider tests where we actually have sufficient balance of the input token to execute it
    const validRoutes = testRoutes.filter((route:any) => {
      const fromAssetBalance = route.fromChain.assets[route.fromAsset.label].balanceUnits;
      const minimumUnits = route.fromChain.assets[route.fromAsset.label].TEST_BRIDGE_AMOUNT_UNITS;
      return fromAssetBalance > (minimumUnits * 1.01); // require slightly more than the bare minimum
    });

    if (validRoutes.length === 0) {
      print('No test routes with sufficient balance found. Unable to perform any tests.');
      return;
    }

    const randomIndex = getRandomInt(0, validRoutes.length - 1);
    const testRoute = validRoutes[randomIndex];

    const countToSend = getRandomInt(
      config.VOLLEY_MIN_COUNT,
      config.VOLLEY_MAX_COUNT
    )

    const batchLabel = config.VOLLEY_MAX_COUNT==1 ? `` : `Batch${(sendCounter + countToSend).toString().padStart(3, '0')} of ${countToSend} : `


    for (let i = 0; i < countToSend; i++) {
      // add a random dust amount to help unq identify each bridge
      const dustFactor = 1 + Math.random() * 0.00089 + 0.00001;
      const sendAmountUnits = Number((testRoute.fromAsset.TEST_BRIDGE_AMOUNT_UNITS * dustFactor).toFixed(testRoute.fromAsset.DECIMALS));

      const printLabel = `%ts ${batchLabel}` +
      `${sendAmountUnits.toString().padEnd(20)} ` +
      `${testRoute.fromChain.label.slice(0,4).padStart(4)}.${testRoute.fromAsset.label.slice(0,6).padEnd(6)}` +
      ` ► ` +
      `${testRoute.toChain.label.slice(0,4).padStart(4)}.${testRoute.toAsset.label.slice(0,6).padEnd(6)}`


      // sendCounter is applied as a tag on the amount just for sloppy tracking purposes. not actually important.
      let txHash_bridge;
      try {
        txHash_bridge = await sendBridge(
          testRoute,
          sendAmountUnits,
          printLabel
        );
      } catch (error: any) {
        print(`${printLabel} sendBridge failed: 
        ${error.message}`);
        continue; // Skip to the next iteration if there's an error
      }
      sendCounter++
      await delay(50)

      // if we're running a relayer test, wait for bridge to land and submit relayer test txn
      if(config.RELAYER_TEST)
      {
        if(config.RELAYER_TEST.toLowerCase() == 'falseprove') await falseProve(printLabel, txHash_bridge, testRoute.fromChain)
      }
    }

    await delay(config.VOLLEY_MILLISECONDS_BETWEEN)
  }
}

async function getRequestParams(
  route: any,
  sendAmountUnits: number
) {

  const requestURL = 
    `${config.BRIDGE_API_URL}/bridge?` +
    `fromChain=${route.fromChain.id}` +
    `&toChain=${route.toChain.id}` +
    `&fromToken=${route.fromAsset.TOKEN_ADDRESS}&toToken=${route.toAsset.TOKEN_ADDRESS}` +
    `&amount=${sendAmountUnits}` +
    `&originUserAddress=${account_bridger.address}` +
    `&destAddress=${account_bridger.address}`

  let responseRFQ: any
  let response: any
  try {
    response = await axios.request({
      url: requestURL,
    })

    if ((response.data?.length ?? 0) === 0) {
      throw new Error(`No data returned from API`)
    }

    if (Array.isArray(response.data)) {
      // find response w/ bridge module that matches those which we are testing
      const matchingModules = response.data.filter((item: any) =>
        config.BRIDGE_MODULES.includes(item.bridgeModuleName)
      );
      // if mult candidates found, randomly pick one
      if (matchingModules.length > 0) {
        const randomIndex = Math.floor(Math.random() * matchingModules.length);
        responseRFQ = matchingModules[randomIndex];
      } else {
        responseRFQ = null;
      }
    } else {
      responseRFQ = null;
    }
    if (!responseRFQ) {
      throw new Error(`No usable BridgeModule found on API response.`)
    }
  } catch (error: any) {
    throw new Error(
      `RFQ Api Fail: ${error.message.substr(0, 50)} -- ${requestURL}`
    )
  }

  return responseRFQ
}

async function sendBridge(
  route:any,
  sendAmountUnits: number,
  printLabel: string
) {

  let _responseRFQ;
  try {
    _responseRFQ = await getRequestParams(route, sendAmountUnits);
  } catch (error: any) {
    throw new Error(`getRequestParams fail: ${error.message}`);
  }

  const sendAmountRaw = parseUnits(sendAmountUnits.toString(), route.fromAsset.DECIMALS);

  const bridgeModuleLabel = _responseRFQ.bridgeModuleName.replace('SynapseIntents', 'SIN').replace('SynapseRFQ', 'RFQ')

  printLabel = printLabel + ` ${bridgeModuleLabel}`

  let estGasUnits

  let txData:any = {
    to: _responseRFQ.callData.to,
    data: _responseRFQ.callData.data,
    account: account_bridger,
    chain: route.fromChain,
    value: route.fromAsset.TOKEN_ADDRESS=='0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE' ? sendAmountRaw : 0n
  }

  try {
    estGasUnits = await route.fromChain.vCliSim.estimateGas(txData)
      
    if (estGasUnits <= 50000n) throw new Error(`Gas units too low ${estGasUnits}. Sim failure?`)

  } catch (error: any) {

    var errMsg = error.cause.cause.cause.cause.cause.data ?? error.cause.details ?? error.details ?? error.message

    if(errMsg.slice(0,2)=='0x') 
      {
        try {
          errMsg = decodeErrorResult({abi: ABI.fastRouterV2, data: errMsg});
        } catch (decodeError) {
          errMsg = `${errMsg} (could not decode)`;
        }
      }

    print (`${printLabel} Bridge Sim error: ${errMsg}`)
    return
  }

  txData.gas = Math.floor(Number(estGasUnits) * 1.3)

  if(config.DEBUG_MODE==true)
  {
    print(`${printLabel} DEBUG MODE - Not actually submitted to chain`)
    return
  }

  let txHash
  try {
    txHash = await route.fromChain.vCliWrite.sendTransaction(txData)
  } catch (error: any) {
    throw new Error(`${printLabel} Send failed: ${error.message}`)
  }

  print(`${printLabel} Submitted ${txHash}`)

  return txHash
}

