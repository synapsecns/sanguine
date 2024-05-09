export const InterchainDBAbi = [
  {
    inputs: [{ internalType: 'uint64', name: 'entryIndex', type: 'uint64' }],
    name: 'BatchingV1__EntryIndexNotZero',
    type: 'error',
  },
  { inputs: [], name: 'BatchingV1__ProofNotEmpty', type: 'error' },
  {
    inputs: [
      { internalType: 'address', name: 'module', type: 'address' },
      { internalType: 'bytes32', name: 'existingBatchRoot', type: 'bytes32' },
      {
        components: [
          { internalType: 'uint64', name: 'srcChainId', type: 'uint64' },
          { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
          { internalType: 'bytes32', name: 'batchRoot', type: 'bytes32' },
        ],
        internalType: 'struct InterchainBatch',
        name: 'newBatch',
        type: 'tuple',
      },
    ],
    name: 'InterchainDB__BatchConflict',
    type: 'error',
  },
  {
    inputs: [
      { internalType: 'uint16', name: 'version', type: 'uint16' },
      { internalType: 'uint16', name: 'required', type: 'uint16' },
    ],
    name: 'InterchainDB__BatchVersionMismatch',
    type: 'error',
  },
  {
    inputs: [{ internalType: 'uint64', name: 'chainId', type: 'uint64' }],
    name: 'InterchainDB__ChainIdNotRemote',
    type: 'error',
  },
  {
    inputs: [
      { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
      { internalType: 'uint64', name: 'entryIndex', type: 'uint64' },
      { internalType: 'uint64', name: 'batchSize', type: 'uint64' },
    ],
    name: 'InterchainDB__EntryIndexOutOfRange',
    type: 'error',
  },
  {
    inputs: [
      { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
      { internalType: 'uint64', name: 'start', type: 'uint64' },
      { internalType: 'uint64', name: 'end', type: 'uint64' },
    ],
    name: 'InterchainDB__EntryRangeInvalid',
    type: 'error',
  },
  {
    inputs: [
      { internalType: 'uint256', name: 'feeAmount', type: 'uint256' },
      { internalType: 'uint256', name: 'minRequired', type: 'uint256' },
    ],
    name: 'InterchainDB__FeeAmountBelowMin',
    type: 'error',
  },
  { inputs: [], name: 'InterchainDB__ModulesNotProvided', type: 'error' },
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
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'uint64',
        name: 'dbNonce',
        type: 'uint64',
      },
      {
        indexed: false,
        internalType: 'bytes32',
        name: 'batchRoot',
        type: 'bytes32',
      },
    ],
    name: 'InterchainBatchFinalized',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'uint64',
        name: 'dstChainId',
        type: 'uint64',
      },
      {
        indexed: true,
        internalType: 'uint64',
        name: 'dbNonce',
        type: 'uint64',
      },
      {
        indexed: false,
        internalType: 'bytes32',
        name: 'batchRoot',
        type: 'bytes32',
      },
      {
        indexed: false,
        internalType: 'address[]',
        name: 'srcModules',
        type: 'address[]',
      },
    ],
    name: 'InterchainBatchVerificationRequested',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'address',
        name: 'module',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'uint64',
        name: 'srcChainId',
        type: 'uint64',
      },
      {
        indexed: true,
        internalType: 'uint64',
        name: 'dbNonce',
        type: 'uint64',
      },
      {
        indexed: false,
        internalType: 'bytes32',
        name: 'batchRoot',
        type: 'bytes32',
      },
    ],
    name: 'InterchainBatchVerified',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
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
        internalType: 'bytes32',
        name: 'srcWriter',
        type: 'bytes32',
      },
      {
        indexed: false,
        internalType: 'bytes32',
        name: 'dataHash',
        type: 'bytes32',
      },
    ],
    name: 'InterchainEntryWritten',
    type: 'event',
  },
  {
    inputs: [],
    name: 'DB_VERSION',
    outputs: [{ internalType: 'uint16', name: '', type: 'uint16' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'address', name: 'dstModule', type: 'address' },
      {
        components: [
          { internalType: 'uint64', name: 'srcChainId', type: 'uint64' },
          { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
          { internalType: 'bytes32', name: 'batchRoot', type: 'bytes32' },
        ],
        internalType: 'struct InterchainBatch',
        name: 'batch',
        type: 'tuple',
      },
    ],
    name: 'checkBatchVerification',
    outputs: [
      { internalType: 'uint256', name: 'moduleVerifiedAt', type: 'uint256' },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'uint64', name: 'dbNonce', type: 'uint64' }],
    name: 'getBatch',
    outputs: [
      {
        components: [
          { internalType: 'uint64', name: 'srcChainId', type: 'uint64' },
          { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
          { internalType: 'bytes32', name: 'batchRoot', type: 'bytes32' },
        ],
        internalType: 'struct InterchainBatch',
        name: '',
        type: 'tuple',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'uint64', name: 'dbNonce', type: 'uint64' }],
    name: 'getBatchLeafs',
    outputs: [{ internalType: 'bytes32[]', name: 'leafs', type: 'bytes32[]' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
      { internalType: 'uint64', name: 'start', type: 'uint64' },
      { internalType: 'uint64', name: 'end', type: 'uint64' },
    ],
    name: 'getBatchLeafsPaginated',
    outputs: [{ internalType: 'bytes32[]', name: 'leafs', type: 'bytes32[]' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        components: [
          { internalType: 'uint64', name: 'srcChainId', type: 'uint64' },
          { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
          { internalType: 'uint64', name: 'entryIndex', type: 'uint64' },
          { internalType: 'bytes32', name: 'srcWriter', type: 'bytes32' },
          { internalType: 'bytes32', name: 'dataHash', type: 'bytes32' },
        ],
        internalType: 'struct InterchainEntry',
        name: 'entry',
        type: 'tuple',
      },
      { internalType: 'bytes32[]', name: 'proof', type: 'bytes32[]' },
    ],
    name: 'getBatchRoot',
    outputs: [{ internalType: 'bytes32', name: '', type: 'bytes32' }],
    stateMutability: 'pure',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'uint64', name: 'dbNonce', type: 'uint64' }],
    name: 'getBatchSize',
    outputs: [{ internalType: 'uint64', name: '', type: 'uint64' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'getDBNonce',
    outputs: [{ internalType: 'uint64', name: '', type: 'uint64' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
      { internalType: 'uint64', name: 'entryIndex', type: 'uint64' },
    ],
    name: 'getEntryProof',
    outputs: [{ internalType: 'bytes32[]', name: 'proof', type: 'bytes32[]' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
      { internalType: 'uint64', name: 'entryIndex', type: 'uint64' },
    ],
    name: 'getEntryValue',
    outputs: [{ internalType: 'bytes32', name: '', type: 'bytes32' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'uint64', name: 'dstChainId', type: 'uint64' },
      { internalType: 'address[]', name: 'srcModules', type: 'address[]' },
    ],
    name: 'getInterchainFee',
    outputs: [{ internalType: 'uint256', name: 'fee', type: 'uint256' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'getNextEntryIndex',
    outputs: [
      { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
      { internalType: 'uint64', name: 'entryIndex', type: 'uint64' },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'uint64', name: 'dbNonce', type: 'uint64' }],
    name: 'getVersionedBatch',
    outputs: [{ internalType: 'bytes', name: 'versionedBatch', type: 'bytes' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'uint64', name: 'dstChainId', type: 'uint64' },
      { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
      { internalType: 'address[]', name: 'srcModules', type: 'address[]' },
    ],
    name: 'requestBatchVerification',
    outputs: [],
    stateMutability: 'payable',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'bytes', name: 'versionedBatch', type: 'bytes' }],
    name: 'verifyRemoteBatch',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'bytes32', name: 'dataHash', type: 'bytes32' }],
    name: 'writeEntry',
    outputs: [
      { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
      { internalType: 'uint64', name: 'entryIndex', type: 'uint64' },
    ],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'uint64', name: 'dstChainId', type: 'uint64' },
      { internalType: 'bytes32', name: 'dataHash', type: 'bytes32' },
      { internalType: 'address[]', name: 'srcModules', type: 'address[]' },
    ],
    name: 'writeEntryWithVerification',
    outputs: [
      { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
      { internalType: 'uint64', name: 'entryIndex', type: 'uint64' },
    ],
    stateMutability: 'payable',
    type: 'function',
  },
] as const
