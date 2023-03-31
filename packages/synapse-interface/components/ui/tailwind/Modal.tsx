import { useEffect } from 'react'
import { useKeyPress } from '@hooks/useKeyPress'

export default function Modal({
  isOpen,
  onClose,
  children,
}: {
  isOpen: boolean
  onClose: () => void
  children: any
}) {
  const escPressed = useKeyPress('Escape')

  function escEffect() {
    if (escPressed) {
      onClose()
    }
  }

  useEffect(escEffect, [escPressed])

  if (isOpen) {
    return (
      <>
        <div className="fixed inset-0 z-50 flex items-start justify-center overflow-x-hidden overflow-y-auto outline-none md:items-center focus:outline-none">
          <div className="relative w-auto max-w-3xl mx-auto my-6">
            <div
              className={`
                border-0 rounded-lg relative flex flex-col w-full outline-none focus:outline-none
                bg-gray-800
              `}
            >
              {children}
            </div>
          </div>
        </div>
        <div className="fixed inset-0 z-40 bg-black opacity-25 md:ml-0 md:space-x-0 -top-2"></div>
      </>
    )
  } else {
    return null
  }
}
