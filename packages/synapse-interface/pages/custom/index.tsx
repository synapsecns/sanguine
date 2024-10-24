import deepmerge from 'deepmerge'
import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'
import { CustomBridge } from '../../components/Custom/CustomBridge'

export async function getStaticProps({ locale }) {
  const userMessages = (await import(`../../messages/${locale}.json`)).default
  const defaultMessages = (await import(`../../messages/en-US.json`)).default
  const messages = deepmerge(defaultMessages, userMessages)

  return {
    props: {
      messages,
    },
  }
}

const CustomPage = () => {
  return (
    <LandingPageWrapper>
      <main
        data-test-id="bridge-page"
        className="relative z-0 flex-1 h-full overflow-y-none focus:outline-none"
      >
        <div className="flex flex-col-reverse justify-center gap-16 px-4 py-20 mx-auto lg:flex-row 2xl:w-3/4 sm:mt-6 sm:px-8 md:px-12">
          <HomeContent />
          <CustomBridge />
        </div>
      </main>
    </LandingPageWrapper>
  )
}

const HomeContent = () => {
  return (
    <div id="custom-home-content" className="w-[500px] text-white space-y-3">
      <h1>Header message</h1>
      <p>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
        tempor incididunt ut labore et dolore magna aliqua.
      </p>
      <p>
        Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi
        ut aliquip ex ea commodo consequat.
      </p>
      <p>
        Duis aute irure dolor in reprehenderit in voluptate velit esse cillum
        dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
        proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
      </p>
    </div>
  )
}

export default CustomPage
