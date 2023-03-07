import useSWR from 'swr'

async function jsonFetcher(url) {
  return fetch(url).then((res) => res.json())
}


// const fetcher = async (...args) => fetch(...args).then(res => res.json())

export function useSwr(endpoint) {

  return useSWR( endpoint, jsonFetcher)
}


