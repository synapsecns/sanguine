import { ConnectButton, AuthenticationStatus } from '@rainbow-me/rainbowkit'
import { useAccount } from 'wagmi'

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
  const { connector: activeConnector } = useAccount()
  const walletId = activeConnector?.id

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
                    className="flex items-center text-md hover:cursor-pointer header-button"
                  >
                    <span className="mr-1 disconnected">● </span>
                    Disconnected
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
                    style={{ display: 'flex', alignItems: 'center' }}
                    type="button"
                  >
                    {chain.hasIcon && (
                      <div
                        style={{
                          background: chain.iconBackground,
                          width: 12,
                          height: 12,
                          borderRadius: 999,
                          overflow: 'hidden',
                          marginRight: 4,

                        }}
                      >
                        {chain.iconUrl && (
                          <img
                            alt={chain.name ?? 'Chain icon'}
                            src={chain.iconUrl}
                            style={{ width: 12, height: 12}}

                          />
                        )}
                      </div>
                    )}
                    {chain.name}
                  </button>

                  <button onClick={openAccountModal} type="button" style={{backgroundColor: 'gray', border: "1px black solid", borderRadius: "5px"}}>
                    {account.displayName}
                    {account.displayBalance
                      ? ` (${account.displayBalance})`
                      : ''}
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
