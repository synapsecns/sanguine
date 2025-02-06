import { IconProps } from '@utils/types'

export const BinanceWalletIcon = ({ className = '', ...props }: IconProps) => {
  return (
    <svg
      width="32"
      height="32"
      viewBox="0 0 32 32"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      className={className}
      {...props}
    >
      <rect width="32" height="32" rx="8" fill="#F3BA2F"/>
      <path
        d="M11.0987 14.0987L16 9.19733L20.904 14.1013L23.4627 11.5427L16 4.08L8.54 11.54L11.0987 14.0987ZM4.08 16L6.63867 13.4413L9.19733 16L6.63867 18.5587L4.08 16ZM11.0987 17.9013L16 22.8027L20.904 17.8987L23.4653 20.456L16 27.92L8.54 20.46L8.53733 20.4573L11.0987 17.9013ZM22.8027 16L25.3613 13.4413L27.92 16L25.3613 18.5587L22.8027 16ZM18.8307 15.9973H18.832L16 13.1653L13.912 15.2533L13.7013 15.464L13.168 15.9973L13.1653 16L13.168 16.0027L16 18.8347L18.832 16.0027L18.8347 16L18.8307 15.9973Z"
        fill="white"
      />
    </svg>
  )
}
