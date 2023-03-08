import { useRef, useEffect } from 'react'

import { commify } from '@ethersproject/units'
import { formatBnMagic } from '@bignumber/format'
import { Token } from '@utils/classes/Token'
import { BigNumber } from '@ethersproject/bignumber'
import { Zero } from '@ethersproject/constants'
import { CHAIN_INFO_MAP } from '@constants/networks'
import Image from 'next/image'
// import { useGenericTokenBalance } from '@hooks/tokens/useTokenBalances'
import { getTokenOnChain } from '@hooks/tokens/useTokenInfo'
import { useAccount, useBalance, useNetwork } from 'wagmi'

import { displaySymbol } from '@utils/displaySymbol'

import {
  getBorderStyleForCoinHover,
  getMenuItemStyleForCoinCombined,
} from '@styles/coins'

export default function TokenMenuItem({
  chainId,
  active,
  coin,
  selected,
  tokenBalance,
  onClick,
}: {
  chainId: number
  active: boolean
  coin: any
  selected: any
  tokenBalance: BigNumber
  onClick: () => void
}) {
  const ref = useRef<any>(null)

  const isCurrentlySelected = selected.symbol === coin.symbol

  useEffect(() => {
    if (active) {
      ref?.current?.focus()
    }
  }, [active])

  let bgClassName
  if (isCurrentlySelected) {
    bgClassName = `bg-bgLight hover:bg-bgLight active:bg-bgLight`
  } else {
    bgClassName = `bg-[#58535B] hover:bg-[#58535B] active:bg-[#58535B]`
  }

  return (
    <div
      ref={ref}
      tabIndex={active ? 1 : 0}
      onClick={onClick}
      className={`
        flex items-center
        transition-all duration-75
        w-full rounded-xl
        px-2 py-3
        cursor-pointer
        border border-transparent
        ${getBorderStyleForCoinHover(coin)}
        ${getMenuItemStyleForCoinCombined(coin)}
        ${bgClassName}
      `}
    >
      <ButtonContent
        token={coin}
        chainId={chainId}
        tokenBalance={tokenBalance}
      />
    </div>
  )
}

function ButtonContent({
  token,
  chainId,
  tokenBalance,
}: {
  token: Token
  chainId: number
  tokenBalance: BigNumber
}) {
  return (
    <div className="flex items-center w-full">
      <Image
        alt="token image"
        className="w-10 h-10 ml-2 mr-4 rounded-full"
        src={token.icon}
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

function CoinOnChain({ token, chainId }: { token: Token; chainId: number }) {
  const { chainImg, chainName } = CHAIN_INFO_MAP[chainId]

  return (
    <div className="flex-col text-left">
      <div className="text-lg font-medium text-white">
        {displaySymbol(chainId, token)}
      </div>
      <div className="flex items-center text-sm text-white">
        <div className="mr-1 opacity-70">{token.name}</div>
        <div className="opacity-60">on</div>
        <Image
          src={chainImg}
          alt={chainName}
          className="w-4 h-4 ml-2 mr-2 rounded-full"
        />
        <div className="hidden md:inline-block opacity-70">{chainName}</div>
      </div>
    </div>
  )
}

function TokenBalance({
  token,
  chainId,
  tokenBalance,
}: {
  token: Token
  chainId: number
  tokenBalance: BigNumber
}) {
  const tokenInfo = getTokenOnChain(chainId, token)
  const formattedBalance = commify(formatBnMagic(tokenBalance, tokenInfo, 2))
  return (
    <div className="ml-auto mr-5 text-lg text-white">
      {!tokenBalance.eq(0) && (
        <p>
          {formattedBalance}
          <span className="text-sm opacity-80">
            {' '}
            {displaySymbol(chainId, token)}
          </span>
        </p>
      )}
    </div>
  )
}
