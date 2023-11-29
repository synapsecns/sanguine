import _ from 'lodash'
import busdLogo from '@assets/icons/busd.svg'
import usdcLogo from '@assets/icons/usdc.svg'
import usdtLogo from '@assets/icons/usdt.svg'
import synapseLogo from '@assets/icons/synapse.svg'
import ethLogo from '@assets/icons/eth.svg'
import wethLogo from '@assets/icons/weth.svg'
import nethLogo from '@assets/icons/neth.svg'
import avwethLogo from '@assets/icons/avweth.svg'
import mimLogo from '@assets/icons/mim.svg'
import ustLogo from '@assets/icons/ust.png'
import fraxLogo from '@assets/icons/frax.svg'
import daiLogo from '@assets/icons/dai.png'
import nusdLogo from '@assets/icons/nusd.svg'
import avaxLogo from '@assets/icons/avalanche.svg'
import movrLogo from '@assets/icons/moonriver.jpeg'
import jewelLogo from '@assets/icons/jewel.png'
import wbtcLogo from '@assets/icons/wbtc.svg'
import noteLogo from '@assets/icons/note.svg'
import klaytnLogo from '@assets/networks/klaytn.jpeg'
import pepeLogo from '@assets/icons/pepe.webp'
import maticLogo from '@assets/icons/matic.svg'
import crvusdLogo from '@assets/icons/crvusd.svg'
import ftmLogo from '@assets/icons/ftm.svg'
import susdLogo from '@assets/icons/susd.svg'
import lusdLogo from '@assets/icons/lusd.svg'
import { ChainId } from '@constants/networks'
import { Token } from '@utils/classes/Token'
import {
  DOG,
  GOHM,
  HIGHSTREET,
  JUMP,
  NFD,
  SOLAR,
  NEWO,
  SDT,
  USDB,
  GMX,
  VSTA,
  SFI,
  H2O,
  L2DAO,
  AGEUR,
  PLS,
  LINK,
  UNIDX,
} from '@constants/tokens/mintable'

export const BUSD = new Token({
  addresses: {
    [ChainId.BSC]: '0xe9e7cea3dedca5984780bafc599bd69add087d56',
    [ChainId.DOGECHAIN]: '0x1555C68Be3b22cdcCa934Ae88Cb929Db40aB311d',
  },
  decimals: 18,
  symbol: 'BUSD',
  name: 'Binance USD',
  logo: busdLogo,
  description: `
    BUSD is a stablecoin that is pegged to the US dollar and
    backed/issued by Binance
  `,
  swapableType: 'USD',
})

export const USDC = new Token({
  addresses: {
    [ChainId.BSC]: '0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d',
    [ChainId.ETH]: '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
    [ChainId.CRONOS]: '0xc21223249ca28397b4b6541dffaecc539bff0c59',
    [ChainId.OPTIMISM]: '0x7f5c764cbc14f9669b88837ca1490cca17c31607',
    [ChainId.POLYGON]: '0x2791bca1f2de4661ed88a30c99a7a9449aa84174',
    [ChainId.FANTOM]: '0x04068da6c83afcfa0e13ba15a6696662335d5b75',
    [ChainId.AVALANCHE]: '0xA7D7079b0FEaD91F3e65f86E8915Cb59c1a4C664',
    [ChainId.ARBITRUM]: '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8',
    [ChainId.HARMONY]: '0x985458e523db3d53125813ed68c274899e9dfab4',
    [ChainId.BOBA]: '0x66a2A913e447d6b4BF33EFbec43aAeF87890FBbc',
    [ChainId.AURORA]: '0xB12BFcA5A55806AaF64E99521918A4bf0fC40802',
    [ChainId.METIS]: '0xEA32A96608495e54156Ae48931A7c20f0dcc1a21',
    [ChainId.KLAYTN]: '0x6270B58BE569a7c0b8f47594F191631Ae5b2C86C',
    [ChainId.CANTO]: '0x80b5a32E4F032B2a058b4F29EC95EEfEEB87aDcd',
  },
  decimals: {
    [ChainId.BSC]: 18,
    [ChainId.ETH]: 6,
    [ChainId.OPTIMISM]: 6,
    [ChainId.POLYGON]: 6,
    [ChainId.FANTOM]: 6,
    [ChainId.AVALANCHE]: 6,
    [ChainId.ARBITRUM]: 6,
    [ChainId.HARMONY]: 6,
    [ChainId.BOBA]: 6,
    [ChainId.AURORA]: 6,
    [ChainId.METIS]: 6,
    [ChainId.CRONOS]: 6,
    [ChainId.KLAYTN]: 6,
    [ChainId.CANTO]: 6,
  },
  symbol: 'USDC',
  name: 'USD Circle',
  logo: usdcLogo,
  description: `
    USD Coin (known by its ticker USDC) is a stablecoin that is pegged to the
    U.S. dollar on a 1:1 basis. Every unit of this cryptocurrency in circulation
    is backed up by $1 that is held in reserve
    `,
  swapableType: 'USD',
})
export const CCTP_USDC = new Token({
  addresses: {
    [ChainId.AVALANCHE]: '0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E',
    [ChainId.ARBITRUM]: '0xaf88d065e77c8cC2239327C5EDb3A432268e5831',
    [ChainId.BASE]: '0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913',
    [ChainId.OPTIMISM]: '0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85',
  },
  decimals: {
    [ChainId.ETH]: 6,
    [ChainId.AVALANCHE]: 6,
    [ChainId.ARBITRUM]: 6,
    [ChainId.BASE]: 6,
    [ChainId.OPTIMISM]: 6,
  },
  symbol: 'USDC',
  name: 'USD Circle',
  logo: usdcLogo,
  description: `
    USD Coin (known by its ticker USDC) is a stablecoin that is pegged to the
    U.S. dollar on a 1:1 basis. Every unit of this cryptocurrency in circulation
    is backed up by $1 that is held in reserve
    `,
  swapableType: 'USD',
})
export const KLAYTN_USDC = new Token({
  addresses: {
    [ChainId.ETH]: '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
    [ChainId.KLAYTN]: '0x6270B58BE569a7c0b8f47594F191631Ae5b2C86C',
    [ChainId.DOGECHAIN]: '0x85C2D3bEBffD83025910985389aB8aD655aBC946',
  },
  decimals: {
    [ChainId.ETH]: 6,
    [ChainId.KLAYTN]: 6,
    [ChainId.DOGECHAIN]: 6,
  },
  symbol: 'USDC  ', // TWO SPACES IS EXTREMELY IMPORTANT
  name: 'USD Circle',
  logo: usdcLogo,
  description: `
    USD Coin (known by its ticker USDC) is a stablecoin that is pegged to the
    U.S. dollar on a 1:1 basis. Every unit of this cryptocurrency in circulation
    is backed up by $1 that is held in reserve
    `,
  swapableType: 'KLAYTN_USDC',
})

export const KLAYTN_USDT = new Token({
  addresses: {
    [ChainId.ETH]: '0xdac17f958d2ee523a2206206994597c13d831ec7',
    [ChainId.KLAYTN]: '0xd6dAb4CfF47dF175349e6e7eE2BF7c40Bb8C05A3',
    [ChainId.DOGECHAIN]: '0x7f8e71DD5A7e445725F0EF94c7F01806299e877A',
  },
  decimals: {
    [ChainId.ETH]: 6,
    [ChainId.KLAYTN]: 6,
    [ChainId.DOGECHAIN]: 6,
  },
  symbol: 'USDT  ', // TWO SPACES IS EXTREMELY IMPORTANT
  name: 'Synapse Tether USDT',
  logo: usdtLogo,
  swapableType: 'KLAYTN_USDT',
})

export const KLAYTN_oUSDT = new Token({
  addresses: {
    [ChainId.KLAYTN]: '0xceE8FAF64bB97a73bb51E115Aa89C17FfA8dD167',
  },
  decimals: {
    [ChainId.KLAYTN]: 6,
  },
  symbol: 'orbitUSDT', // TWO SPACES IS EXTREMELY IMPORTANT
  name: 'Orbit Bridged USDT',
  logo: usdtLogo,
  swapableType: 'KLAYTN_USDT',
})

export const KLAYTN_DAI = new Token({
  addresses: {
    [ChainId.ETH]: '0x6b175474e89094c44da98b954eedeac495271d0f',
    [ChainId.KLAYTN]: '0x078dB7827a5531359f6CB63f62CFA20183c4F10c',
    [ChainId.DOGECHAIN]: '0xB3306f03595490e5cC3a1b1704a5a158D3436ffC',
  },
  decimals: {
    [ChainId.ETH]: 18,
    [ChainId.KLAYTN]: 18,
    [ChainId.DOGECHAIN]: 18,
  },
  symbol: 'DAI  ', // TWO SPACES IS EXTREMELY IMPORTANT
  name: 'DAI',
  logo: daiLogo,
  swapableType: 'KLAYTN_DAI',
})

export const DOGECHAIN_BUSD = new Token({
  addresses: {
    [ChainId.BSC]: '0xe9e7cea3dedca5984780bafc599bd69add087d56',
    [ChainId.DOGECHAIN]: '0x1555C68Be3b22cdcCa934Ae88Cb929Db40aB311d',
  },
  decimals: {
    [ChainId.BSC]: 18,
    [ChainId.DOGECHAIN]: 18,
  },
  symbol: 'BUSD ', // ONE SPACE IS EXTREMELY IMPORTANT
  name: 'Binance USD',
  logo: busdLogo,
  swapableType: 'DOGECHAIN_BUSD',
})

export const USDT = new Token({
  addresses: {
    [ChainId.BSC]: '0x55d398326f99059ff775485246999027b3197955',
    [ChainId.ETH]: '0xdac17f958d2ee523a2206206994597c13d831ec7',
    [ChainId.CRONOS]: '0x66e428c3f67a68878562e79a0234c1f83c208770',
    [ChainId.POLYGON]: '0xc2132d05d31c914a87c6611c10748aeb04b58e8f',
    [ChainId.AVALANCHE]: '0x9702230a8ea53601f5cd2dc00fdbc13d4df4a8c7',
    [ChainId.HARDHAT]: '0x9A9f2CCfdE556A7E9Ff0848998Aa4a0CFD8863AE',
    [ChainId.ARBITRUM]: '0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9',
    [ChainId.FANTOM]: '0x049d68029688eabf473097a2fc38ef61633a3c7a',
    [ChainId.HARMONY]: '0x3c2b8be99c50593081eaa2a724f0b8285f5aba8f',
    [ChainId.BOBA]: '0x5DE1677344D3Cb0D7D465c10b72A8f60699C062d',
    [ChainId.AURORA]: '0x4988a896b1227218e4A686fdE5EabdcAbd91571f',
    [ChainId.KLAYTN]: '0xd6dAb4CfF47dF175349e6e7eE2BF7c40Bb8C05A3',
    [ChainId.CANTO]: '0xd567B3d7B8FE3C79a1AD8dA978812cfC4Fa05e75',
    [ChainId.OPTIMISM]: '0x94b008aA00579c1307B0EF2c499aD98a8ce58e58',
  },
  decimals: {
    [ChainId.BSC]: 18,
    [ChainId.ETH]: 6,
    [ChainId.POLYGON]: 6,
    [ChainId.AVALANCHE]: 6,
    [ChainId.ARBITRUM]: 6,
    [ChainId.FANTOM]: 6,
    [ChainId.HARMONY]: 6,
    [ChainId.BOBA]: 6,
    [ChainId.AURORA]: 6,
    [ChainId.CRONOS]: 6,
    [ChainId.CANTO]: 6,
    [ChainId.OPTIMISM]: 6,
  },
  symbol: 'USDT',
  name: 'USD Tether',
  logo: usdtLogo,
  description: `
    USDT mirrors the price of the U.S. dollar, issued by a Hong Kong-based company Tether.
    The token’s peg to the USD is achieved via maintaining a sum of dollars in reserves equal
    to the number of USDT in circulation.
    `,
  swapableType: 'USD',
})

export const USDTE = new Token({
  addresses: {
    [ChainId.AVALANCHE]: '0xc7198437980c041c805a1edcba50c1ce5db95118',
  },
  decimals: 6,
  symbol: 'USDT.e',
  name: 'Avalanche Bridged USDT.e',
  logo: usdtLogo,
  swapableType: 'USD',
})

export const DAI = new Token({
  addresses: {
    [ChainId.BSC]: '0x1af3f329e8be154074d8769d1ffa4ee058b1dbc3',
    [ChainId.ETH]: '0x6b175474e89094c44da98b954eedeac495271d0f',
    [ChainId.CRONOS]: '0xf2001b145b43032aaf5ee2884e456ccd805f677d',
    [ChainId.POLYGON]: '0x8f3cf7ad23cd3cadbd9735aff958023239c6a063',
    [ChainId.AVALANCHE]: '0xd586E7F844cEa2F87f50152665BCbc2C279D8d70',
    [ChainId.ARBITRUM]: '0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1',
    [ChainId.HARMONY]: '0xef977d2f931c1978db5f6747666fa1eacb0d0339',
    [ChainId.BOBA]: '0xf74195Bb8a5cf652411867c5C2C5b8C2a402be35',
    [ChainId.KLAYTN]: '0x078dB7827a5531359f6CB63f62CFA20183c4F10c',
    [ChainId.BASE]: '0x50c5725949A6F0c72E6C4a641F24049A917DB0Cb',
    [ChainId.OPTIMISM]: '0xda10009cbd5d07dd0cecc66161fc93d7c9000da1',
  },
  decimals: 18,
  symbol: 'DAI',
  name: 'Dai',
  logo: daiLogo,
  swapableType: 'USD',
})

export const SUSD = new Token({
  addresses: {
    [ChainId.OPTIMISM]: '0x8c6f28f2F1A3C87F0f938b96d27520d9751ec8d9',
  },
  decimals: 18,
  symbol: 'sUSD',
  name: 'Synth USD',
  logo: susdLogo,
  swapableType: 'USD',
})

export const LUSD = new Token({
  addresses: {
    [ChainId.ETH]: '0x5f98805A4E8be255a32880FDeC7F6728C6568bA0',
  },
  decimals: 18,
  symbol: 'LUSD',
  name: 'Liquidity USD',
  logo: lusdLogo,
  swapableType: 'USD',
})

export const CRVUSD = new Token({
  addresses: {
    [ChainId.ETH]: '0xf939E0A03FB07F59A73314E73794Be0E57ac1b4E',
    [ChainId.BASE]: '0x417ac0e078398c154edfadd9ef675d30be60af93',
  },
  decimals: 18,
  symbol: 'crvUSD',
  name: 'Curve.fi USD',
  logo: crvusdLogo,
  swapableType: 'USD',
})

export const USDBC = new Token({
  addresses: {
    [ChainId.BASE]: '0xd9aAEc86B65D86f6A7B5B1b0c42FFA531710b6CA',
  },
  decimals: 6,
  symbol: 'USDbC',
  name: 'USD Base Coin',
  logo: usdcLogo,
  swapableType: 'USD',
})

export const WBTC = new Token({
  addresses: {
    [ChainId.ETH]: '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
    [ChainId.KLAYTN]: '0xDCbacF3f7a069922E677912998c8d57423C37dfA',
    [ChainId.DOGECHAIN]: '0xD0c6179c43C00221915f1a61f8eC06A5Aa32F9EC',
  },
  decimals: {
    [ChainId.ETH]: 8,
    [ChainId.KLAYTN]: 8,
    [ChainId.DOGECHAIN]: 8,
  },
  symbol: 'WBTC',
  name: 'Wrapped BTC',
  logo: wbtcLogo,
  swapableType: 'WBTC',
})

export const BTCB = new Token({
  addresses: {
    [ChainId.AVALANCHE]: '0x152b9d0fdc40c096757f570a51e494bd4b943e50',
    [ChainId.DFK]: '0x7516EB8B8Edfa420f540a162335eACF3ea05a247',
  },
  decimals: {
    [ChainId.AVALANCHE]: 8,
    [ChainId.DFK]: 8,
  },
  symbol: 'BTC.b',
  name: 'Bridged Bitcoin on Avalanche',
  logo: wbtcLogo,
  swapableType: 'WBTC',
})

export const UST = new Token({
  addresses: {
    // [ChainId.BSC]:       '0x23396cf899ca06c4472205fc903bdb4de249d6fc',
    // [ChainId.POLYGON]:   '0xed7a89e2d580bdda005ba4cdd64cf0da3c15a5eb',
    // [ChainId.AVALANCHE]: '0x195150ab664795128ac542f26b14c1a12e061ecf',
    // [ChainId.TERRA]:     'uusd',
    [ChainId.ETH]: '0x0261018Aa50E28133C1aE7a29ebdf9Bd21b878Cb',
    [ChainId.BSC]: '0xb7A6c5f0cc98d24Cf4B2011842e64316Ff6d042c',
    [ChainId.POLYGON]: '0x565098CBa693b3325f9fe01D41b7A1cd792Abab1',
    [ChainId.FANTOM]: '0xa0554607e477cdC9d0EE2A6b087F4b2DC2815C22',
    [ChainId.ARBITRUM]: '0x13780E6d5696DD91454F6d3BbC2616687fEa43d0',
    [ChainId.AVALANCHE]: '0xE97097dE8d6A17Be3c39d53AE63347706dCf8f43',
    [ChainId.HARMONY]: '0xa0554607e477cdC9d0EE2A6b087F4b2DC2815C22',
    [ChainId.BOBA]: '0x61A269a9506272D128d79ABfE8E8276570967f00',
    [ChainId.MOONRIVER]: '0xa9D0C0E124F53f4bE1439EBc35A9C73c0e8275fB',
    [ChainId.MOONBEAM]: '0x5CF84397944B9554A278870B510e86667681ff8D',
    [ChainId.OPTIMISM]: '0xFB21B70922B9f6e3C6274BcD6CB1aa8A0fe20B80',
    [ChainId.AURORA]: '0xb1Da21B0531257a7E5aEfa0cd3CbF23AfC674cE1',
    [ChainId.CRONOS]: '0x7Bb5c7e3bF0B2a28fA26359667110bB974fF9359',
    [ChainId.METIS]: '0x0b5740c6b4a97f90eF2F0220651Cca420B868FfB',
    [ChainId.DFK]: '0x360d6DD540E3448371876662FBE7F1aCaf08c5Ab',
    [ChainId.TERRA]: 'uusd',
  },
  decimals: 6,
  symbol: 'UST',
  name: 'TerraUSD',
  logo: ustLogo,
  description: 'ERC-20 Wrapped form of UST',
  swapableType: 'UST',
})

export const MIM = new Token({
  addresses: {
    [ChainId.FANTOM]: '0x82f0b8b456c1a451378467398982d4834b6829c1',
    [ChainId.ARBITRUM]: '0xfea7a6a0b346362bf88a9e4a88416b77a57d6c2a',
  },
  decimals: 18,
  symbol: 'MIM',
  name: 'Magic Internet Money',
  logo: mimLogo,
  swapableType: 'USD',
})

export const WETH = new Token({
  addresses: {
    [ChainId.ETH]: '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
    [ChainId.ARBITRUM]: '0x82af49447d8a07e3bd95bd0d56f35241523fbab1',
    [ChainId.BOBA]: '0xd203De32170130082896b4111eDF825a4774c18E',
    [ChainId.OPTIMISM]: '0x121ab82b49B2BC4c7901CA46B8277962b4350204',
    [ChainId.BASE]: '0x4200000000000000000000000000000000000006',
    // [ChainId.AVALANCHE]: '0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab'
  },
  decimals: 18,
  symbol: 'WETH', // SHOULD BE WETH
  name: 'Wrapped ETH',
  logo: wethLogo,
  description: 'ERC-20 Wrapped form of ETH',
  swapableType: 'ETH',
})

export const WETHE = new Token({
  addresses: {
    [ChainId.AVALANCHE]: '0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab',
  },
  decimals: 18,
  symbol: 'WETH.e', // SHOULD BE WETH
  name: 'Wrapped ETH',
  logo: wethLogo,
  description: 'ERC-20 Wrapped form of ETH',
  swapableType: 'ETH',
})

/**
 * WETH on Moonbeam is nETH on moonbeam.
 * is this stupid & annoying - yes
 */
export const WETHBEAM = new Token({
  addresses: {
    [ChainId.MOONBEAM]: '0x3192Ae73315c3634Ffa217f71CF6CBc30FeE349A',
  },
  decimals: 18,
  symbol: 'WETH ',
  name: 'Wrapped ETH',
  logo: wethLogo,
  description: 'ERC-20 Wrapped form of ETH on Moonbeam',
  swapableType: 'ETH',
})

export const AVWETH = new Token({
  addresses: {
    [ChainId.AVALANCHE]: '0x53f7c5869a859f0aec3d334ee8b4cf01e3492f21',
  },
  decimals: 18,
  symbol: 'AVWETH', // AVALANCHE AAVE WETH
  name: 'Aave Wrapped ETH',
  logo: avwethLogo,
  description: 'Aave Wrapped form of ETH',
  swapableType: 'ETH',
})

export const ONEETH = new Token({
  addresses: {
    [ChainId.HARMONY]: '0x6983d1e6def3690c4d616b13597a09e6193ea013',
  },
  decimals: 18,
  symbol: '1ETH', // SHOULD BE WETH
  name: 'Harmony ETH',
  logo: wethLogo,
  description: 'Harmony ERC-20 Wrapped form of ETH',
  swapableType: 'ETH',
})

export const FTMETH = new Token({
  addresses: {
    [ChainId.FANTOM]: '0x74b23882a30290451A17c44f4F05243b6b58C76d',
  },
  decimals: 18,
  symbol: 'ETH ',
  name: 'Wrapped ETH',
  logo: wethLogo,
  description: 'Fantom Wrapped form of ETH',
  swapableType: 'ETH',
})

export const WFTM = new Token({
  addresses: {
    [ChainId.FANTOM]: '0x21be370D5312f44cB42ce377BC9b8a0cEF1A4C83',
    [ChainId.DFK]: '0x2Df041186C844F8a2e2b63F16145Bc6Ff7d23E25',
  },
  decimals: 18,
  symbol: 'WFTM ',
  name: 'Wrapped Fantom',
  logo: ftmLogo,
  description: 'Fantom Wrapped form of Fantom',
  swapableType: 'FTM',
})

export const CANTOETH = new Token({
  addresses: {
    [ChainId.CANTO]: '0x5FD55A1B9FC24967C4dB09C513C3BA0DFa7FF687',
  },
  decimals: 18,
  symbol: 'ETH ',
  logo: wethLogo,
  name: 'Wrapped ETH',
  description: 'Canto Wrapped form of ETH',
  swapableType: 'ETH',
})

export const METISETH = new Token({
  addresses: {
    [ChainId.METIS]: '0x420000000000000000000000000000000000000A',
  },
  decimals: 18,
  symbol: 'ETH ',
  name: 'Wrapped ETH',
  logo: wethLogo,
  description: 'Metis Wrapped form of ETH',
  swapableType: 'ETH',
})

export const SYN = new Token({
  addresses: {
    [ChainId.ETH]: '0x0f2d719407fdbeff09d87557abb7232601fd9f29',
    [ChainId.BSC]: '0xa4080f1778e69467e905b8d6f72f6e441f9e9484',
    [ChainId.POLYGON]: '0xf8f9efc0db77d8881500bb06ff5d6abc3070e695',
    [ChainId.FANTOM]: '0xE55e19Fb4F2D85af758950957714292DAC1e25B2', // yes this is same as avax swap addr, no its not error
    [ChainId.ARBITRUM]: '0x080f6aed32fc474dd5717105dba5ea57268f46eb',
    [ChainId.AVALANCHE]: '0x1f1E7c893855525b303f99bDF5c3c05Be09ca251',
    [ChainId.HARMONY]: '0xE55e19Fb4F2D85af758950957714292DAC1e25B2',
    [ChainId.BOBA]: '0xb554A55358fF0382Fb21F0a478C3546d1106Be8c',
    [ChainId.METIS]: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
    [ChainId.MOONRIVER]: '0xd80d8688b02B3FD3afb81cDb124F188BB5aD0445',
    [ChainId.MOONBEAM]: '0xF44938b0125A6662f9536281aD2CD6c499F22004',
    [ChainId.OPTIMISM]: '0x5A5fFf6F753d7C11A56A52FE47a177a87e431655',
    [ChainId.CRONOS]: '0xFD0F80899983b8D46152aa1717D76cba71a31616',
    [ChainId.AURORA]: '0xd80d8688b02B3FD3afb81cDb124F188BB5aD0445',
    [ChainId.DOGECHAIN]: '0xDfA53EeBA61D69E1D2b56b36d78449368F0265c1',
    [ChainId.CANTO]: '0x555982d2E211745b96736665e19D9308B615F78e',
    [ChainId.BASE]: '0x432036208d2717394d2614d6697c46DF3Ed69540',
  },
  decimals: 18,
  symbol: 'SYN',
  name: 'Synapse',
  logo: synapseLogo,
  description: 'SYN is the base token behind synapse',
  swapableType: 'SYN',
})

export const FRAX = new Token({
  addresses: {
    [ChainId.ETH]: '0x853d955acef822db058eb8505911ed77f175b99e',
    [ChainId.MOONRIVER]: '0xE96AC70907ffF3Efee79f502C985A7A21Bce407d',
    [ChainId.MOONBEAM]: '0xDd47A348AB60c61Ad6B60cA8C31ea5e00eBfAB4F',
    [ChainId.HARMONY]: '0x1852F70512298d56e9c8FDd905e02581E04ddb2a',
    [ChainId.DOGECHAIN]: '0x10D70831f9C3c11c5fe683b2f1Be334503880DB6',
    [ChainId.ARBITRUM]: '0x17fc002b466eec40dae837fc4be5c67993ddbd6f',
  },
  decimals: 18,
  symbol: 'FRAX',
  name: 'Frax',
  logo: fraxLogo,
  description: 'Frax',
  swapableType: 'FRAX',
})

export const SYN_FRAX = new Token({
  addresses: {
    // [ChainId.FANTOM]:    '0x1852F70512298d56e9c8FDd905e02581E04ddb2a',
    [ChainId.MOONRIVER]: '0xE96AC70907ffF3Efee79f502C985A7A21Bce407d',
    [ChainId.MOONBEAM]: '0xDd47A348AB60c61Ad6B60cA8C31ea5e00eBfAB4F',
    [ChainId.HARMONY]: '0x1852F70512298d56e9c8FDd905e02581E04ddb2a',
  },
  decimals: 18,
  symbol: 'synFRAX',
  name: 'Synapse Frax',
  logo: synapseLogo,
  description: 'Frax',
})

/**
 * nUSD is the token involved in the bridge.
 */
export const NUSD = new Token({
  addresses: {
    [ChainId.BSC]: '0x23b891e5c62e0955ae2bd185990103928ab817b3',
    [ChainId.ETH]: '0x1B84765dE8B7566e4cEAF4D0fD3c5aF52D3DdE4F',
    [ChainId.CRONOS]: '0x396c9c192dd323995346632581BEF92a31AC623b',
    [ChainId.OPTIMISM]: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
    [ChainId.POLYGON]: '0xb6c473756050de474286bed418b77aeac39b02af',
    [ChainId.FANTOM]: '0xED2a7edd7413021d440b09D654f3b87712abAB66',
    [ChainId.AVALANCHE]: '0xCFc37A6AB183dd4aED08C204D1c2773c0b1BDf46',
    [ChainId.ARBITRUM]: '0x2913E812Cf0dcCA30FB28E6Cac3d2DCFF4497688',
    [ChainId.HARMONY]: '0xED2a7edd7413021d440b09D654f3b87712abAB66',
    [ChainId.BOBA]: '0x6B4712AE9797C199edd44F897cA09BC57628a1CF',
    [ChainId.AURORA]: '0x07379565cD8B0CaE7c60Dc78e7f601b34AF2A21c',
    [ChainId.METIS]: '0x961318Fc85475E125B99Cc9215f62679aE5200aB',
    [ChainId.DFK]: '0x3AD9DFE640E1A9Cc1D9B0948620820D975c3803a',
    [ChainId.CANTO]: '0xD8836aF2e565D3Befce7D906Af63ee45a57E8f80',
  },
  decimals: 18,
  symbol: 'nUSD',
  name: 'Synapse nUSD',
  logo: nusdLogo,
  description: 'nUSD',
  swapableType: 'USD',
})

export const NOTE = new Token({
  addresses: {
    [ChainId.CANTO]: '0x4e71a2e537b7f9d9413d3991d37958c0b5e1e503',
  },
  decimals: 18,
  symbol: 'NOTE',
  name: 'Canto Note',
  logo: noteLogo,
  description: 'NOTE',
  swapableType: 'USD',
})

export const DFK_USDC = new Token({
  addresses: {
    [ChainId.DFK]: NUSD.addresses[ChainId.DFK],
  },
  decimals: {
    [ChainId.DFK]: 18,
  },
  symbol: 'USDC ', // SPACE VERY IMPORTANT
  name: 'USD Circle',
  logo: usdcLogo,
  description: '',
  swapableType: 'USD',
})

/**
 * nETH is the token involved in the bridge. it is backed by internet monies...
 */
export const NETH = new Token({
  addresses: {
    [ChainId.FANTOM]: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
    [ChainId.ARBITRUM]: '0x3ea9B0ab55F34Fb188824Ee288CeaEfC63cf908e',
    [ChainId.BOBA]: '0x96419929d7949D6A801A6909c145C8EEf6A40431',
    [ChainId.OPTIMISM]: '0x809DC529f07651bD43A172e8dB6f4a7a0d771036',
    [ChainId.AVALANCHE]: '0x19E1ae0eE35c0404f835521146206595d37981ae',
    [ChainId.HARMONY]: '0x0b5740c6b4a97f90eF2F0220651Cca420B868FfB',
    [ChainId.MOONBEAM]: '0x3192Ae73315c3634Ffa217f71CF6CBc30FeE349A', // THIS OVERLAPS WITH WETHBEAM
    [ChainId.METIS]: '0x931B8f17764362A3325D30681009f0eDd6211231',
    [ChainId.KLAYTN]: '0xCD6f29dC9Ca217d0973d3D21bF58eDd3CA871a86',
    [ChainId.DOGECHAIN]: '0x9F4614E4Ea4A0D7c4B1F946057eC030beE416cbB',
    [ChainId.CANTO]: '0x09fEC30669d63A13c666d2129230dD5588E2e240',
    [ChainId.BASE]: '0xb554A55358fF0382Fb21F0a478C3546d1106Be8c',
  },
  decimals: 18,
  symbol: 'nETH',
  name: 'Synapse nETH',
  logo: nethLogo,
  description: 'nETH',
  swapableType: 'ETH',
})

export const KLAYTN_WETH = new Token({
  addresses: {
    [ChainId.KLAYTN]: NETH.addresses[ChainId.KLAYTN],
    [ChainId.DOGECHAIN]: NETH.addresses[ChainId.DOGECHAIN],
  },
  decimals: {
    [ChainId.KLAYTN]: 18,
    [ChainId.DOGECHAIN]: 18,
  },
  symbol: 'WETH ', // SPACE VERY IMPORTANT
  name: 'Wrapped ETH',
  logo: ethLogo,
  description: '',
  swapableType: 'ETH',
})

export const ETH = new Token({
  addresses: {
    [ChainId.ETH]: '',
    [ChainId.BOBA]: '',
    [ChainId.ARBITRUM]: '',
    [ChainId.OPTIMISM]: '',
    [ChainId.BASE]: '',
    [ChainId.DFK]: '0xfBDF0E31808d0aa7b9509AA6aBC9754E48C58852',
    [ChainId.CANTO]: '0x5FD55A1B9FC24967C4dB09C513C3BA0DFa7FF687',
  },
  decimals: 18,
  symbol: 'ETH',
  name: 'Ethereum',
  logo: ethLogo,
  description: 'ETH',
  isNative: true,
  swapableType: 'ETH',
})

export const MOVR = new Token({
  addresses: {
    [ChainId.MOONRIVER]: '',
  },
  decimals: 18,
  symbol: 'MOVR',
  name: 'MOVR',
  logo: movrLogo,
  description: 'Moonriver',
  isNative: true,
  swapableType: 'MOVR',
})

export const AVAX = new Token({
  addresses: {
    [ChainId.AVALANCHE]: '',
    [ChainId.KLAYTN]: '0xcd8fe44a29db9159db36f96570d7a4d91986f528',
  },
  decimals: 18,
  symbol: 'AVAX',
  name: 'AVAX',
  logo: avaxLogo,
  description: 'AVAX',
  isNative: true,
  swapableType: 'AVAX',
})

export const WMOVR = new Token({
  addresses: {
    [ChainId.MOONBEAM]: '0x1d4C2a246311bB9f827F4C768e277FF5787B7D7E',
    [ChainId.MOONRIVER]: '0x98878b06940ae243284ca214f92bb71a2b032b8a',
  },
  decimals: 18,
  symbol: 'MOVR', // SHOULD BE WETH
  name: 'Wrapped MOVR',
  logo: movrLogo,
  description: 'ERC-20 Wrapped form of MOVR',
  swapableType: 'MOVR',
})

export const WAVAX = new Token({
  addresses: {
    [ChainId.AVALANCHE]: '0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7',
    [ChainId.DFK]: '0xB57B60DeBDB0b8172bb6316a9164bd3C695F133a',
    [ChainId.MOONBEAM]: '0xA1f8890E39b4d8E33efe296D698fe42Fb5e59cC3',
    [ChainId.HARMONY]: '0xD9eAA386cCD65F30b77FF175F6b52115FE454fD6',
  },
  decimals: 18,
  symbol: 'AVAX', // SHOULD BE WETH
  name: 'Wrapped AVAX',
  logo: avaxLogo,
  description: 'ERC-20 Wrapped form of AVAX',
  swapableType: 'AVAX',
})

export const SYNAVAX = new Token({
  addresses: {
    [ChainId.AVALANCHE]: '0xD9eAA386cCD65F30b77FF175F6b52115FE454fD6',
    [ChainId.HARMONY]: '0xD9eAA386cCD65F30b77FF175F6b52115FE454fD6',
  },
  decimals: 18,
  symbol: 'synAVAX', // SHOULD BE WETH
  name: 'Wrapped AVAX',
  logo: avaxLogo,
  description: 'ERC-20 Wrapped form of AVAX',
  swapableType: 'AVAX',
})

export const MULTIAVAX = new Token({
  addresses: {
    [ChainId.HARMONY]: '0xb12c13e66ade1f72f71834f2fc5082db8c091358',
  },
  decimals: 18,
  symbol: 'multiAVAX', // SHOULD BE WETH
  name: 'AnySwap Wrapped AVAX',
  logo: avaxLogo,
  description: 'ERC-20 Wrapped form of AVAX',
  swapableType: 'AVAX',
})

export const JEWEL = new Token({
  addresses: {
    [ChainId.DFK]: '',
  },
  decimals: 18,
  symbol: 'JEWEL',
  name: 'JEWEL',
  logo: jewelLogo,
  description: 'JEWEL',
  isNative: true,
  swapableType: 'JEWEL',
})

export const WJEWEL = new Token({
  addresses: {
    [ChainId.DFK]: '0xCCb93dABD71c8Dad03Fc4CE5559dC3D89F67a260', // from actual jewl
    [ChainId.HARMONY]: '0x72cb10c6bfa5624dd07ef608027e366bd690048f', // from harmony jewel?
    [ChainId.KLAYTN]: '0x30c103f8f5a3a732dfe2dce1cc9446f545527b43',
  },
  decimals: 18,
  symbol: 'JEWEL ', // THE SPACES ARE VERY IMPORTANT
  name: 'JEWEL ', // THE SPACES ARE VERY IMPORTANT
  logo: jewelLogo,
  description: 'JEWEL',
  swapableType: 'JEWEL',
})

export const KLAY = new Token({
  addresses: {
    [ChainId.KLAYTN]: '',
    [ChainId.DFK]: '0x97855Ba65aa7ed2F65Ed832a776537268158B78a',
  },
  decimals: 18,
  symbol: 'KLAY',
  name: 'KLAY',
  logo: klaytnLogo,
  description: 'KLAY',
  isNative: true,
  swapableType: 'KLAY',
})

export const WKLAY = new Token({
  addresses: {
    [ChainId.DFK]: '0x97855Ba65aa7ed2F65Ed832a776537268158B78a',
    [ChainId.KLAYTN]: '0x5819b6af194a78511c79c85ea68d2377a7e9335f',
  },
  decimals: 18,
  symbol: 'KLAY ', // THE SPACES ARE VERY IMPORTANT
  name: 'KLAY ', // THE SPACES ARE VERY IMPORTANT
  logo: klaytnLogo,
  description: 'KLAY',
  swapableType: 'KLAY',
})

export const SYNJEWEL = new Token({
  addresses: {
    [ChainId.AVALANCHE]: '0x997Ddaa07d716995DE90577C123Db411584E5E46',
    [ChainId.HARMONY]: '0x28b42698Caf46B4B012CF38b6C75867E0762186D',
  },
  decimals: 18,
  symbol: 'JEWEL  ', // THE SPACES ARE VERY IMPORTANT
  name: 'JEWEL  ', // THE SPACES ARE VERY IMPORTANT
  logo: jewelLogo,
  description: 'ERC-20 Wrapped form of JEWEL',
  swapableType: 'JEWEL',
})

export const XJEWEL = new Token({
  addresses: {
    [ChainId.DFK]: '0x77f2656d04E158f915bC22f07B779D94c1DC47Ff',
    [ChainId.HARMONY]: '0xA9cE83507D872C5e1273E745aBcfDa849DAA654F',
  },
  decimals: 18,
  symbol: 'xJEWEL',
  name: 'xJEWEL',
  logo: jewelLogo,
  description: 'ERC-20 Wrapped form of xJEWEL',
  swapableType: 'XJEWEL',
})

// synJEWEL harmony
// synJEWEL avax
// JEWEL DFK chain

// xJEWEL Harmony
// xJEWEL DFK Chain

// synAVAX Harmony
// AVAX AVAX
// AVAX DFKChain

// synJEWEL harmony
// xJEWEL Harmony
// synAVAX Harmony

// synJEWEL avax
// AVAX AVAX

// JEWEL DFK chain
// xJEWEL DFK Chain
// AVAX DFKChain

export const WMATIC = new Token({
  addresses: {
    [ChainId.POLYGON]: '0x9b17bAADf0f21F03e35249e0e59723F34994F806',
  },
  decimals: 18,
  symbol: 'MATIC', // SHOULD BE WETH
  name: 'Wrapped MATIC',
  description: 'ERC-20 Wrapped form of MATIC',
  swapableType: 'MATIC',
  logo: maticLogo,
})

export const MATIC = new Token({
  addresses: {
    [ChainId.DFK]: '0xD17a41Cd199edF1093A9Be4404EaDe52Ec19698e',
  },
  decimals: 18,
  symbol: 'MATIC', // SHOULD BE WETH
  name: 'MATIC',
  description: 'ERC-20 MATIC',
  swapableType: 'MATIC',
  logo: maticLogo,
})

export const WBNB = new Token({
  addresses: {
    [ChainId.BSC]: '0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c',
    [ChainId.DOGECHAIN]: '0x1fC532187B4848d2F9c564531b776A4F8e11201d',
  },
  decimals: 18,
  symbol: 'BNB', // SHOULD BE WETH
  name: 'Wrapped BNB',
  description: 'ERC-20 Wrapped form of BNB',
  swapableType: 'BNB',
})

export const DEPRECATED_WKLAY = new Token({
  addresses: {
    [ChainId.KLAYTN]: '0x5819b6af194a78511c79c85ea68d2377a7e9335f',
  },
  decimals: 18,
  symbol: 'WKLAY',
  name: 'Deprecated Wrapped Klay',
  description: 'ERC-20 Wrapped form of KLAY',
  swapableType: 'KLAY',
})

export const PEPE = new Token({
  addresses: {
    [ChainId.ETH]: '0x6982508145454ce325ddbe47a25d4ec3d2311933',
    [ChainId.ARBITRUM]: '0xA54B8e178A49F8e5405A4d44Bb31F496e5564A05',
    [ChainId.BSC]: '0xd2b6F20aa2611e8a7a18e5EeC58ca8369f5D356b',
  },
  decimals: 18,
  symbol: 'PEPE',
  name: 'Pepe',
  logo: pepeLogo,
  description: 'PEPE',
  swapableType: 'PEPE',
})
export const BASIC_TOKENS_BY_CHAIN = {
  [ChainId.ETH]: [
    USDC,
    USDT,
    DAI,
    NUSD,
    SYN,
    WETH,
    ETH,
    HIGHSTREET,
    DOG,
    FRAX,
    GOHM,
    // UST,
    NEWO,
    SDT,
    USDB,
    VSTA,
    SFI,
    H2O,
    AGEUR,
    WBTC,
    KLAYTN_USDC,
    KLAYTN_USDT,
    KLAYTN_DAI,
    LINK,
    PEPE,
    UNIDX,
    LUSD,
    CRVUSD,
  ],
  [ChainId.BSC]: [
    BUSD,
    USDC,
    USDT,
    NUSD,
    SYN,
    NFD,
    HIGHSTREET,
    DOG,
    JUMP,
    GOHM,
    // UST,
    H2O,
    USDB,
    DOGECHAIN_BUSD,
    PEPE,
  ],
  [ChainId.POLYGON]: [
    USDC,
    USDT,
    DAI,
    NUSD,
    SYN,
    GOHM,
    DOG,
    NFD,
    // UST,
    USDB,
    H2O,
    WMATIC,
  ],
  [ChainId.FANTOM]: [
    MIM,
    USDC,
    USDT,
    NUSD,
    // FRAX,
    FTMETH,
    SYN,
    JUMP,
    GOHM,
    FTMETH,
    NETH,
    // UST,
    SDT,
    USDB,
    WFTM,
    UNIDX,
  ],
  [ChainId.BOBA]: [
    SYN,
    NETH,
    WETH,
    ETH,
    USDC,
    USDT,
    DAI,
    NUSD,
    GOHM,
    // UST,
  ],
  [ChainId.MOONBEAM]: [
    // FRAX, #temp
    // WETHBEAM, #temp
    GOHM,
    SOLAR,
    WMOVR,
    WAVAX,
    SYN,
    H2O,
    // UST,
  ],
  [ChainId.MOONRIVER]: [
    SYN,
    FRAX,
    SOLAR,
    GOHM,
    USDB,
    MOVR,
    WMOVR,
    H2O,
    // UST,
  ],
  [ChainId.ARBITRUM]: [
    NETH,
    SYN,
    CCTP_USDC,
    WETH,
    ETH,
    USDC,
    USDT,
    DAI,
    MIM,
    NUSD,
    GOHM,
    GMX,
    UST,
    H2O,
    NEWO,
    VSTA,
    SDT,
    L2DAO,
    PLS,
    AGEUR,
    UNIDX,
    PEPE,
    FRAX,
  ],
  [ChainId.AVALANCHE]: [
    USDC,
    USDT,
    CCTP_USDC,
    // Note that this is Dai.e on Avalanche
    DAI,
    WETHE,
    NETH,
    NUSD,
    SDT,
    SYN,
    NFD,
    GOHM,
    GMX,
    AVAX,
    WAVAX,
    // UST,
    H2O,
    NEWO,
    USDB,
    SFI,
    AVWETH,
    USDTE,
    BTCB,

    SYNJEWEL,
  ],
  [ChainId.DFK]: [
    JEWEL,
    WJEWEL,
    XJEWEL,
    WAVAX,
    DFK_USDC,
    UST,
    KLAY,
    WFTM,
    ETH,
    BTCB,
  ],
  [ChainId.AURORA]: [
    USDC,
    USDT,
    NUSD,
    SYN,
    // UST,
  ],
  [ChainId.HARMONY]: [
    USDC,
    USDT,
    DAI,
    NUSD,
    ONEETH,
    FRAX,
    SDT,
    SYN,
    GOHM,
    NETH,
    SYN_FRAX,
    // UST,

    WJEWEL,
    SYNJEWEL,
    XJEWEL,
    SYNAVAX,
    MULTIAVAX,
  ],
  [ChainId.OPTIMISM]: [
    NETH,
    SYN,
    WETH,
    ETH,
    USDC,
    NUSD,
    UST,
    GOHM,
    H2O,
    L2DAO,
    PLS,
    AGEUR,
    USDT,
    CCTP_USDC,
    UNIDX,
    SUSD,
    DAI,
  ],
  [ChainId.TERRA]: [UST],
  [ChainId.CRONOS]: [
    NUSD,
    DAI,
    USDC,
    USDT,
    GOHM,
    // UST,
    SYN,
  ],
  [ChainId.METIS]: [USDC, METISETH, NETH, NUSD, GOHM, UST, JUMP, SYN],
  [ChainId.KLAYTN]: [
    KLAYTN_USDC,
    KLAYTN_USDT,
    KLAYTN_DAI,
    KLAYTN_WETH,
    LINK,
    WBTC,
    KLAYTN_oUSDT,
    WJEWEL,
    WAVAX,
    WKLAY,
    AVAX,
  ],
  [ChainId.DOGECHAIN]: [
    SYN,
    FRAX,
    KLAYTN_USDC,
    KLAYTN_USDT,
    KLAYTN_DAI,
    KLAYTN_WETH,
    WBTC,
    DOGECHAIN_BUSD,
    NFD,
  ],
  [ChainId.BASE]: [
    SYN,
    ETH,
    NETH,
    WETH,
    // USDC,
    DAI,
    USDBC,
    CRVUSD,
    CCTP_USDC,
    UNIDX,
  ],
  [ChainId.CANTO]: [SYN, NUSD, CANTOETH, NETH, NOTE, USDC, USDT, ETH],
}

const TOKEN_HASH_MAP = {}

for (const [chainId, tokensOnChain] of _.toPairs(BASIC_TOKENS_BY_CHAIN)) {
  TOKEN_HASH_MAP[chainId] = {}
  for (const token of tokensOnChain) {
    TOKEN_HASH_MAP[chainId][_.toLower(token.addresses[chainId])] = token
  }
}

TOKEN_HASH_MAP[ChainId.AVALANCHE][
  _.toLower(GMX.wrapperAddresses[ChainId.AVALANCHE])
] = GMX

export { TOKEN_HASH_MAP }
