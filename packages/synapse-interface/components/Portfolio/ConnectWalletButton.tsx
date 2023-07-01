import { ConnectButton } from '@rainbow-me/rainbowkit'
import { useAccount } from 'wagmi'

export function ConnectWalletButton() {
  const { address } = useAccount()

  const buttonClassName = `
    h-10 border-[#AC8FFF] flex items-center border
    text-base text-white px-6 py-5 hover:opacity-75 rounded-lg
    text-center transform-gpu transition-all duration-75
  `

  const buttonStyle = {
    background:
      'linear-gradient(310.65deg, rgba(255, 0, 255, 0.2) -17.9%, rgba(172, 143, 255, 0.2) 86.48%)',
    borderRadius: '30px',
  }

  return (
    !address && (
      <div data-test-id="connect-wallet-button">
        <ConnectButton.Custom>
          {({ account, chain, openConnectModal, mounted, openChainModal }) => {
            return (
              <>
                {(() => {
                  if (!mounted || !account || !chain) {
                    return (
                      <button
                        className={buttonClassName}
                        style={buttonStyle}
                        onClick={openConnectModal}
                      >
                        Connect Wallet
                      </button>
                    )
                  }
                })()}
              </>
            )
          }}
        </ConnectButton.Custom>
      </div>
    )
  )
}
