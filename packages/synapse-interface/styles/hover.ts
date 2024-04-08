const ColorOptions = {
  ETH: 'eth',
  RED: 'red',
  ORANGE: 'orange',
  AMBER: 'amber',
  YELLOW: 'yellow',
  LIME: 'lime',
  GREEN: 'green',
  EMERALD: 'emerald',
  TEAL: 'teal',
  CYAN: 'cyan',
  SKY: 'sky',
  BLUE: 'blue',
  INDIGO: 'indigo',
  VIOLET: 'violet',
  PURPLE: 'purple',
  FUCHSIA: 'fuchsia',
  PINK: 'pink',
  ROSE: 'rose',
  SLATE: 'slate',
  GRAY: 'gray',
  ZINC: 'zinc',
  NEUTRAL: 'neutral',
  STONE: 'stone',
}

/* New Tailwind color support:
 * Slate, Zinc, Neutral, Stone
 * Amber, Emerald, Violet, Fuchsia, Pink, Rose
* /

/* If border does not appear, set tailwind class 'border' on target element */

export const getHoverStyleForButton = (color: string) => {
  switch (color) {
    case ColorOptions.ETH:
      return 'hover:bg-[#5170ad44] hover:dark:bg-[#5170ad44] hover:border-[#5170ad] hover:dark:border-[#5170ad]'
    case ColorOptions.RED:
      return 'hover:bg-red-500/25 hover:dark:bg-red-500/25 hover:border-red-500 hover:dark:border-red-500'
    case ColorOptions.ORANGE:
      return 'hover:bg-orange-500/25 hover:dark:bg-orange-500/25 hover:border-orange-500 hover:dark:border-orange-500'
    case ColorOptions.AMBER:
      return 'hover:bg-amber-500/25 hover:dark:bg-amber-500/25 hover:border-amber-500 hover:dark:border-amber-500'
    case ColorOptions.YELLOW:
      return 'hover:bg-yellow-500/25 hover:dark:bg-yellow-500/25 hover:border-yellow-500 hover:dark:border-yellow-500'
    case ColorOptions.LIME:
      return 'hover:bg-lime-500/25 hover:dark:bg-lime-500/25 hover:border-lime-500 hover:dark:border-lime-500'
    case ColorOptions.GREEN:
      return 'hover:bg-green-500/25 hover:dark:bg-green-500/25 hover:border-green-500 hover:dark:border-green-500'
    case ColorOptions.EMERALD:
      return 'hover:bg-emerald-500/25 hover:dark:bg-emerald-500/25 hover:border-emerald-500 hover:dark:border-emerald-500'
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
    case ColorOptions.VIOLET:
      return 'hover:bg-violet-500/25 hover:dark:bg-violet-500/25 hover:border-violet-500 hover:dark:border-violet-500'
    case ColorOptions.PURPLE:
      return 'hover:bg-purple-500/25 hover:dark:bg-purple-500/25 hover:border-purple-500 hover:dark:border-purple-500'
    case ColorOptions.FUCHSIA:
      return 'hover:bg-fuchsia-500/25 hover:dark:bg-fuchsia-500/25 hover:border-fuchsia-500 hover:dark:border-fuchsia-500'
    case ColorOptions.PINK:
      return 'hover:bg-pink-500/25 hover:dark:bg-pink-500/25 hover:border-pink-500 hover:dark:border-pink-500'
    case ColorOptions.ROSE:
      return 'hover:bg-rose-500/25 hover:dark:bg-rose-500/25 hover:border-rose-500 hover:dark:border-rose-500'
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
    case ColorOptions.AMBER:
    case ColorOptions.YELLOW:
      return 'bg-yellow-500/25 dark:bg-yellow-500/25 border-yellow-500 dark:border-yellow-500'
    case ColorOptions.LIME:
      return 'bg-lime-500/25 dark:bg-lime-500/25 border-lime-500 dark:border-lime-500'
    case ColorOptions.GREEN:
      return 'bg-green-500/25 dark:bg-green-500/25 border-green-500 dark:border-green-500'
    case ColorOptions.EMERALD:
      return 'bg-emerald-500/25 dark:bg-emerald-500/25 border-emerald-500 dark:border-emerald-500'
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
    case ColorOptions.VIOLET:
      return 'bg-violet-500/25 dark:bg-violet-500/25 border-violet-500 dark:border-violet-500'
    case ColorOptions.PURPLE:
      return 'bg-purple-500/25 dark:bg-purple-500/25 border-purple-500 dark:border-purple-500'
    case ColorOptions.FUCHSIA:
      return 'bg-fuchsia-500/25 dark:bg-fuchsia-500/25 border-fuchsia-500 dark:border-fuchsia-500'
    case ColorOptions.PINK:
      return 'bg-pink-500/25 dark:bg-pink-500/25 border-pink-500 dark:border-pink-500'
    case ColorOptions.ROSE:
      return 'bg-rose-500/25 dark:bg-rose-500/25 border-rose-500 dark:border-rose-500'
    case ColorOptions.GRAY:
    default:
      return 'dark:bg-zinc-700 border-zinc-400 dark:border-zinc-400'
  }
}

export const getStyleForSvg = (color: string) => {
  switch (color) {
    case ColorOptions.ETH:
      return 'fill-[#5170ad] dark:fill-[#5170ad] stroke-[#5170ad] dark:stroke-[#5170ad]'
    case ColorOptions.RED:
      return 'fill-red-500 dark:fill-red-500 stroke-red-500 dark:stroke-red-500'
    case ColorOptions.ORANGE:
      return 'fill-orange-500 dark:fill-orange-500 stroke-orange-500 dark:stroke-orange-500'
    case ColorOptions.AMBER:
      return 'fill-amber-500 dark:fill-amber-500 stroke-amber-500 dark:stroke-amber-500'
    case ColorOptions.YELLOW:
      return 'fill-yellow-500 dark:fill-yellow-500 stroke-yellow-500 dark:stroke-yellow-500'
    case ColorOptions.LIME:
      return 'fill-lime-500 dark:fill-lime-500 stroke-lime-500 dark:stroke-lime-500'
    case ColorOptions.GREEN:
      return 'fill-green-500 dark:fill-green-500 stroke-green-500 dark:stroke-green-500'
    case ColorOptions.EMERALD:
      return 'fill-emerald-500 dark:fill-emerald-500 stroke-emerald-500 dark:stroke-emerald-500'
    case ColorOptions.TEAL:
      return 'fill-teal-500 dark:fill-teal-500 stroke-teal-500 dark:stroke-teal-500'
    case ColorOptions.CYAN:
      return 'fill-cyan-500 dark:fill-cyan-500 stroke-cyan-500 dark:stroke-cyan-500'
    case ColorOptions.SKY:
      return 'fill-sky-500 dark:fill-sky-500 stroke-sky-500 dark:stroke-sky-500'
    case ColorOptions.BLUE:
      return 'fill-blue-500 dark:fill-blue-500 stroke-blue-500 dark:stroke-blue-500'
    case ColorOptions.INDIGO:
      return 'fill-indigo-500 dark:fill-indigo-500 stroke-indigo-500 dark:stroke-indigo-500'
    case ColorOptions.VIOLET:
      return 'fill-violet-500 dark:fill-violet-500 stroke-violet-500 dark:stroke-violet-500'
    case ColorOptions.PURPLE:
      return 'fill-purple-500 dark:fill-purple-500 stroke-purple-500 dark:stroke-purple-500'
    case ColorOptions.FUCHSIA:
      return 'fill-fuchsia-500 dark:fill-fuchsia-500 stroke-fuchsia-500 dark:stroke-fuchsia-500'
    case ColorOptions.PINK:
      return 'fill-pink-500 dark:fill-pink-500 stroke-pink-500 dark:stroke-pink-500'
    case ColorOptions.ROSE:
      return 'fill-rose-500 dark:fill-rose-500 stroke-rose-500 dark:stroke-rose-500'
    case ColorOptions.SLATE:
      return 'dark:fill-slate-700 stroke-slate-400 dark:stroke-slate-400'
    case ColorOptions.GRAY:
    case ColorOptions.ZINC:
    case ColorOptions.NEUTRAL:
    case ColorOptions.STONE:
    default:
      return 'dark:fill-zinc-400 stroke-zinc-400 dark:stroke-zinc-400'
  }
}

export const getStyleForText = (color: string) => {
  switch (color) {
    // case ColorOptions.ETH:
    //   return 'text-[#5170ad] dark:text-[#5170ad]'
    case ColorOptions.RED:
      return 'text-red-500 dark:text-red-500'
    case ColorOptions.ORANGE:
      return 'text-orange-500 dark:text-orange-500'
    case ColorOptions.AMBER:
      return 'text-amber-500 dark:text-amber-500'
    case ColorOptions.YELLOW:
      return 'text-yellow-500 dark:text-yellow-500'
    case ColorOptions.LIME:
      return 'text-lime-500 dark:text-lime-500'
    case ColorOptions.GREEN:
      return 'text-green-500 dark:text-green-500'
    case ColorOptions.EMERALD:
      return 'text-emerald-500 dark:text-emerald-500'
    case ColorOptions.TEAL:
      return 'text-teal-500 dark:text-teal-500'
    case ColorOptions.CYAN:
      return 'text-cyan-500 dark:text-cyan-500'
    case ColorOptions.SKY:
      return 'text-sky-500 dark:text-sky-500'
    case ColorOptions.BLUE:
      return 'text-blue-500 dark:text-blue-500'
    case ColorOptions.INDIGO:
      return 'text-indigo-500 dark:text-indigo-500'
    case ColorOptions.VIOLET:
      return 'text-violet-500 dark:text-violet-500'
    case ColorOptions.PURPLE:
      return 'text-purple-500 dark:text-purple-500'
    case ColorOptions.FUCHSIA:
      return 'text-fuchsia-500 dark:text-fuchsia-500'
    case ColorOptions.PINK:
      return 'text-pink-500 dark:text-pink-500'
    case ColorOptions.ROSE:
      return 'text-rose-500 dark:text-rose-500'
    case ColorOptions.ETH:
    case ColorOptions.SLATE:
      return 'dark:text-slate-400'
    case ColorOptions.GRAY:
    case ColorOptions.ZINC:
    case ColorOptions.NEUTRAL:
    case ColorOptions.STONE:
    default:
      return 'dark:text-zinc-400 stroke-zinc-400 dark:stroke-zinc-400'
  }
}
