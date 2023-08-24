import { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { Address } from 'wagmi'
import { RootState } from '@/store/store'

import LiquidityManagementTabs from '../components/LiquidityManagementTabs'
import Deposit from './Deposit'
import Withdraw from './Withdraw'
import LoadingSpinner from '@/components/ui/tailwind/LoadingSpinner'
import {
  fetchPoolUserData,
  resetPoolUserData,
} from '@/slices/poolUserDataSlice'

const PoolManagement = ({
  address,
  chainId,
}: {
  address: Address
  chainId: number
}) => {
  const [cardNav, setCardNav] = useState(getLiquidityMode('#addLiquidity')) // 'addLiquidity'

  const { pool } = useSelector((state: RootState) => state.poolData)
  const { poolUserData, isLoading } = useSelector(
    (state: RootState) => state.poolUserData
  )

  const dispatch: any = useDispatch()

  useEffect(() => {
    if (pool && address) {
      dispatch(resetPoolUserData())
      dispatch(fetchPoolUserData({ pool, address }))
    }
  }, [pool, address])

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-full">
        <LoadingSpinner />
      </div>
    )
  }

  return (
    <div>
      <div className="rounded-lg text-default">
        <LiquidityManagementTabs
          cardNav={cardNav}
          setCardNav={(val) => {
            setCardNav(val)
          }}
        />
        <div className="mt-4">
          {cardNav === 'addLiquidity' && poolUserData.tokens && (
            <Deposit address={address} chainId={chainId} />
          )}
          {cardNav === 'removeLiquidity' && <Withdraw address={address} />}
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
