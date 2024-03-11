import { DropDownArrowSvg } from '../icons/DropDownArrowSvg'
import {
  getBorderStyleForCoinHover,
  getMenuItemHoverBgForCoin,
} from '@/styles/tokens'

export function BridgeCard({ children }) {
  /* TODOs
   * Lift margin value up to parent
   * Remove need for popoverDependencies styles
   * Adjust interior elements to allow for single p-4 padding value
   */

  const space = 'px-4 pt-4 pb-2 mt-5'
  const bgColor = 'bg-zinc-100 dark:bg-zinc-900/95'
  const popoverDependencies = 'overflow-hidden transform'

  return (
    <div
      className={`${space} ${bgColor} ${popoverDependencies} rounded-[.75rem] shadow-xl`}
    >
      {children}
    </div>
  )
}

export function BridgeContainer({ children }) {
  const space = 'p-2 grid gap-2'
  const bgColor = 'bg-zinc-50 dark:bg-zinc-800'
  const borderColor = 'border border-zinc-300 dark:border-transparent'

  return (
    <section className={`${space} ${bgColor} ${borderColor} rounded-md`}>
      {children}
    </section>
  )
}

export function BridgeCardTokenInput({ children }) {
  const space = 'flex items-center gap-4 p-2 rounded-md'
  const bgColor = 'bg-white dark:bg-inherit'
  const borderColor = 'border border-zinc-200 dark:border-zinc-700'

  return <div className={`${space} ${bgColor} ${borderColor}`}>{children}</div>
}

export function BridgeTokenSelector({
  dataTestId,
  token,
  placeholder,
  onClick,
}) {
  const space = 'p-2 rounded flex-none flex items-center gap-2'
  const bgColor = 'bg-inherit dark:bg-zinc-700'
  const bgHover = getMenuItemHoverBgForCoin(token?.color)
  const borderColor = `border border-transparent`
  const borderHover = getBorderStyleForCoinHover(token?.color)

  return (
    <button
      data-test-id={dataTestId}
      className={`text-lg ${space} ${bgColor} ${bgHover} ${borderColor} ${borderHover}`}
      onClick={onClick}
    >
      {token && (
        <img
          src={token?.icon?.src ?? ''}
          alt={token?.symbol ?? ''}
          className="w-6 h-6"
        />
      )}
      {token?.symbol ?? placeholder}
      <DropDownArrowSvg />
    </button>
  )
}
