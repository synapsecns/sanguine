export const COIN_SLIDE_OVER_PROPS = {
  appear: true,
  unmount: true,
  enter: 'transition duration-150 ease-out',
  enterFrom: 'transform-gpu scale-y-0 ',
  enterTo: 'transform-gpu scale-y-100 opacity-100',
  leave: 'transition duration-150 ease-out ',
  leaveFrom: 'transform-gpu scale-y-100 opacity-100',
  leaveTo: 'transform-gpu scale-y-0 opacity-50',
}