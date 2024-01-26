import { Fragment } from 'react'
import { Popover, Transition } from '@headlessui/react'
import { DocumentTextIcon, MenuIcon, XIcon } from '@heroicons/react/outline'
import React from 'react'

// import { useActiveWeb3React } from '@hooks/wallet/useActiveWeb3React'
import {
  ANALYTICS_PATH,
  BRIDGE_PATH,
  CONTRACTS_PATH,
  DISCORD_URL,
  DOCS_URL,
  FORUM_URL, INTERCHAIN_PATH,
  LANDING_PATH,
  POOLS_PATH,
  PORTFOLIO_PATH,
  STAKE_PATH,
  SWAP_PATH,
  TELEGRAM_URL,
  TWITTER_URL,
} from '@urls'
import Grid from '@components/tailwind/Grid'
import ForumIcon from '@components/icons/ForumIcon'
import TwitterIcon from '@components/icons/TwitterIcon'
import DiscordIcon from '@components/icons/DiscordIcon'
import TelegramIcon from '@components/icons/TelegramIcon'

import { SynapseLogoSvg } from './SynapseLogoSvg'
import TopBarNavLink from './TopBarNavLink'
import WalletNetworkSection from './WalletNetworkSection'
import MoreButton from './MoreButton'
import { PageFooter } from './PageFooter'

export function PageWrapper({ children }) {
  return (
    <div
      className="min-h-screen bg-no-repeat bg-synapse tabular-nums"
      // style={{
      //   background:
      //     'radial-gradient(23.86% 33.62% at 50.97% 47.88%, rgba(255, 0, 255, 0.04) 0%, rgba(172, 143, 255, 0.04) 100%), #111111',
      // }}
    >
      <LandingNav />
      {children}
      <PageFooter />
    </div>
  )
}

export function LandingNav() {
  const topBarBtns = <TopBarButtons />
  const mobileBarBtns = <MobileBarButtons />
  const moreInfoBtns = <MoreInfoButtons />
  const socialBtns = <SocialButtons />

  return (
    <Popover className="relative px-8 pt-6">
      <div className="w-full md:flex-1 md:flex md:items-center md:justify-between">
        <div className="flex justify-between w-full py-4 border-b border-none sm:px-20 ">
          <SynapseTitleLogo showText={true} />
          <div className="items-center justify-center -mr-2 sm:flex lg:hidden">
            <Popover.Button
              className={`
                  rounded-lg p-2 inline-flex items-center justify-center
                  text-gray-400 hover:text-gray-500 hover:bg-gray-100
                  dark:hover:text-gray-400 dark:hover:bg-gray-800
                  focus:outline-none
                `}
            >
              <span className="sr-only">Open menu</span>
              <MenuIcon className="w-8 h-8" aria-hidden="true" />
            </Popover.Button>
          </div>
          <Popover.Group
            as="nav"
            className="hidden lg:flex md:space-x-2 md:justify-space-evenly"
          >
            {topBarBtns}
          </Popover.Group>
          <div className="items-center hidden md:ml-4 lg:flex">
            <div className="space-x-2">
              <WalletNetworkSection />
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
                      {moreInfoBtns}
                      {socialBtns}
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
          <div className="h-full min-h-full  divide-y bg-[#111111] divide-gray-50 dark:divide-gray-600">
            <div className="px-4 pt-1 pb-6">
              <div className="flex items-center justify-between mt-5 ml-3">
                <SynapseTitleLogo showText={false} />
                <div className="-mr-2">
                  <Popover.Button
                    className={`
                        rounded-lg p-2 inline-flex items-center justify-center
                        text-gray-400 hover:text-gray-500 hover:bg-gray-100
                        dark:hover:text-gray-400 dark:hover:bg-gray-900
                        focus:outline-none
                      `}
                  >
                    <span className="sr-only">Close menu</span>
                    <XIcon className="w-6 h-6" aria-hidden="true" />
                  </Popover.Button>
                </div>
              </div>
              <div className="mt-6">
                <Grid cols={{ xs: 1 }} gap={2} className="py-6">
                  {mobileBarBtns}
                </Grid>
              </div>
            </div>
            <Grid cols={{ xs: 1 }} gap={4} className="px-4 py-4">
              <WalletNetworkSection />
            </Grid>
          </div>
        </Popover.Panel>
      </Transition>
    </Popover>
  )
}

function PopoverPanelContainer({ children, className }) {
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
        <div className="overflow-hidden shadow-xl rounded-xl">
          <div className="relative grid gap-3 bg-white dark:bg-[#111111] px-2.5 py-3  sm:p-2">
            {children}
          </div>
        </div>
      </Popover.Panel>
    </Transition>
  )
}

function TopBarButtons() {
  return (
    <>
      <TopBarNavLink to={LANDING_PATH} labelText="About" />
      <TopBarNavLink to={BRIDGE_PATH} labelText="Bridge" />
      <TopBarNavLink to={SWAP_PATH} labelText="Swap" />
      <TopBarNavLink to={POOLS_PATH} labelText="Pools" />
      <TopBarNavLink to={STAKE_PATH} labelText="Stake" />
      <TopBarNavLink to={INTERCHAIN_PATH} labelText={'Interchain Network'} />
      <TopBarNavLink
        className="hidden mdl:block"
        to={ANALYTICS_PATH}
        labelText="Analytics"
      />
    </>
  )
}

function MoreInfoButtons() {
  // const { chainId } = useActiveWeb3React()

  return (
    <>
      <MoreInfoItem
        className="mdl:hidden"
        to={ANALYTICS_PATH}
        labelText="Analytics"
        description="See preliminary analytics of the bridge"
      />
      {/* <MoreInfoItem
        to={getBuySynUrl({ chainId })}
        labelText="Buy $SYN"
        description="Trade and add liquidity to $SYN pools"
      /> */}
      <MoreInfoItem
        to={CONTRACTS_PATH}
        labelText="Contracts"
        description="View contract related information such as contract addresses"
      />
      <MoreInfoItem
        to={PORTFOLIO_PATH}
        labelText="Portfolio"
        description="View your portfolio of related assets on this chain"
      />
    </>
  )
}

function SocialButtons() {
  return (
    <Grid cols={{ xs: 2, sm: 1 }} gapY={1}>
      <MiniInfoItem
        href={DOCS_URL}
        labelText="Docs"
        icon={
          <DocumentTextIcon
            className={`
              w-4 -ml-1 mr-2  inline group-hover:text-blue-700
              dark:text-blue-700
            `}
          />
        }
      />
      <MiniInfoItem
        href={DISCORD_URL}
        labelText="Discord"
        icon={
          <DiscordIcon
            className={`
              w-4 -ml-1 mr-2  inline group-hover:text-indigo-500
              dark:text-indigo-500
            `}
          />
        }
      />
      <MiniInfoItem
        href={TELEGRAM_URL}
        labelText="Telegram"
        icon={
          <TelegramIcon
            className={`
              w-4 -ml-1 mr-2  inline group-hover:text-blue-400
              dark:text-blue-400
            `}
          />
        }
      />
      <MiniInfoItem
        href={TWITTER_URL}
        labelText="Twitter"
        icon={
          <TwitterIcon
            className={`
              w-4 -ml-1 mr-2  inline group-hover:text-sky-400
              dark:text-sky-400
            `}
          />
        }
      />
      <MiniInfoItem
        href={FORUM_URL}
        labelText="Forum"
        icon={
          <ForumIcon
            className={`
              w-4 -ml-1 mr-2  inline group-hover:text-purple-700
              dark:text-purple-700
            `}
          />
        }
      />
    </Grid>
  )
}

function MobileBarButtons() {
  return (
    <>
      <MobileBarItem to={LANDING_PATH} labelText="About" />
      <MobileBarItem to={BRIDGE_PATH} labelText="Bridge" />
      <MobileBarItem to={SWAP_PATH} labelText="Swap" />
      <MobileBarItem to={POOLS_PATH} labelText="Pools" />
      <MobileBarItem to={STAKE_PATH} labelText="Stake" />
      <MobileBarItem to={ANALYTICS_PATH} labelText="Analytics" />
    </>
  )
}

function MobileBarItem({ to, labelText }) {
  return (
    <a
      key={labelText}
      href={to}
      target={to[0] === '/' ? undefined : '_blank'}
      className={`block px-3 pt-2 pb-2 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-900`}
    >
      {to === LANDING_PATH ? (
        <p className="text-2xl font-semibold text-white ">{labelText}</p>
      ) : (
        <p className="text-2xl font-semibold text-white opacity-30 ">
          {labelText}
        </p>
      )}
    </a>
  )
}

interface MoreInfoItemProps {
  to: string,
  labelText: string,
  className?: string
  description?: string
}

const MoreInfoItem: React.FC<MoreInfoItemProps> = ({ to, labelText, description, className }) => {
  return (
    <a
      key={labelText}
      href={to}
      target={to[0] === '/' ? undefined : '_blank'}
      className={`block px-3 pt-2 pb-2 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-900 ${className}`}
    >
      <p className="text-base font-medium text-white ">{labelText}</p>
      <p className="hidden mt-1 text-sm text-gray-500 md:block dark:text-gray-500">
        {description}
      </p>
    </a>
  )
}

function MiniInfoItem({ href, icon, labelText }) {
  return (
    <a
      key={labelText}
      href={href}
      className="block px-3 pt-1 pb-2 text-sm rounded-lg group hover:bg-gray-50 dark:hover:bg-gray-900"
      target="_blank"
    >
      <div>
        <p className="text-sm text-gray-500 dark:text-gray-500 group-hover:text-gray-600 dark:group-hover:text-gray-400">
          {icon}
          <span className="mt-1">{labelText}</span>
        </p>
      </div>
    </a>
  )
}

export function SynapseTitleLogo({ showText }) {
  return (
    <div>
      <a href={'/'} className="flex">
        <div className="flex items-center flex-shrink-0 py-1 ">
          <div className="mr-2">
            <SynapseLogoSvg />
          </div>
          <span
            className={`
              ${showText ? '' : 'hidden'}
              font-medium text-2xl tracking-wide pl-2
              bg-clip-text text-transparent bg-white
              hover:animate-pulse
              transform transition-all
            `}
          >
            Synapse
          </span>
        </div>
      </a>
    </div>
  )
}
