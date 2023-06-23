import { useSelector } from 'react-redux';
import { TransactionButton } from '@/components/buttons/TransactionButton';
import { EMPTY_BRIDGE_QUOTE, EMPTY_BRIDGE_QUOTE_ZERO } from '@/constants/bridge';
import { RootState } from '../../store/store'

export const BridgeTransactionButton = ({ approveTxn, executeBridge, isApproved }) => {
    // Get state from Redux store
    const {
        fromToken,
        fromValue,
        fromChainId,
        isLoading,
        bridgeQuote
    } = useSelector((state: RootState) => state.bridge)

    const isButtonDisabled = isLoading ||
      bridgeQuote === EMPTY_BRIDGE_QUOTE_ZERO ||
      bridgeQuote === EMPTY_BRIDGE_QUOTE;

    let buttonProperties;

    if (!isLoading && bridgeQuote?.feeAmount?.eq(0) && fromValue.gt(0))
    {
        buttonProperties = {
            label: `Amount must be greater than fee`,
            onClick: null
        }
    }
    else if (!isApproved) {
      buttonProperties = {
        onClick: approveTxn,
        label: `Approve ${fromToken.symbol}`,
        pendingLabel: "Approving"
      };
    } else {
      buttonProperties = {
        onClick: executeBridge,
        label: `Bridge ${fromToken.symbol}`,
        pendingLabel: "Bridging"
      };
    }

    return (
      buttonProperties && (
        <TransactionButton
          {...buttonProperties}
          disabled={isButtonDisabled}
          chainId={fromChainId}
        />
      )
    );
};
