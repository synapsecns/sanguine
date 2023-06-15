import { useState } from 'react'

import LiquidityManagementTabs from '../components/LiquidityManagementTabs'
import Deposit from './Deposit'
import Withdraw from './Withdraw'
import { PoolData, PoolUserData } from '@types'
import { Token } from '@types'

const PoolManagement = ({
  pool,
  address,
  chainId,
  poolData,
  poolUserData,
  refetchCallback,
}: {
  pool: Token
  address: string
  chainId: number
  poolData: PoolData
  poolUserData: PoolUserData
  refetchCallback: () => void
}) => {
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
            <Deposit
              pool={pool}
              address={address}
              chainId={chainId}
              poolData={poolData}
              poolUserData={poolUserData}
              refetchCallback={refetchCallback}
            />
          )}
          {cardNav === 'removeLiquidity' && (
            <Withdraw
              pool={pool}
              chainId={chainId}
              address={address}
              poolData={poolData}
              poolUserData={poolUserData}
              refetchCallback={refetchCallback}
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
