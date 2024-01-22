export const getTimeMinutesFromNow = (minutesFromNow: number): number => {
  const currentTimeSeconds = new Date().getTime() / 1000

  return Math.round(currentTimeSeconds + 60 * minutesFromNow)
}
