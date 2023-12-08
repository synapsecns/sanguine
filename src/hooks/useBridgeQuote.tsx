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

interface UseBridgeQuoteArgs {
  originChainId?: number
  originTokenAddress?: string
  destinationChainId?: number
  destinationTokenAddress?: string
  amount?: bigint
  synapseSDK?: any
}

export function useBridgeQuote({
  originChainId,
  destinationChainId,
  originTokenAddress,
  destinationTokenAddress,
  amount,
  synapseSDK,
}: UseBridgeQuoteArgs): {
  state: QuoteCallbackState
  callback: () => Promise<void>
  reset: () => void
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
      console.log('start bridge quote fetch')
      setLoading(true)
      setError(null)
    },
    successFetch: async (quote) => {
      console.log('success bridge quote fetch')
      setLoading(false)
      setQuote(quote)
      setError(null)
    },
    caughtError: (error) => {
      console.log('error bridge quote fetch ', error)
      setLoading(false)
      setError(error)
    },
  }

  const resetQuote = () => setQuote(null)

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
        reset: () => resetQuote(),
        error: QuoteCallbackErrors.REQUIRE_ORIGIN_CHAIN,
      }
    } else if (!originTokenAddress) {
      return {
        state: QuoteCallbackState.INVALID,
        quote: null,
        callback: null,
        reset: () => resetQuote(),
        error: QuoteCallbackErrors.REQUIRE_ORIGIN_TOKEN,
      }
    } else if (!destinationChainId) {
      return {
        state: QuoteCallbackState.INVALID,
        quote: null,
        callback: null,
        reset: () => resetQuote(),
        error: QuoteCallbackErrors.REQUIRE_DEST_CHAIN,
      }
    } else if (!destinationTokenAddress) {
      return {
        state: QuoteCallbackState.INVALID,
        quote: null,
        callback: null,
        reset: () => resetQuote(),
        error: QuoteCallbackErrors.REQUIRE_DEST_TOKEN,
      }
    } else if (!amount) {
      return {
        state: QuoteCallbackState.INVALID,
        quote: null,
        callback: null,
        reset: () => resetQuote(),
        error: QuoteCallbackErrors.REQUIRE_AMOUNT,
      }
    }

    /** Error occuring during Bridge Quote fetch */
    if (error) {
      return {
        state: QuoteCallbackState.INVALID,
        callback: () => getBridgeQuote(),
        reset: () => resetQuote(),
        quote: null,
        error: error,
      }
    }

    /** Bridge Quote fetch in progress */
    if (loading) {
      return {
        state: QuoteCallbackState.LOADING,
        callback: () => getBridgeQuote(),
        reset: () => resetQuote(),
        quote: null,
        error: error,
      }
    }

    /** Valid Quote returned */
    if (quote) {
      return {
        state: QuoteCallbackState.VALID,
        callback: () => getBridgeQuote(),
        reset: () => resetQuote(),
        quote: quote,
        error: error,
      }
    }

    return {
      state: QuoteCallbackState.IDLE,
      callback: () => getBridgeQuote(),
      reset: () => resetQuote(),
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
