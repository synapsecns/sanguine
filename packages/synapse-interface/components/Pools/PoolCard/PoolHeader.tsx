import _ from 'lodash'
import { memo, useEffect, useMemo, useState } from 'react'
import Link from 'next/link'
import { Address, useNetwork } from 'wagmi'

import { getPoolUrl } from '@urls'
import { CHAINS_BY_ID } from '@/constants/chains'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { RightArrow } from '@/components/icons/RightArrow'
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
      const balancesForChain = _.pickBy( balances[pool.chainId],
        (value, _key) => value.balance > 0n
      )

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
      <div className="flex items-center justify-between p-3 pt-2 pb-2 pl-3 border border-transparent ">
        <div className="flex items-center">
          <img
            src={chain.chainImg.src}
            className="w-[16px] h-[16px] rounded-full mr-2"
          />
          <div className="mr-2 text-white text-md">{chain.name}</div>
          <div className="text-sm text-[#BFBCC2] mr-4">{pool.symbol}</div>
          {mounted && connectedChain && connectedChain.id === pool.chainId && (
            <ConnectedIndicator />
          )}
        </div>

        {canDeposit && pool.incentivized ? (
          <Link href={getPoolUrl(pool)}>
            <div className="text-sm text-[#99E6FF] flex items-center space-x-1">
              <div className="hover:underline pr-1">Deposit</div>
              <RightArrow />
            </div>
          </Link>
        ) : (
          <div className="text-bgBase/50">
            <RightArrow className="stroke-bgBase/50 group-hover:stroke-bgBase" />
          </div>
        )}
      </div>
    )
  }
)

const ConnectedIndicator = () => {
  return (
    <div className="flex flex-row space-x-1 text-sm">
      <div className="w-2 h-2 my-auto ml-auto bg-green-500 rounded-full " />
      <div className="text-xs text-secondaryTextColor">Connected</div>
    </div>
  )
}
