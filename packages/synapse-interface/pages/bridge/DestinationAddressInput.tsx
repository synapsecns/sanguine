import { ChainId, CHAIN_INFO_MAP } from '@constants/networks'
import { validateAndParseAddress } from '@utils/validateAndParseAddress'
import { useNetworkController } from '@hooks/wallet/useNetworkController'

import Button from '@tw/Button'

export function DestinationAddressInput({
  nonEvmBridge,
  fromChainId,
  toChainId,
  destinationAddress,
  setDestinationAddress,
}) {
  const { terraAddress, account } = useNetworkController()
  const { chainName } = CHAIN_INFO_MAP[toChainId]

  let placeholder
  if (fromChainId == ChainId.TERRA) {
    if (account) {
      placeholder = `Default: ${account.slice(0, 6)}...${account.slice(
        -4,
        account.length
      )}`
    } else {
      placeholder = 'Connect Wallet to Autofill'
    }
  } else if (toChainId == ChainId.TERRA) {
    if (terraAddress) {
      placeholder = `Terra address: ${terraAddress.slice(
        0,
        12
      )}...${terraAddress.slice(-6, terraAddress.length)}`
    } else {
      placeholder = 'Terra Destination (connect to autofill)'
    }
  } else {
    placeholder = `Enter ${chainName} address...`
  }

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
          {nonEvmBridge && <ConnectChainButton toChainId={toChainId} />}
        </div>
      </div>
    </div>
  )
}

function ConnectChainButton({ toChainId }) {
  const { terraAddress, account, connectToChain, disconnectChain } =
    useNetworkController()

  let label
  let onClick
  let isConnected
  if (toChainId == ChainId.TERRA) {
    if (terraAddress) {
      isConnected = true
      label = 'Disconnect'
      onClick = () => {
        disconnectChain(toChainId)
      }
    } else {
      isConnected = false
      label = 'Connect'
      onClick = () => {
        connectToChain(toChainId)
      }
    }
  } else {
    if (account) {
      isConnected = true
      label = 'Disconnect'
      onClick = () => {
        disconnectChain(toChainId)
      }
    } else {
      isConnected = false
      label = 'Connect'
      onClick = () => {
        connectToChain(toChainId)
      }
    }
  }

  let btnClassName
  if (isConnected) {
    if (toChainId == ChainId.TERRA) {
      btnClassName = `
        hover:text-slate-300
        hover:!border-blue-500
        `
    } else {
      btnClassName = `
        hover:text-slate-300
        `
    }
  } else {
    if (toChainId == ChainId.TERRA) {
      btnClassName = `
        bg-gradient-to-r from-[#CF52FE] to-[#AC8FFF]
        hover:opacity-80
        active:from-[#622e71] active:to-[#564071]
        `
    } else {
      btnClassName = ''
    }
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
