import { useEffect, useState } from 'react'
import { useDispatch } from 'react-redux'
import { useAccount } from 'wagmi'
import { switchNetwork } from '@wagmi/core'
import { ConnectButton } from '@rainbow-me/rainbowkit'

import { setFromChainId } from '@/slices/bridge/reducer'
import { CHAINS_BY_ID } from '@/constants/chains'
import { LoaderIcon } from 'react-hot-toast'
import { useBridgeValidations } from '@/utils/hooks/useBridgeValidations'

export const ConnectedIndicator = () => {
  return (
    <button
      data-test-id="connected-button"
      className={`
        flex items-center justify-center
        text-base text-white px-3 py-1 rounded-lg
        text-center transform-gpu transition-all duration-75
        border border-solid border-transparent
        hover:cursor-default
        h-8
      `}
    >
      <div className="flex flex-row text-sm">
        <div
          className={`
            my-auto ml-auto mr-2 w-2 h-2
            bg-green-500 rounded-full
            `}
        />
        Connected
      </div>
    </button>
  )
}

export const ConnectToNetworkButton = ({ chainId }: { chainId: number }) => {
  const [isConnecting, setIsConnecting] = useState<boolean>(false)
  const dispatch = useDispatch()
  const { hasInputAmount, hasValidRoute, onSelectedChain } =
    useBridgeValidations()

  function scrollToTop(): void {
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }

  const handleConnectNetwork: () => Promise<void> = async () => {
    setIsConnecting(true)
    try {
      await switchNetwork({ chainId: chainId }).then((success) => {
        success && dispatch(setFromChainId(chainId))
        scrollToTop()
      })
    } catch (error) {
      error && setIsConnecting(false)
    }
  }

  return (
    <button
      data-test-id="connect-button"
      className={`
        flex items-center justify-center
        text-base text-white px-3 py-1 rounded-sm
        text-center transform-gpu transition-all duration-75
        border border-solid 
        h-8
        ${
          !onSelectedChain && hasValidRoute
            ? 'border-fuchsia-500 bg-custom-gradient'
            : 'border-gray-500 text-opacity-50 hover:border-gray-400 hover:text-opacity-60'
        }
      `}
      onClick={handleConnectNetwork}
    >
      {isConnecting ? (
        <div className="flex flex-row text-sm">
          <div
            className={`
              my-auto ml-auto mr-2 text-transparent w-2 h-2
              border border-green-300 border-solid rounded-full
            `}
          />
          <div className="flex items-center space-x-2">
            <div>Connecting</div>
            <LoaderIcon />
          </div>
        </div>
      ) : (
        <div className="flex text-sm">Switch chain</div>
      )}
    </button>
  )
}

export function ConnectWalletButton() {
  const [clientReady, setClientReady] = useState<boolean>(false)
  const { address } = useAccount()

  useEffect(() => {
    setClientReady(true)
  }, [])

  return (
    <div data-test-id="">
      {clientReady && (
        <ConnectButton.Custom>
          {({ account, chain, openConnectModal, mounted }) => {
            return (
              <>
                {(() => {
                  if (!mounted || !account || !chain || !address) {
                    return (
                      <button
                        className={`
                          flex items-center justify-center
                          text-sm text-white px-3 py-1 rounded-sm
                          text-center transform-gpu transition-all duration-75
                          border border-solid border-fuchsia-600
                          bg-custom-gradient
                          h-8
                        `}
                        onClick={openConnectModal}
                      >
                        <div
                          className={`
                            my-auto ml-auto mr-2 text-transparent w-2 h-2
                            border border-indigo-300 border-solid rounded-full
                          `}
                        />
                        Connect Wallet
                      </button>
                    )
                  }
                })()}
              </>
            )
          }}
        </ConnectButton.Custom>
      )}
    </div>
  )
}
