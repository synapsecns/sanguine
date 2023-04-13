describe('Bridge', () => {
  beforeEach(() => cy.visit('/'))

  it('should initially load all available origin networks', () => {
    cy.get('[data-test-id="bridge-origin-chain-list-button"]')
      .should('be.visible')
      .click()

    cy.get('')
  })

  it('should load possible origin tokens, given a specific chainId', () => {})

  it('should load possible destination tokens, given a specific chainId and origin token', () => {})
})
