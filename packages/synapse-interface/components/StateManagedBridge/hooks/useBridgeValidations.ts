import { useMemo } from 'react';
import { useAccount } from 'wagmi';
import { useBridgeState } from '@/slices/bridge/hooks';
import { BridgeState } from '@/slices/bridge/reducer';
import { useBridgeQuoteState } from '@/slices/bridgeQuote/hooks';
import { BridgeQuoteState } from '@/slices/bridgeQuote/reducer';
import { EMPTY_BRIDGE_QUOTE } from '@/constants/bridge';
import { hasOnlyZeroes } from '@/utils/hasOnlyZeroes';
import { useBridgeSelections } from './useBridgeSelections';


export const useBridgeValidations = () => {
  const { chainId } = useAccount();
  const {
    fromChainId, toChainId, fromToken, toToken, debouncedFromValue,
  }: BridgeState = useBridgeState();
  const { bridgeQuote }: BridgeQuoteState = useBridgeQuoteState();
  const { fromTokenBalance, debouncedFromValueBigInt } = useBridgeSelections();

  const hasValidInput: boolean = useMemo(() => {
    if (debouncedFromValue === '') return false;
    if (hasOnlyZeroes(debouncedFromValue)) return false;
    return debouncedFromValueBigInt > 0n;
  }, [debouncedFromValue, debouncedFromValueBigInt]);

  const hasValidSelections = useMemo(() => {
    return Boolean(fromChainId && fromToken && toChainId && toToken);
  }, [fromChainId, fromToken, toChainId, toToken]);

  const hasValidQuote: boolean = useMemo(() => {
    return bridgeQuote !== EMPTY_BRIDGE_QUOTE;
  }, [bridgeQuote]);

  const hasSufficientBalance: boolean = useMemo(() => {
    return hasValidSelections
      ? debouncedFromValueBigInt <= fromTokenBalance
      : false;
  }, [hasValidSelections, debouncedFromValueBigInt, fromTokenBalance]);

  const doesChainSelectionsMatchBridgeQuote = useMemo(() => {
    return (
      fromChainId === bridgeQuote.originChainId &&
      toChainId === bridgeQuote.destChainId
    );
  }, [
    fromChainId,
    toChainId,
    bridgeQuote.originChainId,
    bridgeQuote.destChainId,
  ]);

  const isBridgeQuoteAmountGreaterThanInputForRfq = useMemo(() => {
    return (
      bridgeQuote.bridgeModuleName === 'SynapseRFQ' &&
      bridgeQuote.outputAmount > debouncedFromValueBigInt
    );
  }, [
    bridgeQuote.outputAmount,
    bridgeQuote.bridgeModuleName,
    debouncedFromValueBigInt,
  ]);

  const onSelectedChain: boolean = useMemo(() => {
    return chainId === fromChainId;
  }, [fromChainId, chainId]);

  return {
    hasValidInput,
    hasValidSelections,
    hasValidQuote,
    hasSufficientBalance,
    doesChainSelectionsMatchBridgeQuote,
    isBridgeQuoteAmountGreaterThanInputForRfq,
    onSelectedChain,
  };
};
