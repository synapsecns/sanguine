import * as CHAINS from '@constants/chains/master'
import wethLogo from '@assets/icons/weth.svg'
import { Token } from '@types'
import mimLogo from '@assets/icons/mim.svg'

export const WETH = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
    // [CHAINS.ARBITRUM.id]: '0x82af49447d8a07e3bd95bd0d56f35241523fbab1',
    [CHAINS.BOBA.id]: '0xd203De32170130082896b4111eDF825a4774c18E',
    [CHAINS.OPTIMISM.id]: '0x121ab82b49B2BC4c7901CA46B8277962b4350204',
    // [CHAINS.AVALANCHE.id]: '0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab'
  },
  decimals: 18,
  symbol: 'WETH', // SHOULD BE WETH
  name: 'Wrapped ETH',
  logo: wethLogo,
  description: 'ERC-20 Wrapped form of ETH',
  swapableType: 'ETH',
  swapableOn: [CHAINS.ARBITRUM.id, CHAINS.BOBA.id, CHAINS.OPTIMISM.id],
  color: 'sky',
})
export const MIM = new Token({
  addresses: {
    [CHAINS.FANTOM.id]: '0x82f0b8b456c1a451378467398982d4834b6829c1',
    [CHAINS.ARBITRUM.id]: '0xfea7a6a0b346362bf88a9e4a88416b77a57d6c2a',
  },
  decimals: 18,
  symbol: 'MIM',
  name: 'Magic Internet Money',
  logo: mimLogo,
  swapableType: 'USD',
  color: 'indigo',
})
