import useSWR, { SWRConfiguration } from 'swr'

const jsonFetcher = async (url: string) => {
  return fetch(url).then((res) => res.json())
}

export const useSwr = (endpoint: string | null, options?: SWRConfiguration) => {
  return useSWR(endpoint, jsonFetcher, options)
}
