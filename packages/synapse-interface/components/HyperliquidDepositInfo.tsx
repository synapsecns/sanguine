import { ARBITRUM } from '@/constants/chains/master'

export const HyperliquidDepositInfo = ({
  fromChainId,
  isOnArbitrum,
  hasDepositedOnHyperliquid,
}) => {
  if (fromChainId !== ARBITRUM.id) {
    return (
      <div className="flex flex-col p-2 mb-2 space-y-1 text-sm ">
        <div className="flex justify-between">
          <div>
            <div className="text-[#97FCE4] mb-2">Step 1</div>
            <div className="flex items-center space-x-2">
              <GreenStep1Circle />
              <div>Bridge (Arbitrum)</div>
            </div>
          </div>
          <div>
            <div className="mb-2 text-white/65">Step 2</div>
            <div className="flex items-center space-x-2">
              <GrayStep2Circle />
              <div className="text-white/65">Deposit (Hyperliquid)</div>
            </div>
          </div>
        </div>
      </div>
    )
  }

  if (hasDepositedOnHyperliquid) {
    return (
      <div className="flex flex-col p-2 mb-2 space-y-1 text-sm ">
        <div className="flex justify-between">
          <div>
            <div className="mb-2 text-white/65">Step 1</div>
            <div className="flex items-center space-x-2">
              <CompletedCheckMarkCircle />
              <div className="text-white/65">Bridge (Arbitrum)</div>
            </div>
          </div>
          <div>
            <div className="mb-2 text-white/65">Step 2</div>
            <div className="flex items-center space-x-2">
              <CompletedCheckMarkCircle />
              <div className="text-white/65">Deposit (Hyperliquid)</div>
            </div>
          </div>
        </div>
      </div>
    )
  }

  if (fromChainId === ARBITRUM.id && isOnArbitrum) {
    return (
      <div className="flex flex-col p-2 mb-2 space-y-1 text-sm ">
        <div className="flex justify-between">
          <div>
            <div className="mb-2 text-white/65">Step 1</div>
            <div className="flex items-center space-x-2">
              <CompletedCheckMarkCircle />
              <div className="text-white/65">Bridge (Arbitrum)</div>
            </div>
          </div>
          <div>
            <div className="mb-2 text-[#97FCE4]">Step 2</div>
            <div className="flex items-center space-x-2">
              <GreenStep2Circle />
              <div className="text-[#97FCE4]">Deposit (Hyperliquid)</div>
            </div>
          </div>
        </div>
      </div>
    )
  }
}

const CompletedCheckMarkCircle = () => {
  return (
    <svg
      width="34"
      height="34"
      viewBox="0 0 34 34"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <rect
        x="1"
        y="1"
        width="32"
        height="32"
        rx="16"
        stroke="#565058"
        stroke-width="2"
      />
      <path
        d="M13.7112 23.204C13.7112 23.06 13.6952 22.996 13.6312 22.996L13.2632 23.172C13.2632 23.092 13.2152 23.044 13.1352 23.012L13.0072 22.996C12.8952 22.996 12.8472 23.012 12.6872 23.108C12.6392 23.012 12.5752 22.9 12.5272 22.804C12.1112 22.004 11.6792 20.804 11.4872 20.276C11.3912 20.004 11.2952 19.444 11.1832 18.596C11.3112 18.676 11.4072 18.708 11.4552 18.708C11.5192 18.708 11.5992 18.596 11.6632 18.372C11.6952 18.42 11.7592 18.436 11.8392 18.436C11.8872 18.436 11.9512 18.42 11.9832 18.372L12.2392 17.988L12.5272 18.084H12.5432C12.5752 18.084 12.6232 18.036 12.7032 17.988C12.7832 17.94 12.8472 17.908 12.8952 17.908L12.9432 17.924C13.1992 18.052 13.3752 18.276 13.4552 18.628C13.6472 19.444 13.8232 19.844 14.0312 19.844C14.2072 19.844 14.4472 19.636 14.7032 19.236C14.9592 18.836 15.2152 18.292 15.5032 17.636C15.5192 17.764 15.5352 17.828 15.5672 17.828C15.6632 17.828 15.9032 17.268 16.4952 16.324C17.3752 14.9 19.5512 12.164 20.1112 11.78C20.5272 11.492 20.8472 11.22 21.0712 10.98C21.0392 11.14 21.0072 11.252 21.0072 11.3C21.0072 11.348 21.0392 11.364 21.0712 11.364L21.5192 11.14V11.204C21.5192 11.284 21.5352 11.332 21.5832 11.332C21.6472 11.332 21.9032 11.076 21.9352 10.98L21.9032 11.204L22.4472 10.884L22.3192 11.172C22.4792 11.06 22.6072 10.996 22.6872 10.996C22.7672 10.996 22.8152 11.124 22.8152 11.204C22.8152 11.332 22.7032 11.508 22.5272 11.732C22.3352 11.988 21.8552 12.484 20.4152 14.132C19.7912 14.836 17.0232 18.596 16.4952 19.492L15.5032 21.172C15.0712 21.892 14.7992 22.356 14.6552 22.532C14.5112 22.708 14.3352 22.884 14.1272 23.044L13.9832 22.964L13.8552 23.044L13.7112 23.204Z"
        fill="#97FCE4"
      />
      <circle cx="17" cy="17" r="16" stroke="#565058" stroke-width="2" />
    </svg>
  )
}

const GreenStep1Circle = () => {
  return (
    <svg
      width="34"
      height="34"
      viewBox="0 0 34 34"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <rect
        x="1"
        y="1"
        width="32"
        height="32"
        rx="16"
        stroke="#97FCE4"
        strokeWidth="2"
      />
      <path
        d="M17.2346 22.5V13.124C16.4666 13.828 15.6346 14.18 14.5786 14.372V13.076C15.6986 12.836 16.6746 12.26 17.4586 11.3H18.6106V22.5H17.2346Z"
        fill="#FCFCFD"
      />
    </svg>
  )
}

const GreenStep2Circle = () => {
  return (
    <svg
      width="34"
      height="34"
      viewBox="0 0 34 34"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <rect
        x="1"
        y="1"
        width="32"
        height="32"
        rx="16"
        stroke="#97FCE4"
        strokeWidth="2"
      />
      <path
        d="M13.5034 22.5L13.4874 21.236L17.6954 17.604C19.0074 16.468 19.5674 15.54 19.5674 14.532C19.5674 13.284 18.7354 12.468 17.4234 12.468C15.8554 12.468 14.9754 13.572 15.1354 15.236L13.7274 15.14C13.5514 12.788 15.0874 11.172 17.4234 11.172C19.6314 11.172 21.0714 12.5 21.0714 14.532C21.0714 15.892 20.3194 17.14 18.5914 18.596L15.5514 21.188H21.0714V22.5H13.5034Z"
        fill="#97FCE4"
      />
    </svg>
  )
}

const GrayStep2Circle = () => {
  return (
    <svg
      width="34"
      height="34"
      viewBox="0 0 34 34"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <rect
        x="1"
        y="1"
        width="32"
        height="32"
        rx="16"
        stroke="#565058"
        strokeWidth="2"
      />
      <path
        d="M13.5034 22.5L13.4874 21.236L17.6954 17.604C19.0074 16.468 19.5674 15.54 19.5674 14.532C19.5674 13.284 18.7354 12.468 17.4234 12.468C15.8554 12.468 14.9754 13.572 15.1354 15.236L13.7274 15.14C13.5514 12.788 15.0874 11.172 17.4234 11.172C19.6314 11.172 21.0714 12.5 21.0714 14.532C21.0714 15.892 20.3194 17.14 18.5914 18.596L15.5514 21.188H21.0714V22.5H13.5034Z"
        fill="#C0BCC2"
      />
    </svg>
  )
}
