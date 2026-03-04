import { getWithTimeout } from '../utils'
import { clearCctpV2ApiCache, getBurnUSDCFees } from './api'

jest.mock('../utils', () => ({
  getWithTimeout: jest.fn(),
}))

const mockGetWithTimeout = getWithTimeout as jest.MockedFunction<
  typeof getWithTimeout
>

const SOURCE_DOMAIN_ID = 10_001
const DEST_DOMAIN_ID = 20_001

const responseWithJson = (data: unknown): Response =>
  ({
    json: jest.fn().mockResolvedValue(data),
  } as unknown as Response)

describe('getBurnUSDCFees', () => {
  beforeEach(() => {
    mockGetWithTimeout.mockReset()
    clearCctpV2ApiCache()
  })

  it('accepts non-negative decimal minimumFee values', async () => {
    const feeRows = [
      {
        finalityThreshold: 1000,
        minimumFee: 0.125,
        forwardFee: { low: 0, high: 2 },
      },
    ]
    mockGetWithTimeout.mockResolvedValueOnce(responseWithJson(feeRows))

    await expect(
      getBurnUSDCFees(SOURCE_DOMAIN_ID, DEST_DOMAIN_ID)
    ).resolves.toEqual(feeRows)
  })

  it('rejects invalid fee rows while keeping valid rows', async () => {
    const validRows = [
      {
        finalityThreshold: 1200,
        minimumFee: 0.5,
        forwardFee: { basic: 10 },
      },
      {
        finalityThreshold: 2000,
        minimumFee: 1.75,
      },
    ]
    const feeRows = [
      ...validRows,
      { finalityThreshold: 1200.5, minimumFee: 0.5, forwardFee: { basic: 10 } },
      { finalityThreshold: -1, minimumFee: 0.5, forwardFee: { basic: 10 } },
      { finalityThreshold: 1200, minimumFee: -0.5, forwardFee: { basic: 10 } },
      { finalityThreshold: 1200, minimumFee: '0.5', forwardFee: { basic: 10 } },
      { finalityThreshold: 1200, minimumFee: 0.5, forwardFee: { basic: -1 } },
      { finalityThreshold: 1200, minimumFee: 0.5, forwardFee: { basic: 1.5 } },
      { finalityThreshold: 1200, minimumFee: 0.5, forwardFee: {} },
    ]
    mockGetWithTimeout.mockResolvedValueOnce(responseWithJson(feeRows))

    await expect(
      getBurnUSDCFees(SOURCE_DOMAIN_ID, DEST_DOMAIN_ID)
    ).resolves.toEqual(validRows)
  })

  it('returns null when all fee rows are invalid', async () => {
    const feeRows = [
      { finalityThreshold: 0.5, minimumFee: 0.1, forwardFee: { basic: 10 } },
      { finalityThreshold: 1200, minimumFee: 0.1, forwardFee: { basic: 1.5 } },
    ]
    mockGetWithTimeout.mockResolvedValueOnce(responseWithJson(feeRows))

    await expect(
      getBurnUSDCFees(SOURCE_DOMAIN_ID, DEST_DOMAIN_ID)
    ).resolves.toBeNull()
  })
})
