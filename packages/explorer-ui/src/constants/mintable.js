import dogLogo from 'assets/icons/dog.png'
import gmxLogo from 'assets/icons/gmx.png'
import highstreetLogo from 'assets/icons/highstreet.svg'
import hyperjumpLogo from 'assets/icons/hyperjump.png'
import nfdLogo from 'assets/icons/nfd.svg'
import ohmLogo from 'assets/icons/ohm.svg'
import solarbeamLogo from 'assets/icons/solarbeam.png'
import { CHAIN_ID } from './chains'
import { Token } from './Token'

export const GOHM = new Token({
  addresses: {
    [CHAIN_ID.ETH]: '0x0ab87046fBb341D058F17CBC4c1133F25a20a52f',
    [CHAIN_ID.OPTIMISM]: '0x0b5740c6b4a97f90eF2F0220651Cca420B868FfB',
    [CHAIN_ID.BSC]: '0x88918495892BAF4536611E38E75D771Dc6Ec0863',
    [CHAIN_ID.POLYGON]: '0xd8cA34fd379d9ca3C6Ee3b3905678320F5b45195',
    [CHAIN_ID.FANTOM]: '0x91fa20244Fb509e8289CA630E5db3E9166233FDc',
    [CHAIN_ID.ARBITRUM]: '0x8D9bA570D6cb60C7e3e0F31343Efe75AB8E65FB1',
    [CHAIN_ID.AVALANCHE]: '0x321E7092a180BB43555132ec53AaA65a5bF84251',
    [CHAIN_ID.MOONRIVER]: '0x3bF21Ce864e58731B6f28D68d5928BcBEb0Ad172',
    [CHAIN_ID.BOBA]: '0xd22C0a4Af486C7FA08e282E9eB5f30F9AaA62C95',
    [CHAIN_ID.HARMONY]: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
    [CHAIN_ID.MOONBEAM]: '0xD2666441443DAa61492FFe0F37717578714a4521',
    [CHAIN_ID.CRONOS]: '0xbB0A63A6CA2071c6C4bcAC11a1A317b20E3E999C',
    [CHAIN_ID.METIS]: '0xFB21B70922B9f6e3C6274BcD6CB1aa8A0fe20B80',
  },
  decimals: 18,
  symbol: 'gOHM',
  name: 'Olympus DAO',
  logo: ohmLogo,
  description: 'OHM',
  swapableType: 'OHM',
})

export const HIGHSTREET = new Token({
  addresses: {
    [CHAIN_ID.ETH]: '0x71Ab77b7dbB4fa7e017BC15090b2163221420282',
    [CHAIN_ID.BSC]: '0x5f4bde007dc06b867f86ebfe4802e34a1ffeed63',
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
    [CHAIN_ID.BSC]: '0x130025ee738a66e691e6a7a62381cb33c6d9ae83', // redeem
    [CHAIN_ID.FANTOM]: '0x78DE9326792ce1d6eCA0c978753c6953Cdeedd73', // deposit
  },
  decimals: 18,
  symbol: 'JUMP',
  name: 'HyperJump',
  logo: hyperjumpLogo,
  description: 'JUMP is the token behind Hyperjump',
  docUrl: '',
  swapableType: 'JUMP',
})

export const DOG = new Token({
  addresses: {
    [CHAIN_ID.ETH]: '0xBAac2B4491727D78D2b78815144570b9f2Fe8899',
    [CHAIN_ID.BSC]: '0xaa88c603d142c371ea0eac8756123c5805edee03',
    [CHAIN_ID.POLYGON]: '0xeEe3371B89FC43Ea970E908536Fcddd975135D8a',
    // [CHAIN_ID.ARBITRUM]: '0x4425742F1EC8D98779690b5A3A6276Db85Ddc01A'
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
    [CHAIN_ID.BSC]: '0x0fe9778c005a5a6115cbe12b0568a2d50b765a51', // redeem
    [CHAIN_ID.AVALANCHE]: '0xf1293574ee43950e7a8c9f1005ff097a9a713959', // redeem
    [CHAIN_ID.POLYGON]: '0x0a5926027d407222f8fe20f24cb16e103f617046', // deposit
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
    [CHAIN_ID.MOONBEAM]: '0x0DB6729C03C85B0708166cA92801BcB5CAc781fC', // redeem
    [CHAIN_ID.MOONRIVER]: '0x76906411D07815491A5E577022757aD941fb5066', // deposit
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
    [CHAIN_ID.ARBITRUM]: '0xfc5a1a6eb076a2c7ad06ed22c90d7e710e35ad0a', // deposit
    [CHAIN_ID.AVALANCHE]: '0x62edc0692bd897d2295872a9ffcac5425011c661', // redeem
  },
  wrapperAddresses: {
    [CHAIN_ID.AVALANCHE]: '0x20A9DC684B4d0407EF8C9A302BEAaA18ee15F656',
  },
  decimals: 18,
  symbol: 'GMX',
  name: 'GMX',
  logo: gmxLogo,
  description: 'GMX Financial',
  docUrl: '',
  swapableType: 'GMX',
})
