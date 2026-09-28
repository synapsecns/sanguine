// Coalesce only simultaneous, identical SDK requests. Resolved quotes (and
// their fees) are never reused for a later refresh.
const requests = new WeakMap<object, Map<string, Promise<any>>>()

export const bridgeQuoteSingleFlight = <T extends Record<string, unknown>>(
  sdk: { bridgeV2: (params: T) => Promise<any> },
  params: T
): Promise<any> => {
  const key = JSON.stringify(
    Object.keys(params)
      .sort((left, right) => left.localeCompare(right))
      .map((field) => [field, typeof params[field], params[field]])
  )
  let sdkRequests = requests.get(sdk)
  if (!sdkRequests) {
    sdkRequests = new Map()
    requests.set(sdk, sdkRequests)
  }
  const existing = sdkRequests.get(key)
  if (existing !== undefined) return existing

  const request = Promise.resolve().then(() => sdk.bridgeV2(params))
  sdkRequests.set(key, request)
  const clear = () => {
    if (sdkRequests.get(key) === request) sdkRequests.delete(key)
  }
  request.then(clear, clear)
  return request
}
