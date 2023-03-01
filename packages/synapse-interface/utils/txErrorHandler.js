import toast from 'react-hot-toast'

export function txErrorHandler(err) {
  if (err.code === 4001) {
    // EIP-1193 userRejectedRequest error
    toast.error('User denied transaction')
  } else if (err.code === -32603) {
    if (err.data?.code === -32000) {
      toast.error(
        <div>
          <div className="w-full">
            {`Insufficient gas to execute transaction `}
          </div>
        </div>
      )
    }
  }

  console.log(err)
  return err
}