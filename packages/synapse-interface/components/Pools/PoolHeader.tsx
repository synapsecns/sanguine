import _ from 'lodash'
import { memo, useEffect, useMemo, useState } from 'react'
import Link from 'next/link'
import { Address, useNetwork } from 'wagmi'

import { getPoolUrl } from '@urls'
import { CHAINS_BY_ID } from '@/constants/chains'
import { usePortfolioState } from '@/slices/portfolio/hooks'
// import { RightArrow } from '@/components/icons/RightArrow'
import { Token } from '@/utils/types'

export const PoolHeader = memo(
  ({ pool, address }: { pool: Token; address: Address }) => {
    const [mounted, setMounted] = useState(false)
    const { chain: connectedChain } = useNetwork()
    const chain = CHAINS_BY_ID[pool.chainId]
    const { balances } = usePortfolioState()

    useEffect(() => {
      setMounted(true)
    }, [])
    const canDeposit = useMemo(() => {
      const balancesForChain = _(balances[pool.chainId])
        .pickBy((value, _key) => value.balance > 0n)
        .value()

      if (Object.keys(balancesForChain).length === 0) return false

      if (!address) {
        return null
      }

      return _.some(pool.nativeTokens, (poolToken) => {
        const poolAddress = _.get(poolToken, `addresses.${pool.chainId}`)
        return _.some(balancesForChain, (balance) => {
          return (
            poolAddress === _.get(balance, `token.addresses.${pool.chainId}`)
          )
        })
      })
    }, [pool, balances, address])
    return (
      <div className="flex items-center gap-1 px-3 py-2 justify-between">
        <div className="flex items-center gap-1">
          <img
            src={chain.chainImg.src}
            className="w-4 h-4"
          />
          {chain.name}
          <span className="text-sm text-zinc-400 mt-0.5 mr-1.5">{pool.symbol}</span>
          {mounted && connectedChain && connectedChain.id === pool.chainId && (
            <ConnectedIndicator />
          )}
        </div>
        {canDeposit && pool.incentivized ? (
          <Link href={getPoolUrl(pool)}>
            <div className="text-sm text-sky-500 flex items-center space-x-1">
              <div className="hover:underline">Deposit</div>
              <RightArrow className="stroke-sky-500" />
            </div>
          </Link>
        ) : (
          <RightArrow />
        )}
      </div>
    )
  }
)

const ConnectedIndicator = () => {
  return (
    <div className="flex items-center gap-[5px] text-sm text-zinc-400">
      <span className="w-1.5 h-1.5 bg-green-500 rounded-full " />
      Connected
    </div>
  )
}

const RightArrow = ({className = 'stroke-zinc-400'}) => {
  return (
    <svg
      width="7"
      height="9"
      viewBox="0 0 7 12"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      className={`justify-self-end ${className}`}
    >
      <path
        d="M1 1L6 6L1 11"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
    </svg>
  )
}
