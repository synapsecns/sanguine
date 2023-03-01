import { EventEmitter } from 'events'
import { getProvider } from '@utils/getProvider'



export class Transaction extends EventEmitter {
  constructor({
    hash,
    chainId,
    toChainId,
    timestamp,
    token,
    wait,
    pending=true,
  }) {
    super()
    this.hash = (hash || '').trim()//.toLowerCase()
    this.chainId = chainId
    this.wait = wait

    if (toChainId) {
      this.toChainId = toChainId
    }

    this.provider = getProvider(chainId)
    this.timestamp = timestamp || Date.now()
    if (token) {
      this.token = token
    }
    this.pending = pending
    this.wait().then((receipt) => {
      console.log(receipt)
      this.status = !!receipt.status
      this.pending = false
      this.emit('pending', false, this)
    })

    console.debug('transaction:', this.hash)
  }

  // async getTransaction() {
  //   return this.provider.getTransaction(this.hash)
  // }

  toObject() {
    const {
      hash,
      chainId,
      pending,
      timestamp,
      token,
      toChainId,
    } = this
    return {
      hash,
      chainId,
      pending,
      timestamp,
      token,
      toChainId,
    }
  }

  static fromObject({
    hash,
    chainId,
    pending,
    timestamp,
    token,
    toChainId,
  }) {

    return new Transaction({
      hash,
      chainId,
      pending,
      timestamp,
      token,
      toChainId,
    })
  }
}

