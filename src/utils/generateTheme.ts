import { CustomThemeVariables } from 'types'

export function generateTheme(theme: CustomThemeVariables = {}) {
  function str2hsl(color: string) {
    // TODO: Support hex short codes
    if (/^#?[\da-fA-F]{6,8}$/.test(color)) {
      let [r, g, b, a] = color
        .match(/[\da-fA-F]{2}/g)
        .map((a) => parseInt(a, 16) / 255)
      return rgb2hsl({ r, g, b, a })
    } else if (/^rgb/.test(color)) {
      let [r, g, b, a] = color.match(/\d+/g).map((a) => +a / 255)
      return rgb2hsl({ r, g, b, a })
    } else if (/^hsl/.test(color)) {
      let [h, s, l, a] = color.match(/\d+/g).map((a) => +a)
      return { h, s, l, a }
    } else {
      return color === 'dark'
        ? { h: 0, s: 0, l: 0, a: 100 }
        : { h: 0, s: 0, l: 100, a: 100 }
    }
  }

  function rgb2hsl({ r, g, b, a = 1 }: any) {
    // in: r,g,b in [0,1], out: h in [0,360) and s,l in [0,100] // https://stackoverflow.com/a/54071699
    let v = Math.max(r, g, b),
      c = v - Math.min(r, g, b),
      f = 1 - Math.abs(v + v - c - 1)
    let h =
      c && (v === r ? (g - b) / c : v === g ? 2 + (b - r) / c : 4 + (r - g) / c)
    return {
      h: 60 * (h < 0 ? h + 6 : h),
      s: f ? (100 * c) / f : 0,
      l: (100 * (v + v - c)) / 2,
      a: 100 * a,
    }
  }

  const hslString = (h: number, s: number, l: number, a: number, x: number) =>
    colorMode === 'dark' || l < 50
      ? `hsl(${h}deg ${s}% ${x * 100 + l * x}%)`
      : `hsl(${h}deg ${s}% ${Math.min(100, l * (1 + x)) * x}%)`

  const colorMode = theme.bgColor === 'dark' ? 'dark' : 'light'

  const { h, s, l, a } = str2hsl(theme.bgColor)

  const generatedTheme =
    colorMode === 'dark' || l < 50
      ? {
          '--synapse-text': hslString(h, s, l, a, 0.96),
          '--synapse-secondary': hslString(h, s / 2, l, a, 0.6),
          '--synapse-border': hslString(h, s, l, a, 0.12),
          '--synapse-object': hslString(h, s, l, a, 0.25),
          '--synapse-surface': hslString(h, s, l, a, 0.12),
          '--synapse-root': hslString(h, s, l, a, 0.07),

          '--synapse-focus': 'var(--synapse-secondary)',

          '--synapse-select-bg': 'var(--synapse-object)',
          '--synapse-select-text': 'var(--synapse-text)',
          '--synapse-select-border': 'var(--synapse-object)',

          '--synapse-button-bg': 'var(--synapse-surface)',
          '--synapse-button-text': 'var(--synapse-text)',
          '--synapse-button-border': 'var(--synapse-border)',
        }
      : {
          '--synapse-text': hslString(h, s, l, a, 0.04),
          '--synapse-secondary': hslString(h, s / 2, l, a, 0.41),
          '--synapse-border': hslString(h, s, l, a, 0.86),
          '--synapse-object': hslString(h, s, l, a, 0.5),
          '--synapse-surface': hslString(h, s, l, a, 1.0),
          '--synapse-root': hslString(h, s, l, a, 0.96),

          '--synapse-focus': 'var(--synapse-secondary)',

          '--synapse-select-bg': 'var(--synapse-root)',
          '--synapse-select-text': 'var(--synapse-text)',
          '--synapse-select-border': 'var(--synapse-border)',

          '--synapse-button-bg': 'var(--synapse-surface)',
          '--synapse-button-text': 'var(--synapse-text)',
          '--synapse-button-border': 'var(--synapse-border)',
        }

  for (const key in theme) if (/^--/.test(key)) generatedTheme[key] = theme[key]

  return generatedTheme as React.CSSProperties
}
