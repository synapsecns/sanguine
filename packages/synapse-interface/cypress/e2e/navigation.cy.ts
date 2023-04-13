const WINDOW_CONSTANTS = 'window.json'

describe('Navbar ', () => {
  beforeEach(() => cy.visit('/'))

  it('[desktop] will be visible when screen width is >=1024px', () => {
    cy.fixture(WINDOW_CONSTANTS).then((fixture) => {
      const largeBreakpoint = fixture.breakpoints.large
      cy.viewport(largeBreakpoint.width, largeBreakpoint.height)

      cy.get('nav[data-test-id="desktop-nav"]').should('be.visible')
    })
  })

  it('[desktop] will be hidden when screen with is <1024px', () => {
    cy.fixture(WINDOW_CONSTANTS).then((fixture) => {
      const lessThanLargeBreakpoint = fixture.breakpoints.lessThanLarge
      const { width, height } = lessThanLargeBreakpoint
      cy.viewport(width, height)

      cy.get('nav[data-test-id="desktop-nav"]').should('be.hidden')
    })
  })

  it('[mobile] button will be visible when screen with is <1024px', () => {
    cy.fixture(WINDOW_CONSTANTS).then((fixture) => {
      const lessThanLargeBreakpoint = fixture.breakpoints.lessThanLarge
      const { width, height } = lessThanLargeBreakpoint
      cy.viewport(width, height)

      cy.get('button[data-test-id="mobile-navbar-button"]').should('be.visible')
    })
  })

  it('[mobile] button shows routes when clicked', () => {
    cy.fixture(WINDOW_CONSTANTS).then((fixture) => {
      const lessThanLargeBreakpoint = fixture.breakpoints.lessThanLarge
      const { width, height } = lessThanLargeBreakpoint
      cy.viewport(width, height)

      cy.get('button[data-test-id="mobile-navbar-button"]')
        .should('be.visible')
        .click()

      cy.get('div[data-test-id="mobile-nav"]').should('be.visible')
    })
  })

  it('[desktop] shows all routes', () => {})

  it('[mobile] shows all routes', () => {})
})
