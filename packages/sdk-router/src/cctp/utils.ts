import {
  Network,
  Chain,
  ChainId as WormholeChainId,
  chain as _chain,
  nativeChainIds as _nativeChainIds,
  relayInstructionsLayout,
  serializeLayout as _serializeLayout,
} from '@wormhole-foundation/sdk-connect'

const DEFAULT_GAS_LIMIT = 250000n
const DEFAULT_MSG_VALUE = 0n

// Type assertions to work around tsdx's TypeScript type resolution issues
const serializeLayout = _serializeLayout as unknown as (
  layout: unknown,
  data: unknown
) => Uint8Array

const nativeChainIds = _nativeChainIds as unknown as {
  platformNativeChainIdToNetworkChain: (
    platform: string,
    chainId: bigint
  ) => [Network, Chain]
}

const chain = _chain as unknown as {
  chainToChainId: (chain: Chain) => WormholeChainId
}

export const evmChainIdToWormholeChainId = (
  chainId: number
): WormholeChainId => {
  const [, c] = nativeChainIds.platformNativeChainIdToNetworkChain(
    'Evm',
    BigInt(chainId)
  )
  return chain.chainToChainId(c)
}

export const serializeGasInstruction = ({
  gasLimit = DEFAULT_GAS_LIMIT,
  msgValue = DEFAULT_MSG_VALUE,
}: {
  gasLimit?: bigint
  msgValue?: bigint
}): string => {
  const gasInstruction = serializeLayout(relayInstructionsLayout, {
    requests: [
      {
        request: {
          type: 'GasInstruction',
          gasLimit,
          msgValue,
        },
      },
    ],
  })
  return '0x' + Buffer.from(gasInstruction).toString('hex')
}

export const addressToBytes32 = (address: string): string => {
  return '0x' + address.slice(2).padStart(64, '0')
}
