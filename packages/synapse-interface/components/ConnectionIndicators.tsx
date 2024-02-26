import { useEffect, useState } from 'react'
import { useDispatch } from 'react-redux'
import { useAccount } from 'wagmi'
import { switchNetwork } from '@wagmi/core'

import { ConnectButton } from '@rainbow-me/rainbowkit'

import { setFromChainId } from '@/slices/bridge/reducer'

import { CHAINS_BY_ID } from '@/constants/chains'
import {
  getNetworkButtonBgClassNameActive,
  getNetworkButtonBorderActive,
  getNetworkButtonBorderHover,
  getNetworkHover,
} from '@/styles/chains'
import { LoaderIcon } from 'react-hot-toast'
import { useHasMounted } from '@/utils/hooks/useHasMounted'

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
            my-auto ml-auto mr-2 size-2
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
              my-auto ml-auto mr-2 text-transparent size-2
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
              my-auto ml-auto mr-2 text-transparent size-2
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
  const clientReady = useHasMounted()
  const { address } = useAccount()

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
                          flex items-right justify-center
                          text-base text-white px-3 py-1 rounded-lg
                          text-center transform-gpu transition-all duration-75
                          border border-solid border-transparent
                          hover:border-[#3D3D5C]
                        `}
                        onClick={openConnectModal}
                      >
                        <div className="flex flex-row text-sm">
                          <div
                            className={`
                              my-auto ml-auto mr-2 text-transparent size-2
                              border border-indigo-300 border-solid rounded-full
                            `}
                          />
                          Connect Wallet
                        </div>

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
