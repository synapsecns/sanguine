export const checkExists = (value: any): boolean => {
  if (value === null || value === undefined) {
    return false
  }
  return true
}
