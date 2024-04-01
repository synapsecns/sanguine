import React from 'react'
import { Token } from '@/utils/types'
import Image from 'next/image'
import { getParsedBalance } from './PortfolioTokenAsset'
import GasIcon from '@/components/icons/GasIcon'

export const GasTokenAsset = ({
  token,
  balance,
}: {
  token: Token
  balance: bigint
}) => {
  const { icon, symbol, decimals } = token
  const parsedBalance = getParsedBalance(balance, decimals as number, 3)

  return (
    <div
      id="gas-token-asset"
      className={`
        p-2 flex items-center border-y text-white
        justify-between last:rounded-b-md border-transparent
      `}
    >
      <div className="flex items-center gap-2 py-2 pl-2 pr-4 rounded">
        <Image
          loading="lazy"
          alt={`${symbol} img`}
          className="w-6 h-6 rounded-md"
          src={icon}
        />
        {parsedBalance} {symbol}
        <GasIcon className="w-3 pt-px m-auto fill-secondary" />
      </div>

      <div className="p-2 text-sm opacity-70">Not bridgeable</div>
    </div>
  )
}
