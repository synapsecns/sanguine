import * as chains from './chains'

export enum SupportedChainId {
  ETH = chains.mainnet.id,
  OPTIMISM = chains.optimism.id,
  CRONOS = chains.cronos.id,
  BSC = chains.bsc.id,
  POLYGON = chains.polygon.id,
  FANTOM = chains.fantom.id,
  BOBA = chains.boba.id,
  METIS = chains.metis.id,
  MOONBEAM = chains.moonbeam.id,
  MOONRIVER = chains.moonriver.id,
  DOGECHAIN = chains.dogechain.id,
  CANTO = chains.canto.id,
  KLAYTN = chains.klaytn.id,
  BASE = chains.base.id,
  ARBITRUM = chains.arbitrum.id,
  AVALANCHE = chains.avalanche.id,
  DFK = chains.dfk.id,
  AURORA = chains.aurora.id,
  HARMONY = chains.harmonyOne.id,
}

/**
 * List of supported chain ids, where SynapseBridge is deployed.
 */
export const SUPPORTED_CHAIN_IDS: number[] = Object.values(SupportedChainId)
  .map((chainId) => Number(chainId))
  .filter((chainId) => !isNaN(chainId))

/**
 * List of chain ids where SynapseCCTP is deployed, ordered by CCTP's domain:
 * https://developers.circle.com/stablecoin/docs/cctp-protocol-contract#mainnet-contract-addresses
 *
 * Note: This is a subset of SUPPORTED_CHAIN_IDS.
 */
export const CCTP_SUPPORTED_CHAIN_IDS: number[] = [
  SupportedChainId.ETH,
  SupportedChainId.AVALANCHE,
  SupportedChainId.OPTIMISM,
  SupportedChainId.ARBITRUM,
  SupportedChainId.BASE,
]
