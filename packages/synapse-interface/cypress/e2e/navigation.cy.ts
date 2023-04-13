const WINDOW_CONSTANTS = 'window.json'

describe('Navbar ', () => {
  beforeEach(() => cy.visit('/'))

  it('will be visible when screen width is =<1024px', () => {
    cy.fixture(WINDOW_CONSTANTS).then((fixture) => {
      const largeBreakpoint = fixture.breakpoints.large
      cy.viewport(largeBreakpoint.width, largeBreakpoint.height)

      cy.get('nav[data-test-id="desktop-nav"]').should('be.visible')
    })
  })
})
