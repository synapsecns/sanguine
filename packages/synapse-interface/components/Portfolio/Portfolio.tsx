import { usePortfolioBalancesAndAllowances } from '@/utils/hooks/usePortfolioBalances'
import { useAccount, useNetwork } from 'wagmi'
import { useState } from 'react'
import HomeSvg from '../icons/HomeIcon'

enum PortfolioTabs {
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
    </div>
  )
}

type PortfolioTabManagerProps = {
  activeTab: PortfolioTabs
  setTab: React.Dispatch<React.SetStateAction<PortfolioTabs>>
}

const PortfolioTabManager = ({
  activeTab,
  setTab,
}: PortfolioTabManagerProps) => {
  const handleTabChange = (newTab: PortfolioTabs) => {
    setTab(newTab)
  }
  return (
    <div className="flex">
      <Tab
        display={<HomeSvg />}
        activeTab={activeTab}
        tabType={PortfolioTabs.HOME}
        handleTabChange={handleTabChange}
      />
      <Tab
        display="Portfolio"
        activeTab={activeTab}
        tabType={PortfolioTabs.PORTFOLIO}
        handleTabChange={handleTabChange}
      />
    </div>
  )
}

type TabProps = {
  display: string | JSX.Element
  activeTab: PortfolioTabs
  tabType: PortfolioTabs
  handleTabChange: (newTab: PortfolioTabs) => void
}

const Tab = ({ display, activeTab, tabType, handleTabChange }: TabProps) => {
  const isCurrentlyActive: boolean = activeTab === tabType
  const radialGradient = 'radial-gradient text-white'
  return (
    <button
      className={`
      font-medium text-2xl text-gray-500
      border-b-2 border-transparent
      mr-2
      ${isCurrentlyActive && radialGradient}
      `}
      onClick={() => handleTabChange(tabType)}
    >
      {display}
    </button>
  )
}
