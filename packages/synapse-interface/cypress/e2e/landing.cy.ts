import dayjs from 'dayjs'

export const BRIDGE_CONSTANTS = 'bridge.json'
export const WINDOW_CONSTANTS = 'window.json'

describe('Landing (/)', () => {
  beforeEach(() => cy.visit('/'))

  it('initially loads bridge page', () => {
    cy.get('[data-test-id="bridge-page"]').should('be.visible')

    cy.fixture(WINDOW_CONSTANTS).then((fixture) => {
      const nowTime = dayjs().format('MM-DD-YYYY@hh-mm-a')
      const fileName = `${nowTime}`

      const mobileBreakpoint = fixture.screenWidth.mobile
      cy.viewport(mobileBreakpoint.width, mobileBreakpoint.height)
      cy.screenshot(`landing/mobile_breakpoint-${fileName}`)

      const smallBreakpoint = fixture.screenWidth.small
      cy.viewport(smallBreakpoint.width, smallBreakpoint.height)
      cy.screenshot(`landing/small_breakpoint-${fileName}`)

      const mediumBreakpoint = fixture.screenWidth.medium
      cy.viewport(mediumBreakpoint.width, mediumBreakpoint.height)
      cy.screenshot(`landing/medium_breakpoint-${fileName}`)

      const largeBreakpoint = fixture.screenWidth.large
      cy.viewport(largeBreakpoint.width, largeBreakpoint.height)
      cy.screenshot(`landing/large_breakpoint-${fileName}`)

      const xlargeBreakpoint = fixture.screenWidth.xlarge
      cy.viewport(xlargeBreakpoint.width, xlargeBreakpoint.height)
      cy.screenshot(`landing/xlarge_breakpoint-${fileName}`)
    })
  })

  // it('initially loads default origin token', () => {
  //   cy.fixture(BRIDGE_CONSTANTS).then((fixture) => {
  //     cy.get('[data-test-id="bridge-origin-token"]')
  //       .should('be.visible')
  //       .and('contain.text', fixture.defaultOriginToken)
  //   })
  // })

  // it('initially loads default destination token', () => {
  //   cy.fixture(BRIDGE_CONSTANTS).then((fixture) => {
  //     cy.get('[data-test-id="bridge-destination-token"]')
  //       .should('be.visible')
  //       .and('contain.text', fixture.defaultDestinationToken)
  //   })
  // })
})

describe('Landing Page (/landing)', () => {
  beforeEach(() => cy.visit('/landing'))

  it('initially loads landing page', () => {
    cy.get('[data-test-id="landing-page"]').should('be.visible')
  })

  it('displays all landing page sections', () => {
    cy.fixture(WINDOW_CONSTANTS).then((fixture) => {
      const desktop = fixture.screenWidth.large
      cy.viewport(desktop.width, desktop.height)

      cy.get('[data-test-id="landing-hero-section"]').should('be.visible')
      cy.get('[data-test-id="landing-security-section"]').should('be.visible')
      cy.get('[data-test-id="landing-bridge-section"]').should('be.visible')
      cy.get('[data-test-id="landing-explorer-section"]').should('be.visible')
      cy.get('[data-test-id="landing-integration-section"]').should(
        'be.visible'
      )
      cy.get('[data-test-id="landing-how-it-works-section"]').should(
        'be.visible'
      )
      cy.get('[data-test-id="landing-use-cases-section"]').should('be.visible')
      cy.get('[data-test-id="landing-resources-section"]').should('be.visible')
    })
  })
})
