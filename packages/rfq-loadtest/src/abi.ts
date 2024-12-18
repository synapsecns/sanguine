export const ABI = {
  erc20min: [
    {
      "constant": true,
      "inputs": [
        {
          "name": "owner",
          "type": "address"
        },
        {
          "name": "spender",
          "type": "address"
        }
      ],
      "name": "allowance",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "spender",
          "type": "address"
        },
        {
          "name": "value",
          "type": "uint256"
        }
      ],
      "name": "approve",
      "outputs": [
        {
          "name": "",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "account",
          "type": "address"
        }
      ],
      "name": "balanceOf",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    }
  ],
  sinRouter: [{"inputs":[{"internalType":"address","name":"target","type":"address"}],"name":"AddressEmptyCode","type":"error"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"AddressInsufficientBalance","type":"error"},{"inputs":[],"name":"FailedInnerCall","type":"error"},{"inputs":[],"name":"SIR__AmountInsufficient","type":"error"},{"inputs":[],"name":"SIR__DeadlineExceeded","type":"error"},{"inputs":[],"name":"SIR__MsgValueIncorrect","type":"error"},{"inputs":[],"name":"SIR__StepsNotProvided","type":"error"},{"inputs":[],"name":"SIR__TokenNotContract","type":"error"},{"inputs":[],"name":"SIR__UnspentFunds","type":"error"},{"inputs":[],"name":"SIR__ZapIncorrectReturnValue","type":"error"},{"inputs":[],"name":"SIR__ZapNoReturnValue","type":"error"},{"inputs":[{"internalType":"address","name":"token","type":"address"}],"name":"SafeERC20FailedOperation","type":"error"},{"inputs":[],"name":"NATIVE_GAS_TOKEN","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"zapRecipient","type":"address"},{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"minLastStepAmountIn","type":"uint256"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"components":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"},{"internalType":"uint256","name":"msgValue","type":"uint256"},{"internalType":"bytes","name":"zapData","type":"bytes"}],"internalType":"struct ISynapseIntentRouter.StepParams[]","name":"steps","type":"tuple[]"}],"name":"completeIntent","outputs":[],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"address","name":"zapRecipient","type":"address"},{"internalType":"uint256","name":"amountIn","type":"uint256"},{"internalType":"uint256","name":"minLastStepAmountIn","type":"uint256"},{"internalType":"uint256","name":"deadline","type":"uint256"},{"components":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"},{"internalType":"uint256","name":"msgValue","type":"uint256"},{"internalType":"bytes","name":"zapData","type":"bytes"}],"internalType":"struct ISynapseIntentRouter.StepParams[]","name":"steps","type":"tuple[]"}],"name":"completeIntentWithBalanceChecks","outputs":[],"stateMutability":"payable","type":"function"}],
  fastBridgeV2: [
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "defaultAdmin",
          "type": "address"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "inputs": [],
      "name": "AccessControlBadConfirmation",
      "type": "error"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        },
        {
          "internalType": "bytes32",
          "name": "neededRole",
          "type": "bytes32"
        }
      ],
      "name": "AccessControlUnauthorizedAccount",
      "type": "error"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "target",
          "type": "address"
        }
      ],
      "name": "AddressEmptyCode",
      "type": "error"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        }
      ],
      "name": "AddressInsufficientBalance",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "AmountIncorrect",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "BridgeTransactionV2__InvalidEncodedTx",
      "type": "error"
    },
    {
      "inputs": [
        {
          "internalType": "uint16",
          "name": "version",
          "type": "uint16"
        }
      ],
      "name": "BridgeTransactionV2__UnsupportedVersion",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "CancelDelayBelowMin",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "ChainIncorrect",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "DeadlineExceeded",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "DeadlineNotExceeded",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "DeadlineTooShort",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "DisputePeriodNotPassed",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "DisputePeriodPassed",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "ExclusivityParamsIncorrect",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "ExclusivityPeriodNotPassed",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "FailedInnerCall",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "FeeRateAboveMax",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "MsgValueIncorrect",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "MulticallTarget__UndeterminedRevert",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "RecipientIncorrectReturnValue",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "RecipientNoReturnValue",
      "type": "error"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "token",
          "type": "address"
        }
      ],
      "name": "SafeERC20FailedOperation",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "SenderIncorrect",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "StatusIncorrect",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "TokenNotContract",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "TransactionRelayed",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "ZapDataLengthAboveMax",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "ZapNativeNotSupported",
      "type": "error"
    },
    {
      "inputs": [],
      "name": "ZeroAddress",
      "type": "error"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "transactionId",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "relayer",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "address",
          "name": "token",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "BridgeDepositClaimed",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "transactionId",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "address",
          "name": "token",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "BridgeDepositRefunded",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "transactionId",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "relayer",
          "type": "address"
        }
      ],
      "name": "BridgeProofDisputed",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "transactionId",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "relayer",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "bytes32",
          "name": "transactionHash",
          "type": "bytes32"
        }
      ],
      "name": "BridgeProofProvided",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "transactionId",
          "type": "bytes32"
        },
        {
          "indexed": false,
          "internalType": "bytes",
          "name": "quoteId",
          "type": "bytes"
        }
      ],
      "name": "BridgeQuoteDetails",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "transactionId",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "relayer",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "to",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint32",
          "name": "originChainId",
          "type": "uint32"
        },
        {
          "indexed": false,
          "internalType": "address",
          "name": "originToken",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "address",
          "name": "destToken",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "originAmount",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "destAmount",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "chainGasAmount",
          "type": "uint256"
        }
      ],
      "name": "BridgeRelayed",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "transactionId",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "sender",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "bytes",
          "name": "request",
          "type": "bytes"
        },
        {
          "indexed": false,
          "internalType": "uint32",
          "name": "destChainId",
          "type": "uint32"
        },
        {
          "indexed": false,
          "internalType": "address",
          "name": "originToken",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "address",
          "name": "destToken",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "originAmount",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "destAmount",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "bool",
          "name": "sendChainGas",
          "type": "bool"
        }
      ],
      "name": "BridgeRequested",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "oldCancelDelay",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "newCancelDelay",
          "type": "uint256"
        }
      ],
      "name": "CancelDelayUpdated",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "oldFeeRate",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "newFeeRate",
          "type": "uint256"
        }
      ],
      "name": "FeeRateUpdated",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "internalType": "address",
          "name": "token",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "address",
          "name": "recipient",
          "type": "address"
        },
        {
          "indexed": false,
          "internalType": "uint256",
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "FeesSwept",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "role",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "previousAdminRole",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "newAdminRole",
          "type": "bytes32"
        }
      ],
      "name": "RoleAdminChanged",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "role",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "account",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "sender",
          "type": "address"
        }
      ],
      "name": "RoleGranted",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "bytes32",
          "name": "role",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "account",
          "type": "address"
        },
        {
          "indexed": true,
          "internalType": "address",
          "name": "sender",
          "type": "address"
        }
      ],
      "name": "RoleRevoked",
      "type": "event"
    },
    {
      "inputs": [],
      "name": "CANCELER_ROLE",
      "outputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "DEFAULT_ADMIN_ROLE",
      "outputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "DEFAULT_CANCEL_DELAY",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "DISPUTE_PERIOD",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "FEE_BPS",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "FEE_RATE_MAX",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "GOVERNOR_ROLE",
      "outputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "GUARD_ROLE",
      "outputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "MAX_ZAP_DATA_LENGTH",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "MIN_CANCEL_DELAY",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "MIN_DEADLINE_PERIOD",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "NATIVE_GAS_TOKEN",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "PROVER_ROLE",
      "outputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "QUOTER_ROLE",
      "outputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "components": [
            {
              "internalType": "uint32",
              "name": "dstChainId",
              "type": "uint32"
            },
            {
              "internalType": "address",
              "name": "sender",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "to",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "originToken",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "destToken",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "originAmount",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "destAmount",
              "type": "uint256"
            },
            {
              "internalType": "bool",
              "name": "sendChainGas",
              "type": "bool"
            },
            {
              "internalType": "uint256",
              "name": "deadline",
              "type": "uint256"
            }
          ],
          "internalType": "struct IFastBridge.BridgeParams",
          "name": "params",
          "type": "tuple"
        }
      ],
      "name": "bridge",
      "outputs": [],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "transactionId",
          "type": "bytes32"
        }
      ],
      "name": "bridgeProofs",
      "outputs": [
        {
          "internalType": "uint96",
          "name": "timestamp",
          "type": "uint96"
        },
        {
          "internalType": "address",
          "name": "relayer",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "name": "bridgeRelayDetails",
      "outputs": [
        {
          "internalType": "uint48",
          "name": "blockNumber",
          "type": "uint48"
        },
        {
          "internalType": "uint48",
          "name": "blockTimestamp",
          "type": "uint48"
        },
        {
          "internalType": "address",
          "name": "relayer",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "transactionId",
          "type": "bytes32"
        }
      ],
      "name": "bridgeRelays",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "transactionId",
          "type": "bytes32"
        }
      ],
      "name": "bridgeStatuses",
      "outputs": [
        {
          "internalType": "enum IFastBridgeV2.BridgeStatus",
          "name": "status",
          "type": "uint8"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "name": "bridgeTxDetails",
      "outputs": [
        {
          "internalType": "enum IFastBridgeV2.BridgeStatus",
          "name": "status",
          "type": "uint8"
        },
        {
          "internalType": "uint32",
          "name": "destChainId",
          "type": "uint32"
        },
        {
          "internalType": "uint56",
          "name": "proofBlockTimestamp",
          "type": "uint56"
        },
        {
          "internalType": "address",
          "name": "proofRelayer",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "components": [
            {
              "internalType": "uint32",
              "name": "dstChainId",
              "type": "uint32"
            },
            {
              "internalType": "address",
              "name": "sender",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "to",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "originToken",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "destToken",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "originAmount",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "destAmount",
              "type": "uint256"
            },
            {
              "internalType": "bool",
              "name": "sendChainGas",
              "type": "bool"
            },
            {
              "internalType": "uint256",
              "name": "deadline",
              "type": "uint256"
            }
          ],
          "internalType": "struct IFastBridge.BridgeParams",
          "name": "params",
          "type": "tuple"
        },
        {
          "components": [
            {
              "internalType": "address",
              "name": "quoteRelayer",
              "type": "address"
            },
            {
              "internalType": "int256",
              "name": "quoteExclusivitySeconds",
              "type": "int256"
            },
            {
              "internalType": "bytes",
              "name": "quoteId",
              "type": "bytes"
            },
            {
              "internalType": "uint256",
              "name": "zapNative",
              "type": "uint256"
            },
            {
              "internalType": "bytes",
              "name": "zapData",
              "type": "bytes"
            }
          ],
          "internalType": "struct IFastBridgeV2.BridgeParamsV2",
          "name": "paramsV2",
          "type": "tuple"
        }
      ],
      "name": "bridgeV2",
      "outputs": [],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "transactionId",
          "type": "bytes32"
        },
        {
          "internalType": "address",
          "name": "relayer",
          "type": "address"
        }
      ],
      "name": "canClaim",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "cancelDelay",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes",
          "name": "request",
          "type": "bytes"
        }
      ],
      "name": "cancelV2",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "chainGasAmount",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes",
          "name": "request",
          "type": "bytes"
        },
        {
          "internalType": "address",
          "name": "to",
          "type": "address"
        }
      ],
      "name": "claim",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes",
          "name": "request",
          "type": "bytes"
        }
      ],
      "name": "claimV2",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "deployBlock",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "transactionId",
          "type": "bytes32"
        }
      ],
      "name": "dispute",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes",
          "name": "request",
          "type": "bytes"
        }
      ],
      "name": "getBridgeTransaction",
      "outputs": [
        {
          "components": [
            {
              "internalType": "uint32",
              "name": "originChainId",
              "type": "uint32"
            },
            {
              "internalType": "uint32",
              "name": "destChainId",
              "type": "uint32"
            },
            {
              "internalType": "address",
              "name": "originSender",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "destRecipient",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "originToken",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "destToken",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "originAmount",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "destAmount",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "originFeeAmount",
              "type": "uint256"
            },
            {
              "internalType": "bool",
              "name": "sendChainGas",
              "type": "bool"
            },
            {
              "internalType": "uint256",
              "name": "deadline",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "nonce",
              "type": "uint256"
            }
          ],
          "internalType": "struct IFastBridge.BridgeTransaction",
          "name": "",
          "type": "tuple"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes",
          "name": "request",
          "type": "bytes"
        }
      ],
      "name": "getBridgeTransactionV2",
      "outputs": [
        {
          "components": [
            {
              "internalType": "uint32",
              "name": "originChainId",
              "type": "uint32"
            },
            {
              "internalType": "uint32",
              "name": "destChainId",
              "type": "uint32"
            },
            {
              "internalType": "address",
              "name": "originSender",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "destRecipient",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "originToken",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "destToken",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "originAmount",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "destAmount",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "originFeeAmount",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "deadline",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "nonce",
              "type": "uint256"
            },
            {
              "internalType": "address",
              "name": "exclusivityRelayer",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "exclusivityEndTime",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "zapNative",
              "type": "uint256"
            },
            {
              "internalType": "bytes",
              "name": "zapData",
              "type": "bytes"
            }
          ],
          "internalType": "struct IFastBridgeV2.BridgeTransactionV2",
          "name": "",
          "type": "tuple"
        }
      ],
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "role",
          "type": "bytes32"
        }
      ],
      "name": "getRoleAdmin",
      "outputs": [
        {
          "internalType": "bytes32",
          "name": "",
          "type": "bytes32"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "role",
          "type": "bytes32"
        },
        {
          "internalType": "uint256",
          "name": "index",
          "type": "uint256"
        }
      ],
      "name": "getRoleMember",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "role",
          "type": "bytes32"
        }
      ],
      "name": "getRoleMemberCount",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "role",
          "type": "bytes32"
        },
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        }
      ],
      "name": "grantRole",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "role",
          "type": "bytes32"
        },
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        }
      ],
      "name": "hasRole",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes[]",
          "name": "data",
          "type": "bytes[]"
        },
        {
          "internalType": "bool",
          "name": "ignoreReverts",
          "type": "bool"
        }
      ],
      "name": "multicallNoResults",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes[]",
          "name": "data",
          "type": "bytes[]"
        },
        {
          "internalType": "bool",
          "name": "ignoreReverts",
          "type": "bool"
        }
      ],
      "name": "multicallWithResults",
      "outputs": [
        {
          "components": [
            {
              "internalType": "bool",
              "name": "success",
              "type": "bool"
            },
            {
              "internalType": "bytes",
              "name": "returnData",
              "type": "bytes"
            }
          ],
          "internalType": "struct IMulticallTarget.Result[]",
          "name": "results",
          "type": "tuple[]"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "nonce",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "protocolFeeRate",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "protocolFees",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes",
          "name": "request",
          "type": "bytes"
        },
        {
          "internalType": "bytes32",
          "name": "destTxHash",
          "type": "bytes32"
        }
      ],
      "name": "prove",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "transactionId",
          "type": "bytes32"
        },
        {
          "internalType": "bytes32",
          "name": "destTxHash",
          "type": "bytes32"
        },
        {
          "internalType": "address",
          "name": "relayer",
          "type": "address"
        }
      ],
      "name": "proveV2",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes",
          "name": "request",
          "type": "bytes"
        }
      ],
      "name": "refund",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes",
          "name": "request",
          "type": "bytes"
        }
      ],
      "name": "relay",
      "outputs": [],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes",
          "name": "request",
          "type": "bytes"
        },
        {
          "internalType": "address",
          "name": "relayer",
          "type": "address"
        }
      ],
      "name": "relayV2",
      "outputs": [],
      "stateMutability": "payable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "role",
          "type": "bytes32"
        },
        {
          "internalType": "address",
          "name": "callerConfirmation",
          "type": "address"
        }
      ],
      "name": "renounceRole",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes32",
          "name": "role",
          "type": "bytes32"
        },
        {
          "internalType": "address",
          "name": "account",
          "type": "address"
        }
      ],
      "name": "revokeRole",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "name": "senderNonces",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "newCancelDelay",
          "type": "uint256"
        }
      ],
      "name": "setCancelDelay",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "newFeeRate",
          "type": "uint256"
        }
      ],
      "name": "setProtocolFeeRate",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes4",
          "name": "interfaceId",
          "type": "bytes4"
        }
      ],
      "name": "supportsInterface",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "token",
          "type": "address"
        },
        {
          "internalType": "address",
          "name": "recipient",
          "type": "address"
        }
      ],
      "name": "sweepProtocolFees",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ],






/* fastBridgeV1 - deprecated */

  fastBridgeV1: [
    {
      inputs: [{ internalType: 'address', name: '_owner', type: 'address' }],
      stateMutability: 'nonpayable',
      type: 'constructor',
    },
    { inputs: [], name: 'AccessControlBadConfirmation', type: 'error' },
    {
      inputs: [
        { internalType: 'address', name: 'account', type: 'address' },
        { internalType: 'bytes32', name: 'neededRole', type: 'bytes32' },
      ],
      name: 'AccessControlUnauthorizedAccount',
      type: 'error',
    },
    {
      inputs: [{ internalType: 'address', name: 'target', type: 'address' }],
      name: 'AddressEmptyCode',
      type: 'error',
    },
    {
      inputs: [{ internalType: 'address', name: 'account', type: 'address' }],
      name: 'AddressInsufficientBalance',
      type: 'error',
    },
    { inputs: [], name: 'AmountIncorrect', type: 'error' },
    { inputs: [], name: 'ChainIncorrect', type: 'error' },
    { inputs: [], name: 'DeadlineExceeded', type: 'error' },
    { inputs: [], name: 'DeadlineNotExceeded', type: 'error' },
    { inputs: [], name: 'DeadlineTooShort', type: 'error' },
    { inputs: [], name: 'DisputePeriodNotPassed', type: 'error' },
    { inputs: [], name: 'DisputePeriodPassed', type: 'error' },
    { inputs: [], name: 'FailedInnerCall', type: 'error' },
    { inputs: [], name: 'MsgValueIncorrect', type: 'error' },
    {
      inputs: [{ internalType: 'address', name: 'token', type: 'address' }],
      name: 'SafeERC20FailedOperation',
      type: 'error',
    },
    { inputs: [], name: 'SenderIncorrect', type: 'error' },
    { inputs: [], name: 'StatusIncorrect', type: 'error' },
    { inputs: [], name: 'TokenNotContract', type: 'error' },
    { inputs: [], name: 'TransactionRelayed', type: 'error' },
    { inputs: [], name: 'ZeroAddress', type: 'error' },
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
          indexed: true,
          internalType: 'address',
          name: 'relayer',
          type: 'address',
        },
        { indexed: true, internalType: 'address', name: 'to', type: 'address' },
        {
          indexed: false,
          internalType: 'address',
          name: 'token',
          type: 'address',
        },
        {
          indexed: false,
          internalType: 'uint256',
          name: 'amount',
          type: 'uint256',
        },
      ],
      name: 'BridgeDepositClaimed',
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
        { indexed: true, internalType: 'address', name: 'to', type: 'address' },
        {
          indexed: false,
          internalType: 'address',
          name: 'token',
          type: 'address',
        },
        {
          indexed: false,
          internalType: 'uint256',
          name: 'amount',
          type: 'uint256',
        },
      ],
      name: 'BridgeDepositRefunded',
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
          indexed: true,
          internalType: 'address',
          name: 'relayer',
          type: 'address',
        },
      ],
      name: 'BridgeProofDisputed',
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
          indexed: true,
          internalType: 'address',
          name: 'relayer',
          type: 'address',
        },
        {
          indexed: false,
          internalType: 'bytes32',
          name: 'transactionHash',
          type: 'bytes32',
        },
      ],
      name: 'BridgeProofProvided',
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
          indexed: true,
          internalType: 'address',
          name: 'relayer',
          type: 'address',
        },
        { indexed: true, internalType: 'address', name: 'to', type: 'address' },
        {
          indexed: false,
          internalType: 'uint32',
          name: 'originChainId',
          type: 'uint32',
        },
        {
          indexed: false,
          internalType: 'address',
          name: 'originToken',
          type: 'address',
        },
        {
          indexed: false,
          internalType: 'address',
          name: 'destToken',
          type: 'address',
        },
        {
          indexed: false,
          internalType: 'uint256',
          name: 'originAmount',
          type: 'uint256',
        },
        {
          indexed: false,
          internalType: 'uint256',
          name: 'destAmount',
          type: 'uint256',
        },
        {
          indexed: false,
          internalType: 'uint256',
          name: 'chainGasAmount',
          type: 'uint256',
        },
      ],
      name: 'BridgeRelayed',
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
          indexed: true,
          internalType: 'address',
          name: 'sender',
          type: 'address',
        },
        {
          indexed: false,
          internalType: 'bytes',
          name: 'request',
          type: 'bytes',
        },
        {
          indexed: false,
          internalType: 'uint32',
          name: 'destChainId',
          type: 'uint32',
        },
        {
          indexed: false,
          internalType: 'address',
          name: 'originToken',
          type: 'address',
        },
        {
          indexed: false,
          internalType: 'address',
          name: 'destToken',
          type: 'address',
        },
        {
          indexed: false,
          internalType: 'uint256',
          name: 'originAmount',
          type: 'uint256',
        },
        {
          indexed: false,
          internalType: 'uint256',
          name: 'destAmount',
          type: 'uint256',
        },
        {
          indexed: false,
          internalType: 'bool',
          name: 'sendChainGas',
          type: 'bool',
        },
      ],
      name: 'BridgeRequested',
      type: 'event',
    },
    {
      anonymous: false,
      inputs: [
        {
          indexed: false,
          internalType: 'uint256',
          name: 'oldChainGasAmount',
          type: 'uint256',
        },
        {
          indexed: false,
          internalType: 'uint256',
          name: 'newChainGasAmount',
          type: 'uint256',
        },
      ],
      name: 'ChainGasAmountUpdated',
      type: 'event',
    },
    {
      anonymous: false,
      inputs: [
        {
          indexed: false,
          internalType: 'uint256',
          name: 'oldFeeRate',
          type: 'uint256',
        },
        {
          indexed: false,
          internalType: 'uint256',
          name: 'newFeeRate',
          type: 'uint256',
        },
      ],
      name: 'FeeRateUpdated',
      type: 'event',
    },
    {
      anonymous: false,
      inputs: [
        {
          indexed: false,
          internalType: 'address',
          name: 'token',
          type: 'address',
        },
        {
          indexed: false,
          internalType: 'address',
          name: 'recipient',
          type: 'address',
        },
        {
          indexed: false,
          internalType: 'uint256',
          name: 'amount',
          type: 'uint256',
        },
      ],
      name: 'FeesSwept',
      type: 'event',
    },
    {
      anonymous: false,
      inputs: [
        {
          indexed: true,
          internalType: 'bytes32',
          name: 'role',
          type: 'bytes32',
        },
        {
          indexed: true,
          internalType: 'bytes32',
          name: 'previousAdminRole',
          type: 'bytes32',
        },
        {
          indexed: true,
          internalType: 'bytes32',
          name: 'newAdminRole',
          type: 'bytes32',
        },
      ],
      name: 'RoleAdminChanged',
      type: 'event',
    },
    {
      anonymous: false,
      inputs: [
        {
          indexed: true,
          internalType: 'bytes32',
          name: 'role',
          type: 'bytes32',
        },
        {
          indexed: true,
          internalType: 'address',
          name: 'account',
          type: 'address',
        },
        {
          indexed: true,
          internalType: 'address',
          name: 'sender',
          type: 'address',
        },
      ],
      name: 'RoleGranted',
      type: 'event',
    },
    {
      anonymous: false,
      inputs: [
        {
          indexed: true,
          internalType: 'bytes32',
          name: 'role',
          type: 'bytes32',
        },
        {
          indexed: true,
          internalType: 'address',
          name: 'account',
          type: 'address',
        },
        {
          indexed: true,
          internalType: 'address',
          name: 'sender',
          type: 'address',
        },
      ],
      name: 'RoleRevoked',
      type: 'event',
    },
    {
      inputs: [],
      name: 'DEFAULT_ADMIN_ROLE',
      outputs: [{ internalType: 'bytes32', name: '', type: 'bytes32' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [],
      name: 'DISPUTE_PERIOD',
      outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [],
      name: 'FEE_BPS',
      outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [],
      name: 'FEE_RATE_MAX',
      outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [],
      name: 'GOVERNOR_ROLE',
      outputs: [{ internalType: 'bytes32', name: '', type: 'bytes32' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [],
      name: 'GUARD_ROLE',
      outputs: [{ internalType: 'bytes32', name: '', type: 'bytes32' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [],
      name: 'MIN_DEADLINE_PERIOD',
      outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [],
      name: 'REFUNDER_ROLE',
      outputs: [{ internalType: 'bytes32', name: '', type: 'bytes32' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [],
      name: 'REFUND_DELAY',
      outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [],
      name: 'RELAYER_ROLE',
      outputs: [{ internalType: 'bytes32', name: '', type: 'bytes32' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [
        {
          components: [
            { internalType: 'uint32', name: 'dstChainId', type: 'uint32' },
            { internalType: 'address', name: 'sender', type: 'address' },
            { internalType: 'address', name: 'to', type: 'address' },
            { internalType: 'address', name: 'originToken', type: 'address' },
            { internalType: 'address', name: 'destToken', type: 'address' },
            { internalType: 'uint256', name: 'originAmount', type: 'uint256' },
            { internalType: 'uint256', name: 'destAmount', type: 'uint256' },
            { internalType: 'bool', name: 'sendChainGas', type: 'bool' },
            { internalType: 'uint256', name: 'deadline', type: 'uint256' },
          ],
          internalType: 'struct IFastBridge.BridgeParams',
          name: 'params',
          type: 'tuple',
        },
      ],
      name: 'bridge',
      outputs: [],
      stateMutability: 'payable',
      type: 'function',
    },
    {
      inputs: [{ internalType: 'bytes32', name: '', type: 'bytes32' }],
      name: 'bridgeProofs',
      outputs: [
        { internalType: 'uint96', name: 'timestamp', type: 'uint96' },
        { internalType: 'address', name: 'relayer', type: 'address' },
      ],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [{ internalType: 'bytes32', name: '', type: 'bytes32' }],
      name: 'bridgeRelays',
      outputs: [{ internalType: 'bool', name: '', type: 'bool' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [{ internalType: 'bytes32', name: '', type: 'bytes32' }],
      name: 'bridgeStatuses',
      outputs: [
        {
          internalType: 'enum FastBridge.BridgeStatus',
          name: '',
          type: 'uint8',
        },
      ],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [
        { internalType: 'bytes32', name: 'transactionId', type: 'bytes32' },
        { internalType: 'address', name: 'relayer', type: 'address' },
      ],
      name: 'canClaim',
      outputs: [{ internalType: 'bool', name: '', type: 'bool' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [],
      name: 'chainGasAmount',
      outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [
        { internalType: 'bytes', name: 'request', type: 'bytes' },
        { internalType: 'address', name: 'to', type: 'address' },
      ],
      name: 'claim',
      outputs: [],
      stateMutability: 'nonpayable',
      type: 'function',
    },
    {
      inputs: [],
      name: 'deployBlock',
      outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [
        { internalType: 'bytes32', name: 'transactionId', type: 'bytes32' },
      ],
      name: 'dispute',
      outputs: [],
      stateMutability: 'nonpayable',
      type: 'function',
    },
    {
      inputs: [{ internalType: 'bytes', name: 'request', type: 'bytes' }],
      name: 'getBridgeTransaction',
      outputs: [
        {
          components: [
            { internalType: 'uint32', name: 'originChainId', type: 'uint32' },
            { internalType: 'uint32', name: 'destChainId', type: 'uint32' },
            { internalType: 'address', name: 'originSender', type: 'address' },
            { internalType: 'address', name: 'destRecipient', type: 'address' },
            { internalType: 'address', name: 'originToken', type: 'address' },
            { internalType: 'address', name: 'destToken', type: 'address' },
            { internalType: 'uint256', name: 'originAmount', type: 'uint256' },
            { internalType: 'uint256', name: 'destAmount', type: 'uint256' },
            {
              internalType: 'uint256',
              name: 'originFeeAmount',
              type: 'uint256',
            },
            { internalType: 'bool', name: 'sendChainGas', type: 'bool' },
            { internalType: 'uint256', name: 'deadline', type: 'uint256' },
            { internalType: 'uint256', name: 'nonce', type: 'uint256' },
          ],
          internalType: 'struct IFastBridge.BridgeTransaction',
          name: '',
          type: 'tuple',
        },
      ],
      stateMutability: 'pure',
      type: 'function',
    },
    {
      inputs: [{ internalType: 'bytes32', name: 'role', type: 'bytes32' }],
      name: 'getRoleAdmin',
      outputs: [{ internalType: 'bytes32', name: '', type: 'bytes32' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [
        { internalType: 'bytes32', name: 'role', type: 'bytes32' },
        { internalType: 'uint256', name: 'index', type: 'uint256' },
      ],
      name: 'getRoleMember',
      outputs: [{ internalType: 'address', name: '', type: 'address' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [{ internalType: 'bytes32', name: 'role', type: 'bytes32' }],
      name: 'getRoleMemberCount',
      outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [
        { internalType: 'bytes32', name: 'role', type: 'bytes32' },
        { internalType: 'address', name: 'account', type: 'address' },
      ],
      name: 'grantRole',
      outputs: [],
      stateMutability: 'nonpayable',
      type: 'function',
    },
    {
      inputs: [
        { internalType: 'bytes32', name: 'role', type: 'bytes32' },
        { internalType: 'address', name: 'account', type: 'address' },
      ],
      name: 'hasRole',
      outputs: [{ internalType: 'bool', name: '', type: 'bool' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [],
      name: 'nonce',
      outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [],
      name: 'protocolFeeRate',
      outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [{ internalType: 'address', name: '', type: 'address' }],
      name: 'protocolFees',
      outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [
        { internalType: 'bytes', name: 'request', type: 'bytes' },
        { internalType: 'bytes32', name: 'destTxHash', type: 'bytes32' },
      ],
      name: 'prove',
      outputs: [],
      stateMutability: 'nonpayable',
      type: 'function',
    },
    {
      inputs: [{ internalType: 'bytes', name: 'request', type: 'bytes' }],
      name: 'refund',
      outputs: [],
      stateMutability: 'nonpayable',
      type: 'function',
    },
    {
      inputs: [{ internalType: 'bytes', name: 'request', type: 'bytes' }],
      name: 'relay',
      outputs: [],
      stateMutability: 'payable',
      type: 'function',
    },
    {
      inputs: [
        { internalType: 'bytes32', name: 'role', type: 'bytes32' },
        {
          internalType: 'address',
          name: 'callerConfirmation',
          type: 'address',
        },
      ],
      name: 'renounceRole',
      outputs: [],
      stateMutability: 'nonpayable',
      type: 'function',
    },
    {
      inputs: [
        { internalType: 'bytes32', name: 'role', type: 'bytes32' },
        { internalType: 'address', name: 'account', type: 'address' },
      ],
      name: 'revokeRole',
      outputs: [],
      stateMutability: 'nonpayable',
      type: 'function',
    },
    {
      inputs: [
        { internalType: 'uint256', name: 'newChainGasAmount', type: 'uint256' },
      ],
      name: 'setChainGasAmount',
      outputs: [],
      stateMutability: 'nonpayable',
      type: 'function',
    },
    {
      inputs: [
        { internalType: 'uint256', name: 'newFeeRate', type: 'uint256' },
      ],
      name: 'setProtocolFeeRate',
      outputs: [],
      stateMutability: 'nonpayable',
      type: 'function',
    },
    {
      inputs: [{ internalType: 'bytes4', name: 'interfaceId', type: 'bytes4' }],
      name: 'supportsInterface',
      outputs: [{ internalType: 'bool', name: '', type: 'bool' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [
        { internalType: 'address', name: 'token', type: 'address' },
        { internalType: 'address', name: 'recipient', type: 'address' },
      ],
      name: 'sweepProtocolFees',
      outputs: [],
      stateMutability: 'nonpayable',
      type: 'function',
    },
  ],

  fastRouterV2: [
    {
      inputs: [{ internalType: 'address', name: 'owner_', type: 'address' }],
      stateMutability: 'nonpayable',
      type: 'constructor',
    },
    { inputs: [], name: 'DeadlineExceeded', type: 'error' },
    {
      inputs: [],
      name: 'FastBridgeRouterV2__OriginSenderNotSpecified',
      type: 'error',
    },
    { inputs: [], name: 'InsufficientOutputAmount', type: 'error' },
    { inputs: [], name: 'MsgValueIncorrect', type: 'error' },
    { inputs: [], name: 'PoolNotFound', type: 'error' },
    { inputs: [], name: 'TokenAddressMismatch', type: 'error' },
    { inputs: [], name: 'TokenNotContract', type: 'error' },
    { inputs: [], name: 'TokenNotETH', type: 'error' },
    { inputs: [], name: 'TokensIdentical', type: 'error' },
    {
      anonymous: false,
      inputs: [
        {
          indexed: false,
          internalType: 'address',
          name: 'newFastBridge',
          type: 'address',
        },
      ],
      name: 'FastBridgeSet',
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
          internalType: 'address',
          name: 'newSwapQuoter',
          type: 'address',
        },
      ],
      name: 'SwapQuoterSet',
      type: 'event',
    },
    {
      inputs: [],
      name: 'GAS_REBATE_FLAG',
      outputs: [{ internalType: 'bytes1', name: '', type: 'bytes1' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [
        { internalType: 'address', name: 'recipient', type: 'address' },
        { internalType: 'address', name: 'tokenIn', type: 'address' },
        { internalType: 'uint256', name: 'amountIn', type: 'uint256' },
        { internalType: 'address', name: 'tokenOut', type: 'address' },
        { internalType: 'bytes', name: 'rawParams', type: 'bytes' },
      ],
      name: 'adapterSwap',
      outputs: [
        { internalType: 'uint256', name: 'amountOut', type: 'uint256' },
      ],
      stateMutability: 'payable',
      type: 'function',
    },
    {
      inputs: [
        { internalType: 'address', name: 'recipient', type: 'address' },
        { internalType: 'uint256', name: 'chainId', type: 'uint256' },
        { internalType: 'address', name: 'token', type: 'address' },
        { internalType: 'uint256', name: 'amount', type: 'uint256' },
        {
          components: [
            { internalType: 'address', name: 'routerAdapter', type: 'address' },
            { internalType: 'address', name: 'tokenOut', type: 'address' },
            { internalType: 'uint256', name: 'minAmountOut', type: 'uint256' },
            { internalType: 'uint256', name: 'deadline', type: 'uint256' },
            { internalType: 'bytes', name: 'rawParams', type: 'bytes' },
          ],
          internalType: 'struct SwapQuery',
          name: 'originQuery',
          type: 'tuple',
        },
        {
          components: [
            { internalType: 'address', name: 'routerAdapter', type: 'address' },
            { internalType: 'address', name: 'tokenOut', type: 'address' },
            { internalType: 'uint256', name: 'minAmountOut', type: 'uint256' },
            { internalType: 'uint256', name: 'deadline', type: 'uint256' },
            { internalType: 'bytes', name: 'rawParams', type: 'bytes' },
          ],
          internalType: 'struct SwapQuery',
          name: 'destQuery',
          type: 'tuple',
        },
      ],
      name: 'bridge',
      outputs: [],
      stateMutability: 'payable',
      type: 'function',
    },
    {
      inputs: [],
      name: 'fastBridge',
      outputs: [{ internalType: 'address', name: '', type: 'address' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [
        { internalType: 'address', name: 'tokenIn', type: 'address' },
        { internalType: 'address[]', name: 'rfqTokens', type: 'address[]' },
        { internalType: 'uint256', name: 'amountIn', type: 'uint256' },
      ],
      name: 'getOriginAmountOut',
      outputs: [
        {
          components: [
            { internalType: 'address', name: 'routerAdapter', type: 'address' },
            { internalType: 'address', name: 'tokenOut', type: 'address' },
            { internalType: 'uint256', name: 'minAmountOut', type: 'uint256' },
            { internalType: 'uint256', name: 'deadline', type: 'uint256' },
            { internalType: 'bytes', name: 'rawParams', type: 'bytes' },
          ],
          internalType: 'struct SwapQuery[]',
          name: 'originQueries',
          type: 'tuple[]',
        },
      ],
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
      inputs: [
        { internalType: 'address', name: 'fastBridge_', type: 'address' },
      ],
      name: 'setFastBridge',
      outputs: [],
      stateMutability: 'nonpayable',
      type: 'function',
    },
    {
      inputs: [
        { internalType: 'address', name: 'swapQuoter_', type: 'address' },
      ],
      name: 'setSwapQuoter',
      outputs: [],
      stateMutability: 'nonpayable',
      type: 'function',
    },
    {
      inputs: [],
      name: 'swapQuoter',
      outputs: [{ internalType: 'address', name: '', type: 'address' }],
      stateMutability: 'view',
      type: 'function',
    },
    {
      inputs: [{ internalType: 'address', name: 'newOwner', type: 'address' }],
      name: 'transferOwnership',
      outputs: [],
      stateMutability: 'nonpayable',
      type: 'function',
    },
    { stateMutability: 'payable', type: 'receive' },
  ],
}
