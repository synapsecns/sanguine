const COL_SPAN_LOOKUP = {
  auto: 'col-auto ',
  0: 'hidden ',
  1: 'col-span-1 ',
  2: 'col-span-2 ',
  3: 'col-span-3 ',
  4: 'col-span-4 ',
  5: 'col-span-5 ',
  6: 'col-span-6 ',
  7: 'col-span-7 ',
  8: 'col-span-8 ',
  9: 'col-span-9 ',
  10: 'col-span-10 ',
  11: 'col-span-11 ',
  12: 'col-span-12 ',
  24: 'col-span-full ',
  full: 'col-span-full ',
}

const SM_COL_SPAN_LOOKUP = {
  auto: 'sm:col-auto ',
  0: 'sm:hidden ',
  1: 'sm:col-span-1 ',
  2: 'sm:col-span-2 ',
  3: 'sm:col-span-3 ',
  4: 'sm:col-span-4 ',
  5: 'sm:col-span-5 ',
  6: 'sm:col-span-6 ',
  7: 'sm:col-span-7 ',
  8: 'sm:col-span-8 ',
  9: 'sm:col-span-9 ',
  10: 'sm:col-span-10 ',
  11: 'sm:col-span-11 ',
  12: 'sm:col-span-12 ',
  24: 'sm:col-span-full ',
  full: 'sm:col-span-full ',
}

const MD_COL_SPAN_LOOKUP = {
  auto: 'md:col-auto ',
  0: 'md:hidden ',
  1: 'md:col-span-1 ',
  2: 'md:col-span-2 ',
  3: 'md:col-span-3 ',
  4: 'md:col-span-4 ',
  5: 'md:col-span-5 ',
  6: 'md:col-span-6 ',
  7: 'md:col-span-7 ',
  8: 'md:col-span-8 ',
  9: 'md:col-span-9 ',
  10: 'md:col-span-10 ',
  11: 'md:col-span-11 ',
  12: 'md:col-span-12 ',
  24: 'md:col-span-full ',
  full: 'md:col-span-full ',
}

const LG_COL_SPAN_LOOKUP = {
  auto: 'lg:col-auto ',
  0: 'lg:hidden ',
  1: 'lg:col-span-1 ',
  2: 'lg:col-span-2 ',
  3: 'lg:col-span-3 ',
  4: 'lg:col-span-4 ',
  5: 'lg:col-span-5 ',
  6: 'lg:col-span-6 ',
  7: 'lg:col-span-7 ',
  8: 'lg:col-span-8 ',
  9: 'lg:col-span-9 ',
  10: 'lg:col-span-10 ',
  11: 'lg:col-span-11 ',
  12: 'lg:col-span-12 ',
  24: 'lg:col-span-full ',
  full: 'lg:col-span-full ',
}

const XL_COL_SPAN_LOOKUP = {
  auto: 'xl:col-auto ',
  0: 'xl:hidden ',
  1: 'xl:col-span-1 ',
  2: 'xl:col-span-2 ',
  3: 'xl:col-span-3 ',
  4: 'xl:col-span-4 ',
  5: 'xl:col-span-5 ',
  6: 'xl:col-span-6 ',
  7: 'xl:col-span-7 ',
  8: 'xl:col-span-8 ',
  9: 'xl:col-span-9 ',
  10: 'xl:col-span-10 ',
  11: 'xl:col-span-11 ',
  12: 'xl:col-span-12 ',
  24: 'xl:col-span-full ',
  full: 'xl:col-span-full ',
}

export default function Col({
  children,
  xs,
  sm,
  md,
  lg,
  xl,
  className: providedClassName,
  ...props
}) {
  let novelClassName = ''

  if (xs ?? false) {
    novelClassName += COL_SPAN_LOOKUP[xs]
  }
  if (sm ?? false) {
    novelClassName += SM_COL_SPAN_LOOKUP[sm]
  }
  if (md ?? false) {
    novelClassName += MD_COL_SPAN_LOOKUP[md]
  }
  if (lg ?? false) {
    novelClassName += LG_COL_SPAN_LOOKUP[lg]
  }
  if (xl ?? false) {
    novelClassName += XL_COL_SPAN_LOOKUP[xl]
  }

  if (providedClassName ?? false) {
    novelClassName += ' #{providedClassName} '
  }

  return (
    <div className={novelClassName} {...props}>
      {children}
    </div>
  )
}
