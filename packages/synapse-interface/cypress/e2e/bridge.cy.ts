const BRIDGE_CONSTANTS = 'bridge.json'

describe('Bridge', () => {
  beforeEach(() => cy.visit('/'))

  it('should initially load total # of origin networks', () => {
    cy.get('[data-test-id="bridge-origin-chain-list-button"]')
      .should('be.visible')
      .click()

    cy.get('[data-test-id="bridge-origin-chain-list"]').should('be.visible')

    cy.fixture(BRIDGE_CONSTANTS).then((fixture) => {
      cy.get('button[data-test-id="bridge-origin-chain-list-item"]').should(
        'have.length',
        fixture.totalAvailableNetworks
      )
    })
  })

  it('network dropdown should correct list all origin network names', () => {
    cy.get('[data-test-id="bridge-origin-chain-list-button"]')
      .should('be.visible')
      .click()

    cy.get('[data-test-id="bridge-origin-chain-list"]').should('be.visible')

    cy.fixture(BRIDGE_CONSTANTS).then((fixture) => {
      const networksArray = fixture.availableNetworks
      networksArray.forEach((network) => {
        cy.get('button[data-test-id="bridge-origin-chain-list-item"]').should(
          'contain',
          network
        )
      })
    })
  })

  it('should load possible origin tokens, given a specific chainId', () => {})

  it('should load possible destination tokens, given a specific chainId and origin token', () => {})
})
