export const InterchainClientV1Abi = [
  {
    inputs: [
      { internalType: 'address', name: 'interchainDB', type: 'address' },
      { internalType: 'address', name: 'owner_', type: 'address' },
    ],
    stateMutability: 'nonpayable',
    type: 'constructor',
  },
  {
    inputs: [{ internalType: 'uint16', name: 'version', type: 'uint16' }],
    name: 'AppConfigLib__VersionInvalid',
    type: 'error',
  },
  {
    inputs: [{ internalType: 'uint64', name: 'entryIndex', type: 'uint64' }],
    name: 'BatchingV1__EntryIndexNotZero',
    type: 'error',
  },
  { inputs: [], name: 'BatchingV1__ProofNotEmpty', type: 'error' },
  {
    inputs: [{ internalType: 'address', name: 'module', type: 'address' }],
    name: 'InterchainClientV1__BatchConflict',
    type: 'error',
  },
  {
    inputs: [{ internalType: 'uint64', name: 'chainId', type: 'uint64' }],
    name: 'InterchainClientV1__ChainIdNotLinked',
    type: 'error',
  },
  {
    inputs: [{ internalType: 'uint64', name: 'chainId', type: 'uint64' }],
    name: 'InterchainClientV1__ChainIdNotRemote',
    type: 'error',
  },
  {
    inputs: [{ internalType: 'uint64', name: 'chainId', type: 'uint64' }],
    name: 'InterchainClientV1__DstChainIdNotLocal',
    type: 'error',
  },
  {
    inputs: [],
    name: 'InterchainClientV1__ExecutionServiceZeroAddress',
    type: 'error',
  },
  {
    inputs: [
      { internalType: 'uint256', name: 'feeAmount', type: 'uint256' },
      { internalType: 'uint256', name: 'minRequired', type: 'uint256' },
    ],
    name: 'InterchainClientV1__FeeAmountBelowMin',
    type: 'error',
  },
  {
    inputs: [
      { internalType: 'uint256', name: 'gasLeft', type: 'uint256' },
      { internalType: 'uint256', name: 'minRequired', type: 'uint256' },
    ],
    name: 'InterchainClientV1__GasLeftBelowMin',
    type: 'error',
  },
  { inputs: [], name: 'InterchainClientV1__GuardZeroAddress', type: 'error' },
  {
    inputs: [{ internalType: 'bytes32', name: 'client', type: 'bytes32' }],
    name: 'InterchainClientV1__LinkedClientNotEVM',
    type: 'error',
  },
  {
    inputs: [
      { internalType: 'uint256', name: 'msgValue', type: 'uint256' },
      { internalType: 'uint256', name: 'required', type: 'uint256' },
    ],
    name: 'InterchainClientV1__MsgValueMismatch',
    type: 'error',
  },
  {
    inputs: [{ internalType: 'address', name: 'receiver', type: 'address' }],
    name: 'InterchainClientV1__ReceiverNotICApp',
    type: 'error',
  },
  {
    inputs: [],
    name: 'InterchainClientV1__ReceiverZeroAddress',
    type: 'error',
  },
  {
    inputs: [{ internalType: 'address', name: 'receiver', type: 'address' }],
    name: 'InterchainClientV1__ReceiverZeroRequiredResponses',
    type: 'error',
  },
  {
    inputs: [
      { internalType: 'uint256', name: 'responsesAmount', type: 'uint256' },
      { internalType: 'uint256', name: 'minRequired', type: 'uint256' },
    ],
    name: 'InterchainClientV1__ResponsesAmountBelowMin',
    type: 'error',
  },
  {
    inputs: [
      { internalType: 'bytes32', name: 'transactionId', type: 'bytes32' },
    ],
    name: 'InterchainClientV1__TxAlreadyExecuted',
    type: 'error',
  },
  {
    inputs: [
      { internalType: 'bytes32', name: 'transactionId', type: 'bytes32' },
    ],
    name: 'InterchainClientV1__TxNotExecuted',
    type: 'error',
  },
  {
    inputs: [
      { internalType: 'uint16', name: 'txVersion', type: 'uint16' },
      { internalType: 'uint16', name: 'required', type: 'uint16' },
    ],
    name: 'InterchainClientV1__TxVersionMismatch',
    type: 'error',
  },
  {
    inputs: [{ internalType: 'uint16', name: 'version', type: 'uint16' }],
    name: 'OptionsLib__VersionInvalid',
    type: 'error',
  },
  {
    inputs: [{ internalType: 'address', name: 'owner', type: 'address' }],
    name: 'OwnableInvalidOwner',
    type: 'error',
  },
  {
    inputs: [{ internalType: 'address', name: 'account', type: 'address' }],
    name: 'OwnableUnauthorizedAccount',
    type: 'error',
  },
  {
    inputs: [
      { internalType: 'uint8', name: 'bits', type: 'uint8' },
      { internalType: 'uint256', name: 'value', type: 'uint256' },
    ],
    name: 'SafeCastOverflowedUintDowncast',
    type: 'error',
  },
  {
    inputs: [
      { internalType: 'bytes', name: 'versionedPayload', type: 'bytes' },
    ],
    name: 'VersionedPayload__PayloadTooShort',
    type: 'error',
  },
  { inputs: [], name: 'VersionedPayload__PrecompileFailed', type: 'error' },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'address',
        name: 'guard',
        type: 'address',
      },
    ],
    name: 'DefaultGuardSet',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'bytes32',
        name: 'transactionId',
        type: 'bytes32',
      },
      {
        indexed: false,
        internalType: 'uint64',
        name: 'dbNonce',
        type: 'uint64',
      },
      {
        indexed: false,
        internalType: 'uint64',
        name: 'entryIndex',
        type: 'uint64',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'executor',
        type: 'address',
      },
    ],
    name: 'ExecutionProofWritten',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'bytes32',
        name: 'transactionId',
        type: 'bytes32',
      },
      {
        indexed: false,
        internalType: 'uint64',
        name: 'dbNonce',
        type: 'uint64',
      },
      {
        indexed: false,
        internalType: 'uint64',
        name: 'entryIndex',
        type: 'uint64',
      },
      {
        indexed: false,
        internalType: 'uint64',
        name: 'srcChainId',
        type: 'uint64',
      },
      {
        indexed: true,
        internalType: 'bytes32',
        name: 'srcSender',
        type: 'bytes32',
      },
      {
        indexed: true,
        internalType: 'bytes32',
        name: 'dstReceiver',
        type: 'bytes32',
      },
    ],
    name: 'InterchainTransactionReceived',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'bytes32',
        name: 'transactionId',
        type: 'bytes32',
      },
      {
        indexed: false,
        internalType: 'uint64',
        name: 'dbNonce',
        type: 'uint64',
      },
      {
        indexed: false,
        internalType: 'uint64',
        name: 'entryIndex',
        type: 'uint64',
      },
      {
        indexed: false,
        internalType: 'uint64',
        name: 'dstChainId',
        type: 'uint64',
      },
      {
        indexed: true,
        internalType: 'bytes32',
        name: 'srcSender',
        type: 'bytes32',
      },
      {
        indexed: true,
        internalType: 'bytes32',
        name: 'dstReceiver',
        type: 'bytes32',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'verificationFee',
        type: 'uint256',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'executionFee',
        type: 'uint256',
      },
      { indexed: false, internalType: 'bytes', name: 'options', type: 'bytes' },
      { indexed: false, internalType: 'bytes', name: 'message', type: 'bytes' },
    ],
    name: 'InterchainTransactionSent',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'uint64',
        name: 'chainId',
        type: 'uint64',
      },
      {
        indexed: false,
        internalType: 'bytes32',
        name: 'client',
        type: 'bytes32',
      },
    ],
    name: 'LinkedClientSet',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'address',
        name: 'previousOwner',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'newOwner',
        type: 'address',
      },
    ],
    name: 'OwnershipTransferred',
    type: 'event',
  },
  {
    inputs: [],
    name: 'CLIENT_VERSION',
    outputs: [{ internalType: 'uint16', name: '', type: 'uint16' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'INTERCHAIN_DB',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'bytes', name: 'encodedOptions', type: 'bytes' }],
    name: 'decodeOptions',
    outputs: [
      {
        components: [
          { internalType: 'uint256', name: 'gasLimit', type: 'uint256' },
          { internalType: 'uint256', name: 'gasAirdrop', type: 'uint256' },
        ],
        internalType: 'struct OptionsV1',
        name: '',
        type: 'tuple',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'defaultGuard',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        components: [
          { internalType: 'uint64', name: 'srcChainId', type: 'uint64' },
          { internalType: 'uint64', name: 'dstChainId', type: 'uint64' },
          { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
          { internalType: 'uint64', name: 'entryIndex', type: 'uint64' },
          { internalType: 'bytes32', name: 'srcSender', type: 'bytes32' },
          { internalType: 'bytes32', name: 'dstReceiver', type: 'bytes32' },
          { internalType: 'bytes', name: 'options', type: 'bytes' },
          { internalType: 'bytes', name: 'message', type: 'bytes' },
        ],
        internalType: 'struct InterchainTransaction',
        name: 'icTx',
        type: 'tuple',
      },
    ],
    name: 'encodeTransaction',
    outputs: [{ internalType: 'bytes', name: '', type: 'bytes' }],
    stateMutability: 'pure',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'address', name: 'receiver', type: 'address' }],
    name: 'getAppReceivingConfigV1',
    outputs: [
      {
        components: [
          {
            internalType: 'uint256',
            name: 'requiredResponses',
            type: 'uint256',
          },
          {
            internalType: 'uint256',
            name: 'optimisticPeriod',
            type: 'uint256',
          },
          { internalType: 'uint256', name: 'guardFlag', type: 'uint256' },
          { internalType: 'address', name: 'guard', type: 'address' },
        ],
        internalType: 'struct AppConfigV1',
        name: 'config',
        type: 'tuple',
      },
      { internalType: 'address[]', name: 'modules', type: 'address[]' },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'bytes', name: 'encodedTx', type: 'bytes' }],
    name: 'getExecutor',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'bytes32', name: 'transactionId', type: 'bytes32' },
    ],
    name: 'getExecutorById',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'uint64', name: 'dstChainId', type: 'uint64' },
      { internalType: 'address', name: 'srcExecutionService', type: 'address' },
      { internalType: 'address[]', name: 'srcModules', type: 'address[]' },
      { internalType: 'bytes', name: 'options', type: 'bytes' },
      { internalType: 'uint256', name: 'messageLen', type: 'uint256' },
    ],
    name: 'getInterchainFee',
    outputs: [{ internalType: 'uint256', name: 'fee', type: 'uint256' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'uint64', name: 'chainId', type: 'uint64' }],
    name: 'getLinkedClient',
    outputs: [{ internalType: 'bytes32', name: '', type: 'bytes32' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'uint64', name: 'chainId', type: 'uint64' }],
    name: 'getLinkedClientEVM',
    outputs: [
      { internalType: 'address', name: 'linkedClientEVM', type: 'address' },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        components: [
          { internalType: 'uint64', name: 'srcChainId', type: 'uint64' },
          { internalType: 'uint64', name: 'dstChainId', type: 'uint64' },
          { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
          { internalType: 'uint64', name: 'entryIndex', type: 'uint64' },
          { internalType: 'bytes32', name: 'srcSender', type: 'bytes32' },
          { internalType: 'bytes32', name: 'dstReceiver', type: 'bytes32' },
          { internalType: 'bytes', name: 'options', type: 'bytes' },
          { internalType: 'bytes', name: 'message', type: 'bytes' },
        ],
        internalType: 'struct InterchainTransaction',
        name: 'icTx',
        type: 'tuple',
      },
      { internalType: 'bytes32[]', name: 'proof', type: 'bytes32[]' },
    ],
    name: 'getTxReadinessV1',
    outputs: [
      {
        internalType: 'enum IInterchainClientV1.TxReadiness',
        name: 'status',
        type: 'uint8',
      },
      { internalType: 'bytes32', name: 'firstArg', type: 'bytes32' },
      { internalType: 'bytes32', name: 'secondArg', type: 'bytes32' },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'uint256', name: 'gasLimit', type: 'uint256' },
      { internalType: 'bytes', name: 'transaction', type: 'bytes' },
      { internalType: 'bytes32[]', name: 'proof', type: 'bytes32[]' },
    ],
    name: 'interchainExecute',
    outputs: [],
    stateMutability: 'payable',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'uint64', name: 'dstChainId', type: 'uint64' },
      { internalType: 'bytes32', name: 'receiver', type: 'bytes32' },
      { internalType: 'address', name: 'srcExecutionService', type: 'address' },
      { internalType: 'address[]', name: 'srcModules', type: 'address[]' },
      { internalType: 'bytes', name: 'options', type: 'bytes' },
      { internalType: 'bytes', name: 'message', type: 'bytes' },
    ],
    name: 'interchainSend',
    outputs: [
      {
        components: [
          { internalType: 'bytes32', name: 'transactionId', type: 'bytes32' },
          { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
          { internalType: 'uint64', name: 'entryIndex', type: 'uint64' },
        ],
        internalType: 'struct InterchainTxDescriptor',
        name: 'desc',
        type: 'tuple',
      },
    ],
    stateMutability: 'payable',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'uint64', name: 'dstChainId', type: 'uint64' },
      { internalType: 'address', name: 'receiver', type: 'address' },
      { internalType: 'address', name: 'srcExecutionService', type: 'address' },
      { internalType: 'address[]', name: 'srcModules', type: 'address[]' },
      { internalType: 'bytes', name: 'options', type: 'bytes' },
      { internalType: 'bytes', name: 'message', type: 'bytes' },
    ],
    name: 'interchainSendEVM',
    outputs: [
      {
        components: [
          { internalType: 'bytes32', name: 'transactionId', type: 'bytes32' },
          { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
          { internalType: 'uint64', name: 'entryIndex', type: 'uint64' },
        ],
        internalType: 'struct InterchainTxDescriptor',
        name: 'desc',
        type: 'tuple',
      },
    ],
    stateMutability: 'payable',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'bytes', name: 'encodedTx', type: 'bytes' },
      { internalType: 'bytes32[]', name: 'proof', type: 'bytes32[]' },
    ],
    name: 'isExecutable',
    outputs: [{ internalType: 'bool', name: '', type: 'bool' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'owner',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'address', name: 'guard', type: 'address' }],
    name: 'setDefaultGuard',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'uint64', name: 'chainId', type: 'uint64' },
      { internalType: 'bytes32', name: 'client', type: 'bytes32' },
    ],
    name: 'setLinkedClient',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'address', name: 'newOwner', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'bytes32', name: 'transactionId', type: 'bytes32' },
    ],
    name: 'writeExecutionProof',
    outputs: [
      { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
      { internalType: 'uint64', name: 'entryIndex', type: 'uint64' },
    ],
    stateMutability: 'nonpayable',
    type: 'function',
  },
] as const
