import _ from 'lodash'
import { useEffect, useState } from 'react'
import { namehash } from 'ethers/lib/utils'

import { isAddress } from '@utils/isAddress'
import { isZero } from '@utils/isZero'

import { useENSRegistrarContract } from '@hooks/ens/useENSRegistrarContract'
import { useENSResolverContract } from '@hooks/ens/useENSResolverContract'

/**
 * Does a reverse lookup for an address to find its ENS name.
 * Note this is not the same as looking up an ENS name to find an address
 * because that would actually make sense.
 * @param {string} account
 */
export function useENSName(account) {
  const [resolverAddress, setResolverAddress] = useState(undefined)

  const [name, setName] = useState(undefined)

  let ensNodeArgument
  try {
    if (account && isAddress(account)) {
      ensNodeArgument = namehash(
        `${account.toLowerCase().substr(2)}.addr.reverse`
      )
    }
  } catch (error) {
    console.log(error)
  }

  const registrarContract = useENSRegistrarContract(false)

  useEffect(() => {
    if (registrarContract && ensNodeArgument) {
      try {
        registrarContract
          .resolver(ensNodeArgument)
          .then((stuff) => setResolverAddress(stuff))
          .catch((e) => console.log(e))
      } catch (e) {
        console.log(e)
      }
    }
  }, [ensNodeArgument, registrarContract])
  // const resolverAddress = useSingleCallResult(registrarContract, 'resolver', ensNodeArgument)
  // const resolverAddressResult = resolverAddress?.[0]
  const resolverContract = useENSResolverContract(
    resolverAddress && !isZero(resolverAddress) ? resolverAddress : undefined,
    false
  )

  useEffect(() => {
    if (resolverContract && resolverAddress && !isZero(resolverAddress)) {
      try {
        resolverContract
          .name(ensNodeArgument)
          .then((n) => setName(n))
          .catch((e) => console.log(e))
      } catch (e) {
        console.log(e)
      }
    }
  }, [resolverAddress, resolverContract])

  return name
}
