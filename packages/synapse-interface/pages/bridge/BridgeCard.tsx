import { useEffect, useState } from 'react'
import { useSettings } from '@hooks/useSettings'
import { SettingsIcon } from '@icons/SettingsIcon'
import { Transition } from '@headlessui/react'
import { validateAndParseAddress } from '@utils/validateAndParseAddress'
import { TRANSITIONS_PROPS } from '@constants/bridge'
import { COIN_SLIDE_OVER_PROPS } from '@styles/transitions'
import { ORDERED_CHAINS_BY_ID } from '@constants/chains'
import { erc20ABI } from 'wagmi'
import Grid from '@tw/Grid'
import Card from '@tw/Card'
import Button from '@tw/Button'
import ExchangeRateInfo from '@components/ExchangeRateInfo'
import { TransactionButton } from '@components/buttons/SubmitTxButton'
import { PageHeader } from '@components/PageHeader'
import { CoinSlideOver } from '@components/misc/CoinSlideOver'
import { NetworkSlideOver } from '@components/misc/NetworkSlideOver'
import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'
import { formatBNToString } from '@bignumber/format'
import { fetchSigner } from '@wagmi/core'
import { SECTION_TRANSITION_PROPS } from '@styles/transitions'
import { Contract } from 'ethers'

import SettingsSlideOver from './SettingsSlideOver'
import { DestinationAddressInput } from './DestinationAddressInput'
import BridgeInputContainer from './BridgeInputContainer'
import { useSynapseContext } from '@/utils/SynapseProvider'
// import { writeContract, prepareSendTransaction } from '@wagmi/core'
// import { optimism } from '@wagmi/core/chains'

import { Token } from '@/utils/types'
import { BridgeQuote } from '@/utils/types'
// import { useGasDropAmount } from '@hooks/useGasDropAmount'
const BridgeCard = ({
  error,
  address,
  bridgeQuote,
  fromInput,
  fromToken,
  fromTokens,
  fromChainId,
  toToken,
  toChainId,
  toOptions,
  destinationAddress,
  handleChainFlip,
  handleTokenChange,
  onChangeFromAmount,
  onSelectFromChain,
  onSelectToChain,
  setDestinationAddress,
}: {
  error
  address: `0x${string}` | undefined
  bridgeQuote: BridgeQuote
  fromInput: { string: string; bigNum: BigNumber }
  fromToken: Token
  fromTokens: { token: Token; balance: BigNumber }[]
  fromChainId: number
  toToken: Token
  toChainId: number
  toOptions: { tokens: Token[]; chains: string[] }
  destinationAddress: string
  handleChainFlip: () => void
  handleTokenChange: (token: Token, type: 'from' | 'to') => void
  onChangeFromAmount: (amount: string) => void
  onSelectFromChain: (chainId: number) => void
  onSelectToChain: (chainId: number) => void
  setDestinationAddress: (address: string) => void
}) => {
  const SynapseSDK = useSynapseContext()
  const [settings, setSettings] = useSettings()
  const [displayType, setDisplayType] = useState('')
  const [deadlineMinutes, setDeadlineMinutes] = useState('')
  const [fromTokenBalance, setFromTokenBalance] = useState<BigNumber>(Zero)

  useEffect(() => {
    if (fromTokens && fromToken) {
      setFromTokenBalance(
        fromTokens.filter((token) => token.token === fromToken)[0]?.balance
          ? fromTokens.filter((token) => token.token === fromToken)[0]?.balance
          : Zero
      )
    }
  }, [fromToken, fromTokens])

  const fromArgs = {
    address,
    fromTokenBalance,
    isOrigin: true,
    chains: ORDERED_CHAINS_BY_ID.filter((id) => id !== String(fromChainId)),
    tokens: fromTokens,
    chainId: fromChainId,
    inputString: fromInput.string,
    selectedToken: fromToken,
    connectedChainId: fromChainId,
    setDisplayType,
    handleTokenChange,
    onChangeChain: onSelectFromChain,
    onChangeAmount: onChangeFromAmount,
  }

  const toArgs = {
    address,
    isOrigin: false,
    chains: toOptions.chains,
    tokens: toOptions.tokens,
    chainId: toChainId,
    inputString: bridgeQuote.outputAmountString,
    selectedToken: toToken,
    connectedChainId: fromChainId,
    setDisplayType,
    handleChainFlip,
    handleTokenChange,
    onChangeChain: onSelectToChain,
  }

  const settingsArgs = {
    settings,
    setSettings,
    setDisplayType,
    setDestinationAddress,
    deadlineMinutes,
    setDeadlineMinutes,
  }

  const approveToken = async () => {
    // TODO store this erc20 and signer retrieval in a state in a parent component
    const wallet = await fetchSigner({
      chainId: fromChainId,
    })

    const erc20 = new Contract(
      fromToken.addresses[fromChainId],
      erc20ABI,
      wallet
    )
    const approveTx = await erc20.approve(
      bridgeQuote.routerAddress,
      fromInput.bigNum,
      {
        gasPrice: await wallet.provider.getGasPrice(),
      }
    )

    try {
      await approveTx.wait()
      console.log(`Transaction mined succesfully: ${approveTx.hash}`)
    } catch (error) {
      console.log(`Transaction failed with error: ${error}`)
    }
  }
  const executeBridge = async () => {
    const wallet = await fetchSigner({
      chainId: fromChainId,
    })

    // if ()
    const data = await SynapseSDK.bridge(
      address, //To Address
      fromChainId,
      toChainId,
      fromToken.addresses[fromChainId as keyof Token['addresses']], // To token Address **
      fromInput.bigNum,
      bridgeQuote.quotes.originQuery,
      bridgeQuote.quotes.destQuery
    )
      .then((res) => {
        const tx = res
        wallet
          .sendTransaction(tx)
          .then((res) => {
            console.log('sendTransaction', res)
          })
          .catch((err) => console.log('sendTransaction', err))
      })
      .catch((err) => {
        console.log('bridge', err)
      })

    console.log('data', data)
  }
  const isFromBalanceEnough = fromTokenBalance?.gt(fromInput.bigNum)

  let destAddrNotValid
  let btnLabel
  let btnClassName = ''
  let pendingLabel = 'Bridging funds...'
  let buttonAction = executeBridge
  if (error) {
    btnLabel = error
  } else if (!isFromBalanceEnough) {
    btnLabel = `Insufficient ${fromToken.symbol} Balance`
  } else if (bridgeQuote.feeAmount.eq(0) && !fromInput.bigNum.eq(0)) {
    btnLabel = `Amount must be greater than fee`
  } else if (
    bridgeQuote?.allowance &&
    bridgeQuote?.allowance?.gt(Zero) &&
    bridgeQuote?.allowance?.lt(fromInput.bigNum)
  ) {
    buttonAction = approveToken
    btnLabel = `Approve ${fromToken.symbol}`
    pendingLabel = `Approving ${fromToken.symbol}`
    btnClassName = 'from-[#feba06] to-[#FEC737]'
  } else if (
    destinationAddress &&
    !validateAndParseAddress(destinationAddress)
  ) {
    destAddrNotValid = true
    btnLabel = 'Invalid Destination Address'
  } else {
    btnLabel = bridgeQuote.outputAmount.eq(0)
      ? 'Enter amount to bridge'
      : 'Bridge your funds'

    const numExchangeRate = Number(
      formatBNToString(bridgeQuote.exchangeRate, 18, 4)
    )

    if (
      !fromInput.bigNum.eq(0) &&
      (numExchangeRate < 0.95 || numExchangeRate > 1.05)
    ) {
      btnClassName = 'from-[#fe064a] to-[#fe5281]'
      btnLabel = 'Slippage High - Bridge Anyway?'
    }
  }
  const disabled =
    fromChainId === toChainId ||
    bridgeQuote.outputAmount.eq(0) ||
    !isFromBalanceEnough ||
    error != null ||
    destAddrNotValid
  const actionButton = (
    <TransactionButton
      className={btnClassName}
      disabled={disabled}
      onClick={() => buttonAction()}
      onSuccess={() => {
        onChangeFromAmount('')
      }}
      label={btnLabel}
      pendingLabel={pendingLabel}
    />
  )
  return (
    <>
      <div className="flex items-center justify-between mb-5 ml-5 mr-5 space-x-2">
        {displayType !== 'settings' ? (
          <PageHeader
            title="Bridge"
            subtitle="Send your assets across chains."
          />
        ) : (
          <PageHeader title="Settings" subtitle="Customize your experience." />
        )}
        <div>
          <Button
            className="flex items-center p-3 text-opacity-75 bg-bgLight hover:bg-bgLighter text-secondaryTextColor hover:text-white"
            onClick={() => {
              if (displayType !== 'settings') {
                setDisplayType('settings')
              } else {
                setDisplayType('')
              }
            }}
          >
            {displayType !== 'settings' ? (
              <>
                <SettingsIcon className="w-5 h-5 mr-2" />
                <span>Settings</span>
              </>
            ) : (
              <span>Close</span>
            )}
          </Button>
        </div>
      </div>
      <Card
        divider={false}
        className="max-w-lg px-1 pb-0 mb-3 transition-all duration-100 transform rounded-xl bg-bgBase md:px-6 lg:px-6"
      >
        <div>
          <Transition show={displayType === 'from'} {...TRANSITIONS_PROPS}>
            <CoinSlideOver key="fromBlock" {...fromArgs} />
          </Transition>
          <Transition show={displayType === 'to'} {...TRANSITIONS_PROPS}>
            <CoinSlideOver key="toBlock" {...toArgs} />
          </Transition>
          <Transition show={displayType === 'fromChain'} {...TRANSITIONS_PROPS}>
            <NetworkSlideOver key="fromChainBlock" {...fromArgs} />
          </Transition>
          <Transition show={displayType === 'toChain'} {...TRANSITIONS_PROPS}>
            <NetworkSlideOver key="toChainBlock" {...toArgs} />
          </Transition>
          <Transition show={displayType === 'settings'} {...TRANSITIONS_PROPS}>
            <SettingsSlideOver key="settings" {...settingsArgs} />
          </Transition>
          <Grid cols={{ xs: 1 }} gap={10} className="py-1 place-content-center">
            <div className="mt-2">
              <BridgeInputContainer {...fromArgs} />
            </div>
            <BridgeInputContainer {...toArgs} />
          </Grid>
          <Transition
            appear={true}
            unmount={false}
            show={!fromInput.bigNum.eq(0)}
            {...SECTION_TRANSITION_PROPS}
          >
            <ExchangeRateInfo
              fromAmount={fromInput.bigNum}
              toToken={toToken}
              exchangeRate={bridgeQuote.exchangeRate}
              toChainId={toChainId}
            />
          </Transition>
          <Transition
            appear={false}
            unmount={false}
            show={settings.expertMode}
            {...SECTION_TRANSITION_PROPS}
          >
            <DestinationAddressInput
              toChainId={toChainId}
              destinationAddress={destinationAddress}
              setDestinationAddress={setDestinationAddress}
            />
          </Transition>
          <div className="px-2 py-2 -mt-2 md:px-0 md:py-4">
            {actionButton}
            {/* {generateActionButton()} */}
            {/* {bridgeQuote?.allowance &&
            bridgeQuote?.allowance?.lt(fromInput.bigNum) ? (
              <TransactionButton
                onClick={() => approveToken()}
                label={`Approve ${fromToken.symbol}`}
                onSuccess={() => {
                  console.log('YIUGJHGJHGJHGJHGHJ')
                }}
                pendingLabel={`Approving ${fromToken.symbol}  `}
              />
            ) : (
              <TransactionButton
                className={btnClassName}
                disabled={disabled}
                onClick={() => executeBridge()}
                onSuccess={() => {
                  onChangeFromAmount('')
                }}
                label={btnLabel}
                pendingLabel={`Bridging funds...`}
              />
            )} */}
          </div>
          <Transition
            show={
              ['fromChain', 'toChain'].includes(displayType)
              // && feeConfig. .eq(Zero)
            }
            {...COIN_SLIDE_OVER_PROPS}
          ></Transition>
        </div>
      </Card>
    </>
  )
}

// // TODO: Fix the transition post ftm addition
// // TODO: Need to coordnate transition from approval => other action

// function NetworkPausedButton({ networkName }) {
//   return (
//     <Button disabled={true} type="button" className={ACTION_BTN_CLASSNAME}>
//       {networkName} Undergoing Chain Downtime
//     </Button>
//   )
// }
// // Undergoing Network Upgrades

// const PAUSED_BASE_PROPERTIES = `
//     w-full rounded-lg my-2 px-4 py-3
//     text-white text-opacity-100 transition-all
//     hover:opacity-80 disabled:opacity-100 disabled:text-[#88818C]
//     disabled:from-bgLight disabled:to-bgLight
//     bg-gradient-to-r from-[#CF52FE] to-[#AC8FFF]
//   `

// function PausedButton({ networkName }) {
//   return (
//     <Button
//       disabled={true}
//       fancy={true}
//       type="button"
//       className={`${PAUSED_BASE_PROPERTIES}`}
//     >
//      Temporarily paused due to chain connectivity issues
//     </Button>
//   )
// }

// function HeavyLoadButton() {
//   return (
//     <Button
//       disabled={true}
//       fancy={true}
//       type="button"
//       className={ACTION_BTN_CLASSNAME}
//     >
//       Synapse is experiencing heavy load
//     </Button>
//   )
// }

// function AdvancedOptionsButton({ className, onClick }) {
//   return (
//     <div
//       className={`
//         group rounded-lg hover:bg-gray-900 ${className} p-1`}
//       onClick={onClick}
//     >
//       <CogIcon className="w-6 h-6 text-gray-500 group-hover:text-gray-300" />
//     </div>
//   )
// }

export default BridgeCard
