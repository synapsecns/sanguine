export const ConnectedIndicator = () => {
  return (
    <button
      data-test-id="connected-button"
      className={`
      flex items-center justify-center
      text-base text-white px-3 py-1 rounded-3xl
      text-center transform-gpu transition-all duration-75
      border border-solid border-transparent
      `}
    >
      <div className="flex flex-row text-sm">
        <div
          className={`
            my-auto ml-auto mr-2 w-2 h-2
            bg-green-500 rounded-full
            `}
        />
        Connected
      </div>
    </button>
  )
}
