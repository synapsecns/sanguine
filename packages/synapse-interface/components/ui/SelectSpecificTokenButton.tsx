import _ from 'lodash'
import { memo, useEffect, useRef } from 'react'
import Image from 'next/image'

import { type Token, type ActionTypes } from '@/utils/types'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import { findChainIdsWithPausedToken } from '@/constants/tokens'
import LoadingDots from '@/components/ui/tailwind/LoadingDots'
import { getActiveStyleForButton, getHoverStyleForButton } from '@/styles/hover'
import { joinClassNames } from '@/utils/joinClassNames'
import { useSwapState } from '@/slices/swap/hooks'

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

    const parsedBalance = portfolioBalances[chainId]?.find(
      (tb) => tb.token.addresses[chainId] === token.addresses[chainId]
    )?.parsedBalance

    return (
      <div data-test-id="button-content" className="">
        <Coin
          token={token}
          showAllChains={showAllChains}
          isOrigin={isOrigin}
          parsedBalance={parsedBalance}
        />
      </div>
    )
  }
)

const Coin = ({
  token,
  showAllChains,
  isOrigin,
  parsedBalance,
}: {
  token
  showAllChains: boolean
  isOrigin: boolean
  parsedBalance: string | undefined
}) => {
  return (
    <div className="flex justify-between">
      <div className="flex items-center gap-2">
        <Image
          loading="lazy"
          src={token.icon.src}
          alt="Token Image"
          width="20"
          height="20"
          className="w-5 h-5 max-w-fit"
        />
        <div>{token?.symbol}</div>
      </div>

      {!showAllChains && isOrigin && (
        <TokenBalance parsedBalance={parsedBalance} />
      )}
      {showAllChains && <AvailableChains token={token} />}
    </div>
  )
}

const TokenBalance = ({ parsedBalance }: { parsedBalance?: string }) => {
  return (
    <div className="p-1 text-sm">
      {parsedBalance && parsedBalance !== '0.0' && <div>{parsedBalance}</div>}
    </div>
  )
}

const AvailableChains = ({ token }: { token: Token }) => {
  const pausedChainIds = findChainIdsWithPausedToken(token.routeSymbol)
  const chainIds = _.difference(Object.keys(token.addresses), pausedChainIds)
  const hasOneChain = chainIds.length > 0
  const hasMultipleChains = chainIds.length > 1
  const numOverTwoChains = chainIds.length - 2 > 0 ? chainIds.length - 2 : 0

  return (
    <div
      data-test-id="available-chains"
      className="flex items-center space-x-1 text-sm hover-trigger"
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
    </div>
  )
}
