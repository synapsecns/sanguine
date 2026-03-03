import { SupportedChainId } from './chainIds'

export const CCTP_V2_SUPPORTED_CHAIN_IDS: number[] = [
  SupportedChainId.ETH,
  SupportedChainId.AVALANCHE,
  SupportedChainId.OPTIMISM,
  SupportedChainId.ARBITRUM,
  SupportedChainId.BASE,
  SupportedChainId.POLYGON,
]

export const CCTP_V2_DOMAIN_MAP: Record<number, number> = {
  [SupportedChainId.ETH]: 0,
  [SupportedChainId.AVALANCHE]: 1,
  [SupportedChainId.OPTIMISM]: 2,
  [SupportedChainId.ARBITRUM]: 3,
  [SupportedChainId.BASE]: 6,
  [SupportedChainId.POLYGON]: 7,
}

export const CCTP_V2_USDC_ADDRESS_MAP: Record<number, string> = {
  [SupportedChainId.ETH]: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
  [SupportedChainId.AVALANCHE]: '0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E',
  [SupportedChainId.OPTIMISM]: '0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85',
  [SupportedChainId.ARBITRUM]: '0xaf88d065e77c8cC2239327C5EDb3A432268e5831',
  [SupportedChainId.BASE]: '0x833589fCD6EDB6E08f4c7C32D4f71b54bdA02913',
  [SupportedChainId.POLYGON]: '0x3c499c542cEF5E3811e1192ce70d8cc03d5c3359',
}

export const CCTP_V2_TOKEN_MESSENGER_ADDRESS_MAP: Record<number, string> = {
  [SupportedChainId.ETH]: '0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d',
  [SupportedChainId.AVALANCHE]: '0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d',
  [SupportedChainId.OPTIMISM]: '0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d',
  [SupportedChainId.ARBITRUM]: '0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d',
  [SupportedChainId.BASE]: '0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d',
  [SupportedChainId.POLYGON]: '0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d',
}

export const CCTP_V2_FORWARD_SERVICE_FEE_USDC = {
  ETH: 1_250_000,
  NON_ETH: 200_000,
}

export const CCTP_V2_FORWARD_HOOK_DATA =
  '0x636374702d666f72776172640000000000000000000000000000000000000000'

export const CIRCLE_IRIS_API_HOST =
  process.env.CIRCLE_IRIS_API_HOST || 'https://iris-api.circle.com'
