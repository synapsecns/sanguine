const LoadingTokenInput = () => {
  return (
    <div className="mt-4">
      <div className="border-none rounded-xl">
        <div className="flex space-x-2">
          <div
            className={`
            flex flex-grow items-center
            pl-3 sm:pl-4
            w-full h-20
            rounded-lg
            border border-white border-opacity-20
            transform-gpu transition-all duration-75
            hover:border-opacity-30
            bg-white
          `}
          >
            <div className="h-4 my-2 mt-9 bg-slate-700 rounded w-[90%]"></div>
            <div className="h-4 my-2 bg-slate-700 rounded w-[90%]"></div>
          </div>{' '}
        </div>{' '}
      </div>{' '}
    </div>
  )
}
export default LoadingTokenInput
