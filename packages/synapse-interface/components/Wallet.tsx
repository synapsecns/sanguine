import { useMemo } from 'react'
import { ConnectButton } from '@rainbow-me/rainbowkit'
import { useAccount, useNetwork } from 'wagmi'
import { MetamaskIcon } from '@icons/WalletIcons/Metamask'
import { CoinbaseWalletIcon } from '@icons/WalletIcons/CoinbaseWalletIcon'
import { WalletConnectIcon } from '@icons/WalletIcons/WalletConnectIcon'
import { IconProps, WalletId } from '@utils/types'

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
  const { chain } = useNetwork()
  const walletId = activeConnector?.id

  return useMemo(() => {}, [connectedAddress, chain])

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
        // Note: If your app doesn't use authentication, you
        // can remove all 'authenticationStatus' checks
        const ready = mounted && authenticationStatus !== 'loading'
        const connected =
          ready &&
          account &&
          chain &&
          (!authenticationStatus || authenticationStatus === 'authenticated')
        return (
          <div
            {...(!ready && {
              'aria-hidden': true,
            })}
          >
            {(() => {
              if (!connected) {
                return (
                  <button
                    onClick={openConnectModal}
                    type="button"
                    className="text-sm flex items-center group cursor-pointer text-white outline-none active:outline-none ring-none transition-all duration-100 transform-gpu w-full rounded-lg py-2 pl-2.5 pr-2.5 group focus:outline-none focus:ring-0 hover:bg-opacity-70 bg-bgLight hover:bg-bgLightest focus:bg-bgLightest active:bg-bgLightest border-transparent hover:!border-blue-500 flex-shrink border border-none"
                  >
                    Connect Wallet
                  </button>
                )
              }
              if (chain.unsupported) {
                return (
                  <button onClick={openChainModal} type="button">
                    Wrong network
                  </button>
                )
              }
              return (
                <div style={{ display: 'flex', gap: 12 }}>
                  <button
                    onClick={openChainModal}
                    style={{ display: 'flex', alignItems: 'center', gap: 7 }}
                    type="button"
                    className="text-white transition-all duration-100th w-fit cursor-pointer rounded-lg py-2 pl-2.5 pr-2.5 border border-bgLight active:bg-bgLightest/10 hover:bg-bgLightest/10"
                  >
                    {account.displayBalance ? account.displayBalance : 0}
                    {chain.hasIcon && (
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
                    className="flex items-center cursor-pointer text-white transition-all duration-100 w-fit rounded-lg py-2 pl-2.5 pr-2.5  bg-bgLight hover:bg-opacity-70 hover:bg-bgLightest active:bg-bgLightest text-sm"
                  >
                    {account.displayName}
                  </button>
                </div>
              )
            })()}
          </div>
        )
      }}
    </ConnectButton.Custom>

    // <>
    //   <div className='flex items-center'>
    //     <WalletConnectButton
    //       setShowWalletModal={setShowWalletModal}

    //     />
    //   </div>
    //   <Modal isOpen={showWalletModal} onClose={handleClose}>
    //     <ConnectWallet onClose={handleClose} />
    //   </Modal>
    // </>
  )
}

function FormattedDisplayName(displayName: string) {
  const [, hex] = displayName.split('0x')
  return '0x' + hex
}

// :{
//   account?: {
//       address: string;
//       balanceDecimals?: number;
//       balanceFormatted?: string;
//       balanceSymbol?: string;
//       displayBalance?: string;
//       displayName: string;
//       ensAvatar?: string;
//       ensName?: string;
//       hasPendingTransactions: boolean;
//   };
//   chain?: {
//       hasIcon: boolean;
//       iconUrl?: string;
//       iconBackground?: string;
//       id: number;
//       name?: string;
//       unsupported?: boolean;
//   };
//   mounted: boolean;
//   authenticationStatus?: AuthenticationStatus;
//   openAccountModal: () => void;
//   openChainModal: () => void;
//   openConnectModal: () => void;
//   accountModalOpen: boolean;
//   chainModalOpen: boolean;
//   connectModalOpen: boolean;
// }
