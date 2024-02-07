import { setDestinationAddress } from '@/slices/bridge/reducer'
import { CHAINS_BY_ID } from '@constants/chains'
import { useDispatch } from 'react-redux'
import { Address } from 'wagmi'

export const DestinationAddressInput = ({
  toChainId,
  destinationAddress,
}: {
  toChainId: number
  destinationAddress: string
}) => {
  const dispatch = useDispatch()
  const chain = CHAINS_BY_ID[toChainId]
  const chainName = chain?.name || 'nulls'
  let placeholder
  placeholder = `Enter ${chainName} address...`

  return (
    <div>
      <div className="w-[30%]">
        <div className="flex items-center justify-center  h-[26px] -mt-4 p-2 absolute ml-5 md:ml-10 text-sm text-[#D8D1DC] rounded-md bg-bgLight">
          Withdraw to...
        </div>
      </div>
      <div className="h-16 px-2 pb-4 mt-4 space-x-2 text-left sm:px-5">
        <div
          className={`
            h-14 flex flex-grow items-center
            bg-transparent
            border border-bgLight hover:border-bgLightest focus-within:border-bgLightest
            pl-3 sm:pl-4
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
              max-w-[calc(100%-92px)]
              sm:w-full
              text-white text-opacity-80 text-lg
              placeholder:text-[#88818C]
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
