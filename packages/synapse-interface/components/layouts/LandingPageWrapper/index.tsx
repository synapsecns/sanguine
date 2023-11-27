import { Fragment } from 'react'
import { Popover, Transition } from '@headlessui/react'
import { MenuIcon, XIcon } from '@heroicons/react/outline'
import Grid from '@tw/Grid'
import ForumIcon from '@icons/ForumIcon'
import TwitterIcon from '@icons/TwitterIcon'
import DiscordIcon from '@icons/DiscordIcon'
import TelegramIcon from '@icons/TelegramIcon'
import DocumentTextIcon from '@icons/DocsIcon'
import { Wallet } from '@components/Wallet'

import { SynapseLogoSvg, SynapseLogoWithTitleSvg } from './SynapseLogoSvg'
import { TopBarNavLink } from './TopBarNavLink'
import {
  CONTRACTS_PATH,
  DISCORD_URL,
  DOCS_URL,
  FORUM_URL,
  LANDING_PATH,
  TELEGRAM_URL,
  TWITTER_URL,
  getBuySynUrl,
} from '@/constants/urls'
import { NAVIGATION } from '@/constants/routes'
import { MoreButton } from './MoreButton'
import { PageFooter } from './PageFooter'

export function LandingPageWrapper({ children }: { children: any }) {
  return (
    <div
      className="min-h-screen overflow-x-hidden bg-no-repeat"
      style={{
        background:
          'radial-gradient(23.86% 33.62% at 50.97% 47.88%, rgba(255, 0, 255, 0.04) 0%, rgba(172, 143, 255, 0.04) 100%), #111111',
      }}
    >
      <LandingNav />

      <div
        style={{
          backgroundImage: `url('landingBg.svg')`,
          backgroundSize: '60%',
          backgroundPosition: 'top center',
          backgroundRepeat: 'no-repeat',
        }}
      >
        {children}
      </div>
      <PageFooter />
    </div>
  )
}

export function LandingNav() {
  return (
    <Popover className="relative px-8 pt-6">
      <div className="w-full md:flex-1 md:flex md:items-center md:justify-between">
        <div className="flex items-center w-full py-4 lg:px-20">
          <div className="mr-auto">
            <SynapseTitleLogo showText={true} />
          </div>
          <div className="items-center justify-center -mr-2 sm:flex lg:hidden">
            <Popover.Button
              data-test-id="mobile-navbar-button"
              className={`
                  rounded-md p-2 inline-flex items-center justify-center
                  text-gray-400 hover:text-gray-400 hover:bg-gray-800
                  focus:outline-none
                `}
            >
              <span className="sr-only">Open menu</span>
              <MenuIcon className="w-8 h-8" aria-hidden="true" />
            </Popover.Button>
          </div>
          <Popover.Group
            as="nav"
            className="hidden lg:flex md:justify-space-evenly"
            data-test-id="desktop-nav"
          >
            <TopBarButtons />
          </Popover.Group>
          <div className="justify-end hidden ml-auto lg:flex">
            <div className="flex items-center space-x-2">
              <Wallet />
              <Popover className="relative inline-block">
                {({ open }) => (
                  <>
                    <Popover.Button
                      as="div"
                      onMouseEnter={() => {}}
                      className={`
                      ${open ? 'text-gray-900' : 'text-purple-800'}
                       group  rounded-md inline-flex items-center  hover:text-gray-900 focus:outline-none
                      `}
                    >
                      <MoreButton open={open} />
                    </Popover.Button>
                    <PopoverPanelContainer className="-translate-x-full">
                      <MoreInfoButtons />
                      <SocialButtons />
                    </PopoverPanelContainer>
                  </>
                )}
              </Popover>
            </div>
          </div>
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
        <Popover.Panel
          focus
          className="absolute inset-x-0 top-0 z-10 transition origin-top-right transform"
        >
          <div
            className="h-full min-h-full divide-y divide-gray-600 bg-bgLight"
            // data-test-id="mobile-nav"
          >
            <div className="px-4 pt-1 pb-6">
              <div className="flex items-center justify-between mt-5 ml-3">
                <SynapseTitleLogo showText={false} />
                <div className="-mr-2">
                  <Popover.Button
                    className={`
                        rounded-md p-2 inline-flex items-center justify-center
                        text-gray-400 hover:text-gray-400 hover:bg-gray-900
                        focus:outline-none
                      `}
                  >
                    <span className="sr-only">Close menu</span>
                    <XIcon className="w-6 h-6" aria-hidden="true" />
                  </Popover.Button>
                </div>
              </div>
              <div className="mt-6">
                <Grid
                  cols={{ xs: 1 }}
                  gap={2}
                  className="py-6"
                  data-test-id="mobile-nav"
                >
                  <MobileBarButtons />
                </Grid>
              </div>
            </div>
            <Grid cols={{ xs: 1 }} gap={4} className="px-4 py-4">
              <Wallet />
            </Grid>
          </div>
        </Popover.Panel>
      </Transition>
    </Popover>
  )
}

export function PopoverPanelContainer({
  children,
  className,
}: {
  children: any
  className?: string
}) {
  return (
    <Transition
      as={Fragment}
      enter="transition ease-out duration-200"
      enterFrom="opacity-0 translate-y-1"
      enterTo="opacity-100 translate-y-0"
      leave="transition ease-in duration-150"
      leaveFrom="opacity-100 translate-y-0"
      leaveTo="opacity-0 translate-y-1"
    >
      <Popover.Panel
        className={`
          absolute z-10 left-1/2 transform-gpu
          ${className ?? '-translate-x-1/2'}
          mt-3 w-screen max-w-xs sm:px-0
        `}
      >
        <div className="overflow-hidden rounded-md shadow-xl">
          <div className="relative grid gap-3 bg-bgLight px-2.5 py-3  sm:p-2">
            {children}
          </div>
        </div>
      </Popover.Panel>
    </Transition>
  )
}

function TopBarButtons() {
  const topBarNavLinks = Object.entries(NAVIGATION).map(([key, value]) => (
    <TopBarNavLink
      key={key}
      to={value.path}
      labelText={value.text}
      match={value.match}
      className={key === 'Analytics' ? 'hidden mdl:block' : ''}
    />
  ))

  return <>{topBarNavLinks}</>
}

function MoreInfoButtons() {
  return (
    <>
      <MoreInfoItem
        className="mdl:hidden"
        to={NAVIGATION.Analytics.path}
        labelText={NAVIGATION.Analytics.text}
        description="See preliminary analytics of the bridge"
      />
      <MoreInfoItem
        to={NAVIGATION.Contracts.path}
        labelText={NAVIGATION.Contracts.text}
        description="View contract related information such as contract addresses"
      />
    </>
  )
}

function SocialButtons() {
  return (
    <Grid cols={{ xs: 2, sm: 1 }} gapY={'1'}>
      <MiniInfoItem
        href={DOCS_URL}
        labelText="Docs"
        icon={<DocumentTextIcon className="inline w-5 mr-2 -ml-1 " />}
      />
      <MiniInfoItem
        href={DISCORD_URL}
        labelText="Discord"
        icon={<DiscordIcon className="inline w-5 mr-2 -ml-1" />}
      />
      <MiniInfoItem
        href={TELEGRAM_URL}
        labelText="Telegram"
        icon={<TelegramIcon className="inline w-5 mr-2 -ml-1 " />}
      />
      <MiniInfoItem
        href={TWITTER_URL}
        labelText="Twitter"
        icon={<TwitterIcon className="inline w-5 mr-2 -ml-1 " />}
      />
      <MiniInfoItem
        href={FORUM_URL}
        labelText="Forum"
        icon={<ForumIcon className="inline w-5 mr-2 -ml-1" />}
      />
    </Grid>
  )
}

function MobileBarButtons() {
  const mobileBarItems = Object.entries(NAVIGATION).map(([key, value]) => (
    <MobileBarItem key={key} to={value.path} labelText={value.text} />
  ))

  return <>{mobileBarItems}</>
}

function MobileBarItem({ to, labelText }: { to: string; labelText: string }) {
  const match =
    location.pathname.split('/')[1] === to.split('/')[1] && to !== '#'
  const isInternal = to[0] === '/' || to[0] === '#'

  return (
    <a
      key={labelText}
      href={to}
      target={isInternal ? undefined : '_blank'}
      className={`
        block
        px-3 pt-2 pb-2 rounded-md
        text-2xl font-semibold
        hover:text-opacity-100
      `}
    >
      {isInternal && match ? (
        <p className="text-white">{labelText}</p>
      ) : (
        <p className="text-white text-opacity-30">{labelText}</p>
      )}
    </a>
  )
}

function MoreInfoItem({
  to,
  labelText,
  description,
  className,
}: {
  to: string
  labelText: string
  description: string
  className?: string
}) {
  return (
    <a
      key={labelText}
      href={to}
      target={to[0] === '/' ? undefined : '_blank'}
      className={`block px-3 pt-2 pb-2 rounded-md hover:bg-white hover:bg-opacity-10 ${className}`}
    >
      <p className="text-base font-medium text-white">{labelText}</p>
      <p className="hidden mt-1 text-sm text-white text-opacity-60 md:block">
        {description}
      </p>
    </a>
  )
}

function MiniInfoItem({
  href,
  icon,
  labelText,
}: {
  href: string
  icon: JSX.Element
  labelText: string
}) {
  return (
    <a
      key={labelText}
      href={href}
      className="block px-3 pt-1 pb-2 text-sm rounded-md group"
      target="_blank"
    >
      <div>
        <p className="text-base text-white text-opacity-40 group-hover:text-white">
          {icon}
          <span className="mt-1">{labelText}</span>
        </p>
      </div>
    </a>
  )
}

export function SynapseTitleLogo({ showText }: { showText: boolean }) {
  return (
    <a href={LANDING_PATH} className="flex">
      <div className="flex items-center flex-shrink-0 py-1 ">
        {showText ? <SynapseLogoWithTitleSvg /> : <SynapseLogoSvg />}
      </div>
    </a>
  )
}
