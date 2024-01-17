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
import LoadingDots from '@/components/ui/tailwind/LoadingDots'

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
        border-[1px] border-zinc-300 dark:border-zinc-700
        mb-1
        ${alternateBackground && 'bg-zinc-100 dark:bg-zinc-900'}
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
      className="flex px-3 py-0.5 mr-3 text-sm whitespace-nowrap rounded-xl"
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
  const estimatedDurationInMinutes: number =
    estimatedDurationInSeconds < 60
      ? Math.ceil(estimatedDurationInSeconds / 60)
      : Math.floor(estimatedDurationInSeconds / 60)

  return (
    <div data-test-id="option-details" className="flex flex-col">
      <div className="flex items-center font-normal">
        <div className="flex text-sm text-secondary whitespace-nowrap">
          1&nbsp;:&nbsp;
        </div>
        <div className="mb-[1px]">{exchangeRate}</div>
      </div>
      <div className="text-sm text-right text-secondary">
        {estimatedDurationInMinutes} min
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
        <Coin token={token} showAllChains={showAllChains} />
        {isOrigin && (
          <TokenBalance
            token={token}
            chainId={chainId}
            parsedBalance={parsedBalance}
          />
        )}
      </div>
    )
  }
)

const Coin = ({ token, showAllChains }: { token; showAllChains: boolean }) => {
  return (
    <div className="flex-col text-left">
      <div className="text-lg">{token?.symbol}</div>
      <div className="flex items-center space-x-2 text-xs text-zinc-700 dark:text-zinc-400">
        <div>{token?.name}</div>
        {showAllChains && <AvailableChains token={token} />}
      </div>
    </div>
  )
}

const TokenBalance = ({
  token,
  chainId,
  parsedBalance,
}: {
  token: Token
  chainId: number
  parsedBalance?: string
}) => {
  return (
    <div className="ml-auto mr-5 text-md">
      {parsedBalance && parsedBalance !== '0.0' && (
        <div>
          {parsedBalance}
          <span className="text-md text-zinc-700 dark:text-zinc-400">
            {' '}
            {token ? token.symbol : ''}
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
      className="flex flex-row items-center space-x-1 hover-trigger"
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
              border border-solid border-zinc-300 dark:border-zinc-700
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
