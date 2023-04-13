import dayjs from 'dayjs'

const BRIDGE_CONSTANTS = 'bridge.json'

describe('Landing', () => {
  beforeEach(() => cy.visit('/'))

  it('initially loads bridge page', () => {
    cy.get('[data-test-id="bridge-page"]').should('be.visible')

    const nowTime = dayjs().format('MM-DD-YYYY@hh-mm-a')
    const fileName = `${nowTime}`

    cy.screenshot(`landing/${fileName}`)
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
