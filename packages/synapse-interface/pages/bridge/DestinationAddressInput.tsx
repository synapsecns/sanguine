import { CHAIN_INFO_MAP } from '@constants/networks'
import { useAccount, useNetwork, useSwitchNetwork } from 'wagmi'

// import { useNetworkController } from '@hooks/wallet/useNetworkController'

import Button from '@tw/Button'

export function DestinationAddressInput({
  toChainId,
  destinationAddress,
  setDestinationAddress,
}: {
  toChainId: number
  destinationAddress: string
  setDestinationAddress: (val: string) => void
}) {
  const { chainName } = CHAIN_INFO_MAP[toChainId]

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
          py-0.5 rounded-xl
        `}
        >
          <input
            className={`
              focus:outline-none
              bg-transparent
              w-[300px]
              sm:min-w-[300px]
              max-w-[calc(100%-92px)]
              sm:w-full
              text-white text-opacity-80 text-xl
              placeholder:text-[#88818C]
            `}
            placeholder={placeholder}
            onChange={(e) => {
              setDestinationAddress(e.target.value)
            }}
            value={destinationAddress}
          />
        </div>
      </div>
    </div>
  )
}

function ConnectChainButton({ toChainId }: { toChainId: number }) {
  // const { terraAddress, account, connectToChain, disconnectChain } =
  //   useNetworkController()
  const { address } = useAccount()

  const { chain } = useNetwork()
  const { chains, error, isLoading, pendingChainId, switchNetwork } =
    useSwitchNetwork()
  let label
  let onClick
  let isConnected

  if (address) {
    isConnected = true
    label = 'Disconnect'
    onClick = () => {
      switchNetwork?.(toChainId)
    }
  } else {
    isConnected = false
    label = 'Connect'
    onClick = () => {
      switchNetwork?.(toChainId)
    }
  }

  let btnClassName
  if (isConnected) {
    btnClassName = `
        hover:text-slate-300
        `
  } else {
    btnClassName = ''
  }

  return (
    <Button
      className={`
        flex-shrink
        rounded-lg
        !py-0.5
        my-1
        -mr-2
        text-xs
        ${btnClassName}
      `}
      onClick={onClick}
    >
      {label}
    </Button>
  )
}
