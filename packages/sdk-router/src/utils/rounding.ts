export const roundDown = (value: number, precision: number) => {
  return Math.floor(value / precision) * precision
}

export const roundUp = (value: number, precision: number) => {
  return Math.ceil(value / precision) * precision
}
