import { BigNumber } from '@ethersproject/bignumber'

// Default periods for deadlines on origin and destination chains respectively, in seconds
export const TEN_MINUTES = 10 * 60
export const ONE_HOUR = 60 * 60
export const ONE_WEEK = 7 * 24 * 60 * 60

export const calculateDeadline = (seconds: number) => {
  return BigNumber.from(Math.floor(Date.now() / 1000) + seconds)
}

export const applyOptionalDeadline = (
  deadline: BigNumber | undefined,
  defaultPeriod: number
): BigNumber => {
  return deadline ?? calculateDeadline(defaultPeriod)
}
