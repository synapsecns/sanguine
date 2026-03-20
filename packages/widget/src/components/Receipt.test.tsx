import { render, screen } from '@testing-library/react'

import { Receipt } from '@/components/Receipt'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import {
  EMPTY_BRIDGE_QUOTE,
  type BridgeQuote,
} from '@/state/slices/bridgeQuote/reducer'

jest.mock('@/state/slices/bridge/hooks')
jest.mock('@/constants/chains', () => ({
  CHAINS_BY_ID: {
    1: {
      id: 1,
      name: 'Ethereum',
      nativeCurrency: {
        symbol: 'ETH',
        decimals: 18,
      },
    },
    2: {
      id: 2,
      name: 'Arbitrum',
      nativeCurrency: {
        symbol: 'ETH',
        decimals: 18,
      },
    },
    56: {
      id: 56,
      name: 'BNB Chain',
      nativeCurrency: {
        symbol: 'BNB',
        decimals: 18,
      },
    },
  },
}))

const createQuote = (overrides: Partial<BridgeQuote> = {}): BridgeQuote => ({
  ...EMPTY_BRIDGE_QUOTE,
  outputAmount: 1_500_000n,
  outputAmountString: '1.5000',
  delta: 1_500_000n,
  estimatedTime: 45,
  bridgeModuleName: 'SynapseRFQ',
  requestId: 1,
  timestamp: 1,
  ...overrides,
})

const renderReceipt = ({
  quote = createQuote(),
  loading = false,
  send = '1.0000 USDC',
  receive = '1.5000 USDT',
}: {
  quote?: BridgeQuote
  loading?: boolean
  send?: string | undefined
  receive?: string | undefined
} = {}) => {
  return render(
    <Receipt
      quote={quote as any}
      loading={loading}
      send={send}
      receive={receive}
    />
  )
}

const getRowValue = (label: string) => {
  const term = screen.getByText(label)
  const value = term.nextElementSibling

  expect(value).not.toBeNull()

  return value?.textContent
}

describe('Receipt native fee display', () => {
  beforeEach(() => {
    ;(useBridgeState as jest.Mock).mockReturnValue({
      originChainId: 1,
      destinationChainId: 2,
    })
  })

  afterEach(() => {
    jest.clearAllMocks()
  })

  it('renders Bridge fee between Send and Receive for positive native fees', () => {
    const { container } = renderReceipt({
      quote: createQuote({
        nativeFee: 420_000_000_000_000n,
      }),
    })

    expect(getRowValue('Send')).toBe('1.0000 USDC')
    expect(getRowValue('Bridge fee')).toBe('0.00042 ETH')
    expect(getRowValue('Receive')).toBe('1.5000 USDT')
    expect(
      Array.from(container.querySelectorAll('dt')).map(
        (node) => node.textContent
      )
    ).toEqual([
      'Router',
      'Origin',
      'Destination',
      'Send',
      'Bridge fee',
      'Receive',
    ])
  })

  it('keeps the collapsed summary unchanged when a fee is present', () => {
    const { container } = renderReceipt({
      quote: createQuote({
        nativeFee: 420_000_000_000_000n,
      }),
    })

    const summary = container.querySelector('summary')

    expect(summary).toHaveTextContent('45 seconds via SynapseRFQ')
    expect(summary).not.toHaveTextContent('Bridge fee')
    expect(summary).not.toHaveTextContent('0.00042 ETH')
  })

  it.each([
    [
      'the quote is loading',
      createQuote({ nativeFee: 420_000_000_000_000n }),
      true,
    ],
    [
      'the quote is invalid',
      createQuote({ outputAmount: 0n, nativeFee: 420_000_000_000_000n }),
      false,
    ],
    ['nativeFee is zero', createQuote({ nativeFee: 0n }), false],
    [
      'nativeFee is missing',
      createQuote({ nativeFee: undefined as any }),
      false,
    ],
    ['nativeFee is malformed', createQuote({ nativeFee: 'abc' as any }), false],
  ])('hides Bridge fee when %s', (_scenario, quote, loading) => {
    renderReceipt({ quote, loading })

    expect(screen.queryByText('Bridge fee')).not.toBeInTheDocument()
  })

  it.each([
    ['1 wei', 1n, '0.000000000000000001 ETH'],
    ['a tiny non-zero fee', 420_000_000_000_000n, '0.00042 ETH'],
    ['an integer fee', 1_000_000_000_000_000_000n, '1 ETH'],
    ['a fee with trailing zeros', 1_234_000_000_000_000_000n, '1.234 ETH'],
  ])(
    'formats %s without collapsing it to zero',
    (_scenario, nativeFee, expected) => {
      renderReceipt({
        quote: createQuote({
          nativeFee,
        }),
      })

      expect(getRowValue('Bridge fee')).toBe(expected)
      expect(getRowValue('Bridge fee')).not.toMatch(/^0(?:\.0+)? ETH$/)
    }
  )
})
