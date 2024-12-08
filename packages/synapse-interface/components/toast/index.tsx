import React from 'react'
import { Toaster as HotToaster, ToastBar as HotToastBar } from 'react-hot-toast'
import ToastContent from './ToastContent'

const Toaster = HotToaster as React.FC<any>
const ToastBar = HotToastBar as React.FC<any>

export default function CustomToaster(): React.ReactElement {
  return (
    <Toaster
      position="bottom-right"
      containerClassName="pt-8"
      toastOptions={{
        style: {
          background: 'transparent',
          padding: '0px',
        },
        className: `
            shadow-indigo
            bg-gray-800
            text-gray-400
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

export default CustomToaster
