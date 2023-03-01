import { Web3Provider } from '@ethersproject/providers'

import { BLOCK_TIME, ChainId, CHAIN_BLOCK_TIME } from '@constants/networks'

export function getLibrary(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = BLOCK_TIME
  return library
}

export function getLibraryBsc(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.BSC]
  return library
}

export function getLibraryEth(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.ETH]
  return library
}

export function getLibraryPolygon(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.POLYGON]
  return library
}

export function getLibraryFantom(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.FANTOM]
  return library
}

export function getLibraryBoba(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.BOBA]
  return library
}

export function getLibraryMoonbeam(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.MOONBEAM]
  return library
}

export function getLibraryMoonriver(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.MOONRIVER]
  return library
}

export function getLibraryArbitrum(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.ARBITRUM]
  return library
}

export function getLibraryAvalanche(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.AVALANCHE]
  return library
}

export function getLibraryDfk(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.DFK]
  return library
}

export function getLibraryHarmony(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.HARMONY]
  return library
}

export function getLibraryOptimism(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.OPTIMISM]
  return library
}

export function getLibraryAurora(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.AURORA]
  return library
}

export function getLibraryCronos(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.CRONOS]
  return library
}

export function getLibraryMetis(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.METIS]
  return library
}



export function getLibraryKlaytn(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.KLAYTN]
  return library
}

export function getLibraryDogechain(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.DOGECHAIN]
  return library
}


export function getLibraryCanto(provider) {
  const library = new Web3Provider(provider, 'any')
  library.pollingInterval = CHAIN_BLOCK_TIME[ChainId.CANTO]
  return library
}


