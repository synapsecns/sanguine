import { useFromChainListArray } from '@/components/StateManagedBridge/hooks/useFromChainListArray'

jest.mock('../slices/bridge/hooks', () => ({
  useBridgeState: () => ({ fromChainIds: [1, 999] }),
}))
jest.mock('next-intl', () => ({
  useTranslations: () => (key: string) => key,
}))

it('excludes HyperCore from both supported and fallback source-network lists', () => {
  const chains = Object.values(useFromChainListArray()).flat()
  expect(chains.some((chain) => chain.id === 1337)).toBe(false)
  expect(chains.some((chain) => chain.id === 999)).toBe(true)
  expect(chains.some((chain) => chain.id === 1)).toBe(true)
})
