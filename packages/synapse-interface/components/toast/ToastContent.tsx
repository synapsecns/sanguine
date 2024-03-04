import toast from 'react-hot-toast'
import {
  XIcon,
  CheckCircleIcon,
  ExclamationCircleIcon,
} from '@heroicons/react/outline'

export default function ToastContent({ toastData, icon, message }) {
  let borderColor
  let fancyIcon
  if (toastData.type === 'success') {
    borderColor = 'border-green-500'
    fancyIcon = <CheckCircleIcon className="w-5 h-5 text-green-700" />
  } else if (toastData.type === 'error') {
    borderColor = 'border-red-500'
    fancyIcon = <ExclamationCircleIcon className="w-5 h-5 text-red-700" />
  } else {
    borderColor = 'border-indigo-500'
  }

  return (
    <div
      className={`
        flex rounded-md items-center
        min-w-[300px]
        px-2 pt-1 pb-2
        bg-slate-900/10 backdrop-blur-2xl
        text-white
        border ${borderColor}
      `}
    >
      <div className="flex flex-grow pt-1">
        <div className="self-center flex-shrink align-middle justify-items-center">
          {fancyIcon}
        </div>
        <div className="flex-grow">
          <div className="text-sm">{message}</div>
        </div>
      </div>
      <div className="flex-shrink px-2">
        {toastData.type !== 'loading' && (
          <button
            className={`
            rounded-full
            h-6 w-6
            mt-1.5
            focus:outline-none active:outline-none
            hover:bg-gray-900
            text-gray-400 hover:text-gray-300
          `}
            onClick={() => toast.dismiss(toastData.id)}
          >
            <XIcon className="inline w-full h-full p-1 align-middle place-self-center" />
          </button>
        )}
      </div>
    </div>
  )
}
