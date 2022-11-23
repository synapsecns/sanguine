import { Routes, Route } from 'react-router-dom'

import { Address } from '@pages/Address'
import { BridgeTransaction } from '@pages/BridgeTransaction'
import { BridgeTransactions } from '@pages/BridgeTransactions'
import { Chain } from '@pages/Chain'
import { Home } from '@pages/Home'
import { Leaderboard } from '@pages/Leaderboard'
import { TokenAddress } from '@pages/TokenAddress'

import {
  TRANSACTIONS_PATH,
  TRANSACTION_PATH,
  ACCOUNTS_PATH,
  CHAINS_PATH,
  TOKEN_ADDRESSES_PATH,
} from '@urls'

export default function App() {
  return (
    <Routes>
      <Route exact path={'/'} element={<Home />} />
      <Route exact path={TRANSACTIONS_PATH} element={<BridgeTransactions />} />
      <Route
        path={`${TRANSACTION_PATH}/:kappa`}
        element={<BridgeTransaction />}
      />
      <Route path={`${CHAINS_PATH}/:chainId`} element={<Chain />} />
      <Route path={`${ACCOUNTS_PATH}/:address`} element={<Address />} />
      <Route
        path={`${TOKEN_ADDRESSES_PATH}/:tokenAddress`}
        element={<TokenAddress />}
      />
      <Route path={'/leaderboard'} element={<Leaderboard />} />
    </Routes>
  )
}
