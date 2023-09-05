import _ from 'lodash'

import { BRIDGE_MAP } from './bridgeMap'

export const FILTERED = _(BRIDGE_MAP)
  .mapValues((chainObj) => {
    return _(chainObj)
      .pickBy(
        (tokenObj: any) =>
          Array.isArray(tokenObj.swappable) && tokenObj.swappable.length > 0
      )
      .value()
  })
  .pickBy((value, _key) => Object.values(value).length > 0)
  .value()

export const SWAP_CHAIN_IDS = Object.keys(FILTERED).map(Number)
