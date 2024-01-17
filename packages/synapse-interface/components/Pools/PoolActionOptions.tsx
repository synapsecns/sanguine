import { Fragment } from 'react'
import { Popover, Transition } from '@headlessui/react'
import { Token } from '@/utils/types'
import { useRouter } from 'next/router'
import { DepositIcon } from '../icons/DepositIcon'
import { StakeIcon } from '../icons/StakeIcon'
import { UnstakeIcon } from '../icons/UnstakeIcon'
import { WithdrawIcon } from '../icons/WithdrawIcon'
import { ClaimIcon } from '../icons/ClaimIcon'
import { DownArrow } from '../icons/DownArrow'

export const PoolActionOptions = ({
  pool,
  options,
}: {
  pool: Token
  options: string[]
}) => {
  const router = useRouter()

  const handleWithdrawClick = () => {
    router.push(`/pool/${pool.routerIndex}`)
  }

  const handleDepositClick = () => {
    router.push(`/pool/${pool.routerIndex}`)
  }

  const handleStakeClick = () => {
    router.push(`/stake/${pool.routerIndex}`)
  }

  const handleUnstakeClick = () => {
    router.push(`/stake/${pool.routerIndex}`)
  }

  const handleClaimClick = () => {
    router.push(`/stake/${pool.routerIndex}`)
  }

  return (
    <Popover className="relative inline-block flex-shrink-0">
      {({ open }) => (
        <>
          <Popover.Button
            as="div"
            onMouseEnter={() => {}}
            className={`
              rounded-md inline-flex items-center gap-2
              px-2 py-1 text-sm rounded-sm cursor-pointer
              border border-zinc-400 dark:border-zinc-500
              hover:bg-zinc-50 hover:dark:bg-zinc-700
              focus:outline-none
            `}
          >
            Actions
            <DownArrow />
          </Popover.Button>
          <TransactionPopoverContainer>
            {options.includes('Deposit') && (
              <OptionButton
                icon={<DepositIcon />}
                text={'Deposit'}
                onClick={handleDepositClick}
              />
            )}
            {options.includes('Stake') && (
              <OptionButton
                icon={<StakeIcon />}
                text={'Stake'}
                onClick={handleStakeClick}
              />
            )}
            {options.includes('Unstake') && (
              <OptionButton
                icon={<UnstakeIcon />}
                text={'Unstake'}
                onClick={handleUnstakeClick}
              />
            )}
            {options.includes('Withdraw') && (
              <OptionButton
                icon={<WithdrawIcon />}
                text={'Withdraw'}
                onClick={handleWithdrawClick}
              />
            )}
            {options.includes('Claim') && (
              <OptionButton
                icon={<ClaimIcon />}
                text={'Claim'}
                onClick={handleClaimClick}
              />
            )}
          </TransactionPopoverContainer>
        </>
      )}
    </Popover>
  )
}

export function TransactionPopoverContainer({
  children,
  className,
}: {
  children: any
  className?: string
}) {
  return (
    <Transition
      as={Fragment}
      enter="transition ease-out duration-200"
      enterFrom="opacity-0 translate-y-1"
      enterTo="opacity-100 translate-y-0"
      leave="transition ease-in duration-150"
      leaveFrom="opacity-100 translate-y-0"
      leaveTo="opacity-0 translate-y-1"
    >
      <Popover.Panel
        style={{ boxShadow: '0px 4px 4px 0px rgba(0, 0, 0, 0.25)' }}
        className={`
          absolute z-10 top-[-74px] left-[50px] w-[110px] transform-gpu
          -translate-x-full border border-separator bg-surface
          rounded-sm overflow-hidden shadow-md
        `}
      >
        <div className="shadow-xl">
          <div className="relative grid gap-1 p-1">{children}</div>
        </div>
      </Popover.Panel>
    </Transition>
  )
}

export const OptionButton = ({
  icon,
  text,
  onClick,
}: {
  icon: any
  text: string
  onClick: () => void
}) => {
  return (
    <div
      data-test-id="option-button"
      onClick={onClick}
      className="flex hover:cursor-pointer hover:bg-[#0A415C] rounded-sm p-1 text-white"
    >
      <div className="my-auto mr-1">{icon}</div>
      <div className="text-sm">{text}</div>
    </div>
  )
}
