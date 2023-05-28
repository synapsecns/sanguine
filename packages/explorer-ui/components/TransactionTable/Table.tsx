import { TableHeader } from './TableHeader'
import { TableBody } from './TableBody'

export function Table({ header, body }) {
  return (
    <div className="pb-2 px-4 sm:px-6 lg:px-8">
      <div className="mt-8 flex flex-col">
        <div className="-my-2 -mx-4 overflow-x-auto sm:-mx-6 lg:-mx-8">
          <div className="inline-block min-w-full py-2 align-middle">
            <div className="overflow-hidden shadow-sm ring-1 ring-black ring-opacity-5">
              <table className="min-w-full">
                <TableHeader headers={header} />
                <TableBody rows={body} />
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
