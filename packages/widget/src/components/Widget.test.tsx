import { act, fireEvent, render, screen } from '@testing-library/react'

import { useValidations } from '@/hooks/useValidations'
import { useThemeVariables } from '@/hooks/useThemeVariables'
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
jest.mock('@/state/slices/bridgeTransaction/hooks')
jest.mock('@/state/slices/approveTransaction/hooks')
jest.mock('@/components/Receipt', () => ({ Receipt: () => null }))
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
    addresses: { 1: '0x1', 2: '0x2' },
    decimals: { 1: 6, 2: 6 },
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
    1: { id: 1, name: 'Ethereum' },
    2: { id: 2, name: 'Arbitrum' },
  },
}))

const dispatchMock = jest.fn()
const fetchBridgeQuoteMock = fetchBridgeQuote as unknown as jest.Mock
const fetchAndStoreTokenBalancesMock =
  fetchAndStoreTokenBalances as unknown as jest.Mock
const fetchAndStoreAllowanceMock =
  fetchAndStoreAllowance as unknown as jest.Mock

const token = {
  addresses: { 1: '0x1', 2: '0x2' },
  decimals: { 1: 6, 2: 6 },
  symbol: 'USDC',
  name: 'USD Coin',
  swapableType: 'token',
  color: '#fff',
  priorityRank: 1,
  routeSymbol: 'USDC',
  imgUrl: '',
}

let mockBridgeState: any
let mockQuoteState: any
let mockWalletState: any
let mockValidationState: any

const createQuote = (overrides: Partial<BridgeQuote> = {}): BridgeQuote => ({
  ...EMPTY_BRIDGE_QUOTE,
  outputAmount: 1n,
  outputAmountString: '1.00',
  delta: 1n,
  requestId: 1,
  timestamp: 1,
  ...overrides,
})

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

const dispatchMouseMove = () => {
  act(() => {
    document.dispatchEvent(new MouseEvent('mousemove', { bubbles: true }))
  })
}

describe('Widget quote refresh wiring', () => {
  beforeEach(() => {
    jest.useFakeTimers()

    mockBridgeState = {
      debouncedInputAmount: '',
      originChainId: 1,
      originToken: token,
      destinationChainId: 2,
      destinationToken: token,
    }
    mockQuoteState = {
      bridgeQuote: EMPTY_BRIDGE_QUOTE,
      isLoading: false,
    }
    mockWalletState = {
      isWalletPending: false,
    }
    mockValidationState = {
      hasValidSelections: true,
      isInputValid: false,
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
    dispatchMock.mockImplementation((action) => {
      if (action?.type === 'bridge/setDebouncedInputAmount') {
        mockBridgeState = {
          ...mockBridgeState,
          debouncedInputAmount: action.payload,
        }
        mockValidationState = {
          ...mockValidationState,
          isInputValid: action.payload !== '',
        }
      }

      return action
    })
  })

  afterEach(() => {
    jest.clearAllTimers()
    jest.useRealTimers()
    jest.restoreAllMocks()
    jest.clearAllMocks()
  })

  it('keeps normal input-driven fetches and only refreshes stale quotes on mousemove', () => {
    const view = renderWidget()

    expect(fetchBridgeQuote).not.toHaveBeenCalled()

    act(() => {
      fireEvent.change(screen.getByPlaceholderText('0'), {
        target: { value: '1' },
      })
      jest.advanceTimersByTime(300)
    })

    rerenderWidget(view)

    expect(fetchBridgeQuote).toHaveBeenCalledTimes(1)
    expect(fetchBridgeQuoteMock.mock.calls[0][0]).toEqual(
      expect.objectContaining({
        debouncedInputAmount: '1',
        requestId: 1,
      })
    )

    mockQuoteState = {
      bridgeQuote: createQuote({ requestId: 1, timestamp: 1 }),
      isLoading: false,
    }

    rerenderWidget(view)

    act(() => {
      document.dispatchEvent(new MouseEvent('mousemove', { bubbles: true }))
    })
    expect(fetchBridgeQuote).toHaveBeenCalledTimes(1)

    act(() => {
      jest.advanceTimersByTime(15000)
    })
    act(() => {
      document.dispatchEvent(new MouseEvent('mousemove', { bubbles: true }))
    })

    expect(fetchBridgeQuote).toHaveBeenCalledTimes(2)
    expect(fetchBridgeQuoteMock.mock.calls[1][0]).toEqual(
      expect.objectContaining({
        debouncedInputAmount: '1',
        requestId: 2,
      })
    )

    act(() => {
      document.dispatchEvent(new MouseEvent('mousemove', { bubbles: true }))
    })
    expect(fetchBridgeQuote).toHaveBeenCalledTimes(2)
  })

  it.each([
    {
      busyState: 'quote loading',
      startBusy: () => {
        mockQuoteState = {
          ...mockQuoteState,
          isLoading: true,
        }
      },
      endBusy: () => {
        mockQuoteState = {
          ...mockQuoteState,
          isLoading: false,
        }
      },
    },
    {
      busyState: 'wallet pending',
      startBusy: () => {
        mockWalletState = {
          ...mockWalletState,
          isWalletPending: true,
        }
      },
      endBusy: () => {
        mockWalletState = {
          ...mockWalletState,
          isWalletPending: false,
        }
      },
    },
  ])(
    'does not allow immediate stale refreshes after $busyState resumes',
    ({ startBusy, endBusy }) => {
      const view = renderWidget()

      act(() => {
        fireEvent.change(screen.getByPlaceholderText('0'), {
          target: { value: '1' },
        })
        jest.advanceTimersByTime(300)
      })

      rerenderWidget(view)

      expect(fetchBridgeQuote).toHaveBeenCalledTimes(1)

      mockQuoteState = {
        bridgeQuote: createQuote({ requestId: 1, timestamp: 1 }),
        isLoading: false,
      }

      rerenderWidget(view)

      act(() => {
        jest.advanceTimersByTime(15000)
      })

      startBusy()
      rerenderWidget(view)

      dispatchMouseMove()
      expect(fetchBridgeQuote).toHaveBeenCalledTimes(1)

      endBusy()
      rerenderWidget(view)

      dispatchMouseMove()
      expect(fetchBridgeQuote).toHaveBeenCalledTimes(1)

      act(() => {
        jest.advanceTimersByTime(14999)
      })

      dispatchMouseMove()
      expect(fetchBridgeQuote).toHaveBeenCalledTimes(1)

      act(() => {
        jest.advanceTimersByTime(1)
      })

      dispatchMouseMove()
      expect(fetchBridgeQuote).toHaveBeenCalledTimes(2)
      expect(fetchBridgeQuoteMock.mock.calls[1][0]).toEqual(
        expect.objectContaining({
          debouncedInputAmount: '1',
          requestId: 2,
        })
      )
    }
  )
})
