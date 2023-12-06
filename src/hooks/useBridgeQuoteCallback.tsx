import { useEffect, useState, useMemo } from 'react'
import { fetchBridgeQuote } from '@/utils/actions/fetchBridgeQuote'

export enum QuoteCallbackState {
  IDLE = 'idle',
  INVALID = 'invalid',
  LOADING = 'loading',
  VALID = 'valid',
}

export enum QuoteCallbackErrors {
  REQUIRE_ORIGIN_CHAIN = 'Missing Origin Chain',
  REQUIRE_ORIGIN_TOKEN = 'Missing Origin Token',
  REQUIRE_DEST_CHAIN = 'Missing Destination Chain',
  REQUIRE_DEST_TOKEN = 'Missing Destination Token',
  REQUIRE_AMOUNT = 'Missing Input Amount',
}

interface UseBridgeQuoteCallbackArgs {
  originChainId?: number
  originTokenAddress?: string
  destinationChainId?: number
  destinationTokenAddress?: string
  amount?: bigint
  synapseSDK?: any
}

export function useBridgeQuoteCallback({
  originChainId,
  destinationChainId,
  originTokenAddress,
  destinationTokenAddress,
  amount,
  synapseSDK,
}: UseBridgeQuoteCallbackArgs): {
  state: QuoteCallbackState
  callback: () => Promise<void>
  quote: any
  error: any
} {
  const [quote, setQuote] = useState<any>(null)
  const [error, setError] = useState<any>(null)
  const [loading, setLoading] = useState<boolean>(false)

  const FetchStatusCallback: {
    startFetch: () => void
    successFetch: (quote: any) => void
    caughtError: (error: any) => void
  } = {
    startFetch: () => {
      setLoading(true)
      setError(null)
    },
    successFetch: async (quote) => {
      setLoading(false)
      setQuote(quote)
      setError(null)
    },
    caughtError: (error) => {
      setLoading(false)
      setError(error)
    },
  }

  const getBridgeQuote: () => Promise<void> = async () => {
    try {
      FetchStatusCallback.startFetch()
      const bridgeQuote = await fetchBridgeQuote(
        {
          originChainId: originChainId,
          originTokenAddress: originTokenAddress,
          destinationChainId: destinationChainId,
          destinationTokenAddress: destinationTokenAddress,
          amount: amount,
        },
        synapseSDK
      )
      console.log('bridgeQuote:', bridgeQuote)

      FetchStatusCallback.successFetch(bridgeQuote)
    } catch (error) {
      FetchStatusCallback.caughtError(error)
    }
  }

  return useMemo(() => {
    /** Guardrail Checks for Required Quote Inputs */
    if (!originChainId) {
      return {
        state: QuoteCallbackState.INVALID,
        quote: null,
        callback: null,
        error: QuoteCallbackErrors.REQUIRE_ORIGIN_CHAIN,
      }
    } else if (!originTokenAddress) {
      return {
        state: QuoteCallbackState.INVALID,
        quote: null,
        callback: null,
        error: QuoteCallbackErrors.REQUIRE_ORIGIN_TOKEN,
      }
    } else if (!destinationChainId) {
      return {
        state: QuoteCallbackState.INVALID,
        quote: null,
        callback: null,
        error: QuoteCallbackErrors.REQUIRE_DEST_CHAIN,
      }
    } else if (!destinationTokenAddress) {
      return {
        state: QuoteCallbackState.INVALID,
        quote: null,
        callback: null,
        error: QuoteCallbackErrors.REQUIRE_DEST_TOKEN,
      }
    } else if (!amount) {
      return {
        state: QuoteCallbackState.INVALID,
        quote: null,
        callback: null,
        error: QuoteCallbackErrors.REQUIRE_AMOUNT,
      }
    }

    /** Error occuring during Bridge Quote fetch */
    if (error) {
      return {
        state: QuoteCallbackState.INVALID,
        callback: () => getBridgeQuote(),
        quote: null,
        error: error,
      }
    }

    /** Bridge Quote fetch in progress */
    if (loading) {
      return {
        state: QuoteCallbackState.LOADING,
        callback: () => getBridgeQuote(),
        quote: null,
        error: error,
      }
    }

    /** Valid Quote returned */
    if (quote) {
      return {
        state: QuoteCallbackState.VALID,
        callback: () => getBridgeQuote(),
        quote: quote,
        error: error,
      }
    }

    return {
      state: QuoteCallbackState.IDLE,
      callback: () => getBridgeQuote(),
      quote: quote,
      error: error,
    }
  }, [
    originChainId,
    destinationChainId,
    originTokenAddress,
    destinationTokenAddress,
    amount,
    synapseSDK,
    quote,
    loading,
    error,
  ])
}
