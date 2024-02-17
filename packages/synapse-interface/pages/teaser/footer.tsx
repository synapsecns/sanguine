import { SynapseAnchor } from "./SynapseLogo"

const sections = [
  {
    header: 'Functions',
    links: [
      { label: 'Swap', url: '#'},
      { label: 'Bridge', url: '#'},
      { label: 'Pools', url: '#'},
      { label: 'Stake', url: '#'},
    ]
  },
  {
    header: 'Developers',
    links: [
      { label: 'Build on Synapse', url: '#'},
      { label: 'Documentation', url: '#'},
      { label: 'GitHub', url: '#'},
      { label: 'Blog', url: '#'},
    ]
  },
  {
    header: 'Support',
    links: [
      { label: 'Discord', url: '#'},
      { label: 'Twitter', url: '#'},
      { label: 'Forum', url: '#'},
      { label: 'Telegram', url: '#'},
    ]
  }]

export default function Footer() { return (
  <footer className="p-4 sm:p-8 flex flex-wrap gap-8 items-start max-w-7xl mx-auto justify-end sm:justify-between cursor-default">
    <SynapseAnchor />
    <div className="flex flex-wrap gap-4 xs:gap-8 text-right justify-end">
      {sections.map(section => (
        <section>
          <header className="px-2 py-1">{section.header}</header>
          <ul>
            {section.links.map(link => <li>
              <a href={link.url} className="text-zinc-500 hover:text-inherit hover:bg-zinc-200 hover:dark:bg-zinc-900 px-2 py-1 rounded inline-block">
                {link.label}
              </a>
            </li>)}
          </ul>
        </section>
      ))}
    </div>
  </footer>
)}
