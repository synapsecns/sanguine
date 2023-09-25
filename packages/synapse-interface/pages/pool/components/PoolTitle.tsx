const PoolTitle = ({ pool }) => {
  return (
    <div className="mb-5">
      <div className="inline-flex items-center mt-2">
        <div className="items-center hidden mr-4 md:flex lg:flex">
          {pool?.poolTokens &&
            pool.poolTokens.map((token) => (
              <img
                key={token.symbol}
                className="relative inline-block w-8 -mr-2 text-white shadow-solid"
                src={token.icon.src}
              />
            ))}
        </div>
        <h3 className="ml-2 mr-2 text-lg font-medium text-white md:ml-0 md:text-2xl">
          {pool?.name}
        </h3>
      </div>
    </div>
  )
}

export default PoolTitle
