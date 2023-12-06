import { useBridgeState } from '@/state/slices/bridge/hooks'
import { BridgeState } from '@/state/slices/bridge/reducer'
import { TokenMetaData } from 'types'
import { TokenPopoverSelect } from './TokenPopoverSelect'

const exampleTokens: TokenMetaData[] = [
  {
    tokenAddress: '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
    symbol: 'USDC',
    chainId: 1,
    decimals: 6,
  },
  {
    tokenAddress: '0xaf88d065e77c8cc2239327c5edb3a432268e5831',
    symbol: 'USDC',
    chainId: 42161,
    decimals: 6,
  },
  {
    tokenAddress: '0x6b175474e89094c44da98b954eedeac495271d0f',
    symbol: 'DAI',
    chainId: 1,
    decimals: 18,
  },
  {
    tokenAddress: '0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1',
    symbol: 'DAI',
    chainId: 42161,
    decimals: 18,
  },
  {
    tokenAddress: '0xdac17f958d2ee523a2206206994597c13d831ec7',
    symbol: 'USDT',
    chainId: 1,
    decimals: 6,
  },
  {
    tokenAddress: '0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9',
    symbol: 'USDT',
    chainId: 42161,
    decimals: 6,
  },
]

type Props = {
  label: 'Out' | 'In'
  onChange: (newToken: TokenMetaData) => void
  token: TokenMetaData
}

export function TokenSelect({ label, token, onChange }: Props) {
  const { originChain, destinationChain }: BridgeState = useBridgeState()

  let tokens

  if (label === 'In') {
    tokens = exampleTokens.filter((token) => token.chainId === originChain.id)
  } else {
    tokens = exampleTokens.filter(
      (token) => token.chainId === destinationChain.id
    )
  }

  return (
    <TokenPopoverSelect
      selectedChain={label === 'In' ? originChain : destinationChain}
      options={tokens}
      onSelect={(selected) => {
        onChange(selected)
      }}
      selected={token}
      label={label}
    />
  )
}
