import { useState } from 'react'
import numeral from 'numeral'

import Card from '@tw/Card'
import Grid from '@tw/Grid'

export function Stats({ bridgeVolume, transactions, addresses, setChartType }) {
  const [activeState, setState] = useState('BRIDGEVOLUME')

  function updateState(string) {
    setState(string)
    setChartType(string)
  }

  return (
    <Grid cols={{ sm: 1, md: 3, lg: 3 }} gap={4} className="my-3">
      <StatCard
        title="Bridge Volume"
        active={activeState === 'BRIDGEVOLUME'}
        onClick={() => updateState('BRIDGEVOLUME')}
      >
        <div className="text-4xl font-bold text-white">
          {numeral(bridgeVolume / 1000000).format('$0.000')}m
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

export function StatCard({ onClick, title, children, active = false }) {
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
          30-day
        </div>
      </div>
    </Card>
  )
}
