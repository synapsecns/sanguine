import { BRIDGABLE_TOKENS, CHAINS } from '@synapsecns/synapse-constants'

const CHAINS_BY_ID = {}

for (const { chainImg, id, name } of Object.values(CHAINS)) {
  if (id && name) {
    CHAINS_BY_ID[id] = { name, chainImg }
  }
}

export default () =>
  Object.entries(BRIDGABLE_TOKENS).map(([id, tokens]) => {
    const chain = CHAINS_BY_ID[id]
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
          const tokenImg =
            typeof token.icon === 'string' ? (
              <img width="16" height="16" src={token.icon} />
            ) : (
              token.icon({ width: 16, height: 16 })
            )

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
              {tokenImg} {token.symbol}
            </span>
          )
        })}
      </section>
    )
  })
