import { useRef, useEffect } from 'react'

import { commify } from '@ethersproject/units'
import { formatBnMagic } from '@bignumber/format'

import { CHAIN_INFO_MAP } from '@constants/networks'

import { useGenericTokenBalance } from '@hooks/tokens/useTokenBalances'
import { getTokenOnChain } from '@hooks/tokens/useTokenInfo'
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
  onClick,
}) {
  const ref = useRef(null)

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
      tabIndex={active ? '1' : '0'}
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
      <ButtonContent coin={coin} chainId={chainId} />
    </div>
  )
}

function ButtonContent({ coin, chainId }) {
  return (
    <div className="flex items-center w-full">
      <img className="w-10 h-10 ml-2 mr-4 rounded-full" src={coin.icon} />
      <CoinOnChain coin={coin} chainId={chainId} />
      <TokenBalance coin={coin} chainId={chainId} />
    </div>
  )
}

function CoinOnChain({ coin, chainId }) {
  const { chainImg, chainName } = CHAIN_INFO_MAP[chainId]

  return (
    <div className="flex-col text-left">
      <div className="text-lg font-medium text-white">
        {displaySymbol(chainId, coin)}
      </div>
      <div className="flex items-center text-sm text-white">
        <div className="mr-1 opacity-70">{coin.name}</div>
        <div className="opacity-60">on</div>
        <img
          src={chainImg}
          alt={chainName}
          className="w-4 h-4 ml-2 mr-2 rounded-full"
        />
        <div className="hidden md:inline-block opacity-70">{chainName}</div>
      </div>
    </div>
  )
}

function TokenBalance({ coin, chainId }) {
  const tokenBalance = useGenericTokenBalance(chainId, coin)
  const tokenInfo = getTokenOnChain(chainId, coin)

  const formattedBalance = commify(formatBnMagic(tokenBalance, tokenInfo, 2))

  return (
    <div className="ml-auto mr-5 text-lg text-white">
      {!tokenBalance.eq(0) && (
        <p>
          {formattedBalance}
          <span className="text-sm opacity-80">
            {' '}
            {displaySymbol(chainId, coin)}
          </span>
        </p>
      )}
    </div>
  )
}
