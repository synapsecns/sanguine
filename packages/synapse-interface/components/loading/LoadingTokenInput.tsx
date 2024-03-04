export default function LoadingTokenInput() {
  return (
    <div className="mt-4">
      <div className="border-none rounded-md">
        <div className="flex space-x-2">
          <div
            className={`
            flex flex-grow items-center
            px-3 sm:px-4
            w-full h-20
            rounded-md
            border border-white border-opacity-20
            transform-gpu transition-all duration-75
            animate-pulse hover:border-opacity-30
          `}
          >
            <div className="w-40 h-12 px-2 py-1.5 m-2 rounded bg-slate-700"></div>
            <div className="flex-grow h-12 px-2 py-1.5 m-2 rounded bg-slate-700"></div>
          </div>{' '}
        </div>{' '}
      </div>{' '}
    </div>
  )
}

