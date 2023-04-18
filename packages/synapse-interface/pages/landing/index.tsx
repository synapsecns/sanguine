import Grid from '@/components/ui/tailwind/Grid'
import Card from '@/components/ui/tailwind/Card'
import Button from '@/components/ui/tailwind/Button'

import { LandingPageWrapper } from '@/components/layouts/LandingPageWrapper'

function LandingPageContainer({ children }) {
  return (
    <div
      data-test-id="landing-page-container"
      className="relative px-4 md:px-24"
    >
      {children}
    </div>
  )
}

const LandingPage = () => {
  return (
    <LandingPageWrapper>
      <LandingPageContainer>hello</LandingPageContainer>
    </LandingPageWrapper>
  )
}

export default LandingPage
