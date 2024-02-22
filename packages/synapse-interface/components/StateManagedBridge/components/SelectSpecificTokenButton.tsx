import _ from 'lodash'
import { memo, useEffect, useRef, useState } from 'react'

import {
  getBorderStyleForCoin,
  getBorderStyleForCoinHover,
  getMenuItemBgForCoin,
  getMenuItemStyleForCoin,
} from '@styles/tokens'
import { Token } from '@/utils/types'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import { findChainIdsWithPausedToken } from '@/constants/tokens'
import LoadingDots from '@tw/LoadingDots'
import {
  BridgeModules,
  ELIGIBILITY_DEFAULT_TEXT,
  useStipEligibility,
} from '@/utils/hooks/useStipEligibility'
import { getUnderlyingBridgeTokens } from '@/utils/getUnderlyingBridgeTokens'
import { ARBITRUM, AVALANCHE, ETH } from '@/constants/chains/master'
import { AvailableChains } from '@/components/bridgeSwap/TokenButton/AvailableChains'
import { TokenBalance } from '@/components/bridgeSwap/TokenButton/TokenBalance'
export const SelectSpecificTokenButton = ({
  showAllChains,
  isOrigin,
  token,
  active,
  selectedToken,
  onClick,
  alternateBackground = false,
  isLoadingExchangeRate = false,
  exchangeRate,
  isBestExchangeRate = false,
  estimatedDurationInSeconds,
}: {
  showAllChains?: boolean
  isOrigin: boolean
  token: Token
  active: boolean
  selectedToken: Token
  onClick: () => void
  alternateBackground?: boolean
  isLoadingExchangeRate?: boolean
  exchangeRate?: string
  isBestExchangeRate?: boolean
  estimatedDurationInSeconds?: number
}) => {
  const ref = useRef<any>(null)
  const isCurrentlySelected = selectedToken?.routeSymbol === token?.routeSymbol
  const { fromChainId, toChainId, fromToken, toToken } = useBridgeState()

  useEffect(() => {
    if (active) {
      ref?.current?.focus()
    }
  }, [active])

  const chainId = isOrigin ? fromChainId : toChainId

  let bgClassName

  const classNameForMenuItemStyle = getMenuItemStyleForCoin(token?.color)

  if (isCurrentlySelected) {
    bgClassName = `${getMenuItemBgForCoin(
      token?.color
    )} ${getBorderStyleForCoin(token?.color)}`
  } else {
    bgClassName = getBorderStyleForCoinHover(token?.color)
  }

  return (
    <button
      data-test-id="select-specific-token-button"
      ref={ref}
      tabIndex={active ? 1 : 0}
      onClick={onClick}
      className={`
        flex items-center
        transition-all duration-75
        w-full
        px-2 py-1
        cursor-pointer
        rounded-md
        border border-slate-400/10
        mb-1
        ${alternateBackground ? '' : !isCurrentlySelected && 'bg-slate-400/10' }
        ${bgClassName}
        ${classNameForMenuItemStyle}
      `}
    >
      <ButtonContent
        token={token}
        chainId={chainId}
        isOrigin={isOrigin}
        showAllChains={showAllChains}
      />
      {isLoadingExchangeRate ? (
        <LoadingDots className="mr-8 opacity-50" />
      ) : (
        <>
          {exchangeRate && (
            isBestExchangeRate
              ? <OptionTag type={BestOptionType.RATE} />
              :
                <OptionDetails
                  exchangeRate={exchangeRate}
                  estimatedDurationInSeconds={estimatedDurationInSeconds}
                />
          )}
        </>
      )}
    </button>
  )
}

export enum BestOptionType {
  RATE = 'Best rate',
  SPEED = 'Fastest',
}

export const OptionTag = ({ type }: { type: BestOptionType }) => {
  return (
    <div
      data-test-id="option-tag"
      className="flex px-3 py-0.5 mr-3 text-sm whitespace-nowrap text-primary rounded-xl"
      style={{
        background:
          'linear-gradient(to right, rgba(128, 0, 255, 0.2), rgba(255, 0, 191, 0.2))',
      }}
    >{`${type}`}</div>
  )
}

export const OptionDetails = ({
  exchangeRate,
  estimatedDurationInSeconds,
}: {
  exchangeRate: string
  estimatedDurationInSeconds: number
}) => {
  let showTime
  let timeUnit

  if (estimatedDurationInSeconds > 60) {
    showTime = Math.floor(estimatedDurationInSeconds / 60)
    timeUnit = 'min'
  } else {
    showTime = estimatedDurationInSeconds
    timeUnit = 'seconds'
  }

  return (
    <div data-test-id="option-details" className="flex flex-col">
      <div className="flex items-center font-normal">
        <div className="flex text-sm text-secondary whitespace-nowrap">
          1&nbsp;:&nbsp;
        </div>
        <div className="mb-[1px] text-primary">{exchangeRate}</div>
      </div>
      <div className="text-xs text-right text-secondary">
        {showTime} {timeUnit}
      </div>
    </div>
  )
}

const ButtonContent = memo(
  ({
    token,
    chainId,
    isOrigin,
    showAllChains,
  }: {
    token: Token
    chainId: number
    isOrigin: boolean
    showAllChains: boolean
  }) => {
    const portfolioBalances = usePortfolioBalances()

    const parsedBalance = portfolioBalances[chainId]?.find(
      (tb) => tb.token.addresses[chainId] === token.addresses[chainId]
    )?.parsedBalance

    return (
      <div data-test-id="button-content" className="flex items-center w-full">
        <img
          alt="token image"
          className="w-8 h-8 ml-2 mr-4 rounded-full"
          src={token?.icon?.src}
        />
        <Coin token={token} showAllChains={showAllChains} isOrigin={isOrigin} />
        {isOrigin && (
          <TokenBalance token={token} parsedBalance={parsedBalance} />
        )}
      </div>
    )
  }
)

const Coin = ({
  token,
  showAllChains,
  isOrigin,
}: {
  token
  showAllChains: boolean
  isOrigin: boolean
}) => {
  const isEligible = isTokenEligible(token)

  return (
    <div className="flex-col text-left">
      <div className="text-lg text-primaryTextColor">{token?.symbol}</div>
      <div className="flex items-center space-x-2 text-xs text-secondaryTextColor">
        {isOrigin && isEligible ? (
          <div className="text-greenText">{ELIGIBILITY_DEFAULT_TEXT}</div>
        ) : (
          <div>{token?.name}</div>
        )}
        {showAllChains &&
          <AvailableChains
            token={token}
            excludedChainIds={findChainIdsWithPausedToken(token.routeSymbol)}
          />
        }
      </div>
    </div>
  )
}

/*
Synapse:Bridge
  Tokens: nETH, nUSD, GMX
  From Any to ARB: all txs (don't limit this to "user has to receive ETH / USDC / ...")
  ARB to ETH txs: nETH, nUSD
  ARB to AVAX txs: GMX

Synapse:CCTP
  Tokens: USDC
  Any to ARB: all txs
  ARB to ETH: all txs

Synapse: RFQ
  Tokens: USDC
  Any to ARB: all txs
  ARB to ETH: all txs
*/

const isTokenEligible = (token: Token) => {
  const { fromChainId, toChainId, bridgeQuote } = useBridgeState()

  const underlyingBridgeTokens = getUnderlyingBridgeTokens(token, fromChainId)

  if (!underlyingBridgeTokens) {
    return false
  }

  const includesUSDC = underlyingBridgeTokens.includes('USDC')

  return (
    (includesUSDC && toChainId === ARBITRUM.id) ||
    (includesUSDC &&
      fromChainId === ARBITRUM.id &&
      toChainId === ETH.id) ||
    (includesUSDC && toChainId === ARBITRUM.id) ||
    (includesUSDC &&
      fromChainId === ARBITRUM.id &&
      toChainId === ETH.id) ||
    (_.some(['nETH', 'nUSD', 'GMX'], (value) =>
      _.includes(underlyingBridgeTokens, value)
    ) &&
      toChainId === ARBITRUM.id) ||
    (_.some(['nETH', 'nUSD'], (value) =>
      _.includes(underlyingBridgeTokens, value)
    ) &&
      fromChainId === ARBITRUM.id &&
      toChainId === ETH.id) ||
    (_.some(['GMX'], (value) => _.includes(underlyingBridgeTokens, value)) &&
      fromChainId === ARBITRUM.id &&
      toChainId === AVALANCHE.id)
  )
}

