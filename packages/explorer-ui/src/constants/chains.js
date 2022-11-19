import arbitrumLogo from 'assets/icons/arbitrum.svg'
import auroraLogo from 'assets/icons/aurora.svg'
import avalancheLogo from 'assets/icons/avalanche.svg'
import binanceLogo from 'assets/icons/binance.svg'
import bobaLogo from 'assets/icons/boba.svg'
import cronosLogo from 'assets/icons/cronos.svg'
import dfkLogo from 'assets/icons/dfk.svg'
import ethLogo from 'assets/icons/eth.svg'
import fantomLogo from 'assets/icons/fantom.svg'
import harmonyLogo from 'assets/icons/harmonyone.svg'
import klaytnLogo from 'assets/icons/klaytn.svg'
import metisLogo from 'assets/icons/metis.svg'
import optimismLogo from 'assets/icons/optimism.svg'
import polygonLogo from 'assets/icons/polygon.svg'
import arbitrumImg from 'assets/networks/arbitrum.jpg'
import auroraImg from 'assets/networks/aurora.png'
import avalancheImg from 'assets/networks/avalanche.jpg'
import bobaImg from 'assets/networks/boba.png'
import bscImg from 'assets/networks/bsc.jpg'
import cronosImg from 'assets/networks/cronos.png'
import dfkImg from 'assets/networks/dfk.png'
import ethImg from 'assets/networks/eth.jpg'
import fantomImg from 'assets/networks/fantom.jpg'
import harmonyImg from 'assets/networks/harmonyone.jpg'
import klaytnImg from 'assets/networks/klaytn.png'
import metisImg from 'assets/networks/metis.png'
import moonbeamImg from 'assets/networks/moonbeam.jpg'
import moonriverImg from 'assets/networks/moonriver.jpeg'
import optimismImg from 'assets/networks/optimism.png'
import polygonImg from 'assets/networks/polygon.jpg'

export function toHexStr(num) {
  return `0x${num.toString(16)}`
}

export const CHAIN_ID = {
  ETH: 1,
  ROPSTEN: 3,
  RINKEBY: 4,
  GÖRLI: 5,
  OPTIMISM: 10,
  METIS: 25,
  KOVAN: 42,
  BSC: 56,
  POLYGON: 137,
  FANTOM: 250,
  BOBA: 288,
  CRONOS: 1088,
  MOONBEAM: 1284,
  MOONRIVER: 1285,
  KLAYTN: 8217,
  HARDHAT: 31337,
  ARBITRUM: 42161,
  AVALANCHE: 43114,
  DFK: 53935,
  AURORA: 1313161554,
  HARMONY: 1666600000,
}

export const CHAIN_ID_REVERSE = Object.fromEntries(Object.entries(CHAIN_ID).map(([k, v]) => [v, k]))

export const CHAIN_ID_NAMES = {
  [CHAIN_ID.ETH]: 'ethereum',
  [CHAIN_ID.OPTIMISM]: 'optimism',
  [CHAIN_ID.BSC]: 'bsc',
  [CHAIN_ID.POLYGON]: 'polygon',
  [CHAIN_ID.FANTOM]: 'fantom',
  [CHAIN_ID.BOBA]: 'boba',
  [CHAIN_ID.MOONBEAM]: 'moonbeam',
  [CHAIN_ID.MOONRIVER]: 'moonriver',
  [CHAIN_ID.ARBITRUM]: 'arbitrum',
  [CHAIN_ID.AVALANCHE]: 'avalanche',
  [CHAIN_ID.AURORA]: 'aurora',
  [CHAIN_ID.HARMONY]: 'harmony',
  [CHAIN_ID.CRONOS]: 'cronos',
  [CHAIN_ID.METIS]: 'metis',
  [CHAIN_ID.DFK]: 'dfk',
  [CHAIN_ID.KLAYTN]: 'klaytn',
}

export const CHAIN_ID_NAMES_REVERSE = Object.fromEntries(Object.entries(CHAIN_ID_NAMES).map(([k, v]) => [v, k]))

export const CHAIN_INFO = {
  [CHAIN_ID.ETH]: {
    CHAIN_ID: CHAIN_ID.ETH,
    chainSymbol: 'ETH',
    chainName: 'Ethereum',
    chainLogo: ethLogo,
    chainImg: ethImg,
    chainColor: '#28a0f0',
  },
  [CHAIN_ID.OPTIMISM]: {
    CHAIN_ID: CHAIN_ID.OPTIMISM,
    chainSymbol: 'OPTIMISM',
    chainName: 'Optimism',
    chainLogo: optimismLogo,
    chainImg: optimismImg,
    chainColor: '#fe0621',
  },
  [CHAIN_ID.BSC]: {
    CHAIN_ID: CHAIN_ID.BSC,
    chainSymbol: 'BSC',
    chainName: 'Binance Smart Chain',
    chainShortName: 'BSC',
    chainLogo: binanceLogo,
    chainImg: bscImg,
    chainColor: '#efb90b',
  },
  [CHAIN_ID.POLYGON]: {
    CHAIN_ID: CHAIN_ID.POLYGON,
    chainSymbol: 'POLYGON',
    chainName: 'Polygon',
    chainLogo: polygonLogo,
    chainImg: polygonImg,
    chainColor: '#7b3fe4',
  },
  [CHAIN_ID.FANTOM]: {
    CHAIN_ID: CHAIN_ID.FANTOM,
    chainSymbol: 'FANTOM',
    chainName: 'Fantom',
    chainLogo: fantomLogo,
    chainImg: fantomImg,
    chainColor: '#1969ff',
  },
  [CHAIN_ID.BOBA]: {
    CHAIN_ID: CHAIN_ID.BOBA,
    chainSymbol: 'BOBA',
    chainName: 'Boba Network',
    chainLogo: bobaLogo,
    chainImg: bobaImg,
    chainColor: '#cbff00',
  },
  [CHAIN_ID.MOONBEAM]: {
    CHAIN_ID: CHAIN_ID.MOONBEAM,
    chainSymbol: 'MOONBEAM',
    chainName: 'Moonbeam',
    chainLogo: moonbeamImg,
    chainImg: moonbeamImg,
    chainColor: '#f2b707',
  },
  [CHAIN_ID.MOONRIVER]: {
    CHAIN_ID: CHAIN_ID.MOONRIVER,
    chainSymbol: 'MOONRIVER',
    chainName: 'Moonriver',
    chainLogo: moonriverImg,
    chainImg: moonriverImg,
    chainColor: '#f2b707',
  },
  [CHAIN_ID.ARBITRUM]: {
    CHAIN_ID: CHAIN_ID.ARBITRUM,
    chainSymbol: 'ARBITRUM',
    chainName: 'Arbitrum',
    chainLogo: arbitrumLogo,
    chainImg: arbitrumImg,
    chainColor: '#434971',
  },
  [CHAIN_ID.AVALANCHE]: {
    CHAIN_ID: CHAIN_ID.AVALANCHE,
    chainSymbol: 'AVALANCHE',
    chainName: 'Avalanche',
    chainLogo: avalancheLogo,
    chainImg: avalancheImg,
    chainColor: '#e74242',
  },
  [CHAIN_ID.AURORA]: {
    CHAIN_ID: CHAIN_ID.AURORA,
    chainSymbol: 'AURORA',
    chainName: 'Aurora',
    chainLogo: auroraLogo,
    chainImg: auroraImg,
    chainColor: '#78d64b',
  },
  [CHAIN_ID.HARMONY]: {
    CHAIN_ID: CHAIN_ID.HARMONY,
    chainSymbol: 'HARMONY',
    chainName: 'Harmony',
    chainLogo: harmonyLogo,
    chainImg: harmonyImg,
    chainColor: '#39cdd8',
  },
  [CHAIN_ID.CRONOS]: {
    CHAIN_ID: CHAIN_ID.CRONOS,
    chainSymbol: 'CRONOS',
    chainName: 'Cronos',
    chainLogo: cronosLogo,
    chainImg: cronosImg,
    chainColor: '#1711a2',
  },
  [CHAIN_ID.METIS]: {
    CHAIN_ID: CHAIN_ID.METIS,
    chainSymbol: 'METIS',
    chainName: 'Metis',
    chainLogo: metisLogo,
    chainImg: metisImg,
    chainColor: '#22e5f2',
  },
  [CHAIN_ID.DFK]: {
    CHAIN_ID: CHAIN_ID.DFK,
    chainSymbol: 'JEWEL',
    chainName: 'DFK',
    chainLogo: dfkLogo,
    chainImg: dfkImg,
    chainColor: '#ffff83',
  },
  [CHAIN_ID.KLAYTN]: {
    CHAIN_ID: CHAIN_ID.KLAYTN,
    chainSymbol: 'KLAY',
    chainName: 'Klaytn',
    chainLogo: klaytnLogo,
    chainImg: klaytnImg,
    chainColor: '#f9810b',
  },
}
