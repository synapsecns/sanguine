import { useEffect, useState } from 'react'
import { Address, useAccount } from 'wagmi'

import LoadingSpinner from '../ui/tailwind/LoadingSpinner'
import { ToChainSelector } from './ToChainSelector'
import { shortenAddress } from '@/utils/shortenAddress'
import { ToTokenSelector } from './ToTokenSelector'
import { useDispatch } from 'react-redux'
import { setToChainId, setToToken } from '@/slices/bridge/reducer'
import { useBridgeState } from '@/slices/bridge/hooks'

export const OutputContainer = ({}) => {
  const { bridgeQuote, isLoading, toChainId, toToken } = useBridgeState()

  const { address: isConnectedAddress } = useAccount()
  const [address, setAddress] = useState<Address>()

  const dispatch = useDispatch()

  useEffect(() => {
    setAddress(isConnectedAddress)
  }, [isConnectedAddress])

  // update address for destination address if we have a destination address

  return (
    <div
      className={`
        mt-6
        text-left px-2 sm:px-4 pt-4 pb-1 rounded-xl
        bg-bgLight
      `}
    >
      <div className="flex items-center justify-between mb-3">
        <ToChainSelector />
        {/* {address && (
          <div className="h-5">
            <DisplayAddress address={address} />
          </div>
        )} */}
      </div>

      <div className="flex h-16 mb-2 space-x-2">
        <div
          className={`
            flex flex-grow items-center
            pl-3 
            w-full h-16
            rounded-xl
            border border-white border-opacity-20
          `}
        >
          <ToTokenSelector />
          <div className="flex ml-4">
            {isLoading ? (
              <LoadingSpinner className="opacity-50" />
            ) : (
              <input
                pattern="[0-9.]+"
                disabled={true}
                className={`
                  focus:outline-none
                  focus:ring-0
                  focus:border-none
                  border-none
                  p-0
                  bg-transparent
                  max-w-[190px]
                placeholder:text-[#88818C]
                text-white text-opacity-80 text-lg md:text-2xl font-medium
                `}
                placeholder="0.0000"
                value={
                  bridgeQuote.outputAmountString === '0'
                    ? ''
                    : bridgeQuote.outputAmountString
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

const DisplayAddress = ({ address }) => {
  return (
    <div className="border-[0.5px] border-secondaryTextColor rounded-xl pt-1 pb-1 pl-3 pr-3 text-secondaryTextColor text-xxs">
      {shortenAddress(address, 3)}
    </div>
  )
}
