export const formatRGB = (color: string) => {
  return color
    .split(',')
    .map((x) => x.replace(/[^0-9.]/g, ''))
    .join(',')
}
