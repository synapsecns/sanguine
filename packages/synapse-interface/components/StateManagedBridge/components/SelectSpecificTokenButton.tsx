import { displaySymbol } from '@utils/displaySymbol'
import {
  getBorderStyleForCoinHover,
  getMenuItemBgForCoin,
  getMenuItemStyleForCoin,
  getMenuItemStyleForCoinCombined,
} from '@styles/tokens'
import { memo } from 'react'
import { Token } from '@/utils/types'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { useBridgeState } from '@/slices/bridge/hooks'

const SelectSpecificTokenButton = ({
  isOrigin,
  token,
  active,
  selectedToken,
  onClick,
}: {
  isOrigin: boolean
  token: Token
  active: boolean
  selectedToken: Token
  onClick: () => void
}) => {
  const isCurrentlySelected = selectedToken?.symbol === token?.symbol
  const { fromChainId, toChainId } = useBridgeState()

  const chainId = isOrigin ? fromChainId : toChainId

  let bgClassName

  const classNameForBorderStyle = getBorderStyleForCoinHover(token?.color)
  const classNameForMenuItemStyle = getMenuItemStyleForCoinCombined(
    token?.color
  )

  if (isCurrentlySelected) {
    bgClassName = getMenuItemBgForCoin(token?.color)
  }

  return (
    <div
      tabIndex={active ? 1 : 0}
      onClick={onClick}
      className={`
        flex items-center
        transition-all duration-75
        w-full
        px-2 py-1
        cursor-pointer
        border-[1px] border-[#423F44]
        mb-1
        ${bgClassName}
        ${classNameForBorderStyle}
        ${classNameForMenuItemStyle}
      `}
    >
      <ButtonContent token={token} chainId={chainId} />
    </div>
  )
}

const ButtonContent = memo(
  ({ token, chainId }: { token: Token; chainId: number }) => {
    const portfolioBalances = usePortfolioBalances()

    const parsedBalance = portfolioBalances[chainId]?.find(
      (tb) => tb.token === token
    )?.parsedBalance

    return (
      <div className="flex items-center w-full">
        <img
          alt="token image"
          className="w-8 h-8 ml-2 mr-4 rounded-full"
          src={token?.icon?.src}
        />
        <Coin token={token} />
        <TokenBalance
          token={token}
          chainId={chainId}
          parsedBalance={parsedBalance}
        />
      </div>
    )
  }
)

const Coin = ({ token }: { token }) => {
  return (
    <div className="flex-col text-left">
      <div className="text-lg text-primaryTextColor">{token?.symbol}</div>
      <div className="flex items-center text-xs text-secondaryTextColor">
        {token?.name}
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
            {token ? displaySymbol(chainId, token) : ''}
          </span>
        </div>
      )}
    </div>
  )
}

export default SelectSpecificTokenButton
