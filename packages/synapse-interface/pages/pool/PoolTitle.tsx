import { POOLS_MAP } from '@hooks/pools/usePools'

export default function PoolTitle({ poolName, poolChainId }) {
  const coins = POOLS_MAP[poolChainId]?.[poolName]

  return (
    <div className="inline-flex items-center mt-2">
      <div className="items-center hidden mr-4 md:flex lg:flex">
        {coins.map((coin) => (
          <img
            key={coin.symbol}
            className="relative inline-block w-8 -mr-2 text-white shadow-solid"
            src={coin.icon}
          />
        ))}
      </div>
      <h3 className="ml-2 mr-2 text-lg font-medium text-white md:ml-0 md:text-2xl">
        {poolName}
      </h3>
    </div>
  )
}
