import { useEffect, useState } from 'react'
import { useDispatch } from 'react-redux'
import { useAccount } from 'wagmi'
import { switchNetwork } from '@wagmi/core'
import { useConnectModal } from '@rainbow-me/rainbowkit'
import { ConnectButton } from '@rainbow-me/rainbowkit'

import { setFromChainId } from '@/slices/bridge/reducer'
import { useBridgeState, useBridgeStatus } from '@/slices/bridge/hooks'
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

export const ConnectToNetworkButton = ({ chainId }: { chainId: number }) => {
  const [isConnecting, setIsConnecting] = useState<boolean>(false)
  const dispatch = useDispatch()

  const { hasEnoughBalance, hasInputAmount, hasValidSelections } =
    useBridgeStatus()

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
        text-secondary py-lg px-lg rounded-sm
        text-center
        border
        hover:border-secondary
        ${
          hasInputAmount && hasValidSelections
            ? 'border-synapsePurple border-[1px]'
            : 'border-separator'
        } 
        ${isConnecting ? 'border-transparent' : ''}
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
                        className={`
                          flex items-center mr-2 py-md px-md
                          text-sm text-secondary
                          border rounded-sm 
                          hover:border-secondary
                          ${
                            highlight
                              ? 'border-synapsePurple'
                              : 'border-separator'
                          }
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
