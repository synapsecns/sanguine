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

/* If border does not appear, set tailwind class 'border' on target element */

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

/* If border does not appear, set tailwind class 'border' on target element */

export const getActiveStyleForButton = (color: string) => {
  switch (color) {
    case ColorOptions.ETH:
      return 'bg-[#5170ad44] dark:bg-[#5170ad44] border-[#5170ad] dark:border-[#5170ad]'
    case ColorOptions.RED:
      return 'bg-red-500/25 dark:bg-red-500/25 border-red-500 dark:border-red-500'
    case ColorOptions.ORANGE:
      return 'bg-orange-500/25 dark:bg-orange-500/25 border-orange-500 dark:border-orange-500'
    case ColorOptions.YELLOW:
      return 'bg-yellow-500/25 dark:bg-yellow-500/25 border-yellow-500 dark:border-yellow-500'
    case ColorOptions.LIME:
      return 'bg-lime-500/25 dark:bg-lime-500/25 border-lime-500 dark:border-lime-500'
    case ColorOptions.GREEN:
      return 'bg-green-500/25 dark:bg-green-500/25 border-green-500 dark:border-green-500'
    case ColorOptions.TEAL:
      return 'bg-teal-500/25 dark:bg-teal-500/25 border-teal-500 dark:border-teal-500'
    case ColorOptions.CYAN:
      return 'bg-cyan-500/25 dark:bg-cyan-500/25 border-cyan-500 dark:border-cyan-500'
    case ColorOptions.SKY:
      return 'bg-sky-500/25 dark:bg-sky-500/25 border-sky-500 dark:border-sky-500'
    case ColorOptions.BLUE:
      return 'bg-blue-500/25 dark:bg-blue-500/25 border-blue-500 dark:border-blue-500'
    case ColorOptions.INDIGO:
      return 'bg-indigo-500/25 dark:bg-indigo-500/25 border-indigo-500 dark:border-indigo-500'
    case ColorOptions.PURPLE:
      return 'bg-purple-500/25 dark:bg-purple-500/25 border-purple-500 dark:border-purple-500'
    case ColorOptions.GRAY:
    default:
      return 'dark:bg-zinc-700 border-zinc-400 dark:border-zinc-400'
  }
}
