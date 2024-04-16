export const SynapseModuleAbi = [
  {
    inputs: [
      { internalType: 'address', name: 'interchainDB', type: 'address' },
      { internalType: 'address', name: 'owner_', type: 'address' },
    ],
    stateMutability: 'nonpayable',
    type: 'constructor',
  },
  {
    inputs: [{ internalType: 'address', name: 'account', type: 'address' }],
    name: 'AddressInsufficientBalance',
    type: 'error',
  },
  { inputs: [], name: 'FailedInnerCall', type: 'error' },
  {
    inputs: [{ internalType: 'uint64', name: 'chainId', type: 'uint64' }],
    name: 'InterchainModule__IncorrectSourceChainId',
    type: 'error',
  },
  {
    inputs: [
      { internalType: 'uint256', name: 'actual', type: 'uint256' },
      { internalType: 'uint256', name: 'required', type: 'uint256' },
    ],
    name: 'InterchainModule__InsufficientFee',
    type: 'error',
  },
  {
    inputs: [{ internalType: 'address', name: 'caller', type: 'address' }],
    name: 'InterchainModule__NotInterchainDB',
    type: 'error',
  },
  {
    inputs: [{ internalType: 'uint64', name: 'chainId', type: 'uint64' }],
    name: 'InterchainModule__SameChainId',
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
      { internalType: 'uint256', name: 'claimFeeFraction', type: 'uint256' },
    ],
    name: 'SynapseModule__ClaimFeeFractionExceedsMax',
    type: 'error',
  },
  { inputs: [], name: 'SynapseModule__FeeCollectorNotSet', type: 'error' },
  {
    inputs: [{ internalType: 'address', name: 'gasOracle', type: 'address' }],
    name: 'SynapseModule__GasOracleNotContract',
    type: 'error',
  },
  { inputs: [], name: 'SynapseModule__GasOracleNotSet', type: 'error' },
  { inputs: [], name: 'SynapseModule__NoFeesToClaim', type: 'error' },
  {
    inputs: [{ internalType: 'address', name: 'account', type: 'address' }],
    name: 'ThresholdECDSA__AlreadySigner',
    type: 'error',
  },
  {
    inputs: [{ internalType: 'uint256', name: 'length', type: 'uint256' }],
    name: 'ThresholdECDSA__IncorrectSignaturesLength',
    type: 'error',
  },
  {
    inputs: [{ internalType: 'bytes', name: 'signature', type: 'bytes' }],
    name: 'ThresholdECDSA__InvalidSignature',
    type: 'error',
  },
  {
    inputs: [
      { internalType: 'uint256', name: 'provided', type: 'uint256' },
      { internalType: 'uint256', name: 'threshold', type: 'uint256' },
    ],
    name: 'ThresholdECDSA__NotEnoughSignatures',
    type: 'error',
  },
  {
    inputs: [{ internalType: 'address', name: 'account', type: 'address' }],
    name: 'ThresholdECDSA__NotSigner',
    type: 'error',
  },
  {
    inputs: [],
    name: 'ThresholdECDSA__RecoveredSignersNotSorted',
    type: 'error',
  },
  { inputs: [], name: 'ThresholdECDSA__ZeroAddress', type: 'error' },
  { inputs: [], name: 'ThresholdECDSA__ZeroThreshold', type: 'error' },
  { inputs: [], name: 'VersionedPayload__PrecompileFailed', type: 'error' },
  {
    inputs: [
      { internalType: 'bytes', name: 'versionedPayload', type: 'bytes' },
    ],
    name: 'VersionedPayload__TooShort',
    type: 'error',
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
      { indexed: false, internalType: 'bytes', name: 'batch', type: 'bytes' },
      {
        indexed: false,
        internalType: 'bytes32',
        name: 'ethSignedBatchHash',
        type: 'bytes32',
      },
    ],
    name: 'BatchVerificationRequested',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'uint64',
        name: 'srcChainId',
        type: 'uint64',
      },
      { indexed: false, internalType: 'bytes', name: 'batch', type: 'bytes' },
      {
        indexed: false,
        internalType: 'bytes32',
        name: 'ethSignedBatchHash',
        type: 'bytes32',
      },
    ],
    name: 'BatchVerified',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'uint256',
        name: 'claimFeeFraction',
        type: 'uint256',
      },
    ],
    name: 'ClaimFeeFractionSet',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'address',
        name: 'feeCollector',
        type: 'address',
      },
    ],
    name: 'FeeCollectorSet',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'address',
        name: 'feeCollector',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'collectedFees',
        type: 'uint256',
      },
      {
        indexed: false,
        internalType: 'address',
        name: 'claimer',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'claimerFee',
        type: 'uint256',
      },
    ],
    name: 'FeesClaimed',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'uint64',
        name: 'srcChainId',
        type: 'uint64',
      },
      { indexed: false, internalType: 'bytes', name: 'data', type: 'bytes' },
    ],
    name: 'GasDataReceived',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'uint64',
        name: 'dstChainId',
        type: 'uint64',
      },
      { indexed: false, internalType: 'bytes', name: 'data', type: 'bytes' },
    ],
    name: 'GasDataSent',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'address',
        name: 'gasOracle',
        type: 'address',
      },
    ],
    name: 'GasOracleSet',
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
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'uint256',
        name: 'threshold',
        type: 'uint256',
      },
    ],
    name: 'ThresholdSet',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'address',
        name: 'verifier',
        type: 'address',
      },
    ],
    name: 'VerifierAdded',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'address',
        name: 'verifier',
        type: 'address',
      },
    ],
    name: 'VerifierRemoved',
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
        internalType: 'uint256',
        name: 'gasLimit',
        type: 'uint256',
      },
    ],
    name: 'VerifyGasLimitSet',
    type: 'event',
  },
  {
    inputs: [],
    name: 'DEFAULT_VERIFY_GAS_LIMIT',
    outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
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
    inputs: [{ internalType: 'address', name: 'verifier', type: 'address' }],
    name: 'addVerifier',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'address[]', name: 'verifiers', type: 'address[]' },
    ],
    name: 'addVerifiers',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'claimFees',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'feeCollector',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'gasOracle',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'getClaimFeeAmount',
    outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'getClaimFeeFraction',
    outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'uint64', name: 'dstChainId', type: 'uint64' },
      { internalType: 'uint64', name: 'dbNonce', type: 'uint64' },
    ],
    name: 'getModuleFee',
    outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'getThreshold',
    outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'getVerifiers',
    outputs: [{ internalType: 'address[]', name: '', type: 'address[]' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'uint64', name: 'chainId', type: 'uint64' }],
    name: 'getVerifyGasLimit',
    outputs: [{ internalType: 'uint256', name: 'gasLimit', type: 'uint256' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'address', name: 'account', type: 'address' }],
    name: 'isVerifier',
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
    inputs: [{ internalType: 'address', name: 'verifier', type: 'address' }],
    name: 'removeVerifier',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'address[]', name: 'verifiers', type: 'address[]' },
    ],
    name: 'removeVerifiers',
    outputs: [],
    stateMutability: 'nonpayable',
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
      { internalType: 'uint64', name: 'dstChainId', type: 'uint64' },
      { internalType: 'bytes', name: 'versionedBatch', type: 'bytes' },
    ],
    name: 'requestBatchVerification',
    outputs: [],
    stateMutability: 'payable',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'uint256', name: 'claimFeeFraction', type: 'uint256' },
    ],
    name: 'setClaimFeeFraction',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'address', name: 'feeCollector_', type: 'address' },
    ],
    name: 'setFeeCollector',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'address', name: 'gasOracle_', type: 'address' }],
    name: 'setGasOracle',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'uint256', name: 'threshold', type: 'uint256' }],
    name: 'setThreshold',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'uint64', name: 'chainId', type: 'uint64' },
      { internalType: 'uint256', name: 'gasLimit', type: 'uint256' },
    ],
    name: 'setVerifyGasLimit',
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
      { internalType: 'bytes', name: 'encodedBatch', type: 'bytes' },
      { internalType: 'bytes', name: 'signatures', type: 'bytes' },
    ],
    name: 'verifyRemoteBatch',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
] as const
