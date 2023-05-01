import { SYNAPSE_BASE_URL } from '@/constants/urls'

const IPFS_URI_REGEXP = /^ipfs:(\/\/)?(.*)$/i
const IPNS_URI_REGEXP = /^ipns:(\/\/)?(.*)$/i

/**
 * handle IPFS bullshit to be usable
 * Given a URI that may be ipfs, ipns, http, or https protocol,
 * return the fetch-able http(s) URLs for the same content
 * @param {string} uri to convert to fetch-able http url
 * @return {string[]} array of urls
 */
export function uriToHttp(uri: string) {
  let protocol
  if (uri[0] === '/') {
    protocol = '/'
  } else {
    protocol = uri.split(':')[0].toLowerCase()
  }

  switch (protocol) {
    case 'https':
      return [uri]
    case 'http':
      // udpate
      return [`https${uri.substr(4)}`, uri]
    case 'ipfs':
      const hash = uri.match(IPFS_URI_REGEXP)?.[2]
      return [
        `https://cloudflare-ipfs.com/ipfs/${hash}/`,
        `https://ipfs.io/ipfs/${hash}/`,
      ]
    case 'ipns':
      const name = uri.match(IPNS_URI_REGEXP)?.[2]
      return [
        `https://cloudflare-ipfs.com/ipns/${name}/`,
        `https://ipfs.io/ipns/${name}/`,
      ]
    case 'data':
      return [uri]
    case '/':
      return [`${SYNAPSE_BASE_URL}${uri}`]
    default:
      return []
  }
}
