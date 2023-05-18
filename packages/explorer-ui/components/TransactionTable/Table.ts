import { TableHeader} from "./TableHeader";
import { TableBody} from "./TableBody";

export function Table({header, body}) {
  return (
    // @ts-expect-error TS(2304): Cannot find name 'div'.
    <div className="pb-2 px-4 sm:px-6 lg:px-8">
      // @ts-expect-error TS(2304): Cannot find name 'div'.
      <div className="mt-8 flex flex-col">
        // @ts-expect-error TS(2304): Cannot find name 'div'.
        <div className="-my-2 -mx-4 overflow-x-auto sm:-mx-6 lg:-mx-8">
          // @ts-expect-error TS(2304): Cannot find name 'div'.
          <div className="inline-block min-w-full py-2 align-middle">
            // @ts-expect-error TS(2304): Cannot find name 'div'.
            <div className="overflow-hidden shadow-sm ring-1 ring-black ring-opacity-5">
              // @ts-expect-error TS(2304): Cannot find name 'table'.
              <table className="min-w-full">
                // @ts-expect-error TS(2749): 'TableHeader' refers to a value, but is being used... Remove this comment to see the full error message
                <TableHeader headers={header} />
                // @ts-expect-error TS(2304): Cannot find name 'rows'.
                <TableBody rows={body} />
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  )

}
