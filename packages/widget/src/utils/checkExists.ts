export const checkExists = (value: any): boolean => {
  if (value === null) return false
  if (value === undefined) return false
  return true
}
