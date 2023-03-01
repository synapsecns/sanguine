export function promisify(promise) {
  let _resolve, _reject

  let wrap = new Promise(async (resolve, reject) => {
    _resolve = resolve
    _reject = reject
    let result = await promise
    resolve(result)
  })

  wrap.resolve = _resolve
  wrap.reject = _reject

  return wrap
}
