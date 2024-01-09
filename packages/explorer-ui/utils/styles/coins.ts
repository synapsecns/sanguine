import {
  BUSD,
  USDT,
  USDC,
  DAI,
  NUSD,
  NETH,
  SYN,
  ETH,
  FRAX,
  WETH,
  WETHE,
  ONEETH,
  WMOVR,
  MOVR,
  AVAX,
  WAVAX,
  JEWEL,
  WJEWEL,
  SYNJEWEL,
  XJEWEL,
  WBTC,
  DOG,
  GOHM,
  HIGH,
  LINK,
  JUMP,
  NFD,
  SOLAR,
  NEWO,
  VSTA,
  SDT,
  SFI,
  H2O,
  L2DAO,
  PLS,
  AGEUR,
  NOTE,
  // MIM,
  // AVWETH,
} from '@synapse-constants'

const COLOR_COIN_MAP = {
  gray: [FRAX, GOHM, SDT, VSTA],
  yellow: [BUSD, DOG, NFD, NEWO, AGEUR, WBTC],
  green: [USDT, PLS, NOTE],
  lime: [JEWEL, WJEWEL, XJEWEL, SYNJEWEL],
  sky: [ETH, WETH, WETHE, ONEETH],
  blue: [USDC, LINK],
  orange: [DAI, SOLAR],
  purple: [NUSD, NETH, SYN, WMOVR, MOVR],
  // need to figure out imports for the following
  // indigo: [MIM],
  // cyan: [HIGH, JUMP, AVWETH, H2O, L2DAO],
  cyan: [HIGH, JUMP, H2O, L2DAO],
  red: [AVAX, WAVAX, SFI],
}

const COIN_COLORS = {}

for (const [colorName, coinArr] of Object.entries(COLOR_COIN_MAP)) {
  for (const someCoin of coinArr) {
    COIN_COLORS[someCoin.symbol] = colorName
  }
}

const SHARED_OPACITY_OVERRIDE = `
  dark:hover:bg-opacity-20
  dark:focus:bg-opacity-20
  dark:active:bg-opacity-20
  `

/**
 * @param {Token} coin
 */
export function getButtonStyleForCoin(coin) {
  switch (COIN_COLORS[coin.symbol]) {
    case 'yellow':
      return 'shadow-yellow-xl border-yellow-100 dark:border-opacity-50 dark:border-yellow-700'
    case 'green':
      return 'shadow-green-xl border-green-100 dark:border-opacity-50 dark:border-green-700'
    case 'lime':
      return 'shadow-lime-xl border-lime-100 dark:border-opacity-50 dark:border-lime-700'
    case 'sky':
      return 'shadow-sky-xl border-sky-100 dark:border-opacity-50 dark:border-sky-700'
    case 'blue':
      return 'shadow-blue-xl border-blue-100 dark:border-opacity-50 dark:border-blue-700'
    case 'orange':
      return 'shadow-orange-xl border-orange-100 dark:border-opacity-50 dark:border-orange-700'
    case 'purple':
      return 'shadow-purple-xl border-purple-100 dark:border-opacity-50 dark:border-purple-700'
    case 'indigo':
      return 'shadow-indigo-xl border-indigo-100 dark:border-opacity-50 dark:border-indigo-700'
    case 'cyan':
      return 'shadow-cyan-xl border-cyan-100 dark:border-opacity-50 dark:border-cyan-700'
    case 'red':
      return 'shadow-red-xl border-red-100 dark:border-opacity-50 dark:border-red-700'
    default:
      return 'shadow-lg border-gray-300'
  }
}
/**
 * @param {Token} coin
 */
export function getMenuItemStyleForCoin(coin) {
  switch (COIN_COLORS[coin.symbol]) {
    case 'yellow':
      return 'hover:bg-yellow-50 focus:bg-yellow-50 active:bg-yellow-50'
    case 'green':
      return 'hover:bg-green-50 focus:bg-green-50 active:bg-green-50'
    case 'lime':
      return 'hover:bg-lime-50 focus:bg-lime-50 active:bg-lime-50'
    case 'sky':
      return 'hover:bg-sky-50 focus:bg-sky-50 active:bg-sky-50'
    case 'blue':
      return 'hover:bg-blue-50 focus:bg-blue-50 active:bg-blue-50'
    case 'orange':
      return 'hover:bg-orange-50 focus:bg-orange-50 active:bg-orange-50'
    case 'purple':
      return 'hover:bg-purple-50 focus:bg-purple-50 active:bg-purple-50'
    case 'indigo':
      return 'hover:bg-indigo-50 focus:bg-indigo-50 active:bg-indigo-50'
    case 'cyan':
      return 'hover:bg-cyan-50 focus:bg-cyan-50 active:bg-cyan-50'
    case 'red':
      return 'hover:bg-red-50 focus:bg-red-50 active:bg-red-50'
    default:
      return 'hover:bg-gray-200 focus:bg-gray-200 active:bg-gray-200'
  }
}

/**
 * @param {Token} coin
 */
export function getMenuItemStyleForCoinDark(coin) {
  switch (COIN_COLORS[coin.symbol]) {
    case 'yellow':
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-yellow-500 dark:focus:bg-yellow-500 dark:active:bg-yellow-500`
    case 'green':
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-green-500 dark:focus:bg-green-500 dark:active:bg-green-500`
    case 'lime':
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-lime-500 dark:focus:bg-lime-500 dark:active:bg-lime-500`
    case 'sky':
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-sky-500 dark:focus:bg-sky-500 dark:active:bg-sky-500`
    case 'blue':
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-blue-500 dark:focus:bg-blue-500 dark:active:bg-blue-500`
    case 'orange':
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-orange-500 dark:focus:bg-orange-500 dark:active:bg-orange-500`
    case 'purple':
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-purple-500 dark:focus:bg-purple-500 dark:active:bg-purple-500`
    case 'indigo':
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-indigo-500 dark:focus:bg-indigo-500 dark:active:bg-indigo-500`
    case 'cyan':
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-cyan-500 dark:focus:bg-cyan-500 dark:active:bg-cyan-500`
    case 'red':
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-red-500 dark:focus:bg-red-500 dark:active:bg-red-500`
    default:
      return `${SHARED_OPACITY_OVERRIDE} dark:hover:bg-gray-700 dark:focus:bg-gray-700 dark:active:bg-gray-700`
  }
}

export function getMenuItemStyleForCoinCombined(coin) {
  return `${getMenuItemStyleForCoin(coin)} ${getMenuItemStyleForCoinDark(coin)}`
}

/**
 * @param {Token} coin
 */
export function getSwapHoverStyleForCoin(coin) {
  switch (COIN_COLORS[coin.symbol]) {
    case 'yellow':
      return 'hover:shadow-yellow-xl border-yellow-100'
    case 'green':
      return 'hover:shadow-green-xl border-green-100'
    case 'lime':
      return 'hover:shadow-lime-xl border-lime-100'
    case 'sky':
      return 'hover:shadow-sky-xl border-sky-100'
    case 'blue':
      return 'hover:shadow-blue-xl border-blue-100'
    case 'orange':
      return 'hover:shadow-orange-xl border-orange-100'
    case 'purple':
      return 'hover:shadow-purple-xl border-purple-100'
    case 'indigo':
      return 'hover:shadow-indigo-xl border-indigo-100'
    case 'cyan':
      return 'hover:shadow-cyan-xl border-cyan-100'
    case 'red':
      return 'hover:shadow-red-xl border-red-100'
    default:
      return 'hover:shadow-lg border-gray-300'
  }
}

/**
 * @param {Token} coin
 */
export function getMenuItemBgForCoin(coin) {
  switch (COIN_COLORS[coin.symbol]) {
    case 'yellow':
      return 'bg-yellow-50 dark:bg-opacity-20 dark:bg-yellow-700'
    case 'green':
      return 'bg-green-50 dark:bg-opacity-20 dark:bg-green-700'
    case 'lime':
      return 'bg-lime-50 dark:bg-opacity-20 dark:bg-lime-700'
    case 'sky':
      return 'bg-sky-50 dark:bg-opacity-20 dark:bg-sky-700'
    case 'blue':
      return 'bg-blue-50 dark:bg-opacity-20 dark:bg-blue-700'
    case 'orange':
      return 'bg-orange-50 dark:bg-opacity-20 dark:bg-orange-700'
    case 'purple':
      return 'bg-purple-50 dark:bg-opacity-20 dark:bg-purple-700'
    case 'indigo':
      return 'bg-indigo-50 dark:bg-opacity-20 dark:bg-indigo-700'
    case 'cyan':
      return 'bg-cyan-50 dark:bg-opacity-20 dark:bg-cyan-700'
    case 'red':
      return 'bg-red-50 dark:bg-opacity-20 dark:bg-red-700'
    default:
      return 'bg-gray-200 dark:bg-opacity-20 dark:bg-gray-700'
  }
}
/**
 * @param {Token} coin
 */
export function getMenuItemHoverBgForCoin(coin) {
  switch (COIN_COLORS[coin.symbol]) {
    case 'yellow':
      return 'hover:bg-yellow-100 dark:hover:bg-opacity-20 dark:hover:bg-yellow-700'
    case 'green':
      return 'hover:bg-green-100 dark:hover:bg-opacity-20 dark:hover:bg-green-700'
    case 'lime':
      return 'hover:bg-lime-100 dark:hover:bg-opacity-20 dark:hover:bg-lime-700'
    case 'sky':
      return 'hover:bg-sky-100 dark:hover:bg-opacity-20 dark:hover:bg-sky-700'
    case 'blue':
      return 'hover:bg-blue-100 dark:hover:bg-opacity-20 dark:hover:bg-blue-700'
    case 'orange':
      return 'hover:bg-orange-100 dark:hover:bg-opacity-20 dark:hover:bg-orange-700'
    case 'purple':
      return 'hover:bg-purple-100 dark:hover:bg-opacity-20 dark:hover:bg-purple-700'
    case 'indigo':
      return 'hover:bg-indigo-100 dark:hover:bg-opacity-20 dark:hover:bg-indigo-700'
    case 'cyan':
      return 'hover:bg-cyan-100 dark:hover:bg-opacity-20 dark:hover:bg-cyan-700'
    case 'red':
      return 'hover:bg-red-100 dark:hover:bg-opacity-20 dark:hover:bg-red-700'
    default:
      return 'hover:bg-gray-300 dark:hover:bg-opacity-20 '
  }
}
/**
 * @param {Token} coin
 */
export function getCoinTextColor(coin) {
  switch (COIN_COLORS[coin.symbol]) {
    case 'yellow':
      return 'text-yellow-500 group-hover:text-yellow-400'
    case 'green':
      return 'text-green-600 group-hover:text-green-500'
    case 'lime':
      return 'text-lime-600 group-hover:text-lime-500'
    case 'sky':
      return 'text-sky-600 group-hover:text-sky-500'
    case 'blue':
      return 'text-blue-600 group-hover:text-blue-500'
    case 'orange':
      return 'text-orange-600 group-hover:text-orange-500'
    case 'purple':
      return 'text-purple-600 group-hover:text-purple-500'
    case 'indigo':
      return 'text-indigo-600 group-hover:text-indigo-500'
    case 'cyan':
      return 'text-cyan-600 group-hover:text-cyan-500'
    case 'red':
      return 'text-red-600 group-hover:text-red-500'
    case 'gray':
      return 'dark:text-gray-400 dark:group-hover:text-gray-300'
    default:
      return 'text-indigo-600 group-hover:text-indigo-500'
  }
}

export function getCoinTextColorDark(coin) {
  switch (COIN_COLORS[coin.symbol]) {
    case 'yellow':
      return 'dark:text-yellow-500 dark:group-hover:text-yellow-400'
    case 'green':
      return 'dark:text-green-600 dark:group-hover:text-green-500'
    case 'lime':
      return 'dark:text-lime-600 dark:group-hover:text-lime-500'
    case 'sky':
      return 'dark:text-sky-600 dark:group-hover:text-sky-500'
    case 'blue':
      return 'dark:text-blue-600 dark:group-hover:text-blue-500'
    case 'orange':
      return 'dark:text-orange-600 dark:group-hover:text-orange-500'
    case 'purple':
      return 'dark:text-purple-600 dark:group-hover:text-purple-500'
    case 'indigo':
      return 'dark:text-indigo-600 dark:group-hover:text-indigo-500'
    case 'cyan':
      return 'dark:text-cyan-600 dark:group-hover:text-cyan-500'
    case 'red':
      return 'dark:text-red-600 dark:group-hover:text-red-500'
    case 'gray':
      return 'dark:text-gray-400 dark:group-hover:text-gray-300'
    default:
      return 'dark:text-indigo-600 dark:group-hover:text-indigo-500'
  }
}

export function getCoinTextColorAlt(coin) {
  switch (COIN_COLORS[coin.symbol]) {
    case 'yellow':
      return 'dark:text-yellow-500 dark:group-hover:text-yellow-400'
    case 'green':
      return 'dark:text-green-500 dark:group-hover:text-green-400'
    case 'lime':
      return 'dark:text-lime-500 dark:group-hover:text-lime-400'
    case 'sky':
      return 'dark:text-sky-500 dark:group-hover:text-sky-400'
    case 'blue':
      return 'dark:text-blue-500 dark:group-hover:text-blue-400'
    case 'orange':
      return 'dark:text-orange-500 dark:group-hover:text-orange-400'
    case 'purple':
      return 'dark:text-purple-500 dark:group-hover:text-purple-400'
    case 'indigo':
      return 'dark:text-indigo-500 dark:group-hover:text-indigo-400'
    case 'cyan':
      return 'dark:text-cyan-500 dark:group-hover:text-cyan-400'
    case 'red':
      return 'dark:text-red-500 dark:group-hover:text-red-400'
    case 'gray':
      return 'dark:text-gray-400 dark:group-hover:text-gray-300'
    default:
      return 'dark:text-indigo-600 dark:group-hover:text-indigo-500'
  }
}

export function getCoinTextColorCombined(coin) {
  return `${getCoinTextColor(coin)} ${getCoinTextColorDark(coin)}`
}

/**
 * @param {Token} coin
 */
export function getInputBorderFocusStyleForCoin(coin) {
  switch (COIN_COLORS[coin.symbol]) {
    case 'yellow':
      return 'focus-within:border-yellow-200 dark:focus-within:border-yellow-500'
    case 'green':
      return 'focus-within:border-green-200 dark:focus-within:border-green-500'
    case 'lime':
      return 'focus-within:border-lime-200 dark:focus-within:border-lime-500'
    case 'sky':
      return 'focus-within:border-sky-200 dark:focus-within:border-sky-500'
    case 'blue':
      return 'focus-within:border-blue-200 dark:focus-within:border-blue-500'
    case 'orange':
      return 'focus-within:border-orange-200 dark:focus-within:border-orange-500'
    case 'purple':
      return 'focus-within:border-purple-200 dark:focus-within:border-purple-500'
    case 'indigo':
      return 'focus-within:border-indigo-200 dark:focus-within:border-indigo-500'
    case 'cyan':
      return 'focus-within:border-cyan-200 dark:focus-within:border-cyan-500'
    case 'red':
      return 'focus-within:border-red-200 dark:focus-within:border-red-500'
    default:
      return 'focus-within:border-gray-300 dark:focus-within:border-gray-500'
  }
}
/**
 * @param {Token} coin
 */
export function getSwapBorderStyleForCoin(coin) {
  switch (COIN_COLORS[coin.symbol]) {
    case 'yellow':
      return 'border-yellow-50 dark:border-opacity-20 dark:border-yellow-700 '
    case 'green':
      return 'border-green-50 dark:border-opacity-20 dark:border-green-700'
    case 'lime':
      return 'border-lime-50 dark:border-opacity-20 dark:border-lime-700'
    case 'sky':
      return 'border-sky-50 dark:border-opacity-20 dark:border-sky-700'
    case 'blue':
      return 'border-blue-50 dark:border-opacity-20 dark:border-blue-700'
    case 'orange':
      return 'border-orange-50 dark:border-opacity-20 dark:border-orange-700'
    case 'purple':
      return 'border-purple-50 dark:border-opacity-20 dark:border-purple-700'
    case 'indigo':
      return 'border-indigo-50 dark:border-opacity-20 dark:border-indigo-700'
    case 'cyan':
      return 'border-cyan-50 dark:border-opacity-20 dark:border-cyan-700'
    case 'red':
      return 'border-red-50 dark:border-opacity-20 dark:border-red-700'
    default:
      return 'border-gray-200 dark:border-opacity-20 dark:border-gray-700'
  }
}
/**
 * @param {Token} coin
 */
export function getSwapBorderHoverStyleForCoin(coin) {
  switch (COIN_COLORS[coin?.symbol]) {
    case 'yellow':
      return 'hover:border-yellow-100 dark:hover:border-opacity-50 dark:hover:!border-yellow-700'
    case 'green':
      return 'hover:border-green-100 dark:hover:border-opacity-50 dark:hover:!border-green-700'
    case 'lime':
      return 'hover:border-lime-100 dark:hover:border-opacity-50 dark:hover:!border-lime-700'
    case 'sky':
      return 'hover:border-sky-100 dark:hover:border-opacity-50 dark:hover:!border-sky-700'
    case 'blue':
      return 'hover:border-blue-100 dark:hover:border-opacity-50 dark:hover:!border-blue-700'
    case 'orange':
      return 'hover:border-orange-100 dark:hover:border-opacity-50 dark:hover:!border-orange-700'
    case 'purple':
      return 'hover:border-purple-100 dark:hover:border-opacity-50 dark:hover:!border-purple-700'
    case 'indigo':
      return 'hover:border-indigo-100 dark:hover:border-opacity-50 dark:hover:!border-indigo-700'
    case 'cyan':
      return 'hover:border-cyan-100 dark:hover:border-opacity-50 dark:hover:!border-cyan-700'
    case 'red':
      return 'hover:border-red-100 dark:hover:border-opacity-50 dark:hover:!border-red-700'
    default:
      return 'hover:border-gray-300 dark:hover:border-opacity-50 dark:hover:border-gray-700'
  }
}
/**
 * @param {Token} coin
 */
export function getBorderStyleForCoin(coin) {
  switch (COIN_COLORS[coin.symbol]) {
    case 'yellow':
      return 'border-yellow-300'
    case 'green':
      return 'border-green-300'
    case 'lime':
      return 'border-lime-300'
    case 'sky':
      return 'border-sky-300'
    case 'blue':
      return 'border-blue-300'
    case 'orange':
      return 'border-orange-300'
    case 'purple':
      return 'border-purple-300'
    case 'indigo':
      return 'border-indigo-300'
    case 'cyan':
      return 'border-cyan-300'
    case 'red':
      return 'border-red-300'
    default:
      return 'border-gray-200'
  }
}
/**
 * @param {Token} coin
 */
export function getFromStyleForCoin(coin) {
  switch (COIN_COLORS[coin.symbol]) {
    case 'yellow':
      return 'from-yellow-100'
    case 'green':
      return 'from-green-100'
    case 'lime':
      return 'from-lime-100'
    case 'sky':
      return 'from-sky-100'
    case 'blue':
      return 'from-blue-100'
    case 'orange':
      return 'from-orange-100'
    case 'purple':
      return 'from-purple-100'
    case 'indigo':
      return 'from-indigo-100'
    case 'cyan':
      return 'from-cyan-100'
    case 'red':
      return 'from-red-100'
    default:
      return 'from-gray-100'
  }
}
/**
 * @param {Token} coin
 */
export function getToStyleForCoin(coin) {
  switch (COIN_COLORS[coin.symbol]) {
    case 'yellow':
      return 'to-yellow-100'
    case 'green':
      return 'to-green-100'
    case 'lime':
      return 'to-lime-100'
    case 'sky':
      return 'to-sky-100'
    case 'blue':
      return 'to-blue-100'
    case 'orange':
      return 'to-orange-100'
    case 'purple':
      return 'to-purple-100'
    case 'indigo':
      return 'to-indigo-100'
    case 'cyan':
      return 'to-cyan-100'
    case 'red':
      return 'to-red-100'
    default:
      return 'to-gray-100'
  }
}
/**
 * @param {Token} coin
 */
export function getSwapCardShadowStyleForCoin(coin) {
  switch (COIN_COLORS[coin.symbol]) {
    case 'yellow':
      return 'shadow-yellow-xl hover:shadow-yellow-2xl'
    case 'green':
      return 'shadow-green-xl hover:shadow-green-2xl'
    case 'lime':
      return 'shadow-lime-xl hover:shadow-lime-2xl'
    case 'sky':
      return 'shadow-sky-xl hover:shadow-sky-2xl'
    case 'blue':
      return 'shadow-blue-xl hover:shadow-blue-2xl'
    case 'orange':
      return 'shadow-orange-xl hover:shadow-orange-2xl'
    case 'purple':
      return 'shadow-purple-xl hover:shadow-purple-2xl'
    case 'indigo':
      return 'shadow-indigo-xl hover:shadow-indigo-2xl'
    case 'cyan':
      return 'shadow-cyan-xl hover:shadow-cyan-2xl'
    case 'red':
      return 'shadow-red-xl hover:shadow-red-2xl'
    default:
      return 'shadow-xl hover:shadow-2xl'
  }
}

/**
 * @param {Token} coin
 */
export function getBorderStyleForCoinHover(coin) {
  switch (COIN_COLORS[coin.symbol]) {
    case 'yellow':
      return 'hover:border-yellow-300'
    case 'green':
      return 'hover:border-green-300'
    case 'lime':
      return 'hover:border-lime-300'
    case 'sky':
      return 'hover:border-sky-300'
    case 'blue':
      return 'hover:border-blue-300'
    case 'orange':
      return 'hover:border-orange-300'
    case 'purple':
      return 'hover:border-purple-300'
    case 'indigo':
      return 'hover:border-indigo-300'
    case 'cyan':
      return 'hover:border-cyan-300'
    case 'red':
      return 'hover:border-red-300'
    default:
      return 'hover:border-gray-200'
  }
}

export { COIN_COLORS }
