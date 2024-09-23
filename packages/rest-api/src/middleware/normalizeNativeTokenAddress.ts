import { Request, Response, NextFunction } from 'express'
import { isAddress, getAddress } from 'ethers/lib/utils'

import { NativeGasAddress, ZeroAddress } from '../constants'

export const normalizeNativeTokenAddress = (addressFields: string[]) => {
  return (req: Request, _res: Response, next: NextFunction) => {
    for (const field of addressFields) {
      const address = req.query[field]
      if (typeof address === 'string' && isAddress(address)) {
        const checksumAddress = getAddress(address)
        req.query[field] =
          checksumAddress === ZeroAddress ? NativeGasAddress : checksumAddress
      }
    }
    next()
  }
}
