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
      style={{
        background:
          'radial-gradient(23.86% 33.62% at 50.97% 47.88%, rgba(255, 0, 255, 0.04) 0%, rgba(172, 143, 255, 0.04) 100%), #111111',
      }}
    >
      <LandingNav />

      <div
        style={{
          backgroundImage: `url('landingBg.svg')`,
          backgroundSize: '800px',
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
    <Popover>
      <div className="flex gap-4 place-content-between p-8 max-w-[1440px] m-auto">
        <SynapseTitleLogo showText={true} />
        <div className="lg:hidden">
          <Popover.Button
            data-test-id="mobile-navbar-button"
            className="p-2 text-gray-400 rounded-md hover:bg-gray-800 focus:outline-none"
          >
            <span className="sr-only">Open menu</span>
            <MenuIcon className="w-8 h-8" aria-hidden="true" />
          </Popover.Button>
        </div>
        <Popover.Group
          as="nav"
          className="flex-wrap justify-center hidden lg:flex"
          data-test-id="desktop-nav"
        >
          <TopBarButtons />
        </Popover.Group>
        <div className="hidden lg:flex h-fit">
          <div className="flex items-center space-x-2">
            <Wallet />
            <Popover className="relative">
              {({ open }) => (
                <>
                  <Popover.Button as="div" onMouseEnter={() => {}}>
                    <MoreButton open={open} />
                  </Popover.Button>
                  <PopoverPanelContainer className="-translate-x-full left-full">
                    <MoreInfoButtons />
                    <SocialButtons />
                  </PopoverPanelContainer>
                </>
              )}
            </Popover>
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
        <Popover.Panel focus className="absolute top-0 z-10 w-screen">
          <div
            className="bg-bgLight"
            // data-test-id="mobile-nav"
          >
            <div className="flex items-center px-4 pt-4 place-content-between">
              <SynapseTitleLogo showText={true} />
              <Popover.Button className="p-2 text-gray-400 rounded-md hover:bg-gray-900 focus:outline-none">
                <span className="sr-only">Close menu</span>
                <XIcon className="w-8 h-8" aria-hidden="true" />
              </Popover.Button>
            </div>
            <div className="flex flex-col gap-2 py-4" data-test-id="mobile-nav">
              <MobileBarButtons />
            </div>
            <div className="px-2 py-4 bg-white/10">
              <Wallet />
            </div>
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
        px-4 py-2 text-2xl font-medium text-white
        ${!(isInternal && match) && 'opacity-30 hover:opacity-100'}`}
    >
      {labelText}
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
    <a href={LANDING_PATH}>
      {showText ? <SynapseLogoWithTitleSvg /> : <SynapseLogoSvg />}
    </a>
  )
}
