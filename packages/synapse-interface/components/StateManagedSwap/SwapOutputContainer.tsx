import { useEffect, useState } from 'react'
import { Address, useAccount } from 'wagmi'

import LoadingDots from '../ui/tailwind/LoadingDots'
// import { SwapToTokenSelector } from './SwapToTokenSelector'
import { useSwapState } from '@/slices/swap/hooks'
import {
  BridgeAmountContainer,
  BridgeSectionContainer,
  ChainSelector,
  TokenSelector,
} from '../ui/BridgeCardComponents'
import { CHAINS_BY_ID } from '@/constants/chains'
import { useBridgeState } from '@/slices/bridge/hooks'
import { SwapChainListArray } from './SwapChainListOverlay'
import {
  setSwapChainId,
  setSwapFromToken,
  setSwapToToken,
} from '@/slices/swap/reducer'
// import { SwapFromTokenListOverlay } from './SwapFromTokenListOverlay'
import { SwapToTokenListArray } from './SwapToTokenListOverlay'

export const SwapOutputContainer = ({}) => {
  const { swapQuote, isLoading, swapToToken } = useSwapState()

  const { address: isConnectedAddress } = useAccount()
  const [address, setAddress] = useState<Address>()

  useEffect(() => {
    setAddress(isConnectedAddress)
  }, [isConnectedAddress])

  const join = (a) => Object.values(a).join(' ')

  const inputClassName = join({
    unset: 'bg-transparent border-none p-0',
    layout: 'flex-1', // required for Swap Output – different from other inputs for some reason
    placeholder: 'placeholder:text-zinc-500 placeholder:dark:text-zinc-400',
    font: 'text-xl md:text-2xl font-medium',
    focus: 'focus:outline-none focus:ring-0 focus:border-none',
  })

  return (
    <BridgeSectionContainer>
      <BridgeAmountContainer>
        <SwapToTokenSelector />
        {isLoading ? (
          <LoadingDots className="opacity-50" />
        ) : (
          <input
            pattern="[0-9.]+"
            readOnly={true}
            disabled={true}
            className={inputClassName}
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
      </BridgeAmountContainer>
    </BridgeSectionContainer>
  )
}

const SwapToTokenSelector = () => (
  <TokenSelector
    dataTestId="swap-destination-token"
    isOrigin={false}
    selectedItem={useSwapState().swapToToken}
    placeholder="Out"
    itemListFunction={SwapToTokenListArray}
    setFunction={setSwapToToken}
  />
)