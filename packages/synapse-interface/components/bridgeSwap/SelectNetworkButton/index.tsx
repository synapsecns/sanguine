import { useEffect, useRef } from 'react'
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
import { sortTokensByBalanceDescending } from '@/utils/actions/fetchPortfolioBalances'
import { ELIGIBILITY_DEFAULT_TEXT } from '@/utils/hooks/useStipEligibility'

import { ChainTokens } from '@/components/Portfolio/components/ChainTokens'

export const SelectNetworkButton = ({
  itemChainId,
  isCurrentChain,
  active,
  onClick,
  dataId,
  isOrigin,
  isEligible=false,
  alternateBackground = false,
}: {
  itemChainId: number
  isCurrentChain: boolean
  active: boolean
  onClick: () => void
  dataId?: string
  isOrigin: boolean
  isEligible?: boolean
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
      <ButtonContent
        chainId={itemChainId}
        isOrigin={isOrigin}
        isEligible={isEligible}
      />
    </button>
  )
}

function ButtonContent({
  chainId,
  isOrigin,
  isEligible
}: {
  chainId: number
  isOrigin: boolean
  isEligible?: boolean
}) {
  const chain = CHAINS_BY_ID[chainId]
  const { balances } = usePortfolioState()

  const balanceTokens =
    balances &&
    balances[chainId] &&
    sortTokensByBalanceDescending(
      balances[chainId].filter((bt) => bt.balance > 0n)
    )

  return chain && (
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
      {isOrigin && balanceTokens && balanceTokens.length > 0 && (
        <ChainTokens
          balanceTokens={balanceTokens}
          hoverClassName="bg-slate-900/70"
        />
      )}
    </>
  )
}
