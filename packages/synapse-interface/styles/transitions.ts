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
  enter: 'duration-25 transition-opacity ease-out',
  enterFrom: 'transform-gpu transform opacity-0',
  enterTo: 'transform-gpu transform opacity-100',
  leave: 'duration-25 transition-opacity ease-out',
  leaveFrom: 'transform-gpu transform opacity-100',
  leaveTo: 'transform-gpu transform opacity-0 hidden',
}

export const SECTION_TRANSITION_PROPS: TransitionClassesProps = {
  appear: true,
  unmount: true,
  enter: 'transition duration-75 ease-out',
  enterFrom: 'transform-gpu scale-y-0',
  enterTo: 'transform-gpu scale-y-100 opacity-100',
  leave: 'transition duration-75 ease-out ',
  leaveFrom: 'transform-gpu scale-y-100 opacity-100',
  leaveTo: 'transform-gpu scale-y-0 hidden',
  className: 'origin-top -mx-0 md:-mx-6',
}

export const TRANSITION_PROPS: TransitionClassesProps = {
  ...COIN_SLIDE_OVER_PROPS,
  className: `
    origin-top absolute
    transition-all
    w-full h-full
    backdrop-blur-lg
    z-20 rounded-lg
    left-0
  `,
}
