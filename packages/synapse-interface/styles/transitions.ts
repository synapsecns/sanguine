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

/**
 * `transform-gpu transform` were used on the *To/*From classes
 * but it appears it doesnt have support in conjunction with opacity
 * in firefox 72-102, 109 and other mis versions easiest fix is to
 * just not use it for now till either tw or firefox fixes it
 * whoever wrote firefox deserves to get shot, then brought back
 * to life and then shot again
 */
export const COIN_SLIDE_OVER_PROPS: TransitionClassesProps = {
  appear: true,
  unmount: true,
  enter: 'duration-0 transition-opacity ease-out', //duration-25
  enterFrom: ' opacity-0',
  enterTo: ' opacity-100',
  leave: 'duration-0 transition-opacity ease-out', //duration-25
  leaveFrom: ' opacity-100',
  leaveTo: ' opacity-0 hidden',
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
