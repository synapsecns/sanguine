import fs from 'fs'
import path from 'path'

import { SupportedChainId } from '../constants'
import { ETH_NATIVE_TOKEN_ADDRESS } from '../utils'
import { getSbaRemoteToken, getSbaSupportedTokens } from './supportedTokens'

type AdapterTokenConfig = Record<
  string,
  Record<string, { tokenAddress: string; isUnderlying: boolean }>
>

const adapterTokens = JSON.parse(
  fs.readFileSync(
    path.resolve(
      __dirname,
      '../../../contracts-adapter/configs/global/tokens.json'
    ),
    'utf8'
  )
) as AdapterTokenConfig

const chainIds: Record<string, SupportedChainId> = {
  arbitrum: SupportedChainId.ARBITRUM,
  aurora: SupportedChainId.AURORA,
  avalanche: SupportedChainId.AVALANCHE,
  base: SupportedChainId.BASE,
  blast: SupportedChainId.BLAST,
  bnb: SupportedChainId.BSC,
  canto: SupportedChainId.CANTO,
  cronos: SupportedChainId.CRONOS,
  dfk: SupportedChainId.DFK,
  ethereum: SupportedChainId.ETH,
  fantom: SupportedChainId.FANTOM,
  harmony: SupportedChainId.HARMONY,
  kaia: SupportedChainId.KLAYTN,
  metis: SupportedChainId.METIS,
  moonbeam: SupportedChainId.MOONBEAM,
  moonriver: SupportedChainId.MOONRIVER,
  optimism: SupportedChainId.OPTIMISM,
  polygon: SupportedChainId.POLYGON,
}

// These underlying tokens are the chain's wrapped native asset. The lookup
// exposes the native address when they are the destination.
const wrappedNativeChains: Record<string, string> = {
  AVAX: 'avalanche',
  FTM: 'fantom',
  JEWEL: 'dfk',
  KLAY: 'kaia',
  MATIC: 'polygon',
  Metis: 'metis',
  nETH: 'ethereum',
}

describe('SBA supported token snapshot', () => {
  it('matches adapter tokens on supported chains, excluding GMX', () => {
    const expectedByPair = new Map<string, string[]>()

    Object.entries(adapterTokens).forEach(([tokenId, chains]) => {
      // Moonbeam and Moonriver remain outside the supported token snapshot.
      const chainEntries = Object.entries(chains).filter(
        ([chain]) => chain !== 'moonbeam' && chain !== 'moonriver'
      )
      if (tokenId === 'GMX' || chainEntries.length < 2) {
        return
      }

      chainEntries.forEach(([originChain, origin]) => {
        chainEntries.forEach(([destChain, dest]) => {
          if (originChain === destChain) {
            return
          }

          const originChainId = chainIds[originChain]
          const destChainId = chainIds[destChain]
          expect(originChainId).toBeDefined()
          expect(destChainId).toBeDefined()

          const effectiveDestToken =
            wrappedNativeChains[tokenId] === destChain
              ? ETH_NATIVE_TOKEN_ADDRESS
              : dest.tokenAddress
          const pairKey = `${originChainId}:${destChainId}`
          const expected = expectedByPair.get(pairKey) ?? []
          expected.push(
            `${origin.tokenAddress.toLowerCase()}:${effectiveDestToken.toLowerCase()}`
          )
          expectedByPair.set(pairKey, expected)

          expect(
            getSbaRemoteToken(
              originChainId,
              origin.tokenAddress,
              destChainId
            )?.toLowerCase()
          ).toBe(effectiveDestToken.toLowerCase())
        })
      })
    })

    Object.values(chainIds).forEach((originChainId) => {
      Object.values(chainIds).forEach((destChainId) => {
        const pairKey = `${originChainId}:${destChainId}`
        const actual = getSbaSupportedTokens(originChainId, destChainId)
          .map(
            ({ originToken, destToken }) =>
              `${originToken.toLowerCase()}:${destToken.toLowerCase()}`
          )
          .sort()
        expect(actual).toEqual((expectedByPair.get(pairKey) ?? []).sort())
      })
    })
  })
})
