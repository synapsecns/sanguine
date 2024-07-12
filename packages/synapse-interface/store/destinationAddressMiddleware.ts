import { Middleware } from '@reduxjs/toolkit'

import { screenAddress } from '@/utils/screenAddress'
import { setDestinationAddress } from '@/slices/bridge/reducer'

export const destinationAddressMiddleware: Middleware =
  (store) => (next) => (action) => {
    if (setDestinationAddress.match(action) && action.payload !== null) {
      screenAddress(action.payload)
    }
    return next(action)
  }
