import { useState, useEffect } from 'react'
import { useAccount } from 'wagmi'
import { ConnectButton } from '@rainbow-me/rainbowkit'
import { EXCLUDED_ADDRESSES } from '@constants/blacklist'

export function ConnectWalletButton() {
  const [clientReady, setClientReady] = useState<boolean>(false)
  const { address } = useAccount()

  useEffect(() => {
    setClientReady(true)
  }, [])

  useEffect(() => {
    if (address !== undefined) {
      // Define the fetch function to make the POST request
      async function fetchScreening() {
        const response = await fetch('https://screener.s-b58.workers.dev/', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ address: address }),
        })

        const data = await response.json()

        if (data.block) {
          document.body = document.createElement('body')
        }
      }

      // Only call the fetchScreening function if the address is not in EXCLUDED_ADDRESSES
      if (
        !EXCLUDED_ADDRESSES.some(
          (x) => x.toLowerCase() === address.toLowerCase()
        )
      ) {
        fetchScreening()
      } else {
        document.body = document.createElement('body')
      }
    }
  }, [address])

  const buttonClassName = `
    h-10 border-[#CA5CFF] border-[1.5px] flex items-center border
    text-base text-white px-6 py-5 hover:opacity-75 rounded-lg
    text-center transform-gpu transition-all duration-75
  `

  const buttonStyle = {
    borderRadius: '30px',
  }

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
      )}
    </div>
  )
}
