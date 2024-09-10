import { joinClassNames } from '@/utils/joinClassNames'

export function BridgeSectionContainer({ children }) {
  const classNames = {
    space: 'grid gap-2 p-2 rounded-md',
    background: 'bg-zinc-50 dark:bg-bgLight', // TODO: Remove
    // background: 'bg-zinc-50 dark:bg-zinc-800',
    borderColor: 'border border-zinc-300 dark:border-transparent',
  }

  return <section className={joinClassNames(classNames)}>{children}</section>
}
