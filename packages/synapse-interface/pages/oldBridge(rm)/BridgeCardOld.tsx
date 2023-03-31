import { useEffect, useState } from 'react'
import { Zero } from '@ethersproject/constants'
import { parseUnits } from '@ethersproject/units'
import { SettingsIcon } from '@icons/SettingsIcon'
import { Transition } from '@headlessui/react'
import { useSettings } from '@hooks/settings/useSettings'
import { useGasDropAmount } from '@hooks/useGasDropAmount'
import { useBridgeSwap } from '@hooks/actions/useBridgeSwap'
import { useTerraUstBalance } from '@hooks/terra/useTerraUstBalance'
import { useSynapseContract } from '@hooks/contracts/useSynapseContract'
import { useBridgeZapContract } from '@hooks/contracts/useBridgeZapContract'
import { APPROVAL_STATE, useApproveToken } from '@hooks/actions/useApproveToken'
import { useTokenBalance } from '@hooks/tokens/useTokenBalances'
import { ChainId } from '@constants/networks'
import { validateAndParseAddress } from '@utils/validateAndParseAddress'
import { BRIDGABLE_TOKENS } from '@constants/bridge'
import { GMX } from '@constants/tokens/mintable'
import { COIN_SLIDE_OVER_PROPS } from '@styles/transitions'
import Grid from '@tw/Grid'
import Card from '@tw/Card'
import Button from '@tw/Button'
import ExchangeRateInfo from '@components/ExchangeRateInfo'
import { TransactionButton } from '@components/buttons/SubmitTxButton'
import { PageHeader } from '@components/PageHeader'
import { CoinSlideOver } from '@components/misc/CoinSlideOver'
import { NetworkSlideOver } from '@components/misc/NetworkSlideOver'
import { validateTerraAddress } from '@utils/validateTerraAddress'
import { BigNumber } from '@ethersproject/bignumber'
import { formatBNToPercentString, formatBNToString } from '@bignumber/format'

import { DestinationAddressInput } from './DestinationAddressInput'
import SettingsSlideOver from './SettingsSlideOver'
import BridgeInputContainer from './BridgeInputContainer'
import { sanitizeValue } from '@/utils/stringToBigNum'

const ACTION_BTN_CLASSNAME = `
  w-full rounded-lg my-2 px-4 py-3 tracking-wide
  text-white disabled:bg-gray-300 transition-all
  `

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
  fromChainId,
  toChainId,
  fromCoin,
  fromValue,
  toCoin,
  toValue,
  onSelectFromChain,
  onSelectToChain,
  swapFromToChains,
  onSelectFromCoin,
  onSelectToCoin,
  onChangeFromAmount,
  onChangeToAmount,
  error,
  priceImpact,
  exchangeRate,
  feeAmount,
  fromRef,
  toRef,
  destinationAddress,
  setDestinationAddress,
}) {
  // populates the selectable tokens using the from and to chain ids
  const fromChainTokens = BRIDGABLE_TOKENS[fromChainId]
  const toChainTokens = BRIDGABLE_TOKENS[toChainId]

  // can be replaced by get bridge quote
  const gasDropAmount = useGasDropAmount(toChainId)

  // augment settings
  const [displayType, setDisplayType] = useState(undefined)

  // settings
  const [settings, setSettings] = useSettings()

  // deadline set in settings
  const [deadlineMinutes, setDeadlineMinutes] = useState()

  // gets the from amount from the props
  let fromAmount
  try {
    fromAmount = parseUnits(
      sanitizeValue(fromValue),
      fromCoin.decimals[fromChainId]
    )
  } catch (e) {
    fromAmount = Zero
  }

  // gets the to amount from the props
  let toAmount
  try {
    toAmount = parseUnits(sanitizeValue(toValue), toCoin.decimals[toChainId])
  } catch (e) {
    toAmount = Zero
  }

  // nuke this bc no more terra
  const bridgeZapContract = useBridgeZapContract()

  // gets the bridge contract (may no need bc of sdk)
  const synapseBridgeContract = useSynapseContract()

  const bridgeSwap = useBridgeSwap({ amount: fromAmount, token: fromCoin })

  // targetApprovalContract = synapseBridgeContract
  let targetApprovalContract
  // nuke this
  if (fromCoin.swapableType == 'UST' && fromChainId != ChainId.TERRA) {
    targetApprovalContract = synapseBridgeContract
  } else {
    targetApprovalContract = bridgeZapContract
  }

  const [approvalState, approveToken] = useApproveToken(
    fromCoin,
    targetApprovalContract.address,
    fromAmount
  )

  const evmFromTokenBalance = useTokenBalance(fromCoin)
  const terraUstBalance = useTerraUstBalance()

  let fromTokenBalance
  if (fromCoin.symbol == 'UST' && fromChainId == ChainId.TERRA) {
    fromTokenBalance = terraUstBalance
  } else {
    fromTokenBalance = evmFromTokenBalance
  }

  // start nonevm dest NUKE EVERYTHING RELATED TO THIS
  let nonEvmBridge
  if ([fromChainId, toChainId].includes(ChainId.TERRA)) {
    nonEvmBridge = true
  } else {
    nonEvmBridge = false
  }

  // const expertMode = useMemo(() => settings.expertMode, [settings.expertMode])

  // nuke
  useEffect(() => {
    if (!nonEvmBridge) {
      setDestinationAddress('')
    }
  }, [nonEvmBridge])

  // nuke
  useEffect(() => {
    if (!settings.expertMode) {
      setDeadlineMinutes(undefined)
    }
  }, [settings.expertMode])

  // end nonevm dest THE FUCK DOE Sjhfaskjh

  const fromArgs = {
    isSwapFrom: true,
    selected: fromCoin,
    onChangeSelected: onSelectFromCoin,
    onChangeAmount: onChangeFromAmount,
    inputValue: fromValue,
    inputRef: fromRef,
    tokens: fromChainTokens,
    chainId: fromChainId,
    setDisplayType,
    onChangeChain: onSelectFromChain,
  }

  const toArgs = {
    isSwapFrom: false,
    selected: toCoin,
    onChangeSelected: onSelectToCoin,
    onChangeAmount: onChangeToAmount,
    inputValue: toValue,
    inputRef: toRef,
    tokens: toChainTokens,
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

  const settingsArgs = {
    settings,
    setSettings,
    setDisplayType,
    setDestinationAddress,
    deadlineMinutes,
    setDeadlineMinutes,
  }

  const approvalBtn = (
    <TransactionButton
      onClick={approveToken}
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
    !(
      validateAndParseAddress(destinationAddress) ||
      validateTerraAddress(destinationAddress)
    )
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
    error ||
    destAddrNotValid

  const swapBtn = (
    <TransactionButton
      className={btnClassName}
      disabled={disabled}
      onClick={() => {
        return bridgeSwap({
          destinationAddress,
          fromChainId,
          toChainId,
          fromAmount,
          fromCoin,
          toAmount,
          toCoin,
          deadlineMinutes,
        })
      }}
      onSuccess={() => {
        onChangeFromAmount('')
      }}
      label={btnLabel}
      pendingLabel={`Bridging funds...`}
    />
  )

  let approvalRequired
  if (fromChainId == ChainId.TERRA) {
    approvalRequired = false
  } else if (
    fromChainId === ChainId.AVALANCHE &&
    GMX.addresses[ChainId.AVALANCHE] === fromCoin.addresses[ChainId.AVALANCHE]
  ) {
    approvalRequired = false
  } else {
    approvalRequired = true
  }

  let actionBtn
  if (approvalState === APPROVAL_STATE.NOT_APPROVED && approvalRequired) {
    actionBtn = approvalBtn
    //    } else if ([fromChainId, toChainId].includes(ChainId.POLYGON)) {
    //} else if ([toChainId, fromChainId].includes(ChainId.CANTO)) {
    //     actionBtn = <NetworkPausedButton networkName="Polygon" />
  } else {
    //   actionBtn = <PausedButton/> // PAUSE OVERRIDE
    actionBtn = swapBtn
  }
  //  }
  // let actionBtn = <PausedButton/> // PAUSE OVERRIDE

  const bridgeCardMainContent = (
    <>
      <Grid cols={{ xs: 1 }} gap={10} className="py-1 place-content-center">
        <div className="mt-2">
          <BridgeInputContainer {...fromArgs} />
        </div>
        <BridgeInputContainer {...toArgs} />
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
        priceImpact={priceImpact}
        feeAmount={feeAmount}
        gasDropAmount={gasDropAmount}
        fromChainId={fromChainId}
        toChainId={toChainId}
      />
      {/* </Transition> */}
      <Transition
        appear={false}
        unmount={false}
        show={settings.expertMode || nonEvmBridge}
        {...SECTION_TRANSITION_PROPS}
      >
        <DestinationAddressInput
          fromChainId={fromChainId}
          toChainId={toChainId}
          nonEvmBridge={nonEvmBridge}
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
    <NetworkSlideOver key="fromChainBlock" {...fromChainArgs} />
  )
  const toChainCardContent = (
    <NetworkSlideOver key="toChainBlock" {...toChainArgs} />
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
                setDisplayType(undefined)
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

// TODO: Fix the transition post ftm addition
// TODO: Need to coordnate transition from approval => other action

function NetworkPausedButton({ networkName }) {
  return (
    <Button disabled={true} type="button" className={ACTION_BTN_CLASSNAME}>
      {networkName} Undergoing Chain Downtime
    </Button>
  )
}
// Undergoing Network Upgrades

const PAUSED_BASE_PROPERTIES = `
    w-full rounded-lg my-2 px-4 py-3
    text-white text-opacity-100 transition-all
    hover:opacity-80 disabled:opacity-100 disabled:text-[#88818C]
    disabled:from-bgLight disabled:to-bgLight
    bg-gradient-to-r from-[#CF52FE] to-[#AC8FFF]
  `

function PausedButton({ networkName }) {
  return (
    <Button
      disabled={true}
      fancy={true}
      type="button"
      className={`${PAUSED_BASE_PROPERTIES}`}
    >
      Temporarily paused due to chain connectivity issues
    </Button>
  )
}

function HeavyLoadButton() {
  return (
    <Button
      disabled={true}
      fancy={true}
      type="button"
      className={ACTION_BTN_CLASSNAME}
    >
      Synapse is experiencing heavy load
    </Button>
  )
}

function AdvancedOptionsButton({ className, onClick }) {
  return (
    <div
      className={`
        group rounded-lg hover:bg-gray-900 ${className} p-1`}
      onClick={onClick}
    >
      <CogIcon className="w-6 h-6 text-gray-500 group-hover:text-gray-300" />
    </div>
  )
}
