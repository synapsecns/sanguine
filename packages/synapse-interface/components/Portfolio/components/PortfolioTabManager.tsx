import { useAccount } from 'wagmi'
import { useTranslations } from 'next-intl'

import { useAppDispatch } from '@/store/hooks'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { PortfolioTabs, setActiveTab } from '@/slices/portfolio/actions'
import { SearchBar } from './SearchBar'
import { _Transactions } from '../../_Transaction/_Transactions'

export const PortfolioTabManager = () => {
  const dispatch = useAppDispatch()
  const { activeTab } = usePortfolioState()
  const { address } = useAccount()

  const handleTabChange = (newTab: PortfolioTabs) => {
    dispatch(setActiveTab(newTab))
  }

  const t = useTranslations()

  return (
    <div data-test-id="portfolio-tab-manager" className="flex flex-col">
      <div className="flex items-center">
        <Tab
          display={t('Portfolio.Portfolio')}
          activeTab={activeTab}
          tabType={PortfolioTabs.PORTFOLIO}
          handleTabChange={handleTabChange}
        />
        <Tab
          display={t('Activity.Activity')}
          activeTab={activeTab}
          tabType={PortfolioTabs.ACTIVITY}
          handleTabChange={handleTabChange}
        />
        <SearchBar />
      </div>
      {address && <_Transactions connectedAddress={address} />}
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
  return (
    <button
      id="tab"
      onClick={() => handleTabChange(tabType)}
      className={`
        font-medium text-2xl text-gray-500
        border-b-2 border-transparent
        focus:outline-none focus:ring-0 active:outline-none active:ring:0 outline-none
        hover:text-white transform-gpu transition-all duration-75
        ${isCurrentlyActive && 'text-white'}
      `}
      style={{
        borderImage: isCurrentlyActive
          ? 'linear-gradient(to right, rgba(255, 0, 255, 1), rgba(172, 143, 255, 1)) 1'
          : 'none',
      }}
    >
      <div className="p-2">{display}</div>
    </button>
  )
}
