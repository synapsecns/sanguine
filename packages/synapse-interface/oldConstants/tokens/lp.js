import sushiLogo from '@assets/icons/sushi.svg'

import { ChainId } from '@constants/networks'

import { Token } from '@utils/classes/Token'


export const SYN_ETH_SUSHI_TOKEN = new Token({
  addresses: {
    [ChainId.ETH]: '0x4a86c01d67965f8cb3d0aaa2c655705e64097c31',
  },
  decimals:    18,
  symbol:      'SYN/ETH-SLP',
  name:        'SYN/ETH Sushi LP',
  logo:        sushiLogo,
  poolName:    'SYN/ETH Sushiswap LP',
  poolId:      0,
  poolType:    "EXTERNAL_LP",
  description: 'The SYN/ETH Sushiswap LP Token',
})


export const ETH_USDC_SUSHI_TOKEN = new Token({
  addresses: {
    [ChainId.ETH]: '0x397ff1542f962076d0bfe58ea045ffa2d347aca0',
  },
  decimals:    18,
  symbol:      'ETH/USDC-SLP',
  name:        'ETH/USDC Sushi LP',
  logo:        sushiLogo,
  poolName:    'ETH/USDC Sushiswap LP',
  poolId:      666,
  poolType:    "EXTERNAL_LP",
  description: 'The ETH/USDC Sushiswap LP Token',
})