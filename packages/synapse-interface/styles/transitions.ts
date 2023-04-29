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
  enterFrom: 'transform-gpu transform translate-y-50 opacity-0',
  enterTo: 'transform-gpu transform translate-y-0 opacity-100',
  leave: 'duration-200 transition ease-out',
  leaveFrom: 'transform-gpu transform translate-y-0  opacity-100',
  leaveTo: 'transform-gpu transform translate-y-100 opacity-0',
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
    w-full h-full
    md:w-[95%] md:h-[95%]
    -ml-0 md:-ml-3
    md:mt-3
    bg-bgBase
    z-20 rounded-3xl
  `,
}
