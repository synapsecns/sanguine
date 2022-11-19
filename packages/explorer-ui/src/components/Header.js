import { Popover, Transition } from '@headlessui/react'
import { ExternalLinkIcon, MenuIcon, XIcon } from '@heroicons/react/outline'
import { Fragment } from 'react'
import SynapseLogoSvg from './icons/SynapseLogoSvg'

export const SynapseTitleLogo = ({ showText }) => {
  return (
    <div>
      <a href={'/'} className="flex">
        <div className="flex items-center flex-shrink-0 py-1 ">
          <div className="mr-2">
            <SynapseLogoSvg />
          </div>
          <span
            className={`${
              showText ? '' : 'hidden'
            } pl-2 text-2xl font-normal tracking-wider text-transparent transition-all transform  bg-clip-text bg-gradient-to-r from-purple-600 to-blue-600 active:from-purple-700 active:to-blue-700 hover:animate-pulse`}
          >
            Synapse
          </span>
        </div>
      </a>
    </div>
  )
}

export const Header = () => {
  return (
    <Popover className="relative">
      <div className="bg-[#111827] flex items-center justify-between px-4 py-1 border-b border-gray-300 sm:px-6 md:justify-start md:space-x-10 dark:border-gray-800 min-h-[64px]">
        <SynapseTitleLogo showText={false} />
        <div className="hidden md:block md:ml-6">
          <div className="flex space-x-4">
            <a
              href="/"
              className="px-3 py-2 text-base font-normal text-gray-400 rounded-md hover:bg-gray-800 hover:text-gray-200"
            >
              Overview
            </a>
            <a
              href="/bridge"
              className="px-3 py-2 text-base font-normal text-gray-400 rounded-md hover:bg-gray-800 hover:text-gray-200"
            >
              Bridge Statistics
            </a>
            <a
              href="/pool"
              className="px-3 py-2 text-base font-normal text-gray-400 rounded-md hover:bg-gray-800 hover:text-gray-200"
            >
              Pool Statistics
            </a>
            <a
              href="/fee"
              className="px-3 py-2 text-base font-normal text-gray-400 rounded-md hover:bg-gray-800 hover:text-gray-200"
            >
              Fee Statistics
            </a>
            <div className="flex items-center px-3 py-2 text-base font-normal text-gray-400 rounded-md hover:bg-gray-800 hover:text-gray-200">
              <a href="https://synapseprotocol.com" className="">
                Synapse Protocol
              </a>
              <ExternalLinkIcon className="w-5 h-5 ml-2 text-gray-400" aria-hidden="true" />
            </div>
          </div>
        </div>
        <div className="-my-2 -mr-2 md:hidden">
          <Popover.Button className="inline-flex items-center justify-center p-2 text-gray-400 rounded-lg hover:text-gray-500 hover:bg-gray-100 dark:hover:text-gray-400 dark:hover:bg-gray-800 focus:outline-none">
            <span className="sr-only">Open menu</span>
            <MenuIcon className="w-6 h-6" aria-hidden="true" />
          </Popover.Button>
        </div>
      </div>
      <Transition
        as={Fragment}
        enter="duration-100 ease-out"
        enterFrom=" opacity-0"
        enterTo=" opacity-100"
        leave="duration-75 ease-in"
        leaveFrom=" opacity-100"
        leaveTo=" opacity-0"
      >
        <Popover.Panel focus className="absolute inset-x-0 top-0 z-10 transition origin-top-right transform md:hidden">
          <div className="h-full min-h-full bg-white divide-y shadow-lg ring-1 ring-opacity-5 dark:bg-gray-800 divide-gray-50 dark:divide-gray-600">
            <div className="px-4 pt-2 pb-6">
              <div className="flex items-center justify-between">
                <SynapseTitleLogo showText={true} />
                <div className="-mr-2">
                  <Popover.Button className="inline-flex items-center justify-center p-2 text-gray-400 rounded-lg hover:text-gray-500 hover:bg-gray-100 dark:hover:text-gray-400 dark:hover:bg-gray-900 focus:outline-none">
                    <span className="sr-only">Close menu</span>
                    <XIcon className="w-6 h-6" aria-hidden="true" />
                  </Popover.Button>
                </div>
              </div>
              <div className="mt-4">
                <div>
                  <div className="flex px-3 py-2 text-base font-normal text-gray-400 rounded-md hover:bg-gray-800 hover:text-gray-200">
                    <a href="/">Overview</a>
                  </div>
                  <div className="flex px-3 py-2 text-base font-normal text-gray-400 rounded-md hover:bg-gray-800 hover:text-gray-200">
                    <a href="/bridge">Bridge Statistics</a>
                  </div>
                  <div className="flex px-3 py-2 text-base font-normal text-gray-400 rounded-md hover:bg-gray-800 hover:text-gray-200">
                    <a href="/pool">Pool Statistics</a>
                  </div>
                  <div className="flex px-3 py-2 text-base font-normal text-gray-400 rounded-md hover:bg-gray-800 hover:text-gray-200">
                    <a href="/fee">Fee Statistics</a>
                  </div>
                  <div className="flex items-center px-3 py-2 text-base font-normal text-gray-400 rounded-md hover:bg-gray-800 hover:text-gray-200">
                    <a href="https://synapseprotocol.com">Synapse Protocol</a>
                    <ExternalLinkIcon className="w-5 h-5 ml-2 text-gray-400" aria-hidden="true" />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </Popover.Panel>
      </Transition>
    </Popover>
  )
}
