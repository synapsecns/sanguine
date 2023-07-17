import { PortfolioTabs } from './Portfolio'
import HomeSvg from '@icons/HomeIcon'

type PortfolioTabManagerProps = {
  activeTab: PortfolioTabs
  setTab: React.Dispatch<React.SetStateAction<PortfolioTabs>>
}

export const PortfolioTabManager = ({
  activeTab,
  setTab,
}: PortfolioTabManagerProps) => {
  const handleTabChange = (newTab: PortfolioTabs) => {
    setTab(newTab)
  }
  return (
    <div data-test-id="portfolio-tab-manager" className="flex">
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
  return (
    <button
      className={`
      font-medium text-2xl text-gray-500
      border-b-2 border-transparent
      mr-2 pb-2
      ${isCurrentlyActive && 'text-white'}
      `}
      onClick={() => handleTabChange(tabType)}
      style={{
        borderImage: isCurrentlyActive
          ? 'linear-gradient(to right, rgba(255, 0, 255, 1), rgba(172, 143, 255, 1)) 1'
          : 'none',
      }}
    >
      {display}
    </button>
  )
}
