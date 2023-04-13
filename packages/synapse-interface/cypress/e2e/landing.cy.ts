import dayjs from 'dayjs'

const BRIDGE_CONSTANTS = 'bridge.json'
const WINDOW_CONSTANTS = 'window.json'

describe('Landing', () => {
  beforeEach(() => cy.visit('/'))

  it('initially loads bridge page', () => {
    cy.get('[data-test-id="bridge-page"]').should('be.visible')

    cy.fixture(WINDOW_CONSTANTS).then((fixture) => {
      const nowTime = dayjs().format('MM-DD-YYYY@hh-mm-a')
      const fileName = `${nowTime}`

      const largeBreakpoint = fixture.screenWidth.large
      cy.viewport(largeBreakpoint.width, largeBreakpoint.height)
      cy.screenshot(`landing/Desktop-${fileName}`)

      const smallBreakpoint = fixture.screenWidth.small
      cy.viewport(smallBreakpoint.width, smallBreakpoint.height)
      cy.screenshot(`landing/Mobile-${fileName}`)
    })
  })

  it('initially loads default origin token', () => {
    cy.fixture(BRIDGE_CONSTANTS).then((fixture) => {
      cy.get('[data-test-id="bridge-origin-token"]')
        .should('be.visible')
        .and('contain.text', fixture.defaultOriginToken)
    })
  })

  it('initially loads default destination token', () => {
    cy.fixture(BRIDGE_CONSTANTS).then((fixture) => {
      cy.get('[data-test-id="bridge-destination-token"]')
        .should('be.visible')
        .and('contain.text', fixture.defaultDestinationToken)
    })
  })
})
