import { useEffect, useState } from 'react'
import { useDispatch } from 'react-redux'
import { useAccount } from 'wagmi'
import { switchNetwork } from '@wagmi/core'
import { useConnectModal } from '@rainbow-me/rainbowkit'
import { ConnectButton } from '@rainbow-me/rainbowkit'

import { setFromChainId } from '@/slices/bridge/reducer'
import { useBridgeState } from '@/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import {
  getNetworkButtonBgClassNameActive,
  getNetworkButtonBorderActive,
  getNetworkButtonBorderHover,
  getNetworkHover,
} from '@/styles/chains'
import { LoaderIcon } from 'react-hot-toast'

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

const DisconnectedIndicator = () => {
  const { openConnectModal } = useConnectModal()
  const { fromChainId } = useBridgeState()
  const chain = CHAINS_BY_ID[fromChainId]

  return (
    <button
      data-test-id="disconnected-button"
      className={`
        flex items-center justify-center
        text-base text-white px-3 py-1 rounded-md
        text-center transform-gpu transition-all duration-75
        border border-solid border-transparent
        h-8
        ${getNetworkHover(chain?.color)}
        ${getNetworkButtonBgClassNameActive(chain?.color)}
        ${getNetworkButtonBorderActive(chain?.color)}
        ${getNetworkButtonBorderHover(chain?.color)}
      `}
      onClick={openConnectModal}
    >
      <div className="flex flex-row text-sm">
        <div
          className={`
            my-auto ml-auto mr-2 w-2 h-2
            bg-red-500 rounded-full
            `}
        />
        Disconnected
      </div>
    </button>
  )
}

export const ConnectToNetworkButton = ({ chainId }: { chainId: number }) => {
  const [isConnecting, setIsConnecting] = useState<boolean>(false)
  const dispatch = useDispatch()
  const chain = CHAINS_BY_ID[chainId]

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
        text-base text-white px-3 py-1 rounded-lg
        text-center transform-gpu transition-all duration-75
        border border-solid border-transparent
        h-8
        ${getNetworkHover(chain?.color)}
        ${getNetworkButtonBgClassNameActive(chain?.color)}
        ${getNetworkButtonBorderActive(chain?.color)}
        ${getNetworkButtonBorderHover(chain?.color)}
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
        <div className="flex flex-row text-sm">
          <div
            className={`
              my-auto ml-auto mr-2 text-transparent w-2 h-2
              border border-indigo-300 border-solid rounded-full
            `}
          />
          Switch Network
        </div>
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
          {({ account, chain, openConnectModal, mounted, openChainModal }) => {
            return (
              <>
                {(() => {
                  if (!mounted || !account || !chain || !address) {
                    return (
                      <button
                        className={`
                          flex items-center text-sm text-white mr-2
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
