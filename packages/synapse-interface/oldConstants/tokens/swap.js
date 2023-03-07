
import { ChainId } from '@constants/networks'

import * from '@constants/tokens/basic'


export const SWAPABLE_TOKENS = {
  [ChainId.BSC]: [
    BUSD,
    USDT,
    USDC,
    NUSD,
  ],
  [ChainId.ETH]: [
    DAI,
    USDC,
    USDT
  ],
  [ChainId.POLYGON]: [
    DAI,
    USDC,
    USDT,
    NUSD,
  ],
  [ChainId.FANTOM]: [
    USDC,
    USDT,
    NUSD,

    FTMETH,
    NETH,

  ],
  [ChainId.ARBITRUM]: [
    WETH,
    NETH,

    USDC,
    USDT,
    NUSD,

  ],
  [ChainId.AVALANCHE]: [
    DAI,
    USDC,
    USDT,
    NUSD,

    WETHE,
    NETH,
  ],
  [ChainId.HARMONY]: [
    DAI,
    USDC,
    USDT,
    NUSD,

    ONEETH,
    NETH,

    WJEWEL,
    SYNJEWEL,

    SYNAVAX,
    MULTIAVAX
  ],
  [ChainId.AURORA]: [
    USDC,
    USDT,
    NUSD,
  ],
  [ChainId.BOBA]: [
    WETH,
    NETH,

    DAI,
    USDC,
    USDT,
    NUSD
  ],
  [ChainId.OPTIMISM]: [
    WETH,
    NETH,

    USDC,
    NUSD,
  ],
  [ChainId.MOONRIVER]: [
    FRAX,
  ],
  [ChainId.MOONBEAM]: [
    FRAX,
  ],
  [ChainId.METIS]: [
    USDC,
    NUSD,

    METISETH,
    NETH,
  ],
  [ChainId.CRONOS]: [
    USDC,
    NUSD,
  ],
  [ChainId.KLAYTN]: [
    KLAYTN_USDT,
    KLAYTN_oUSDT,
  ],
  [ChainId.DOGECHAIN]: [],
  [ChainId.CANTO]: [
    NOTE,
    NUSD,
    USDC,
    USDT,
    NETH,
    CANTOETH
  ],
}
