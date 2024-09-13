import { Request, Response, NextFunction } from 'express'
import { validationResult } from 'express-validator'

export const showFirstValidationError = (
  req: Request,
  res: Response,
  next: NextFunction
): void => {
  const errors = validationResult(req)
  if (!errors.isEmpty()) {
    const firstError = errors.array({ onlyFirstError: true })[0]

    res.status(400).json({
      error: {
        value: (firstError as any).value,
        message: firstError.msg,
        field: firstError.type === 'field' ? firstError.path : undefined,
        location: (firstError as any).location,
      },
    })
    return
  }
  next()
}
