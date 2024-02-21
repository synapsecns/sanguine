// @ts-nocheck
/* @ts-ignore */
import NodeCache from "node-cache"

export const CACHE_TIMES = {
  ONE_SECOND: 1,
  THIRTY_SECONDS: 30,
  ONE_MINUTE: 60,
  FIVE_MINUTES: 5 * 60,
  TEN_MINUTES: 10 * 60,
  ONE_HOUR: 60 * 60,
  ONE_DAY: 24 * 60 * 60,
  ONE_WEEK: 7 * 24 * 60 * 60,
  INFINITE: 0
}

export function RouterCache(maxAge: number) {
  /* @ts-ignore */
  return function(
    target: Object, // target
    propertyKey: string,
    descriptor: PropertyDescriptor
  ) {
    const originalMethod = descriptor.value
    const cache = new NodeCache({
      stdTTL: maxAge,
      checkperiod: maxAge,
      useClones: false // this is to handle promises + performance
    })

    descriptor.value = function(...args: any[]) {
      const key = JSON.stringify({
        args,
        propertyKey,
        name:target.constructor.name,
        address: this.address,
        chainId: this.chainId
      })
      const debugDetails = `${propertyKey}(${args}) on ${this.chainId} (${this.address})`
      if (cache.has(key)) {
        console.log(`Returning cached result for ${debugDetails})`)
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
