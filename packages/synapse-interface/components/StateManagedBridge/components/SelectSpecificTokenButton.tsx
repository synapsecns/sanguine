import _ from 'lodash'
import { memo, useEffect, useRef, useState } from 'react'
import Image from 'next/image'

import { type Token, type ActionTypes } from '@/utils/types'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import { findChainIdsWithPausedToken } from '@/constants/tokens'
import LoadingDots from '@/components/ui/tailwind/LoadingDots'
import { ELIGIBILITY_DEFAULT_TEXT } from '@/utils/hooks/useStipEligibility'
import { getUnderlyingBridgeTokens } from '@/utils/getUnderlyingBridgeTokens'
import { ARBITRUM, AVALANCHE, ETH } from '@/constants/chains/master'
import { getActiveStyleForButton, getHoverStyleForButton } from '@/styles/hover'
import { joinClassNames } from '@/utils/joinClassNames'
import { useSwapState } from '@/slices/swap/hooks'

const SelectSpecificTokenButton = ({
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
  action,
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
  action: ActionTypes
}) => {
  const ref = useRef<any>(null)
  const isCurrentToken = selectedToken?.routeSymbol === token?.routeSymbol
  const { fromChainId, toChainId } = useBridgeState()

  const { swapChainId } = useSwapState()

  useEffect(() => {
    if (active) {
      ref?.current?.focus()
    }
  }, [active])

  const buttonClass = joinClassNames({
    other: 'whitespace-nowrap',
    grid: 'grid gap-0.5',
    space: 'pl-2 pr-1.5 py-2.5 w-full',
    border: 'border border-transparent',
    transition: 'transition-all duration-75',
    hover: getHoverStyleForButton(token.color),
    activeStyle: isCurrentToken
      ? getActiveStyleForButton(isCurrentToken && token.color)
      : '',
  })

  const chainId =
    action === 'Swap' ? swapChainId : isOrigin ? fromChainId : toChainId

  return (
    <button
      data-test-id="select-specific-token-button"
      ref={ref}
      tabIndex={active ? 1 : 0}
      onClick={onClick}
      className={buttonClass}
    >
      <ButtonContent
        token={token}
        chainId={chainId}
        isOrigin={isOrigin}
        showAllChains={showAllChains}
        action={action}
      />
      {isLoadingExchangeRate ? (
        <LoadingDots className="mr-8 opacity-50" />
      ) : (
        <>
          {exchangeRate && isBestExchangeRate && (
            <OptionTag type={BestOptionType.RATE} />
          )}

          {exchangeRate && (
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
    action,
  }: {
    token: Token
    chainId: number
    isOrigin: boolean
    showAllChains: boolean
    action: ActionTypes
  }) => {
    const portfolioBalances = usePortfolioBalances()

    console.log(`chainId`, chainId)

    const parsedBalance = portfolioBalances[chainId]?.find(
      (tb) => tb.token.addresses[chainId] === token.addresses[chainId]
    )?.parsedBalance

    return (
      <div data-test-id="button-content" className="">
        <div className="flex items-center justify-between">
          <span className="flex items-center gap-2">
            <Image
              loading="lazy"
              src={token.icon.src}
              alt="Token Image"
              width="20"
              height="20"
              className="w-5 h-5 max-w-fit"
            />
            <Coin
              token={token}
              showAllChains={showAllChains}
              isOrigin={isOrigin}
            />
          </span>
          {showAllChains && <AvailableChains token={token} />}
          {isOrigin && (
            <TokenBalance token={token} parsedBalance={parsedBalance} />
          )}
        </div>
        <div className="flex items-center space-x-2 text-sm text-secondary">
          {action === 'Bridge' && isOrigin && isTokenEligible(token) ? (
            <div className="text-greenText">{ELIGIBILITY_DEFAULT_TEXT}</div>
          ) : (
            <></>
          )}
        </div>
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
  return (
    <div>
      <div className="flex justify-between text-left">
        <div className="">{token?.symbol}</div>

        {/* {showAllChains && <AvailableChains token={token} />} */}
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

  return (
    (underlyingBridgeTokens.includes('USDC') && toChainId === ARBITRUM.id) ||
    (underlyingBridgeTokens.includes('USDC') &&
      fromChainId === ARBITRUM.id &&
      toChainId === ETH.id) ||
    (underlyingBridgeTokens.includes('USDC') && toChainId === ARBITRUM.id) ||
    (underlyingBridgeTokens.includes('USDC') &&
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

const TokenBalance = ({
  token,
  parsedBalance,
}: {
  token: Token
  parsedBalance?: string
}) => {
  return (
    <div className="p-1 text-sm">
      {parsedBalance && parsedBalance !== '0.0' && (
        <div>
          {parsedBalance}
          <span className="text-md text-secondaryTextColor">
            {/* {' '}
            {token ? token.symbol : ''} */}
          </span>
        </div>
      )}
    </div>
  )
}

const AvailableChains = ({ token }: { token: Token }) => {
  const [isHovered, setIsHovered] = useState(false)
  const pausedChainIds = findChainIdsWithPausedToken(token.routeSymbol)
  const chainIds = _.difference(Object.keys(token.addresses), pausedChainIds)
  const hasOneChain = chainIds.length > 0
  const hasMultipleChains = chainIds.length > 1
  const numOverTwoChains = chainIds.length - 2 > 0 ? chainIds.length - 2 : 0

  return (
    <div
      data-test-id="available-chains"
      className="flex items-center space-x-1 text-sm hover-trigger"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
    >
      {hasOneChain && (
        <img
          className="w-3 h-3 rounded-md"
          alt={`${CHAINS_BY_ID[chainIds[0]].name} img`}
          src={`${CHAINS_BY_ID[chainIds[0]].chainImg.src}`}
        />
      )}
      {hasMultipleChains && (
        <img
          className="w-3 h-3 rounded-md"
          alt={`${CHAINS_BY_ID[chainIds[1]].name} img`}
          src={`${CHAINS_BY_ID[chainIds[1]].chainImg.src}`}
        />
      )}
      {numOverTwoChains > 0 && (
        <div className="ml-1 text-white">+ {numOverTwoChains}</div>
      )}
      <div className="relative inline-block">
        {isHovered && (
          <div
            className={`
              absolute z-50 hover-content p-2 text-white
              border border-solid border-[#252537]
              bg-[#101018] rounded-md
            `}
          >
            {chainIds.map((chainId) => {
              const chainName = CHAINS_BY_ID[chainId].name
              return <div className="whitespace-nowrap">{chainName}</div>
            })}
          </div>
        )}
      </div>
    </div>
  )
}

export default SelectSpecificTokenButton
