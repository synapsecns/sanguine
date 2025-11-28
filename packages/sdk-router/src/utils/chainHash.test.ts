import { marshallChainHash, unmarshallChainHash, ChainHash } from './chainHash'

describe('chainHash', () => {
  describe('marshallChainHash', () => {
    it('converts ChainHash to string format', () => {
      const chainHash: ChainHash = { chainId: 1, hash: '0xabc123' }
      expect(marshallChainHash(chainHash)).toBe('0xabc123:1')
    })

    it('handles chainId of 0', () => {
      const chainHash: ChainHash = { chainId: 0, hash: '0xdef456' }
      expect(marshallChainHash(chainHash)).toBe('0xdef456:0')
    })
  })

  describe('unmarshallChainHash', () => {
    it('parses valid chain hash string', () => {
      const result = unmarshallChainHash('0xhash:42')
      expect(result).toEqual({ chainId: 42, hash: '0xhash' })
    })

    it('handles chainId of 0', () => {
      const result = unmarshallChainHash('0xhash:0')
      expect(result).toEqual({ chainId: 0, hash: '0xhash' })
    })

    it('throws on missing colon', () => {
      expect(() => unmarshallChainHash('invalid')).toThrow(
        'Invalid chain hash format'
      )
    })

    it('throws on multiple colons', () => {
      expect(() => unmarshallChainHash('0xhash:1:2')).toThrow(
        'Invalid chain hash format'
      )
    })

    it('throws on empty string', () => {
      expect(() => unmarshallChainHash('')).toThrow('Invalid chain hash format')
    })
  })
})
