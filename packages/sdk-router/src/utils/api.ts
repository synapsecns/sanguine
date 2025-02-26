import { logger } from './logger'

const fetchWithTimeout = async (
  name: string,
  url: string,
  timeout: number,
  params: any = {},
  init?: RequestInit
): Promise<Response | null> => {
  const controller = new AbortController()
  const timeoutId = setTimeout(() => controller.abort(), timeout)
  try {
    const response = await fetch(url, {
      signal: controller.signal,
      ...init,
    })
    if (!response.ok) {
      const text = await response.text()
      logger.info(
        { name, url, params, response, text },
        `${name}: response not OK`
      )
      return null
    }
    return response
  } catch (error) {
    if (error instanceof Error && error.name === 'AbortError') {
      logger.info({ name, url, timeout, params }, `${name}: timed out`)
    } else {
      logger.error(
        { name, url, params, error },
        `${name}: was not able to get a response`
      )
    }
    return null
  } finally {
    clearTimeout(timeoutId)
  }
}

export const getWithTimeout = async (
  name: string,
  url: string,
  timeout: number,
  params: any = {},
  headers: any = {}
): Promise<Response | null> => {
  const urlWithParams = Object.keys(params).length
    ? `${url}?${new URLSearchParams(params)}`
    : url
  return fetchWithTimeout(name, urlWithParams, timeout, params, {
    method: 'GET',
    headers,
  })
}

export const postWithTimeout = async (
  name: string,
  url: string,
  timeout: number,
  params: any,
  headers: any = {}
): Promise<Response | null> => {
  return fetchWithTimeout(name, url, timeout, params, {
    method: 'POST',
    body: JSON.stringify(params),
    headers: {
      ...headers,
      'Content-Type': 'application/json',
    },
  })
}

export const putWithTimeout = async (
  name: string,
  url: string,
  timeout: number,
  params: any,
  headers: any = {}
): Promise<Response | null> => {
  return fetchWithTimeout(name, url, timeout, params, {
    method: 'PUT',
    body: JSON.stringify(params),
    headers: {
      ...headers,
      'Content-Type': 'application/json',
    },
  })
}
