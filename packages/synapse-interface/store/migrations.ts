import { createMigrate } from 'redux-persist'

import { AppState } from './reducer'

const isLegacyHyperliquidDeposit = (tx: {
  originChain?: { id?: number }
  originToken?: { routeSymbol?: string }
  destinationChain?: { chainSymbol?: string }
  destinationToken?: { routeSymbol?: string }
  bridgeModuleName?: string
}) =>
  tx.originChain?.id === 42161 &&
  tx.destinationChain?.chainSymbol === 'HYPERLIQUID' &&
  (tx.destinationToken ?? tx.originToken)?.routeSymbol === 'USDC' &&
  !tx.bridgeModuleName

export const migratePersistedState = createMigrate({
  4: (state) => {
    if (!state) return state
    const { transactions, _transactions } = state as typeof state &
      Partial<Pick<AppState, 'transactions' | '_transactions'>>
    return {
      ...state,
      ...(transactions && {
        transactions: {
          ...transactions,
          pendingBridgeTransactions: transactions.pendingBridgeTransactions.filter(
            (tx) => !isLegacyHyperliquidDeposit(tx)
          ),
        },
      }),
      ...(_transactions && {
        _transactions: {
          ..._transactions,
          transactions: _transactions.transactions.filter(
            (tx) => !isLegacyHyperliquidDeposit(tx)
          ),
        },
      }),
    }
  },
})
