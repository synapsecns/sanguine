/*
@NOTE: super repetitive code here is actually necessary...
*/

interface StyleLookUpInterface {
  [key: string]: string
}
interface StyleGapLookUpInterface {
  [key: string]: { [key: string]: string }
}

const GRID_COLS_LOOKUP: StyleLookUpInterface = {
  0: 'grid-cols-0 ',
  1: 'grid-cols-1 ',
  2: 'grid-cols-2 ',
  3: 'grid-cols-3 ',
  4: 'grid-cols-4 ',
  5: 'grid-cols-5 ',
  6: 'grid-cols-6 ',
  7: 'grid-cols-7 ',
  8: 'grid-cols-8 ',
  9: 'grid-cols-9 ',
  10: 'grid-cols-10 ',
  11: 'grid-cols-11 ',
  12: 'grid-cols-12 ',
  24: 'grid-cols-24 ',
}

const SM_GRID_COLS_LOOKUP: StyleLookUpInterface = {
  0: 'sm:grid-cols-0 ',
  1: 'sm:grid-cols-1 ',
  2: 'sm:grid-cols-2 ',
  3: 'sm:grid-cols-3 ',
  4: 'sm:grid-cols-4 ',
  5: 'sm:grid-cols-5 ',
  6: 'sm:grid-cols-6 ',
  7: 'sm:grid-cols-7 ',
  8: 'sm:grid-cols-8 ',
  9: 'sm:grid-cols-9 ',
  10: 'sm:grid-cols-10 ',
  11: 'sm:grid-cols-11 ',
  12: 'sm:grid-cols-12 ',
  24: 'sm:grid-cols-24 ',
}

const MD_GRID_COLS_LOOKUP: StyleLookUpInterface = {
  0: 'md:grid-cols-0 ',
  1: 'md:grid-cols-1 ',
  2: 'md:grid-cols-2 ',
  3: 'md:grid-cols-3 ',
  4: 'md:grid-cols-4 ',
  5: 'md:grid-cols-5 ',
  6: 'md:grid-cols-6 ',
  7: 'md:grid-cols-7 ',
  8: 'md:grid-cols-8 ',
  9: 'md:grid-cols-9 ',
  10: 'md:grid-cols-10 ',
  11: 'md:grid-cols-11 ',
  12: 'md:grid-cols-12 ',
  24: 'md:grid-cols-24 ',
}

const LG_GRID_COLS_LOOKUP: StyleLookUpInterface = {
  0: 'lg:grid-cols-0 ',
  1: 'lg:grid-cols-1 ',
  2: 'lg:grid-cols-2 ',
  3: 'lg:grid-cols-3 ',
  4: 'lg:grid-cols-4 ',
  5: 'lg:grid-cols-5 ',
  6: 'lg:grid-cols-6 ',
  7: 'lg:grid-cols-7 ',
  8: 'lg:grid-cols-8 ',
  9: 'lg:grid-cols-9 ',
  10: 'lg:grid-cols-10 ',
  11: 'lg:grid-cols-11 ',
  12: 'lg:grid-cols-12 ',
  24: 'lg:grid-cols-24 ',
}

const XL_GRID_COLS_LOOKUP: StyleLookUpInterface = {
  0: 'xl:grid-cols-0 ',
  1: 'xl:grid-cols-1 ',
  2: 'xl:grid-cols-2 ',
  3: 'xl:grid-cols-3 ',
  4: 'xl:grid-cols-4 ',
  5: 'xl:grid-cols-5 ',
  6: 'xl:grid-cols-6 ',
  7: 'xl:grid-cols-7 ',
  8: 'xl:grid-cols-8 ',
  9: 'xl:grid-cols-9 ',
  10: 'xl:grid-cols-10 ',
  11: 'xl:grid-cols-11 ',
  12: 'xl:grid-cols-12 ',
  24: 'xl:grid-cols-24 ',
}

const GAP_LOOKUP: StyleLookUpInterface = {
  0: 'gap-0 ',
  1: 'gap-1 ',
  2: 'gap-2 ',
  3: 'gap-3 ',
  4: 'gap-4 ',
  5: 'gap-5 ',
  6: 'gap-6 ',
  8: 'gap-8 ',
  10: 'gap-10 ',
  12: 'gap-12 ',
  16: 'gap-16 ',
  20: 'gap-20 ',
  24: 'gap-24 ',
}

const SIZE_GAP_LOOKUP: StyleGapLookUpInterface = {
  xs: {
    0: 'gap-0 ',
    1: 'gap-1 ',
    2: 'gap-2 ',
    3: 'gap-3 ',
    4: 'gap-4 ',
    5: 'gap-5 ',
    6: 'gap-6 ',
    8: 'gap-8 ',
    10: 'gap-10 ',
    12: 'gap-12 ',
    16: 'gap-16 ',
    20: 'gap-20 ',
    24: 'gap-24 ',
  },
  sm: {
    0: 'sm:gap-0 ',
    1: 'sm:gap-1 ',
    2: 'sm:gap-2 ',
    3: 'sm:gap-3 ',
    4: 'sm:gap-4 ',
    5: 'sm:gap-5 ',
    6: 'sm:gap-6 ',
    8: 'sm:gap-8 ',
    10: 'sm:gap-10 ',
    12: 'sm:gap-12 ',
    16: 'sm:gap-16 ',
    20: 'sm:gap-20 ',
    24: 'sm:gap-24 ',
  },
  md: {
    0: 'md:gap-0 ',
    1: 'md:gap-1 ',
    2: 'md:gap-2 ',
    3: 'md:gap-3 ',
    4: 'md:gap-4 ',
    5: 'md:gap-5 ',
    6: 'md:gap-6 ',
    8: 'md:gap-8 ',
    10: 'md:gap-10 ',
    12: 'md:gap-12 ',
    16: 'md:gap-16 ',
    20: 'md:gap-20 ',
    24: 'md:gap-24 ',
  },
  lg: {
    0: 'lg:gap-0 ',
    1: 'lg:gap-1 ',
    2: 'lg:gap-2 ',
    3: 'lg:gap-3 ',
    4: 'lg:gap-4 ',
    5: 'lg:gap-5 ',
    6: 'lg:gap-6 ',
    8: 'lg:gap-8 ',
    10: 'lg:gap-10 ',
    12: 'lg:gap-12 ',
    16: 'lg:gap-16 ',
    20: 'lg:gap-20 ',
    24: 'lg:gap-24 ',
  },
}

const GAP_X_LOOKUP: StyleLookUpInterface = {
  0: 'gap-x-0 ',
  1: 'gap-x-1 ',
  2: 'gap-x-2 ',
  3: 'gap-x-3 ',
  4: 'gap-x-4 ',
  5: 'gap-x-5 ',
  6: 'gap-x-6 ',
  8: 'gap-x-8 ',
  10: 'gap-x-10 ',
  12: 'gap-x-12 ',
  16: 'gap-x-16 ',
  20: 'gap-x-20 ',
  24: 'gap-x-24 ',
}

const GAP_Y_LOOKUP: StyleLookUpInterface = {
  0: 'gap-y-0 ',
  1: 'gap-y-1 ',
  2: 'gap-y-2 ',
  3: 'gap-y-3 ',
  4: 'gap-y-4 ',
  5: 'gap-y-5 ',
  6: 'gap-y-6 ',
  8: 'gap-y-8 ',
  10: 'gap-y-10 ',
  12: 'gap-y-12 ',
  16: 'gap-y-16 ',
  20: 'gap-y-20 ',
  24: 'gap-y-24 ',
}

export default function Grid({
  children,
  cols,
  gap,
  gapX,
  gapY,
  className: providedClassName,
  ...props
}: {
  children?: any
  cols?: any
  gap?: any
  gapX?: string
  gapY?: string
  // as?: string
  className?: string
}) {
  let novelClassName = 'grid '

  const { xs, sm, md, lg, xl } = cols ?? {}

  if (cols) {
    if (xs) {
      novelClassName += GRID_COLS_LOOKUP[xs]
    }
    if (sm ?? false) {
      novelClassName += SM_GRID_COLS_LOOKUP[sm]
    }
    if (md ?? false) {
      novelClassName += MD_GRID_COLS_LOOKUP[md]
    }
    if (lg ?? false) {
      novelClassName += LG_GRID_COLS_LOOKUP[lg]
    }
    if (xl ?? false) {
      novelClassName += XL_GRID_COLS_LOOKUP[xl]
    }
  } else {
    novelClassName += GRID_COLS_LOOKUP[12]
  }

  if (!((gap ?? false) || (gapX ?? false) || (gapY ?? false))) {
    novelClassName += GAP_LOOKUP[0]
  }

  if (gap) {
    if (Number.isInteger(gap)) {
      novelClassName += GAP_LOOKUP[gap]
    } else {
      for (const [screenSize, gapSize] of gap.entries()) {
        novelClassName += SIZE_GAP_LOOKUP[screenSize][gapSize]
      }
    }
  }

  if (gapX) {
    novelClassName += GAP_X_LOOKUP[gapX]
  }

  if (gapY) {
    novelClassName += GAP_Y_LOOKUP[gapY]
  }

  if (providedClassName ?? false) {
    novelClassName = `${novelClassName} ${providedClassName} `
  }

  return (
    <div className={novelClassName} {...props}>
      {children ? children : null}
    </div>
  )
}
