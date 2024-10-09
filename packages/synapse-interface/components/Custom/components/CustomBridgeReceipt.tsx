export const CustomBridgeReceipt = ({ bridgeQuote, toToken }) => {
  return (
    <>
      {bridgeQuote?.outputAmountString !== '' && (
        <div className="flex justify-end mt-3 text-sm">
          <div>
            <div className="opacity-75">
              {bridgeQuote?.estimatedTime} seconds via{' '}
              {bridgeQuote?.bridgeModuleName}
            </div>
            <div className="flex justify-end space-x-2">
              <div className="opacity-75">Receive:</div>
              <div>
                {bridgeQuote?.outputAmountString} {toToken?.symbol}
              </div>
            </div>
          </div>
        </div>
      )}
      {bridgeQuote?.outputAmountString === '' && (
        <div className="flex justify-end mt-3 text-sm">Powered by Synapse</div>
      )}
    </>
  )
}
