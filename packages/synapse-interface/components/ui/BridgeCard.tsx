/* TODOs
 * Lift margin value up to parent
 * Remove need for popoverDependencies styles
 * Adjust interior elements to allow for single p-4 padding value
 */

const space = 'px-4 pt-4 pb-2 mt-5'
const bgColor = 'bg-zinc-100 dark:bg-zinc-900/95'
const popoverDependencies = 'overflow-hidden transform'

export default function BridgeCard({ children }) {
  return (
    <div
      className={`${space} ${bgColor} ${popoverDependencies} rounded-[.75rem] shadow-xl`}
    >
      {children}
    </div>
  )
}
