import { useState } from 'react'

import LiquidityManagementTabs from './LiquidityManagementTabs'
import PoolManagementDeposit from './PoolManagementDeposit'
import PoolManagementWithdraw from './PoolManagementWithdraw'
import { Token } from '@types'
const PoolManagement = ({
  pool,
  address,
  chainId,
  poolData,
  poolUserData,
}: {
  pool: Token
  address: string
  chainId: number
  poolData: any
  poolUserData: any
}) => {
  // const [cardNav, setCardNav] = useState(getLiquidityMode(location.hash)) // 'addLiquidity'

  const [cardNav, setCardNav] = useState(getLiquidityMode('#addLiquidity')) // 'addLiquidity'
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
              pool={pool}
              address={address}
              chainId={chainId}
              poolData={poolData}
              poolUserData={poolUserData}
            />
          )}
          {cardNav === 'removeLiquidity' && (
            <PoolManagementWithdraw
              pool={pool}
              chainId={chainId}
              address={address}
              poolUserData={poolUserData}
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
