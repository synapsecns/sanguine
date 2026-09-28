import { migratePersistedState } from '@/store/migrations'

describe('persisted transaction migration', () => {
  it('removes only old Arbitrum USDC deposits to HyperCore', async () => {
    const legacyPending = {
      originChain: { id: 42161 },
      originToken: { routeSymbol: 'USDC' },
      destinationChain: { id: 998, chainSymbol: 'HYPERLIQUID' },
    }
    const legacyActivity = {
      ...legacyPending,
      destinationToken: { routeSymbol: 'USDC' },
      bridgeModuleName: '',
      status: 'pending',
    }
    const relayDeposit = { ...legacyActivity, bridgeModuleName: 'Relay' }
    const synDeposit = {
      ...legacyActivity,
      originToken: { routeSymbol: 'SYN' },
      destinationToken: { routeSymbol: 'SYN' },
      bridgeModuleName: 'SYN',
    }
    const unrelated = {
      ...legacyActivity,
      destinationChain: { id: 1, chainSymbol: 'ETHEREUM' },
    }
    const state = {
      _persist: { version: 3, rehydrated: false },
      application: { lastConnectedAddress: '0xabc' },
      transactions: {
        pendingBridgeTransactions: [legacyPending, relayDeposit, unrelated],
        userHistoricalTransactions: ['history'],
      },
      _transactions: {
        transactions: [legacyActivity, relayDeposit, synDeposit, unrelated],
      },
    }

    const migrated = (await migratePersistedState(state, 4)) as typeof state

    expect(migrated.transactions.pendingBridgeTransactions).toEqual([
      relayDeposit,
      unrelated,
    ])
    expect(migrated._transactions.transactions).toEqual([
      relayDeposit,
      synDeposit,
      unrelated,
    ])
    expect(migrated.application).toEqual(state.application)
    expect(migrated.transactions.userHistoricalTransactions).toEqual([
      'history',
    ])
    expect(
      await migratePersistedState(
        { ...state, _persist: { version: 4, rehydrated: true } },
        4
      )
    ).toEqual({ ...state, _persist: { version: 4, rehydrated: true } })
  })
})
