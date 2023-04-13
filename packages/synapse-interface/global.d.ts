declare namespace Cypress {
  interface Chainable {
    visitWithWallet(url: string): Chainable<any>
  }
}
