import { Middleware } from '@reduxjs/toolkit'

import { screenAddress } from '@/utils/screenAddress'
import { setDestinationAddress } from '@/slices/bridge/reducer'

export const destinationAddressMiddleware: Middleware =
  (_store) => (next) => async (action) => {
    if (setDestinationAddress.match(action) && action.payload !== null) {
      const isRisky = await screenAddress(action.payload)
      if (isRisky) {
        return
      }
    }
    return next(action)
  }
