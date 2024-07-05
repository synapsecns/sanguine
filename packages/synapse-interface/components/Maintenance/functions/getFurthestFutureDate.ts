export const getFurthestFutureDate = (
  date1: Date | null,
  date2: Date | null
): Date | null => {
  if (date1 === null && date2 === null) return null
  if (date1 === null) return date2
  if (date2 === null) return date1
  return date1 > date2 ? date1 : date2
}
