import {
  getBorderStyleForCoin,
  getBorderStyleForCoinHover,
  getMenuItemBgForCoin,
  getMenuItemStyleForCoin,
} from '@styles/tokens'
import { memo, useEffect, useRef, useState } from 'react'
import { Token } from '@/utils/types'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import { useSwapState } from '@/slices/swap/hooks'

export const SelectSpecificTokenButton = ({
  showAllChains,
  isOrigin,
  token,
  active,
  selectedToken,
  onClick,
  alternateBackground = false,
}: {
  showAllChains?: boolean
  isOrigin: boolean
  token: Token
  active: boolean
  selectedToken: Token
  onClick: () => void
  alternateBackground?: boolean
}) => {
  const ref = useRef<any>(null)
  const isCurrentlySelected = selectedToken?.routeSymbol === token?.routeSymbol
  const { swapChainId, swapFromToken, swapToToken } = useSwapState()

  useEffect(() => {
    if (active) {
      ref?.current?.focus()
    }
  }, [active])

  const chainId = swapChainId

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
    </button>
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
      <div className="flex items-center w-full">
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
      <div className="text-lg text-primaryTextColor">{token?.symbol}</div>
      <div className="flex items-center space-x-2 text-xs text-secondaryTextColor">
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
    <div className="ml-auto mr-5 text-md text-primaryTextColor">
      {parsedBalance && parsedBalance !== '0.0' && (
        <div>
          {parsedBalance}
          <span className="text-md text-secondaryTextColor">
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
  const chainIds = Object.keys(token.addresses)
  const hasOneChain = chainIds.length > 0
  const hasMultipleChains = chainIds.length > 1
  const numOverTwoChains = chainIds.length - 2 > 0 ? chainIds.length - 2 : 0

  return (
    <div
      data-test-id="portfolio-token-visualizer"
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
