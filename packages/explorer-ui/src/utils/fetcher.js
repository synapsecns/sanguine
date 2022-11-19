import _ from 'lodash'

export const fetcher = (...args) => fetch(...args).then((res) => res.json())

export const fetcherMany = (...urls) => Promise.all(_.compact(urls).map((url) => fetcher(url)))
