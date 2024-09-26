export const overrideJsonBigIntSerialization = () => {
    const originalJSONStringify = JSON.stringify
  
    JSON.stringify = function (value: any, replacer, space: number): string {
      const bigIntReplacer = (_key: string, value: any): any => {
        if (typeof value === 'bigint') {
          return parseInt(value.toString())
        }
        return value
      }
  
      const customReplacer = (key: string, value: any): any => {
        if (Array.isArray(replacer) && !replacer.includes(key) && key !== '') {
          return undefined
        }
  
        const modifiedValue = bigIntReplacer(key, value)
  
        if (typeof replacer === 'function') {
          return replacer(key, modifiedValue)
        }
  
        return modifiedValue
      }
  
      return originalJSONStringify(
        value,
        replacer != null ? customReplacer : bigIntReplacer,
        space
      )
    }
  }