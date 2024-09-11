import { useEffect, useState } from 'react'
import { useDispatch } from 'react-redux'
import { useAccount } from 'wagmi'
import { switchChain } from '@wagmi/core'
import { LoaderIcon } from 'react-hot-toast'
import { ConnectButton } from '@rainbow-me/rainbowkit'
import { useTranslations } from 'next-intl'

import { CHAINS_BY_ID } from '@/constants/chains'
import { setFromChainId } from '@/slices/bridge/reducer'
import { getNetworkButtonBorderHover, getNetworkHover } from '@/styles/chains'
import { joinClassNames } from '@/utils/joinClassNames'
import { wagmiConfig } from '@/wagmiConfig'

const Indicator = ({ className }) => (
  <span
    className={`w-2 h-2 rounded-full ${
      className.match(/^border-/) ? `border` : ''
    } ${className}`}
  />
)

export const ConnectedIndicator = () => {
  const classNames = {
    flex: 'flex items-center gap-2',
    space: 'px-3 py-1 rounded-full',
    hover: 'hover:opacity-80',
    font: 'text-sm',
  }

  const t = useTranslations('Wallet')

  return (
    <button
      data-test-id="connected-button"
      disabled
      className={joinClassNames(classNames)}
    >
      <Indicator className="bg-green-500 dark:bg-green-400" />
      {t('Connected')}
    </button>
  )
}

export const ConnectToNetworkButton = ({ chainId }: { chainId: number }) => {
  const [isConnecting, setIsConnecting] = useState<boolean>(false)
  const dispatch = useDispatch()
  const chain = CHAINS_BY_ID[chainId]

  const t = useTranslations('Wallet')

  function scrollToTop(): void {
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }

  const handleConnectNetwork: () => Promise<void> = async () => {
    setIsConnecting(true)
    try {
      await switchChain(wagmiConfig, { chainId }).then((success) => {
        success && dispatch(setFromChainId(chainId))
        scrollToTop()
      })
    } catch (error) {
      error && setIsConnecting(false)
    }
  }

  const classNames = {
    flex: 'flex items-center gap-2',
    space: 'px-3 py-1 rounded-full',
    border: 'border border-transparent',
    font: 'text-sm',
    bgHover: getNetworkHover(chain?.color),
    borderHover: getNetworkButtonBorderHover(chain?.color),
    active: 'hover:active:opacity-80',
  }

  return (
    <button
      data-test-id="connect-button"
      className={joinClassNames(classNames)}
      onClick={handleConnectNetwork}
    >
      {isConnecting ? (
        <>
          <Indicator className="border-green-500 dark:border-green-400" />
          {t('Connecting')}
          <LoaderIcon />
        </>
      ) : (
        <>
          <Indicator className="border-indigo-500 dark:border-indigo-300" />
          {t('Switch Network')}
        </>
      )}
    </button>
  )
}

export function ConnectWalletButton() {
  const [clientReady, setClientReady] = useState<boolean>(false)
  const { address } = useAccount()

  const t = useTranslations('Wallet')

  useEffect(() => {
    setClientReady(true)
  }, [])

  const classNames = {
    flex: 'flex items-center gap-2',
    space: 'px-3 py-1 rounded-full',
    border: 'border border-transparent',
    hover:
      'hover:bg-fuchsia-50 hover:border-fuchsia-500 hover:dark:bg-fuchsia-950',
    font: 'text-sm',
    active: 'active:opacity-80',
  }

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
                        className={joinClassNames(classNames)}
                        onClick={openConnectModal}
                      >
                        <Indicator className="border-fuchsia-500" />
                        {t('Connect Wallet')}
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
