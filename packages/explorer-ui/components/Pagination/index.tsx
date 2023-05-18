import {
  ChevronDoubleLeftIcon,
  ChevronDoubleRightIcon,
  ChevronLeftIcon,
  ChevronRightIcon,
} from '@heroicons/react/outline'
import { useRouter } from 'next/router'
import Link from 'next/link'

export function Pagination({}) {
  const router = useRouter()
  let { p } = router.query
  // @ts-ignore
  p = p ?? 1

  return (
    <div className="flex items-center justify-center mt-3 text-sm text-gray-500">
      <Link
        href={{
          pathname: router.basePath,
          query: { ...router.query, p: 1 },
        }}
        scroll={false}
        // @ts-ignore
        className={p === 1 ? `pointer-events-none opacity-50` : ''}
      >
        <ChevronDoubleLeftIcon className="w-5 h-5" strokeWidth={1} />
      </Link>
      <Link
        href={{
          pathname: router.basePath,
          // @ts-ignore
          query: { ...router.query, p: p - 1 },
        }}
        scroll={false}
        // @ts-ignore
        className={p === 1 ? `pointer-events-none opacity-50` : ''}
      >
        <ChevronLeftIcon className="w-5 h-5" strokeWidth={1} />
      </Link>
      <span className="font-light">Page {p}</span>
      <Link
        href={{
          pathname: router.basePath,
          query: { ...router.query, p: Number(p) + 1 },
        }}
        scroll={false}
        className={``}
      >
        <ChevronRightIcon className="w-5 h-5" strokeWidth={1} />
      </Link>
    </div>
  )
}
