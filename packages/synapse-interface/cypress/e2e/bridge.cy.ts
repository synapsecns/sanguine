import { mock } from '@depay/web3-mock'

export const BRIDGE_CONSTANTS = 'bridge.json'

describe('Bridge without wallet connected', () => {
  beforeEach(() => cy.visit('/'))

  it('should initially load total # of origin networks', () => {
    cy.get('[data-test-id="bridge-origin-chain-list-button"]')
      .should('be.visible')
      .click()

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

    cy.get('[data-test-id="bridge-origin-chain-list"]')
      .its('0.offsetHeight')
      .should('be.gt', 0)

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
})

describe('Bridge with metamask wallet connected to ethereum network', () => {
  const network = 'ethereum'
  const wallet = 'metamask'
  const account = ['0xF080B794AbF6BB905F2330d25DF545914e6027F8']

  beforeEach(() => cy.visit('/'))
  beforeEach(() =>
    mock({
      blockchain: network,
      wallet,
      accounts: {
        return: account,
      },
    })
  )

  it('should be connected to metamask', () => {
    cy.window().then(() => {
      expect(global.ethereum).to.exist
      expect(global.ethereum.isMetaMask).to.be.true
      expect(global.ethereum.isCoinbaseWallet).to.be.undefined
      expect(global.ethereum.isBraveWallet).to.be.undefined
    })
  })

  it('should connect to ethereum network', () => {
    cy.window().then(async () => {
      expect(global.ethereum).to.exist
      const currentChainId = await global.ethereum.request({
        method: 'eth_chainId',
      })
      expect(currentChainId).to.equal('0x1')
    })
  })

  // @dev TO-DO: write tests once logic has been implemented in frontend
  // require origin tokens to load when wallet connected initially
  // it('should load possible origin tokens, given a specific chainId', () => {})
  // it('should load possible destination tokens, given a specific chainId and origin token', () => {})
})
