import sushiLogo from '@assets/icons/sushi.svg'
import * as CHAINS from '@constants-new/chains/master'

import { Token } from '../../utils/types'
import { MINICHEF_ADDRESSES } from '../minichef'

export const SYN_ETH_SUSHI_TOKEN = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0x4a86c01d67965f8cb3d0aaa2c655705e64097c31',
  },
  decimals: 18,
  symbol: 'SYN/ETH-SLP',
  name: 'SYN/ETH Sushi LP',
  logo: sushiLogo,
  poolName: 'SYN/ETH Sushiswap LP',
  poolId: 0,
  poolType: 'EXTERNAL_LP',
  description: 'The SYN/ETH Sushiswap LP Token',
  priorityRank: 6,
  chainId: CHAINS.ETH.id,
  miniChefAddress: MINICHEF_ADDRESSES[CHAINS.ETH.id],
})

export const ETH_USDC_SUSHI_TOKEN = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0x397ff1542f962076d0bfe58ea045ffa2d347aca0',
  },
  decimals: 18,
  symbol: 'ETH/USDC-SLP',
  name: 'ETH/USDC Sushi LP',
  logo: sushiLogo,
  poolName: 'ETH/USDC Sushiswap LP',
  poolId: 666,
  poolType: 'EXTERNAL_LP',
  description: 'The ETH/USDC Sushiswap LP Token',
  priorityRank: 6,
  chainId: CHAINS.ETH.id,
})
