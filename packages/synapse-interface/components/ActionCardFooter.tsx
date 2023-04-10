import { ArrowSmRightIcon, ChartSquareBarIcon } from '@heroicons/react/outline'

import { ANALYTICS_PATH } from '@/constants/urls'

export const ActionCardFooter = ({ link }: { link: string }) => {
  return (
    <div className="flex-wrap items-center justify-between ml-5 mr-5 text-[15px] md:flex lg:flex">
      <div className="flex items-center text-secondaryTextColor">
        <span className="mr-1 opacity-50">Need help? Read </span>
        <a href={link} target="_blank" className="">
          <span className="transition-all duration-75 cursor-pointer hover:text-opacity-100 hover:text-white transform-gpu">
            this guide{' '}
          </span>
        </a>
        <ArrowSmRightIcon className="w-5 h-5" />
      </div>
      <a href={ANALYTICS_PATH} target="_blank">
        <div className="flex items-center text-opacity-50 transition-all duration-75 text-secondaryTextColor hover:text-opacity-100 hover:text-white transform-gpu ">
          <ChartSquareBarIcon className="w-5 h-5 mr-2" />
          <span className="cursor-pointer">Explorer</span>
        </div>
      </a>
    </div>
  )
}
