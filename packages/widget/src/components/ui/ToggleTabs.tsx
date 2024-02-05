import { useBridgeState } from '@/state/slices/bridge/hooks'

export type TabOption = 'All' | 'Target'

type ToggleTabsProps = {
  selectedTab: TabOption
  onTabSelect: (tab: TabOption) => void
  isOrigin: boolean
}

export const ToggleTabs: React.FC<ToggleTabsProps> = ({
  selectedTab,
  onTabSelect,
  isOrigin,
}) => {
  const { protocolName } = useBridgeState()
  const baseTabClass =
    'flex-grow text-sm font-medium text-center text-[--synapse-primary] rounded-sm p-1 '

  const activeTabClass = 'bg-[var(--synapse-surface)]'
  const inactiveTabClass =
    'bg-[var(--synapse-select-bg)] hover:bg-[var(--synapse-surface)] hover:cursor-pointer'

  return (
    <div className="flex mt-2 mb-2" role="group">
      {isOrigin ? (
        <>
          <div
            className={`${baseTabClass} ${
              selectedTab === 'All' ? activeTabClass : inactiveTabClass
            }`}
            onClick={() => onTabSelect('All')}
          >
            All
          </div>
          <div
            className={`${baseTabClass} ${
              selectedTab === 'Target' ? activeTabClass : inactiveTabClass
            }`}
            onClick={() => onTabSelect('Target')}
          >
            {protocolName ?? 'Target'}
          </div>
        </>
      ) : (
        <>
          <div
            className={`${baseTabClass} ${
              selectedTab === 'Target' ? activeTabClass : inactiveTabClass
            }`}
            onClick={() => onTabSelect('Target')}
          >
            {protocolName ?? 'Target'}
          </div>
          <div
            className={`${baseTabClass} ${
              selectedTab === 'All' ? activeTabClass : inactiveTabClass
            }`}
            onClick={() => onTabSelect('All')}
          >
            All
          </div>
        </>
      )}
    </div>
  )
}
