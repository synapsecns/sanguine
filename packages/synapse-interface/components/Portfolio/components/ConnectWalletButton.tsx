import { useState, useEffect } from 'react'
import { useAccount } from 'wagmi'
import { ConnectButton } from '@rainbow-me/rainbowkit'
import { useTranslations } from 'next-intl'

export function ConnectWalletButton() {
  const [clientReady, setClientReady] = useState<boolean>(false)
  const { address } = useAccount()

  const t = useTranslations('Wallet')

  useEffect(() => {
    setClientReady(true)
  }, [])

  const buttonClassName = `
    h-10 border-[#CA5CFF] border-[1.5px] flex items-center border
    text-base text-white px-6 py-5 hover:opacity-75 rounded-lg
    text-center transform-gpu transition-all duration-75
  `

  const buttonStyle = {
    borderRadius: '30px',
  }

  return (
    <div id="connect-wallet-button">
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
