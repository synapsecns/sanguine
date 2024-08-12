import { TableHeader } from './TableHeader'
import { TableBody } from './TableBody'

export const Table = ({ header, body }) => {
  return (
    <div className="px-4 pb-2 sm:px-6 lg:px-8">
      <div className="flex flex-col">
        <div className="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
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
