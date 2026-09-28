import { renderToStaticMarkup } from 'react-dom/server'

import { _Transactions as Transactions } from '@/components/_Transaction/_Transactions'
import { use_TransactionsState } from '@/slices/_transactions/hooks'
import { SYN, USDC } from '@/constants/tokens/bridgeable'
import transactionsReducer, {
  completeTransaction,
} from '@/slices/_transactions/reducer'
import { ETH, HYPEREVM, HYPERLIQUID } from '@/constants/chains/master'

jest.mock('../slices/_transactions/hooks', () => ({
  use_TransactionsState: jest.fn(),
}))
jest.mock('../slices/wallet/hooks', () => ({
  useWalletState: () => ({ isWalletPending: false }),
}))
jest.mock('../utils/hooks/useIntervalTimer', () => ({
  useIntervalTimer: () => 123,
}))
jest.mock('../components/_Transaction/_Transaction', () => ({
  _Transaction: ({ status, destinationChain }) => (
    <span data-chain={destinationChain.id}>{status}</span>
  ),
}))

describe('HyperCore transaction display', () => {
  it.each([
    [SYN, 'SYN', 'pending', 'pending', 1337],
    [SYN, 'SYN', 'completed', 'completed', 1337],
    [USDC, 'Relay', 'pending', 'pending', 1337],
  ])(
    'uses the tracked status for %s through %s (%s)',
    (token, bridgeModuleName, storedStatus, displayedStatus, storedChainId) => {
      jest.mocked(use_TransactionsState).mockReturnValue({
        transactions: [
          {
            address: '0xabc',
            originChain: { id: 1 },
            destinationChain: { id: storedChainId, chainSymbol: 'HYPERLIQUID' },
            originToken: token,
            destinationToken: token,
            originTxHash: '0xsyn',
            originValue: '1',
            bridgeModuleName,
            timestamp: 123,
            status: storedStatus,
          },
        ],
      } as any)
      expect(
        renderToStaticMarkup(<Transactions connectedAddress="0xabc" />)
      ).toContain(`<span data-chain="1337">${displayedStatus}</span>`)
    }
  )

  it('persists an observed HyperEVM fallback as the completed destination', () => {
    const originTxHash = '0xfallback'
    const state = {
      transactions: [
        {
          originTxHash,
          originChain: ETH,
          originToken: SYN,
          destinationChain: HYPERLIQUID,
          destinationToken: SYN,
          status: 'pending',
        },
      ],
    }

    const nextState = transactionsReducer(
      state,
      completeTransaction({
        originTxHash,
        kappa: originTxHash,
        destinationChain: HYPEREVM,
      })
    )

    expect(nextState.transactions[0]).toMatchObject({
      status: 'completed',
      destinationChain: HYPEREVM,
    })
  })
})
