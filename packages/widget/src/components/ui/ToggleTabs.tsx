import { useBridgeState } from '@/state/slices/bridge/hooks'

export type TabOption = 'All' | 'Target'

type ToggleTabsProps = {
  selectedTab: TabOption
  onTabSelect: (tab: TabOption) => void
}

export const ToggleTabs: React.FC<ToggleTabsProps> = ({
  selectedTab,
  onTabSelect,
}) => {
  const { protocolName } = useBridgeState()

  const baseTabClass = 'basis-full rounded p-1.5'
  const activeTabClass = `${baseTabClass} cursor-pointer hover:opacity-70`
  const activeStyle = {
    background: 'var(--synapse-button-bg)',
    color: 'var(--synapse-button-text)',
  }

  return (
    <div
      className="flex text-sm text-center cursor-default rounded border border-solid border-[--synapse-border]"
      style={{ background: 'var(--synapse-root)' }}
      role="group"
    >
      <div
        className={selectedTab === 'All' ? baseTabClass : activeTabClass}
        style={selectedTab === 'All' ? activeStyle : {}}
        onClick={() => onTabSelect('All')}
      >
        All
      </div>
      <div
        className={selectedTab === 'Target' ? baseTabClass : activeTabClass}
        style={selectedTab === 'Target' ? activeStyle : {}}
        onClick={() => onTabSelect('Target')}
      >
        {protocolName ?? 'Target'}
      </div>
    </div>
  )
}
