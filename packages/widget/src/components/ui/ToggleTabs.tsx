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
  const baseTabClass = 'basis-full rounded p-1.5'
  const activeStyle = {
    background: 'var(--synapse-button-bg)',
    color: 'var(--synapse-button-text',
    cursor: 'default',
  }

  return (
    <div
      className="flex mx-1 my-2 text-sm text-center cursor-pointer"
      style={{ background: 'var(--synapse-root'}}
      role="group"
    >
      <div
        className={`${baseTabClass} ${selectedTab !== 'All' && 'hover:opacity-70'}`}
        style={selectedTab === 'All' ? activeStyle : {}}
        onClick={() => onTabSelect('All')}
      >
        All
      </div>
      <div
        className={`${baseTabClass} ${selectedTab !== 'Target' && 'hover:opacity-70'}`}
        style={selectedTab === 'Target' ? activeStyle : {}}
        onClick={() => onTabSelect('Target')}
      >
        {protocolName ?? 'Target'}
      </div>
    </div>
  )
}
