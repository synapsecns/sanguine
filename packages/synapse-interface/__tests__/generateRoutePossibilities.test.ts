import { expect } from '@jest/globals'

import {
  generateRoutePossibilities,
  getPossibleFromChainIds,
  getPossibleFromChainIdsByFromToken,
  getPossibleFromChainIdsByToChainId,
  getPossibleFromChainIdsByToToken,
  getPossibleFromTokens,
  getPossibleFromTokensByFromChainId,
  getPossibleFromTokensByFromChainIdAndToChainId,
  getPossibleFromTokensByFromToken,
  getPossibleFromTokensByToChainId,
  getPossibleFromTokensByToToken,
  getPossibleToChainIds,
  getPossibleToChainIdsByFromChainId,
  getPossibleToChainIdsByFromToken,
  getPossibleToChainIdsByToToken,
  getPossibleToTokens,
  getPossibleToTokensByFromChainId,
  getPossibleToTokensByFromToken,
  getPossibleToTokensByFromTokenAndToChainId,
  getPossibleToTokensByToChainId,
  getPossibleToTokensByToToken,
} from '@/utils/generateRoutePossibilities'

jest.mock('../constants/existing-bridge-routes', () => ({
  __esModule: true,
  EXISTING_BRIDGE_ROUTES: {
    'GOHM-1': ['GOHM-10', 'GOHM-25', 'GOHM-56'],
    'GOHM-10': ['GOHM-1', 'GOHM-25', 'GOHM-56'],
    'GOHM-25': ['GOHM-1', 'GOHM-10', 'GOHM-56'],
    'GOHM-56': ['GOHM-1', 'GOHM-10', 'GOHM-25'],
    'HIGHSTREET-1': ['HIGHSTREET-56'],
    'HIGHSTREET-56': ['HIGHSTREET-1'],
    'USDC-1': ['USDC-10', 'USDC-25', 'USDC-56'],
    'USDC-10': ['USDC-1', 'USDC-25', 'USDC-56'],
    'USDC-25': ['USDC-1', 'USDC-10', 'USDC-56'],
    'USDC-56': ['USDC-1', 'USDC-10', 'USDC-25'],
    'SYN-1': ['SYN-10', 'SYN-25', 'SYN-56'],
    'SYN-10': ['SYN-1', 'SYN-25', 'SYN-56'],
    'SYN-25': ['SYN-1', 'SYN-10', 'SYN-56'],
    'SYN-56': ['SYN-1', 'SYN-10', 'SYN-25'],
    'XYZ-50': ['XYZ-1'],
    'XYZ-1': ['XYZ-50'],
  },
}))

jest.mock('../utils/flattenPausedTokens', () => ({
  __esModule: true,
  flattenPausedTokens: jest.fn(() => {
    return []
  }),
}))

describe('generateRoutePossibilities', () => {
  it('takes all null values', () => {
    const {
      fromChainId,
      fromToken,
      toChainId,
      toToken,
      fromChainIds,
      fromTokens,
      toChainIds,
      toTokens,
    } = generateRoutePossibilities({
      fromChainId: null,
      fromToken: null,
      toChainId: null,
      toToken: null,
    })

    expect(fromChainId).toEqual(null)
    expect(fromToken).toEqual(null)
    expect(toChainId).toEqual(null)
    expect(toToken).toEqual(null)
    expect(fromChainIds).toEqual([1, 10, 25, 56, 50])
    expect(fromTokens).toEqual([
      'GOHM-1',
      'GOHM-10',
      'GOHM-25',
      'GOHM-56',
      'HIGHSTREET-1',
      'HIGHSTREET-56',
      'USDC-1',
      'USDC-10',
      'USDC-25',
      'USDC-56',
      'SYN-1',
      'SYN-10',
      'SYN-25',
      'SYN-56',
      'XYZ-50',
      'XYZ-1',
    ])
    expect(toChainIds).toEqual([10, 25, 56, 1, 50])
    expect(toTokens).toEqual([
      'GOHM-10',
      'GOHM-25',
      'GOHM-56',
      'GOHM-1',
      'HIGHSTREET-56',
      'HIGHSTREET-1',
      'USDC-10',
      'USDC-25',
      'USDC-56',
      'USDC-1',
      'SYN-10',
      'SYN-25',
      'SYN-56',
      'SYN-1',
      'XYZ-1',
      'XYZ-50',
    ])
  })

  it('only toChainId', () => {
    const {
      fromChainId,
      fromToken,
      toChainId,
      toToken,
      fromChainIds,
      fromTokens,
      toChainIds,
      toTokens,
    } = generateRoutePossibilities({
      fromChainId: null,
      fromToken: null,
      toChainId: 10,
      toToken: null,
    })

    expect(fromChainId).toEqual(null)
    expect(fromToken).toEqual(null)
    expect(toChainId).toEqual(10)
    expect(toToken).toEqual(null)
    expect(fromChainIds).toEqual([1, 25, 56])
    expect(fromTokens).toEqual([
      'GOHM-1',
      'GOHM-25',
      'GOHM-56',
      'USDC-1',
      'USDC-25',
      'USDC-56',
      'SYN-1',
      'SYN-25',
      'SYN-56',
    ])
    expect(toChainIds).toEqual([10, 25, 56, 1, 50])
    expect(toTokens).toEqual(['GOHM-10', 'USDC-10', 'SYN-10'])
  })
})

it('has all inputs', () => {
  const {
    fromChainId,
    fromToken,
    toChainId,
    toToken,
    fromChainIds,
    fromTokens,
    toChainIds,
    toTokens,
  } = generateRoutePossibilities({
    fromChainId: 10,
    fromToken: 'USDC-10',
    toChainId: 25,
    toToken: 'USDC-25',
  })

  expect(fromChainId).toEqual(10)
  expect(fromToken).toEqual('USDC-10')
  expect(toChainId).toEqual(25)
  expect(toToken).toEqual('USDC-25')
  expect(fromChainIds).toEqual([1, 10, 25, 56, 50])
  expect(fromTokens).toEqual(['GOHM-10', 'USDC-10', 'SYN-10'])
  expect(toChainIds).toEqual([1, 25, 56])
  expect(toTokens).toEqual(['USDC-25'])
})

it('no toToken', () => {
  const {
    fromChainId,
    fromToken,
    toChainId,
    toToken,
    fromChainIds,
    fromTokens,
    toChainIds,
    toTokens,
  } = generateRoutePossibilities({
    fromChainId: 10,
    fromToken: 'USDC-10',
    toChainId: 25,
    toToken: null,
  })

  expect(fromChainId).toEqual(10)
  expect(fromToken).toEqual('USDC-10')
  expect(toChainId).toEqual(25)
  expect(toToken).toEqual(null)
  expect(fromChainIds).toEqual([1, 10, 25, 56, 50])
  expect(fromTokens).toEqual(['GOHM-10', 'USDC-10', 'SYN-10'])
  expect(toChainIds).toEqual([1, 25, 56])
  expect(toTokens).toEqual(['USDC-25'])
})

it('only toChainId and toToken', () => {
  const {
    fromChainId,
    fromToken,
    toChainId,
    toToken,
    fromChainIds,
    fromTokens,
    toChainIds,
    toTokens,
  } = generateRoutePossibilities({
    fromChainId: null,
    fromToken: null,
    toChainId: 25,
    toToken: 'GOHM-25',
  })

  expect(fromChainId).toEqual(null)
  expect(fromToken).toEqual(null)
  expect(toChainId).toEqual(25)
  expect(toToken).toEqual('GOHM-25')
  expect(fromChainIds).toEqual([1, 10, 56])
  expect(fromTokens).toEqual(['GOHM-1', 'GOHM-10', 'GOHM-56'])
  expect(toChainIds).toEqual([1, 10, 56])
  expect(toTokens).toEqual(['GOHM-25', 'USDC-25', 'SYN-25'])
})

describe('getPossibleFromChainIds()', () => {
  it('returns all possible from chain ids', () => {
    const fromChainIds = getPossibleFromChainIds()

    expect(fromChainIds).toEqual([1, 10, 25, 56, 50])
  })
})

describe('getPossibleFromChainIdsByFromToken', () => {
  it('returns from chain ids that the fromToken can be sent from', () => {
    const fromChainIds = getPossibleFromChainIdsByFromToken('USDC-25')

    expect(fromChainIds).toEqual([1, 10, 25, 56])
  })
})

describe('getPossibleFromChainIdsByToChainId', () => {
  it('returns from chain ids with tokens that can be sent to a toChainId', () => {
    const fromChainIds = getPossibleFromChainIdsByToChainId(25)

    expect(fromChainIds).toEqual([1, 10, 56])
  })
})

describe('getPossibleFromChainIdsByToToken', () => {
  it('returns fromChainIds that have an end path of the toToken, single', () => {
    const fromChainIds = getPossibleFromChainIdsByToToken('XYZ-50')

    expect(fromChainIds).toEqual([1])
  })

  it('returns fromChainIds that have an end path of the toToken, multiple', () => {
    const fromChainIds = getPossibleFromChainIdsByToToken('USDC-25')

    expect(fromChainIds).toEqual([1, 10, 56])
  })
})

describe('getPossibleFromTokens', () => {
  it('returns all possible from tokens', () => {
    const fromTokens = getPossibleFromTokens()

    expect(fromTokens).toEqual([
      'GOHM-1',
      'GOHM-10',
      'GOHM-25',
      'GOHM-56',
      'HIGHSTREET-1',
      'HIGHSTREET-56',
      'USDC-1',
      'USDC-10',
      'USDC-25',
      'USDC-56',
      'SYN-1',
      'SYN-10',
      'SYN-25',
      'SYN-56',
      'XYZ-50',
      'XYZ-1',
    ])
  })
})

describe('getPossibleFromTokensByFromChainId', () => {
  it('returns all possible from tokens by fromChainId', () => {
    const fromTokens = getPossibleFromTokensByFromChainId(1)

    expect(fromTokens).toEqual([
      'GOHM-1',
      'HIGHSTREET-1',
      'USDC-1',
      'SYN-1',
      'XYZ-1',
    ])
  })
})

describe('getPossibleFromTokensByFromToken', () => {
  it('returns all from tokens that can come from the same chain id as the fromToken', () => {
    const fromTokens = getPossibleFromTokensByFromToken('USDC-1')

    expect(fromTokens).toEqual([
      'GOHM-1',
      'HIGHSTREET-1',
      'USDC-1',
      'SYN-1',
      'XYZ-1',
    ])
  })
})

describe('getPossibleFromTokensByFromChainIdAndToChainId', () => {
  it('returns all possible from tokens that can go fromChainID -> toChainId', () => {
    const fromTokens = getPossibleFromTokensByFromChainIdAndToChainId(1, 56)

    expect(fromTokens).toEqual(['GOHM-1', 'HIGHSTREET-1', 'USDC-1', 'SYN-1'])
  })
})

describe('getPossbileFromTokensByToChainId', () => {
  it('returns all possible from tokens that can go to toChainId, single', () => {
    const fromTokens = getPossibleFromTokensByToChainId(50)

    expect(fromTokens).toEqual(['XYZ-1'])
  })

  it('returns all possible from tokens that can go to toChainId, multiple', () => {
    const fromTokens = getPossibleFromTokensByToChainId(56)

    expect(fromTokens).toEqual([
      'GOHM-1',
      'GOHM-10',
      'GOHM-25',
      'HIGHSTREET-1',
      'USDC-1',
      'USDC-10',
      'USDC-25',
      'SYN-1',
      'SYN-10',
      'SYN-25',
    ])
  })
})

describe('getPossibleFromTokensByToToken', () => {
  it('returns all possible from tokens that can go to toToken, single', () => {
    const fromTokens = getPossibleFromTokensByToToken('XYZ-1')

    expect(fromTokens).toEqual(['XYZ-50'])
  })

  it('returns all possible from tokens that can go to toToken, multiple', () => {
    const fromTokens = getPossibleFromTokensByToToken('USDC-25')

    expect(fromTokens).toEqual(['USDC-1', 'USDC-10', 'USDC-56'])
  })
})

describe('getPossibleToChainIds', () => {
  it('returns all possible toChainIds', () => {
    const toChainIds = getPossibleToChainIds()

    expect(toChainIds.sort()).toEqual([1, 10, 25, 56, 50].sort())
  })
})

describe('getPossibleToChainIdsByFromChainId', () => {
  it('returns all possible toChainIds that a fromChainId can send to, single', () => {
    const toChainIds = getPossibleToChainIdsByFromChainId(50)

    expect(toChainIds).toEqual([1])
  })

  it('returns all possible toChainIds that a fromChainId can send to, multiple', () => {
    const toChainIds = getPossibleToChainIdsByFromChainId(25)

    expect(toChainIds).toEqual([1, 10, 56])
  })
})

describe('getPossibleToChainIdsByFromToken', () => {
  it('returns all possible toChainIds that a fromToken can send to, single', () => {
    const toChainIds = getPossibleToChainIdsByFromToken('XYZ-50')

    expect(toChainIds).toEqual([1])
  })
  it('returns all possible toChainIds that a fromToken can send to, multiple', () => {
    const toChainIds = getPossibleToChainIdsByFromToken('USDC-25')

    expect(toChainIds).toEqual([1, 10, 56])
  })
})

describe('getPossibleToChainIdsByToToken', () => {
  it('returns all possible toChainIds by a toToken, single', () => {
    const toChainIds = getPossibleToChainIdsByToToken('XYZ-50')

    expect(toChainIds).toEqual([1])
  })

  it('returns all possible toChainIds by a toToken', () => {
    const toChainIds = getPossibleToChainIdsByToToken('USDC-25')

    expect(toChainIds.sort()).toEqual([1, 10, 56].sort())
  })
})

describe('getPossibleToTokens', () => {
  it('returns all possible toTokens', () => {
    const toTokens = getPossibleToTokens()

    expect(toTokens.sort()).toEqual([
      'GOHM-1',
      'GOHM-10',
      'GOHM-25',
      'GOHM-56',
      'HIGHSTREET-1',
      'HIGHSTREET-56',
      'SYN-1',
      'SYN-10',
      'SYN-25',
      'SYN-56',
      'USDC-1',
      'USDC-10',
      'USDC-25',
      'USDC-56',
      'XYZ-1',
      'XYZ-50',
    ])
  })
})

describe('getPossibleToTokensByFromChainId', () => {
  it('returns possible toTokens by a fromChainId, single', () => {
    const toTokens = getPossibleToTokensByFromChainId(50)

    expect(toTokens).toEqual(['XYZ-1'])
  })

  it('returns possible toTokens by a fromChainId, multiple', () => {
    const toTokens = getPossibleToTokensByFromChainId(10)

    expect(toTokens.sort()).toEqual([
      'GOHM-1',
      'GOHM-25',
      'GOHM-56',
      'SYN-1',
      'SYN-25',
      'SYN-56',
      'USDC-1',
      'USDC-25',
      'USDC-56',
    ])
  })
})

describe('getPossibleToTokensByFromToken', () => {
  it('returns possible toTokens given a fromToken, single', () => {
    const toTokens = getPossibleToTokensByFromToken('XYZ-1')

    expect(toTokens).toEqual(['XYZ-50'])
  })

  it('returns possible toTokens given a fromToken, multiple', () => {
    const toTokens = getPossibleToTokensByFromToken('USDC-1')

    expect(toTokens).toEqual(['USDC-10', 'USDC-25', 'USDC-56'])
  })
})

describe('getPossibleToTokensByFromTokenAndToChainId', () => {
  it('returns possible toTokens given a fromToken and a toChainId, single', () => {
    const toTokens = getPossibleToTokensByFromTokenAndToChainId('XYZ-1', 50)

    expect(toTokens).toEqual(['XYZ-50'])
  })

  it('returns possible toTokens given a fromToken and a toChainId, multiple', () => {
    const toTokens = getPossibleToTokensByFromTokenAndToChainId('USDC-1', 25)

    expect(toTokens).toEqual(['USDC-25'])
  })

  it('returns possible toTokens given a fromToken and a toChainId, no path', () => {
    const toTokens = getPossibleToTokensByFromTokenAndToChainId('USDC-1', 50)

    expect(toTokens).toEqual([])
  })
})

describe('getPossibleToTokensByToChainId', () => {
  it('returns possible toTokens based on toChainId, single', () => {
    const toTokens = getPossibleToTokensByToChainId(50)

    expect(toTokens).toEqual(['XYZ-50'])
  })

  it('returns possible toTokens based on toChainId, multiple', () => {
    const toTokens = getPossibleToTokensByToChainId(25)

    expect(toTokens).toEqual(['GOHM-25', 'USDC-25', 'SYN-25'])
  })
})

describe('getPossibleToTokensByToToken', () => {
  it('returns possible toTokens based on a toToken, single', () => {
    const toTokens = getPossibleToTokensByToToken('XYZ-50')

    expect(toTokens).toEqual(['XYZ-50'])
  })

  it('returns possible toTokens based on a toToken, multiple', () => {
    const toTokens = getPossibleToTokensByToToken('USDC-25')

    expect(toTokens).toEqual(['GOHM-25', 'USDC-25', 'SYN-25'])
  })
})
