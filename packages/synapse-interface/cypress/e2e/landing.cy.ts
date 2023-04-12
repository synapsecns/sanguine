import dayjs from 'dayjs'

describe('Landing Page', () => {
  beforeEach(() => cy.visit('http://localhost:3000/')) // hardcode local for now, update later to include env

  it('loads bridge page', () => {
    cy.get('[data-test-id="bridge-page"]').should('be.visible')

    const nowTime = dayjs().format('MM-DD-YYYY@hh-mm-a')
    const fileName = `${nowTime}`

    cy.screenshot(`landing/${fileName}`)
  })
})
