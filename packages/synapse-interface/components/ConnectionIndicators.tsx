import { useEffect, useMemo, useState } from 'react'
import { useDispatch } from 'react-redux'
import { useAccount } from 'wagmi'
import { switchNetwork } from '@wagmi/core'
import { ConnectButton } from '@rainbow-me/rainbowkit'

import { setFromChainId } from '@/slices/bridge/reducer'
import { useBridgeStatus } from '@/slices/bridge/hooks'
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

export const ConnectToNetworkButton = ({ chainId }: { chainId: number }) => {
  const [isConnecting, setIsConnecting] = useState<boolean>(false)
  const dispatch = useDispatch()

  const { hasInputAmount, hasValidSelections } = useBridgeStatus()

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

  const borderStyle = useMemo(() => {
    if (hasInputAmount && hasValidSelections && !isConnecting) {
      return {
        borderColor: '#D747FF',
        background:
          'linear-gradient(90deg, rgba(128, 0, 255, 0.2) 0%, rgba(255, 0, 191, 0.2) 100%)',
      }
    } else {
      return {
        borderColor: 'transparent',
        background: '',
      }
    }
  }, [hasInputAmount, hasValidSelections, isConnecting])

  return (
    <button
      data-test-id="connect-button"
      style={borderStyle}
      className={`
        flex items-center justify-center
        text-white py-lg px-lg rounded-sm
        text-center
        border
        hover:border-secondary
        h-8
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
        <div className="flex text-sm">Connect chain</div>
      )}
    </button>
  )
}

export function ConnectWalletButton({ highlight }: { highlight?: boolean }) {
  const [clientReady, setClientReady] = useState<boolean>(false)
  const { address } = useAccount()

  useEffect(() => {
    setClientReady(true)
  }, [])

  const borderStyle: {} = useMemo(() => {
    if (highlight) {
      return {
        borderColor: '#D747FF',
        background:
          'linear-gradient(90deg, rgba(128, 0, 255, 0.2) 0%, rgba(255, 0, 191, 0.2) 100%)',
      }
    } else {
      return {
        borderColor: '#565058',
        background: '',
      }
    }
  }, [highlight])

  return (
    <div data-test-id="connect-wallet-button">
      {clientReady && (
        <ConnectButton.Custom>
          {({ account, chain, openConnectModal, mounted, openChainModal }) => {
            return (
              <>
                {(() => {
                  if (!mounted || !account || !chain || !address) {
                    return (
                      <button
                        style={borderStyle}
                        className={`
                          flex items-center mr-2 py-md px-md
                          text-sm text-white
                          border rounded-sm 
                          hover:border-secondary
                        `}
                        onClick={openConnectModal}
                      >
                        Connect wallet
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
