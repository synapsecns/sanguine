import { useEffect, useMemo, useRef, useState } from 'react'
import { CHAINS_BY_ID } from '@constants/chains'
import Image from 'next/image'
// import {
//   getNetworkHover,
//   getNetworkButtonBorder,
//   getNetworkButtonBorderHover,
//   getNetworkButtonBgClassName,
//   getNetworkButtonBgClassNameActive,
//   getNetworkButtonBorderActive,
//   getMenuItemStyleForChain,
// } from '@/styles/chains'
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

  useEffect(() => {
    if (active) {
      ref?.current?.focus()
    }
  }, [active])

  const bgClassName = isCurrentChain
    ? getActiveStyleForButton(isCurrentChain && chain.color)
    : ''

  return (
    <button
      ref={ref}
      tabIndex={active ? 1 : 0}
      className={`
        flex items-center justify-between
        transition-all duration-75
        w-full py-2 my-1

        ${alternateBackground && 'bg-[#282328]'}
        ${bgClassName}
        ${getHoverStyleForButton(chain.color)}
      `}
      onClick={onClick}
      data-test-id={`${dataId}-item`}
    >
      <ButtonContent chainId={itemChainId} isOrigin={isOrigin} />
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
  const { fromChainId, fromToken } = useBridgeState()

  const balanceTokens =
    balances &&
    balances[chainId] &&
    sortTokensByBalanceDescending(
      balances[chainId].filter((bt) => bt.balance > 0n)
    )

  const isEligible = isChainEligible(fromChainId, chain.id, fromToken)

  return chain ? (
    <>
      <div className="flex items-center space-x-2">
        <Image
          src={chain.chainImg}
          alt="Switch Network"
          className="ml-2 rounded-full w-5 h-5"
        />
        <div className="flex-col text-left">
          {chain.name}
          {!isOrigin && isEligible && (
            <div className="text-sm text-greenText">
              {ELIGIBILITY_DEFAULT_TEXT}
            </div>
          )}
        </div>
      </div>
      {isOrigin && balanceTokens && balanceTokens.length > 0 ? (
        <ChainTokens balanceTokens={balanceTokens} />
      ) : null}
    </>
  ) : null
}

const ChainTokens = ({
  balanceTokens = [],
}: {
  balanceTokens: TokenAndBalance[]
}) => {
  const [isT1Hovered, setIsT1Hovered] = useState<boolean>(false)
  const [isT2Hovered, setIsT2Hovered] = useState<boolean>(false)
  const [isT3Hovered, setIsT3Hovered] = useState<boolean>(false)

  const hasNoTokens: boolean =
    !balanceTokens || (balanceTokens && balanceTokens.length === 0)
  const hasOneToken: boolean = balanceTokens && balanceTokens.length > 0
  const hasTwoTokens: boolean = balanceTokens && balanceTokens.length > 1
  const numOverTwoTokens: number =
    balanceTokens && balanceTokens.length - 2 > 0 ? balanceTokens.length - 2 : 0
  const hasOnlyOneToken: boolean = balanceTokens && balanceTokens.length === 1
  const hasOnlyTwoTokens: boolean = balanceTokens && balanceTokens.length === 2

  if (hasNoTokens) {
    return (
      <div
        data-test-id="portfolio-token-visualizer"
        className="flex flex-row items-center mr-4 cursor-pointer hover-trigger text-secondary"
      >
        -
      </div>
    )
  }
  return (
    <div
      data-test-id="portfolio-token-visualizer"
      className="flex flex-row items-center space-x-2 cursor-pointer hover-trigger"
    >
      {hasOneToken && (
        <div>
          <Image
            loading="lazy"
            className="w-5 h-5 rounded-md"
            alt={`${balanceTokens[0].token.symbol} img`}
            src={balanceTokens[0].token.icon}
            onMouseEnter={() => setIsT1Hovered(true)}
            onMouseLeave={() => setIsT1Hovered(false)}
          />
          <div className="relative">
            <HoverContent isHovered={isT1Hovered}>
              <div className="whitespace-nowrap">
                {balanceTokens[0]?.parsedBalance}{' '}
                {balanceTokens[0]?.token.symbol}
              </div>
            </HoverContent>
          </div>
        </div>
      )}
      {hasOnlyOneToken && (
        <div className="text-white whitespace-nowrap">
          {balanceTokens[0].parsedBalance} {balanceTokens[0].token.symbol}
        </div>
      )}
      {hasTwoTokens && (
        <div>
          <Image
            loading="lazy"
            className="w-5 h-5 rounded-md"
            alt={`${balanceTokens[1].token.symbol} img`}
            src={balanceTokens[1].token.icon}
            onMouseEnter={() => setIsT2Hovered(true)}
            onMouseLeave={() => setIsT2Hovered(false)}
          />
          <div className="relative">
            <HoverContent isHovered={isT2Hovered}>
              <div className="whitespace-nowrap">
                {balanceTokens[1]?.parsedBalance}{' '}
                {balanceTokens[1]?.token.symbol}
              </div>
            </HoverContent>
          </div>
        </div>
      )}
      {numOverTwoTokens > 0 && (
        <div
          className="text-white"
          onMouseEnter={() => setIsT3Hovered(true)}
          onMouseLeave={() => setIsT3Hovered(false)}
        >
          + {numOverTwoTokens}
        </div>
      )}
      <div className="relative inline-block">
        <HoverContent isHovered={isT3Hovered}>
          {balanceTokens?.map((token: TokenAndBalance, key: number) => {
            if (key > 1) {
              const tokenSymbol = token.token.symbol
              const balance = token.parsedBalance
              return (
                <div className="whitespace-nowrap" key={key}>
                  {balance} {tokenSymbol}
                </div>
              )
            }
          })}
        </HoverContent>
      </div>
    </div>
  )
}

export const HoverContent = ({
  isHovered,
  children,
}: {
  isHovered: boolean
  children: React.ReactNode
}) => {
  if (isHovered) {
    return (
      <div
        className={`
          absolute -ml-28 z-50 hover-content p-2 text-white
          border border-solid border-[#252537]
          bg-[#101018] rounded-md text-left
        `}
      >
        {children}
      </div>
    )
  }
}
