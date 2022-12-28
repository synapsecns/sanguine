import {ChevronDoubleLeftIcon, ChevronLeftIcon, ChevronRightIcon,} from '@heroicons/react/outline'

import {PAGINATION_COUNT} from '@constants'

export function Pagination({
  page,
  resetPage,
  prevPage,
  nextPage,
  totalCount,
}) {
  let showNextPage = !totalCount
    ? true
    : totalCount - page * PAGINATION_COUNT > 0

  return (
    <div className="flex items-center justify-center mt-3 text-sm text-gray-500">
      <button
        className={page === 1 ? `pointer-events-none opacity-50` : ''}
        onClick={resetPage}
      >
        <ChevronDoubleLeftIcon className="w-5 h-5" strokeWidth={1} />
      </button>
      <button
        className={page === 1 ? `pointer-events-none opacity-50` : ''}
        onClick={prevPage}
      >
        <ChevronLeftIcon className="w-5 h-5" strokeWidth={1} />
      </button>
      <span className="ml-2 font-light">Page {page}</span>
      <button
        onClick={nextPage}
        className={showNextPage ? '' : 'pointer-events-none opacity-50'}
      >
        <ChevronRightIcon className="w-5 h-5" strokeWidth={1} />
      </button>
    </div>
  )
}
