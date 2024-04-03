import { useAccount } from 'wagmi'

import LoadingDots from '../ui/tailwind/LoadingDots'
import { ToChainSelector } from './ToChainSelector'
import { ToTokenSelector } from './ToTokenSelector'
import { useBridgeState, useBridgeDisplayState } from '@/slices/bridge/hooks'
import { DestinationAddressInput } from './DestinationAddressInput'

export const OutputContainer = ({}) => {
  const { address } = useAccount()
  const { bridgeQuote, isLoading } = useBridgeState()
  const { showDestinationAddress } = useBridgeDisplayState()

  return (
    <div className="relative text-left rounded-md p-md bg-bgLight">
      <div className="flex items-center justify-between mb-3">
        <ToChainSelector />

        {showDestinationAddress ? (
          <DestinationAddressInput connectedAddress={address} />
        ) : null}
      </div>

      <div className="flex h-16 mb-2 space-x-2">
        <div
          className={`
            flex flex-grow items-center pl-md w-full h-16
            rounded-md border border-white border-opacity-20
          `}
        >
          <ToTokenSelector />
          <div className="flex ml-4">
            {isLoading ? (
              <LoadingDots className="opacity-50" />
            ) : (
              <input
                pattern="[0-9.]+"
                disabled={true}
                className={`
                text-white text-opacity-80 text-xl font-medium
                  border-none p-0 bg-transparent
                  focus:outline-none focus:ring-0 focus:border-none
                  max-w-[190px] md:text-2xl placeholder:text-[#88818C]
                `}
                placeholder="0.0000"
                value={
                  bridgeQuote?.outputAmountString === '0'
                    ? ''
                    : bridgeQuote?.outputAmountString
                }
                name="inputRow"
                autoComplete="off"
              />
            )}
          </div>
        </div>
      </div>
    </div>
  )
}
