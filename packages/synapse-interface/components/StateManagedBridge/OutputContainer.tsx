import { useEffect, useState } from 'react'
import { Address, useAccount } from 'wagmi'

import LoadingDots from '../ui/tailwind/LoadingDots'
import { ToChainSelector } from './ToChainSelector'
import { shortenAddress } from '@/utils/shortenAddress'
import { ToTokenSelector } from './ToTokenSelector'
import { useDispatch } from 'react-redux'
import { setToChainId, setToToken } from '@/slices/bridge/reducer'
import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeCardTokenInput, BridgeContainer } from '../ui/BridgeCard'

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
    <BridgeContainer>
      <div className="flex items-center justify-between">
        <ToChainSelector />
        {/* {address && (
          <div className="h-5">
            <DisplayAddress address={address} />
          </div>
        )} */}
      </div>

      <BridgeCardTokenInput>
        <ToTokenSelector />
        {isLoading ? (
          <LoadingDots className="opacity-50" />
        ) : (
          <input
            pattern="[0-9.]+"
            disabled={true}
            className={`
              focus:outline-none focus:ring-0 focus:border-none
              border-none
              p-0 flex-1
              bg-transparent
              placeholder:text-[#88818C]
              text-white text-opacity-80 text-xl md:text-2xl font-medium
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
      </BridgeCardTokenInput>
    </BridgeContainer>
  )
}

const DisplayAddress = ({ address }) => {
  return (
    <div className="border-[0.5px] border-secondaryTextColor rounded-md pt-1 pb-1 pl-3 pr-3 text-secondaryTextColor text-xxs">
      {shortenAddress(address)}
    </div>
  )
}
