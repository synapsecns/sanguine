import { CHAINS_BY_ID } from '@constants/chains'
import Image from 'next/image'
import {
  getBorderStyleForCoinHover,
  getMenuItemStyleForCoinCombined,
} from '@styles/tokens'
import { memo } from 'react'
import { Token } from '@/utils/types'
import TokenBalance from '@components/TokenBalance'
const TokenMenuItem = ({
  token,
  active,
  chainId,
  selectedToken,
  tokenBalance,
  onClick,
}: {
  token: Token
  active: boolean
  chainId: number
  selectedToken: Token
  tokenBalance: bigint | undefined
  onClick: () => void
}) => {
  const isCurrentlySelected = selectedToken?.symbol === token?.symbol

  let bgClassName

  if (isCurrentlySelected) {
    bgClassName = `bg-bgLight hover:bg-bgLight active:bg-bgLight`
  } else {
    bgClassName = `bg-[#58535B] hover:bg-[#58535B] active:bg-[#58535B]`
  }

  const classNameForBorderStyle = getBorderStyleForCoinHover(token?.color)
  const classNameForMenuItemStyle = getMenuItemStyleForCoinCombined(
    token?.color
  )

  return (
    <div
      tabIndex={active ? 1 : 0}
      onClick={onClick}
      className={`
      flex items-center
      transition-all duration-75
      w-full rounded-md
      px-2 py-3
      cursor-pointer
      border border-transparent
      ${bgClassName}
      ${classNameForBorderStyle}
      ${classNameForMenuItemStyle}
        `}
    >
      <ButtonContent
        token={token}
        chainId={chainId}
        tokenBalance={tokenBalance ? tokenBalance : 0n}
      />
    </div>
  )
}

const ButtonContent = memo(
  ({
    token,
    chainId,
    tokenBalance,
  }: {
    token: Token
    chainId: number
    tokenBalance: bigint
  }) => {
    return (
      <div className="flex items-center w-full">
        <img
          alt="token image"
          className="w-10 h-10 ml-2 mr-4 rounded-full"
          src={token?.icon?.src}
        />
        <CoinOnChain token={token} chainId={chainId} />
        <TokenBalance
          token={token}
          chainId={chainId}
          tokenBalance={tokenBalance}
        />
      </div>
    )
  }
)

const CoinOnChain = ({ token, chainId }: { token: Token; chainId: number }) => {
  const chain = CHAINS_BY_ID?.[chainId]

  return chain ? (
    <div className="flex-col text-left">
      <div className="text-lg font-medium text-white">
        {token ? token.symbol : ''}
      </div>
      <div className="flex items-center text-sm text-white">
        <div className="mr-1 opacity-70">{token?.name}</div>
        <div className="opacity-60">on</div>
        <Image
          src={chain?.chainImg}
          alt={chain?.name}
          className="w-4 h-4 ml-2 mr-2 rounded-full"
        />
        <div className="hidden md:inline-block opacity-70">{chain?.name}</div>
      </div>
    </div>
  ) : null
}

export default TokenMenuItem
