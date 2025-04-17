import { BRIDGABLE_TOKENS, CHAINS } from '@synapsecns/synapse-constants'

export default () =>
  Object.entries(BRIDGABLE_TOKENS)
    .filter(([id]) => {
      const chainId = Number(id)
      return ![288, 1666600000, 1284, 1285, 1313161554, 25, 2000].includes(chainId)
    })
    .map(([id, tokens]) => {
      const chain = CHAINS.CHAINS_BY_ID[id]
      const chainImg = chain.chainImg

      return (
        <section key={id}>
          <h2
            style={{
              fontSize: '1.5rem',
              display: 'flex',
              gap: '.75rem',
              alignItems: 'center',
            }}
          >
            <img width="28" height="28" src={chainImg} alt={chain.name} />
            {chain.name} <code>{id}</code>
          </h2>
          {Object.values(tokens).map((token) => {
            return (
              <span
                key={token.addresses[id]}
                style={{
                  display: 'inline-flex',
                  gap: '.5rem',
                  alignItems: 'center',
                  padding: '.25rem .5rem',
                }}
              >
                <img width="16" height="16" src={token.icon} alt={token.symbol} />{' '}
                {token.symbol}
              </span>
            )
          })}
        </section>
      )
    })