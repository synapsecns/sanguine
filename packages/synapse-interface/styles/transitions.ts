interface TransitionClasses {
  enter?: string
  enterFrom?: string
  enterTo?: string
  entered?: string
  leave?: string
  leaveFrom?: string
  leaveTo?: string
}

export interface TransitionClassesProps extends TransitionClasses {
  appear?: boolean
  unmount?: boolean
  className?: string
}

export const COIN_SLIDE_OVER_PROPS: TransitionClassesProps = {
  appear: true,
  unmount: true,
  enter: 'duration-200 transition ease-out',
  enterFrom: 'transform-gpu transform-y-200 opacity-20',
  entered: 'h-100',
  enterTo: 'transform-gpu transform-y-0 opacity-100',
  leave: 'duration-150 transition ease-out',
  leaveFrom: 'transform-gpu transform-y-0 opacity-100',
  leaveTo: 'transform-gpu transform-y-200 opacity-0',
}

export const TEST_COIN_SLIDE_OVER_PROPS: TransitionClassesProps = {
  appear: true,
  unmount: true,
  enter: 'transition duration-150 ease-out',
  enterFrom: 'transform-gpu -scale-y-30 ',
  enterTo: 'transform-gpu scale-y-100 opacity-100',
  leave: 'transition duration-150 ease-out ',
  leaveFrom: 'transform-gpu scale-y-100 opacity-100',
  leaveTo: 'transform-gpu -scale-y-30 opacity-50',
}

export const SECTION_TRANSITION_PROPS: TransitionClassesProps = {
  enter: 'transition duration-75 ease-out',
  enterFrom: 'transform-gpu scale-y-0 ',
  enterTo: 'transform-gpu scale-y-100 opacity-100',
  leave: 'transition duration-75 ease-out ',
  leaveFrom: 'transform-gpu scale-y-100 opacity-100',
  leaveTo: 'transform-gpu scale-y-0 ',
  className: 'origin-top -mx-0 md:-mx-6',
}

export const TRANSITION_PROPS: TransitionClassesProps = {
  ...COIN_SLIDE_OVER_PROPS,
  className: `
    origin-top absolute
    transition-all
    w-full
    md:w-[95%] md:h-[95%]
    -ml-0 md:-ml-3
    md:mt-3
    bg-bgBase
    z-20 rounded-3xl
  `,
}

export const TEST_TRANSITION_PROPS: TransitionClassesProps = {
  ...TEST_COIN_SLIDE_OVER_PROPS,
  className: `
    origin-bottom absolute
    w-full h-full
    md:w-[95%] md:h-[95%]
    -ml-0 md:-ml-3
    md:mt-3
    bg-bgBase
    z-20 rounded-3xl
  `,
}
