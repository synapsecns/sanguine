import { useState } from 'react'
import { useLocation, useHistory } from 'react-router-dom'

import LiquidityManagementTabs from './LiquidityManagementTabs'
import PoolManagementDeposit from './PoolManagementDeposit'
import PoolManagementWithdraw from './PoolManagementWithdraw'

export default function PoolManagement({
  poolName,
  poolStakingLink,
  poolStakingLinkText,
}) {
  const location = useLocation()
  const history = useHistory()

  const [cardNav, setCardNav] = useState(getLiquidityMode(location.hash)) // 'addLiquidity'

  return (
    <div>
      <div className="rounded-lg text-default">
        <LiquidityManagementTabs
          cardNav={cardNav}
          setCardNav={(val) => {
            history.push(`#${val}`)
            setCardNav(val)
          }}
        />
        <div className="mt-4">
          {cardNav === 'addLiquidity' && (
            <PoolManagementDeposit
              poolName={poolName}
              poolStakingLink={poolStakingLink}
              poolStakingLinkText={poolStakingLinkText}
            />
          )}
          {cardNav === 'removeLiquidity' && (
            <PoolManagementWithdraw poolName={poolName} />
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
