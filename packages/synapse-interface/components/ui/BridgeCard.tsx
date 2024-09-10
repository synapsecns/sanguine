import React from 'react'
import { joinClassNames } from '@/utils/joinClassNames'
import { type BridgeCardTypes } from '@/components/ui/types'

export function BridgeCard({ bridgeRef, children }: BridgeCardTypes) {
  /* TODOs
   * Lift margin value up to parent
   * Remove need for popoverDependencies styles (in progress)
   */
  const classNames = {
    grid: 'grid gap-2',
    space: 'p-3 mt-5 rounded-[.75rem]',
    background: 'bg-zinc-100 dark:bg-bgBase', // TODO: Remove
    // background: 'bg-zinc-100 dark:bg-zinc-900/95 shadow-xl',
    // popoverDependencies: 'overflow-hidden transform',
  }

  return (
    <div ref={bridgeRef} className={joinClassNames(classNames)}>
      {children}
    </div>
  )
}
