const ColorOptions = {
  ETH: 'eth',
  RED: 'red',
  ORANGE: 'orange',
  YELLOW: 'yellow',
  LIME: 'lime',
  GREEN: 'green',
  TEAL: 'teal',
  CYAN: 'cyan',
  SKY: 'sky',
  BLUE: 'blue',
  INDIGO: 'indigo',
  PURPLE: 'purple',
  GRAY: 'gray',
}

/* TODO: Combine with chain.ts function, remove all other functions in this file */
export const getHoverStyleForButton = (color: string) => {
  switch (color) {
    case ColorOptions.ETH:
      return 'hover:bg-[#5170ad44] hover:dark:bg-[#5170ad44] hover:border-[#5170ad] hover:dark:border-[#5170ad]'
    case ColorOptions.RED:
      return 'hover:bg-red-500/25 hover:dark:bg-red-500/25 hover:border-red-500 hover:dark:border-red-500'
    case ColorOptions.ORANGE:
      return 'hover:bg-orange-500/25 hover:dark:bg-orange-500/25 hover:border-orange-500 hover:dark:border-orange-500'
    case ColorOptions.YELLOW:
      return 'hover:bg-yellow-500/25 hover:dark:bg-yellow-500/25 hover:border-yellow-500 hover:dark:border-yellow-500'
    case ColorOptions.LIME:
      return 'hover:bg-lime-500/25 hover:dark:bg-lime-500/25 hover:border-lime-500 hover:dark:border-lime-500'
    case ColorOptions.GREEN:
      return 'hover:bg-green-500/25 hover:dark:bg-green-500/25 hover:border-green-500 hover:dark:border-green-500'
    case ColorOptions.TEAL:
      return 'hover:bg-teal-500/25 hover:dark:bg-teal-500/25 hover:border-teal-500 hover:dark:border-teal-500'
    case ColorOptions.CYAN:
      return 'hover:bg-cyan-500/25 hover:dark:bg-cyan-500/25 hover:border-cyan-500 hover:dark:border-cyan-500'
    case ColorOptions.SKY:
      return 'hover:bg-sky-500/25 hover:dark:bg-sky-500/25 hover:border-sky-500 hover:dark:border-sky-500'
    case ColorOptions.BLUE:
      return 'hover:bg-blue-500/25 hover:dark:bg-blue-500/25 hover:border-blue-500 hover:dark:border-blue-500'
    case ColorOptions.INDIGO:
      return 'hover:bg-indigo-500/25 hover:dark:bg-indigo-500/25 hover:border-indigo-500 hover:dark:border-indigo-500'
    case ColorOptions.PURPLE:
      return 'hover:bg-purple-500/25 hover:dark:bg-purple-500/25 hover:border-purple-500 hover:dark:border-purple-500'
    case ColorOptions.GRAY:
    default:
      return 'hover:dark:bg-zinc-700 hover:border-zinc-400 hover:dark:border-zinc-400'
  }
}
