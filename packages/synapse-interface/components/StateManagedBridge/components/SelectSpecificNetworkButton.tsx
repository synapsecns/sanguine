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
  TokenWithBalanceAndAllowances,
  sortTokensByBalanceDescending,
} from '@/utils/actions/fetchPortfolioBalances'

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
        w-full
        px-2 py-4
        cursor-pointer
        border-[1px] border-[#423F44]
        mb-1
        ${alternateBackground && 'bg-[#282328]'}
        ${bgClassName}
        ${getNetworkButtonBorderHover(chain.color)}
        ${getNetworkHover(chain.color)}
        ${getNetworkButtonBgClassNameActive(chain.color)}
        ${getNetworkButtonBorderActive(chain.color)}
      `}
      onClick={onClick}
      data-test-id={`${dataId}-item`}
    >
      <ButtonContent chainId={itemChainId} />
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
  const { balancesAndAllowances } = usePortfolioState()

  const balanceTokens =
    balancesAndAllowances &&
    balancesAndAllowances[chainId] &&
    sortTokensByBalanceDescending(
      balancesAndAllowances[chainId].filter((bt) => bt.balance > 0n)
    )

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
        </div>
      </div>
      {balanceTokens && balanceTokens.length > 0 ? (
        <ChainTokens balanceTokens={balanceTokens} />
      ) : null}
    </>
  ) : null
}

const ChainTokens = ({
  balanceTokens = [],
}: {
  balanceTokens: TokenWithBalanceAndAllowances[]
}) => {
  const [isHovered, setIsHovered] = useState(false)
  const hasOneToken = useMemo(
    () => balanceTokens && balanceTokens.length > 0,
    [balanceTokens]
  )
  const hasTwoTokens = useMemo(
    () => balanceTokens && balanceTokens.length > 1,
    [balanceTokens]
  )
  const numOverTwoTokens = useMemo(
    () =>
      balanceTokens && balanceTokens.length - 2 > 0
        ? balanceTokens.length - 2
        : 0,
    [balanceTokens]
  )

  return (
    <div
      data-test-id="chain-tokens"
      className="flex flex-row items-center cursor-pointer hover-trigger"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
    >
      {hasOneToken && (
        <Image
          loading="lazy"
          className="w-5 h-5 rounded-md"
          alt={`${balanceTokens[0].token.symbol} img`}
          src={balanceTokens[0].token.icon}
        />
      )}
      {hasTwoTokens && (
        <Image
          loading="lazy"
          className="w-5 h-5 ml-1.5 rounded-md"
          alt={`${balanceTokens[1].token.symbol} img`}
          src={balanceTokens[1].token.icon}
        />
      )}
      {numOverTwoTokens > 0 && (
        <div className="ml-1.5 text-white">+ {numOverTwoTokens}</div>
      )}
      <div className="relative inline-block">
        {isHovered && (
          <div
            className={`
              absolute -ml-28 z-50 hover-content p-2 text-white
              border border-solid border-[#252537]
              bg-[#101018] rounded-md
            `}
          >
            {balanceTokens.map(
              (token: TokenWithBalanceAndAllowances, key: number) => {
                const tokenSymbol = token.token.symbol
                const balance = token.parsedBalance
                return (
                  <div className="whitespace-nowrap" key={key}>
                    {balance} {tokenSymbol}
                  </div>
                )
              }
            )}
          </div>
        )}
      </div>
    </div>
  )
}
