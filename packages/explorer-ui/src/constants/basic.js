import avaxLogo from 'assets/icons/avalanche.svg'
import avwethLogo from 'assets/icons/avweth.svg'
import busdLogo from 'assets/icons/busd.svg'
import daiLogo from 'assets/icons/dai.png'
import ethLogo from 'assets/icons/eth.svg'
import fraxLogo from 'assets/icons/frax.svg'
import jewelLogo from 'assets/icons/jewel.svg'
import mimLogo from 'assets/icons/mim.svg'
import movrLogo from 'assets/icons/moonriver.jpeg'
import nethLogo from 'assets/icons/neth.svg'
import nusdLogo from 'assets/icons/nusd.svg'
import synapseLogo from 'assets/icons/synapse.svg'
import usdcLogo from 'assets/icons/usdc.svg'
import usdtLogo from 'assets/icons/usdt.svg'
import ustLogo from 'assets/icons/ust.png'
import wethLogo from 'assets/icons/weth.svg'
import wbtcLogo from 'assets/icons/wbtc.svg'
import _ from 'lodash'
import { CHAIN_ID } from './chains'
import { DOG, GMX, GOHM, HIGHSTREET, JUMP, NFD, SOLAR } from './mintable'
import { Token } from './Token'

export const BUSD = new Token({
  addresses: {
    [CHAIN_ID.BSC]: '0xe9e7cea3dedca5984780bafc599bd69add087d56',
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
    [CHAIN_ID.BSC]: '0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d',
    [CHAIN_ID.ETH]: '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
    [CHAIN_ID.POLYGON]: '0x2791bca1f2de4661ed88a30c99a7a9449aa84174',
    [CHAIN_ID.FANTOM]: '0x04068da6c83afcfa0e13ba15a6696662335d5b75',
    [CHAIN_ID.AVALANCHE]: '0xA7D7079b0FEaD91F3e65f86E8915Cb59c1a4C664',
    [CHAIN_ID.ARBITRUM]: '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8',
    [CHAIN_ID.HARMONY]: '0x985458e523db3d53125813ed68c274899e9dfab4',
    [CHAIN_ID.BOBA]: '0x66a2A913e447d6b4BF33EFbec43aAeF87890FBbc',
    [CHAIN_ID.AURORA]: '0xB12BFcA5A55806AaF64E99521918A4bf0fC40802',
    [CHAIN_ID.METIS]: '0xEA32A96608495e54156Ae48931A7c20f0dcc1a21',
    [CHAIN_ID.DFK]: '0x3ad9dfe640e1a9cc1d9b0948620820d975c3803a',
    [CHAIN_ID.KLAYTN]: '0x6270b58be569a7c0b8f47594f191631ae5b2c86c',
  },
  decimals: {
    [CHAIN_ID.BSC]: 18,
    [CHAIN_ID.ETH]: 6,
    [CHAIN_ID.POLYGON]: 6,
    [CHAIN_ID.FANTOM]: 6,
    [CHAIN_ID.AVALANCHE]: 6,
    [CHAIN_ID.ARBITRUM]: 6,
    [CHAIN_ID.HARMONY]: 6,
    [CHAIN_ID.BOBA]: 6,
    [CHAIN_ID.AURORA]: 6,
    [CHAIN_ID.METIS]: 6,
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

export const USDT = new Token({
  addresses: {
    [CHAIN_ID.BSC]: '0x55d398326f99059ff775485246999027b3197955',
    [CHAIN_ID.ETH]: '0xdac17f958d2ee523a2206206994597c13d831ec7',
    [CHAIN_ID.POLYGON]: '0xc2132d05d31c914a87c6611c10748aeb04b58e8f',
    [CHAIN_ID.AVALANCHE]: '0xc7198437980c041c805a1edcba50c1ce5db95118',
    [CHAIN_ID.HARDHAT]: '0x9A9f2CCfdE556A7E9Ff0848998Aa4a0CFD8863AE',
    [CHAIN_ID.ARBITRUM]: '0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9',
    [CHAIN_ID.FANTOM]: '0x049d68029688eabf473097a2fc38ef61633a3c7a',
    [CHAIN_ID.HARMONY]: '0x3c2b8be99c50593081eaa2a724f0b8285f5aba8f',
    [CHAIN_ID.BOBA]: '0x5DE1677344D3Cb0D7D465c10b72A8f60699C062d',
    [CHAIN_ID.AURORA]: '0x4988a896b1227218e4A686fdE5EabdcAbd91571f',
    [CHAIN_ID.KLAYTN]: '0xd6dab4cff47df175349e6e7ee2bf7c40bb8c05a3',
  },
  decimals: {
    [CHAIN_ID.BSC]: 18,
    [CHAIN_ID.ETH]: 6,
    [CHAIN_ID.POLYGON]: 6,
    [CHAIN_ID.AVALANCHE]: 6,
    [CHAIN_ID.ARBITRUM]: 6,
    [CHAIN_ID.FANTOM]: 6,
    [CHAIN_ID.HARMONY]: 6,
    [CHAIN_ID.BOBA]: 6,
    [CHAIN_ID.AURORA]: 6,
    [CHAIN_ID.KLAYTN]: 6,
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

export const DAI = new Token({
  addresses: {
    [CHAIN_ID.BSC]: '0x1af3f329e8be154074d8769d1ffa4ee058b1dbc3',
    [CHAIN_ID.ETH]: '0x6b175474e89094c44da98b954eedeac495271d0f',
    [CHAIN_ID.POLYGON]: '0x8f3cf7ad23cd3cadbd9735aff958023239c6a063',
    [CHAIN_ID.AVALANCHE]: '0xd586E7F844cEa2F87f50152665BCbc2C279D8d70',
    [CHAIN_ID.ARBITRUM]: '0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1',
    [CHAIN_ID.HARMONY]: '0xef977d2f931c1978db5f6747666fa1eacb0d0339',
    [CHAIN_ID.BOBA]: '0xf74195Bb8a5cf652411867c5C2C5b8C2a402be35',
    [CHAIN_ID.KLAYTN]: '0x078db7827a5531359f6cb63f62cfa20183c4f10c',
  },
  decimals: 18,
  symbol: 'DAI',
  name: 'Dai',
  logo: daiLogo,
  swapableType: 'USD',
})

export const UST = new Token({
  addresses: {
    // [CHAIN_ID.BSC]:       '0x23396cf899ca06c4472205fc903bdb4de249d6fc',
    // [CHAIN_ID.POLYGON]:   '0xed7a89e2d580bdda005ba4cdd64cf0da3c15a5eb',
    // [CHAIN_ID.AVALANCHE]: '0x195150ab664795128ac542f26b14c1a12e061ecf',
    // [CHAIN_ID.TERRA]:     'uusd',
    [CHAIN_ID.ETH]: '0x0261018Aa50E28133C1aE7a29ebdf9Bd21b878Cb',
    [CHAIN_ID.BSC]: '0xb7A6c5f0cc98d24Cf4B2011842e64316Ff6d042c',
    [CHAIN_ID.POLYGON]: '0x565098CBa693b3325f9fe01D41b7A1cd792Abab1',
    [CHAIN_ID.FANTOM]: '0xa0554607e477cdC9d0EE2A6b087F4b2DC2815C22',
    [CHAIN_ID.ARBITRUM]: '0x13780E6d5696DD91454F6d3BbC2616687fEa43d0',
    [CHAIN_ID.AVALANCHE]: '0xE97097dE8d6A17Be3c39d53AE63347706dCf8f43',
    [CHAIN_ID.HARMONY]: '0xa0554607e477cdC9d0EE2A6b087F4b2DC2815C22',
    [CHAIN_ID.BOBA]: '0x61A269a9506272D128d79ABfE8E8276570967f00',
    [CHAIN_ID.MOONRIVER]: '0xa9D0C0E124F53f4bE1439EBc35A9C73c0e8275fB',
    [CHAIN_ID.MOONBEAM]: '0x5CF84397944B9554A278870B510e86667681ff8D',
    [CHAIN_ID.OPTIMISM]: '0xFB21B70922B9f6e3C6274BcD6CB1aa8A0fe20B80',
    [CHAIN_ID.AURORA]: '0xb1Da21B0531257a7E5aEfa0cd3CbF23AfC674cE1',
    [CHAIN_ID.CRONOS]: '0x7Bb5c7e3bF0B2a28fA26359667110bB974fF9359',
    [CHAIN_ID.METIS]: '0x0b5740c6b4a97f90eF2F0220651Cca420B868FfB',
    [CHAIN_ID.TERRA]: 'uusd',
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
    [CHAIN_ID.FANTOM]: '0x82f0b8b456c1a451378467398982d4834b6829c1',
    [CHAIN_ID.ARBITRUM]: '0xfea7a6a0b346362bf88a9e4a88416b77a57d6c2a',
  },
  decimals: 18,
  symbol: 'MIM',
  name: 'Magic Internet Money',
  logo: mimLogo,
  swapableType: 'USD',
})

export const WETH = new Token({
  addresses: {
    [CHAIN_ID.ETH]: '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
    [CHAIN_ID.ARBITRUM]: '0x82af49447d8a07e3bd95bd0d56f35241523fbab1',
    [CHAIN_ID.BOBA]: '0xd203De32170130082896b4111eDF825a4774c18E',
    [CHAIN_ID.OPTIMISM]: '0x121ab82b49B2BC4c7901CA46B8277962b4350204',
    [CHAIN_ID.KLAYTN]: '0xCD6f29dC9Ca217d0973d3D21bF58eDd3CA871a86',
    // [CHAIN_ID.AVALANCHE]: '0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab'
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
    [CHAIN_ID.AVALANCHE]: '0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab',
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
    [CHAIN_ID.MOONBEAM]: '0x3192Ae73315c3634Ffa217f71CF6CBc30FeE349A',
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
    [CHAIN_ID.AVALANCHE]: '0x53f7c5869a859f0aec3d334ee8b4cf01e3492f21',
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
    [CHAIN_ID.HARMONY]: '0x6983d1e6def3690c4d616b13597a09e6193ea013',
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
    [CHAIN_ID.FANTOM]: '0x74b23882a30290451A17c44f4F05243b6b58C76d',
  },
  decimals: 18,
  symbol: 'ETH ', // AVALANCHE AAVE WETH
  name: 'Wrapped ETH',
  logo: wethLogo,
  description: 'Fantom Wrapped form of ETH',
  swapableType: 'ETH',
})

export const SYN = new Token({
  addresses: {
    [CHAIN_ID.ETH]: '0x0f2d719407fdbeff09d87557abb7232601fd9f29',
    [CHAIN_ID.BSC]: '0xa4080f1778e69467e905b8d6f72f6e441f9e9484',
    [CHAIN_ID.POLYGON]: '0xf8f9efc0db77d8881500bb06ff5d6abc3070e695',
    [CHAIN_ID.FANTOM]: '0xE55e19Fb4F2D85af758950957714292DAC1e25B2', // yes this is same as avax swap addr, no its not error
    [CHAIN_ID.ARBITRUM]: '0x080f6aed32fc474dd5717105dba5ea57268f46eb',
    [CHAIN_ID.AVALANCHE]: '0x1f1E7c893855525b303f99bDF5c3c05Be09ca251',
    [CHAIN_ID.HARMONY]: '0xE55e19Fb4F2D85af758950957714292DAC1e25B2',
    [CHAIN_ID.BOBA]: '0xb554A55358fF0382Fb21F0a478C3546d1106Be8c',
    [CHAIN_ID.METIS]: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
    [CHAIN_ID.MOONRIVER]: '0xd80d8688b02B3FD3afb81cDb124F188BB5aD0445',
    [CHAIN_ID.MOONBEAM]: '0xF44938b0125A6662f9536281aD2CD6c499F22004',
    [CHAIN_ID.OPTIMISM]: '0x5A5fFf6F753d7C11A56A52FE47a177a87e431655',
    [CHAIN_ID.CRONOS]: '0xFD0F80899983b8D46152aa1717D76cba71a31616',
    [CHAIN_ID.AURORA]: '0xd80d8688b02B3FD3afb81cDb124F188BB5aD0445',
    [CHAIN_ID.SYN]: '0xF5f3650f54dA85e4A4D8E490139C77275B167c53',
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
    [CHAIN_ID.ETH]: '0x853d955acef822db058eb8505911ed77f175b99e',
    [CHAIN_ID.MOONRIVER]: '0x1a93b23281cc1cde4c4741353f3064709a16197d',
    [CHAIN_ID.MOONBEAM]: '',
    [CHAIN_ID.HARMONY]: '0xFa7191D292d5633f702B0bd7E3E3BcCC0e633200',
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
    // [CHAIN_ID.FANTOM]:    '0x1852F70512298d56e9c8FDd905e02581E04ddb2a',
    [CHAIN_ID.MOONRIVER]: '0xE96AC70907ffF3Efee79f502C985A7A21Bce407d',
    [CHAIN_ID.MOONBEAM]: '0xDd47A348AB60c61Ad6B60cA8C31ea5e00eBfAB4F',
    [CHAIN_ID.HARMONY]: '0x1852F70512298d56e9c8FDd905e02581E04ddb2a',
    [CHAIN_ID.FANTOM]: '0x1852f70512298d56e9c8fdd905e02581e04ddb2a',
  },
  decimals: 18,
  symbol: 'synFRAX',
  name: 'Synapse Frax',
  logo: synapseLogo,
  description: 'Frax',
})

/**
 * nUSD is the token involved in the bridge. it is backed by pixie dust...
 */
export const NUSD = new Token({
  addresses: {
    [CHAIN_ID.BSC]: '0x23b891e5c62e0955ae2bd185990103928ab817b3',
    [CHAIN_ID.ETH]: '0x1B84765dE8B7566e4cEAF4D0fD3c5aF52D3DdE4F',
    [CHAIN_ID.POLYGON]: '0xb6c473756050de474286bed418b77aeac39b02af',
    [CHAIN_ID.FANTOM]: '0xED2a7edd7413021d440b09D654f3b87712abAB66',
    [CHAIN_ID.AVALANCHE]: '0xCFc37A6AB183dd4aED08C204D1c2773c0b1BDf46',
    [CHAIN_ID.ARBITRUM]: '0x2913E812Cf0dcCA30FB28E6Cac3d2DCFF4497688',
    [CHAIN_ID.HARMONY]: '0xED2a7edd7413021d440b09D654f3b87712abAB66',
    [CHAIN_ID.BOBA]: '0x6B4712AE9797C199edd44F897cA09BC57628a1CF',
    [CHAIN_ID.AURORA]: '0x07379565cD8B0CaE7c60Dc78e7f601b34AF2A21c',
    [CHAIN_ID.METIS]: '0x961318Fc85475E125B99Cc9215f62679aE5200aB',
  },
  decimals: 18,
  symbol: 'nUSD',
  name: 'Synapse nUSD',
  logo: nusdLogo,
  description: 'nUSD',
  swapableType: 'USD',
})

/**
 * nETH is the token involved in the bridge. it is backed by internet monies...
 */
export const NETH = new Token({
  addresses: {
    [CHAIN_ID.FANTOM]: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
    [CHAIN_ID.ARBITRUM]: '0x3ea9B0ab55F34Fb188824Ee288CeaEfC63cf908e',
    [CHAIN_ID.BOBA]: '0x96419929d7949D6A801A6909c145C8EEf6A40431',
    [CHAIN_ID.OPTIMISM]: '0x809DC529f07651bD43A172e8dB6f4a7a0d771036',
    [CHAIN_ID.AVALANCHE]: '0x19E1ae0eE35c0404f835521146206595d37981ae',
    [CHAIN_ID.HARMONY]: '0x0b5740c6b4a97f90eF2F0220651Cca420B868FfB',
    [CHAIN_ID.MOONBEAM]: '0x3192Ae73315c3634Ffa217f71CF6CBc30FeE349A', // THIS OVERLAPS WITH WETHBEAM
  },
  decimals: 18,
  symbol: 'nETH',
  name: 'Synapse nETH',
  logo: nethLogo,
  description: 'nETH',
  swapableType: 'ETH',
})

export const ETH = new Token({
  addresses: {
    [CHAIN_ID.ETH]: '',
    [CHAIN_ID.BOBA]: '',
    [CHAIN_ID.ARBITRUM]: '',
    [CHAIN_ID.OPTIMISM]: '',
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
    [CHAIN_ID.MOONRIVER]: '',
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
    [CHAIN_ID.AVALANCHE]: '',
    [CHAIN_ID.DFK]: '0xb57b60debdb0b8172bb6316a9164bd3c695f133a',
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
    [CHAIN_ID.MOONBEAM]: '0x1d4C2a246311bB9f827F4C768e277FF5787B7D7E',
    [CHAIN_ID.MOONRIVER]: '0x98878b06940ae243284ca214f92bb71a2b032b8a',
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
    [CHAIN_ID.AVALANCHE]: '0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7',
    [CHAIN_ID.MOONBEAM]: '0xA1f8890E39b4d8E33efe296D698fe42Fb5e59cC3',
  },
  decimals: 18,
  symbol: 'AVAX', // SHOULD BE WETH
  name: 'Wrapped AVAX',
  logo: avaxLogo,
  description: 'ERC-20 Wrapped form of AVAX',
  swapableType: 'AVAX',
})

export const WMATIC = new Token({
  addresses: {
    [CHAIN_ID.POLYGON]: '0x9b17bAADf0f21F03e35249e0e59723F34994F806',
  },
  decimals: 18,
  symbol: 'MATIC', // SHOULD BE WETH
  name: 'Wrapped MATIC',
  description: 'ERC-20 Wrapped form of MATIC',
  swapableType: 'MATIC',
})

export const WBNB = new Token({
  addresses: {
    [CHAIN_ID.BSC]: '0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c',
  },
  decimals: 18,
  symbol: 'BNB', // SHOULD BE WETH
  name: 'Wrapped BNB',
  description: 'ERC-20 Wrapped form of BNB',
  swapableType: 'BNB',
})

export const SYN_AVAX = new Token({
  name: 'Wrapped AVAX',
  symbol: 'synAVAX',
  logo: avaxLogo,
  decimals: 18,
  addresses: {
    [CHAIN_ID.HARMONY]: '0xD9eAA386cCD65F30b77FF175F6b52115FE454fD6',
  },
})

export const MULTI_AVAX = new Token({
  name: 'AnySwap/Multi Wrapped AVAX',
  symbol: 'multiAVAX',
  logo: avaxLogo,
  decimals: 18,
  addresses: {
    [CHAIN_ID.HARMONY]: '0xb12c13e66ade1f72f71834f2fc5082db8c091358',
  },
})

export const JEWEL = new Token({
  name: 'JEWEL',
  symbol: 'JEWEL',
  decimals: 18,
  addresses: {
    [CHAIN_ID.DFK]: '',
  },
  isNative: true,
  logo: jewelLogo,
})

export const WJEWEL = new Token({
  name: 'Wrapped JEWEL',
  symbol: 'wJEWEL',
  decimals: 18,
  addresses: {
    [CHAIN_ID.DFK]: '0xCCb93dABD71c8Dad03Fc4CE5559dC3D89F67a260',
    [CHAIN_ID.HARMONY]: '0x72Cb10C6bfA5624dD07Ef608027E366bd690048F',
  },
  logo: jewelLogo,
})

export const SYN_JEWEL = new Token({
  name: 'Synapse JEWEL',
  symbol: 'synJEWEL',
  decimals: 18,
  addresses: {
    [CHAIN_ID.AVALANCHE]: '0x997Ddaa07d716995DE90577C123Db411584E5E46',
    [CHAIN_ID.HARMONY]: '0x28b42698Caf46B4B012CF38b6C75867E0762186D',
  },
  logo: jewelLogo,
})

export const XJEWEL = new Token({
  name: 'xJEWEL',
  symbol: 'xJEWEL',
  decimals: 18,
  addresses: {
    [CHAIN_ID.DFK]: '0x77f2656d04E158f915bC22f07B779D94c1DC47Ff',
    [CHAIN_ID.HARMONY]: '0xA9cE83507D872C5e1273E745aBcfDa849DAA654F',
  },
  logo: jewelLogo,
})

export const WBTC = new Token({
  name: 'Wrapped Bitcoin',
  symbol: 'WBTC',
  decimals: 8,
  addresses: {
    [CHAIN_ID.KLAYTN]: '0xDCbacF3f7a069922E677912998c8d57423C37dfA',
  },
  logo: wbtcLogo,
})

export const BASIC_TOKENS_BY_CHAIN = {
  [CHAIN_ID.ETH]: [USDC, USDT, DAI, NUSD, SYN, WETH, ETH, HIGHSTREET, DOG, FRAX, GOHM, UST],
  [CHAIN_ID.BSC]: [BUSD, USDC, USDT, NUSD, SYN, NFD, HIGHSTREET, DOG, JUMP, GOHM, UST],
  [CHAIN_ID.POLYGON]: [USDC, USDT, DAI, NUSD, SYN, GOHM, DOG, NFD, UST],
  [CHAIN_ID.FANTOM]: [
    MIM,
    USDC,
    USDT,
    NUSD,
    // FRAX,
    SYN,
    JUMP,
    GOHM,
    FTMETH,
    NETH,
    UST,
    SYN_FRAX,
  ],
  [CHAIN_ID.BOBA]: [SYN, NETH, WETH, ETH, USDC, USDT, DAI, NUSD, GOHM, UST],
  [CHAIN_ID.MOONBEAM]: [
    // FRAX, #temp
    // WETHBEAM, #temp
    GOHM,
    SOLAR,
    WMOVR,
    WAVAX,
    SYN,
    UST,
  ],
  [CHAIN_ID.MOONRIVER]: [SYN, FRAX, SOLAR, GOHM, MOVR, WMOVR, UST, SYN_FRAX],
  [CHAIN_ID.ARBITRUM]: [
    NETH,
    SYN,
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
    // DOG
  ],
  [CHAIN_ID.AVALANCHE]: [USDC, USDT, DAI, WETHE, NETH, NUSD, SYN, NFD, GOHM, GMX, AVAX, WAVAX, UST, SYN_JEWEL],
  [CHAIN_ID.AURORA]: [USDC, USDT, NUSD, SYN, UST],
  [CHAIN_ID.HARMONY]: [USDC, USDT, DAI, NUSD, FRAX, SYN, GOHM, ONEETH, NETH, UST, SYN_FRAX, SYN_JEWEL, SYN_AVAX],
  [CHAIN_ID.OPTIMISM]: [NETH, SYN, WETH, ETH, UST, GOHM],
  [CHAIN_ID.TERRA]: [UST],
  [CHAIN_ID.CRONOS]: [GOHM, UST, SYN],
  [CHAIN_ID.METIS]: [USDC, NUSD, GOHM, UST, SYN],
  [CHAIN_ID.DFK]: [XJEWEL, AVAX, WJEWEL, USDC, SYN],
  [CHAIN_ID.KLAYTN]: [SYN, USDC, USDT, WETH, DAI, WBTC],
}

let TOKEN_HASH_MAP = {}

for (const [CHAIN_ID, tokensOnChain] of _.toPairs(BASIC_TOKENS_BY_CHAIN)) {
  TOKEN_HASH_MAP[CHAIN_ID] = {}
  for (const token of tokensOnChain) {
    TOKEN_HASH_MAP[CHAIN_ID][_.toLower(token.addresses[CHAIN_ID])] = token
  }
}

// TOKEN_HASH_MAP[CHAIN_ID.AVALANCHE][_.toLower(GMX.wrapperAddresses[CHAIN_ID.AVALANCHE])] = GMX

export { TOKEN_HASH_MAP }
