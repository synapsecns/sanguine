import { fetchPortfolioBalances } from '@/utils/actions/fetchPortfolioBalances'
import { sortByTokenBalance } from '@/utils/sortTokens'

jest.mock('../utils/sortTokens', () => ({
  sortByTokenBalance: jest.fn().mockResolvedValue([]),
}))
jest.mock('../constants/tokens', () => ({
  BRIDGABLE_TOKENS: { 1: [], 1337: [], 999: [] },
  POOLS_BY_CHAIN: {},
  NON_BRIDGEABLE_GAS_TOKENS: {},
}))

beforeEach(() => jest.clearAllMocks())

it('fetches SYN EVM balances without querying the virtual HyperCore chain', async () => {
  const result = await fetchPortfolioBalances(
    '0x1111111111111111111111111111111111111111'
  )
  expect(
    jest.mocked(sortByTokenBalance).mock.calls.map((args) => args[1])
  ).toEqual([1, 999])
  expect(Object.keys(result.balances)).toEqual(['1', '999'])
})

it('skips EVM balance queries when refreshing a HyperCore destination', async () => {
  const result = await fetchPortfolioBalances(
    '0x1111111111111111111111111111111111111111',
    1337
  )
  expect(sortByTokenBalance).not.toHaveBeenCalled()
  expect(result.balances).toEqual({})
})
