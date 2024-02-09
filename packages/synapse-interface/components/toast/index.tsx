import toast, { Toaster, ToastBar } from 'react-hot-toast'
import ToastContent from './ToastContent'

export default function CustomToaster() {
  return (
    <Toaster
      position="bottom-right" // top-right
      containerClassName="pt-8"
      toastOptions={{
        // Define default options
        style: {
          background: 'transparent',
          padding: '0px',
        },
        className: `
            shadow-indigo
            bg-gray-800
            text-gray-400
            backdrop-blur-2xl
          `,
        duration: 10000,
      }}
    >
      {(toastData) => (
        <ToastBar toast={toastData} style={{}}>
          {(props) => <ToastContent toastData={toastData} {...props} />}
        </ToastBar>
      )}
    </Toaster>
  )
}
