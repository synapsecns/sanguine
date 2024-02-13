import { useEffect, useRef, useState } from 'react'
import { DownArrow } from '@/components/icons/DownArrow'
import { useKeyPress } from '@/utils/hooks/useKeyPress'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'

export const DropdownMenu = ({ menuTitleElement, children }) => {
  const [open, setOpen] = useState<boolean>(false)
  const ref = useRef(null)
  const handleClick = () => {
    setOpen(!open)
  }

  const escPressed = useKeyPress('Escape')

  function escFunc() {
    if (escPressed) {
      handleClose()
    }
  }

  useEffect(escFunc, [escPressed])

  const handleClose = () => {
    setOpen(false)
  }

  useCloseOnOutsideClick(ref, handleClose)

  return (
    <div id="dropdown-menu" className="relative" ref={ref}>
      <div className="space-x-2">
        <div className='inline-block'>
          {menuTitleElement}
        </div>
        <div className='inline-block'>
          <div
            onClick={handleClick}
            className={`
              flex place-items-center justify-center
              w-5 h-[21px] rounded
              border border-solid border-white/10
              cursor-pointer hover:border-white/50 active:border-white/70
              group
            `}
          >

            <DownArrow className="fill-white/40 group-hover:fill-white/80 group-active:fill-white/100 " />
          </div>
        </div>

      </div>


      {open && (
        <ul
          className={`
            absolute z-50 mt-1 p-1 -right-1
            border border-white/20 bg-bgBase/10 backdrop-blur-lg
            rounded-md overflow-hidden shadow-md
            popover list-none text-left text-sm
          `}
        >
          {children}
        </ul>
      )}
    </div>
  )
}
