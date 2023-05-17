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
}: {
  pool: Token
  address: string
  chainId: number
  poolData: PoolData
  poolUserData: PoolUserData
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
          {cardNav === 'addLiquidity' &&
            (pool && poolUserData && poolData && address ? (
              <Deposit
                pool={pool}
                address={address}
                chainId={chainId}
                poolData={poolData}
                poolUserData={poolUserData}
              />
            ) : (
              <div className="w-full text-center mt-[80px] text-sm text-white">
                <p>connect wallet</p>
              </div>
            ))}
          {cardNav === 'removeLiquidity' &&
            (pool && poolUserData && poolData && address ? (
              <Withdraw
                pool={pool}
                chainId={chainId}
                address={address}
                poolData={poolData}
                poolUserData={poolUserData}
              />
            ) : (
              <div className="w-full text-center mt-[80px] text-sm text-white">
                <p>connect wallet</p>
              </div>
            ))}
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
