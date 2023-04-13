describe('Bridge', () => {
  beforeEach(() => cy.visit('/'))

  it('should initially load all available origin networks', () => {
    cy.get('[data-test-id="bridge-origin-chain-list-button"]')
      .should('be.visible')
      .click()

    cy.get('[data-test-id="bridge-origin-chain-list"]').should('be.visible')

    cy.get('button[data-test-id="bridge-origin-chain-list-item"]').should(
      'have.length',
      18
    )
  })

  it('should load possible origin tokens, given a specific chainId', () => {})

  it('should load possible destination tokens, given a specific chainId and origin token', () => {})
})
