import { useState } from 'react'
import numeral from 'numeral'
import Card from '@components/tailwind/Card'
import Grid from '@components/tailwind/Grid'

export function Stats({
  bridgeVolume,
  transactions,
  addresses,
  setChartType,
  allTime,
}) {
  const [activeState, setState] = useState('VOLUME')

  function updateState(string) {
    setState(string)
    setChartType(string)
  }

  if (allTime) {
    return (
      <Grid cols={{ sm: 1, md: 3, lg: 3 }} gap={4} className="my-3">
        <AllTimeStatCard
          title="Bridge Volume"
          active={activeState === 'VOLUME'}
          onClick={() => updateState('VOLUME')}
        >
          <div className="text-4xl font-bold text-white">
            {numeral(bridgeVolume / 1000000000).format('$0.000')}B
          </div>
        </AllTimeStatCard>
        <AllTimeStatCard
          title="Transactions"
          active={activeState === 'TRANSACTIONS'}
          onClick={() => updateState('TRANSACTIONS')}
        >
          <div className="text-4xl font-bold text-white">
            {numeral(transactions).format('0,0')}
          </div>
        </AllTimeStatCard>
        <AllTimeStatCard
          title="Addresses"
          active={activeState === 'ADDRESSES'}
          onClick={() => updateState('ADDRESSES')}
        >
          <div className="text-4xl font-bold text-white">
            {numeral(addresses).format('0,0')}
          </div>
        </AllTimeStatCard>
      </Grid>
    )
  } else {
    return (
      <Grid cols={{ sm: 1, md: 3, lg: 3 }} gap={4} className="my-3">
        <StatCard
          title="Bridge Volume"
          active={activeState === 'VOLUME'}
          onClick={() => updateState('VOLUME')}
        >
          <div className="text-4xl font-bold text-white">
            {numeral(bridgeVolume / 1000000).format('$0.000')}M
          </div>
        </StatCard>
        <StatCard
          title="Transactions"
          active={activeState === 'TRANSACTIONS'}
          onClick={() => updateState('TRANSACTIONS')}
        >
          <div className="text-4xl font-bold text-white">
            {numeral(transactions).format('0,0')}
          </div>
        </StatCard>
        <StatCard
          title="Addresses"
          active={activeState === 'ADDRESSES'}
          onClick={() => updateState('ADDRESSES')}
        >
          <div className="text-4xl font-bold text-white">
            {numeral(addresses).format('0,0')}
          </div>
        </StatCard>
      </Grid>
    )
  }
}

export function StatCard({
  onClick,
  title,
  children,
  active = false,
  duration = '30-day',
}) {
  const activeClass = active ? 'opacity-100' : 'opacity-20 hover:opacity-100'
  return (
    <Card
      className={`px-0 pb-2 space-y-3 text-white bg-transparent cursor-pointer ${activeClass}`}
      onClick={onClick}
    >
      <div className="text-xl">{title}</div>
      {children}
      <div className="flex space-x-2 text-sm font-medium">
        <div className="text-transparent bg-clip-text bg-gradient-to-r from-purple-500 to-purple-400">
          {duration}
        </div>
      </div>
    </Card>
  )
}

export function AllTimeStatCard({
  onClick,
  title,
  children,
  active = false,
  duration = 'All-Time',
}) {
  const activeClass = active ? 'opacity-100' : 'opacity-20 hover:opacity-100'
  return (
    <Card
      className={`px-0 pb-2 space-y-3 text-white bg-transparent cursor-pointer ${activeClass}`}
      onClick={onClick}
    >
      <div className="text-xl">{title}</div>
      {children}
      <div className="flex space-x-2 text-sm font-medium">
        <div className="text-transparent bg-clip-text bg-gradient-to-r from-purple-500 to-purple-400">
          {duration}
        </div>
      </div>
    </Card>
  )
}
