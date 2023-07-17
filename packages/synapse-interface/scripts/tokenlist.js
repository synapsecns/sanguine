const { ethers } = require('ethers')
const fs = require('fs')

// Providers
const providers = {
  1: new ethers.providers.JsonRpcProvider('https://eth.llamarpc.com'),
  43114: new ethers.providers.JsonRpcProvider(
    'https://api.avax.network/ext/bc/C/rpc'
  ),
  42161: new ethers.providers.JsonRpcProvider('https://arb1.arbitrum.io/rpc'),
  56: new ethers.providers.JsonRpcProvider(
    'https://bsc-dataseed1.ninicoin.io/'
  ),
  10: new ethers.providers.JsonRpcProvider('https://mainnet.optimism.io'),
  137: new ethers.providers.JsonRpcProvider('https://polygon.llamarpc.com'),
  250: new ethers.providers.JsonRpcProvider('https://rpc3.fantom.network'),
  1285: new ethers.providers.JsonRpcProvider(
    'https://moonriver.public.blastapi.io'
  ),
  288: new ethers.providers.JsonRpcProvider(
    'https://boba-ethereum.gateway.tenderly.co'
  ),
  1284: new ethers.providers.JsonRpcProvider(
    'https://moonbeam.public.blastapi.io'
  ),
  25: new ethers.providers.JsonRpcProvider('https://evm.cronos.org'),
  1088: new ethers.providers.JsonRpcProvider(
    'https://andromeda.metis.io/?owner=1088'
  ),
  1666600000: new ethers.providers.JsonRpcProvider('https://api.s0.t.hmny.io'),
  53935: new ethers.providers.JsonRpcProvider(
    'https://subnets.avax.network/defi-kingdoms/dfk-chain/rpc'
  ),
  1313161554: new ethers.providers.JsonRpcProvider(
    'https://mainnet.aurora.dev'
  ),
  2000: new ethers.providers.JsonRpcProvider('https://rpc.ankr.com/dogechain'),
  7700: new ethers.providers.JsonRpcProvider(
    'https://mainnode.plexnode.org:8545'
  ),
  8217: new ethers.providers.JsonRpcProvider(
    'https://klaytn.api.onfinality.io/public'
  ),

  // Add more providers for the chains you want to support
}

// Contract ABIs
const SynapseRouterABI = require('./abi/SynapseRouter.json')
const SynapseCCTPRouterABI = require('./abi/SynapseCCTPRouter.json')

// Contract addresses
const SynapseRouterAddress = '0x7e7a0e201fd38d3adaa9523da6c109a07118c96a'
const SynapseCCTPRouterAddress = '0xd359bc471554504f683fbd4f6e36848612349ddf'

// Contract instances
const SynapseRouters = {}
const SynapseCCTPRouters = {}

// Chain IDs where SynapseCCTPRouter is allowed
const allowedChainIdsForSynapseCCTPRouter = [1, 43114, 42161]

for (let chainName in providers) {
  SynapseRouters[chainName] = new ethers.Contract(
    SynapseRouterAddress,
    SynapseRouterABI,
    providers[chainName]
  )

  console.log(chainName)
  // Only initialize SynapseCCTPRouters for allowed chains
  if (allowedChainIdsForSynapseCCTPRouter.includes(Number(chainName))) {
    SynapseCCTPRouters[chainName] = new ethers.Contract(
      SynapseCCTPRouterAddress,
      SynapseCCTPRouterABI,
      providers[chainName]
    )
  }
}

// Your token mapping from earlier
const tokens = {
  GOHM: {
    1: '0x0ab87046fBb341D058F17CBC4c1133F25a20a52f',
    10: '0x0b5740c6b4a97f90eF2F0220651Cca420B868FfB',
    56: '0x88918495892BAF4536611E38E75D771Dc6Ec0863',
    137: '0xd8cA34fd379d9ca3C6Ee3b3905678320F5b45195',
    250: '0x91fa20244Fb509e8289CA630E5db3E9166233FDc',
    42161: '0x8D9bA570D6cb60C7e3e0F31343Efe75AB8E65FB1',
    43114: '0x321E7092a180BB43555132ec53AaA65a5bF84251',
    1285: '0x3bF21Ce864e58731B6f28D68d5928BcBEb0Ad172',
    288: '0xd22C0a4Af486C7FA08e282E9eB5f30F9AaA62C95',
    1666600000: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
    1284: '0xD2666441443DAa61492FFe0F37717578714a4521',
    25: '0xbB0A63A6CA2071c6C4bcAC11a1A317b20E3E999C',
    1088: '0xFB21B70922B9f6e3C6274BcD6CB1aa8A0fe20B80',
  },
  LINK: {
    1: '0x514910771af9ca656af840dff83e8264ecf986ca',
    8217: '0xfbed1abb3ad0f8c467068de9fde905887e8c9118',
  },
  HIGHSTREET: {
    1: '0x71Ab77b7dbB4fa7e017BC15090b2163221420282',
    56: '0x5f4bde007dc06b867f86ebfe4802e34a1ffeed63',
  },
  JUMP: {
    56: '0x130025ee738a66e691e6a7a62381cb33c6d9ae83',
    250: '0x78DE9326792ce1d6eCA0c978753c6953Cdeedd73',
    1088: '0xE3c82A836Ec85311a433fBd9486EfAF4b1AFbF48',
  },
  SFI: {
    1: '0xb753428af26e81097e7fd17f40c88aaa3e04902c',
    43114: '0xc2Bf0A1f7D8Da50D608bc96CF701110d4A438312',
  },
  DOG: {
    1: '0xBAac2B4491727D78D2b78815144570b9f2Fe8899',
    56: '0xaa88c603d142c371ea0eac8756123c5805edee03',
    137: '0xeEe3371B89FC43Ea970E908536Fcddd975135D8a',
    42161: '0x4425742F1EC8D98779690b5A3A6276Db85Ddc01A',
  },
  NFD: {
    56: '0x0fe9778c005a5a6115cbe12b0568a2d50b765a51',
    43114: '0xf1293574ee43950e7a8c9f1005ff097a9a713959',
    2000: '0x868055ADFA27D331d5b69b1BF3469aDAAc3ba7f2',
    137: '0x0a5926027d407222f8fe20f24cb16e103f617046',
  },
  SOLAR: {
    1284: '0x0DB6729C03C85B0708166cA92801BcB5CAc781fC',
    1285: '0x76906411D07815491A5E577022757aD941fb5066',
  },
  GMX: {
    42161: '0xfc5a1a6eb076a2c7ad06ed22c90d7e710e35ad0a',
    43114: '0x62edc0692bd897d2295872a9ffcac5425011c661',
  },
  SDT: {
    1: '0x73968b9a57c6e53d41345fd57a6e6ae27d6cdb2f',
    43114: '0xCCBf7c451F81752F7d2237F2c18C371E6e089E69',
    42161: '0x087d18A77465c34CDFd3a081a2504b7E86CE4EF8',
    250: '0xE3c82A836Ec85311a433fBd9486EfAF4b1AFbF48',
    1666600000: '0xE3c82A836Ec85311a433fBd9486EfAF4b1AFbF48',
  },
  NEWO: {
    1: '0x98585dFc8d9e7D48F0b1aE47ce33332CF4237D96',
    43114: '0x4Bfc90322dD638F81F034517359BD447f8E0235a',
    42161: '0x0877154a755B24D499B8e2bD7ecD54d3c92BA433',
  },
  USDB: {
    1: '0x02b5453d92b730f29a86a0d5ef6e930c4cf8860b',
    56: '0xc8699abbba90c7479dedccef19ef78969a2fc608',
    137: '0xfa1fbb8ef55a4855e5688c0ee13ac3f202486286',
    250: '0x6fc9383486c163fa48becdec79d6058f984f62ca',
    43114: '0x5ab7084cb9d270c2cb052dd30dbecbca42f8620c',
    1285: '0x3e193c39626bafb41ebe8bdd11ec7cca9b3ec0b2',
  },
  PEPE: {
    1: '0x6982508145454ce325ddbe47a25d4ec3d2311933',
    42161: '0xA54B8e178A49F8e5405A4d44Bb31F496e5564A05',
  },
  VSTA: {
    1: '0xA8d7F5e7C78ed0Fa097Cc5Ec66C1DC3104c9bbeb',
    42161: '0xa684cd057951541187f288294a1e1c2646aa2d24',
  },
  H2O: {
    1: '0x0642026e7f0b6ccac5925b4e7fa61384250e1701',
    42161: '0xD1c6f989e9552DB523aBAE2378227fBb059a3976',
    43114: '0xC6b11a4Fd833d1117E9D312c02865647cd961107',
    56: '0x03eFca7CEb108734D3777684F3C0A0d8ad652f79',
    1284: '0xA46aDF6D5881ca0b8596EDadF8f058F8c16d8B68',
    1285: '0x9c0a820bb01e2807aCcd1c682d359e92DDd41403',
    10: '0xE3c82A836Ec85311a433fBd9486EfAF4b1AFbF48',
    137: '0x32ba7cF7d681357529013de6a2CDF93933C0dF3f',
  },
  L2DAO: {
    42161: '0x2CaB3abfC1670D1a452dF502e216a66883cDf079',
    10: '0xd52f94DF742a6F4B4C8b033369fE13A41782Bf44',
  },
  PLS: {
    42161: '0x51318b7d00db7acc4026c88c3952b66278b6a67f',
    10: '0xD9eAA386cCD65F30b77FF175F6b52115FE454fD6',
  },
  AGEUR: {
    1: '0x1a7e4e63778B4f12a199C062f3eFdD288afCBce8',
    42161: '0x16BFc5fe024980124bEf51d1D792dC539d1B5Bf0',
    10: '0xa0554607e477cdC9d0EE2A6b087F4b2DC2815C22',
  },
  UNIDX: {
    1: '0x95b3497bbcccc46a8f45f5cf54b0878b39f8d96c',
    42161: '0x5429706887FCb58a595677B73E9B0441C25d993D',
  },
  BUSD: {
    56: '0xe9e7cea3dedca5984780bafc599bd69add087d56',
    2000: '0x1555C68Be3b22cdcCa934Ae88Cb929Db40aB311d',
  },
  USDC: {
    56: '0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d',
    1: '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
    25: '0xc21223249ca28397b4b6541dffaecc539bff0c59',
    10: '0x7f5c764cbc14f9669b88837ca1490cca17c31607',
    137: '0x2791bca1f2de4661ed88a30c99a7a9449aa84174',
    250: '0x04068da6c83afcfa0e13ba15a6696662335d5b75',
    43114: '0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e',
    42161: '0xaf88d065e77c8cc2239327c5edb3a432268e5831',
    1666600000: '0x985458e523db3d53125813ed68c274899e9dfab4',
    288: '0x66a2A913e447d6b4BF33EFbec43aAeF87890FBbc',
    1313161554: '0xB12BFcA5A55806AaF64E99521918A4bf0fC40802',
    1088: '0xEA32A96608495e54156Ae48931A7c20f0dcc1a21',
    7700: '0x80b5a32E4F032B2a058b4F29EC95EEfEEB87aDcd',
    8217: '0x6270B58BE569a7c0b8f47594F191631Ae5b2C86C',
    2000: '0x85C2D3bEBffD83025910985389aB8aD655aBC946',
  },
  KLAYTN_USDC: {
    1: '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
    8217: '0x6270B58BE569a7c0b8f47594F191631Ae5b2C86C',
    2000: '0x85C2D3bEBffD83025910985389aB8aD655aBC946',
  },
  KLAYTN_USDT: {
    1: '0xdac17f958d2ee523a2206206994597c13d831ec7',
    8217: '0xd6dAb4CfF47dF175349e6e7eE2BF7c40Bb8C05A3',
    2000: '0x7f8e71DD5A7e445725F0EF94c7F01806299e877A',
  },
  KLAYTN_oUSDT: { 8217: '0xceE8FAF64bB97a73bb51E115Aa89C17FfA8dD167' },
  KLAYTN_DAI: {
    1: '0x6b175474e89094c44da98b954eedeac495271d0f',
    8217: '0x078dB7827a5531359f6CB63f62CFA20183c4F10c',
    2000: '0xB3306f03595490e5cC3a1b1704a5a158D3436ffC',
  },
  DOGECHAIN_BUSD: {
    56: '0xe9e7cea3dedca5984780bafc599bd69add087d56',
    2000: '0x1555C68Be3b22cdcCa934Ae88Cb929Db40aB311d',
  },
  USDT: {
    56: '0x55d398326f99059ff775485246999027b3197955',
    1: '0xdac17f958d2ee523a2206206994597c13d831ec7',
    25: '0x66e428c3f67a68878562e79a0234c1f83c208770',
    137: '0xc2132d05d31c914a87c6611c10748aeb04b58e8f',
    43114: '0x9702230a8ea53601f5cd2dc00fdbc13d4df4a8c7',
    0: '0x9A9f2CCfdE556A7E9Ff0848998Aa4a0CFD8863AE',
    42161: '0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9',
    250: '0x049d68029688eabf473097a2fc38ef61633a3c7a',
    1666600000: '0x3c2b8be99c50593081eaa2a724f0b8285f5aba8f',
    288: '0x5DE1677344D3Cb0D7D465c10b72A8f60699C062d',
    1313161554: '0x4988a896b1227218e4A686fdE5EabdcAbd91571f',
    7700: '0xd567B3d7B8FE3C79a1AD8dA978812cfC4Fa05e75',
    8217: '0xd6dAb4CfF47dF175349e6e7eE2BF7c40Bb8C05A3',
    2000: '0x7f8e71DD5A7e445725F0EF94c7F01806299e877A',
  },
  DAI: {
    56: '0x1af3f329e8be154074d8769d1ffa4ee058b1dbc3',
    1: '0x6b175474e89094c44da98b954eedeac495271d0f',
    25: '0xf2001b145b43032aaf5ee2884e456ccd805f677d',
    137: '0x8f3cf7ad23cd3cadbd9735aff958023239c6a063',
    43114: '0xd586E7F844cEa2F87f50152665BCbc2C279D8d70',
    42161: '0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1',
    1666600000: '0xef977d2f931c1978db5f6747666fa1eacb0d0339',
    288: '0xf74195Bb8a5cf652411867c5C2C5b8C2a402be35',
    8217: '0x078dB7827a5531359f6CB63f62CFA20183c4F10c',
    2000: '0xB3306f03595490e5cC3a1b1704a5a158D3436ffC',
  },
  WBTC: {
    1: '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599',
    8217: '0xDCbacF3f7a069922E677912998c8d57423C37dfA',
    2000: '0xD0c6179c43C00221915f1a61f8eC06A5Aa32F9EC',
  },
  MIM: {
    250: '0x82f0b8b456c1a451378467398982d4834b6829c1',
    42161: '0xfea7a6a0b346362bf88a9e4a88416b77a57d6c2a',
  },
  WETHE: { 43114: '0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab' },
  WETHBEAM: { 1284: '0x3192Ae73315c3634Ffa217f71CF6CBc30FeE349A' },
  AVWETH: { 43114: '0x53f7c5869a859f0aec3d334ee8b4cf01e3492f21' },
  ONEETH: { 1666600000: '0x6983d1e6def3690c4d616b13597a09e6193ea013' },
  FTMETH: { 250: '0x74b23882a30290451A17c44f4F05243b6b58C76d' },
  CANTOETH: { 7700: '0x5FD55A1B9FC24967C4dB09C513C3BA0DFa7FF687' },
  METISETH: { 1088: '0x420000000000000000000000000000000000000A' },
  SYN: {
    1: '0x0f2d719407fdbeff09d87557abb7232601fd9f29',
    56: '0xa4080f1778e69467e905b8d6f72f6e441f9e9484',
    137: '0xf8f9efc0db77d8881500bb06ff5d6abc3070e695',
    250: '0xE55e19Fb4F2D85af758950957714292DAC1e25B2',
    42161: '0x080f6aed32fc474dd5717105dba5ea57268f46eb',
    43114: '0x1f1E7c893855525b303f99bDF5c3c05Be09ca251',
    1666600000: '0xE55e19Fb4F2D85af758950957714292DAC1e25B2',
    288: '0xb554A55358fF0382Fb21F0a478C3546d1106Be8c',
    1088: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
    1285: '0xd80d8688b02B3FD3afb81cDb124F188BB5aD0445',
    1284: '0xF44938b0125A6662f9536281aD2CD6c499F22004',
    10: '0x5A5fFf6F753d7C11A56A52FE47a177a87e431655',
    25: '0xFD0F80899983b8D46152aa1717D76cba71a31616',
    1313161554: '0xd80d8688b02B3FD3afb81cDb124F188BB5aD0445',
    2000: '0xDfA53EeBA61D69E1D2b56b36d78449368F0265c1',
    7700: '0x555982d2E211745b96736665e19D9308B615F78e',
  },
  FRAX: {
    1: '0x853d955acef822db058eb8505911ed77f175b99e',
    1285: '0xE96AC70907ffF3Efee79f502C985A7A21Bce407d',
    1666600000: '0xFa7191D292d5633f702B0bd7E3E3BcCC0e633200',
    2000: '0x10D70831f9C3c11c5fe683b2f1Be334503880DB6',
  },
  SYN_FRAX: {
    1285: '0xE96AC70907ffF3Efee79f502C985A7A21Bce407d',
    1284: '0xDd47A348AB60c61Ad6B60cA8C31ea5e00eBfAB4F',
    1666600000: '0x1852F70512298d56e9c8FDd905e02581E04ddb2a',
  },
  NUSD: {
    56: '0x23b891e5c62e0955ae2bd185990103928ab817b3',
    1: '0x1B84765dE8B7566e4cEAF4D0fD3c5aF52D3DdE4F',
    25: '0x396c9c192dd323995346632581BEF92a31AC623b',
    10: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
    137: '0xb6c473756050de474286bed418b77aeac39b02af',
    250: '0xED2a7edd7413021d440b09D654f3b87712abAB66',
    43114: '0xCFc37A6AB183dd4aED08C204D1c2773c0b1BDf46',
    42161: '0x2913E812Cf0dcCA30FB28E6Cac3d2DCFF4497688',
    1666600000: '0xED2a7edd7413021d440b09D654f3b87712abAB66',
    288: '0x6B4712AE9797C199edd44F897cA09BC57628a1CF',
    1313161554: '0x07379565cD8B0CaE7c60Dc78e7f601b34AF2A21c',
    1088: '0x961318Fc85475E125B99Cc9215f62679aE5200aB',
    53935: '0x3AD9DFE640E1A9Cc1D9B0948620820D975c3803a',
    7700: '0xD8836aF2e565D3Befce7D906Af63ee45a57E8f80',
  },
  NOTE: { 7700: '0x4e71a2e537b7f9d9413d3991d37958c0b5e1e503' },
  DFK_USDC: {},
  NETH: {
    250: '0x67C10C397dD0Ba417329543c1a40eb48AAa7cd00',
    42161: '0x3ea9B0ab55F34Fb188824Ee288CeaEfC63cf908e',
    288: '0x96419929d7949D6A801A6909c145C8EEf6A40431',
    10: '0x809DC529f07651bD43A172e8dB6f4a7a0d771036',
    43114: '0x19E1ae0eE35c0404f835521146206595d37981ae',
    1666600000: '0x0b5740c6b4a97f90eF2F0220651Cca420B868FfB',
    1284: '0x3192Ae73315c3634Ffa217f71CF6CBc30FeE349A',
    1088: '0x931B8f17764362A3325D30681009f0eDd6211231',
    8217: '0xCD6f29dC9Ca217d0973d3D21bF58eDd3CA871a86',
    2000: '0x9F4614E4Ea4A0D7c4B1F946057eC030beE416cbB',
    7700: '0x09fEC30669d63A13c666d2129230dD5588E2e240',
  },
  KLAYTN_WETH: {},
  1: {},
  MOVR: {},
  AVAX: {},
  WMOVR: {
    1284: '0x1d4C2a246311bB9f827F4C768e277FF5787B7D7E',
    1285: '0x98878b06940ae243284ca214f92bb71a2b032b8a',
  },
  WAVAX: {
    43114: '0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7',
    53935: '0xB57B60DeBDB0b8172bb6316a9164bd3C695F133a',
    1284: '0xA1f8890E39b4d8E33efe296D698fe42Fb5e59cC3',
    1666600000: '0xD9eAA386cCD65F30b77FF175F6b52115FE454fD6',
  },
  SYNAVAX: { 1666600000: '0xD9eAA386cCD65F30b77FF175F6b52115FE454fD6' },
  MULTIAVAX: { 1666600000: '0xb12c13e66ade1f72f71834f2fc5082db8c091358' },
  JEWEL: {
    1666600000: '0x72cb10c6bfa5624dd07ef608027e366bd690048f',
    8217: '0x30C103f8f5A3A732DFe2dCE1Cc9446f545527b43',
  },
  WJEWEL: { 53935: '0xCCb93dABD71c8Dad03Fc4CE5559dC3D89F67a260' },
  SYNJEWEL: {
    43114: '0x997Ddaa07d716995DE90577C123Db411584E5E46',
    1666600000: '0x28b42698Caf46B4B012CF38b6C75867E0762186D',
  },
  XJEWEL: {
    53935: '0x77f2656d04E158f915bC22f07B779D94c1DC47Ff',
    1666600000: '0xA9cE83507D872C5e1273E745aBcfDa849DAA654F',
  },
  USDCe: {
    43114: '0xA7D7079b0FEaD91F3e65f86E8915Cb59c1a4C664',
    42161: '0xff970a61a04b1ca14834a43f5de4533ebddb5cc8',
  },
  USDTe: { 43114: '0xc7198437980c041c805a1edcba50c1ce5db95118' },
  WMATIC: { 137: '0x9b17bAADf0f21F03e35249e0e59723F34994F806' },
  WBNB: {
    56: '0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c',
    2000: '0x1fC532187B4848d2F9c564531b776A4F8e11201d',
  },
  DEPRECATED_WKLAY: { 8217: '0x5819b6af194a78511c79c85ea68d2377a7e9335f' },
}

const tokenMapping = {}

let connectedTokensList = {}

async function getConnectedTokens() {
  for (let tokenName in tokens) {
    for (let chainId in tokens[tokenName]) {
      const tokenAddress = tokens[tokenName][chainId]

      if (!tokenAddress) {
        console.error(
          `Token address for token ${tokenName} on chain ${chainId} does not exist`
        )
        continue
      }

      let connectedTokensSynapse = []
      let connectedTokensCCTP = []

      if (SynapseRouters[chainId]) {
        try {
          connectedTokensSynapse = await SynapseRouters[
            chainId
          ].getConnectedBridgeTokens(tokenAddress)
        } catch (error) {
          console.error(
            `Failed to get connectedTokensSynapse for token ${tokenName} on chain ${chainId}`,
            error
          )
        }
      } else {
        console.error(`SynapseRouters for chain ${chainId} is not defined`)
      }
      if (SynapseCCTPRouters[chainId]) {
        try {
          connectedTokensCCTP = await SynapseCCTPRouters[
            chainId
          ].getConnectedBridgeTokens(tokenAddress)
        } catch (error) {
          console.error(
            `Failed to get connectedTokensCCTP for token ${tokenName} on chain ${chainId}`,
            error
          )
        }
      } else {
        console.error(`SynapseCCTPRouters for chain ${chainId} is not defined`)
      }

      // Extract the symbols and merge the lists
      const connectedTokensSymbols = [
        ...connectedTokensSynapse,
        ...connectedTokensCCTP,
      ].map((item) => item[0])

      // Check if tokenName already exists in connectedTokensList
      if (!connectedTokensList[tokenName]) {
        // If not, create a new entry
        connectedTokensList[tokenName] = {}
      }

      // Add the new chain data
      connectedTokensList[tokenName][chainId] = connectedTokensSymbols
    }
  }
  console.log(JSON.stringify(connectedTokensList, null, 2))
}

getConnectedTokens()
