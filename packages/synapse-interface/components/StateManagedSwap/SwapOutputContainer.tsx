import { useEffect, useState } from 'react'
import { Address, useAccount } from 'wagmi'

import LoadingDots from '../ui/tailwind/LoadingDots'
import { SwapToTokenSelector } from './SwapToTokenSelector'
import { useSwapState } from '@/slices/swap/hooks'
import { BridgeCardTokenInput, BridgeContainer } from '../ui/BridgeCard'

export const SwapOutputContainer = ({}) => {
  const { swapQuote, isLoading, swapToToken } = useSwapState()

  const { address: isConnectedAddress } = useAccount()
  const [address, setAddress] = useState<Address>()

  useEffect(() => {
    setAddress(isConnectedAddress)
  }, [isConnectedAddress])

  return (
    <BridgeContainer>
      <BridgeCardTokenInput>
        <SwapToTokenSelector />
        {isLoading ? (
          <LoadingDots className="opacity-50" />
        ) : (
          <input
            pattern="[0-9.]+"
            disabled={true}
            className={`
              focus:outline-none focus:ring-0 focus:border-none
              border-none
              p-0 w-full
              bg-transparent
              placeholder:text-[#88818C]
              text-white text-opacity-80 text-xl md:text-2xl font-medium
            `}
            placeholder="0.0000"
            value={
              swapQuote.outputAmountString === '0'
                ? ''
                : swapQuote.outputAmountString
            }
            name="inputRow"
            autoComplete="off"
          />
        )}
      </BridgeCardTokenInput>
    </BridgeContainer>
  )
}
