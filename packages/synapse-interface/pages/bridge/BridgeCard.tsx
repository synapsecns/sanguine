import React, { useEffect, useState } from 'react'
import { useSettings } from '@hooks/useSettings'
import { SettingsIcon } from '@icons/SettingsIcon'
import { Transition } from '@headlessui/react'
import { validateAndParseAddress } from '@utils/validateAndParseAddress'
// import { TRANSITIONS_PROPS } from '@constants/bridge'
import { COIN_SLIDE_OVER_PROPS } from '@styles/transitions'
import { ORDERED_CHAINS_BY_ID } from '@constants/chains'
import Grid from '@tw/Grid'
import Card from '@tw/Card'
import Button from '@tw/Button'
import ExchangeRateInfo from '@components/ExchangeRateInfo'
import { TransactionButton } from '@components/buttons/SubmitTxButton'
import { PageHeader } from '@components/PageHeader'
import { TokenSlideOver } from '@/components/misc/TokenSlideOver'
import { ChainSlideOver } from '@/components/misc/ChainSlideOver'
import { BigNumber } from '@ethersproject/bignumber'
import { Zero, MaxInt256 } from '@ethersproject/constants'
import { formatBNToString } from '@bignumber/format'
import { SECTION_TRANSITION_PROPS } from '@styles/transitions'
import { approveToken } from '@/utils/approveToken'
import SettingsSlideOver from './SettingsSlideOver'
import { DestinationAddressInput } from '../../components/input/DestinationAddressInput'
import BridgeInputContainer from '../../components/input/TokenAmountInput'
import { TransactionResponse } from '@ethersproject/providers'

import { Token } from '@/utils/types'
import { BridgeQuote } from '@/utils/types'
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
  handleChainChange,
  handleTokenChange,
  onChangeFromAmount,
  setDestinationAddress,
  executeBridge,
  resetRates,
  setTime,
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
  handleChainChange: (
    chainId: number,
    flip: boolean,
    type: 'from' | 'to'
  ) => void
  handleTokenChange: (token: Token, type: 'from' | 'to') => void
  onChangeFromAmount: (amount: string) => void
  setDestinationAddress: (address: string) => void
  executeBridge: () => Promise<TransactionResponse>
  resetRates: () => void
  setTime: (time: number) => void
}) => {
  const [settings, setSettings] = useSettings()
  const [displayType, setDisplayType] = useState('')
  const [deadlineMinutes, setDeadlineMinutes] = useState('')
  const [fromTokenBalance, setFromTokenBalance] = useState<BigNumber>(Zero)

  /*
  useEffect Trigger: fromToken, fromTokens
  - When either the from token or list of from tokens are mutated, the selected token's balance is set in state
  this is for checking max bridge possible as well as for producing the option to select max bridge
  */
  useEffect(() => {
    if (fromTokens && fromToken) {
      setFromTokenBalance(
        fromTokens.filter((token) => token.token === fromToken)[0]?.balance
          ? fromTokens.filter((token) => token.token === fromToken)[0]?.balance
          : Zero
      )
    }
  }, [fromToken, fromTokens])

  /*
  Constant: fromArgs, toArgs
  - Define various arguments to numerous bridge ui components. Defined here to prevent messy code.
   */
  const fromArgs = {
    address,
    fromTokenBalance,
    isOrigin: true,
    isSwap: false,
    chains: ORDERED_CHAINS_BY_ID.filter((id) => id !== String(fromChainId)),
    tokens: fromTokens,
    chainId: fromChainId,
    inputString: fromInput?.string,
    selectedToken: fromToken,
    connectedChainId: fromChainId,
    setDisplayType,
    handleTokenChange,
    onChangeChain: handleChainChange,
    onChangeAmount: onChangeFromAmount,
  }

  const toArgs = {
    address,
    isOrigin: false,
    isSwap: false,
    chains: toOptions?.chains,
    tokens: toOptions?.tokens,
    chainId: toChainId,
    inputString: bridgeQuote?.outputAmountString,
    selectedToken: toToken,
    connectedChainId: fromChainId,
    setDisplayType,
    handleTokenChange,
    onChangeChain: handleChainChange,
  }

  // TODO move this away and into the actual component
  const settingsArgs = {
    settings,
    setSettings,
    setDisplayType,
    setDestinationAddress,
    deadlineMinutes,
    setDeadlineMinutes,
  }

  // some messy button gen stuff (will re-write)
  // maybe just put everything in index without the card
  const isFromBalanceEnough = fromTokenBalance.gt(fromInput?.bigNum ?? Zero)
  let destAddrNotValid
  let btnLabel
  let btnClassName = ''
  let pendingLabel = 'Bridging funds...'
  let buttonAction = () => executeBridge()
  let postButtonAction = () => resetRates()
  if (error) {
    btnLabel = error
  } else if (!isFromBalanceEnough) {
    btnLabel = `Insufficient ${fromToken?.symbol} Balance`
  } else if (bridgeQuote.feeAmount.eq(0) && !fromInput?.bigNum?.eq(0)) {
    btnLabel = `Amount must be greater than fee`
  } else if (
    bridgeQuote?.allowance &&
    bridgeQuote?.allowance?.lt(fromInput?.bigNum)
  ) {
    buttonAction = () =>
      approveToken(
        bridgeQuote?.routerAddress,
        fromChainId,
        fromToken.addresses[fromChainId]
      )
    btnLabel = `Approve ${fromToken?.symbol}`
    pendingLabel = `Approving ${fromToken?.symbol}`
    btnClassName = 'from-[#feba06] to-[#FEC737]'
    postButtonAction = () => setTime(0)
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
      formatBNToString(bridgeQuote?.exchangeRate, 18, 4)
    )

    if (
      !fromInput?.bigNum?.eq(0) &&
      (numExchangeRate < 0.95 || numExchangeRate > 1.05)
    ) {
      btnClassName = 'from-[#fe064a] to-[#fe5281]'
      btnLabel = 'Slippage High - Bridge Anyway?'
    }
  }

  const actionBtn = (
    <TransactionButton
      className={btnClassName}
      disabled={
        fromChainId === toChainId ||
        bridgeQuote.outputAmount.eq(0) ||
        !isFromBalanceEnough ||
        error != null ||
        destAddrNotValid
      }
      onClick={() => buttonAction()}
      onSuccess={() => {
        postButtonAction()
      }}
      label={btnLabel}
      pendingLabel={pendingLabel}
    />
  )
  const TRANSITION_PROPS = {
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
          <Transition show={displayType === 'from'} {...TRANSITION_PROPS}>
            <TokenSlideOver key="fromBlock" {...fromArgs} />{' '}
          </Transition>
          <Transition show={displayType === 'to'} {...TRANSITION_PROPS}>
            <TokenSlideOver key="toBlock" {...toArgs} />
          </Transition>
          <Transition show={displayType === 'fromChain'} {...TRANSITION_PROPS}>
            <ChainSlideOver key="fromChainBlock" {...fromArgs} />
          </Transition>
          <Transition show={displayType === 'toChain'} {...TRANSITION_PROPS}>
            <ChainSlideOver key="fromChainBlock" {...toArgs} />
          </Transition>
          <Transition show={displayType === 'settings'} {...TRANSITION_PROPS}>
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
            show={!fromInput?.bigNum?.eq(0)}
            {...SECTION_TRANSITION_PROPS}
          >
            <ExchangeRateInfo
              fromAmount={fromInput?.bigNum}
              toToken={toToken}
              exchangeRate={bridgeQuote?.exchangeRate}
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
export default BridgeCard
