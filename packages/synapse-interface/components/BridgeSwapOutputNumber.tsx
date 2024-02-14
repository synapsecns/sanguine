export function BridgeSwapOutputNumber({quote}) {
  const isQuoteZero = quote && (quote?.outputAmountString === '0')

  return (
    <div
      className={`
        max-w-[200px]
      text-white/90 text-2xl md:text-3xl font-medium
      ${isQuoteZero && "text-white/40"}
      -mt-1

      `}
    >
      {
      convertNumStrToSized(
        isQuoteZero
          ? "0.0000"
          : quote?.outputAmountString
      )

      }
    </div>
  )
}

function convertNumStrToSized(numStr) {
  const [integer, decimal] = numStr.split('.')
  return (
    <>
      {integer}
      {decimal && (
        <>
          .
          <span className="text-xl md:text-2xl ">
            {decimal}
          </span>
        </>
      )}
    </>
  )
}

