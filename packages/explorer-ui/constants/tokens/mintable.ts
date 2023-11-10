import { ChainId } from '@constants/networks'
import ohmLogo from '@assets/icons/ohm.svg'
import highstreetLogo from '@assets/icons/highstreet.svg'
import hyperjumpLogo from '@assets/icons/hyperjump.png'
import dogLogo from '@assets/icons/dog.png'
import nfdLogo from '@assets/icons/nfd.svg'
import sdtLogo from '@assets/icons/sdt.svg'
import newoLogo from '@assets/icons/newo.svg'
import vstaLogo from '@assets/icons/vsta.svg'
import usdbLogo from '@assets/icons/usdb.png'
import solarbeamLogo from '@assets/icons/solarbeam.png'
import gmxLogo from '@assets/icons/gmx.png'
import sfiLogo from '@assets/icons/sfi.png'
import h2oLogo from '@assets/icons/h2o.svg'
import l2daoLogo from '@assets/icons/l2dao.svg'
import plsLogo from '@assets/icons/pls.svg'
import ageurLogo from '@assets/icons/ageur.svg'
import chainlinkLogo from '@assets/icons/chainlink.svg'
import unidexLogo from '@assets/icons/unidex.png'
import { Token } from '@utils/classes/Token'

export const GOHM = new Token({
  addresses: {
    [ChainId.ETH]: '0x0ab87046fBb341D058F17CBC4c1133F25a20a52f',
    [ChainId.OPTIMISM]: '0x0b5740c6b4a97f90eF2F0220651Cca420B868FfB',
    [ChainId.BSC]: '0x88918495892BAF4536611E38E75D771Dc6Ec0863',
    [ChainId.POLYGON]: '0xd8cA34fd379d9ca3C6Ee3b3905678320F5b45195',
    [ChainId.FANTOM]: '0x91fa20244Fb509e8289CA630E5db3E9166233FDc',
    [ChainId.ARBITRUM]: '0x8D9bA570D6cb60C7e3e0F31343Efe75AB8E65FB1',
    [ChainId.AVALANCHE]: '0x321E7092a180BB43555132ec53AaA65a5bF84251',
    [ChainId.MOONRIVER]: '0x3bF21Ce864e58731B6f28D68d5928BcBEb0Ad172',
    [ChainId.BOBA]: '0xd22C0a4Af486C7FA08e282E9eB5f30F9AaA62C95',
    [ChainId.HARMONY]: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
    [ChainId.MOONBEAM]: '0xD2666441443DAa61492FFe0F37717578714a4521',
    [ChainId.CRONOS]: '0xbB0A63A6CA2071c6C4bcAC11a1A317b20E3E999C',
    [ChainId.METIS]: '0xFB21B70922B9f6e3C6274BcD6CB1aa8A0fe20B80',
  },
  decimals: 18,
  symbol: 'gOHM',
  name: 'Olympus DAO',
  logo: ohmLogo,
  description: 'OHM',
  swapableType: 'OHM',
})

export const LINK = new Token({
  addresses: {
    [ChainId.ETH]: '0x514910771af9ca656af840dff83e8264ecf986ca',
    [ChainId.KLAYTN]: '0xfbed1abb3ad0f8c467068de9fde905887e8c9118',
  },
  decimals: 18,
  symbol: 'LINK',
  name: 'ChainLink Token',
  logo: chainlinkLogo,
  description: 'LINK',
  swapableType: 'LINK',
})

export const HIGHSTREET = new Token({
  addresses: {
    [ChainId.ETH]: '0x71Ab77b7dbB4fa7e017BC15090b2163221420282',
    [ChainId.BSC]: '0x5f4bde007dc06b867f86ebfe4802e34a1ffeed63',
  },
  decimals: 18,
  symbol: 'HIGH',
  name: 'Highstreet',
  logo: highstreetLogo,
  description: 'HIGH is the token behind Highstreet',
  swapableType: 'HIGHSTREET',
})

export const JUMP = new Token({
  addresses: {
    [ChainId.BSC]: '0x130025ee738a66e691e6a7a62381cb33c6d9ae83', // redeem
    [ChainId.FANTOM]: '0x78DE9326792ce1d6eCA0c978753c6953Cdeedd73', // deposit
    [ChainId.METIS]: '0xE3c82A836Ec85311a433fBd9486EfAF4b1AFbF48', // redeem
  },
  decimals: 18,
  symbol: 'JUMP',
  name: 'HyperJump',
  logo: hyperjumpLogo,
  description: 'JUMP is the token behind Hyperjump',
  docUrl: '',
  swapableType: 'JUMP',
})

export const SFI = new Token({
  addresses: {
    [ChainId.ETH]: '0xb753428af26e81097e7fd17f40c88aaa3e04902c',
    [ChainId.AVALANCHE]: '0xc2Bf0A1f7D8Da50D608bc96CF701110d4A438312', // deposit
  },
  decimals: 18,
  symbol: 'SFI',
  name: 'Saffron Finance',
  logo: sfiLogo,
  description: '',
  docUrl: '',
  swapableType: 'SFI',
})

export const DOG = new Token({
  addresses: {
    [ChainId.ETH]: '0xBAac2B4491727D78D2b78815144570b9f2Fe8899',
    [ChainId.BSC]: '0xaa88c603d142c371ea0eac8756123c5805edee03',
    [ChainId.POLYGON]: '0xeEe3371B89FC43Ea970E908536Fcddd975135D8a',
    // [ChainId.ARBITRUM]: '0x4425742F1EC8D98779690b5A3A6276Db85Ddc01A'
  },
  decimals: 18,
  symbol: 'DOG',
  name: 'The Doge NFT',
  logo: dogLogo,
  description: 'DOG is the token behind the Doge NFT',
  docUrl: '',
  swapableType: 'DOG',
})

export const NFD = new Token({
  addresses: {
    [ChainId.BSC]: '0x0fe9778c005a5a6115cbe12b0568a2d50b765a51', // redeem
    [ChainId.AVALANCHE]: '0xf1293574ee43950e7a8c9f1005ff097a9a713959', // redeem
    [ChainId.DOGECHAIN]: '0x868055ADFA27D331d5b69b1BF3469aDAAc3ba7f2', // redeem
    [ChainId.POLYGON]: '0x0a5926027d407222f8fe20f24cb16e103f617046', // deposit
  },
  decimals: 18,
  symbol: 'NFD',
  name: 'Feisty Doge',
  logo: nfdLogo,
  description: 'Feisty Doge NFT',
  docUrl: '',
  swapableType: 'NFD',
})

export const SOLAR = new Token({
  addresses: {
    [ChainId.MOONBEAM]: '0x0DB6729C03C85B0708166cA92801BcB5CAc781fC', // redeem
    [ChainId.MOONRIVER]: '0x76906411D07815491A5E577022757aD941fb5066', // deposit
  },
  decimals: 18,
  symbol: 'veSOLAR',
  name: 'Vested SolarBeam',
  logo: solarbeamLogo,
  description: 'Vested SolarBeam',
  docUrl: '',
  swapableType: 'SOLAR',
})

export const GMX = new Token({
  addresses: {
    [ChainId.ARBITRUM]: '0xfc5a1a6eb076a2c7ad06ed22c90d7e710e35ad0a', // deposit
    [ChainId.AVALANCHE]: '0x62edc0692bd897d2295872a9ffcac5425011c661', // redeem
  },
  wrapperAddresses: {
    [ChainId.AVALANCHE]: '0x20A9DC684B4d0407EF8C9A302BEAaA18ee15F656',
  },
  decimals: 18,
  symbol: 'GMX',
  name: 'GMX',
  logo: gmxLogo,
  description: 'GMX Financial',
  docUrl: '',
  swapableType: 'GMX',
})

export const SDT = new Token({
  addresses: {
    [ChainId.ETH]: '0x73968b9a57c6e53d41345fd57a6e6ae27d6cdb2f', // deposit
    [ChainId.AVALANCHE]: '0xCCBf7c451F81752F7d2237F2c18C371E6e089E69', // redeem
    [ChainId.ARBITRUM]: '0x087d18A77465c34CDFd3a081a2504b7E86CE4EF8',
    [ChainId.FANTOM]: '0xE3c82A836Ec85311a433fBd9486EfAF4b1AFbF48', // redeem
    [ChainId.HARMONY]: '0xE3c82A836Ec85311a433fBd9486EfAF4b1AFbF48', // redeem
  },
  decimals: 18,
  symbol: 'SDT',
  name: 'Stake DAO',
  logo: sdtLogo,
  description: 'Stake DAO',
  docUrl: '',
  swapableType: 'SDT',
})

export const NEWO = new Token({
  addresses: {
    [ChainId.ETH]: '0x98585dFc8d9e7D48F0b1aE47ce33332CF4237D96', // deposit
    [ChainId.AVALANCHE]: '0x4Bfc90322dD638F81F034517359BD447f8E0235a', // redeem
    [ChainId.ARBITRUM]: '0x0877154a755B24D499B8e2bD7ecD54d3c92BA433', // redeem
  },
  decimals: 18,
  symbol: 'NEWO',
  name: 'New Order',
  logo: newoLogo,
  description: 'New Order',
  docUrl: '',
  swapableType: 'NEWO',
})

export const USDB = new Token({
  addresses: {
    [ChainId.ETH]: '0x02b5453d92b730f29a86a0d5ef6e930c4cf8860b',
    [ChainId.BSC]: '0xc8699abbba90c7479dedccef19ef78969a2fc608',
    [ChainId.POLYGON]: '0xfa1fbb8ef55a4855e5688c0ee13ac3f202486286',
    [ChainId.FANTOM]: '0x6fc9383486c163fa48becdec79d6058f984f62ca',
    [ChainId.AVALANCHE]: '0x5ab7084cb9d270c2cb052dd30dbecbca42f8620c',
    [ChainId.MOONRIVER]: '0x3e193c39626bafb41ebe8bdd11ec7cca9b3ec0b2',
  },
  decimals: 18,
  symbol: 'USDB',
  name: 'USDB',
  logo: usdbLogo,
  description: 'USDB',
  docUrl: '',
  swapableType: 'USDB',
})

export const VSTA = new Token({
  addresses: {
    [ChainId.ETH]: '0xA8d7F5e7C78ed0Fa097Cc5Ec66C1DC3104c9bbeb', // redeem
    [ChainId.ARBITRUM]: '0xa684cd057951541187f288294a1e1c2646aa2d24', // deposit
  },
  decimals: 18,
  symbol: 'VSTA',
  name: 'Vesta',
  logo: vstaLogo,
  description: 'Vesta Finance',
  docUrl: '',
  swapableType: 'VSTA',
})

export const H2O = new Token({
  addresses: {
    [ChainId.ETH]: '0x0642026e7f0b6ccac5925b4e7fa61384250e1701', // deposit
    [ChainId.ARBITRUM]: '0xD1c6f989e9552DB523aBAE2378227fBb059a3976', // redeem
    [ChainId.AVALANCHE]: '0xC6b11a4Fd833d1117E9D312c02865647cd961107', // redeem
    [ChainId.BSC]: '0x03eFca7CEb108734D3777684F3C0A0d8ad652f79', // redeem
    [ChainId.MOONBEAM]: '0xA46aDF6D5881ca0b8596EDadF8f058F8c16d8B68', // redeem
    [ChainId.MOONRIVER]: '0x9c0a820bb01e2807aCcd1c682d359e92DDd41403', // redeem
    [ChainId.OPTIMISM]: '0xE3c82A836Ec85311a433fBd9486EfAF4b1AFbF48', // redeem
    [ChainId.POLYGON]: '0x32ba7cF7d681357529013de6a2CDF93933C0dF3f', // redeem
  },
  decimals: 18,
  symbol: 'H2O',
  name: 'H2O',
  logo: h2oLogo,
  description: 'H2O',
  docUrl: '',
  swapableType: 'H2O',
})

export const L2DAO = new Token({
  addresses: {
    [ChainId.ARBITRUM]: '0x2CaB3abfC1670D1a452dF502e216a66883cDf079', // deposit
    [ChainId.OPTIMISM]: '0xd52f94DF742a6F4B4C8b033369fE13A41782Bf44', // redeem
  },
  decimals: 18,
  symbol: 'L2DAO',
  name: 'Layer2DAO',
  logo: l2daoLogo,
  description: 'Layer2DAO',
  docUrl: '',
  swapableType: 'L2DAO',
})

export const PLS = new Token({
  addresses: {
    [ChainId.ARBITRUM]: '0x51318b7d00db7acc4026c88c3952b66278b6a67f', // deposit
    [ChainId.OPTIMISM]: '0xD9eAA386cCD65F30b77FF175F6b52115FE454fD6', // redeem
  },
  decimals: 18,
  symbol: 'PLS',
  name: 'Plutus',
  logo: plsLogo,
  description: 'PlutusDao',
  docUrl: '',
  swapableType: 'PLS',
})

export const AGEUR = new Token({
  addresses: {
    [ChainId.ETH]: '0x1a7e4e63778B4f12a199C062f3eFdD288afCBce8', // deposit
    [ChainId.ARBITRUM]: '0x16BFc5fe024980124bEf51d1D792dC539d1B5Bf0', // redeem
    [ChainId.OPTIMISM]: '0xa0554607e477cdC9d0EE2A6b087F4b2DC2815C22', // redeem
  },
  decimals: 18,
  symbol: 'agEUR',
  name: 'Angle Euro',
  logo: ageurLogo,
  description: 'Angle Euro',
  docUrl: '',
  swapableType: 'AGEUR',
})

export const UNIDX = new Token({
  addresses: {
    [ChainId.ETH]: '0xf0655dcee37e5c0b70fffd70d85f88f8edf0aff6', // deposit
    [ChainId.ARBITRUM]: '0x5429706887FCb58a595677B73E9B0441C25d993D', // redeem
    [ChainId.BASE]: '0x6B4712AE9797C199edd44F897cA09BC57628a1CF',
    [ChainId.OPTIMISM]: '0x28b42698Caf46B4B012CF38b6C75867E0762186D',
    [ChainId.FANTOM]: '0x0483a76d80d0afec6bd2afd12c1ad865b9df1471',
  },
  decimals: 18,
  symbol: 'UNIDX',
  name: 'Unidex',
  logo: unidexLogo,
  description: 'Unidex',
  docUrl: '',
  swapableType: 'UNIDX',
})
