export const checkCleanedValue = (cleanedValue) => {
  return cleanedValue === 0 || cleanedValue === '' || isNaN(+cleanedValue)
}
