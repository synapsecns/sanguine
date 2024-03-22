export const ListSectionWrapper = ({ sectionKey, children }) => (
  <section key={sectionKey} className="rounded bg-bgBase first:bg-bgLight">
    <header
      className="sticky top-0 z-10 p-2 text-sm cursor-default text-secondary bg-inherit"
      onClick={(e) => e.stopPropagation()}
    >
      {sectionKey}
    </header>
    {children}
  </section>
)
