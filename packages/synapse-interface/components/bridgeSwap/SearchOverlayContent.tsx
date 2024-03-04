import { CloseButton } from "@/components/buttons/CloseButton"
import { NoSearchResultsFound } from "@/components/bridgeSwap/NoSearchResultsFound"
import { SlideSearchBox } from "@/components/bridgeSwap/SlideSearchBox"


const PLACEHOLDERS = {
  token: "Filter by symbol, contract, or name...",
  chain: 'Filter by chain name, id, or native currency'
}

export function SearchOverlayContent({
  overlayRef,
  searchStr,
  onSearch,
  onClose,
  type,
  children
}: {
  overlayRef: React.RefObject<HTMLDivElement>
  searchStr: string
  onSearch: (searchStr: string) => void
  onClose: () => void
  type: 'token' | 'chain'
  children: React.ReactNode
}) {
  return (
    <div
      ref={overlayRef}
      data-test-id="search-overlay-content"
      className="max-h-full pb-4 overflow-auto scrollbar-hide"
    >
      <div className="z-10 w-full px-2 ">
        <div className="relative flex items-center my-2 font-medium">
          <SlideSearchBox
            placeholder={PLACEHOLDERS[type]}
            searchStr={searchStr}
            onSearch={onSearch}
          />
          <CloseButton onClick={onClose} />
        </div>
      </div>
      <div data-test-id="search-overlay-content-children" className="px-2 pt-2 pb-8 md:px-2">
        {children}
        <NoSearchResultsFound searchStr={searchStr} type={type} />
      </div>
    </div>
  )
}

