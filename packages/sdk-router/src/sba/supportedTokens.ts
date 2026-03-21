import { getAddress } from '@ethersproject/address'

import { SupportedChainId } from '../constants'
import { BridgeTokenCandidate } from '../module'
import { ETH_NATIVE_TOKEN_ADDRESS, isSameAddress } from '../utils'
import { SBA_CHAIN_METADATA } from './metadata'

type SupportedTokenSnapshotEntry = {
  readonly tokenId: string
  readonly chains: Partial<Record<SupportedChainId, string>>
  readonly wrappedNativeChainId?: SupportedChainId
}

type SupportedTokenPairLookup = Partial<
  Record<
    SupportedChainId,
    Partial<Record<SupportedChainId, BridgeTokenCandidate[]>>
  >
>

type SupportedTokenRemoteLookup = Partial<
  Record<
    SupportedChainId,
    Record<string, Partial<Record<SupportedChainId, string>>>
  >
>

const EXPECTED_TOKEN_ID_COUNT = 37
const EXPECTED_ORIGIN_ENTRY_COUNT = 149

// wrappedNativeChainId marks the snapshot chain entry confirmed to be the local
// wrapped native token via live on-chain research on March 18, 2026.
const SBA_SUPPORTED_TOKEN_SNAPSHOT: readonly SupportedTokenSnapshotEntry[] = [
  {
    tokenId: 'AVAX',
    chains: {
      [SupportedChainId.AVALANCHE]:
        '0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7',
      [SupportedChainId.DFK]: '0xB57B60DeBDB0b8172bb6316a9164bd3C695F133a',
      [SupportedChainId.HARMONY]: '0xD9eAA386cCD65F30b77FF175F6b52115FE454fD6',
      [SupportedChainId.KLAYTN]: '0xCd8fE44A29Db9159dB36f96570d7A4d91986f528',
    },
    wrappedNativeChainId: SupportedChainId.AVALANCHE,
  },
  {
    tokenId: 'BTCB',
    chains: {
      [SupportedChainId.AVALANCHE]:
        '0x152b9d0FdC40C096757F570A51E494bd4b943E50',
      [SupportedChainId.DFK]: '0x7516EB8B8Edfa420f540a162335eACF3ea05a247',
      [SupportedChainId.KLAYTN]: '0xe82f87ba4E97b2796aA0Fa4eFB06e8f0d2EB4FE1',
    },
  },
  {
    tokenId: 'DAI',
    chains: {
      [SupportedChainId.ETH]: '0x6B175474E89094C44Da98b954EedeAC495271d0F',
      [SupportedChainId.KLAYTN]: '0x078dB7827a5531359f6CB63f62CFA20183c4F10c',
    },
  },
  {
    tokenId: 'DOG',
    chains: {
      [SupportedChainId.BSC]: '0xaA88C603d142C371eA0eAC8756123c5805EdeE03',
      [SupportedChainId.ETH]: '0xBAac2B4491727D78D2b78815144570b9f2Fe8899',
      [SupportedChainId.POLYGON]: '0xeEe3371B89FC43Ea970E908536Fcddd975135D8a',
    },
  },
  {
    tokenId: 'FLEX',
    chains: {
      [SupportedChainId.ARBITRUM]: '0x14f37e4d9822BAf2e956C65f797B76Fa7a914B9F',
      [SupportedChainId.BSC]: '0xC9b30623C09f47D5a009D47Bc90D19e0Fd7ec372',
      [SupportedChainId.ETH]: '0xFcF8eda095e37A41e002E266DaAD7efC1579bc0A',
    },
  },
  {
    tokenId: 'FTM',
    chains: {
      [SupportedChainId.DFK]: '0x2Df041186C844F8a2e2b63F16145Bc6Ff7d23E25',
      [SupportedChainId.FANTOM]: '0x21be370D5312f44cB42ce377BC9b8a0cEF1A4C83',
    },
    wrappedNativeChainId: SupportedChainId.FANTOM,
  },
  {
    tokenId: 'H2O',
    chains: {
      [SupportedChainId.ARBITRUM]: '0xD1c6f989e9552DB523aBAE2378227fBb059a3976',
      [SupportedChainId.AVALANCHE]:
        '0xC6b11a4Fd833d1117E9D312c02865647cd961107',
      [SupportedChainId.BSC]: '0x03eFca7CEb108734D3777684F3C0A0d8ad652f79',
      [SupportedChainId.ETH]: '0x0642026E7f0B6cCaC5925b4E7Fa61384250e1701',
      [SupportedChainId.OPTIMISM]: '0xE3c82A836Ec85311a433fBd9486EfAF4b1AFbF48',
      [SupportedChainId.POLYGON]: '0x32ba7cF7d681357529013de6a2CDF93933C0dF3f',
    },
  },
  {
    tokenId: 'HIGH',
    chains: {
      [SupportedChainId.BSC]: '0x5f4Bde007Dc06b867f86EBFE4802e34A1fFEEd63',
      [SupportedChainId.ETH]: '0x71Ab77b7dbB4fa7e017BC15090b2163221420282',
    },
  },
  {
    tokenId: 'JEWEL',
    chains: {
      [SupportedChainId.AVALANCHE]:
        '0x997Ddaa07d716995DE90577C123Db411584E5E46',
      [SupportedChainId.DFK]: '0xCCb93dABD71c8Dad03Fc4CE5559dC3D89F67a260',
      [SupportedChainId.ETH]: '0x12f79f8c1A6e47a9b5F0796FDb008Bdc182fa19e',
      [SupportedChainId.HARMONY]: '0x28b42698Caf46B4B012CF38b6C75867E0762186D',
      [SupportedChainId.KLAYTN]: '0x30C103f8f5A3A732DFe2dCE1Cc9446f545527b43',
      [SupportedChainId.METIS]: '0x17C09cfC96C865CF546d73365Cedb6dC66986963',
    },
    wrappedNativeChainId: SupportedChainId.DFK,
  },
  {
    tokenId: 'JUMP',
    chains: {
      [SupportedChainId.BSC]: '0x130025eE738A66E691E6A7a62381CB33c6d9Ae83',
      [SupportedChainId.FANTOM]: '0x78DE9326792ce1d6eCA0c978753c6953Cdeedd73',
      [SupportedChainId.METIS]: '0xE3c82A836Ec85311a433fBd9486EfAF4b1AFbF48',
    },
  },
  {
    tokenId: 'KLAY',
    chains: {
      [SupportedChainId.DFK]: '0x97855Ba65aa7ed2F65Ed832a776537268158B78a',
      [SupportedChainId.KLAYTN]: '0x5819b6af194A78511c79C85Ea68D2377a7e9335f',
    },
    wrappedNativeChainId: SupportedChainId.KLAYTN,
  },
  {
    tokenId: 'L2DAO',
    chains: {
      [SupportedChainId.ARBITRUM]: '0x2CaB3abfC1670D1a452dF502e216a66883cDf079',
      [SupportedChainId.OPTIMISM]: '0xd52f94DF742a6F4B4C8b033369fE13A41782Bf44',
    },
  },
  {
    tokenId: 'LINK',
    chains: {
      [SupportedChainId.ETH]: '0x514910771AF9Ca656af840dff83E8264EcF986CA',
      [SupportedChainId.KLAYTN]: '0xfbEd1AbB3aD0f8C467068De9fDE905887e8C9118',
    },
  },
  {
    tokenId: 'MATIC',
    chains: {
      [SupportedChainId.DFK]: '0xD17a41Cd199edF1093A9Be4404EaDe52Ec19698e',
      [SupportedChainId.POLYGON]: '0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270',
    },
    wrappedNativeChainId: SupportedChainId.POLYGON,
  },
  {
    tokenId: 'Metis',
    chains: {
      [SupportedChainId.DFK]: '0x43E8E55792D6317328e5c6B0A0C89eF4b8102Fa2',
      [SupportedChainId.KLAYTN]: '0x543f1b2176F7E677a95D01ca83551FAa08F83D9F',
      [SupportedChainId.METIS]: '0x75cb093E4D61d2A2e65D8e0BBb01DE8d89b53481',
    },
    wrappedNativeChainId: SupportedChainId.METIS,
  },
  {
    tokenId: 'NEWO',
    chains: {
      [SupportedChainId.ARBITRUM]: '0x0877154a755B24D499B8e2bD7ecD54d3c92BA433',
      [SupportedChainId.AVALANCHE]:
        '0x4Bfc90322dD638F81F034517359BD447f8E0235a',
      [SupportedChainId.ETH]: '0x98585dFc8d9e7D48F0b1aE47ce33332CF4237D96',
    },
  },
  {
    tokenId: 'NFD',
    chains: {
      [SupportedChainId.AVALANCHE]:
        '0xf1293574EE43950E7a8c9F1005Ff097A9A713959',
      [SupportedChainId.BSC]: '0x0FE9778c005a5A6115cBE12b0568a2d50b765A51',
      [SupportedChainId.POLYGON]: '0x0A5926027d407222F8fe20f24cB16e103f617046',
    },
  },
  {
    tokenId: 'PEPE',
    chains: {
      [SupportedChainId.ARBITRUM]: '0xA54B8e178A49F8e5405A4d44Bb31F496e5564A05',
      [SupportedChainId.BSC]: '0xd2b6F20aa2611e8a7a18e5EeC58ca8369f5D356b',
      [SupportedChainId.ETH]: '0x6982508145454Ce325dDbE47a25d4ec3d2311933',
    },
  },
  {
    tokenId: 'PLS',
    chains: {
      [SupportedChainId.ARBITRUM]: '0x51318B7D00db7ACc4026C88c3952B66278B6A67F',
      [SupportedChainId.OPTIMISM]: '0xD9eAA386cCD65F30b77FF175F6b52115FE454fD6',
    },
  },
  {
    tokenId: 'SDT',
    chains: {
      [SupportedChainId.ARBITRUM]: '0x087d18A77465c34CDFd3a081a2504b7E86CE4EF8',
      [SupportedChainId.AVALANCHE]:
        '0xCCBf7c451F81752F7d2237F2c18C371E6e089E69',
      [SupportedChainId.ETH]: '0x73968b9a57c6E53d41345FD57a6E6ae27d6CDB2F',
      [SupportedChainId.FANTOM]: '0xE3c82A836Ec85311a433fBd9486EfAF4b1AFbF48',
      [SupportedChainId.HARMONY]: '0xE3c82A836Ec85311a433fBd9486EfAF4b1AFbF48',
    },
  },
  {
    tokenId: 'SFI',
    chains: {
      [SupportedChainId.AVALANCHE]:
        '0xc2Bf0A1f7D8Da50D608bc96CF701110d4A438312',
      [SupportedChainId.ETH]: '0xb753428af26E81097e7fD17f40c88aaA3E04902c',
    },
  },
  {
    tokenId: 'SPEC',
    chains: {
      [SupportedChainId.BASE]: '0x96419929d7949D6A801A6909c145C8EEf6A40431',
      [SupportedChainId.ETH]: '0xAdF7C35560035944e805D98fF17d58CDe2449389',
    },
  },
  {
    tokenId: 'SYN',
    chains: {
      [SupportedChainId.ARBITRUM]: '0x080F6AEd32Fc474DD5717105Dba5ea57268F46eb',
      [SupportedChainId.AVALANCHE]:
        '0x1f1E7c893855525b303f99bDF5c3c05Be09ca251',
      [SupportedChainId.BASE]: '0x432036208d2717394d2614d6697c46DF3Ed69540',
      [SupportedChainId.BLAST]: '0x9592f08387134e218327E6E8423400eb845EdE0E',
      [SupportedChainId.BSC]: '0xa4080f1778e69467E905B8d6F72f6e441f9e9484',
      [SupportedChainId.CANTO]: '0x555982d2E211745b96736665e19D9308B615F78e',
      [SupportedChainId.ETH]: '0x0f2D719407FdBeFF09D87557AbB7232601FD9F29',
      [SupportedChainId.FANTOM]: '0xE55e19Fb4F2D85af758950957714292DAC1e25B2',
      [SupportedChainId.HARMONY]: '0xE55e19Fb4F2D85af758950957714292DAC1e25B2',
      [SupportedChainId.METIS]: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
      [SupportedChainId.OPTIMISM]: '0x5A5fFf6F753d7C11A56A52FE47a177a87e431655',
      [SupportedChainId.POLYGON]: '0xf8F9efC0db77d8881500bb06FF5D6ABc3070E695',
    },
  },
  {
    tokenId: 'UNIDX',
    chains: {
      [SupportedChainId.ARBITRUM]: '0x5429706887FCb58a595677B73E9B0441C25d993D',
      [SupportedChainId.BASE]: '0x6B4712AE9797C199edd44F897cA09BC57628a1CF',
      [SupportedChainId.ETH]: '0xf0655DcEE37E5C0b70Fffd70D85f88F8eDf0AfF6',
      [SupportedChainId.FANTOM]: '0x0483a76D80D0aFEC6bd2afd12C1AD865b9DF1471',
      [SupportedChainId.OPTIMISM]: '0x28b42698Caf46B4B012CF38b6C75867E0762186D',
    },
  },
  {
    tokenId: 'USDC',
    chains: {
      [SupportedChainId.ETH]: '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48',
      [SupportedChainId.KLAYTN]: '0x6270B58BE569a7c0b8f47594F191631Ae5b2C86C',
    },
  },
  {
    tokenId: 'USDT',
    chains: {
      [SupportedChainId.ETH]: '0xdAC17F958D2ee523a2206206994597C13D831ec7',
      [SupportedChainId.KLAYTN]: '0xd6dAb4CfF47dF175349e6e7eE2BF7c40Bb8C05A3',
    },
  },
  {
    tokenId: 'UST',
    chains: {
      [SupportedChainId.ARBITRUM]: '0x13780E6d5696DD91454F6d3BbC2616687fEa43d0',
      [SupportedChainId.AVALANCHE]:
        '0xE97097dE8d6A17Be3c39d53AE63347706dCf8f43',
      [SupportedChainId.BSC]: '0xb7A6c5f0cc98d24Cf4B2011842e64316Ff6d042c',
      [SupportedChainId.DFK]: '0x360d6DD540E3448371876662FBE7F1aCaf08c5Ab',
      [SupportedChainId.ETH]: '0x0261018Aa50E28133C1aE7a29ebdf9Bd21b878Cb',
      [SupportedChainId.FANTOM]: '0xa0554607e477cdC9d0EE2A6b087F4b2DC2815C22',
      [SupportedChainId.HARMONY]: '0xa0554607e477cdC9d0EE2A6b087F4b2DC2815C22',
      [SupportedChainId.METIS]: '0x0b5740c6b4a97f90eF2F0220651Cca420B868FfB',
      [SupportedChainId.OPTIMISM]: '0xFB21B70922B9f6e3C6274BcD6CB1aa8A0fe20B80',
      [SupportedChainId.POLYGON]: '0x565098CBa693b3325f9fe01D41b7A1cd792Abab1',
    },
  },
  {
    tokenId: 'VSTA',
    chains: {
      [SupportedChainId.ARBITRUM]: '0xa684cd057951541187f288294a1e1C2646aA2d24',
      [SupportedChainId.ETH]: '0xA8d7F5e7C78ed0Fa097Cc5Ec66C1DC3104c9bbeb',
    },
  },
  {
    tokenId: 'WBTC',
    chains: {
      [SupportedChainId.ETH]: '0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599',
      [SupportedChainId.KLAYTN]: '0xDCbacF3f7a069922E677912998c8d57423C37dfA',
    },
  },
  {
    tokenId: 'agEUR',
    chains: {
      [SupportedChainId.ARBITRUM]: '0x16BFc5fe024980124bEf51d1D792dC539d1B5Bf0',
      [SupportedChainId.ETH]: '0x1a7e4e63778B4f12a199C062f3eFdD288afCBce8',
      [SupportedChainId.OPTIMISM]: '0xa0554607e477cdC9d0EE2A6b087F4b2DC2815C22',
    },
  },
  {
    tokenId: 'fantOHM',
    chains: {
      [SupportedChainId.AVALANCHE]:
        '0x5aB7084CB9d270c2Cb052dd30dbecBCA42F8620c',
      [SupportedChainId.BSC]: '0xc8699AbBba90C7479dedcCEF19eF78969a2fc608',
      [SupportedChainId.ETH]: '0x02B5453D92B730F29a86A0D5ef6e930c4Cf8860B',
      [SupportedChainId.FANTOM]: '0x6Fc9383486c163fA48becdEC79d6058f984f62cA',
      [SupportedChainId.POLYGON]: '0xfa1FBb8Ef55A4855E5688C0eE13aC3f202486286',
    },
  },
  {
    tokenId: 'gOHM',
    chains: {
      [SupportedChainId.ARBITRUM]: '0x8D9bA570D6cb60C7e3e0F31343Efe75AB8E65FB1',
      [SupportedChainId.AVALANCHE]:
        '0x321E7092a180BB43555132ec53AaA65a5bF84251',
      [SupportedChainId.BSC]: '0x88918495892BAF4536611E38E75D771Dc6Ec0863',
      [SupportedChainId.ETH]: '0x0ab87046fBb341D058F17CBC4c1133F25a20a52f',
      [SupportedChainId.FANTOM]: '0x91fa20244Fb509e8289CA630E5db3E9166233FDc',
      [SupportedChainId.HARMONY]: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
      [SupportedChainId.METIS]: '0xFB21B70922B9f6e3C6274BcD6CB1aa8A0fe20B80',
      [SupportedChainId.OPTIMISM]: '0x0b5740c6b4a97f90eF2F0220651Cca420B868FfB',
      [SupportedChainId.POLYGON]: '0xd8cA34fd379d9ca3C6Ee3b3905678320F5b45195',
    },
  },
  {
    tokenId: 'nETH',
    chains: {
      [SupportedChainId.ARBITRUM]: '0x3ea9B0ab55F34Fb188824Ee288CeaEfC63cf908e',
      [SupportedChainId.AVALANCHE]:
        '0x19E1ae0eE35c0404f835521146206595d37981ae',
      [SupportedChainId.BASE]: '0xb554A55358fF0382Fb21F0a478C3546d1106Be8c',
      [SupportedChainId.BLAST]: '0xce971282fAAc9faBcF121944956da7142cccC855',
      [SupportedChainId.CANTO]: '0x09fEC30669d63A13c666d2129230dD5588E2e240',
      [SupportedChainId.DFK]: '0xfBDF0E31808d0aa7b9509AA6aBC9754E48C58852',
      [SupportedChainId.ETH]: '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2',
      [SupportedChainId.FANTOM]: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
      [SupportedChainId.HARMONY]: '0x0b5740c6b4a97f90eF2F0220651Cca420B868FfB',
      [SupportedChainId.KLAYTN]: '0xCD6f29dC9Ca217d0973d3D21bF58eDd3CA871a86',
      [SupportedChainId.METIS]: '0x931B8f17764362A3325D30681009f0eDd6211231',
      [SupportedChainId.OPTIMISM]: '0x809DC529f07651bD43A172e8dB6f4a7a0d771036',
    },
    wrappedNativeChainId: SupportedChainId.ETH,
  },
  {
    tokenId: 'nUSD',
    chains: {
      [SupportedChainId.ARBITRUM]: '0x2913E812Cf0dcCA30FB28E6Cac3d2DCFF4497688',
      [SupportedChainId.AVALANCHE]:
        '0xCFc37A6AB183dd4aED08C204D1c2773c0b1BDf46',
      [SupportedChainId.BLAST]: '0x3194B0A295D87fDAA54DF852c248F7a6BAF6c6e0',
      [SupportedChainId.BSC]: '0x23b891e5C62E0955ae2bD185990103928Ab817b3',
      [SupportedChainId.CANTO]: '0xD8836aF2e565D3Befce7D906Af63ee45a57E8f80',
      [SupportedChainId.DFK]: '0x3AD9DFE640E1A9Cc1D9B0948620820D975c3803a',
      [SupportedChainId.ETH]: '0x1B84765dE8B7566e4cEAF4D0fD3c5aF52D3DdE4F',
      [SupportedChainId.FANTOM]: '0xED2a7edd7413021d440b09D654f3b87712abAB66',
      [SupportedChainId.HARMONY]: '0xED2a7edd7413021d440b09D654f3b87712abAB66',
      [SupportedChainId.METIS]: '0x961318Fc85475E125B99Cc9215f62679aE5200aB',
      [SupportedChainId.OPTIMISM]: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
      [SupportedChainId.POLYGON]: '0xB6c473756050dE474286bED418B77Aeac39B02aF',
    },
  },
  {
    tokenId: 'synFRAX',
    chains: {
      [SupportedChainId.ETH]: '0x853d955aCEf822Db058eb8505911ED77F175b99e',
      [SupportedChainId.FANTOM]: '0x1852F70512298d56e9c8FDd905e02581E04ddb2a',
      [SupportedChainId.HARMONY]: '0x1852F70512298d56e9c8FDd905e02581E04ddb2a',
    },
  },
  {
    tokenId: 'wsOHM',
    chains: {
      [SupportedChainId.ARBITRUM]: '0x30bD4e574a15994B35EF9C7a5bc29002F1224821',
      [SupportedChainId.AVALANCHE]:
        '0x240E332Cd26AaE10622B24160D23425A17256F5d',
      [SupportedChainId.ETH]: '0xCa76543Cf381ebBB277bE79574059e32108e3E65',
    },
  },
  {
    tokenId: 'xJEWEL',
    chains: {
      [SupportedChainId.DFK]: '0x77f2656d04E158f915bC22f07B779D94c1DC47Ff',
      [SupportedChainId.HARMONY]: '0xA9cE83507D872C5e1273E745aBcfDa849DAA654F',
    },
  },
] as const

const buildSupportedTokenLookups = (
  snapshot: readonly SupportedTokenSnapshotEntry[]
): {
  pairLookup: SupportedTokenPairLookup
  remoteLookup: SupportedTokenRemoteLookup
} => {
  const pairLookup: SupportedTokenPairLookup = {}
  const remoteLookup: SupportedTokenRemoteLookup = {}
  const tokenIds = new Set<string>()
  let originEntryCount = 0

  snapshot.forEach(({ tokenId, chains, wrappedNativeChainId }) => {
    const chainEntries = Object.entries(chains)
      .filter((entry): entry is [string, string] => entry[1] !== undefined)
      .map(
        ([chainId, token]) =>
          [Number(chainId) as SupportedChainId, getAddress(token)] as const
      )
      .sort((a, b) => a[0] - b[0])

    if (chainEntries.length < 2) {
      throw new Error(
        `SBA supported token snapshot entry ${tokenId} must exist on at least two supported chains`
      )
    }

    tokenIds.add(tokenId)

    if (wrappedNativeChainId !== undefined) {
      const chainId = wrappedNativeChainId
      if (!SBA_CHAIN_METADATA[chainId]) {
        throw new Error(
          `SBA supported token snapshot entry ${tokenId} uses unsupported wrapped-native chain ${chainId}`
        )
      }
      const chainToken = chains[chainId]
      if (!chainToken) {
        throw new Error(
          `SBA supported token snapshot entry ${tokenId} marks chain ${chainId} as wrapped native but no token is defined`
        )
      }
    }

    chainEntries.forEach(([originChainId, originToken]) => {
      if (!SBA_CHAIN_METADATA[originChainId]) {
        throw new Error(
          `SBA supported token snapshot entry ${tokenId} uses unsupported origin chain ${originChainId}`
        )
      }

      const originLookup = remoteLookup[originChainId] || {}
      remoteLookup[originChainId] = originLookup
      const originTokenKey = originToken.toLowerCase()
      if (originLookup[originTokenKey]) {
        throw new Error(
          `SBA supported token snapshot has duplicate origin token mapping for chain ${originChainId} and token ${originToken}`
        )
      }

      const remoteTokens: Partial<Record<SupportedChainId, string>> = {}

      chainEntries.forEach(([destChainId, destToken]) => {
        if (destChainId === originChainId) {
          return
        }
        if (!SBA_CHAIN_METADATA[destChainId]) {
          throw new Error(
            `SBA supported token snapshot entry ${tokenId} uses unsupported destination chain ${destChainId}`
          )
        }
        const effectiveDestToken =
          destChainId === wrappedNativeChainId
            ? ETH_NATIVE_TOKEN_ADDRESS
            : destToken
        const existingDestToken = remoteTokens[destChainId]
        if (
          existingDestToken &&
          !isSameAddress(existingDestToken, effectiveDestToken)
        ) {
          throw new Error(
            `SBA supported token snapshot has conflicting remote mappings for ${originChainId}:${originToken} -> ${destChainId}`
          )
        }
        remoteTokens[destChainId] = effectiveDestToken
        const originPairLookup = pairLookup[originChainId] || {}
        pairLookup[originChainId] = originPairLookup
        const pairCandidates = originPairLookup[destChainId] || []
        originPairLookup[destChainId] = pairCandidates
        pairCandidates.push({
          originChainId,
          destChainId,
          originToken,
          destToken: effectiveDestToken,
        })
      })

      originLookup[originTokenKey] = remoteTokens
      originEntryCount += 1
    })
  })

  if (tokenIds.size !== EXPECTED_TOKEN_ID_COUNT) {
    throw new Error(
      `SBA supported token snapshot expected ${EXPECTED_TOKEN_ID_COUNT} token IDs but found ${tokenIds.size}`
    )
  }
  if (originEntryCount !== EXPECTED_ORIGIN_ENTRY_COUNT) {
    throw new Error(
      `SBA supported token snapshot expected ${EXPECTED_ORIGIN_ENTRY_COUNT} origin entries but found ${originEntryCount}`
    )
  }

  Object.values(pairLookup).forEach((pairDestinations) => {
    Object.values(pairDestinations ?? {}).forEach((candidates) => {
      candidates?.sort((left, right) =>
        left.originToken
          .toLowerCase()
          .localeCompare(right.originToken.toLowerCase())
      )
    })
  })

  return {
    pairLookup,
    remoteLookup,
  }
}

const {
  pairLookup: SBA_SUPPORTED_TOKEN_PAIR_LOOKUP,
  remoteLookup: SBA_SUPPORTED_TOKEN_REMOTE_LOOKUP,
} = buildSupportedTokenLookups(SBA_SUPPORTED_TOKEN_SNAPSHOT)

export const getSbaSupportedTokens = (
  fromChainId: number,
  toChainId: number,
  toToken?: string
): BridgeTokenCandidate[] => {
  const candidates =
    SBA_SUPPORTED_TOKEN_PAIR_LOOKUP[fromChainId as SupportedChainId]?.[
      toChainId as SupportedChainId
    ] ?? []
  const filteredCandidates = toToken
    ? candidates.filter((candidate) =>
        isSameAddress(candidate.destToken, toToken)
      )
    : candidates
  return filteredCandidates.map((candidate) => ({ ...candidate }))
}

export const getSbaRemoteToken = (
  fromChainId: number,
  originToken: string,
  toChainId: number
): string | undefined => {
  try {
    return SBA_SUPPORTED_TOKEN_REMOTE_LOOKUP[fromChainId as SupportedChainId]?.[
      getAddress(originToken).toLowerCase()
    ]?.[toChainId as SupportedChainId]
  } catch {
    return undefined
  }
}
