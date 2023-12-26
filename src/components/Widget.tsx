import {
  useMemo,
  useEffect,
  useContext,
  useCallback,
  useRef,
  useState,
} from 'react'
import { Web3Context } from '@/providers/Web3Provider'

import { formatBigIntToString } from '@/utils/formatBigIntToString'
import { stringToBigInt } from '@/utils/stringToBigInt'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import { Receipt } from '@/components/Receipt'

import { BridgeableToken, Chain, WidgetProps } from 'types'
import { ChainSelect } from '@/components/ui/ChainSelect'
import { TokenSelect } from '@/components/ui/TokenSelect'

import { useAppDispatch } from '@/state/hooks'
import {
  setDestinationChainId,
  setOriginChainId,
  setOriginToken,
  setDestinationToken,
  setTokens,
  setDebouncedInputAmount,
} from '@/state/slices/bridge/reducer'
import { useBridgeState } from '@/state/slices/bridge/hooks'
import {
  fetchAndStoreAllowance,
  fetchAndStoreTokenBalances,
} from '@/state/slices/wallet/hooks'
import { BridgeButton } from './BridgeButton'

import { AvailableBalance } from './AvailableBalance'
import { ZeroAddress } from 'ethers'
import { useValidations } from '@/hooks/useValidations'

import {
  fetchBridgeQuote,
  useBridgeQuoteState,
} from '@/state/slices/bridgeQuote/hooks'

import {
  EMPTY_BRIDGE_QUOTE,
  resetQuote,
} from '@/state/slices/bridgeQuote/reducer'

import {
  executeBridgeTxn,
  useBridgeTransactionState,
} from '@/state/slices/bridgeTransaction/hooks'
import { BridgeTransactionStatus } from '@/state/slices/bridgeTransaction/reducer'
import {
  executeApproveTxn,
  useApproveTransactionState,
} from '@/state/slices/approveTransaction/hooks'
import { ApproveTransactionStatus } from '@/state/slices/approveTransaction/reducer'
import { useThemeVariables } from '@/hooks/useThemeVariables'

import { Transactions } from './Transactions'

import { CHAINS_BY_ID } from '@/constants/chains'
import { useSynapseContext } from '@/providers/SynapseProvider'

export const Widget = ({
  chainIds,
  networkProviders,
  theme,
  customTheme,
  container,
  tokens,
  toChainId,
}) => {
  const dispatch = useAppDispatch()
  const currentSDKRequestID = useRef(0)

  const { synapseSDK } = useSynapseContext()

  const web3Context = useContext(Web3Context)
  const { connectedAddress, signer, provider, networkId } =
    web3Context.web3Provider

  const [inputAmount, setInputAmount] = useState('')

  const {
    debouncedInputAmount,
    originChainId,
    originToken,
    destinationChainId,
    destinationToken,
    tokens: allTokens,
  } = useBridgeState()

  const { bridgeQuote, isLoading } = useBridgeQuoteState()

  const { isInputValid, hasValidSelections } = useValidations()

  const { bridgeTxnStatus } = useBridgeTransactionState()
  const { approveTxnStatus } = useApproveTransactionState()

  const themeVariables = useThemeVariables(theme, customTheme)

  /** Select Consumer networkProvider based on Origin ChainId */
  /** TODO: Move to helper hook */
  const originChainProvider = useMemo(() => {
    if (!Array.isArray(networkProviders)) return null
    if (!originChainId) return null
    const _provider = networkProviders.find(
      (provider) => Number(provider?._network?.chainId) === originChainId
    )
    return _provider
  }, [originChainId, networkProviders])

  useEffect(() => {
    dispatch(setTokens(tokens))
    dispatch(setDestinationChainId(toChainId))
    dispatch(setDestinationToken(tokens[0]))
  }, [tokens, toChainId])

  /** Debounce user input to fetch bridge quote (in ms) */
  /** TODO: Can this be moved to the input component? */
  useEffect(() => {
    const DEBOUNCE_DELAY = 300
    const debounceTimer = setTimeout(() => {
      dispatch(setDebouncedInputAmount(inputAmount))
    }, DEBOUNCE_DELAY)

    return () => {
      clearTimeout(debounceTimer)
    }
  }, [dispatch, inputAmount])

  /** Fetch token balances when signer/address connected */
  /** TODO: Can this be moved into a level above? */
  useEffect(() => {
    if (!signer && !originChainProvider) return
    if (originChainId && tokens && connectedAddress) {
      dispatch(
        fetchAndStoreTokenBalances({
          address: connectedAddress,
          chainId: originChainId,
          tokens: allTokens,
          signerOrProvider: originChainProvider ?? signer,
        })
      )
    }
  }, [originChainId, allTokens, connectedAddress, signer, originChainProvider])

  const executeApproval = () => {
    dispatch(
      executeApproveTxn({
        spenderAddress: bridgeQuote?.routerAddress,
        tokenAddress: originToken?.addresses[originChainId],
        amount: stringToBigInt(
          debouncedInputAmount,
          originToken?.decimals[originChainId]
        ),
        signer,
      })
    )
  }

  const executeBridge = () => {
    dispatch(
      executeBridgeTxn({
        destinationAddress: connectedAddress,
        originRouterAddress: bridgeQuote?.routerAddress,
        originChainId: originChainId,
        destinationChainId: destinationChainId,
        tokenAddress: originToken?.addresses[originChainId],
        amount: stringToBigInt(
          debouncedInputAmount,
          originToken?.decimals[originChainId]
        ),
        originQuery: bridgeQuote?.quotes.originQuery,
        destinationQuery: bridgeQuote?.quotes.destQuery,
        bridgeModuleName: bridgeQuote?.bridgeModuleName,
        estimatedTime: bridgeQuote?.estimatedTime,
        synapseSDK,
        signer,
      })
    )
  }

  /** fetch and store token allowance */

  useEffect(() => {
    if (
      originToken?.addresses[originChainId] !== ZeroAddress &&
      bridgeQuote?.routerAddress
    ) {
      dispatch(
        fetchAndStoreAllowance({
          spenderAddress: bridgeQuote?.routerAddress,
          ownerAddress: connectedAddress,
          chainId: originChainId,
          token: originToken,
          provider: originChainProvider ?? provider,
        })
      )
    }
  }, [
    originToken?.routeSymbol,
    originChainId,
    connectedAddress,
    bridgeQuote?.routerAddress,
  ])

  /** Handle refreshing quotes */
  useEffect(() => {
    if (isInputValid && hasValidSelections) {
      currentSDKRequestID.current += 1
      const thisRequestId = currentSDKRequestID.current

      if (thisRequestId === currentSDKRequestID.current) {
        dispatch(
          fetchBridgeQuote({
            originChainId,
            destinationChainId,
            originToken,
            destinationToken,
            amount: stringToBigInt(
              debouncedInputAmount,
              originToken.decimals[originChainId]
            ),
            debouncedInputAmount,
            synapseSDK,
          })
        )
      }
    } else {
      dispatch(resetQuote())
    }
  }, [
    debouncedInputAmount,
    originToken?.routeSymbol,
    destinationToken?.routeSymbol,
    originChainId,
    destinationChainId,
    isInputValid,
    hasValidSelections,
  ])

  const handleUserInput = useCallback(
    (event: React.ChangeEvent<HTMLInputElement>) => {
      const value = cleanNumberInput(event.target.value)
      setInputAmount(value)
    },
    []
  )

  const handleOriginChainSelection = useCallback(
    (newOriginChain: Chain) => {
      dispatch(setOriginChainId(newOriginChain.id))
    },
    [dispatch]
  )

  const handleDestinationChainSelection = useCallback(
    (newDestinationChain: Chain) => {
      dispatch(setDestinationChainId(newDestinationChain.id))
    },
    [dispatch]
  )

  const handleOriginTokenSelection = useCallback(
    (newOriginToken: BridgeableToken) => {
      dispatch(setOriginToken(newOriginToken))
    },
    [dispatch]
  )

  const handleDestinationTokenSelection = useCallback(
    (newDestinationToken: BridgeableToken) => {
      dispatch(setDestinationToken(newDestinationToken))
    },
    [dispatch]
  )

  const containerStyle = `
    ${container === false ? 'p-0' : 'p-2 rounded-lg'}`

  const cardStyle = `
    grid grid-cols-[1fr_auto] 
    bg-[--synapse-surface] rounded-md p-2 gap-1
    border border-solid border-[--synapse-border]
  `

  const inputStyle = `
    text-3xl w-full font-semibold bg-transparent border-none py-1
    text-[--synapse-text] placeholder:text-[--synapse-secondary] focus:outline-none disabled:cursor-not-allowed font-sans
  `

  return (
    <div
      style={themeVariables}
      className={`synapse-widget ${container && 'max-w-400px'}`}
    >
      <div
        className={`
        bg-[--synapse-root] grid gap-2 text-[--synapse-text] w-full ${containerStyle}
        `}
      >
        <Transactions
          synapseSDK={synapseSDK}
          connectedAddress={connectedAddress}
        />
        <section className={cardStyle}>
          <ChainSelect
            label="From"
            isOrigin={true}
            chain={CHAINS_BY_ID[originChainId]}
            onChange={handleOriginChainSelection}
          />
          <input
            className={inputStyle}
            placeholder="0"
            value={inputAmount}
            onChange={handleUserInput}
          />
          <div className="flex flex-col items-end justify-center">
            <TokenSelect
              label="In"
              isOrigin={true}
              token={originToken}
              onChange={handleOriginTokenSelection}
            />
            <AvailableBalance
              connectedAddress={connectedAddress}
              setInputAmount={setInputAmount}
            />
          </div>
        </section>
        <section className={cardStyle}>
          <ChainSelect
            label="To"
            isOrigin={false}
            chain={CHAINS_BY_ID[destinationChainId]}
            onChange={handleDestinationChainSelection}
          />
          <input
            className={inputStyle}
            disabled={true}
            placeholder=""
            value={
              isLoading || !bridgeQuote || !destinationToken
                ? '...'
                : formatBigIntToString(
                    bridgeQuote?.delta,
                    destinationToken?.decimals[destinationChainId],
                    4
                  ) || ''
            }
          />
          <div className="flex flex-col items-end justify-center">
            <TokenSelect
              label="Out"
              isOrigin={false}
              token={destinationToken}
              onChange={handleDestinationTokenSelection}
            />
          </div>
        </section>
        <Receipt
          quote={bridgeQuote ?? null}
          send={formatBigIntToString(
            stringToBigInt(
              debouncedInputAmount,
              originToken?.decimals[originChainId]
            ),
            originToken?.decimals[originChainId],
            4
          )}
          receive={formatBigIntToString(
            bridgeQuote?.delta,
            destinationToken?.decimals[destinationChainId],
            4
          )}
        />
        <BridgeButton
          originChain={CHAINS_BY_ID[originChainId]}
          isValidQuote={
            Boolean(bridgeQuote) && bridgeQuote !== EMPTY_BRIDGE_QUOTE
          }
          handleApprove={executeApproval}
          handleBridge={executeBridge}
          isApprovalPending={
            approveTxnStatus === ApproveTransactionStatus.PENDING
          }
          isBridgePending={bridgeTxnStatus === BridgeTransactionStatus.PENDING}
        />
      </div>
    </div>
  )
}
