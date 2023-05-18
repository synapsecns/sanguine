import {ChevronDoubleLeftIcon, ChevronDoubleRightIcon, ChevronLeftIcon, ChevronRightIcon,} from '@heroicons/react/outline'

import {useRouter} from "next/router";
import Link from 'next/link';


export function Pagination({
}) {
  const router = useRouter()
  let { p } = router.query
  // @ts-expect-error TS(2322): Type 'string | string[] | 1' is not assignable to ... Remove this comment to see the full error message
  p = p ?? 1

  return (
    // @ts-expect-error TS(2304): Cannot find name 'div'.
    <div className="flex items-center justify-center mt-3 text-sm text-gray-500">
      // @ts-expect-error TS(2749): 'Link' refers to a value, but is being used as a t... Remove this comment to see the full error message
      <Link href={{
        pathname: router.basePath,
        // @ts-expect-error TS(2304): Cannot find name 'query'.
        query: {...router.query, p: 1}
      // @ts-expect-error TS(2630): Cannot assign to 'scroll' because it is a function... Remove this comment to see the full error message
      }} scroll={false} className={p === 1 ? `pointer-events-none opacity-50` : ''}>
          // @ts-expect-error TS(2749): 'ChevronDoubleLeftIcon' refers to a value, but is ... Remove this comment to see the full error message
          <ChevronDoubleLeftIcon className="w-5 h-5" strokeWidth={1} />
      </Link>
      // @ts-expect-error TS(2304): Cannot find name 'href'.
      <Link href={{
        // @ts-expect-error TS(2304): Cannot find name 'router'.
        pathname: router.basePath,
        // @ts-expect-error TS(2304): Cannot find name 'query'.
        query: {...router.query, p: p-1}
      // @ts-expect-error TS(2630): Cannot assign to 'scroll' because it is a function... Remove this comment to see the full error message
      }} scroll={false}
            // @ts-expect-error TS(2304): Cannot find name 'className'.
            className={p === 1 ? `pointer-events-none opacity-50` : ''}>
        // @ts-expect-error TS(2749): 'ChevronLeftIcon' refers to a value, but is being ... Remove this comment to see the full error message
        <ChevronLeftIcon className="w-5 h-5" strokeWidth={1} />
      </Link>
      // @ts-expect-error TS(2304): Cannot find name 'span'.
      <span className="font-light">Page {p}</span>
      // @ts-expect-error TS(2304): Cannot find name 'href'.
      <Link href={{
        // @ts-expect-error TS(2304): Cannot find name 'router'.
        pathname: router.basePath,
        // @ts-expect-error TS(2304): Cannot find name 'query'.
        query: {...router.query, p: Number(p)+1}
      // @ts-expect-error TS(2630): Cannot assign to 'scroll' because it is a function... Remove this comment to see the full error message
      }} scroll={false} className={``}>
        // @ts-expect-error TS(2749): 'ChevronRightIcon' refers to a value, but is being... Remove this comment to see the full error message
        <ChevronRightIcon className="w-5 h-5" strokeWidth={1} />
      </Link>
    </div>
  )
}
