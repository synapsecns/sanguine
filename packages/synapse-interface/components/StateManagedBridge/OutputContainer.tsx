import { useEffect, useState } from 'react'
import { Address, useAccount } from 'wagmi'

import LoadingDots from '../ui/tailwind/LoadingDots'
import { ChainSelector } from './ChainSelector'
import { ToChainSelector } from './ToChainSelector'
import { shortenAddress } from '@/utils/shortenAddress'
import { TokenSelector } from './TokenSelector'
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
    <div className="text-left rounded-md p-2 flex flex-col gap-2 bg-white dark:bg-zinc-700 border border-zinc-200 dark:border-transparent">
      <div className="flex items-center justify-between">
        <ChainSelector side="to" />
        {/* {address && (
          <div className="h-5">
            <DisplayAddress address={address} />
          </div>
        )} */}
      </div>
      <div className={`
        flex gap-2 items-center
        px-2 py-1 min-h-[4rem] rounded-md border
        border-zinc-200 dark:border-zinc-600
      `}>
        <TokenSelector side="to" />
        {isLoading ? (
          <LoadingDots className="opacity-50" />
        ) : (
          <input
            pattern="[0-9.]+"
            disabled={true}
            className={`
              focus:outline-none
              focus:ring-0
              focus:border-none
              border-none
              bg-transparent
              w-full
              p-0
              placeholder:text-zinc-400 placeholder:dark:text-zinc-500
              text-xl xs:text-2xl font-medium block
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
  )
}

const DisplayAddress = ({ address }) => {
  return (
    <div className="border border-secondaryTextColor rounded-md py-1 px-3 text-secondaryTextColor text-sm">
      {shortenAddress(address)}
    </div>
  )
}
