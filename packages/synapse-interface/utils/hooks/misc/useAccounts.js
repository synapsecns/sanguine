import { useEffect, useState } from 'react'
import { useActiveWeb3React } from  '@hooks/wallet/useActiveWeb3React'


export function useAccounts() {
  const { library } = useActiveWeb3React()
  const [accounts, setAccounts] = useState([])
  useEffect( async () => {
      const requestedAccounts = await library.provider.request({ method: 'eth_requestAccounts' })
      setAccounts(requestedAccounts)
    },
    [library]

  )

  console.log({accounts})
  return { accounts }
}