import { useMemo, useState, useEffect } from 'react'
import { ConnectButton } from '@rainbow-me/rainbowkit'
import { useAccount } from 'wagmi'
import { useTranslations } from 'next-intl'

import { MetamaskIcon } from '@icons/WalletIcons/Metamask'
import { CoinbaseWalletIcon } from '@icons/WalletIcons/CoinbaseWalletIcon'
import { WalletConnectIcon } from '@icons/WalletIcons/WalletConnectIcon'
import { IconProps, WalletId } from '@utils/types'
import Spinner from './icons/Spinner'

const WALLETS = [
  {
    id: WalletId.MetaMask,
    icon: MetamaskIcon,
  },
  {
    id: WalletId.WalletConnect,
    icon: WalletConnectIcon,
  },
  {
    id: WalletId.CoinbaseWallet,
    icon: CoinbaseWalletIcon,
  },
]

export const WalletIcon = ({
  walletId,
  ...rest
}: IconProps): JSX.Element | null => {
  const SelectedIcon = Object.values(WALLETS).find(
    ({ id }) => id === walletId
  )?.icon

  return SelectedIcon ? <SelectedIcon {...rest} /> : null
}

export const Wallet = () => {
  const { connector: activeConnector, address: connectedAddress } = useAccount()
  const { chain: currentChain } = useAccount()
  const walletId = activeConnector?.id

  const [mounted, setMounted] = useState(false)

  const t = useTranslations('Wallet')

  useEffect(() => {
    setMounted(true)
  }, [])

  const render = useMemo(() => {
    return (
      <ConnectButton.Custom>
        {({
          account,
          chain,
          openAccountModal,
          openChainModal,
          openConnectModal,
          authenticationStatus,
          mounted,
        }) => {
          const ready = mounted && authenticationStatus !== 'loading'
          return (
            <div
              {...(!ready && {
                'aria-hidden': true,
              })}
            >
              {(() => {
                if (!connectedAddress) {
                  return (
                    <button
                      onClick={openConnectModal}
                      type="button"
                      className={`
                        text-sm text-white outline-none active:outline-none
                        ring-none transition-all duration-100 transform-gpu
                        rounded-md py-2 px-2.5
                        focus:outline-none focus:ring-0 hover:bg-opacity-70
                        bg-bgLight hover:bg-bgLightest focus:bg-bgLightest
                        active:bg-bgLightest hover:!border-blue-500
                        border border-none border-transparent
                        whitespace-nowrap
                      `}
                    >
                      {t('Connect Wallet')}
                    </button>
                  )
                }
                if (chain?.unsupported) {
                  return (
                    <button
                      onClick={openChainModal}
                      type="button"
                      className={`
                        text-white transition-all duration-100th
                        rounded-md py-2 px-2.5 border border-bgLight
                        active:bg-bgLightest/10 hover:bg-bgLightest/10
                        whitespace-nowrap
                      `}
                    >
                      {t('Wrong Network')}
                    </button>
                  )
                }
                return (
                  <div className="flex gap-3">
                    <button
                      onClick={openChainModal}
                      type="button"
                      className={`
                        flex items-center gap-2 text-white transition-all duration-100th
                        rounded-md py-2 px-2.5 border border-bgLight
                        active:bg-bgLightest/10 hover:bg-bgLightest/10
                        whitespace-nowrap
                        `}
                    >
                      {account?.displayBalance ? (
                        account.displayBalance
                      ) : (
                        <Spinner />
                      )}
                      {chain && chain.hasIcon && (
                        <div
                          style={{
                            backgroundImage: chain.iconBackground,
                            width: 20,
                            height: 20,
                            borderRadius: 999,
                            overflow: 'hidden',
                            backgroundPosition: 'center',
                            color: '#ffffff',
                          }}
                        >
                          {chain.iconUrl && (
                            <img
                              alt={chain.name ?? 'Chain icon'}
                              src={chain.iconUrl}
                              style={{ width: 20, height: 20 }}
                            />
                          )}
                        </div>
                      )}
                    </button>

                    <button
                      onClick={openAccountModal}
                      type="button"
                      className={`
                        text-white transition-all duration-100 rounded-md
                        py-2 px-2.5 bg-bgLight hover:bg-opacity-70
                        hover:bg-bgLightest active:bg-bgLightest text-sm
                        whitespace-nowrap font-bold
                      `}
                    >
                      {account ? account.displayName : <Spinner />}
                    </button>
                  </div>
                )
              })()}
            </div>
          )
        }}
      </ConnectButton.Custom>
    )
  }, [connectedAddress, currentChain, walletId, t])

  return mounted && render
}
