import { useEffect, useMemo, useRef, useState } from 'react'
import { CHAINS_BY_ID } from '@constants/chains'
import Image from 'next/image'
import {
  getNetworkHover,
  getNetworkButtonBorder,
  getNetworkButtonBorderHover,
  getNetworkButtonBgClassName,
  getNetworkButtonBgClassNameActive,
  getNetworkButtonBorderActive,
  getMenuItemStyleForChain,
} from '@/styles/chains'
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
import { HoverTokenAndBalance } from '@/components/Portfolio/components/HoverTokenAndBalance'

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

  let bgClassName

  if (isCurrentChain) {
    bgClassName = `
      ${getNetworkButtonBgClassName(chain.color)}
      ${getNetworkButtonBorder(chain.color)}
      bg-opacity-30
    `
  }

  return (
    <button
      ref={ref}
      tabIndex={active ? 1 : 0}
      className={`
        flex items-center justify-between
        transition-all duration-75
        w-full h-[62px]
        px-2 py-4
        cursor-pointer
        rounded-md
        border border-slate-400/10
        mb-1
        ${alternateBackground ? '' : !isCurrentChain && 'bg-slate-400/10'}
        ${bgClassName}
        ${getNetworkButtonBorderHover(chain.color)}
        ${getNetworkHover(chain.color)}
        ${getNetworkButtonBgClassNameActive(chain.color)}
        ${getNetworkButtonBorderActive(chain.color)}
        ${getMenuItemStyleForChain(chain.color)}
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
          className="ml-2 rounded-full w-7 h-7"
        />
        <div className="flex-col text-left">
          <div className="text-lg font-normal text-white">{chain.name}</div>
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
            className="w-6 h-6 rounded-md"
            alt={`${balanceTokens[0].token.symbol} img`}
            src={balanceTokens[0].token.icon}
            onMouseEnter={() => setIsT1Hovered(true)}
            onMouseLeave={() => setIsT1Hovered(false)}
          />
          <HoverTokenAndBalance
            isHovered={isT1Hovered}
            tokens={[balanceTokens[0]]}
            hoverClassName="bg-slate-900/70"
          />
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
            className="w-6 h-6 rounded-md"
            alt={`${balanceTokens[1].token.symbol} img`}
            src={balanceTokens[1].token.icon}
            onMouseEnter={() => setIsT2Hovered(true)}
            onMouseLeave={() => setIsT2Hovered(false)}
          />
          <HoverTokenAndBalance
            isHovered={isT2Hovered}
            tokens={[balanceTokens[1]]}
            hoverClassName="bg-slate-900/70"
          />
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
      <HoverTokenAndBalance
        isHovered={isT3Hovered}
        tokens={balanceTokens}
        startFrom={2}
        hoverClassName="bg-slate-900/70"
      />
    </div>
  )
}