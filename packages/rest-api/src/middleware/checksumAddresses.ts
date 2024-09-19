import { Request, Response, NextFunction } from 'express'
import { getAddress, isAddress } from 'ethers/lib/utils'

export const checksumAddresses = (addressFields: string[]) => {
  return (req: Request, _res: Response, next: NextFunction) => {
    for (const field of addressFields) {
      const address = req.query[field]
      if (typeof address === 'string' && isAddress(address)) {
        req.query[field] = getAddress(address)
      }
    }
    next()
  }
}
