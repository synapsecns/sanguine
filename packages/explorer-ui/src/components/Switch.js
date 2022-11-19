import { Switch } from '@headlessui/react'

function classNames(...classes) {
  return classes.filter(Boolean).join(' ')
}

export const LabelSwitch = ({ text, enabled, setEnabled }) => {
  return (
    <Switch.Group as="div" className="flex items-center justify-end">
      <span className="mr-4 text-gray-400">{text}</span>
      <Switch
        checked={enabled}
        onChange={setEnabled}
        className={classNames(
          enabled ? 'bg-indigo-600' : 'bg-gray-900',
          'relative inline-flex flex-shrink-0 h-6 w-11 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none',
        )}
      >
        <span
          aria-hidden="true"
          className={classNames(
            enabled ? 'translate-x-5' : 'translate-x-0',
            'pointer-events-none inline-block h-5 w-5 rounded-full bg-gray-400 shadow transform ring-0 transition ease-in-out duration-200',
          )}
        />
      </Switch>
    </Switch.Group>
  )
}
