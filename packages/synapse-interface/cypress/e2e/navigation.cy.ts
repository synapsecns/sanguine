export const WINDOW_CONSTANTS = 'window.json'
export const NAVIGATION_CONSTANTS = 'navigation.json'

describe('Navbar', () => {
  beforeEach(() => cy.visit('/'))

  it('[desktop] will be visible when screen width >=1024px', () => {
    cy.fixture(WINDOW_CONSTANTS).then((fixture) => {
      const largeBreakpoint = fixture.screenWidth.large
      cy.viewport(largeBreakpoint.width, largeBreakpoint.height)

      cy.get('nav[data-test-id="desktop-nav"]').should('be.visible')

      const xlargeBreakpoint = fixture.screenWidth.xlarge
      cy.viewport(xlargeBreakpoint.width, xlargeBreakpoint.height)

      cy.get('nav[data-test-id="desktop-nav"]').should('be.visible')
    })
  })

  it('[desktop] will be hidden when screen width <1024px', () => {
    cy.fixture(WINDOW_CONSTANTS).then((fixture) => {
      const medium_screen = fixture.screenWidth.medium
      cy.viewport(medium_screen.width, medium_screen.height)

      cy.get('nav[data-test-id="desktop-nav"]').should('be.hidden')

      const small_screen = fixture.screenWidth.small
      cy.viewport(small_screen.width, small_screen.height)

      cy.get('nav[data-test-id="desktop-nav"]').should('be.hidden')
    })
  })

  it('[desktop] shows all routes, in the correct order', () => {
    cy.fixture(WINDOW_CONSTANTS).then((fixture) => {
      const medium_screen = fixture.screenWidth.medium
      cy.viewport(medium_screen.width, medium_screen.height)
    })

    cy.fixture(NAVIGATION_CONSTANTS).then((fixture) => {
      const routes = fixture.routes
      cy.get('nav[data-test-id="desktop-nav"]')
        .children('a')
        .should('have.length', routes.length)
        .each(($a, index) => {
          expect($a.text()).to.equal(routes[index])
        })
    })
  })

  it('[mobile] button will be visible when screen with is <1024px', () => {
    cy.fixture(WINDOW_CONSTANTS).then((fixture) => {
      const medium_screen = fixture.screenWidth.medium
      cy.viewport(medium_screen.width, medium_screen.height)

      cy.get('button[data-test-id="mobile-navbar-button"]').should('be.visible')

      const small_screen = fixture.screenWidth.small
      cy.viewport(small_screen.width, small_screen.height)

      cy.get('button[data-test-id="mobile-navbar-button"]').should('be.visible')
    })
  })

  it('[mobile] button shows routes when clicked', () => {
    cy.fixture(WINDOW_CONSTANTS).then((fixture) => {
      const smallBreakpoint = fixture.screenWidth.small
      const { width, height } = smallBreakpoint
      cy.viewport(width, height)

      cy.get('button[data-test-id="mobile-navbar-button"]')
        .should('be.visible')
        .click()

      cy.get('div[data-test-id="mobile-nav"]').should('be.visible')
    })
  })

  it('[mobile] shows all routes, in the correct order', () => {
    cy.fixture(WINDOW_CONSTANTS).then((fixture) => {
      const medium_screen = fixture.screenWidth.medium
      cy.viewport(medium_screen.width, medium_screen.height)
    })

    cy.get('button[data-test-id="mobile-navbar-button"]').click()

    cy.fixture(NAVIGATION_CONSTANTS).then((fixture) => {
      const routes = fixture.routes
      cy.get('div[data-test-id="mobile-nav"]')
        .children('a')
        .should('have.length', routes.length)
        .each(($a, index) => {
          expect($a.text()).to.equal(routes[index])
        })
    })
  })
})
