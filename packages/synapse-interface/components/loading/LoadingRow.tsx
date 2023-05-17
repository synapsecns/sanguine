const LoadingRow = () => {
  return (
    <div className="p-1.5">
      <div
        className={`
            flex flex-grow items-center
            w-full
            rounded-lg
            transform-gpu transition-all duration-75
            animate-pulse hover:border-opacity-30
          `}
      >
        <div className="flex-grow py-2 my-1.5 rounded bg-slate-700"></div>
      </div>
    </div>
  )
}

export default LoadingRow
