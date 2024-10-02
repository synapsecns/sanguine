import * as CHAINS from '../constants/chains/master'

export const makeMultiChainObj = (valOrObj: any) => {
  if (typeof valOrObj === 'object') {
    return valOrObj
  } else {
    const obj: { [key: number]: any } = {}
    for (const chain of Object.values(CHAINS)) {
      obj[chain.id] = valOrObj
    }
    return obj
  }
}
