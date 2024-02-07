const LoadingText = () => {
  return (
    <div
      className={`
            rounded-md
            transform-gpu transition-all duration-75
            animate-pulse hover:border-opacity-30
          `}
    >
      <div className="w-10 py-2 m-1.5 rounded bg-slate-700"></div>
    </div>
  )
}

export default LoadingText
