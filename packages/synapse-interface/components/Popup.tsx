import { useState, useEffect } from 'react'
import { IMPAIRED_CHAINS } from '@/constants/impairedChains'
const Popup = ({ chainId }) => {
  const [active, setActive] = useState(false)

  useEffect(() => {
    if (chainId && IMPAIRED_CHAINS[chainId]) {
      setActive(true)
    } else {
      setActive(false)
    }
  }, [chainId])

  if (!active) return null
  return (
    <div className="w-full mb-[100px]">
      <div className="bg-slate-600/40 shadow-lg pt-3 px-6 pb-6 rounded-lg text-white relative">
        <div className="pr-3 text-center">
          {IMPAIRED_CHAINS[chainId] && IMPAIRED_CHAINS[chainId].content()}
        </div>
        <div className="absolute right-3 top-2">
          <button
            className="text-white text-sm hover:text-gray-400"
            onClick={() => setActive(false)}
          >
            ✕
          </button>
        </div>
      </div>
    </div>
  )
}

export default Popup
