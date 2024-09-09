import _ from 'lodash'
import { memo, useRef } from 'react'
import Image from 'next/image'

import { type Token, type ActionTypes } from '@/utils/types'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import { findChainIdsWithPausedToken } from '@/constants/tokens'
import { getActiveStyleForButton, getHoverStyleForButton } from '@/styles/hover'
import { joinClassNames } from '@/utils/joinClassNames'
import { useSwapState } from '@/slices/swap/hooks'
import { formatAmount } from '@/utils/formatAmount'
import { getParsedBalance } from '@/utils/getParsedBalance'

export const SelectSpecificTokenButton = ({
  showAllChains,
  isOrigin,
  token,
  onClick,
  action,
  isSelected,
  isActive,
}: {
  showAllChains?: boolean
  isOrigin: boolean
  token: Token
  onClick: () => void
  action: ActionTypes
  isSelected: boolean
  isActive: boolean
}) => {
  const ref = useRef<any>(null)
  const { fromChainId, toChainId } = useBridgeState()

  const { swapChainId } = useSwapState()

  const buttonClasses = {
    other: 'whitespace-nowrap',
    grid: 'grid gap-0.5',
    space: 'pl-2 pr-1.5 py-2.5 w-full',
    border: 'border border-transparent',
    transition: 'transition-all duration-75',
    hover: getHoverStyleForButton(token?.color),
    activeStyle:
      isActive || isSelected ? getActiveStyleForButton(token?.color) : '',
  }

  const chainId =
    action === 'Swap' ? swapChainId : isOrigin ? fromChainId : toChainId

  return (
    <button
      data-test-id="select-specific-token-button"
      ref={ref}
      onClick={onClick}
      className={joinClassNames(buttonClasses)}
    >
      <ButtonContent
        token={token}
        chainId={chainId}
        isOrigin={isOrigin}
        showAllChains={showAllChains}
        action={action}
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
    action,
  }: {
    token: Token
    chainId: number
    isOrigin: boolean
    showAllChains: boolean
    action: ActionTypes
  }) => {
    const portfolioBalances = usePortfolioBalances()

    const tokenData = portfolioBalances[chainId]?.find(
      (tb) => tb.token.addresses[chainId] === token.addresses[chainId]
    )
    const decimals = tokenData?.token?.decimals[chainId]
    const balance = tokenData?.balance
    const parsedBalance = getParsedBalance(balance, decimals)
    const formattedBalance = formatAmount(parsedBalance)

    return (
      <div data-test-id="button-content" className="">
        <Coin
          token={token}
          showAllChains={showAllChains}
          isOrigin={isOrigin}
          parsedBalance={formattedBalance}
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
