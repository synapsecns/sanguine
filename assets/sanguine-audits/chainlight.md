# (Synapse Protocol) Interchain App Smart Contract Audit Issues (DRAFT)


1. [SYNAPSE-001] Make sure `threshold` is equal to or less than the length of `_signers`
    - Add `require(threshold <= self._signers.length());` on `modifyThreshold` of `contracts/libs/ThresholdECDSA.sol`
1. [SYNAPSE-002] Early return of `verifySignedHash` once the required threshold is met could save gas costs
    - Return function early if `if (validSignatures == threshold) {` check is satisfied.
1. [SYNAPSE-003] Ensure timely updates to remote chain gas prices
    - The gas oracle contract is responsible for calculating the charges for interchain messages using the gas prices of the remote chains. If the gas prices stored in the Oracle contract for the remote chain are not updated for a certain period of time, there may be significant discrepancies between the stored and current market prices. This discrepancy can result in significant financial loss to either the protocol or the users. To mitigate this risk, the Gas Oracle contract should track the last update time of the gas prices. If the gas prices have not been updated within a defined threshold (e.g., 1 hour), the Oracle contract should not return the outdated prices. Implement this by adding a timestamp for the last update and logic to check if the elapsed time exceeds the allowed threshold before returning the gas price.
2. [SYNAPSE-004] Make sure the execution service uses the entire requested gas limit
    - The execution service sends an interchain transaction on the counterparty chain on behalf of the user. They collect the execution fee on the source chain for sending the transaction. Therefore, the gas limit of the transaction on the destination chain must be the same as the gas limit on the message requested by the user. However, the check in `contracts/InterchainClientV1.sol` does not force that the left gas meets the gas limit because there are additional operations after the application contract is invoked.
    - Change `(a)` to `gasLeft <= gasleft() + baseGas`. (or `bufferGas`)

    - ```solidity
        uint256 gasLeft = gasleft();
        if (gasLeft <= gasLimit) { // (a)
            revert InterchainClientV1__GasLeftBelowMin(gasLeft, gasLimit);
        }
        // Pass the full msg.value to the app: we have already checked that it matches the requested gas airdrop.
        IInterchainApp(icTx.dstReceiver.bytes32ToAddress()).appReceive{gas: gasLimit```

1. [SYNAPSE-005] `_convertRemoteValueToLocalUnits` does not take into account the decimal difference
    - `_remoteGasData[remoteChainId].nativePrice` is the price of 1 wei of the native token of the remote chain relative to ETH (i.e., if the ETH price is $5,000 and the chain X's native token is $250 then this value will be 0.05 ETH, 5 * 1e16).
    - `_convertRemoteValueToLocalUnits` does not take into account the decimal difference between the two chains, and this can result in significant loss.
    - Let's assume the user sends 2e13 X's native tokens from the ETH chain to the X chain, then the calculation would be `int(int(2e13) * 1e6) * (0.05 * int(1e18)) // int(1e18) == 1e18`. It means that if user pays 1 ETH ($5,000), they can get 20,000,000 X Token ($5,000,000,000).
1. [SYNAPSE-006] Consider implementing rate limiting for compliance and to prevent DoS attacks.
    - Other omnichain messaging protocols (i.e., Layer Zero) support message rate limiting feature.
    - Consider implementing rate limiting for compliance and to prevent DoS attacks.
        - Ref: https://docs.layerzero.network/v2/developers/evm/oapp/rate-limiter
1. [SYNAPSE-007] Implement configurable message ordering by app contracts
    - Currently, all interchain messages are processed out of order, making them vulnerable to front-running attacks. To incrase security and flexibility, it is recommended that app contracts be allowed to choose between out-of-order and ordered message processing. Implement a configurable option in the contract that allows app contracts to choose their preferred and appropriate message ordering method. This will help mitigate the risk of front-running by enabling ordered processing where message sequence integrity is maintained, providing an additional layer of protection for interchain communications.
    - References:
        - Out of order execution risks: https://blog.trailofbits.com/2024/03/01/when-try-try-try-again-leads-to-out-of-order-execution-bugs/
        - Force ordering risks: https://medium.com/@Heuss/layerzeros-cross-chain-messaging-vulnerability-e5ef48c5ccec
        - Reference implementation by LZ: https://docs.layerzero.network/v2/developers/evm/oapp/message-ordering
        - Common message ordering patterns: https://docs.layerzero.network/v2/developers/evm/oapp/message-design-patterns
1. [SYNAPSE-008] Allow users to pay the additional gas limit to minimize the risk of messaging failure
    - Depending on the app contract implementation, the gas limit option passed to `_sendToLinkedApp()` can be a fixed value or a user provided value. However, both methods have the possibility of message failure if the app or user sets an incorrect value. To minimize the possibility of message failure, it is recommended that app contracts set a minimum gas limit required at the destination chain, and support the ability for the user to specify additional gas limits.
    - It should be very careful not to use the existing app balance to pay for the additional gas limits.
1. [SYNAPSE-009] Add deadline field to the `InterchainTransaction` message
    - If a message fails in `appReceive()` due to an out-of-gas or unexpected reverts on the app contract (e.g. the minimum amount on the interchain swap request is not met), anyone can always try to send the message again. Let's assume that the behavior of the swap message is to pull all the tokens of an authorized address and then swap the tokens for the other token. The pulled token amount could be significantly different at the two points in time, which could lead to uncertain results and cause security issues. Therefore it is recommended to add a timeout to the `InterchainTransaction` field and `appReceive()` check if `block.timestamp` exceeds the timeout.
1. [SYNAPSE-010] `InterchainClientV1._getFinalizedResponsesCount()` should consider each chain's finalized time
    - `InterchainClientV1._getFinalizedResponsesCount()` checks if the message is finalized with `confirmedAt + optimisticSeconds < block.timestamp`. However, the finalized time is different for each chain. Therefore, a message may be executed in the destination chain before the corresponding message request is finalized in the source chain. In the worst case, the message request may be canceled due to a re-org in the source chain and the message request may be processed in the destination chain. It is recommended to set the finalized time for each source chain be set separately instead of using `optimisticSeconds` globally for the verification.
1. [SYNAPSE-011] `SynapseModule` has excessive permissions for `Owner`
    - `SynapseModule` verifies the interchain messages with the verifiers set by `Owner`. However, the `Owner` can add or remove verifiers at will. In the script code, since `Owner` is EOA, if the private key is leaked, all privileges of `SynapseModule` will be stolen. Therefore, it is recommended to set `Owner` to a multi-sig wallet or restrict to call the sensitive functions to be called when a certain quorum of verifiers is exceeded.
1. [SYNAPSE-012] Delegated interchain messaging fee is vulnerable to the Denial of Service attack
    - The SYNAPSE protocol allows app contracts to pay messaging fees on behalf of end users. The idea that the app pays the messaging fee is vulnerable to a Denial of Service attack. If messages are repeatedly sent from a chain with lower gas fees (i.e. Arbitrum) to a chain with higher gas fees (i.e. Ethereum mainnet), the native tokens held by the app contract will easily be depleted (i.e. it costs $0.1 to make a transaction, but $10 worth of tokens will be depleted). Therefore, it is recommended to take measures such as only paying gas to trusted addresses, or only paying when sending messages from a chain with higher gas fees to a chain with lower gas fees.
1. [SYNAPSE-013] Use two-step ownership transfer
    - The owner can change the owner with the transferOwnership function. However, if the owner is changed incorrectly by entering the wrong address, it cannot be taken back. Also, the contract will become un-upgradable because only the owner can perform the upgrade, so it is recommended to add a fail-safe. Use OZ.Ownable2StepUpgradeable instead of OZ.OwnableUpgradeable at `SynapseModule`, `SynapseGasOracleV1` and `InterchainClientV1`.
1. [SYNAPSE-014] Recommend apps to use the patterns for handling messaging failures
    - To ensure the apps can effectively manage messaging failures, it is advisable to recommend robust handling patterns on the app side. Additionally, at the protocol level (i.e., client), it would be great to have more states for messages that allow them to be deleted or blocked. This approach will prevent any side effects from the stuck messages.

1. [TOKENCONTROLLER-001] (Minor) In TokenController.setGasLimit() should check that gasLimit is greater than or equal to getMinGasLimit().
1. [TOKENCONTROLLER-002] (Minor) In TokenController._getBasicGasLimit(), when gasLimit is 0, it sets gasLimit to DEFAULT_GAS_LIMIT, which might be better to remove.
Making setGasLimit() a mandatory call could be beneficial, as each CApp deployed on different chains may require different gas amounts. By removing this logic from _getBasicGasLimit(), you ensure that messages are only sent to dstChainIds where the gas limit has been explicitly set using setGasLimit().
