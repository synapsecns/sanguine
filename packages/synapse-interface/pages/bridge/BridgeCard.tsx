import { useEffect, useContext, useState } from 'react'

import { Zero, One } from '@ethersproject/constants'
import { parseUnits } from '@ethersproject/units'
import { useBalance } from 'wagmi'
import { useSettings } from '@hooks/settings/useSettings'

import { SettingsIcon } from '@icons/SettingsIcon'
import { Transition } from '@headlessui/react'
// import { useSettings } from '@hooks/settings/useSettings'
// import { useGasDropAmount } from '@hooks/useGasDropAmount'
// import { useBridgeSwap } from '@hooks/actions/useBridgeSwap'
// import { useSynapseContract } fromnpm i '@hooks/contracts/useSynapseContract'

// import { APPROVAL_STATE, useApproveToken } from '@hooks/actions/useApproveToken'
import { useSynapseContext } from '@/utils/SynapseProvider'

import { sanitizeValue } from '@utils/sanitizeValue'
import { validateAndParseAddress } from '@utils/validateAndParseAddress'

import { BRIDGABLE_TOKENS } from '@constants/tokens'

import { COIN_SLIDE_OVER_PROPS } from '@styles/transitions'
import { Token } from '@utils/classes/Token'

import Grid from '@tw/Grid'
import Card from '@tw/Card'
import Button from '@tw/Button'

import ExchangeRateInfo from '@components/ExchangeRateInfo'
import { TransactionButton } from '@components/buttons/SubmitTxButton'
import { PageHeader } from '@components/PageHeader'

import BridgeInputContainer from './BridgeInputContainer'
import { CoinSlideOver } from '@components/misc/CoinSlideOver'
import { NetworkSlideOver } from '@components/misc/NetworkSlideOver'
import SettingsSlideOver from './SettingsSlideOver'
import { DestinationAddressInput } from './DestinationAddressInput'

import { BigNumber } from '@ethersproject/bignumber'
import { formatBNToString } from '@bignumber/format'

// import { getBridgeQuote } from '@hooks/synapse'

//  console.log(getBridgeQuote())

// const ACTION_BTN_CLASSNAME = `
//   w-full rounded-lg my-2 px-4 py-3 tracking-wide
//   text-white disabled:bg-gray-300 transition-all
//   `

const SECTION_TRANSITION_PROPS = {
  enter: 'transition duration-100 ease-out',
  enterFrom: 'transform-gpu scale-y-0 ',
  enterTo: 'transform-gpu scale-y-100 opacity-100',
  leave: 'transition duration-75 ease-out ',
  leaveFrom: 'transform-gpu scale-y-100 opacity-100',
  leaveTo: 'transform-gpu scale-y-0 ',
  className: 'origin-top -mx-0 md:-mx-6',
}

export default function BridgeCard({
  address,
  fromChainId,
  toChainId,
  fromCoin,
  fromValue,
  toCoin,
  toValue,
  possibleChains,
  onSelectFromChain,
  onSelectToChain,
  swapFromToChains,
  // onSelectFromCoin,
  // onSelectToCoin,
  onChangeFromAmount,
  onChangeToAmount,
  error,
  priceImpact,
  exchangeRate,
  feeAmount,
  destinationAddress,
  setDestinationAddress,
  handleTokenChange,
  toBridgeableTokens,
  quotes,
}: {
  address: `0x${string}` | undefined
  fromChainId: number
  toChainId: number
  fromCoin: Token
  fromValue: string
  toCoin: Token
  toValue: string
  possibleChains: string[]
  onSelectFromChain: (v: number) => void
  onSelectToChain: (v: number) => void
  swapFromToChains: () => void
  // onSelectFromCoin: (v: Token) => void
  // onSelectToCoin: (v: Token) => void

  onChangeFromAmount: (v: string) => void
  onChangeToAmount: (v: string) => void

  error?: string
  priceImpact: BigNumber
  exchangeRate: BigNumber
  feeAmount: BigNumber
  destinationAddress: string
  setDestinationAddress: (v: string) => void
  handleTokenChange: (token: Token, type: 'from' | 'to') => void
  toBridgeableTokens: Token[]
  quotes: any
}) {
  const SynapseSDK = useSynapseContext()
  // populates the selectable tokens using the from and to chain ids
  const fromChainTokens = BRIDGABLE_TOKENS[Number(fromChainId)]

  // can be replaced by get bridge quote
  // const gasDropAmount = useGasDropAmount(toChainId)

  // augment settings
  const [displayType, setDisplayType] = useState('')

  // settings
  // const [settings, setSettings] = useSettings()

  // deadline set in settings
  const [deadlineMinutes, setDeadlineMinutes] = useState('')

  // gets the from amount from the props
  let fromAmount: BigNumber
  try {
    fromAmount = parseUnits(
      sanitizeValue(fromValue),
      fromCoin.decimals?.[fromChainId as keyof Token['decimals']]
    )
  } catch (e) {
    fromAmount = Zero
  }

  // gets the to amount from the props
  let toAmount: BigNumber
  try {
    toAmount = parseUnits(
      sanitizeValue(toValue),
      toCoin.decimals?.[toChainId as keyof Token['decimals']]
    )
  } catch (e) {
    toAmount = Zero
  }

  // SDK
  // const bridgeSwap = useBridgeSwap({ amount: fromAmount, token: fromCoin })
  const bridgeSwap = null

  // let targetApprovalContract = useSynapseContract()

  // const { approvalState, approveToken } = useApproveToken(
  //   fromCoin,
  //   String(targetApprovalContract?.address),
  //   fromAmount
  // )
  const tokenAddr = fromCoin.addresses[fromChainId as keyof Token['addresses']]
  let fromTokenBalance: BigNumber
  if (!tokenAddr) {
    const { data: rawTokenBalance } = useBalance({
      chainId: fromChainId,
      address: address,
    })
    fromTokenBalance = rawTokenBalance?.value ?? Zero
  } else {
    const { data: rawTokenBalance } = useBalance({
      chainId: fromChainId,
      address: address,
      token: `0x${tokenAddr.slice(2)}`,
    })
    fromTokenBalance = rawTokenBalance?.value ?? Zero
  }

  // end nonevm dest
  const fromArgs = {
    isSwapFrom: true,
    selected: fromCoin,
    address: address,
    connectedChainId: fromChainId,
    // onChangeSelected: onSelectFromCoin,
    handleTokenChange: handleTokenChange,
    onChangeAmount: onChangeFromAmount,
    inputValue: fromValue,
    tokens: fromChainTokens,
    chainId: fromChainId,
    setDisplayType,
    onChangeChain: onSelectFromChain,
  }

  const toArgs = {
    isSwapFrom: false,
    selected: toCoin,
    address: address,
    connectedChainId: fromChainId,
    // onChangeSelected: onSelectToCoin,
    handleTokenChange: handleTokenChange,
    onChangeAmount: onChangeToAmount,
    inputValue: toValue,
    tokens: toBridgeableTokens,
    chainId: toChainId,
    swapFromToChains,
    setDisplayType,
    onChangeChain: onSelectToChain,
  }

  const fromChainArgs = {
    isSwapFrom: true,
    selectedChainId: fromChainId,
    onChangeChain: onSelectFromChain,
    setDisplayType,
  }

  const toChainArgs = {
    isSwapFrom: false,
    selectedChainId: toChainId,
    onChangeChain: onSelectToChain,
    setDisplayType,
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
      label={`Approve ${fromCoin.symbol}`}
      pendingLabel={`Approving ${fromCoin.symbol}  `}
    />
  )

  const isFromBalanceEnough = fromTokenBalance.gte(fromAmount) // && !fromAmount.eq(0)

  let destAddrNotValid
  let btnLabel
  let btnClassName = ''

  if (error) {
    btnLabel = error
  } else if (!isFromBalanceEnough) {
    btnLabel = `Insufficient ${fromCoin.symbol} Balance`
  } else if (exchangeRate.eq(0) && !fromAmount.eq(0)) {
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
    btnLabel = toAmount.eq(0) ? 'Enter amount to bridge' : 'Bridge your funds'

    const formattedExchangeRate = formatBNToString(exchangeRate, 18, 4)
    const numExchangeRate = Number(formattedExchangeRate)

    if (
      !fromAmount.eq(0) &&
      (numExchangeRate < 0.95 || numExchangeRate > 1.05)
    ) {
      btnClassName = 'from-[#fe064a] to-[#fe5281]'
      btnLabel = 'Slippage High - Bridge Anyway?'
    }
  }

  const disabled =
    fromChainId == toChainId ||
    toAmount.eq(0) ||
    !isFromBalanceEnough ||
    error != null ||
    destAddrNotValid
  const sss = async (): Promise<any> => {
    await console.log('s')
  }
  const executeBridge = async () => {
    await SynapseSDK.bridge(
      address, //To Address
      42161,
      43114,
      '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8', // To token Address **
      BigNumber.from('20000000'),
      quotes.originQuery,
      quotes.destQuery
    )
  }
  const swapBtn = (
    <TransactionButton
      className={btnClassName}
      disabled={disabled}
      onClick={sss}
      // onClick={() => {
      //   return bridgeSwap({
      //     destinationAddress,
      //     fromChainId,
      //     toChainId,
      //     fromAmount,
      //     fromCoin,
      //     toAmount,
      //     toCoin,
      //     deadlineMinutes,
      //   })
      // }}
      onSuccess={() => {
        onChangeFromAmount('')
      }}
      label={btnLabel}
      pendingLabel={`Bridging funds...`}
    />
  )

  let approvalRequired = true
  // if (
  //   fromChainId === ChainId.AVALANCHE &&
  //   GMX.addresses[ChainId.AVALANCHE] === fromCoin.addresses[ChainId.AVALANCHE]
  // ) {
  //   approvalRequired = false
  // } else {
  //   approvalRequired = true
  // }

  let actionBtn = swapBtn
  // if (approvalState === APPROVAL_STATE.NOT_APPROVED && approvalRequired) {
  //   actionBtn = approvalBtn
  //   //    } else if ([fromChainId, toChainId].includes(ChainId.POLYGON)) {
  //   //} else if ([toChainId, fromChainId].includes(ChainId.CANTO)) {
  //   //     actionBtn = <NetworkPausedButton networkName="Polygon" />
  // } else {
  //   //   actionBtn = <PausedButton/> // PAUSE OVERRIDE
  //   actionBtn = swapBtn
  // }
  // //  }
  // // let actionBtn = <PausedButton/> // PAUSE OVERRIDE

  const bridgeCardMainContent = (
    <>
      <Grid cols={{ xs: 1 }} gap={10} className="py-1 place-content-center">
        <div className="mt-2">
          <BridgeInputContainer possibleChains={possibleChains} {...fromArgs} />
        </div>
        <BridgeInputContainer possibleChains={possibleChains} {...toArgs} />
      </Grid>
      {/* <Transition
        appear={true}
        unmount={false}
        show={!fromAmount.eq(0)}
        {...SECTION_TRANSITION_PROPS}
      > */}
      <ExchangeRateInfo
        fromAmount={fromAmount}
        fromCoin={fromCoin}
        toCoin={toCoin}
        exchangeRate={exchangeRate}
        gasDropAmount={One}
        fromChainId={fromChainId}
        toChainId={toChainId}
      />
      {/* </Transition> */}
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
    <NetworkSlideOver
      key="fromChainBlock"
      possibleChains={possibleChains}
      {...fromChainArgs}
    />
  )
  const toChainCardContent = (
    <NetworkSlideOver
      key="toChainBlock"
      possibleChains={possibleChains}
      {...toChainArgs}
    />
  )

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
              ['fromChain', 'toChain'].includes(displayType) &&
              feeAmount.eq(Zero)
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
