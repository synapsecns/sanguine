type Props = {
  header: string
  links: { label: string, url: string }[]
}

function FooterSection({ header, links }: Props) {

  return <section>

    <header>{header}</header>

    {links.map(link => 

      <a key={link.label} href={link.url}>
        <span>{link.label}</span>
      </a>

    )}

  </section>

}

export default function Footer() {

  return (
    <footer>
      <div>
        <a href="https://synapseprotocol.com/landing">
          <picture>
            <source srcSet="synapse-logo-onLight.svg" media="(prefers-color-scheme: light)" />
            <source srcSet="synapse-logo-onDark.svg" media="(prefers-color-scheme: dark)" />
            <img src="synapse-logo-onDark.svg" width="160" alt="Synapse Logo" />
          </picture>
        </a>
        <nav>
          <FooterSection header='Functions' links={[
            { label: 'Swap', url: 'https://synapseprotocol.com/swap', },
            { label: 'Bridge', url: 'https://synapseprotocol.com/', },
            { label: 'Pools', url: 'https://synapseprotocol.com/pools', },
            { label: 'Stake', url: 'https://synapseprotocol.com/stake', },
          ]}/>
          <FooterSection header='Developers' links={[
            { label: 'Build on Synapse', url: 'https://docs.synapseprotocol.com/synapse-interchain-network-sin/build-on-the-synapse-interchain-network', },
            { label: 'Documentation', url: 'https://docs.synapseprotocol.com/', },
            { label: 'Github', url: 'https://github.com/synapsecns', },
            { label: 'Blog', url: 'https://synapse.mirror.xyz/', },
          ]}/>
          <FooterSection header='Support' links={[
            { label: 'Discord', url: 'https://discord.com/invite/synapseprotocol', },
            { label: 'Twitter', url: 'https://twitter.com/SynapseProtocol', },
            { label: 'Forum', url: 'https://forum.synapseprotocol.com/', },
            { label: 'Telegram', url: 'https://t.me/synapseprotocol', },
          ]}/>
        </nav>
      </div>
    </footer>
  )
}