import { usePortfolioBalancesAndAllowances } from '@/utils/hooks/usePortfolioBalances'
import { PortfolioTabManager } from './PortfolioTabManager'
import { useAccount, useNetwork } from 'wagmi'
import { useState } from 'react'
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
        {tab === PortfolioTabs.HOME && <HomeTabContent />}
      </div>
    </div>
  )
}

const HomeTabContent = () => {
  return (
    <>
      <div className="flex items-center mb-4 space-x-2 text-lg">
        Synapse is the most widely used, extensible, and secure cross-chain
        communications network.
      </div>
      <div className="mb-5">
        Preview your route in the Bridge panel, and connect your wallet when
        you're ready to authorize your transaction.
      </div>
    </>
  )
}
