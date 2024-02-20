/* @ts-ignore */
import NodeCache from "node-cache"

export const MS_TIMES = {
  ONE_SECOND: 1000,
  THIRTY_SECONDS: 30 * 1000,
  ONE_MINUTE: 60 * 1000,
  FIVE_MINUTES: 5 * 60 * 1000,
  TEN_MINUTES: 10 * 60 * 1000,
  ONE_HOUR: 60 * 60 * 1000,
  ONE_DAY: 24 * 60 * 60 * 1000,
  ONE_WEEK: 7 * 24 * 60 * 60 * 1000,
}
export function SimpleCache(maxAge: number) {
  /* @ts-ignore */
  return function(
    _: Object, // target
    propertyKey: string,
    descriptor: PropertyDescriptor
  ) {
    const originalMethod = descriptor.value
    const cache = new NodeCache({
      stdTTL: maxAge / 1000,
      checkperiod: maxAge / 1000,
      useClones: false // this is to handle promises + performance
    })

    descriptor.value = function(...args: any[]) {
      const key = JSON.stringify(args)

      if (cache.has(key)) {
        console.log(`Returning cached result for ${propertyKey}`)
        return cache.get(key)
      } else {
        console.log(`Calculating result for ${propertyKey}`)
        const result = originalMethod.apply(this, args)
        result.then((res: any) => {
          cache.set(key, res)
          return res
        })
        return result
      }
    }
  }
}
