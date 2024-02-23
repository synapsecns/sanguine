import _ from 'lodash'
import { useState } from 'react'

import type { Token } from '@/utils/types'
import { CHAINS_BY_ID } from '@/constants/chains'

export const AvailableChains = ({
    token,
    pausedChainIds=[]
}: {
  token: Token
  pausedChainIds?: string[]
}) => {
  const [isHovered, setIsHovered] = useState(false)
  const chainIds = _.difference(Object.keys(token.addresses), pausedChainIds)
  const hasOneChain = chainIds.length > 0
  const hasMultipleChains = chainIds.length > 1
  const numOverTwoChains = chainIds.length - 2 > 0 ? chainIds.length - 2 : 0

  return (
    <div
      data-test-id="available-chains"
      className="flex flex-row items-center space-x-1 hover-trigger"
      onMouseEnter={() => setIsHovered(true)}
      onMouseLeave={() => setIsHovered(false)}
    >
      {hasOneChain && (
        <img
          className="w-3 h-3 rounded-md"
          alt={`${CHAINS_BY_ID[chainIds[0]].name} img`}
          src={`${CHAINS_BY_ID[chainIds[0]].chainImg.src}`}
        />
      )}
      {hasMultipleChains && (
        <img
          className="w-3 h-3 rounded-md"
          alt={`${CHAINS_BY_ID[chainIds[1]].name} img`}
          src={`${CHAINS_BY_ID[chainIds[1]].chainImg.src}`}
        />
      )}
      {numOverTwoChains > 0 && (
        <div className="ml-1 text-white">+ {numOverTwoChains}</div>
      )}
      <div className="relative inline-block">
        {isHovered && (
          <div
            className={`
              absolute z-50 hover-content p-2 text-white
              border border-solid border-[#252537]
              bg-[#101018] rounded-md
            `}
          >
            {chainIds.map((chainId) => {
              const chainName = CHAINS_BY_ID[chainId].name
              return <div className="whitespace-nowrap">{chainName}</div>
            })}
          </div>
        )}
      </div>
    </div>
  )
}


