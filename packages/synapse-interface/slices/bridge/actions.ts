import { createAsyncThunk } from '@reduxjs/toolkit'
import { 
  setFromChainId, 
  setFromToken, 
  setToChainId, 
  setToToken, 
  updateDebouncedFromValue 
} from './reducer'
import { CHAINS_BY_ID } from '@/constants/chains'
import { getRoutePossibilities } from '@/utils/routeMaker/generateRoutePossibilities'
import { findTokenByRouteSymbol } from '@/utils/findTokenByRouteSymbol'
import { toast } from 'react-hot-toast'

// Real API integration to parse bridge requests
async function extractParameters(input: string): Promise<any> {
  try {
    const response = await fetch('http://localhost:4000/api/parse-bridge-request', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ query: input }),
    });
    
    const data = await response.json();
    
    if (!data.success) {
      throw new Error(data.error || 'Failed to parse bridge request');
    }
    
    return data.parameters;
  } catch (error) {
    console.error('Error calling bridge parser API:', error);
    // Return null to indicate failure - the calling code will handle this
    return null;
  }
}

export const updateBridgeParameters = createAsyncThunk(
  'bridge/updateBridgeParameters',
  async (input: string, { getState, dispatch }) => {
    const state = getState() as any
    const { fromChainId, fromToken, toChainId, toToken } = state.bridge

    // Get parameters from the LLM API
    const params = await extractParameters(input)
    
    if (!params) {
      toast.error('Could not understand your bridge request. Please try again with a clearer description.')
      return
    }

    // Map chain names to IDs - handle null values gracefully
    const mapChain = (name: string | null) => {
      if (!name) return null;
      return Object.values(CHAINS_BY_ID).find(c => 
        c.name.toLowerCase() === name.toLowerCase())?.id;
    }

    // Extract parameters from the API response, handling null values
    const extracted = {
      amount: params.amount || '',
      fromChainId: mapChain(params.fromChain) || fromChainId,
      fromTokenSymbol: params.fromToken,
      toChainId: mapChain(params.toChain) || toChainId,
      toTokenSymbol: params.toToken === 'unspecified' ? null : params.toToken,
    }
    
    // Handle case where only chains are specified but no tokens
    const onlyChainsSpecified = params.fromChain && params.toChain && 
                                !params.fromToken && !params.toToken;
    
    // Toast to let user know we're using the chains they specified
    if (onlyChainsSpecified) {
      toast.success(
        `Setting up bridge from ${params.fromChain} to ${params.toChain}. Please select tokens and amount.`,
        { duration: 3000 }
      );
    }

    // Get available tokens for these chains
    const { fromTokens, toTokens } = getRoutePossibilities({
      fromChainId: extracted.fromChainId,
      fromToken: null,
      toChainId: extracted.toChainId,
      toToken: null,
    })

    // Find matching token objects - handle null values
    const findToken = (symbol: string | null, tokens: any[]) => {
      if (!symbol) return null;
      return tokens.find(t => t.symbol.toLowerCase() === symbol.toLowerCase());
    }

    // Handle both specific tokens and null values
    let validFromToken;
    let validToToken;
    
    // If fromToken specified, use it; otherwise use default for that chain
    if (extracted.fromTokenSymbol) {
      validFromToken = findToken(extracted.fromTokenSymbol, fromTokens) || fromTokens[0] || fromToken;
    } else {
      // Choose a suitable default token - prefer stablecoins like USDC if available
      const preferredDefault = fromTokens.find(t => t.symbol === 'USDC') || 
                              fromTokens.find(t => t.symbol === 'USDT') ||
                              fromTokens.find(t => t.symbol === 'ETH') ||
                              fromTokens[0];
      validFromToken = preferredDefault || fromToken;
    }
      
    // If toToken specified, use it; otherwise try to match fromToken symbol on destination chain
    if (extracted.toTokenSymbol) {
      validToToken = findToken(extracted.toTokenSymbol, toTokens) || toTokens[0] || toToken;
    } else {
      // Try to find the same token on destination chain
      const sameSymbolToken = validFromToken ? 
        toTokens.find(t => t.symbol === validFromToken.symbol) : null;
        
      validToToken = sameSymbolToken || toTokens[0] || toToken;
    }

    // Validate the route
    const routePossibilities = getRoutePossibilities({
      fromChainId: extracted.fromChainId,
      fromToken: validFromToken,
      toChainId: extracted.toChainId,
      toToken: validToToken,
    })
    
    const isValidRoute = routePossibilities.toTokens.length > 0

    // Check if we have mappings for the chains provided by the LLM
    if (!extracted.fromChainId) {
      toast.error(`Sorry, couldn't recognize the source chain "${params.fromChain}". Please try a different chain.`)
      return
    }

    if (!extracted.toChainId) {
      toast.error(`Sorry, couldn't recognize the destination chain "${params.toChain}". Please try a different chain.`)
      return
    }
    
    // If the route is invalid but we have chains specified, try to find a valid token pair
    if (!isValidRoute && extracted.fromChainId && extracted.toChainId) {
      // Try to find any valid token pair between these chains
      for (const fromTokenOption of fromTokens) {
        const testPossibilities = getRoutePossibilities({
          fromChainId: extracted.fromChainId,
          fromToken: fromTokenOption,
          toChainId: extracted.toChainId,
          toToken: null,
        });
        
        if (testPossibilities.toTokens.length > 0) {
          // Found a valid token pair, use it instead
          validFromToken = fromTokenOption;
          validToToken = testPossibilities.toTokens[0];
          
          // Show a message that we're using different tokens
          toast.success(`Using ${validFromToken.symbol} and ${validToToken.symbol} for your bridge between ${CHAINS_BY_ID[extracted.fromChainId]?.name} and ${CHAINS_BY_ID[extracted.toChainId]?.name}.`);
          
          break;
        }
      }
      
      // Re-check if we found a valid route
      const retryPossibilities = getRoutePossibilities({
        fromChainId: extracted.fromChainId,
        fromToken: validFromToken,
        toChainId: extracted.toChainId,
        toToken: validToToken,
      });
      
      if (retryPossibilities.toTokens.length === 0) {
        // Still no valid route
        toast.error(`Sorry, bridging between ${CHAINS_BY_ID[extracted.fromChainId]?.name} and ${CHAINS_BY_ID[extracted.toChainId]?.name} is not supported.`)
        return
      }
    } else if (!isValidRoute) {
      // Regular invalid route error
      toast.error(`Sorry, bridging from ${validFromToken.symbol} on ${CHAINS_BY_ID[extracted.fromChainId]?.name} to ${validToToken.symbol} on ${CHAINS_BY_ID[extracted.toChainId]?.name} is not supported.`)
      return
    }

    // Success toast with the interpreted parameters
    const fromChainName = CHAINS_BY_ID[extracted.fromChainId]?.name
    const toChainName = CHAINS_BY_ID[extracted.toChainId]?.name
    
    // Different toast message depending on what parameters were provided
    if (extracted.amount) {
      toast.success(
        `Setting up bridge: ${extracted.amount} ${validFromToken.symbol} from ${fromChainName} to ${validToToken.symbol} on ${toChainName}`,
        { duration: 4000 }
      );
    } else if (onlyChainsSpecified) {
      // We already showed a toast for this case
    } else {
      toast.success(
        `Setting up bridge from ${validFromToken.symbol} on ${fromChainName} to ${validToToken.symbol} on ${toChainName}`,
        { duration: 4000 }
      );
    }

    // Update state
    dispatch(setFromChainId(extracted.fromChainId))
    dispatch(setFromToken(validFromToken))
    dispatch(setToChainId(extracted.toChainId))
    dispatch(setToToken(validToToken))
    dispatch(updateDebouncedFromValue(extracted.amount))

    return {
      fromChainId: extracted.fromChainId,
      fromToken: validFromToken,
      toChainId: extracted.toChainId,
      toToken: validToToken,
      amount: extracted.amount
    }
  }
)