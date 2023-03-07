import _ from 'lodash'

import toast from 'react-hot-toast'

import bech32 from 'bech32'

import { ChainId, CHAIN_INFO_MAP } from '@constants/networks'
import {
  SYN,
  NUSD,
  NETH,
  WETH,
  ETH,
  FRAX,
  WETHE,
  ONEETH,
  WETHBEAM,
  WAVAX,
  WMOVR,
  AVAX,
  MOVR,
  FTMETH,
  METISETH,
  CANTOETH,
  WJEWEL,
  XJEWEL,
  JEWEL,
  SYNJEWEL,
  SYNAVAX,
  DFK_USDC,
  UST,
  MULTIAVAX,
  WBTC,
  KLAYTN_USDT,
  KLAYTN_DAI,
  KLAYTN_USDC,
  KLAYTN_WETH,
  DOGECHAIN_BUSD
} from '@constants/tokens/basic'
import {
  DOG,
  GMX,
  GOHM,
  LINK, HIGHSTREET,
  JUMP,
  NFD,
  NEWO,
  SOLAR,
  SDT,
  USDB,
  VSTA,
  SFI,
  H2O,
  L2DAO,
  PLS,
  AGEUR,
  UNIDX
} from '@constants/tokens/mintable'
import { BRIDGE_SWAPABLE_TYPE_POOLS_BY_CHAIN } from '@constants/tokens/tokenGroups'

import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'
import { useBlockNumber } from '@hooks/useBlockNumber'
import { useTxHistory } from '@hooks/store/useTxHistory'
import { useGetTxArgs } from '@hooks/useGetTxArgs'
import { useAllContracts } from '@hooks/contracts/useAllContracts'
import { useBridgeZapContract } from '@hooks/contracts/useBridgeZapContract'

import { Slippages, subtractSlippage } from '@utils/slippage'

import ExplorerToastLink from '@components/ExplorerToastLink'
import { txErrorHandler } from '@utils/txErrorHandler'

import { validateAndParseAddress } from '@utils/validateAndParseAddress'
import { useDestinationInfo } from '@hooks/store/useDestinationInfo'
import { matchSymbolWithinPool } from '@utils/matchSymbolWithinPool'
import { useSynapseContract } from '@hooks/contracts/useSynapseContract'

import { useTerraWallet } from '@hooks/terra/useTerraWallet'
import { Fee, MsgSend, MsgExecuteContract, Coins } from '@terra-money/terra.js'
import { SYNAPSE_BRIDGE_ADDRESSES } from '@constants/bridge'
import { validateTerraAddress } from '@utils/validateTerraAddress'

export function useBridgeSwap() {
  const bridgeZapContract = useBridgeZapContract()
  const synapseBridgeContract = useSynapseContract()

  const tokenContracts = useAllContracts()
  const { account, chainId } = useActiveWeb3React()
  const { addTransaction } = useTxHistory()
  const [blockNumber, setBlockNumber] = useBlockNumber(chainId)
  const getTxArgs = useGetTxArgs()

  const [addressesForAccount, setAddressesForAccount] = useDestinationInfo()

  const { terraAddress, ...terraWalletRest } = useTerraWallet()

  return async function bridgeSwap({
    destinationAddress,
    fromChainId,
    toChainId,
    fromCoin,
    toCoin,
    fromAmount,
    toAmount,
    deadlineMinutes,
  }) {
    try {
      if (!account) throw new Error('Wallet must be connected')
      if (!bridgeZapContract)
        throw new Error('Bridge Zap contract is not loaded')
      if (fromChainId != ChainId.TERRA && chainId != fromChainId)
        throw new Error('wallet not connected or on diff chain than connected')
      let destAddr
      if (destinationAddress && destinationAddress != '') {
        if (
          toChainId != ChainId.TERRA &&
          validateAndParseAddress(destinationAddress)
        ) {
          destAddr = destinationAddress
        } else if (
          toChainId == ChainId.TERRA &&
          validateTerraAddress(destinationAddress)
        ) {
          destAddr = destinationAddress
        } else {
          throw new Error('Destnation Address is invalid')
        }
      } else {
        if (toChainId == ChainId.TERRA && validateTerraAddress(terraAddress)) {
          destAddr = terraAddress
        } else {
          destAddr = account
        }
      }

      let fromTokenSymbol = fromCoin.symbol
      let toTokenSymbol = toCoin.symbol

      const fromChainName = CHAIN_INFO_MAP[fromChainId].chainName
      const toChainName = CHAIN_INFO_MAP[toChainId].chainName

      const fromChainTokens =
        BRIDGE_SWAPABLE_TYPE_POOLS_BY_CHAIN[fromChainId][fromCoin.swapableType]
          .poolTokens
      const toChainTokens =
        BRIDGE_SWAPABLE_TYPE_POOLS_BY_CHAIN[toChainId][toCoin.swapableType]
          .poolTokens

      const fromSymbol = fromTokenSymbol
      const toSymbol = toTokenSymbol
      if (fromCoin.swapableType == 'ETH') {
        if (fromTokenSymbol == ETH.symbol) {
          fromTokenSymbol = WETH.symbol
        }

        if (toTokenSymbol == ETH.symbol) {
          toTokenSymbol = WETH.symbol
        }
      } else if (fromCoin.swapableType == 'AVAX') {
        if (fromTokenSymbol == AVAX.symbol) {
          fromTokenSymbol = WAVAX.symbol
        }

        if (toTokenSymbol == AVAX.symbol) {
          toTokenSymbol = WAVAX.symbol
        }
      } else if (fromCoin.swapableType == 'JEWEL') {
        if (fromTokenSymbol == JEWEL.symbol) {
          fromTokenSymbol = WJEWEL.symbol
        }

        if (toTokenSymbol == JEWEL.symbol) {
          toTokenSymbol = WJEWEL.symbol
        }
      } else if (fromCoin.swapableType == 'MOVR') {
        if (fromTokenSymbol == MOVR.symbol) {
          fromTokenSymbol = WMOVR.symbol
        }

        if (toTokenSymbol == MOVR.symbol) {
          toTokenSymbol = WMOVR.symbol
        }
      }

      const tokenIndexFrom = fromChainTokens.findIndex((i) =>
        matchSymbolWithinPool(i, fromCoin)
      )
      const tokenIndexTo = toChainTokens.findIndex((i) =>
        matchSymbolWithinPool(i, toCoin)
      )
      // For each token being deposited, check the allowance and approve it if necessary

      const tokenContract = tokenContracts?.[fromTokenSymbol]
      if (tokenContract == null && fromChainId != ChainId.TERRA)
        throw new Error('no token contract loaded')
      const {
        slippageCustom,
        slippageSelected,
        transactionDeadline, // in minutes
        bridgeTransactionDeadline,
      } = getTxArgs({ deadlineMinutes })

      const selectedGasArgs = [slippageSelected, slippageCustom]
      const twoTenthGasArgs = [Slippages.TwoTenth, slippageCustom]
      const quarterGasArgs = [Slippages.Quarter, slippageCustom]

      const minToSwapOrigin = subtractSlippage(fromAmount, ...selectedGasArgs)
      const minToSwapDest = subtractSlippage(toAmount, ...selectedGasArgs)
      const minToSwapDestFromOrigin = subtractSlippage(
        minToSwapDest,
        ...selectedGasArgs
      )

      const minToSwapOriginMediumSlippage = subtractSlippage(
        fromAmount,
        ...twoTenthGasArgs
      )
      const minToSwapDestMediumSlippage = subtractSlippage(
        toAmount,
        ...twoTenthGasArgs
      )
      const minToSwapDestFromOriginMediumSlippage = subtractSlippage(
        minToSwapDestMediumSlippage,
        ...twoTenthGasArgs
      )

      const minToSwapOriginHighSlippage = subtractSlippage(
        fromAmount,
        ...quarterGasArgs
      )
      const minToSwapDestHighSlippage = subtractSlippage(
        toAmount,
        ...quarterGasArgs
      )
      const minToSwapDestFromOriginHighSlippage = subtractSlippage(
        minToSwapDestHighSlippage,
        ...quarterGasArgs
      )

      let bridgeZapSwapTransaction
      if (fromCoin.swapableType == 'UST') {
        if (fromChainId == ChainId.TERRA) {
          bridgeZapSwapTransaction = await terraWalletRest.post({
            // fee: new Fee(1000000, '200000uusd'),
            msgs: [
              new MsgExecuteContract(
                terraAddress,
                SYNAPSE_BRIDGE_ADDRESSES[ChainId.TERRA],
                {
                  deposit_denom: {
                    to_address: destAddr,
                    chain_id: `${toChainId}`,
                    denom: 'uusd',
                  },
                },
                new Coins({ uusd: fromAmount.toNumber() })
              ),
            ],
          })
        } else if (toChainId == ChainId.TERRA) {
          bridgeZapSwapTransaction = await synapseBridgeContract.redeemV2(
            bech32.decode(destAddr).words, // to address
            toChainId, // to chainId
            UST.addresses[fromChainId],
            fromAmount
          )
        } else {
          bridgeZapSwapTransaction = await synapseBridgeContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            UST.addresses[fromChainId],
            fromAmount
          )
        }
      } else if (fromChainId == ChainId.BSC && toTokenSymbol == DOGECHAIN_BUSD.symbol) {
          bridgeZapSwapTransaction = await bridgeZapContract.deposit(
            destAddr, // to address
            toChainId, // to chainId
            DOGECHAIN_BUSD.addresses[fromChainId],
            fromAmount
          )
      } else if (fromChainId == ChainId.ETH) {
        if (toTokenSymbol == SYN.symbol) {
          // This is the new part added while sleep deprived.
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            SYN.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == USDB.symbol) {
          // This is the new part added while sleep deprived.
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            USDB.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == VSTA.symbol) {
          // This is the new part added while sleep deprived.
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            VSTA.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == GOHM.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.deposit(
            destAddr, // to address
            toChainId, // to chainId
            GOHM.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == HIGHSTREET.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.deposit(
            destAddr, // to address
            toChainId, // to chainId
            HIGHSTREET.addresses[fromChainId],
            fromAmount,
          )
        } else if (toTokenSymbol == LINK.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.deposit(
            destAddr, // to address
            toChainId, // to chainId
            LINK.addresses[fromChainId],  
            fromAmount,
          )
        } else if (toTokenSymbol == UNIDX.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.deposit(
            destAddr, // to address
            toChainId, // to chainId
            UNIDX.addresses[fromChainId],  
            fromAmount,
          )
        } else if (toTokenSymbol == DOG.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.deposit(
            destAddr, // to address
            toChainId, // to chainId
            DOG.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == NEWO.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.deposit(
            destAddr, // to address
            toChainId, // to chainId
            NEWO.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == SFI.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.deposit(
            destAddr, // to address
            toChainId, // to chainId
            SFI.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == SDT.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.deposit(
            destAddr, // to address
            toChainId, // to chainId
            SDT.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == H2O.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.deposit(
            destAddr, // to address
            toChainId, // to chainId
            H2O.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == AGEUR.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.deposit(
            destAddr, // to address
            toChainId, // to chainId
            AGEUR.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == KLAYTN_DAI.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.deposit(
            destAddr, // to address
            toChainId, // to chainId
            KLAYTN_DAI.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == KLAYTN_USDC.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.deposit(
            destAddr, // to address
            toChainId, // to chainId
            KLAYTN_USDC.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == KLAYTN_USDT.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.deposit(
            destAddr, // to address
            toChainId, // to chainId
            KLAYTN_USDT.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == WBTC.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.deposit(
            destAddr, // to address
            toChainId, // to chainId
            WBTC.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == FRAX.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.deposit(
            destAddr, // to address
            toChainId, // to chainId
            FRAX.addresses[fromChainId],
            fromAmount
          )
        } else if (
          toTokenSymbol == NETH.symbol ||
          toTokenSymbol == KLAYTN_WETH.symbol
        ) {
          bridgeZapSwapTransaction = await bridgeZapContract.depositETH(
            destAddr, // to address
            toChainId, // to chainId
            fromAmount,
            { value: fromAmount }
          )
        } else if (
          [WETH, WETHE, ONEETH, FTMETH, METISETH, CANTOETH]
            .map((i) => i.symbol)
            .includes(toTokenSymbol)
        ) {
          //toTokenSymbol == WETH.symbol
          bridgeZapSwapTransaction = await bridgeZapContract.depositETHAndSwap(
            destAddr, // to address
            toChainId, // to chainId
            fromAmount,
            0, // tokenIndexFrom for nusd
            tokenIndexTo, // tokenIndexTo + 1,
            minToSwapDestFromOrigin, // minDy
            bridgeTransactionDeadline,
            { value: fromAmount }
          )
        } else if (
          [NUSD, DFK_USDC].map((i) => i.symbol).includes(toTokenSymbol)
        ) {
          // } else if (toTokenSymbol == nusd.symbol) {
          if (fromTokenSymbol == NUSD.symbol) {
            bridgeZapSwapTransaction = await bridgeZapContract.deposit(
              destAddr, // to address
              toChainId, // to chainId
              NUSD.addresses[fromChainId],
              fromAmount
            )
          } else {
            const liquidityAmounts = fromChainTokens.map((t) => {
              if (t.symbol === fromTokenSymbol) {
                return fromAmount
              } else {
                return 0
              }
            })
            bridgeZapSwapTransaction = await bridgeZapContract.zapAndDeposit(
              destAddr, // to address
              toChainId, // to chainId
              NUSD.addresses[fromChainId],
              liquidityAmounts,
              minToSwapDest,
              transactionDeadline
            )
          }
        } else {
          if (fromTokenSymbol == NUSD.symbol) {
            bridgeZapSwapTransaction = await bridgeZapContract.depositAndSwap(
              destAddr,
              toChainId,
              NUSD.addresses[fromChainId],
              fromAmount,
              0, // tokenIndexFrom for nusd
              tokenIndexTo, // tokenIndexTo + 1,
              minToSwapDestFromOriginMediumSlippage, //, minToSwapDestFromOrigin, // minDy
              bridgeTransactionDeadline
            )
          } else {
            const liquidityAmounts = fromChainTokens.map((t) => {
              if (t.symbol === fromTokenSymbol) {
                return fromAmount
              } else {
                return 0
              }
            })
            /** coin on ETH -> coin L2  */
            bridgeZapSwapTransaction =
              await bridgeZapContract.zapAndDepositAndSwap(
                destAddr,
                toChainId,
                NUSD.addresses[fromChainId],
                liquidityAmounts,
                minToSwapOriginMediumSlippage, // minToSwapOrigin,
                transactionDeadline,
                0, // tokenIndexFrom for nusd
                tokenIndexTo, // tokenIndexTo + 1,
                minToSwapDestFromOriginMediumSlippage, //, minToSwapDestFromOrigin, // minDy
                bridgeTransactionDeadline
              )
          }
        }
      } else if (fromChainId == ChainId.DOGECHAIN) {
        if (toTokenSymbol == KLAYTN_DAI.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            KLAYTN_DAI.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == KLAYTN_USDC.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            KLAYTN_USDC.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == KLAYTN_USDT.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            KLAYTN_USDT.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == WBTC.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            WBTC.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == DOGECHAIN_BUSD.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            DOGECHAIN_BUSD.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == NFD.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            NFD.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == WETH.symbol && toChainId == ChainId.ETH) {
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            KLAYTN_WETH.addresses[fromChainId],
            fromAmount
          )
        } else if (fromTokenSymbol == KLAYTN_WETH.symbol) {
          if (
            [WETHE, ONEETH, FTMETH, METISETH, WETH, CANTOETH]
              .map((i) => i.symbol)
              .includes(toTokenSymbol)
          ) {
            // Destination asset is WETH.e / 1ETH / Fantom WETH / Metis WETH / ETH on Ethereum/L2s
            bridgeZapSwapTransaction = await bridgeZapContract.redeemAndSwap(
              destAddr, // to address
              toChainId, // to chainId
              KLAYTN_WETH.addresses[fromChainId],
              fromAmount,
              0, // tokenIndexFrom, // tokenIndexFrom + 1, // tokenIndexFrom
              tokenIndexTo, // tokenIndexTo + 1, //swapTokenIndex
              minToSwapDest,
              transactionDeadline
            )
          } else if (
            [NETH, KLAYTN_WETH].map((i) => i.symbol).includes(toTokenSymbol)
          ) {
            // Destination asset is nETH (most chains) / Klaytn's WETH
            bridgeZapSwapTransaction = await bridgeZapContract.redeem(
              destAddr, // to address
              toChainId, // to chainId
              KLAYTN_WETH.addresses[fromChainId],
              fromAmount
            )
          }
        } else if (toTokenSymbol == SYN.symbol) {
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            SYN.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == FRAX.symbol) {
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            FRAX.addresses[fromChainId],
            fromAmount
          )
        }
      } else if (fromChainId == ChainId.KLAYTN) {
        if (toTokenSymbol == KLAYTN_DAI.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            KLAYTN_DAI.addresses[fromChainId],
            fromAmount,
            { gasPrice: 250000000000 }
          )
        } else if (toTokenSymbol == KLAYTN_USDC.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            KLAYTN_USDC.addresses[fromChainId],
            fromAmount,
            { gasPrice: 250000000000 }
          )
        } else if (toTokenSymbol == KLAYTN_USDT.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            KLAYTN_USDT.addresses[fromChainId],
            fromAmount,
            { gasPrice: 250000000000 }
          )
        } else if (toTokenSymbol == WBTC.symbol) {
          // needs to be merged w/ syn conditional
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            WBTC.addresses[fromChainId],
            fromAmount,
            { gasPrice: 250000000000 }
          )
        } else if (toTokenSymbol == LINK.symbol) {
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            LINK.addresses[fromChainId],
            fromAmount,
            { gasPrice: 250000000000 }
          ) 
        } else if (toTokenSymbol == WETH.symbol && toChainId == ChainId.ETH) {
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            KLAYTN_WETH.addresses[fromChainId],
            fromAmount,
            { gasPrice: 250000000000 }
          )
        } else if (fromTokenSymbol == KLAYTN_WETH.symbol) {
          if (
            [WETHE, ONEETH, FTMETH, METISETH, WETH, CANTOETH]
              .map((i) => i.symbol)
              .includes(toTokenSymbol)
          ) {
            // Destination asset is WETH.e / 1ETH / Fantom WETH / Metis WETH / ETH on Ethereum/L2s
            bridgeZapSwapTransaction = await bridgeZapContract.redeemAndSwap(
              destAddr, // to address
              toChainId, // to chainId
              KLAYTN_WETH.addresses[fromChainId],
              fromAmount,
              0, // tokenIndexFrom, // tokenIndexFrom + 1, // tokenIndexFrom
              tokenIndexTo, // tokenIndexTo + 1, //swapTokenIndex
              minToSwapDest,
              transactionDeadline,
              { gasPrice: 250000000000 }
            )
          } else if (
            [NETH, KLAYTN_WETH].map((i) => i.symbol).includes(toTokenSymbol)
          ) {
            // Destination asset is nETH (most chains) / Dogecoin's WETH
            bridgeZapSwapTransaction = await bridgeZapContract.redeem(
              destAddr, // to address
              toChainId, // to chainId
              KLAYTN_WETH.addresses[fromChainId],
              fromAmount,
              { gasPrice: 250000000000 }
            )
          }
        }
      } else {
        if (toCoin.swapableType == 'JEWEL') {
          // DFK jewel -> avax jewel deposit()
          // DFK jewel -> harmony jewel depositAndSwap()
          // avax jewel -> dfk jewel redeem()
          // Avax jewel -> harmony jewel redeemAndSwap()
          // harmony jewel -> avax jewel swapAndRedeem()
          // harmony jewel -> dfk jewel swapAndRedeem()
          if (fromChainId == ChainId.DFK) {
            if (toChainId == ChainId.HARMONY) {
              bridgeZapSwapTransaction =
                await bridgeZapContract.depositETHAndSwap(
                  destAddr, // to address
                  toChainId, // to chainId
                  fromAmount,
                  1,
                  0,
                  minToSwapDestFromOrigin, // minDy
                  bridgeTransactionDeadline,
                  { value: fromAmount }
                )
            } else {
              bridgeZapSwapTransaction = await bridgeZapContract.depositETH(
                destAddr, // to address
                toChainId, // to chainId
                fromAmount,
                { value: fromAmount }
              )
            }
          } else if (fromChainId == ChainId.HARMONY) {
            // console.log("GOT HERE")
            // bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            //   destAddr, // to address
            //   toChainId, // to chainId
            //   SYNJEWEL.addresses[fromChainId],
            //   fromAmount,
            // )
            bridgeZapSwapTransaction = await bridgeZapContract.swapAndRedeem(
              destAddr, // to address
              toChainId, // to chainId
              SYNJEWEL.addresses[fromChainId],
              0, // tokenIndexFrom
              1, // tokenIndexTo, // token
              fromAmount,
              minToSwapOriginHighSlippage,
              transactionDeadline
            )
          } else if (fromChainId == ChainId.AVALANCHE) {
            // avax jewel -> dfk jewel redeem()
            // Avax jewel -> harmony jewel redeemAndSwap()
            if (toChainId == ChainId.HARMONY) {
              bridgeZapSwapTransaction = await bridgeZapContract.redeemAndSwap(
                destAddr, // to address
                toChainId, // to chainId
                SYNJEWEL.addresses[fromChainId],
                fromAmount,
                1,
                0,
                minToSwapDest,
                transactionDeadline
              )
            } else if (toChainId == ChainId.DFK) {
              bridgeZapSwapTransaction = await bridgeZapContract.redeem(
                destAddr, // to address
                toChainId, // to chainId
                SYNJEWEL.addresses[fromChainId],
                fromAmount
              )
            }
          }
        } else if (toCoin.swapableType == 'XJEWEL') {
          if (toChainId == ChainId.HARMONY) {
            bridgeZapSwapTransaction = await bridgeZapContract.redeem(
              destAddr, // to address
              toChainId, // to chainId
              XJEWEL.addresses[fromChainId],
              fromAmount
            )
          } else {
            bridgeZapSwapTransaction = await bridgeZapContract.deposit(
              destAddr, // to address
              toChainId, // to chainId
              XJEWEL.addresses[fromChainId],
              fromAmount
            )
          }
        } else if (toTokenSymbol == SYN.symbol) {
          // This is the new part added while sleep deprived.
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            SYN.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == USDB.symbol) {
          // This is the new part added while sleep deprived.
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            USDB.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == GOHM.symbol) {
          // This is the new part added while sleep deprived.
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            GOHM.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == SFI.symbol) {
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            SFI.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == HIGHSTREET.symbol) {
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            HIGHSTREET.addresses[fromChainId],
            fromAmount,
          )
        } else if (toTokenSymbol == LINK.symbol) {
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            LINK.addresses[fromChainId],
            fromAmount,
          )
        } else if (toTokenSymbol == DOG.symbol) {
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            DOG.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == NEWO.symbol) {
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            NEWO.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == SDT.symbol) {
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            SDT.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == H2O.symbol) {
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            H2O.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == AGEUR.symbol) {
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            AGEUR.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == L2DAO.symbol) {
          if (fromChainId == ChainId.ARBITRUM) {
            bridgeZapSwapTransaction = await bridgeZapContract.deposit(
              destAddr, // to address
              toChainId, // to chainId
              L2DAO.addresses[fromChainId],
              fromAmount
            )
          } else {
            bridgeZapSwapTransaction = await bridgeZapContract.redeem(
              destAddr, // to address
              toChainId, // to chainId
              L2DAO.addresses[fromChainId],
              fromAmount
            )
          }
        } else if (toTokenSymbol == PLS.symbol) {
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            UNIDX.addresses[fromChainId],
            fromAmount,
          )
        } else if (toTokenSymbol == UNIDX.symbol) {
          if (fromChainId == ChainId.ARBITRUM) {
            bridgeZapSwapTransaction = await bridgeZapContract.deposit(
              destAddr, // to address
              toChainId, // to chainId
              PLS.addresses[fromChainId],
              fromAmount
            )
          } else {
            bridgeZapSwapTransaction = await bridgeZapContract.redeem(
              destAddr, // to address
              toChainId, // to chainId
              PLS.addresses[fromChainId],
              fromAmount
            )
          }
        } else if (toTokenSymbol == FRAX.symbol) {
          bridgeZapSwapTransaction = await bridgeZapContract.redeem(
            destAddr, // to address
            toChainId, // to chainId
            FRAX.addresses[fromChainId],
            fromAmount
          )
        } else if (toTokenSymbol == VSTA.symbol) {
          if (fromChainId == ChainId.ARBITRUM) {
            bridgeZapSwapTransaction = await bridgeZapContract.deposit(
              destAddr, // to address
              toChainId, // to chainId
              VSTA.addresses[fromChainId],
              fromAmount
            )
          } else {
            bridgeZapSwapTransaction = await bridgeZapContract.redeem(
              destAddr, // to address
              toChainId, // to chainId
              VSTA.addresses[fromChainId],
              fromAmount
            )
          }
        } else if (toTokenSymbol == JUMP.symbol) {
          if (fromChainId == ChainId.FANTOM) {
            bridgeZapSwapTransaction = await bridgeZapContract.deposit(
              destAddr, // to address
              toChainId, // to chainId
              JUMP.addresses[fromChainId],
              fromAmount
            )
          } else {
            bridgeZapSwapTransaction = await bridgeZapContract.redeem(
              destAddr, // to address
              toChainId, // to chainId
              JUMP.addresses[fromChainId],
              fromAmount
            )
          }
        } else if (toTokenSymbol == NFD.symbol) {
          if (fromChainId == ChainId.POLYGON) {
            bridgeZapSwapTransaction = await bridgeZapContract.deposit(
              destAddr, // to address
              toChainId, // to chainId
              NFD.addresses[fromChainId],
              fromAmount
            )
          } else {
            bridgeZapSwapTransaction = await bridgeZapContract.redeem(
              destAddr, // to address
              toChainId, // to chainId
              NFD.addresses[fromChainId],
              fromAmount
            )
          }
        } else if (toTokenSymbol == SOLAR.symbol) {
          if (fromChainId == ChainId.MOONRIVER) {
            bridgeZapSwapTransaction = await bridgeZapContract.deposit(
              destAddr, // to address
              toChainId, // to chainId
              SOLAR.addresses[fromChainId],
              fromAmount
            )
          } else {
            bridgeZapSwapTransaction = await bridgeZapContract.redeem(
              destAddr, // to address
              toChainId, // to chainId
              SOLAR.addresses[fromChainId],
              fromAmount
            )
          }
        } else if (toTokenSymbol == GMX.symbol) {
          if (fromChainId == ChainId.ARBITRUM) {
            bridgeZapSwapTransaction = await bridgeZapContract.deposit(
              destAddr, // to address
              toChainId, // to chainId
              GMX.addresses[fromChainId],
              fromAmount
            )
          } else {
            bridgeZapSwapTransaction = await synapseBridgeContract.redeem(
              destAddr, // to address
              toChainId, // to chainId
              GMX.wrapperAddresses[fromChainId],
              fromAmount
            )
          }
        } else if (fromCoin.swapableType == 'AVAX') {
          if (fromChainId == ChainId.AVALANCHE) {
            if (toCoin.symbol == MULTIAVAX.symbol) {
              bridgeZapSwapTransaction =
                await bridgeZapContract.depositETHAndSwap(
                  destAddr, // to address
                  toChainId, // to chainId
                  fromAmount,
                  0, // tokenIndexFrom for nusd
                  tokenIndexTo, // tokenIndexTo + 1,
                  minToSwapDestFromOrigin, // minDy
                  bridgeTransactionDeadline,
                  { value: fromAmount }
                )
            } else {
              bridgeZapSwapTransaction = await bridgeZapContract.depositETH(
                destAddr, // to address
                toChainId, // to chainId
                fromAmount,
                { value: fromAmount }
              )
            }
          } else if ([ChainId.DFK, ChainId.MOONBEAM].includes(fromChainId)) {
            if (toCoin.symbol == MULTIAVAX.symbol) {
              bridgeZapSwapTransaction = await bridgeZapContract.redeemAndSwap(
                destAddr, // to address
                toChainId, // to chainId
                WAVAX.addresses[fromChainId],
                fromAmount,
                0, // tokenIndexFrom, // tokenIndexFrom + 1, // tokenIndexFrom
                tokenIndexTo, // tokenIndexTo + 1, //swapTokenIndex
                minToSwapDest,
                transactionDeadline
              )
            } else {
              bridgeZapSwapTransaction = await bridgeZapContract.redeem(
                destAddr, // to address
                toChainId, // to chainId
                WAVAX.addresses[fromChainId],
                fromAmount
              )
            }
          } else if (fromChainId == ChainId.HARMONY) {
            if (fromCoin.symbol == MULTIAVAX.symbol) {
              bridgeZapSwapTransaction = await bridgeZapContract.swapAndRedeem(
                destAddr, // to address
                toChainId, // to chainId
                SYNAVAX.addresses[fromChainId],
                tokenIndexFrom, // tokenIndexFrom
                0, // tokenIndexTo, // token
                fromAmount,
                minToSwapOriginHighSlippage,
                transactionDeadline
              )
            } else {
              bridgeZapSwapTransaction = await bridgeZapContract.redeem(
                destAddr, // to address
                toChainId, // to chainId
                SYNAVAX.addresses[fromChainId],
                fromAmount
              )
            }
          }
          // @TODO: NEED TO ADD HARMONY AVAX
        } else if (fromCoin.swapableType == 'MOVR') {
          if (fromChainId == ChainId.MOONRIVER) {
            bridgeZapSwapTransaction = await bridgeZapContract.depositETH(
              destAddr, // to address
              toChainId, // to chainId
              fromAmount,
              { value: fromAmount }
            )
          } else if (fromChainId == ChainId.MOONBEAM) {
            bridgeZapSwapTransaction = await bridgeZapContract.redeem(
              destAddr, // to address
              toChainId, // to chainId
              WMOVR.addresses[fromChainId],
              fromAmount
            )
          }
        } else if (
          [NUSD, DFK_USDC].map((i) => i.symbol).includes(toTokenSymbol)
        ) {
          /** basic token on L2 -> NUSD on ETH */
          if ([NUSD, DFK_USDC].map((i) => i.symbol).includes(fromTokenSymbol)) {
            /** NUSD on L2 -> NUSD on L2 */
            bridgeZapSwapTransaction = await bridgeZapContract.redeem(
              destAddr, // to address
              toChainId, // to chainId
              NUSD.addresses[fromChainId],
              fromAmount
            )
          } else {
            bridgeZapSwapTransaction = await bridgeZapContract.swapAndRedeem(
              destAddr, // to address
              toChainId, // to chainId
              NUSD.addresses[fromChainId],
              tokenIndexFrom, // tokenIndexFrom
              0, // tokenIndexTo, // token
              fromAmount,
              minToSwapOriginHighSlippage,
              transactionDeadline
            )
          }
        } else if (
          toTokenSymbol == NETH.symbol ||
          toTokenSymbol == KLAYTN_WETH.symbol
        ) {
          /** basic token on L2 -> NUSD on ETH */
          if (fromTokenSymbol == NETH.symbol) {
            /** NUSD on L2 -> NUSD on L2 */
            bridgeZapSwapTransaction = await bridgeZapContract.redeem(
              destAddr, // to address
              toChainId, // to chainId
              NETH.addresses[fromChainId],
              fromAmount
            )
          } else if (
            [WETHE, ONEETH, FTMETH, METISETH, CANTOETH]
              .map((i) => i.symbol)
              .includes(fromTokenSymbol)
          ) {
            // } else if ((fromTokenSymbol === wethe.symbol) || (fromTokenSymbol === oneeth.symbol) || (fromTokenSymbol === ftmeth.symbol) ) {
            bridgeZapSwapTransaction = await bridgeZapContract.swapAndRedeem(
              destAddr, // to address
              toChainId, // to chainId
              NETH.addresses[fromChainId],
              tokenIndexFrom, // tokenIndexFrom
              0, // tokenIndexTo, // token
              fromAmount,
              minToSwapOriginHighSlippage,
              transactionDeadline
            )
          } else {
            bridgeZapSwapTransaction = await bridgeZapContract.swapETHAndRedeem(
              destAddr, // to address
              toChainId, // to chainId
              NETH.addresses[fromChainId],
              tokenIndexFrom, // tokenIndexFrom
              0, // tokenIndexTo, // token
              fromAmount,
              minToSwapOriginHighSlippage,
              transactionDeadline,
              { value: fromAmount }
            )
          }
        } else if (toChainId == ChainId.ETH) {
          /** basic token on L2 -> ETH */
          if (
            [
              ChainId.BOBA,
              ChainId.ARBITRUM,
              ChainId.OPTIMISM,
              ChainId.AVALANCHE,
              ChainId.HARMONY,
              ChainId.FANTOM,
              ChainId.METIS,
              ChainId.CANTO
            ].includes(fromChainId) &&
            fromCoin.swapableType == 'ETH'
          ) {
            if (fromTokenSymbol == NETH.symbol) {
              bridgeZapSwapTransaction = await bridgeZapContract.redeem(
                destAddr, // to address
                toChainId, // to chainId
                NETH.addresses[fromChainId],
                fromAmount
              )
              // } else if ((fromTokenSymbol === wethe.symbol) || (fromTokenSymbol === oneeth.symbol)) {
            } else if (
              [WETHE, ONEETH, FTMETH, METISETH, CANTOETH]
                .map((i) => i.symbol)
                .includes(fromTokenSymbol)
            ) {
              bridgeZapSwapTransaction = await bridgeZapContract.swapAndRedeem(
                destAddr, // to address
                toChainId, // to chainId
                NETH.addresses[fromChainId],
                tokenIndexFrom, // tokenIndexFrom
                0, // tokenIndexTo, // token
                fromAmount,
                minToSwapOriginHighSlippage, // minToSwapOrigin, // minToSwapOriginHighSlippage,
                transactionDeadline
              )
            } else {
              bridgeZapSwapTransaction =
                await bridgeZapContract.swapETHAndRedeem(
                  destAddr, // to address
                  toChainId, // to chainId
                  NETH.addresses[fromChainId],
                  tokenIndexFrom, // tokenIndexFrom
                  0, // tokenIndexTo, // token
                  fromAmount,
                  minToSwapOriginHighSlippage, // minToSwapOrigin, // minToSwapOriginHighSlippage,
                  transactionDeadline,
                  { value: fromAmount }
                )
            }
          } else {
            if (
              [NUSD, DFK_USDC].map((i) => i.symbol).includes(fromTokenSymbol)
            ) {
              bridgeZapSwapTransaction =
                await bridgeZapContract.redeemAndRemove(
                  destAddr, // to address
                  toChainId, // to chainId
                  NUSD.addresses[fromChainId],
                  fromAmount,
                  tokenIndexTo, // tokenIndexTo + 1, //swapTokenIndex
                  minToSwapDest,
                  transactionDeadline
                )
            } else {
              bridgeZapSwapTransaction =
                await bridgeZapContract.swapAndRedeemAndRemove(
                  destAddr, // to address
                  toChainId, // to chainId
                  NUSD.addresses[fromChainId],
                  tokenIndexFrom,
                  0, // tokenIndexTo
                  fromAmount,
                  minToSwapOriginHighSlippage,
                  transactionDeadline,
                  tokenIndexTo, //swapTokenIndex
                  minToSwapDestFromOriginHighSlippage, // swapMinAmount
                  bridgeTransactionDeadline // toSwapDeadline, // swapDeadline
                )
            }
          }
        } else {
          /** L2 -> L2 */
          if ([NUSD, DFK_USDC].map((i) => i.symbol).includes(fromTokenSymbol)) {
            /** NUSD on L2 -> basic token on L2 */
            bridgeZapSwapTransaction = await bridgeZapContract.redeemAndSwap(
              destAddr, // to address
              toChainId, // to chainId
              NUSD.addresses[fromChainId],
              fromAmount,
              0, // tokenIndexFrom, // tokenIndexFrom + 1, // tokenIndexFrom
              tokenIndexTo, // tokenIndexTo + 1, //swapTokenIndex
              minToSwapDest,
              transactionDeadline
            )
          } else if (fromTokenSymbol === NETH.symbol) {
            /** NUSD on L2 -> basic token on L2 */
            bridgeZapSwapTransaction = await bridgeZapContract.redeemAndSwap(
              destAddr, // to address
              toChainId, // to chainId
              NETH.addresses[fromChainId],
              fromAmount,
              0, // tokenIndexFrom, // tokenIndexFrom + 1, // tokenIndexFrom
              tokenIndexTo, // tokenIndexTo + 1, //swapTokenIndex
              minToSwapDest,
              transactionDeadline
            )
          } else if (fromCoin.swapableType == 'ETH') {
            /** ETH on L2 -> ETH on L2 */
            // if ((fromTokenSymbol === wethe.symbol) || (fromTokenSymbol === oneeth.symbol)) {
            if (
              [WETHE, ONEETH, FTMETH, METISETH, CANTOETH]
                .map((i) => i.symbol)
                .includes(fromTokenSymbol)
            ) {
              bridgeZapSwapTransaction =
                await bridgeZapContract.swapAndRedeemAndSwap(
                  destAddr, // to address
                  toChainId, // to chainId
                  NETH.addresses[fromChainId],
                  tokenIndexFrom, // tokenIndexFrom + 1, // tokenIndexFrom
                  0, // tokenIndexTo, // token
                  fromAmount,
                  minToSwapOriginHighSlippage,
                  transactionDeadline,
                  0, // swapTokenIndexFrom
                  tokenIndexTo, // tokenIndexTo + 1, //swapTokenIndex
                  minToSwapDestFromOriginHighSlippage, // swapMinAmount
                  bridgeTransactionDeadline // toSwapDeadline, // swapDeadline
                )
            } else {
              bridgeZapSwapTransaction =
                await bridgeZapContract.swapETHAndRedeemAndSwap(
                  destAddr, // to address
                  toChainId, // to chainId
                  NETH.addresses[fromChainId],
                  tokenIndexFrom, // tokenIndexFrom + 1, // tokenIndexFrom
                  0, // tokenIndexTo, // token
                  fromAmount,
                  minToSwapOriginHighSlippage,
                  transactionDeadline,
                  0, // swapTokenIndexFrom
                  tokenIndexTo, // tokenIndexTo + 1, //swapTokenIndex
                  minToSwapDestFromOriginHighSlippage, // swapMinAmount
                  bridgeTransactionDeadline, // toSwapDeadline, // swapDeadline
                  { value: fromAmount }
                )
            }
          } else {
            /** stablecoin on L2 -> stablecoin on L2 */
            bridgeZapSwapTransaction =
              await bridgeZapContract.swapAndRedeemAndSwap(
                destAddr, // to address
                toChainId, // to chainId
                NUSD.addresses[fromChainId],
                tokenIndexFrom, // tokenIndexFrom + 1, // tokenIndexFrom
                0, // tokenIndexTo, // token
                fromAmount,
                minToSwapOriginHighSlippage,
                transactionDeadline,
                0, // swapTokenIndexFrom
                tokenIndexTo, // tokenIndexTo + 1, //swapTokenIndex
                minToSwapDestFromOriginHighSlippage, // swapMinAmount
                bridgeTransactionDeadline // toSwapDeadline, // swapDeadline
              )
          }
        }
      }

      console.log(bridgeZapSwapTransaction)

      addTransaction({
        ...bridgeZapSwapTransaction,
        chainId,
      })

      toast(`
          Bridging from ${fromSymbol} on ${fromChainName}
          to ${toSymbol} on ${toChainName}
      `)

      const tx = bridgeZapSwapTransaction.wait
        ? await bridgeZapSwapTransaction.wait()
        : bridgeZapSwapTransaction

      addTransaction({ ...tx, chainId })

      if (destAddr != account) {
        setAddressesForAccount(destAddr)
      }

      if (tx?.status === 1) {
        toast.success(
          <div>
            <div className="w-full">
              Successfully initiated bridge from {fromSymbol} on {fromChainName}{' '}
              to {toSymbol} on {toChainName}
            </div>
            <ExplorerToastLink {...tx} chainId={fromChainId} />
          </div>
        )
      } else if (fromChainId == ChainId.TERRA && tx?.success) {
        toast.success(
          <div>
            <div className="w-full">
              Successfully initiated bridge from {fromSymbol} on {fromChainName}{' '}
              to {toSymbol} on {toChainName}
            </div>
            <ExplorerToastLink
              transactionHash={tx.result?.txhash}
              chainId={fromChainId}
            />
          </div>
        )
      }
      setBlockNumber(tx.blockNumber)
      return tx
    } catch (err) {
      txErrorHandler(err)
    }
  }
}
