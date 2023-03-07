import { isAddress } from '@utils/isAddress'
import useENSAddress from '@hooks/ens/useENSAddress'
import useENSName from '@hooks/ens/useENSName'

/**
 * Given a name or address, does a lookup to resolve to an address and name
 * @param nameOrAddress ENS name or address
 */
export default function useENS(nameOrAddress) {
  const validated = isAddress(nameOrAddress)
  const reverseLookup = useENSName(validated ? validated : undefined)
  const lookup = useENSAddress(nameOrAddress)

  return {
    loading: reverseLookup.loading || lookup.loading,
    address: validated ? validated : lookup.address,
    name: reverseLookup.ENSName
      ? reverseLookup.ENSName
      : !validated && lookup.address
        ? nameOrAddress || null
        : null
  }
}