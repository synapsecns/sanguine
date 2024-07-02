export const calculateTimeBetween = (
  timeBefore: number,
  timeAfter: number
): number => {
  return Math.abs(timeBefore - timeAfter) * 1000
}
