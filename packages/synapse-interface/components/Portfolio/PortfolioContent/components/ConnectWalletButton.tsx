import { useState, useEffect } from 'react'
import { useAccount } from 'wagmi'
import { ConnectButton } from '@rainbow-me/rainbowkit'

export function ConnectWalletButton() {
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
                        className="px-6 py-2 border border-purple-500 rounded-full hover:opacity-70"
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
      )}
    </div>
  )
}
