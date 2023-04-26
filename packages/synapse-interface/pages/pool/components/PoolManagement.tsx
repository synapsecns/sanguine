import { useState } from 'react'

import LiquidityManagementTabs from './LiquidityManagementTabs'
import PoolManagementDeposit from './PoolManagementDeposit'
import PoolManagementWithdraw from './PoolManagementWithdraw'
import { Token } from '@types'
const PoolManagement = ({
  pool,
  address,
  chainId,
}: {
  pool: Token
  address: string
  chainId: number
}) => {
  // const [cardNav, setCardNav] = useState(getLiquidityMode(location.hash)) // 'addLiquidity'

  const [cardNav, setCardNav] = useState('#addLiquidity') // 'addLiquidity'
  return (
    <div>
      <div className="rounded-lg text-default">
        <LiquidityManagementTabs
          cardNav={cardNav}
          setCardNav={(val) => {
            // history.push(`#${val}`) TODO
            setCardNav(val)
          }}
        />
        <div className="mt-4">
          {cardNav === 'addLiquidity' && (
            <PoolManagementDeposit
              poolName={pool.name}
              address={address}
              chainId={chainId}
            />
          )}
          {cardNav === 'removeLiquidity' && (
            <PoolManagementWithdraw
              pool={pool}
              chainId={chainId}
              address={address}
            />
          )}
        </div>
      </div>
    </div>
  )
}

function getLiquidityMode(browserHash) {
  switch (browserHash) {
    case '#addLiquidity':
      return 'addLiquidity'
    case '#removeLiquidity':
      return 'removeLiquidity'
    default:
      return 'addLiquidity'
  }
}
export default PoolManagement
