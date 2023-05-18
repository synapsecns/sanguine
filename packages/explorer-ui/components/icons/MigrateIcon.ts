import {LibraryIcon} from '@heroicons/react/outline'


export default function MigrateIcon({ className, ...props }) {
  return (
    // @ts-expect-error TS(2749): 'LibraryIcon' refers to a value, but is being used... Remove this comment to see the full error message
    <LibraryIcon
      // @ts-expect-error TS(2349): This expression is not callable.
      className={`h-6 w-6 ${className}`}
      // @ts-expect-error TS(2304): Cannot find name 'props'.
      {...props}
    />
  )
}


