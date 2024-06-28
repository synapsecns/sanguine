export const messageBusAbi = [
  {
    inputs: [
      {
        internalType: 'address',
        name: '_gasFeePricing',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_authVerifier',
        type: 'address',
      },
    ],
    stateMutability: 'nonpayable',
    type: 'constructor',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'string',
        name: 'reason',
        type: 'string',
      },
    ],
    name: 'CallReverted',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'bytes32',
        name: 'messageId',
        type: 'bytes32',
      },
      {
        indexed: false,
        internalType: 'enum MessageBusReceiver.TxStatus',
        name: 'status',
        type: 'uint8',
      },
      {
        indexed: true,
        internalType: 'address',
        name: '_dstAddress',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'uint64',
        name: 'srcChainId',
        type: 'uint64',
      },
      {
        indexed: false,
        internalType: 'uint64',
        name: 'srcNonce',
        type: 'uint64',
      },
    ],
    name: 'Executed',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'address',
        name: 'sender',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'srcChainID',
        type: 'uint256',
      },
      {
        indexed: false,
        internalType: 'bytes32',
        name: 'receiver',
        type: 'bytes32',
      },
      {
        indexed: true,
        internalType: 'uint256',
        name: 'dstChainId',
        type: 'uint256',
      },
      {
        indexed: false,
        internalType: 'bytes',
        name: 'message',
        type: 'bytes',
      },
      {
        indexed: false,
        internalType: 'uint64',
        name: 'nonce',
        type: 'uint64',
      },
      {
        indexed: false,
        internalType: 'bytes',
        name: 'options',
        type: 'bytes',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'fee',
        type: 'uint256',
      },
      {
        indexed: true,
        internalType: 'bytes32',
        name: 'messageId',
        type: 'bytes32',
      },
    ],
    name: 'MessageSent',
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
    name: 'authVerifier',
    outputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: '_srcAddress',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: '_srcChainId',
        type: 'uint256',
      },
      {
        internalType: 'bytes32',
        name: '_dstAddress',
        type: 'bytes32',
      },
      {
        internalType: 'uint256',
        name: '_dstChainId',
        type: 'uint256',
      },
      {
        internalType: 'uint256',
        name: '_srcNonce',
        type: 'uint256',
      },
      {
        internalType: 'bytes',
        name: '_message',
        type: 'bytes',
      },
    ],
    name: 'computeMessageId',
    outputs: [
      {
        internalType: 'bytes32',
        name: '',
        type: 'bytes32',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'uint256',
        name: '_dstChainId',
        type: 'uint256',
      },
      {
        internalType: 'bytes',
        name: '_options',
        type: 'bytes',
      },
    ],
    name: 'estimateFee',
    outputs: [
      {
        internalType: 'uint256',
        name: '',
        type: 'uint256',
      },
    ],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'uint256',
        name: '_srcChainId',
        type: 'uint256',
      },
      {
        internalType: 'bytes32',
        name: '_srcAddress',
        type: 'bytes32',
      },
      {
        internalType: 'address',
        name: '_dstAddress',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: '_gasLimit',
        type: 'uint256',
      },
      {
        internalType: 'uint256',
        name: '_nonce',
        type: 'uint256',
      },
      {
        internalType: 'bytes',
        name: '_message',
        type: 'bytes',
      },
      {
        internalType: 'bytes32',
        name: '_messageId',
        type: 'bytes32',
      },
    ],
    name: 'executeMessage',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'fees',
    outputs: [
      {
        internalType: 'uint256',
        name: '',
        type: 'uint256',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'gasFeePricing',
    outputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'bytes32',
        name: '_messageId',
        type: 'bytes32',
      },
    ],
    name: 'getExecutedMessage',
    outputs: [
      {
        internalType: 'enum MessageBusReceiver.TxStatus',
        name: '',
        type: 'uint8',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'nonce',
    outputs: [
      {
        internalType: 'uint64',
        name: '',
        type: 'uint64',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'owner',
    outputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
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
    inputs: [
      {
        internalType: 'address payable',
        name: 'to',
        type: 'address',
      },
    ],
    name: 'rescueGas',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'bytes32',
        name: '_receiver',
        type: 'bytes32',
      },
      {
        internalType: 'uint256',
        name: '_dstChainId',
        type: 'uint256',
      },
      {
        internalType: 'bytes',
        name: '_message',
        type: 'bytes',
      },
      {
        internalType: 'bytes',
        name: '_options',
        type: 'bytes',
      },
    ],
    name: 'sendMessage',
    outputs: [],
    stateMutability: 'payable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'newOwner',
        type: 'address',
      },
    ],
    name: 'transferOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: '_authVerifier',
        type: 'address',
      },
    ],
    name: 'updateAuthVerifier',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: '_gasFeePricing',
        type: 'address',
      },
    ],
    name: 'updateGasFeePricing',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'bytes32',
        name: '_messageId',
        type: 'bytes32',
      },
      {
        internalType: 'enum MessageBusReceiver.TxStatus',
        name: '_status',
        type: 'uint8',
      },
    ],
    name: 'updateMessageStatus',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address payable',
        name: 'to',
        type: 'address',
      },
    ],
    name: 'withdrawGasFees',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
] as const
