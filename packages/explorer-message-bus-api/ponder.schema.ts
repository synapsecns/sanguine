import { createSchema } from '@ponder/core'

import { STATUSES } from './src/status'

export default createSchema((p) => ({
  TxStatus: p.createEnum(STATUSES),
  Message: p.createTable({
    id: p.hex(),
    srcChainId: p.bigint(),
    srcTxHash: p.hex().optional(),
    srcTimestamp: p.string().optional(),
    dstChainId: p.bigint(),
    dstTxHash: p.hex().optional(),
    dstTimestamp: p.string().optional(),
    status: p.enum('TxStatus'),
  }),
}))
