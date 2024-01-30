
type Props = {
  links: { label: string; url?: string; selected?: boolean }[]
}

function HeaderSection({ links }: Props) {

  return (
    <nav>
      {links.map((link, i) => {
        return (
          <a
            key={i}
            href={link.url}
            target="_blank" rel="noreferrer"
            className={`cursor-pointer ${link.selected ? 'selected' : ''}`}
          >
            {link.label}
          </a>
        )
      })}

      {/* <NavOverflow links={links} /> */}
      {/* <a onClick={handleClick}>More…</a> */}
    </nav>
  )
}

export default function Header() {
  return (
    <header>
      <a href="https://synapseprotocol.com/landing">
        <picture>
          <source srcSet="synapse-logo-onLight.svg" media="(prefers-color-scheme: light)" />
          <source srcSet="synapse-logo-onDark.svg" media="(prefers-color-scheme: dark)" />
          <img src="synapse-logo-onDark.svg" width="180" alt="Synapse Logo" />
        </picture>
      </a>

      <HeaderSection
        links={[
          { label: 'Widget', selected: true },
          { label: 'Docs', url: 'https://docs.synapseprotocol.com/' },
          { label: 'EVM Bridge', url: 'https://synapseprotocol.com' },
          // { label: 'Synapse Interchain Network', url: '#', },
        ]}
      />
    </header>
  )
}
