import { useState } from 'react'
import { useSettings } from '@hooks/useSettings'
import { SettingsIcon } from '@icons/SettingsIcon'
import { Transition } from '@headlessui/react'
import { BridgeQuote } from '@/utils/types'
import { validateAndParseAddress } from '@utils/validateAndParseAddress'
import { COIN_SLIDE_OVER_PROPS } from '@styles/transitions'
import { Token } from '@/utils/types'
import Grid from '@tw/Grid'
import Card from '@tw/Card'
import Button from '@tw/Button'
import ExchangeRateInfo from '@components/ExchangeRateInfo'
import { TransactionButton } from '@components/buttons/SubmitTxButton'
import { PageHeader } from '@components/PageHeader'
import { CoinSlideOver } from '@components/misc/CoinSlideOver'
import { NetworkSlideOver } from '@components/misc/NetworkSlideOver'
import { BigNumber } from '@ethersproject/bignumber'
import { formatBNToString } from '@bignumber/format'
import SettingsSlideOver from './SettingsSlideOver'
import { DestinationAddressInput } from './DestinationAddressInput'
import BridgeInputContainer from './BridgeInputContainer'
import { useSynapseContext } from '@/utils/SynapseProvider'

import { SECTION_TRANSITION_PROPS } from '@styles/transitions'
// import { useGasDropAmount } from '@hooks/useGasDropAmount'

export default function BridgeCard({
  error,
  address,
  bridgeQuote,
  fromInput,
  fromToken,
  fromTokens,
  fromChainId,
  toToken,
  toChainId,
  toBridgeableChains,
  toBridgeableTokens,
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
  fromTokens: { token: Token; balance: number }[]
  fromChainId: number
  toToken: Token
  toChainId: number
  toBridgeableChains: string[]
  toBridgeableTokens: Token[]
  destinationAddress: string
  handleChainFlip: () => void
  handleTokenChange: (token: Token, type: 'from' | 'to') => void
  onChangeFromAmount: (amount: string) => void
  onSelectFromChain: (chainId: number) => void
  onSelectToChain: (chainId: number) => void
  setDestinationAddress: (address: string) => void
}) {
  const SynapseSDK = useSynapseContext()
  // const [settings, setSettings] = useSettings()
  const [displayType, setDisplayType] = useState('')
  const [deadlineMinutes, setDeadlineMinutes] = useState('')

  const tokenAddr = fromToken.addresses[fromChainId as keyof Token['addresses']]
  let fromTokenBalance = fromTokens.filter(
    (token) => token.token === fromToken
  )[0]?.balance

  const fromArgs = {
    address,
    isOrigin: true,
    chains: toBridgeableChains,
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
    chains: toBridgeableChains,
    tokens: toBridgeableTokens,
    chainId: toChainId,
    inputString: bridgeQuote.outputAmountString,
    selectedToken: toToken,
    connectedChainId: fromChainId,
    setDisplayType,
    handleChainFlip,
    handleTokenChange,
    onChangeChain: onSelectToChain,
  }

  const [settings, setSettings] = useSettings()
  const settingsArgs = {
    settings,
    setSettings,
    setDisplayType,
    setDestinationAddress,
    deadlineMinutes,
    setDeadlineMinutes,
  }
  const deleteme = async () => {
    await console.log('deleteme')
  }
  const approvalBtn = (
    <TransactionButton
      onClick={deleteme}
      label={`Approve ${fromToken.symbol}`}
      pendingLabel={`Approving ${fromToken.symbol}  `}
    />
  )

  const isFromBalanceEnough = fromTokenBalance > Number(fromInput.string)

  let destAddrNotValid
  let btnLabel
  let btnClassName = ''

  if (error) {
    btnLabel = error
  } else if (!isFromBalanceEnough) {
    btnLabel = `Insufficient ${fromToken.symbol} Balance`
  } else if (bridgeQuote.exchangeRate.eq(0) && !fromInput.bigNum.eq(0)) {
    btnLabel = `Amount must be greater than fee`
  } else if (fromChainId == toChainId) {
    btnLabel = 'Why are you bridging to the same network?'
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

    const formattedExchangeRate = formatBNToString(
      bridgeQuote.exchangeRate,
      18,
      4
    )
    const numExchangeRate = Number(formattedExchangeRate)

    if (
      !fromInput.bigNum.eq(0) &&
      (numExchangeRate < 0.95 || numExchangeRate > 1.05)
    ) {
      btnClassName = 'from-[#fe064a] to-[#fe5281]'
      btnLabel = 'Slippage High - Bridge Anyway?'
    }
  }

  const disabled =
    fromChainId == toChainId ||
    bridgeQuote.outputAmount.eq(0) ||
    !isFromBalanceEnough ||
    error != null ||
    destAddrNotValid

  const executeBridge = async () => {
    await SynapseSDK.bridge(
      address, //To Address
      42161,
      43114,
      '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8', // To token Address **
      BigNumber.from('20000000'),
      bridgeQuote.quotes.originQuery,
      bridgeQuote.quotes.destQuery
    )
  }
  const swapBtn = (
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
  )

  const actionBtn = swapBtn
  const bridgeCardMainContent = (
    <>
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
      <div className="px-2 py-2 -mt-2 md:px-0 md:py-4">{actionBtn}</div>
    </>
  )

  const fromCardContent = <CoinSlideOver key="fromBlock" {...fromArgs} />
  const toCardContent = <CoinSlideOver key="toBlock" {...toArgs} />

  const fromChainCardContent = (
    <NetworkSlideOver key="fromChainBlock" {...fromArgs} />
  )
  const toChainCardContent = <NetworkSlideOver key="toChainBlock" {...toArgs} />

  const settingsCardContent = (
    <SettingsSlideOver key="settings" {...settingsArgs} />
  )

  const transitionProps = {
    ...COIN_SLIDE_OVER_PROPS,
    className: `
      origin-bottom absolute
      w-full h-full
      md:w-[95%] md:h-[95%]
      -ml-0 md:-ml-3
      md:mt-3
      bg-bgBase
      z-20 rounded-3xl
    `,
  }

  const settingsTransitionProps = {
    ...COIN_SLIDE_OVER_PROPS,
    className: `
      origin-bottom absolute
      w-full h-full
      md:w-[95%]
      -ml-0 md:-ml-3
      md:-mt-3
      bg-bgBase
      z-20 rounded-3xl
    `,
  }

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
          <Transition show={displayType === 'from'} {...transitionProps}>
            {fromCardContent}
          </Transition>
          <Transition show={displayType === 'to'} {...transitionProps}>
            {toCardContent}
          </Transition>
          <Transition show={displayType === 'fromChain'} {...transitionProps}>
            {fromChainCardContent}
          </Transition>
          <Transition show={displayType === 'toChain'} {...transitionProps}>
            {toChainCardContent}
          </Transition>
          <Transition
            show={displayType === 'settings'}
            {...settingsTransitionProps}
          >
            {settingsCardContent}
          </Transition>
          {bridgeCardMainContent}
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
