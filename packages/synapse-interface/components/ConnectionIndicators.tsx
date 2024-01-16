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

const buttonBaseStyles = `
  flex gap-2 items-center
  px-3 py-1 rounded-full text-sm
  transition-all duration-75
  border border-transparent
`
const dotBaseStyles = "w-2 h-2 rounded-full"

export const ConnectedIndicator = () => {
  return (
    <button
      data-test-id="connected-button"
      className={buttonBaseStyles}
    >
      <div className="w-2 h-2 bg-green-500 rounded-full" />
      Connected
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
        ${buttonBaseStyles}
        ${getNetworkHover(chain?.color)}
        ${getNetworkButtonBgClassNameActive(chain?.color)}
        ${getNetworkButtonBorderActive(chain?.color)}
        ${getNetworkButtonBorderHover(chain?.color)}
      `}
      onClick={openConnectModal}
    >
        <div className={`${dotBaseStyles} bg-red-500`} />
        Disconnected
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
        ${buttonBaseStyles}
        ${getNetworkHover(chain?.color)}
        ${getNetworkButtonBgClassNameActive(chain?.color)}
        ${getNetworkButtonBorderActive(chain?.color)}
        ${getNetworkButtonBorderHover(chain?.color)}
      `}
      onClick={handleConnectNetwork}
    >
      {isConnecting ? (
        <>
          <div className={`${dotBaseStyles} border border-green-300`} />
          Connecting
          <LoaderIcon />
        </>
      ) : (
        <>
          <div className={`${dotBaseStyles} border border-indigo-300`} />
          Switch Network
        </>
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
                        className={buttonBaseStyles}
                        onClick={openConnectModal}
                      >
                        <div className={`${dotBaseStyles} border border-indigo-300`} />
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
