export function BridgeSwapOutputNumber({quote}) {
  return (
    <input
      pattern="[0-9.]+"
      disabled={true}
      className={`
        focus:outline-none
        focus:ring-0
        focus:border-none
        border-none
        p-0
        bg-transparent
        max-w-[200px]
      placeholder:text-[#88818C]
      text-white/90 text-2xl md:text-3xl font-medium
      -mt-1
      `}
      placeholder="0.0000"
      value={
        quote?.outputAmountString === '0'
          ? ''
          : quote?.outputAmountString
      }
      name="inputRow"
      autoComplete="off"
    />
  )

}

