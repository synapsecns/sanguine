import { Fragment } from 'react'
import { useRouter } from 'next/router'
import { Popover, Transition } from '@headlessui/react'
import { MenuIcon, XIcon } from '@heroicons/react/outline'
import { useTranslations } from 'next-intl'

import Grid from '@tw/Grid'
import ForumIcon from '@icons/ForumIcon'
import TwitterIcon from '@icons/TwitterIcon'
import DiscordIcon from '@icons/DiscordIcon'
import TelegramIcon from '@icons/TelegramIcon'
import DocumentTextIcon from '@icons/DocsIcon'
import { Wallet } from '@components/Wallet'

import { SynapseLogoSvg, SynapseLogoWithTitleSvg } from './SynapseLogoSvg'
import { TopBarNavLink, checkIsRouteMatched } from './TopBarNavLink'
import {
  DISCORD_URL,
  SYNAPSE_DOCS_URL,
  FORUM_URL,
  LANDING_PATH,
  TELEGRAM_URL,
  TWITTER_URL,
  STAKE_SYN_FOR_CX_URL,
} from '@/constants/urls'
import { NAVIGATION } from '@/constants/routes'
import { MoreButton } from './MoreButton'
import { PageFooter } from './PageFooter'
import { joinClassNames } from '@/utils/joinClassNames'
import {
  MaintenanceBanners,
  useMaintenance,
} from '@/components/Maintenance/Maintenance'
import { AnnouncementBanner } from '@/components/Maintenance/components/AnnouncementBanner'
import { LanguageSelector } from '@/components/LanguageSelector'

const wrapperClassNames = {
  textColor: 'text-zinc-800 dark:text-zinc-200',
  font: 'tracking-wide',
  bgColor: 'bg-gradient-to-b',
  bgGradient: 'from-white to-[hsl(235deg_75%_96%)]',
  bgGradientDark: 'dark:from-black dark:to-[hsl(265deg_25%_7.5%)]',
  // bgFrame: 'w-screen h-screen overflow-scroll', // TODO: Enable once wrapperStyle is removed
}

const TODO_REMOVE_wrapperStyle = {
  background:
    'radial-gradient(23.86% 33.62% at 50.97% 47.88%, rgba(255, 0, 255, 0.04) 0%, rgba(172, 143, 255, 0.04) 100%), #111111',
  backgroundImage: `url('landingBg.svg')`,
  backgroundSize: '800px',
  backgroundPosition: 'center 150px',
  backgroundRepeat: 'no-repeat',
}


const HypercallWordmark = ({ height = 16 }: { height?: number }) => (
  <svg
    viewBox="0 0 180 36"
    height={height}
    width={(180 / 36) * height}
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
    style={{ display: 'block' }}
    aria-label="Hypercall"
  >
    <g fill="currentColor">
      <path d="M169.741 28.5L174.221 0.179993H179.261L174.781 28.5H169.741Z" />
      <path d="M174 24.5H177L176.5 28.5H173.5L174 24.5Z" />
      <path d="M172.5 0.174988H175.5L175 4.17499H172L172.5 0.174988Z" />
      <path d="M158.819 28.5L163.299 0.179993H168.339L163.859 28.5H158.819Z" />
      <path d="M163 24.5H166L165.5 28.5H162.5L163 24.5Z" />
      <path d="M161.5 0.174988H164.5L164 4.17499H161L161.5 0.174988Z" />
      <path d="M133.25 19.46C133.25 16 134.75 12.48 137.25 10.25C139.75 8.02002 142.661 8.02002 145.341 8.02002C150.421 8.02002 153.261 11.46 153.261 16.58C153.261 23.66 148.581 28.82 142.021 28.82C139.461 28.82 137.25 27.25 136 26.25C134.75 25.25 133.25 22.92 133.25 19.46ZM143.601 24.46C147.121 24.46 149.681 21.54 149.681 17.38C149.681 14.38 148.001 12.38 145.201 12.38C141.521 12.38 138.801 15.3 138.801 19.46C138.801 22.46 140.641 24.46 143.601 24.46Z" />
      <path d="M156.341 8.33997H151.301L148.097 28.5H153.137L156.341 8.33997Z" />
      <path d="M152.5 24.5H155.5L155 28.5H152L152.5 24.5Z" />
      <path d="M120.281 28.82C114.881 28.82 110.961 24.9 110.961 19.62C110.961 12.86 115.801 8.02002 122.641 8.02002C127.681 8.02002 131.321 11.06 131.401 15.3L126.241 16.18C126.201 14.02 124.481 12.38 122.121 12.38C118.561 12.38 116.281 15.1 116.281 19.18C116.281 22.22 118.081 24.46 120.761 24.46C123.161 24.46 124.921 23.18 125.441 21.06L130.361 22.06C129.361 26.18 125.521 28.82 120.281 28.82Z" />
      <path d="M101.215 21.1001L100.055 28.5001H95.0146L98.2147 8.25H103.255H111L110.25 12.1401H102.655L101.215 21.1001Z" />
      <path d="M93.0411 16.42C93.0411 17.46 92.8411 18.7 92.5611 19.66H77.6011C77.7211 22.78 79.4811 24.58 82.4411 24.58C84.7211 24.58 86.4011 23.5 87.0011 21.5L91.7611 22.78C90.6011 26.38 86.6411 28.82 81.8011 28.82C76.4011 28.82 72.6411 25.06 72.6411 19.62C72.6411 13.98 76.4411 8.02002 83.9611 8.02002C90.0011 8.02002 93.0411 12.22 93.0411 16.42ZM87.8411 16.06C88.0011 13.66 86.2811 11.86 83.3611 11.86C80.7211 11.86 78.8811 13.54 78.0811 16.06H87.8411Z" />
      <path d="M46.8625 35.78L49.4557 19.46C49.4557 19.46 50 13.5 53 10.75C56 8.00003 59.5 7.75001 62.25 8.00001C66 8.25001 70.1826 11.46 70.1826 16.58C70.1826 23.66 65.5026 28.82 58.9426 28.82C56.3826 28.82 54.3426 27.9 53.3826 26.26L51.9026 35.78H46.8625ZM59.0226 24.46C62.5426 24.46 65.1026 21.54 65.1026 17.38C65.1026 14.38 63.4226 12.38 60.6226 12.38C56.9426 12.38 54.2226 15.3 54.2226 19.46C54.2226 22.46 56.0625 24.46 59.0226 24.46Z" />
      <path d="M45.3174 8.34003H50.3174L35.3574 35.78H30.5174L35.0374 27.42L30.3574 8.34003H35.5574L38.3174 21.3L45.3174 8.34003Z" />
      <path d="M34.5 31.78H40.5L40 35.78H34L34.5 31.78Z" />
      <path d="M29.5 8.34998H35.5L35 12.35H29L29.5 8.34998Z" />
      <path d="M23.2802 0.5H28.7202L24.2802 28.5H18.8402L20.7602 16.3H8.52016L6.60016 28.5H1.16016L5.60016 0.5H11.0402L9.28016 11.62H21.5202L23.2802 0.5Z" />
      <path d="M3 0.5H6L5.5 4.5H2.5L3 0.5Z" />
      <path d="M24.5 24.5H27.5L27 28.5H24L24.5 24.5Z" />
    </g>
  </svg>
)

const HYPERCALL_APP_URL = 'https://app.hypercall.xyz/'
const HC_LIME = '#A9FA38'

const HypercallBannerContent = () => {
  return (
    <div className="flex flex-wrap items-center justify-center gap-x-3 gap-y-1 px-2">
      <span
        className="inline-block w-2 h-2 rounded-full animate-pulse shrink-0"
        style={{ background: HC_LIME, boxShadow: `0 0 8px ${HC_LIME}` }}
        aria-hidden="true"
      />
      <span className="inline-flex items-baseline gap-2">
        <span className="inline-flex text-white translate-y-[3px]">
          <HypercallWordmark height={16} />
        </span>
        <span className="text-xs text-white/50 whitespace-nowrap">
          by Synapse
        </span>
      </span>
      <span className="font-medium text-white/90">
        is live — options on anything, starting with{' '}
        <span style={{ color: HC_LIME }}>SpaceX.</span>
      </span>
      <span className="text-white/60">
        Defined risk. No liquidations. On Hyperliquid.
      </span>
      <a
        href={HYPERCALL_APP_URL}
        target="_blank"
        rel="noopener noreferrer"
        className="inline-flex items-center gap-1.5 px-4 py-1.5 rounded-[10px] font-semibold whitespace-nowrap hover:opacity-90 transition-opacity"
        style={{ background: HC_LIME, color: '#0b0f0a' }}
      >
        Trade now
        <span aria-hidden="true">→</span>
      </a>
    </div>
  )
}

export function LandingPageWrapper({ children }: { children: any }) {
  return (
    <div className="dark">
      <div
        className={joinClassNames(wrapperClassNames)}
        style={TODO_REMOVE_wrapperStyle}
      >
        <AnnouncementBanner
          bannerId="07102026-hypercall-live"
          bannerContent={<HypercallBannerContent />}
          startDate={new Date('2026-07-10T00:00:00+00:00')}
          endDate={null}
        />
        <MaintenanceBanners />
        <LandingNav />
        {children}
        <PageFooter />
      </div>
    </div>
  )
}

export function LandingNav() {
  const t = useTranslations('Nav')

  return (
    <Popover>
      <div className="flex gap-4 place-content-between p-8 max-w-[1440px] m-auto">
        <SynapseTitleLogo showText={true} />
        <div className="lg:hidden">
          <Popover.Button
            data-test-id="mobile-navbar-button"
            className="p-2 text-gray-400 rounded-md hover:bg-gray-800 focus:outline-none"
          >
            <span className="sr-only">{t('Open menu')}</span>
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
            <LanguageSelector />
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
        as={Fragment as any}
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
                <span className="sr-only">{t('Close menu')}</span>
                <XIcon className="w-8 h-8" aria-hidden="true" />
              </Popover.Button>
            </div>
            <div className="flex flex-col gap-2 py-4" data-test-id="mobile-nav">
              <MobileBarButtons />
            </div>
            <div className="flex items-center px-2 py-4 space-x-2 bg-white/10">
              <LanguageSelector />
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
      as={Fragment as any}
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
  const t = useTranslations('Nav')

  const topBarNavLinks = Object.entries(NAVIGATION).map(([key, value]) => (
    <TopBarNavLink
      key={key}
      to={value.path}
      labelText={t(value.text)}
      match={value.match}
    />
  ))

  return <>{topBarNavLinks}</>
}

function MoreInfoButtons() {
  return (
    <>
      <MoreInfoItem
        to={NAVIGATION.SYN.path}
        labelText={NAVIGATION.SYN.text}
        description="View $SYN related information such as contract addresses"
      />
    </>
  )
}

function SocialButtons() {
  return (
    <Grid cols={{ xs: 2, sm: 1 }} gapY={'1'}>
      <MiniInfoItem
        href={SYNAPSE_DOCS_URL}
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
    <MobileBarItem
      key={key}
      to={value.path}
      labelText={value.text}
      match={value.match}
    />
  ))

  return <>{mobileBarItems}</>
}

function MobileBarItem({
  to,
  labelText,
  match,
}: {
  to: string
  labelText: string
  match?: string | { startsWith: string }
}) {
  const router = useRouter()

  const isRouteMatched = checkIsRouteMatched(router, match)

  const isInternal = to[0] === '/' || to[0] === '#'

  return (
    <a
      key={labelText}
      href={to}
      target={isInternal ? undefined : '_blank'}
      className={`
        px-4 py-2 text-2xl font-medium text-white
        ${isRouteMatched ? 'text-opacity-100' : 'text-opacity-30'}
      `}
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
