import { usePortfolioBalancesAndAllowances } from '@/utils/hooks/usePortfolioBalances'
import { PortfolioTabManager } from './PortfolioTabManager'
import { useAccount, useNetwork, Address } from 'wagmi'
import { useState } from 'react'
import { ConnectWalletButton } from './ConnectWalletButton'
import HomeSvg from '../icons/HomeIcon'

export enum PortfolioTabs {
  HOME = 'home',
  PORTFOLIO = 'portfolio',
}

export const Portfolio = () => {
  const [tab, setTab] = useState<PortfolioTabs>(PortfolioTabs.HOME)

  const portfolioData = usePortfolioBalancesAndAllowances()

  const { address } = useAccount()
  const { chain } = useNetwork()

  return (
    <div className="flex flex-col w-1/2">
      <PortfolioTabManager activeTab={tab} setTab={setTab} />
      <div className="border-t border-solid border-[#28282F] mt-4">
        {tab === PortfolioTabs.HOME && <HomeContent />}
        {tab === PortfolioTabs.PORTFOLIO && (
          <PortfolioContent connectedAddress={address} />
        )}
      </div>
    </div>
  )
}

type PortfolioContentProps = {
  connectedAddress: Address | string
}

const PortfolioContent = ({ connectedAddress }: PortfolioContentProps) => {
  return (
    <div className="">
      <PortfolioAssetHeader />
      {connectedAddress ? (
        <></>
      ) : (
        <>
          <p
            className={`
            text-[#CCCAD3BF] mt-6 mb-4 pb-6
            border-b border-solid border-[#28282F]
            `}
          >
            Your bridgable assets appear here when your wallet is connected.
          </p>
          <ConnectWalletButton />
        </>
      )}
    </div>
  )
}

const PortfolioAssetHeader = () => {
  return (
    <div className="flex text-[#CCCAD3BF] my-2">
      <div className="w-1/2 text-left">Token</div>
      <div className="w-1/2 text-left">Amount</div>
    </div>
  )
}

const HomeContent = () => {
  return (
    <div className="my-4 font-thin text-white">
      <p className="mb-3">
        Synapse is the most widely used, extensible, and secure cross-chain
        communications network.
      </p>
      <p className="mb-5">
        Preview your route in the Bridge panel, and connect your wallet when
        you're ready to authorize your transaction.
      </p>
      <ConnectWalletButton />
    </div>
  )
}
