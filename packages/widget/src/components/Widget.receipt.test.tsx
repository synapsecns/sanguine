import { act, render, screen } from '@testing-library/react'

import { useValidations } from '@/hooks/useValidations'
import { useThemeVariables } from '@/hooks/useThemeVariables'
import { useBridgeQuoteUpdater } from '@/hooks/useBridgeQuoteUpdater'
import { useAppDispatch } from '@/state/hooks'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import {
  fetchBridgeQuote,
  useBridgeQuoteState,
} from '@/state/slices/bridgeQuote/hooks'
import {
  EMPTY_BRIDGE_QUOTE,
  type BridgeQuote,
} from '@/state/slices/bridgeQuote/reducer'
import {
  fetchAndStoreAllowance,
  fetchAndStoreTokenBalances,
  useWalletState,
} from '@/state/slices/wallet/hooks'
import { useSynapseContext } from '@/providers/SynapseProvider'
import { Web3Context } from '@/providers/Web3Provider'
import { useMaintenance } from '@/components/Maintenance/Maintenance'
import { useApproveTransactionState } from '@/state/slices/approveTransaction/hooks'
import { useBridgeTransactionState } from '@/state/slices/bridgeTransaction/hooks'
import { Widget } from '@/components/Widget'

jest.mock('ethers', () => ({
  ZeroAddress: '0x0000000000000000000000000000000000000000',
}))
jest.mock('@/state/hooks')
jest.mock('@/state/slices/bridge/hooks')
jest.mock('@/state/slices/bridgeQuote/hooks')
jest.mock('@/state/slices/wallet/hooks')
jest.mock('@/hooks/useValidations')
jest.mock('@/providers/SynapseProvider')
jest.mock('@/components/Maintenance/Maintenance')
jest.mock('@/hooks/useThemeVariables')
jest.mock('@/hooks/useBridgeQuoteUpdater')
jest.mock('@/state/slices/bridgeTransaction/hooks')
jest.mock('@/state/slices/approveTransaction/hooks')
jest.mock('@/components/BridgeButton', () => ({ BridgeButton: () => null }))
jest.mock('@/components/AvailableBalance', () => ({
  AvailableBalance: () => null,
}))
jest.mock('@/components/Transactions', () => ({ Transactions: () => null }))
jest.mock('@/components/ui/ChainSelect', () => ({ ChainSelect: () => null }))
jest.mock('@/components/ui/TokenSelect', () => ({ TokenSelect: () => null }))
jest.mock('@/components/ui/SwitchButton', () => ({ SwitchButton: () => null }))
jest.mock('@/utils/routeMaker/getFromTokens', () => ({
  getFromTokens: () => [{ routeSymbol: 'USDC' }],
}))
jest.mock('@/utils/routeMaker/generateRoutePossibilities', () => ({
  getSymbol: (routeToken: { routeSymbol: string }) => routeToken.routeSymbol,
}))
jest.mock('@/utils/findTokenByRouteSymbol', () => ({
  findTokenByRouteSymbol: () => ({
    addresses: { 1: '0x1', 2: '0x2', 56: '0x56' },
    decimals: { 1: 6, 2: 6, 56: 6 },
    symbol: 'USDC',
    name: 'USD Coin',
    swapableType: 'token',
    color: '#fff',
    priorityRank: 1,
    routeSymbol: 'USDC',
    imgUrl: '',
  }),
}))
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

const dispatchMock = jest.fn()
const fetchBridgeQuoteMock = fetchBridgeQuote as unknown as jest.Mock
const fetchAndStoreTokenBalancesMock =
  fetchAndStoreTokenBalances as unknown as jest.Mock
const fetchAndStoreAllowanceMock =
  fetchAndStoreAllowance as unknown as jest.Mock

const createToken = (symbol?: string) => ({
  addresses: { 1: '0x1', 2: '0x2', 56: '0x56' },
  decimals: { 1: 6, 2: 6, 56: 6 },
  symbol,
  name: symbol ?? 'Token',
  swapableType: 'token',
  color: '#fff',
  priorityRank: 1,
  routeSymbol: symbol ?? 'TOKEN',
  imgUrl: '',
})

const createQuote = (overrides: Partial<BridgeQuote> = {}): BridgeQuote => ({
  ...EMPTY_BRIDGE_QUOTE,
  outputAmount: 1_500_000n,
  outputAmountString: '1.5000',
  delta: 1_500_000n,
  nativeFee: 420_000_000_000_000n,
  bridgeModuleName: 'SynapseRFQ',
  estimatedTime: 45,
  requestId: 1,
  timestamp: 1,
  ...overrides,
})

let mockBridgeState: any
let mockQuoteState: any
let mockWalletState: any
let mockValidationState: any

const renderWidget = () => {
  return render(
    <Web3Context.Provider
      value={{
        web3Provider: {
          connectedAddress: '',
          networkId: 1,
          signer: null,
          provider: null,
        },
        setWeb3Provider: jest.fn() as any,
      }}
    >
      <Widget customTheme={{} as any} />
    </Web3Context.Provider>
  )
}

const rerenderWidget = (view: ReturnType<typeof renderWidget>) => {
  act(() => {
    view.rerender(
      <Web3Context.Provider
        value={{
          web3Provider: {
            connectedAddress: '',
            networkId: 1,
            signer: null,
            provider: null,
          },
          setWeb3Provider: jest.fn() as any,
        }}
      >
        <Widget customTheme={{} as any} />
      </Web3Context.Provider>
    )
  })
}

const getRowValue = (label: string) => {
  const term = screen.getByText(label)
  const value = term.nextElementSibling

  expect(value).not.toBeNull()

  return value?.textContent
}

describe('Widget receipt integration', () => {
  beforeEach(() => {
    mockBridgeState = {
      debouncedInputAmount: '1',
      originChainId: 1,
      originToken: createToken('USDC'),
      destinationChainId: 2,
      destinationToken: createToken('USDT'),
    }
    mockQuoteState = {
      bridgeQuote: createQuote(),
      isLoading: false,
    }
    mockWalletState = {
      isWalletPending: false,
    }
    mockValidationState = {
      hasValidSelections: true,
      isInputValid: true,
    }
    ;(useAppDispatch as jest.Mock).mockReturnValue(dispatchMock)
    ;(useBridgeState as jest.Mock).mockImplementation(() => mockBridgeState)
    ;(useBridgeQuoteState as jest.Mock).mockImplementation(() => mockQuoteState)
    ;(useWalletState as jest.Mock).mockImplementation(() => mockWalletState)
    ;(useValidations as jest.Mock).mockImplementation(() => ({
      hasValidSelections: mockValidationState.hasValidSelections,
      isInputValid: mockValidationState.isInputValid,
    }))
    ;(useSynapseContext as jest.Mock).mockReturnValue({
      synapseSDK: {},
      synapseProviders: [],
    })
    ;(useMaintenance as jest.Mock).mockReturnValue({
      isBridgePaused: false,
      pausedModulesList: [],
      BridgeMaintenanceProgressBar: () => null,
      BridgeMaintenanceWarningMessage: () => null,
    })
    ;(useThemeVariables as jest.Mock).mockReturnValue({})
    ;(useBridgeQuoteUpdater as jest.Mock).mockImplementation(() => undefined)
    ;(useBridgeTransactionState as jest.Mock).mockReturnValue({
      bridgeTxnStatus: null,
    })
    ;(useApproveTransactionState as jest.Mock).mockReturnValue({
      approveTxnStatus: null,
    })

    fetchBridgeQuoteMock.mockImplementation((args) => ({
      type: 'bridgeQuote/fetchBridgeQuote',
      meta: { arg: args },
    }))
    fetchAndStoreTokenBalancesMock.mockImplementation((args) => ({
      type: 'wallet/fetchAndStoreTokenBalances',
      meta: { arg: args },
    }))
    fetchAndStoreAllowanceMock.mockImplementation((args) => ({
      type: 'wallet/fetchAndStoreAllowance',
      meta: { arg: args },
    }))

    dispatchMock.mockReset()
    dispatchMock.mockImplementation((action) => action)
  })

  afterEach(() => {
    jest.clearAllMocks()
  })

  it('appends the selected token symbols to the Send and Receive rows', () => {
    renderWidget()

    expect(getRowValue('Send')).toBe('1.0000 USDC')
    expect(getRowValue('Receive')).toBe('1.5000 USDT')
  })

  it('falls back to amount-only Send and Receive rows when token symbols are unavailable', () => {
    mockBridgeState = {
      ...mockBridgeState,
      originToken: createToken(undefined),
      destinationToken: createToken(undefined),
    }

    renderWidget()

    expect(getRowValue('Send')).toBe('1.0000')
    expect(getRowValue('Receive')).toBe('1.5000')
    expect(screen.queryByText(/undefined/)).not.toBeInTheDocument()
  })

  it('does not relabel a stale quote with newly selected token symbols before a new quote arrives', () => {
    const view = renderWidget()

    expect(getRowValue('Send')).toBe('1.0000 USDC')
    expect(getRowValue('Receive')).toBe('1.5000 USDT')

    mockBridgeState = {
      ...mockBridgeState,
      originToken: createToken('DAI'),
      destinationToken: createToken('FRAX'),
    }

    rerenderWidget(view)

    expect(screen.queryByText('1.0000 DAI')).not.toBeInTheDocument()
    expect(screen.queryByText('1.5000 FRAX')).not.toBeInTheDocument()
  })

  it('does not relabel a stale bridge fee with the new origin native symbol before a new quote arrives', () => {
    const view = renderWidget()

    expect(getRowValue('Bridge fee')).toBe('0.00042 ETH')

    mockBridgeState = {
      ...mockBridgeState,
      originChainId: 56,
      originToken: createToken('USDC'),
    }

    rerenderWidget(view)

    expect(screen.queryByText('0.00042 BNB')).not.toBeInTheDocument()
  })

  it('does not keep stale receive and bridge fee values after only the amount changes', () => {
    const view = renderWidget()

    expect(getRowValue('Send')).toBe('1.0000 USDC')
    expect(getRowValue('Bridge fee')).toBe('0.00042 ETH')
    expect(getRowValue('Receive')).toBe('1.5000 USDT')

    mockBridgeState = {
      ...mockBridgeState,
      debouncedInputAmount: '2',
    }

    rerenderWidget(view)

    expect(getRowValue('Send')).toBe('-')
    expect(screen.queryByText('Bridge fee')).not.toBeInTheDocument()
    expect(getRowValue('Receive')).toBe('-')
    expect(screen.queryByText('2.0000 USDC')).not.toBeInTheDocument()
    expect(screen.queryByText('1.5000 USDT')).not.toBeInTheDocument()
    expect(screen.queryByText('0.00042 ETH')).not.toBeInTheDocument()
  })
})
