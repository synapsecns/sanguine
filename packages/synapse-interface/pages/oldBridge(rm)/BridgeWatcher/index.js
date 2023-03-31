import { useCallback, useEffect } from 'react'
import { useWeb3React } from '@web3-react/core'
import { id } from '@ethersproject/hash'
import { BigNumber } from '@ethersproject/bignumber'
import bech32 from 'bech32'
import { toHexStr } from '@utils/toHexStr'
import {
  AGEUR,
  DOG,
  GMX,
  GOHM,
  H2O,
  HIGHSTREET,
  JUMP,
  L2DAO,
  LINK,
  NEWO,
  NFD,
  PLS,
  SDT,
  SFI,
  UNIDX,
  USDB,
  VSTA,
} from '@constants/tokens/mintable'
import {
  DOGECHAIN_BUSD,
  FRAX,
  KLAYTN_DAI,
  KLAYTN_USDC,
  KLAYTN_USDT,
  KLAYTN_WETH,
  NETH,
  NOTE,
  NUSD,
  SYN,
  SYN_FRAX,
  TOKEN_HASH_MAP,
  UST,
  WBTC,
  WETH,
} from '@constants/tokens/basic'
import { INCLUDED_BRIDGE_EVENTS } from '@constants/events'
import { ChainId } from '@constants/networks'
import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'
import { useTxHistory } from '@hooks/store/useTxHistory'
import { useGenericSynapseContract } from '@hooks/contracts/useSynapseContract'
import { usePoller } from '@hooks/usePoller'
import Grid from '@tw/Grid'
import Card from '@tw/Card'
import { useDestinationInfo } from '@hooks/store/useDestinationInfo'
import { useGetTerraBridgeEvent } from '@hooks/terra/useGetTerraBridgeEvent'
import { useTerraWallet } from '@hooks/terra/useTerraWallet'
import { useGetTerraBridgeEvents } from '@hooks/terra/useGetTerraBridgeEvents'

import PairedTransactionItem from './PairedTransactionItem'
import { pairTxKappa } from './pairTxKappa'

export default function BridgeWatcher() {
  const { account, chainId } = useActiveWeb3React()

  const { library: ethLibrary } = useWeb3React(`${ChainId.ETH}`)
  const { library: bscLibrary } = useWeb3React(`${ChainId.BSC}`)
  const { library: polygonLibrary } = useWeb3React(`${ChainId.POLYGON}`)
  const { library: fantomLibrary } = useWeb3React(`${ChainId.FANTOM}`)
  // const { library: bobaLibrary }      = useWeb3React(`${ChainId.BOBA}`)
  const { library: moonbeamLibrary } = useWeb3React(`${ChainId.MOONBEAM}`)
  const { library: moonriverLibrary } = useWeb3React(`${ChainId.MOONRIVER}`)
  const { library: arbitrumLibrary } = useWeb3React(`${ChainId.ARBITRUM}`)
  const { library: avalancheLibrary } = useWeb3React(`${ChainId.AVALANCHE}`)
  const { library: harmonyLibrary } = useWeb3React(`${ChainId.HARMONY}`)
  const { library: optimismLibrary } = useWeb3React(`${ChainId.OPTIMISM}`)
  const { library: cronosLibrary } = useWeb3React(`${ChainId.CRONOS}`)
  const { library: dfkLibrary } = useWeb3React(`${ChainId.DFK}`)
  const { library: metisLibrary } = useWeb3React(`${ChainId.METIS}`)
  const { library: klaytnLibrary } = useWeb3React(`${ChainId.KLAYTN}`)
  const { library: dogechainLibrary } = useWeb3React(`${ChainId.DOGECHAIN}`)
  // const { library: cantoLibrary } = useWeb3React(`${ChainId.CANTO}`)
  // const { library: auroraLibrary }    = useWeb3React(`${ChainId.AURORA}`)

  const { transactions, updateTransactions, clear } = useTxHistory()

  const { terraAddress } = useTerraWallet()

  const ethSynapseContract = useGenericSynapseContract(ChainId.ETH)
  const bscSynapseContract = useGenericSynapseContract(ChainId.BSC)
  const polygonSynapseContract = useGenericSynapseContract(ChainId.POLYGON)
  const fantomSynapseContract = useGenericSynapseContract(ChainId.FANTOM)
  // const bobaSynapseContract      = useGenericSynapseContract(ChainId.BOBA)
  const moonbeamSynapseContract = useGenericSynapseContract(ChainId.MOONBEAM)
  const moonriverSynapseContract = useGenericSynapseContract(ChainId.MOONRIVER)
  const arbitrumSynapseContract = useGenericSynapseContract(ChainId.ARBITRUM)
  const avalancheSynapseContract = useGenericSynapseContract(ChainId.AVALANCHE)
  const harmonySynapseContract = useGenericSynapseContract(ChainId.HARMONY)
  const optimismSynapseContract = useGenericSynapseContract(ChainId.OPTIMISM)
  const cronosSynapseContract = useGenericSynapseContract(ChainId.CRONOS)
  const dfkSynapseContract = useGenericSynapseContract(ChainId.DFK)
  const metisSynapseContract = useGenericSynapseContract(ChainId.METIS)
  const klaytnSynapseContract = useGenericSynapseContract(ChainId.KLAYTN)
  //  const dogechainSynapseContract     = useGenericSynapseContract(ChainId.DOGECHAIN)
  // const cantoSynapseContract     = useGenericSynapseContract(ChainId.CANTO)
  // const auroraSynapseContract    = useGenericSynapseContract(ChainId.AURORA)

  const [addressesForAccount, setAddressesForAccount] = useDestinationInfo()

  const getTerraBridgeEvents = useGetTerraBridgeEvents()
  // console.log({terraEventObj})

  const getAllTransactions = useCallback(
    _.throttle(
      () => {
        const sharedArgs = {
          account,
          terraAddress,
          transactions,
          addressesForAccount,
        }

        getBridgeZapsOnChain({
          library: ethLibrary,
          synapseContract: ethSynapseContract,
          chainId: ChainId.ETH,
          ...sharedArgs,
        }).then(updateTransactions)

        getBridgeZapsOnChain({
          library: bscLibrary,
          synapseContract: bscSynapseContract,
          chainId: ChainId.BSC,
          ...sharedArgs,
        }).then(updateTransactions)

        getBridgeZapsOnChain({
          library: polygonLibrary,
          synapseContract: polygonSynapseContract,
          chainId: ChainId.POLYGON,
          ...sharedArgs,
        }).then(updateTransactions)

        getBridgeZapsOnChain({
          library: fantomLibrary,
          synapseContract: fantomSynapseContract,
          chainId: ChainId.FANTOM,
          ...sharedArgs,
        }).then(updateTransactions)

        // getBridgeZapsOnChain({
        //   library:         bobaLibrary,
        //   synapseContract: bobaSynapseContract,
        //   chainId:         ChainId.BOBA,
        //   ...sharedArgs
        // }).then(updateTransactions)

        getBridgeZapsOnChain({
          library: moonbeamLibrary,
          synapseContract: moonbeamSynapseContract,
          chainId: ChainId.MOONBEAM,
          ...sharedArgs,
        }).then(updateTransactions)

        getBridgeZapsOnChain({
          library: moonriverLibrary,
          synapseContract: moonriverSynapseContract,
          chainId: ChainId.MOONRIVER,
          ...sharedArgs,
        }).then(updateTransactions)

        getBridgeZapsOnChain({
          library: arbitrumLibrary,
          synapseContract: arbitrumSynapseContract,
          chainId: ChainId.ARBITRUM,
          ...sharedArgs,
        }).then(updateTransactions)

        getBridgeZapsOnChain({
          library: optimismLibrary,
          synapseContract: optimismSynapseContract,
          chainId: ChainId.OPTIMISM,
          ...sharedArgs,
        }).then(updateTransactions)

        getBridgeZapsOnChain({
          library: avalancheLibrary,
          synapseContract: avalancheSynapseContract,
          chainId: ChainId.AVALANCHE,
          ...sharedArgs,
        }).then(updateTransactions)

        getBridgeZapsOnChain({
          library: harmonyLibrary,
          synapseContract: harmonySynapseContract,
          chainId: ChainId.HARMONY,
          ...sharedArgs,
        }).then(updateTransactions)

        getBridgeZapsOnChain({
          library: cronosLibrary,
          synapseContract: cronosSynapseContract,
          chainId: ChainId.CRONOS,
          ...sharedArgs,
        }).then(updateTransactions)

        getBridgeZapsOnChain({
          library: dfkLibrary,
          synapseContract: dfkSynapseContract,
          chainId: ChainId.DFK,
          ...sharedArgs,
        }).then(updateTransactions)

        getBridgeZapsOnChain({
          library: metisLibrary,
          synapseContract: metisSynapseContract,
          chainId: ChainId.METIS,
          ...sharedArgs,
        }).then(updateTransactions)

        getBridgeZapsOnChain({
          library: klaytnLibrary,
          synapseContract: klaytnSynapseContract,
          chainId: ChainId.KLAYTN,
          ...sharedArgs,
        }).then(updateTransactions)

        //getBridgeZapsOnChain({
        // library:         dogechainLibrary,
        // synapseContract: dogechainSynapseContract,
        // chainId:         ChainId.DOGECHAIN,
        // ...sharedArgs
        //}).then(updateTransactions)

        // getBridgeZapsOnChain({
        //   library:         cantoLibrary,
        //   synapseContract: cantoSynapseContract,
        //   chainId:         ChainId.CANTO,
        //   ...sharedArgs
        // }).then(updateTransactions)

        getTerraBridgeEvents({ account, terraAddress }).then(updateTransactions)
      },
      10000,
      { leading: true }
    )
  )

  useEffect(getAllTransactions, [chainId, account])

  usePoller(getAllTransactions, 10000)

  return (
    <div className="space-y-2">
      {transactions?.length > 0 && (
        <Card title="Bridge Watcher" divider={false}>
          <SynapseEvents transactions={transactions} />
        </Card>
      )}
    </div>
  )
}

function SynapseEvents({ transactions }) {
  const pairedTransactions = pairTxKappa(transactions)

  return (
    <Grid cols={{ xs: 1 }} gap={2}>
      {pairedTransactions.map(([inputTx, outputTx]) => {
        return (
          <PairedTransactionItem
            key={`${inputTx?.transactionHash}-${outputTx?.transactionHash}`}
            inputTx={inputTx}
            outputTx={outputTx}
          />
        )
      })}
    </Grid>
  )
}

async function getBridgeZapsOnChain({
  library,
  synapseContract,
  account,
  terraAddress,
  chainId,
  transactions,
  addressesForAccount,
}) {
  const currentBlock = await library.getBlock()
  // console.log(synapseContract.filters)
  const synapseContractEventHandles = Object.keys(
    synapseContract.filters
  ).filter((key) => INCLUDED_BRIDGE_EVENTS.includes(key))

  const evmAddrsForAccount = (addressesForAccount ?? []).filter(
    (s) => s.slice(0, 2) == '0x'
  )
  const terraAddrsForAccount = (addressesForAccount ?? []).filter(
    (s) => s.slice(0, 6) == 'terra'
  )

  const pastEventsByHandle = await Promise.all(
    synapseContractEventHandles.map((eventHandle) => {
      let accountList
      if (eventHandle == 'TokenRedeemV2' && terraAddress) {
        const terraAddrs = [...terraAddrsForAccount, terraAddress].map(
          (addr) => bech32.decode(addr).words
        )
        accountList = [account, ...evmAddrsForAccount, ...terraAddrs]
      } else {
        accountList = [account, ...evmAddrsForAccount]
      }

      return synapseContract.queryFilter(
        synapseContract.filters[eventHandle](accountList),
        toHexStr(currentBlock.number - 2040)
      )
    })
  )

  const pastEvents = _.flatten(pastEventsByHandle)

  const [inputTimestamps, transactionReceipts] = await Promise.all([
    Promise.all(
      pastEvents.map((eventObj) => library.getBlock(eventObj.blockNumber))
    ),
    Promise.all(pastEvents.map((eventObj) => eventObj.getTransactionReceipt())),
  ])

  const transactionsWithTimestamps = _.zip(
    pastEvents,
    inputTimestamps,
    transactionReceipts
  ).map(([eventObj, timestampObj, txReceipt]) => {
    return mergeTxObj(chainId, eventObj, timestampObj, txReceipt)
  })

  return transactionsWithTimestamps
}

function mergeTxObj(chainId, eventObj, timestampObj, txReceipt) {
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
      UST,
      KLAYTN_USDT,
      KLAYTN_USDC,
      KLAYTN_DAI,
      WBTC,
      KLAYTN_WETH,
      DOGECHAIN_BUSD,
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
    if (chainId == ChainId.ETH) {
      outputTokenAddr = txReceipt.logs[txReceipt.logs.length - 2].address
    } else {
      outputTokenAddr = txReceipt.logs[txReceipt.logs.length - 1].address
    }
  } else if (swapTokenAddr == _.toLower(NUSD.addresses[chainId])) {
    if (chainId == ChainId.ETH) {
      if (eventObj.event == 'TokenWithdraw') {
        outputTokenAddr = txReceipt.logs[txReceipt.logs.length - 1].address
      } else {
        outputTokenAddr = txReceipt.logs[txReceipt.logs.length - 2].address
      }
    } else if (chainId == ChainId.POLYGON) {
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
      _.toLower(GMX.addresses[ChainId.ARBITRUM]) ||
    _.toLower(txReceipt.logs[1]?.address) ==
      _.toLower(GMX.addresses[ChainId.AVALANCHE])
  ) {
    inputTokenAmount = txReceipt.logs[1].data
  } else if (
    _.toLower(txReceipt.logs[1].address) == _.toLower(UST.addresses[chainId])
  ) {
    inputTokenAmount = txReceipt.logs[2].data
  } else {
    inputTokenAmount = txReceipt.logs[0].data
  }

  let inputTokenAddr
  if (
    _.toLower(txReceipt.logs[1].address) ==
    _.toLower(GMX.addresses[ChainId.AVALANCHE])
  ) {
    inputTokenAddr = _.toLower(GMX.addresses[ChainId.AVALANCHE])
  } else if (
    _.toLower(txReceipt.logs[1].address) == _.toLower(UST.addresses[chainId])
  ) {
    inputTokenAddr = _.toLower(UST.addresses[chainId])
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
    inputTokenAmount,
    inputTokenAddr,
    outputTokenAddr,
    inputToken: TOKEN_HASH_MAP[chainId][inputTokenAddr],
    outputToken: TOKEN_HASH_MAP[chainId][outputTokenAddr],
    txReceipt,
  }
}
