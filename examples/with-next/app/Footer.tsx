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

  return <footer>

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
        { label: 'Swap', url: '#', },
        { label: 'Bridge', url: '#', },
        { label: 'Pools', url: '#', },
        { label: 'Stake', url: '#', },
      ]}/>
      
      <FooterSection header='Developers' links={[
        { label: 'Build on Synapse', url: '#', },
        { label: 'Documentation', url: '#', },
        { label: 'Github', url: '#', },
        { label: 'Blog', url: '#', },
      ]}/>

      <FooterSection header='Support' links={[
        { label: 'Discord', url: '#', },
        { label: 'Twitter', url: '#', },
        { label: 'Forum', url: '#', },
        { label: 'Telegram', url: '#', },
      ]}/>

    </nav>

    </div>

  </footer>

}