import { useEffect, useRef } from 'react'
import { CHAINS_BY_ID } from '@constants/chains'
import Image from 'next/image'
import { getHoverStyleForButton, getActiveStyleForButton } from '@/styles/hover'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import {
  TokenAndBalance,
  sortTokensByBalanceDescending,
} from '@/utils/actions/fetchPortfolioBalances'
import {
  ELIGIBILITY_DEFAULT_TEXT,
  isChainEligible,
  useStipEligibility,
} from '@/utils/hooks/useStipEligibility'
import { useBridgeState } from '@/slices/bridge/hooks'
import { joinClassNames } from '@/utils/joinClassNames'

export const SelectSpecificNetworkButton = ({
  itemChainId,
  isCurrentChain,
  active,
  onClick,
  dataId,
  isOrigin,
  alternateBackground = false,
}: {
  itemChainId: number
  isCurrentChain: boolean
  active: boolean
  onClick: () => void
  dataId: string
  isOrigin: boolean
  alternateBackground?: boolean
}) => {
  const ref = useRef<any>(null)
  const chain = CHAINS_BY_ID[itemChainId]

  // useEffect(() => {
  //   if (active) {
  //     ref?.current?.focus()
  //   }
  // }, [active])

  const buttonClass = joinClassNames({
    other: 'whitespace-nowrap',
    grid: 'grid gap-0.5',
    space: 'pl-2 pr-1.5 py-2.5 w-full',
    border: 'border border-transparent',
    transition: 'transition-all duration-75',
    hover: getHoverStyleForButton(chain.color),
    activeStyle: isCurrentChain
      ? getActiveStyleForButton(isCurrentChain && chain.color)
      : '',
  })

  const { fromChainId, fromToken } = useBridgeState()
  const isEligible =
    !isOrigin && isChainEligible(fromChainId, chain.id, fromToken)

  return (
    <button
      ref={ref}
      // tabIndex={active ? 1 : 0}
      className={buttonClass}
      onClick={onClick}
      data-test-id={`${dataId}-item`}
    >
      <ButtonContent chainId={itemChainId} isOrigin={isOrigin} />
      {isEligible && (
        <div className="text-sm text-left text-green-500 dark:text-green-400">
          {ELIGIBILITY_DEFAULT_TEXT}
        </div>
      )}
    </button>
  )
}

function ButtonContent({
  chainId,
  isOrigin,
}: {
  chainId: number
  isOrigin: boolean
}) {
  const chain = CHAINS_BY_ID[chainId]
  const { balances } = usePortfolioState()
  // const { fromChainId, fromToken } = useBridgeState()

  const balanceTokens =
    balances?.[chainId] &&
    sortTokensByBalanceDescending(
      balances[chainId].filter((bt) => bt.balance > 0n)
    )

  // const isEligible = isChainEligible(fromChainId, chain.id, fromToken)

  return (
    chain && (
      <div className="flex items-center justify-between gap-6">
        <span className="flex items-center gap-2">
          <Image
            loading="lazy"
            src={chain.chainImg}
            alt="Switch Network"
            width="20"
            height="20"
            className="w-5 h-5 max-w-fit"
          />
          {chain.name}
        </span>
        {isOrigin && <ChainTokens balanceTokens={balanceTokens} />}
      </div>
    )
  )
}

const ChainTokens = ({
  balanceTokens,
}: {
  balanceTokens: TokenAndBalance[]
}) => {
  const max = balanceTokens?.length === 2 ? 2 : 1
  const remainder: number =
    balanceTokens && balanceTokens.length - max > 0
      ? balanceTokens.length - max
      : 0

  return (
    <span
      data-test-id="portfolio-token-visualizer"
      className="flex items-center cursor-pointer hover-trigger text-sm text-secondary -space-x-1.5"
    >
      {balanceTokens?.slice(0, max).map((token: TokenAndBalance) => {
        return <HoverIcon token={token} />
      })}
      {remainder > 0 && (
        <span className="relative">
          <div className="peer h-6 w-6 text-[13px] mb-px text-center grid place-content-center bg-bgBase rounded-full">
            {remainder}
          </div>
          <ul className="hidden peer-hover:block absolute z-50 bottom-6 -right-1 -mr-px pl-1 pr-1.5 py-1.5 bg-bgBase rounded text-right space-y-0.5 whitespace-normal max-w-40">
            {balanceTokens?.slice(max).map((token: TokenAndBalance) => (
              <li key={token.token.symbol} className="px-0.5">
                {token.token.symbol}
              </li>
            ))}
          </ul>
        </span>
      )}
    </span>
  )
}

function HoverIcon({ token }) {
  const symbol = token.token.symbol
  const src = token.token.icon
  const parsedBalance = token?.parsedBalance

  return (
    <span
      key={token.tokenAddress}
      className="relative flex justify-center text-center justify-items-center"
    >
      <Image
        loading="lazy"
        width="20"
        height="20"
        alt={`${symbol} img`}
        src={src}
        className="peer max-w-fit"
      />
      <div className="absolute z-50 hidden px-2 py-1 -mr-px rounded peer-hover:block bottom-6 -right-2 bg-bgBase">
        {parsedBalance} {symbol}
      </div>
    </span>
  )
}
