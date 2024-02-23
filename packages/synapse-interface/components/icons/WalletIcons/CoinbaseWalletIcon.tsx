import type { IconProps } from '@types'

export const CoinbaseWalletIcon: React.FC<IconProps> = (props) => {
  return (
    <svg
      width="50"
      height="50"
      viewBox="0 0 32 32"
      fill="black"
      xmlns="http://www.w3.org/2000/svg"
      {...props}
    >
      <g clipPath="url(#coinbasewallet-svg-clipPath)">
        <path
          fillRule="evenodd"
          clipRule="evenodd"
          d="M32 16C32 24.8366 24.8366 32 16 32C7.16345 32 0 24.8366 0 16C0 7.16345 7.16345 0 16 0C24.8366 0 32 7.16345 32 16ZM25.4787 16C25.4787 21.235 21.235 25.4787 16 25.4787C10.7651 25.4787 6.52133 21.235 6.52133 16C6.52133 10.7651 10.7651 6.52134 16 6.52134C21.235 6.52134 25.4787 10.7651 25.4787 16ZM13.6493 12.891C13.3143 12.891 13.0426 13.1626 13.0426 13.4977V18.5024C13.0426 18.8374 13.3143 19.109 13.6493 19.109H18.3507C18.6857 19.109 18.9574 18.8374 18.9574 18.5024V13.4977C18.9574 13.1626 18.6857 12.891 18.3507 12.891H13.6493Z"
          fill="#2F5CE2"
        ></path>
      </g>
      <defs>
        <clipPath id="coinbasewallet-svg-clipPath">
          <rect width="32" height="32" fill="white"></rect>
        </clipPath>
      </defs>
    </svg>
  )
}
