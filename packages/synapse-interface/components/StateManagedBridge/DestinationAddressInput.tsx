import type { Address } from 'wagmi'
import { useDispatch } from 'react-redux'

import { setDestinationAddress } from '@/slices/bridge/reducer'
import { CHAINS_BY_ID } from '@constants/chains'

export const DestinationAddressInput = ({
  toChainId,
  destinationAddress,
}: {
  toChainId: number
  destinationAddress: string
}) => {
  const dispatch = useDispatch()
  const chain = CHAINS_BY_ID[toChainId]
  const chainName = chain?.name || ''
  const placeholder = `Enter ${chainName} address...`

  return (
    <div>
      <div className="w-[30%]">
        <div className="flex items-center justify-center  h-[26px] -mt-4 p-2 absolute ml-5 md:ml-10 text-sm text-[#D8D1DC] rounded-md bg-slate-700">
          Withdraw to...
        </div>
      </div>
      <div className="h-16 pb-4 mt-4 text-left sm:px-2">
        <div
          className={`
            h-14 flex flex-grow items-center
            bg-bgBase/10
            border border-white/10 hover:border-white/50 focus-within:border-white/50
            pl-3 sm:pl-3
            py-0.5 rounded-md
          `}
        >
          <input
            className={`
              focus:outline-none
              focus:ring-0
              focus:border-none
              border-none
              bg-transparent
              w-[300px]
              sm:min-w-[400px]
              max-w-[calc(100%-88px)]
              sm:w-full
              text-white text-opacity-80
              placeholder:text-slate-400
            `}
            placeholder={placeholder}
            onChange={(e) => {
              dispatch(setDestinationAddress(e.target.value as Address))
            }}
            value={destinationAddress}
          />
        </div>
      </div>
    </div>
  )
}
