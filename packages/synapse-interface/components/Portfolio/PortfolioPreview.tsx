import { useState } from 'react'
import { PageHeader } from '../PageHeader'
import { ConnectButton } from '@rainbow-me/rainbowkit'
import Card from '../ui/tailwind/Card'
import HomeSvg from '../icons/HomeIcon'

export const PortfolioPreview = () => {
  const [showContent, setShowContent] = useState('home')

  const renderCardContent = () => {
    if (showContent === 'home') {
      return (
        <div className="">
          <div className="flex items-center mb-4 space-x-2 text-lg">
            Synapse is the most widely used, extensible, and secure cross-chain
            communications network.
          </div>
          <div className="mb-5">
            Preview your route in the Bridge panel, and connect your wallet when
            you're ready to authorize your transaction.
          </div>
          <ConnectWallet />
        </div>
      )
    } else if (showContent === 'portfolio') {
      return (
        <div>
          <div className="flex items-center mb-5 space-x-2 text-lg">
            Your bridgable assets appear here when your wallet is connected.
          </div>
          <ConnectWallet />
        </div>
      )
    }
  }

  return (
    <div className="flex flex-col w-1/3">
      <div className="flex items-center space-x-2">
        <div
          onClick={() => setShowContent('home')}
          className="hover:opacity-80 hover:cursor-pointer"
        >
          <div className="flex items-center h-[32px]">
            <HomeSvg />
          </div>
          <div
            className={`mt-2 border-b-3 ${
              showContent === 'home'
                ? 'border-indigo-500'
                : 'border-transparent'
            }`}
          />
        </div>
        <div
          onClick={() => setShowContent('portfolio')}
          className="hover:opacity-80 hover:cursor-pointer"
        >
          <div className="flex items-center h-[32px]">
            <PageHeader title="Portfolio" subtitle="" />
          </div>
          <div
            className={`mt-2 border-b-3 ${
              showContent === 'portfolio'
                ? 'border-indigo-500'
                : 'border-transparent'
            }`}
          />
        </div>
      </div>
      <Card
        divider={false}
        className={`
          max-w-lg px-1 pb-0 mt-3 overflow-hidden
          rounded-xl
          bg-transparent
        text-white
        `}
      >
        {renderCardContent()}
      </Card>
    </div>
  )
}

export function ConnectWallet() {
  const buttonClassName = `
    h-10 border-[#AC8FFF] flex items-center border
    text-base px-4 py-3 hover:opacity-75 rounded-lg
    text-center transform-gpu transition-all duration-75
  `

  const buttonStyle = {
    background:
      'linear-gradient(310.65deg, rgba(255, 0, 255, 0.2) -17.9%, rgba(172, 143, 255, 0.2) 86.48%)',
    borderRadius: '16px',
  }
  return (
    <div>
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
}
