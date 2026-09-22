import { BigNumber, Contract, providers, utils } from 'ethers'

import { SYN_COMPOSER_ADDRESS, SYN_CORE_TOKEN_INDEX } from '../constants'
import { OftSendParams, UsdtModule } from './usdtModule'

export type SynSendParams = OftSendParams & {
  minAmount?: BigNumber
  hyperCoreRecipient?: string
}

export const SYN_COMPOSER_ABI = [
  'function coreUserExists(address) view returns (tuple(bool exists))',
  'function ERC20_ASSET_BRIDGE() view returns (address)',
  'function quoteHyperCoreAmount(uint64,int8,address,uint256) view returns (tuple(uint256 evm,uint64 core,uint64 coreBalanceAssetBridge))',
]

export class SynModule extends UsdtModule {
  public getSendParamTuple(
    params: SynSendParams
  ): ReturnType<UsdtModule['getSendParamTuple']> {
    return [
      params.toEid,
      utils.hexZeroPad(
        params.hyperCoreRecipient ? SYN_COMPOSER_ADDRESS : params.toRecipient,
        32
      ),
      params.amount,
      params.minAmount ?? 0,
      '0x', // The adapter enforces receive/compose execution options.
      params.hyperCoreRecipient
        ? utils.defaultAbiCoder.encode(
            ['uint256', 'address'],
            [0, params.hyperCoreRecipient]
          )
        : '0x',
      '0x',
    ]
  }
}

export const getSynComposer = (provider: providers.Provider): Contract => {
  return new Contract(SYN_COMPOSER_ADDRESS, SYN_COMPOSER_ABI, provider)
}

export const quoteSynHyperCore = async (
  composer: Contract,
  amount: BigNumber
): Promise<BigNumber> => {
  const assetBridge = await composer.ERC20_ASSET_BRIDGE()
  const result = await composer.quoteHyperCoreAmount(
    SYN_CORE_TOKEN_INDEX,
    10,
    assetBridge,
    amount
  )
  // Destination quotes use native HyperCore units; OFT calldata remains in EVM units.
  return result.core
}
