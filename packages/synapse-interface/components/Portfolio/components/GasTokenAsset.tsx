import React from 'react'
import Image from 'next/image'
import { Token } from '@/utils/types'
import { getParsedBalance } from '@/utils/getParsedBalance'
import { HoverTooltip } from '@/components/HoverTooltip'
import { formatAmount } from '@/utils/formatAmount'
import GasIcon from '@/components/icons/GasIcon'

export const GasTokenAsset = ({
  token,
  balance,
}: {
  token: Token
  balance: bigint
}) => {
  const { icon, symbol, decimals } = token
  const parsedBalance = getParsedBalance(balance, decimals as number)

  return (
    <div
      id="gas-token-asset"
      className={`
        p-2 flex items-center border-y text-white
        justify-between last:rounded-b-md border-transparent
      `}
    >
      <div className="relative flex items-center gap-2 py-2 pl-2 pr-4 rounded">
        <Image
          loading="lazy"
          alt={`${symbol} img`}
          className="w-6 h-6 rounded-md"
          src={icon}
        />
        <HoverTooltip
          hoverContent={
            <div className="whitespace-nowrap">
              {parsedBalance} {symbol}
            </div>
          }
        >
          <div>
            {formatAmount(parsedBalance)} {symbol}
          </div>
        </HoverTooltip>
        <HoverTooltip
          hoverContent={<div className="whitespace-nowrap">Gas token</div>}
        >
          <GasIcon className="pt-0.5 m-auto fill-secondary" />
        </HoverTooltip>
      </div>

      <div className="p-2 text-sm opacity-70">Not bridgeable</div>
    </div>
  )
}
