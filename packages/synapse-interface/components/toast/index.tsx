// @ts-nocheck
import React from 'react'
import toast, { Toaster, ToastBar } from 'react-hot-toast'
import ToastContent from './ToastContent'

const CustomToaster: React.FC = () => {
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
