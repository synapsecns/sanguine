import toast from 'react-hot-toast'

export const txErrorHandler = (err: any) => {
  if (err.code === 4001) {
    // EIP-1193 userRejectedRequest error
    return toast.error('User denied transaction')
  } else if (err.code === -32603) {
    if (err.data?.code === -32000) {
      return toast.error(
        <div>
          <div className="w-full">
            {`Insufficient gas to execute transaction `}
          </div>
        </div>
      )
    }
  }

  return toast.error(err?.message)
}
