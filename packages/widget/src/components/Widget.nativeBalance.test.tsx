import { fireEvent, render, screen, waitFor } from '@testing-library/react'

import { useCurrentTokenBalance } from '@/hooks/useCurrentTokenBalance'
import { useThemeVariables } from '@/hooks/useThemeVariables'
import { useAppDispatch } from '@/state/hooks'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import { useBridgeQuoteState } from '@/state/slices/bridgeQuote/hooks'
import { EMPTY_BRIDGE_QUOTE } from '@/state/slices/bridgeQuote/reducer'
import { FetchState } from '@/state/slices/wallet/reducer'
import { useWalletState } from '@/state/slices/wallet/hooks'
import { useValidations } from '@/hooks/useValidations'
import { useSynapseContext } from '@/providers/SynapseProvider'
import { Web3Context } from '@/providers/Web3Provider'
import { useMaintenance } from '@/components/Maintenance/Maintenance'
import { useBridgeQuoteUpdater } from '@/hooks/useBridgeQuoteUpdater'
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
jest.mock('@/hooks/useCurrentTokenBalance')
jest.mock('@/hooks/useValidations')
jest.mock('@/providers/SynapseProvider')
jest.mock('@/components/Maintenance/Maintenance')
jest.mock('@/hooks/useThemeVariables')
jest.mock('@/hooks/useBridgeQuoteUpdater')
jest.mock('@/state/slices/bridgeTransaction/hooks')
jest.mock('@/state/slices/approveTransaction/hooks')
jest.mock('@/components/Receipt', () => ({ Receipt: () => null }))
jest.mock('@/components/BridgeButton', () => ({ BridgeButton: () => null }))
jest.mock('@/components/Transactions', () => ({ Transactions: () => null }))
jest.mock('@/components/ui/ChainSelect', () => ({ ChainSelect: () => null }))
jest.mock('@/components/ui/TokenSelect', () => ({ TokenSelect: () => null }))
jest.mock('@/components/ui/SwitchButton', () => ({ SwitchButton: () => null }))
jest.mock('@/utils/routeMaker/getFromTokens', () => ({
  getFromTokens: () => [{ routeSymbol: 'ETH' }],
}))
jest.mock('@/utils/routeMaker/generateRoutePossibilities', () => ({
  getSymbol: (routeToken: { routeSymbol: string }) => routeToken.routeSymbol,
}))
jest.mock('@/utils/findTokenByRouteSymbol', () => ({
  findTokenByRouteSymbol: () => ({
    addresses: { 1: '0x0000000000000000000000000000000000000000' },
    decimals: { 1: 18 },
    symbol: 'ETH',
    name: 'Ether',
    swapableType: 'token',
    color: '#fff',
    priorityRank: 1,
    routeSymbol: 'ETH',
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
  },
}))

const dispatchMock = jest.fn()
let mockBridgeV2: jest.Mock
let mockOriginChainProvider: {
  _network: { chainId: number }
  estimateGas: jest.Mock
  getFeeData: jest.Mock
}

const nativeToken = {
  addresses: { 1: '0x0000000000000000000000000000000000000000' },
  decimals: { 1: 18 },
  symbol: 'ETH',
  name: 'Ether',
  swapableType: 'token',
  color: '#fff',
  priorityRank: 1,
  routeSymbol: 'ETH',
  imgUrl: '',
}

const destinationToken = {
  addresses: { 2: '0x2' },
  decimals: { 2: 18 },
  symbol: 'nETH',
  name: 'nETH',
  swapableType: 'token',
  color: '#fff',
  priorityRank: 1,
  routeSymbol: 'NETH',
  imgUrl: '',
}

const createQuote = (nativeFee: string) => ({
  moduleNames: ['SynapseRFQ'],
  nativeFee,
  tx: {
    data: '0x1234',
    to: '0xrouter',
    value: '0',
  },
})

describe('Widget native balance control integration', () => {
  beforeEach(() => {
    mockOriginChainProvider = {
      _network: { chainId: 1 },
      estimateGas: jest
        .fn()
        .mockResolvedValueOnce(100_000n)
        .mockResolvedValueOnce(80_000n),
      getFeeData: jest.fn().mockResolvedValue({
        gasPrice: 10_000_000_000n,
        maxFeePerGas: 1_000_000_000n,
      }),
    }
    mockBridgeV2 = jest
      .fn()
      .mockResolvedValueOnce([
        createQuote('100000000000000000'),
        {
          ...createQuote('1'),
          moduleNames: ['SynapseBridge'],
        },
      ])
      .mockResolvedValueOnce([createQuote('200000000000000000')])
    ;(useAppDispatch as jest.Mock).mockReturnValue(dispatchMock)
    ;(useBridgeState as jest.Mock).mockReturnValue({
      debouncedInputAmount: '',
      destinationChainId: 2,
      destinationToken,
      originChainId: 1,
      originToken: nativeToken,
    })
    ;(useBridgeQuoteState as jest.Mock).mockReturnValue({
      bridgeQuote: EMPTY_BRIDGE_QUOTE,
      isLoading: false,
    })
    ;(useWalletState as jest.Mock).mockReturnValue({
      balancesFetchStatus: FetchState.VALID,
      isWalletPending: false,
    })
    ;(useCurrentTokenBalance as jest.Mock).mockReturnValue({
      rawBalance: '1000000000000000000',
      parsedBalance: '1.0000',
      decimals: 18,
    })
    ;(useValidations as jest.Mock).mockReturnValue({
      hasEnoughBalance: true,
      hasValidSelections: true,
      isInputValid: false,
    })
    ;(useSynapseContext as jest.Mock).mockReturnValue({
      synapseSDK: { bridgeV2: mockBridgeV2 },
      synapseProviders: [mockOriginChainProvider],
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

    dispatchMock.mockReset()
    dispatchMock.mockImplementation((action) => action)
  })

  afterEach(() => {
    jest.clearAllMocks()
  })

  it('renders the real native bridgeable balance and fills the refined safe max on click', async () => {
    render(
      <Web3Context.Provider
        value={{
          web3Provider: {
            connectedAddress: '0xabc',
            networkId: 1,
            provider: null,
            signer: null,
          },
          setWeb3Provider: jest.fn() as any,
        }}
      >
        <Widget customTheme={{} as any} />
      </Web3Context.Provider>
    )

    expect(screen.getByText('Available 1.0000')).toBeInTheDocument()

    await waitFor(() => {
      expect(screen.getByText('Bridgeable 0.7998')).toBeInTheDocument()
    })

    fireEvent.click(screen.getByText('Bridgeable 0.7998'))

    expect(screen.getByDisplayValue('0.799880000000000000')).toBeInTheDocument()
    expect(mockBridgeV2).toHaveBeenCalledTimes(2)
    expect(mockBridgeV2).toHaveBeenNthCalledWith(
      1,
      expect.objectContaining({
        fromAmount: '500000000000000000',
        fromSender: '0xabc',
        toRecipient: '0xabc',
      })
    )
    expect(mockOriginChainProvider.getFeeData).toHaveBeenCalledTimes(2)
    expect(mockOriginChainProvider.estimateGas).toHaveBeenNthCalledWith(
      1,
      expect.objectContaining({
        data: '0x1234',
        from: '0xabc',
        to: '0xrouter',
      })
    )
  })
})
