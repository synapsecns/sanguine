import synapseLogo from '@assets/icons/synapse.svg'
import busdLogo from '@assets/icons/busd.svg'
import usdtLogo from '@assets/icons/usdt.svg'
import ethLogo from '@assets/icons/eth.svg'
import nethLogo from '@assets/icons/neth.svg'
import avwethLogo from '@assets/icons/avweth.svg'
import fraxLogo from '@assets/icons/frax.svg'
import daiLogo from '@assets/icons/dai.png'
import nusdLogo from '@assets/icons/nusd.svg'
import avaxLogo from '@assets/icons/avalanche.svg'
import movrLogo from '@assets/icons/moonriver.jpeg'
import jewelLogo from '@assets/icons/jewel.png'
import wbtcLogo from '@assets/icons/wbtc.svg'
import noteLogo from '@assets/icons/note.svg'
import ohmLogo from '@assets/icons/ohm.svg'
import highstreetLogo from '@assets/icons/highstreet.svg'
import hyperjumpLogo from '@assets/icons/hyperjump.png'
import dogLogo from '@assets/icons/dog.png'
import nfdLogo from '@assets/icons/nfd.svg'
import sdtLogo from '@assets/icons/sdt.svg'
import sfiLogo from '@assets/icons/sfi.png'
import newoLogo from '@assets/icons/newo.svg'
import vstaLogo from '@assets/icons/vsta.svg'
import usdbLogo from '@assets/icons/usdb.png'
import l2daoLogo from '@assets/icons/l2dao.svg'
import plsLogo from '@assets/icons/pls.svg'
import chainlinkLogo from '@assets/icons/chainlink.svg'
import unidexLogo from '@assets/icons/unidex.png'
import wethLogo from '@assets/icons/weth.svg'
import usdcLogo from '@assets/icons/usdc.svg'
import solarbeamLogo from '@assets/icons/solarbeam.png'
import h2oLogo from '@assets/icons/h2o.svg'
import gmxLogo from '@assets/icons/gmx.png'
import ageurLogo from '@assets/icons/ageur.svg'
import pepeLogo from '@assets/icons/pepe-token.webp'
import { AddressZero } from '@ethersproject/constants'

import { Token } from '@/utils/types'
import * as CHAINS from '@/constants/chains/master'

// MINTABLE TOKENS
export const GOHM = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0x0ab87046fBb341D058F17CBC4c1133F25a20a52f',
    [CHAINS.OPTIMISM.id]: '0x0b5740c6b4a97f90eF2F0220651Cca420B868FfB',
    [CHAINS.BNB.id]: '0x88918495892BAF4536611E38E75D771Dc6Ec0863',
    [CHAINS.POLYGON.id]: '0xd8cA34fd379d9ca3C6Ee3b3905678320F5b45195',
    [CHAINS.FANTOM.id]: '0x91fa20244Fb509e8289CA630E5db3E9166233FDc',
    [CHAINS.ARBITRUM.id]: '0x8D9bA570D6cb60C7e3e0F31343Efe75AB8E65FB1',
    [CHAINS.AVALANCHE.id]: '0x321E7092a180BB43555132ec53AaA65a5bF84251',
    [CHAINS.MOONRIVER.id]: '0x3bF21Ce864e58731B6f28D68d5928BcBEb0Ad172',
    [CHAINS.BOBA.id]: '0xd22C0a4Af486C7FA08e282E9eB5f30F9AaA62C95',
    [CHAINS.HARMONY.id]: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
    [CHAINS.MOONBEAM.id]: '0xD2666441443DAa61492FFe0F37717578714a4521',
    [CHAINS.CRONOS.id]: '0xbB0A63A6CA2071c6C4bcAC11a1A317b20E3E999C',
    [CHAINS.METIS.id]: '0xFB21B70922B9f6e3C6274BcD6CB1aa8A0fe20B80',
  },
  decimals: 18,
  symbol: 'gOHM',
  name: 'Olympus DAO',
  logo: ohmLogo,
  description: 'OHM',
  swapableType: 'OHM',
  color: 'gray',
  visibilityRank: 40,
  priorityRank: 6,
})

export const LINK = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0x514910771af9ca656af840dff83e8264ecf986ca',
    [CHAINS.KLAYTN.id]: '0xfbed1abb3ad0f8c467068de9fde905887e8c9118',
  },
  decimals: 18,
  symbol: 'LINK',
  name: 'ChainLink Token',
  logo: chainlinkLogo,
  description: 'LINK',
  swapableType: 'LINK',
  color: 'blue',
  priorityRank: 6,
})

export const HIGHSTREET = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0x71Ab77b7dbB4fa7e017BC15090b2163221420282',
    [CHAINS.BNB.id]: '0x5f4bde007dc06b867f86ebfe4802e34a1ffeed63',
  },
  decimals: 18,
  symbol: 'HIGH',
  name: 'Highstreet',
  logo: highstreetLogo,
  description: 'HIGH is the token behind Highstreet',
  swapableType: 'HIGHSTREET',
  color: 'cyan',
  priorityRank: 6,
})

export const JUMP = new Token({
  addresses: {
    [CHAINS.BNB.id]: '0x130025ee738a66e691e6a7a62381cb33c6d9ae83', // redeem
    [CHAINS.FANTOM.id]: '0x78DE9326792ce1d6eCA0c978753c6953Cdeedd73', // deposit
    [CHAINS.METIS.id]: '0xE3c82A836Ec85311a433fBd9486EfAF4b1AFbF48', // redeem
  },
  decimals: 18,
  symbol: 'JUMP',
  name: 'HyperJump',
  logo: hyperjumpLogo,
  description: 'JUMP is the token behind Hyperjump',
  docUrl: '',
  swapableType: 'JUMP',
  color: 'cyan',
  priorityRank: 6,
})

export const SFI = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0xb753428af26e81097e7fd17f40c88aaa3e04902c',
    [CHAINS.AVALANCHE.id]: '0xc2Bf0A1f7D8Da50D608bc96CF701110d4A438312', // deposit
  },
  decimals: 18,
  symbol: 'SFI',
  name: 'Saffron Finance',
  logo: sfiLogo,
  description: '',
  docUrl: '',
  swapableType: 'SFI',
  color: 'red',
  priorityRank: 6,
})

export const DOG = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0xBAac2B4491727D78D2b78815144570b9f2Fe8899',
    [CHAINS.BNB.id]: '0xaa88c603d142c371ea0eac8756123c5805edee03',
    [CHAINS.POLYGON.id]: '0xeEe3371B89FC43Ea970E908536Fcddd975135D8a',
    // [CHAINS.ARBITRUM.id]: '0x4425742F1EC8D98779690b5A3A6276Db85Ddc01A'
  },
  decimals: 18,
  symbol: 'DOG',
  name: 'The Doge NFT',
  logo: dogLogo,
  description: 'DOG is the token behind the Doge NFT',
  docUrl: '',
  swapableType: 'DOG',
  color: 'yellow',
  priorityRank: 6,
})

export const NFD = new Token({
  addresses: {
    [CHAINS.BNB.id]: '0x0fe9778c005a5a6115cbe12b0568a2d50b765a51', // redeem
    [CHAINS.AVALANCHE.id]: '0xf1293574ee43950e7a8c9f1005ff097a9a713959', // redeem
    [CHAINS.DOGE.id]: '0x868055ADFA27D331d5b69b1BF3469aDAAc3ba7f2', // redeem
    [CHAINS.POLYGON.id]: '0x0a5926027d407222f8fe20f24cb16e103f617046', // deposit
  },
  decimals: 18,
  symbol: 'NFD',
  name: 'Feisty Doge',
  logo: nfdLogo,
  description: 'Feisty Doge NFT',
  docUrl: '',
  swapableType: 'NFD',
  color: 'yellow',
  priorityRank: 6,
})

export const SOLAR = new Token({
  addresses: {
    [CHAINS.MOONBEAM.id]: '0x0DB6729C03C85B0708166cA92801BcB5CAc781fC', // redeem
    [CHAINS.MOONRIVER.id]: '0x76906411D07815491A5E577022757aD941fb5066', // deposit
  },
  decimals: 18,
  symbol: 'veSOLAR',
  name: 'Vested SolarBeam',
  logo: solarbeamLogo,
  description: 'Vested SolarBeam',
  docUrl: '',
  swapableType: 'SOLAR',
  color: 'orange',
  priorityRank: 6,
})

export const GMX = new Token({
  addresses: {
    [CHAINS.ARBITRUM.id]: '0xfc5a1a6eb076a2c7ad06ed22c90d7e710e35ad0a', // deposit
    [CHAINS.AVALANCHE.id]: '0x62edc0692bd897d2295872a9ffcac5425011c661', // redeem
  },
  wrapperAddresses: {
    [CHAINS.AVALANCHE.id]: '0x20A9DC684B4d0407EF8C9A302BEAaA18ee15F656',
  },
  decimals: 18,
  symbol: 'GMX',
  name: 'GMX',
  logo: gmxLogo,
  description: 'GMX Financial',
  docUrl: '',
  swapableType: 'GMX',
  priorityRank: 6,
})

export const SDT = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0x73968b9a57c6e53d41345fd57a6e6ae27d6cdb2f', // deposit
    [CHAINS.AVALANCHE.id]: '0xCCBf7c451F81752F7d2237F2c18C371E6e089E69', // redeem
    [CHAINS.ARBITRUM.id]: '0x087d18A77465c34CDFd3a081a2504b7E86CE4EF8',
    [CHAINS.FANTOM.id]: '0xE3c82A836Ec85311a433fBd9486EfAF4b1AFbF48', // redeem
    [CHAINS.HARMONY.id]: '0xE3c82A836Ec85311a433fBd9486EfAF4b1AFbF48', // redeem
  },
  decimals: 18,
  symbol: 'SDT',
  name: 'Stake DAO',
  logo: sdtLogo,
  description: 'Stake DAO',
  docUrl: '',
  swapableType: 'SDT',
  color: 'gray',
  priorityRank: 6,
})

export const NEWO = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0x98585dFc8d9e7D48F0b1aE47ce33332CF4237D96', // deposit
    [CHAINS.AVALANCHE.id]: '0x4Bfc90322dD638F81F034517359BD447f8E0235a', // redeem
    [CHAINS.ARBITRUM.id]: '0x0877154a755B24D499B8e2bD7ecD54d3c92BA433', // redeem
  },
  decimals: 18,
  symbol: 'NEWO',
  name: 'New Order',
  logo: newoLogo,
  description: 'New Order',
  docUrl: '',
  swapableType: 'NEWO',
  color: 'yellow',
  priorityRank: 6,
})

export const USDB = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0x02b5453d92b730f29a86a0d5ef6e930c4cf8860b',
    [CHAINS.BNB.id]: '0xc8699abbba90c7479dedccef19ef78969a2fc608',
    [CHAINS.POLYGON.id]: '0xfa1fbb8ef55a4855e5688c0ee13ac3f202486286',
    [CHAINS.FANTOM.id]: '0x6fc9383486c163fa48becdec79d6058f984f62ca',
    [CHAINS.AVALANCHE.id]: '0x5ab7084cb9d270c2cb052dd30dbecbca42f8620c',
    [CHAINS.MOONRIVER.id]: '0x3e193c39626bafb41ebe8bdd11ec7cca9b3ec0b2',
  },
  decimals: 18,
  symbol: 'USDB',
  name: 'USDB',
  logo: usdbLogo,
  description: 'USDB',
  docUrl: '',
  swapableType: 'USDB',
  priorityRank: 6,
})

export const PEPE = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0x6982508145454ce325ddbe47a25d4ec3d2311933',
    [CHAINS.ARBITRUM.id]: '0xA54B8e178A49F8e5405A4d44Bb31F496e5564A05',
  },
  decimals: 18,
  symbol: 'PEPE',
  name: 'Pepe',
  logo: pepeLogo,
  description: 'PEPE',
  swapableType: 'PEPE',
  priorityRank: 6,
})

export const VSTA = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0xA8d7F5e7C78ed0Fa097Cc5Ec66C1DC3104c9bbeb', // redeem
    [CHAINS.ARBITRUM.id]: '0xa684cd057951541187f288294a1e1c2646aa2d24', // deposit
  },
  decimals: 18,
  symbol: 'VSTA',
  name: 'Vesta',
  logo: vstaLogo,
  description: 'Vesta Finance',
  docUrl: '',
  swapableType: 'VSTA',
  color: 'gray',
  priorityRank: 6,
})

export const H2O = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0x0642026e7f0b6ccac5925b4e7fa61384250e1701', // deposit
    [CHAINS.ARBITRUM.id]: '0xD1c6f989e9552DB523aBAE2378227fBb059a3976', // redeem
    [CHAINS.AVALANCHE.id]: '0xC6b11a4Fd833d1117E9D312c02865647cd961107', // redeem
    [CHAINS.BNB.id]: '0x03eFca7CEb108734D3777684F3C0A0d8ad652f79', // redeem
    [CHAINS.MOONBEAM.id]: '0xA46aDF6D5881ca0b8596EDadF8f058F8c16d8B68', // redeem
    [CHAINS.MOONRIVER.id]: '0x9c0a820bb01e2807aCcd1c682d359e92DDd41403', // redeem
    [CHAINS.OPTIMISM.id]: '0xE3c82A836Ec85311a433fBd9486EfAF4b1AFbF48', // redeem
    [CHAINS.POLYGON.id]: '0x32ba7cF7d681357529013de6a2CDF93933C0dF3f', // redeem
  },
  decimals: 18,
  symbol: 'H2O',
  name: 'H2O',
  logo: h2oLogo,
  description: 'H2O',
  docUrl: '',
  swapableType: 'H2O',
  color: 'cyan',
  priorityRank: 6,
})

export const L2DAO = new Token({
  addresses: {
    [CHAINS.ARBITRUM.id]: '0x2CaB3abfC1670D1a452dF502e216a66883cDf079', // deposit
    [CHAINS.OPTIMISM.id]: '0xd52f94DF742a6F4B4C8b033369fE13A41782Bf44', // redeem
  },
  decimals: 18,
  symbol: 'L2DAO',
  name: 'Layer2DAO',
  logo: l2daoLogo,
  description: 'Layer2DAO',
  docUrl: '',
  swapableType: 'L2DAO',
  color: 'cyan',
  priorityRank: 6,
})

export const PLS = new Token({
  addresses: {
    [CHAINS.ARBITRUM.id]: '0x51318b7d00db7acc4026c88c3952b66278b6a67f', // deposit
    [CHAINS.OPTIMISM.id]: '0xD9eAA386cCD65F30b77FF175F6b52115FE454fD6', // redeem
  },
  decimals: 18,
  symbol: 'PLS',
  name: 'Plutus',
  logo: plsLogo,
  description: 'PlutusDao',
  docUrl: '',
  swapableType: 'PLS',
  color: 'green',
  priorityRank: 6,
})

export const AGEUR = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0x1a7e4e63778B4f12a199C062f3eFdD288afCBce8', // deposit
    [CHAINS.ARBITRUM.id]: '0x16BFc5fe024980124bEf51d1D792dC539d1B5Bf0', // redeem
    [CHAINS.OPTIMISM.id]: '0xa0554607e477cdC9d0EE2A6b087F4b2DC2815C22', // redeem
  },
  decimals: 18,
  symbol: 'agEUR',
  name: 'Angle Euro',
  logo: ageurLogo,
  description: 'Angle Euro',
  docUrl: '',
  swapableType: 'AGEUR',
  color: 'yellow',
  priorityRank: 6,
})

export const UNIDX = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0x95b3497bbcccc46a8f45f5cf54b0878b39f8d96c', // deposit
    [CHAINS.ARBITRUM.id]: '0x5429706887FCb58a595677B73E9B0441C25d993D', // redeem
  },
  decimals: 18,
  symbol: 'UNIDX',
  name: 'Unidex',
  logo: unidexLogo,
  description: 'Unidex',
  docUrl: '',
  swapableType: 'UNIDX',
  color: 'gray',
  priorityRank: 6,
})

// BASIC TOKENS
export const BUSD = new Token({
  addresses: {
    [CHAINS.BNB.id]: '0xe9e7cea3dedca5984780bafc599bd69add087d56',
    [CHAINS.DOGE.id]: '0x1555C68Be3b22cdcCa934Ae88Cb929Db40aB311d',
  },
  decimals: 18,
  symbol: 'BUSD',
  name: 'Binance USD',
  logo: busdLogo,
  description: `
    BUSD is a stablecoin that is pegged to the US dollar and
    backed/issued by Binance
  `,
  swapableType: 'BUSD',
  swapableOn: [CHAINS.BNB.id],
  color: 'yellow',
  priorityRank: 2,
})

export const USDC = new Token({
  visibilityRank: 101,
  addresses: {
    [CHAINS.BNB.id]: '0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d',
    [CHAINS.ETH.id]: '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
    [CHAINS.CRONOS.id]: '0xc21223249ca28397b4b6541dffaecc539bff0c59',
    [CHAINS.OPTIMISM.id]: '0x7f5c764cbc14f9669b88837ca1490cca17c31607',
    [CHAINS.POLYGON.id]: '0x2791bca1f2de4661ed88a30c99a7a9449aa84174',
    [CHAINS.FANTOM.id]: '0x04068da6c83afcfa0e13ba15a6696662335d5b75',
    [CHAINS.AVALANCHE.id]: '0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e',
    [CHAINS.ARBITRUM.id]: '0xaf88d065e77c8cc2239327c5edb3a432268e5831',
    [CHAINS.HARMONY.id]: '0x985458e523db3d53125813ed68c274899e9dfab4',
    [CHAINS.BOBA.id]: '0x66a2A913e447d6b4BF33EFbec43aAeF87890FBbc',
    [CHAINS.AURORA.id]: '0xB12BFcA5A55806AaF64E99521918A4bf0fC40802',
    [CHAINS.METIS.id]: '0xEA32A96608495e54156Ae48931A7c20f0dcc1a21',
    [CHAINS.CANTO.id]: '0x80b5a32E4F032B2a058b4F29EC95EEfEEB87aDcd',
    [CHAINS.KLAYTN.id]: '0x6270B58BE569a7c0b8f47594F191631Ae5b2C86C',
    [CHAINS.DOGE.id]: '0x85C2D3bEBffD83025910985389aB8aD655aBC946',
  },
  decimals: {
    [CHAINS.BNB.id]: 18,
    [CHAINS.ETH.id]: 6,
    [CHAINS.OPTIMISM.id]: 6,
    [CHAINS.POLYGON.id]: 6,
    [CHAINS.FANTOM.id]: 6,
    [CHAINS.AVALANCHE.id]: 6,
    [CHAINS.ARBITRUM.id]: 6,
    [CHAINS.HARMONY.id]: 6,
    [CHAINS.BOBA.id]: 6,
    [CHAINS.AURORA.id]: 6,
    [CHAINS.METIS.id]: 6,
    [CHAINS.CRONOS.id]: 6,
    [CHAINS.KLAYTN.id]: 6,
    [CHAINS.CANTO.id]: 6,
  },
  swapExceptions: {
    [CHAINS.KLAYTN.id]: [CHAINS.ETH.id, CHAINS.DOGE.id],
    [CHAINS.DOGE.id]: [CHAINS.ETH.id, CHAINS.DOGE.id],
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
  swapableOn: [
    CHAINS.BNB.id,
    CHAINS.ETH.id,
    CHAINS.POLYGON.id,
    CHAINS.FANTOM.id,
    // CHAINS.ARBITRUM.id,
    // CHAINS.AVALANCHE.id,
    CHAINS.HARMONY.id,
    CHAINS.AURORA.id,
    CHAINS.BOBA.id,
    CHAINS.OPTIMISM.id,
    CHAINS.METIS.id,
    CHAINS.CRONOS.id,
    CHAINS.CANTO.id,
  ],
  color: 'blue',
  priorityRank: 1,
})

// export const KLAYTN_USDC = new Token({
//   addresses: {
//     [CHAINS.ETH.id]: '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
//     [CHAINS.KLAYTN.id]: '0x6270B58BE569a7c0b8f47594F191631Ae5b2C86C',
//     [CHAINS.DOGE.id]: '0x85C2D3bEBffD83025910985389aB8aD655aBC946',
//   },
//   decimals: {
//     [CHAINS.ETH.id]: 6,
//     [CHAINS.KLAYTN.id]: 6,
//     [CHAINS.DOGE.id]: 6,
//   },
//   symbol: 'USDC  ', // TWO SPACES IS EXTREMELY IMPORTANT
//   name: 'USD Circle',
//   logo: usdcLogo,
//   description: `
//     USD Coin (known by its ticker USDC) is a stablecoin that is pegged to the
//     U.S. dollar on a 1:1 basis. Every unit of this cryptocurrency in circulation
//     is backed up by $1 that is held in reserve
//     `,
//   swapableType: 'KLAYTN_USDC',
// })

// export const KLAYTN_USDT = new Token({
//   addresses: {
//     [CHAINS.ETH.id]: '0xdac17f958d2ee523a2206206994597c13d831ec7',
//     [CHAINS.KLAYTN.id]: '0xd6dAb4CfF47dF175349e6e7eE2BF7c40Bb8C05A3',
//     [CHAINS.DOGE.id]: '0x7f8e71DD5A7e445725F0EF94c7F01806299e877A',
//   },
//   decimals: {
//     [CHAINS.ETH.id]: 6,
//     [CHAINS.KLAYTN.id]: 6,
//     [CHAINS.DOGE.id]: 6,
//   },
//   symbol: 'USDT  ', // TWO SPACES IS EXTREMELY IMPORTANT
//   name: 'Synapse Tether USDT',
//   logo: usdtLogo,
//   swapableType: 'KLAYTN_USDT',
// })

export const KLAYTN_oUSDT = new Token({
  addresses: {
    [CHAINS.KLAYTN.id]: '0xceE8FAF64bB97a73bb51E115Aa89C17FfA8dD167',
  },
  decimals: {
    [CHAINS.KLAYTN.id]: 6,
  },
  symbol: 'orbitUSDT', // TWO SPACES IS EXTREMELY IMPORTANT
  name: 'Orbit Bridged USDT',
  logo: usdtLogo,
  swapableType: 'KLAYTN_USDT',
  swapableOn: [CHAINS.KLAYTN.id],
  priorityRank: 6,
})

// export const KLAYTN_DAI = new Token({
//   addresses: {
//     [CHAINS.ETH.id]: '0x6b175474e89094c44da98b954eedeac495271d0f',
//     [CHAINS.KLAYTN.id]: '0x078dB7827a5531359f6CB63f62CFA20183c4F10c',
//     [CHAINS.DOGE.id]: '0xB3306f03595490e5cC3a1b1704a5a158D3436ffC',
//   },
//   decimals: {
//     [CHAINS.ETH.id]: 18,
//     [CHAINS.KLAYTN.id]: 18,
//     [CHAINS.DOGE.id]: 18,
//   },
//   symbol: 'DAI  ', // TWO SPACES IS EXTREMELY IMPORTANT
//   name: 'DAI',
//   logo: daiLogo,
//   swapableType: 'KLAYTN_DAI',
// })

// export const DOGECHAIN_BUSD = new Token({
//   addresses: {
//     [CHAINS.BNB.id]: '0xe9e7cea3dedca5984780bafc599bd69add087d56',
//     [CHAINS.DOGE.id]: '0x1555C68Be3b22cdcCa934Ae88Cb929Db40aB311d',
//   },
//   decimals: {
//     [CHAINS.BNB.id]: 18,
//     [CHAINS.DOGE.id]: 18,
//   },
//   symbol: 'BUSD ', // ONE SPACE IS EXTREMELY IMPORTANT
//   name: 'Binance USD',
//   logo: busdLogo,
//   swapableType: 'DOGECHAIN_BUSD',
// })

export const USDT = new Token({
  addresses: {
    [CHAINS.BNB.id]: '0x55d398326f99059ff775485246999027b3197955',
    [CHAINS.ETH.id]: '0xdac17f958d2ee523a2206206994597c13d831ec7',
    [CHAINS.CRONOS.id]: '0x66e428c3f67a68878562e79a0234c1f83c208770',
    [CHAINS.POLYGON.id]: '0xc2132d05d31c914a87c6611c10748aeb04b58e8f',
    // [CHAINS.AVALANCHE.id]: '0x9702230a8ea53601f5cd2dc00fdbc13d4df4a8c7',
    // [CHAINS.HARDHAT.id]: '0x9A9f2CCfdE556A7E9Ff0848998Aa4a0CFD8863AE',
    [CHAINS.ARBITRUM.id]: '0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9',
    [CHAINS.FANTOM.id]: '0x049d68029688eabf473097a2fc38ef61633a3c7a',
    [CHAINS.HARMONY.id]: '0x3c2b8be99c50593081eaa2a724f0b8285f5aba8f',
    [CHAINS.BOBA.id]: '0x5DE1677344D3Cb0D7D465c10b72A8f60699C062d',
    [CHAINS.AURORA.id]: '0x4988a896b1227218e4A686fdE5EabdcAbd91571f',
    [CHAINS.CANTO.id]: '0xd567B3d7B8FE3C79a1AD8dA978812cfC4Fa05e75',
    [CHAINS.KLAYTN.id]: '0xd6dAb4CfF47dF175349e6e7eE2BF7c40Bb8C05A3',
    [CHAINS.DOGE.id]: '0x7f8e71DD5A7e445725F0EF94c7F01806299e877A',
  },
  swapExceptions: {
    [CHAINS.KLAYTN.id]: [CHAINS.ETH.id, CHAINS.DOGE.id],
    [CHAINS.DOGE.id]: [CHAINS.ETH.id, CHAINS.DOGE.id],
  },
  decimals: {
    [CHAINS.BNB.id]: 18,
    [CHAINS.ETH.id]: 6,
    [CHAINS.CRONOS.id]: 6,
    [CHAINS.POLYGON.id]: 6,
    // Commenting out as currently unsupported above
    // [CHAINS.AVALANCHE.id]: 6,
    [CHAINS.ARBITRUM.id]: 6,
    [CHAINS.FANTOM.id]: 6,
    [CHAINS.HARMONY.id]: 6,
    [CHAINS.BOBA.id]: 6,
    [CHAINS.AURORA.id]: 6,
    [CHAINS.CANTO.id]: 6,
    [CHAINS.KLAYTN.id]: 6,
    [CHAINS.DOGE.id]: 6,
  },
  symbol: 'USDT',
  name: 'USD Tether',
  logo: usdtLogo,
  color: 'lime',
  description: `
    USDT mirrors the price of the U.S. dollar, issued by a Hong Kong-based company Tether.
    The token’s peg to the USD is achieved via maintaining a sum of dollars in reserves equal
    to the number of USDT in circulation.
    `,
  swapableType: 'USD',
  swapableOn: [
    CHAINS.BNB.id,
    CHAINS.ETH.id,
    CHAINS.POLYGON.id,
    CHAINS.FANTOM.id,
    CHAINS.ARBITRUM.id,
    // CHAINS.AVALANCHE.id,
    CHAINS.HARMONY.id,
    CHAINS.AURORA.id,
    CHAINS.BOBA.id,
    CHAINS.CANTO.id,
  ],
  visibilityRank: 100,
  priorityRank: 1,
})

export const DAI = new Token({
  addresses: {
    //[CHAINS.BNB.id]: '0x1af3f329e8be154074d8769d1ffa4ee058b1dbc3',
    [CHAINS.ETH.id]: '0x6b175474e89094c44da98b954eedeac495271d0f',
    [CHAINS.CRONOS.id]: '0xf2001b145b43032aaf5ee2884e456ccd805f677d',
    [CHAINS.POLYGON.id]: '0x8f3cf7ad23cd3cadbd9735aff958023239c6a063',
    [CHAINS.AVALANCHE.id]: '0xd586E7F844cEa2F87f50152665BCbc2C279D8d70',
    // [CHAINS.ARBITRUM.id]: '0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1',
    [CHAINS.HARMONY.id]: '0xef977d2f931c1978db5f6747666fa1eacb0d0339',
    [CHAINS.BOBA.id]: '0xf74195Bb8a5cf652411867c5C2C5b8C2a402be35',
    [CHAINS.KLAYTN.id]: '0x078dB7827a5531359f6CB63f62CFA20183c4F10c',
    [CHAINS.DOGE.id]: '0xB3306f03595490e5cC3a1b1704a5a158D3436ffC',
  },
  swapExceptions: {
    [CHAINS.KLAYTN.id]: [CHAINS.ETH.id, CHAINS.DOGE.id],
    [CHAINS.DOGE.id]: [CHAINS.ETH.id, CHAINS.DOGE.id],
  },
  decimals: 18,
  symbol: 'DAI',
  name: 'Dai',
  logo: daiLogo,
  swapableType: 'USD',
  swapableOn: [
    CHAINS.ETH.id,
    CHAINS.POLYGON.id,
    CHAINS.AVALANCHE.id,
    CHAINS.HARMONY.id,
    CHAINS.BOBA.id,
  ],
  color: 'orange',
  visibilityRank: 100,
  priorityRank: 1,
})

export const WBTC = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
    [CHAINS.KLAYTN.id]: '0xDCbacF3f7a069922E677912998c8d57423C37dfA',
    [CHAINS.DOGE.id]: '0xD0c6179c43C00221915f1a61f8eC06A5Aa32F9EC',
  },
  decimals: {
    [CHAINS.ETH.id]: 8,
    [CHAINS.KLAYTN.id]: 8,
    [CHAINS.DOGE.id]: 8,
  },
  symbol: 'WBTC',
  name: 'Wrapped BTC',
  logo: wbtcLogo,
  swapableType: 'WBTC',
  color: 'yellow',
  priorityRank: 3,
})

// export const MIM = new Token({
//   addresses: {
//     [CHAINS.FANTOM.id]: '0x82f0b8b456c1a451378467398982d4834b6829c1',
//     [CHAINS.ARBITRUM.id]: '0xfea7a6a0b346362bf88a9e4a88416b77a57d6c2a',
//   },
//   decimals: 18,
//   symbol: 'MIM',
//   name: 'Magic Internet Money',
//   logo: mimLogo,
//   swapableType: 'USD',
// })

export const WETHE = new Token({
  addresses: {
    [CHAINS.AVALANCHE.id]: '0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab',
  },
  decimals: 18,
  symbol: 'WETH.e', // SHOULD BE WETH
  name: 'Wrapped ETH',
  logo: wethLogo,
  description: 'ERC-20 Wrapped form of ETH',
  swapableType: 'ETH',
  swapableOn: [CHAINS.AVALANCHE.id],
  color: 'sky',
  priorityRank: 2,
})

/**
 * WETH on Moonbeam is nETH on moonbeam.
 * is this stupid & annoying - yes
 */
export const WETHBEAM = new Token({
  addresses: {
    [CHAINS.MOONBEAM.id]: '0x3192Ae73315c3634Ffa217f71CF6CBc30FeE349A',
  },
  decimals: 18,
  symbol: 'WETH ',
  name: 'Wrapped ETH',
  logo: wethLogo,
  description: 'ERC-20 Wrapped form of ETH on Moonbeam',
  swapableType: 'ETH',
  color: 'sky',
  priorityRank: 2,
})

export const AVWETH = new Token({
  addresses: {
    [CHAINS.AVALANCHE.id]: '0x53f7c5869a859f0aec3d334ee8b4cf01e3492f21',
  },
  decimals: 18,
  symbol: 'AVWETH', // AVALANCHE AAVE WETH
  name: 'Aave Wrapped ETH',
  logo: avwethLogo,
  description: 'Aave Wrapped form of ETH',
  swapableType: 'ETH',
  color: 'cyan',
  priorityRank: 2,
})

export const ONEETH = new Token({
  addresses: {
    [CHAINS.HARMONY.id]: '0x6983d1e6def3690c4d616b13597a09e6193ea013',
  },
  decimals: 18,
  symbol: '1ETH', // SHOULD BE WETH
  name: 'Harmony ETH',
  logo: wethLogo,
  description: 'Harmony ERC-20 Wrapped form of ETH',
  swapableType: 'ETH',
  swapableOn: [CHAINS.HARMONY.id],
  color: 'sky',
  priorityRank: 2,
})

export const FTMETH = new Token({
  addresses: {
    [CHAINS.FANTOM.id]: '0x74b23882a30290451A17c44f4F05243b6b58C76d',
  },
  decimals: 18,
  symbol: 'ETH ',
  name: 'Wrapped ETH',
  logo: wethLogo,
  description: 'Fantom Wrapped form of ETH',
  swapableType: 'ETH',
  swapableOn: [CHAINS.FANTOM.id],
  color: 'sky',
  priorityRank: 2,
})

export const CANTOETH = new Token({
  addresses: {
    [CHAINS.CANTO.id]: '0x5FD55A1B9FC24967C4dB09C513C3BA0DFa7FF687',
  },
  decimals: 18,
  symbol: 'ETH ',
  logo: wethLogo,
  name: 'Wrapped ETH',
  description: 'Canto Wrapped form of ETH',
  swapableType: 'ETH',
  swapableOn: [CHAINS.CANTO.id],
  color: 'sky',
  visibilityRank: 100,
  priorityRank: 2,
})

export const METISETH = new Token({
  addresses: {
    [CHAINS.METIS.id]: '0x420000000000000000000000000000000000000A',
  },
  decimals: 18,
  symbol: 'ETH ',
  name: 'Wrapped ETH',
  logo: wethLogo,
  description: 'Metis Wrapped form of ETH',
  swapableType: 'ETH',
  swapableOn: [CHAINS.METIS.id],
  color: 'sky',
  priorityRank: 2,
})

export const SYN = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0x0f2d719407fdbeff09d87557abb7232601fd9f29',
    [CHAINS.BNB.id]: '0xa4080f1778e69467e905b8d6f72f6e441f9e9484',
    [CHAINS.POLYGON.id]: '0xf8f9efc0db77d8881500bb06ff5d6abc3070e695',
    [CHAINS.FANTOM.id]: '0xE55e19Fb4F2D85af758950957714292DAC1e25B2', // yes this is same as avax swap addr, no its not error
    [CHAINS.ARBITRUM.id]: '0x080f6aed32fc474dd5717105dba5ea57268f46eb',
    [CHAINS.AVALANCHE.id]: '0x1f1E7c893855525b303f99bDF5c3c05Be09ca251',
    [CHAINS.HARMONY.id]: '0xE55e19Fb4F2D85af758950957714292DAC1e25B2',
    [CHAINS.BOBA.id]: '0xb554A55358fF0382Fb21F0a478C3546d1106Be8c',
    [CHAINS.METIS.id]: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
    [CHAINS.MOONRIVER.id]: '0xd80d8688b02B3FD3afb81cDb124F188BB5aD0445',
    [CHAINS.MOONBEAM.id]: '0xF44938b0125A6662f9536281aD2CD6c499F22004',
    [CHAINS.OPTIMISM.id]: '0x5A5fFf6F753d7C11A56A52FE47a177a87e431655',
    [CHAINS.CRONOS.id]: '0xFD0F80899983b8D46152aa1717D76cba71a31616',
    [CHAINS.AURORA.id]: '0xd80d8688b02B3FD3afb81cDb124F188BB5aD0445',
    [CHAINS.DOGE.id]: '0xDfA53EeBA61D69E1D2b56b36d78449368F0265c1',
    [CHAINS.CANTO.id]: '0x555982d2E211745b96736665e19D9308B615F78e',
  },
  decimals: 18,
  symbol: 'SYN',
  name: 'Synapse',
  logo: synapseLogo,
  description: 'SYN is the base token behind synapse',
  swapableType: 'SYN',
  color: 'purple',
  visibilityRank: 90,
  priorityRank: 2,
})

export const FRAX = new Token({
  addresses: {
    [CHAINS.ETH.id]: '0x853d955acef822db058eb8505911ed77f175b99e',
    // [CHAINS.MOONRIVER.id]: '0x1a93b23281cc1cde4c4741353f3064709a16197d',
    // [CHAINS.MOONRIVER.id]: '0xE96AC70907ffF3Efee79f502C985A7A21Bce407d',

    // [CHAINS.MOONBEAM.id]: '',
    // [CHAINS.HARMONY.id]: '0xFa7191D292d5633f702B0bd7E3E3BcCC0e633200',
    [CHAINS.DOGE.id]: '0x10D70831f9C3c11c5fe683b2f1Be334503880DB6',
  },
  decimals: 18,
  symbol: 'FRAX',
  name: 'Frax',
  logo: fraxLogo,
  description: 'Frax',
  swapableType: 'FRAX',
  // swapableOn: [CHAINS.MOONRIVER.id, CHAINS.MOONBEAM.id],
  color: 'gray',
  priorityRank: 6,
})

export const SYN_FRAX = new Token({
  addresses: {
    // [CHAINS.FANTOM.id]:    '0x1852F70512298d56e9c8FDd905e02581E04ddb2a',
    [CHAINS.MOONRIVER.id]: '0xE96AC70907ffF3Efee79f502C985A7A21Bce407d',
    // [CHAINS.MOONBEAM.id]: '0xDd47A348AB60c61Ad6B60cA8C31ea5e00eBfAB4F',
    [CHAINS.HARMONY.id]: '0x1852F70512298d56e9c8FDd905e02581E04ddb2a',
  },
  decimals: 18,
  symbol: 'synFRAX',
  name: 'Synapse Frax',
  logo: synapseLogo,
  description: 'Frax',
  swapableType: 'FRAX',
  priorityRank: 4,
})

/**
 * nUSD is the token involved in the bridge.
 */
export const NUSD = new Token({
  addresses: {
    [CHAINS.BNB.id]: '0x23b891e5c62e0955ae2bd185990103928ab817b3',
    [CHAINS.ETH.id]: '0x1B84765dE8B7566e4cEAF4D0fD3c5aF52D3DdE4F',
    [CHAINS.CRONOS.id]: '0x396c9c192dd323995346632581BEF92a31AC623b',
    [CHAINS.OPTIMISM.id]: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
    [CHAINS.POLYGON.id]: '0xb6c473756050de474286bed418b77aeac39b02af',
    [CHAINS.FANTOM.id]: '0xED2a7edd7413021d440b09D654f3b87712abAB66',
    [CHAINS.AVALANCHE.id]: '0xCFc37A6AB183dd4aED08C204D1c2773c0b1BDf46',
    [CHAINS.ARBITRUM.id]: '0x2913E812Cf0dcCA30FB28E6Cac3d2DCFF4497688',
    [CHAINS.HARMONY.id]: '0xED2a7edd7413021d440b09D654f3b87712abAB66',
    [CHAINS.BOBA.id]: '0x6B4712AE9797C199edd44F897cA09BC57628a1CF',
    [CHAINS.AURORA.id]: '0x07379565cD8B0CaE7c60Dc78e7f601b34AF2A21c',
    [CHAINS.METIS.id]: '0x961318Fc85475E125B99Cc9215f62679aE5200aB',
    [CHAINS.DFK.id]: '0x3AD9DFE640E1A9Cc1D9B0948620820D975c3803a',
    [CHAINS.CANTO.id]: '0xD8836aF2e565D3Befce7D906Af63ee45a57E8f80',
  },
  decimals: 18,
  symbol: 'nUSD',
  name: 'Synapse nUSD',
  logo: nusdLogo,
  description: 'nUSD',
  swapableType: 'USD',
  swapableOn: [
    CHAINS.BNB.id,
    CHAINS.POLYGON.id,
    CHAINS.FANTOM.id,
    CHAINS.ARBITRUM.id,
    CHAINS.AVALANCHE.id,
    CHAINS.HARMONY.id,
    CHAINS.AURORA.id,
    CHAINS.BOBA.id,
    CHAINS.OPTIMISM.id,
    CHAINS.METIS.id,
    CHAINS.CRONOS.id,
    CHAINS.CANTO.id,
  ],
  color: 'purple',
  visibilityRank: 50,
  priorityRank: 5,
})

export const NOTE = new Token({
  addresses: {
    [CHAINS.CANTO.id]: '0x4e71a2e537b7f9d9413d3991d37958c0b5e1e503',
  },
  decimals: 18,
  symbol: 'NOTE',
  name: 'Canto Note',
  logo: noteLogo,
  description: 'NOTE',
  swapableType: 'USD',
  swapableOn: [CHAINS.CANTO.id],
  color: 'green',
  visibilityRank: 90,
  priorityRank: 3,
})

export const DFK_USDC = new Token({
  addresses: {
    [CHAINS.DFK.id]: NUSD.addresses[CHAINS.DFK.id],
  },
  decimals: {
    [CHAINS.DFK.id]: 18,
  },
  symbol: 'USDC ', // SPACE VERY IMPORTANT
  name: 'USD Circle',
  logo: usdcLogo,
  description: '',
  swapableType: 'USD',
  color: 'blue',
  visibilityRank: 100,
  priorityRank: 1,
})

/**
 * nETH is the token involved in the bridge. it is backed by internet monies...
 */
export const NETH = new Token({
  addresses: {
    [CHAINS.FANTOM.id]: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
    [CHAINS.ARBITRUM.id]: '0x3ea9B0ab55F34Fb188824Ee288CeaEfC63cf908e',
    [CHAINS.BOBA.id]: '0x96419929d7949D6A801A6909c145C8EEf6A40431',
    [CHAINS.OPTIMISM.id]: '0x809DC529f07651bD43A172e8dB6f4a7a0d771036',
    [CHAINS.AVALANCHE.id]: '0x19E1ae0eE35c0404f835521146206595d37981ae',
    [CHAINS.HARMONY.id]: '0x0b5740c6b4a97f90eF2F0220651Cca420B868FfB',
    [CHAINS.MOONBEAM.id]: '0x3192Ae73315c3634Ffa217f71CF6CBc30FeE349A', // THIS OVERLAPS WITH WETHBEAM
    [CHAINS.METIS.id]: '0x931B8f17764362A3325D30681009f0eDd6211231',
    [CHAINS.KLAYTN.id]: '0xCD6f29dC9Ca217d0973d3D21bF58eDd3CA871a86',
    [CHAINS.DOGE.id]: '0x9F4614E4Ea4A0D7c4B1F946057eC030beE416cbB',
    [CHAINS.CANTO.id]: '0x09fEC30669d63A13c666d2129230dD5588E2e240',
  },
  decimals: 18,
  symbol: 'nETH',
  name: 'Synapse nETH',
  logo: nethLogo,
  description: 'nETH',
  swapableType: 'ETH',
  swapableOn: [
    CHAINS.ARBITRUM.id,
    CHAINS.AVALANCHE.id,
    CHAINS.HARMONY.id,
    CHAINS.BOBA.id,
    CHAINS.OPTIMISM.id,
    CHAINS.METIS.id,
    CHAINS.CANTO.id,
  ],
  color: 'purple',
  visibilityRank: 50,
  priorityRank: 5,
})

export const KLAYTN_WETH = new Token({
  addresses: {
    [CHAINS.KLAYTN.id]: NETH.addresses[CHAINS.KLAYTN.id],
    [CHAINS.DOGE.id]: NETH.addresses[CHAINS.DOGE.id],
  },
  decimals: {
    [CHAINS.KLAYTN.id]: 18,
    [CHAINS.DOGE.id]: 18,
  },
  symbol: 'WETH ', // SPACE VERY IMPORTANT
  name: 'Wrapped ETH',
  logo: ethLogo,
  description: '',
  swapableType: 'ETH',
  priorityRank: 3,
})

export const ETH = new Token({
  addresses: {
    [CHAINS.ETH.id]: AddressZero,
    [CHAINS.BOBA.id]: AddressZero,
    [CHAINS.ARBITRUM.id]: AddressZero,
    [CHAINS.OPTIMISM.id]: AddressZero,
  },
  decimals: 18,
  symbol: 'ETH',
  name: 'Ethereum',
  logo: ethLogo,
  description: 'ETH',
  isNative: true,
  swapableType: 'ETH',
  color: 'sky',
  visibilityRank: 101,
  priorityRank: 2,
})

export const MOVR = new Token({
  addresses: {
    [CHAINS.MOONRIVER.id]: AddressZero,
  },
  decimals: 18,
  symbol: 'MOVR',
  name: 'MOVR',
  logo: movrLogo,
  description: 'Moonriver',
  isNative: true,
  swapableType: 'MOVR',
  color: 'purple',
  priorityRank: 3,
})

export const AVAX = new Token({
  addresses: {
    [CHAINS.AVALANCHE.id]: AddressZero,
  },
  decimals: 18,
  symbol: 'AVAX',
  name: 'AVAX',
  logo: avaxLogo,
  description: 'AVAX',
  isNative: true,
  swapableType: 'AVAX',
  color: 'red',
  priorityRank: 3,
})

export const WMOVR = new Token({
  addresses: {
    [CHAINS.MOONBEAM.id]: '0x1d4C2a246311bB9f827F4C768e277FF5787B7D7E',
    [CHAINS.MOONRIVER.id]: '0x98878b06940ae243284ca214f92bb71a2b032b8a',
  },
  decimals: 18,
  symbol: 'MOVR', // SHOULD BE WETH
  name: 'Wrapped MOVR',
  logo: movrLogo,
  description: 'ERC-20 Wrapped form of MOVR',
  swapableType: 'MOVR',
  color: 'purple',
  priorityRank: 3,
})

export const WAVAX = new Token({
  addresses: {
    [CHAINS.AVALANCHE.id]: '0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7',
    [CHAINS.DFK.id]: '0xB57B60DeBDB0b8172bb6316a9164bd3C695F133a',
    [CHAINS.MOONBEAM.id]: '0xA1f8890E39b4d8E33efe296D698fe42Fb5e59cC3',
    [CHAINS.HARMONY.id]: '0xD9eAA386cCD65F30b77FF175F6b52115FE454fD6',
  },
  decimals: 18,
  symbol: 'AVAX',
  name: 'Wrapped AVAX',
  logo: avaxLogo,
  description: 'ERC-20 Wrapped form of AVAX',
  swapableType: 'AVAX',
  color: 'red',
  visibilityRank: 90,
  priorityRank: 3,
})

export const SYNAVAX = new Token({
  addresses: {
    [CHAINS.HARMONY.id]: '0xD9eAA386cCD65F30b77FF175F6b52115FE454fD6',
  },
  decimals: 18,
  symbol: 'synAVAX',
  name: 'Wrapped AVAX',
  logo: avaxLogo,
  description: 'ERC-20 Wrapped form of AVAX',
  swapableType: 'AVAX',
  swapableOn: [CHAINS.HARMONY.id],
  color: 'red',
  priorityRank: 4,
})

export const MULTIAVAX = new Token({
  addresses: {
    [CHAINS.HARMONY.id]: '0xb12c13e66ade1f72f71834f2fc5082db8c091358',
  },
  decimals: 18,
  symbol: 'multiAVAX', // SHOULD BE WETH
  name: 'AnySwap Wrapped AVAX',
  logo: avaxLogo,
  description: 'ERC-20 Wrapped form of AVAX',
  swapableType: 'AVAX',
  swapableOn: [CHAINS.HARMONY.id],
  color: 'red',
  priorityRank: 3,
})

export const JEWEL = new Token({
  addresses: {
    [CHAINS.DFK.id]: AddressZero,
    [CHAINS.HARMONY.id]: '0x72cb10c6bfa5624dd07ef608027e366bd690048f', // from harmony jewel?
    [CHAINS.KLAYTN.id]: '0x30C103f8f5A3A732DFe2dCE1Cc9446f545527b43',
  },
  decimals: 18,
  symbol: 'JEWEL',
  name: 'JEWEL',
  logo: jewelLogo,
  description: 'JEWEL',
  isNative: true,
  swapableType: 'JEWEL',
  priorityRank: 3,
})

export const WJEWEL = new Token({
  addresses: {
    [CHAINS.DFK.id]: '0xCCb93dABD71c8Dad03Fc4CE5559dC3D89F67a260', // from actual jewl
  },
  decimals: 18,
  symbol: 'WJEWEL',
  name: 'Wrapped JEWEL',
  logo: jewelLogo,
  description: 'JEWEL',
  swapableType: 'JEWEL',
  swapableOn: [CHAINS.HARMONY.id],
  color: 'lime',
  priorityRank: 3,
})

export const SYNJEWEL = new Token({
  addresses: {
    [CHAINS.AVALANCHE.id]: '0x997Ddaa07d716995DE90577C123Db411584E5E46',
    [CHAINS.HARMONY.id]: '0x28b42698Caf46B4B012CF38b6C75867E0762186D',
  },
  decimals: 18,
  symbol: 'JEWEL  ', // THE SPACES ARE VERY IMPORTANT
  name: 'JEWEL  ', // THE SPACES ARE VERY IMPORTANT
  logo: jewelLogo,
  description: 'ERC-20 Wrapped form of JEWEL',
  swapableType: 'JEWEL',
  swapableOn: [CHAINS.HARMONY.id],
  color: 'lime',
  priorityRank: 4,
})

export const XJEWEL = new Token({
  addresses: {
    [CHAINS.DFK.id]: '0x77f2656d04E158f915bC22f07B779D94c1DC47Ff',
    [CHAINS.HARMONY.id]: '0xA9cE83507D872C5e1273E745aBcfDa849DAA654F',
  },
  decimals: 18,
  symbol: 'xJEWEL',
  name: 'xJEWEL',
  logo: jewelLogo,
  description: 'ERC-20 Wrapped form of xJEWEL',
  swapableType: 'XJEWEL',
  color: 'lime',
  priorityRank: 3,
})

export const USDCe = new Token({
  addresses: {
    [CHAINS.AVALANCHE.id]: '0xA7D7079b0FEaD91F3e65f86E8915Cb59c1a4C664',
    [CHAINS.ARBITRUM.id]: '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8',
  },
  decimals: 6,
  symbol: 'USDCe',
  name: 'Bridged USDC',
  logo: usdcLogo,
  description: `
    USD Coin (known by its ticker USDC) is a stablecoin that is pegged to the
    U.S. dollar on a 1:1 basis. Every unit of this cryptocurrency in circulation
    is backed up by $1 that is held in reserve
  `,
  swapableType: 'USD',
  swapableOn: [CHAINS.AVALANCHE.id, CHAINS.ARBITRUM.id],
  visibilityRank: 100,
  priorityRank: 1,
})

export const USDTe = new Token({
  addresses: {
    [CHAINS.AVALANCHE.id]: '0xc7198437980c041c805a1edcba50c1ce5db95118',
  },
  decimals: 6,
  symbol: 'USDTe',
  name: 'Tether',
  logo: usdtLogo,
  description: `
    USDT mirrors the price of the U.S. dollar, issued by a Hong Kong-based company Tether.
    The token’s peg to the USD is achieved via maintaining a sum of dollars in reserves equal
    to the number of USDT in circulation.
  `,
  swapableType: 'USD',
  swapableOn: [CHAINS.AVALANCHE.id],
  visibilityRank: 100,
  priorityRank: 1,
})

// export const WMATIC = new Token({
//   addresses: {
//     [CHAINS.POLYGON.id]: '0x9b17bAADf0f21F03e35249e0e59723F34994F806',
//   },
//   decimals: 18,
//   symbol: 'MATIC', // SHOULD BE WETH
//   name: 'Wrapped MATIC',
//   description: 'ERC-20 Wrapped form of MATIC',
//   swapableType: 'MATIC',
// })

// export const WBNB = new Token({
//   addresses: {
//     [CHAINS.BNB.id]: '0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c',
//     [CHAINS.DOGE.id]: '0x1fC532187B4848d2F9c564531b776A4F8e11201d',
//   },
//   decimals: 18,
//   symbol: 'BNB', // SHOULD BE WETH
//   name: 'Wrapped BNB',
//   description: 'ERC-20 Wrapped form of BNB',
//   swapableType: 'BNB',
// })

// export const DEPRECATED_WKLAY = new Token({
//   addresses: {
//     [CHAINS.KLAYTN.id]: '0x5819b6af194a78511c79c85ea68d2377a7e9335f',
//   },
//   decimals: 18,
//   symbol: 'WKLAY',
//   name: 'Deprecated Wrapped Klay',
//   description: 'ERC-20 Wrapped form of KLAY',
//   swapableType: 'KLAY',
// })
