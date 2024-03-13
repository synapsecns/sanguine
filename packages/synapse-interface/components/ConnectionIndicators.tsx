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

const join = (a) => Object.values(a).join(' ')

const Indicator = ({ className }) => (
  <span
    className={`w-2 h-2 rounded-full ${
      className.match(/^border-/) ? `border` : ''
    } ${className}`}
  />
)

export const ConnectedIndicator = () => {
  const className = join({
    flex: 'flex items-center gap-2',
    space: 'px-3 py-1 rounded-full',
    hover: 'hover:opacity-70 cursor-default',
    font: 'text-sm',
  })
  return (
    <button data-test-id="connected-button" disabled className={className}>
      <Indicator className="bg-green-500 dark:bg-green-300" />
      Connected
    </button>
  )
}

const DisconnectedIndicator = () => {
  const { openConnectModal } = useConnectModal()
  const { fromChainId } = useBridgeState()
  const chain = CHAINS_BY_ID[fromChainId]

  const className = join({
    flex: 'flex items-center gap-2',
    space: 'px-3 py-1 rounded-full',
    border: 'border border-transparent',
    font: 'text-sm',
    bgHover: getNetworkHover(chain?.color),
    // bgActive: getNetworkButtonBgClassNameActive(chain?.color),
    borderHover: getNetworkButtonBorderHover(chain?.color),
    // borderActive: getNetworkButtonBorderActive(chain?.color),
    active: 'hover:active:opacity-80',
  })

  return (
    <button
      data-test-id="disconnected-button"
      className={className}
      onClick={openConnectModal}
    >
      <Indicator className="bg-red-500" />
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

  const className = join({
    flex: 'flex items-center gap-2',
    space: 'px-3 py-1 rounded-full',
    border: 'border border-transparent',
    font: 'text-sm',
    bgHover: getNetworkHover(chain?.color),
    // bgActive: getNetworkButtonBgClassNameActive(chain?.color),
    borderHover: getNetworkButtonBorderHover(chain?.color),
    // borderActive: getNetworkButtonBorderActive(chain?.color),
    active: 'hover:active:opacity-80',
  })

  return (
    <button
      data-test-id="connect-button"
      className={className}
      onClick={handleConnectNetwork}
    >
      {isConnecting ? (
        <>
          <Indicator className="border-green-500 dark:border-green-300" />
          Connecting
          <LoaderIcon />
        </>
      ) : (
        <>
          <Indicator className="border-indigo-500 dark:border-indigo-300" />
          Switch Network
        </>
      )}
      {/* {isConnecting && <LoaderIcon />} */}
    </button>
  )
}

export function ConnectWalletButton() {
  const [clientReady, setClientReady] = useState<boolean>(false)
  const { address } = useAccount()

  useEffect(() => {
    setClientReady(true)
  }, [])

  const className = join({
    flex: 'flex items-center gap-2',
    space: 'px-3 py-1 rounded-full',
    border: 'border border-transparent',
    hover:
      'hover:bg-fuchsia-50 hover:border-fuchsia-500 hover:dark:bg-fuchsia-950',
    font: 'text-sm',
    active: 'active:opacity-80',
  })

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
                      <button className={className} onClick={openConnectModal}>
                        <Indicator className="border-fuchsia-500" />
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
