import { Request, Response, NextFunction } from 'express'
import { getAddress } from 'ethers/lib/utils'

export const checksumAddresses = (addressFields: string[]) => {
  return (req: Request, _res: Response, next: NextFunction) => {
    for (const field of addressFields) {
      if (req.query[field] && typeof req.query[field] === 'string') {
        try {
          req.query[field] = getAddress(req.query[field] as string)
        } catch (error) {
          console.log(error)
        }
      }
    }
    next()
  }
}
