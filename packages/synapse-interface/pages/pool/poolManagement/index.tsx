import { useState } from 'react'

import LiquidityManagementTabs from '../components/LiquidityManagementTabs'
import Deposit from './Deposit'
import Withdraw from './Withdraw'
import { PoolData, PoolUserData } from '@types'
import { Token } from '@types'
import { useSelector } from 'react-redux'
import { RootState } from '@/store/store'

const PoolManagement = ({
  pool,
  address,
  chainId,
}: {
  pool: Token
  address: string
  chainId: number
}) => {
  const [cardNav, setCardNav] = useState(getLiquidityMode('#addLiquidity')) // 'addLiquidity'

  const { poolData } = useSelector((state: RootState) => state.poolData)
  const { poolUserData } = useSelector((state: RootState) => state.poolUserData)

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
          {cardNav === 'addLiquidity' && poolUserData.tokens && (
            <Deposit address={address} chainId={chainId} />
          )}
          {cardNav === 'removeLiquidity' && (
            <Withdraw
              pool={pool}
              chainId={chainId}
              address={address}
              poolData={poolData}
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
