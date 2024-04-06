import { Fragment } from 'react'

const sections = [
  {
    label: 'About',
    url: '#',
    links: [
      {
        label: 'Vision',
        url: '#',
        description:
          'Vision lorem ipsum dolor sit amet consecteteur adipisicing elit.',
      },
      {
        label: 'Philosophy',
        url: '#',
        description:
          'Philosophy lorem ipsum dolor sit amet consecteteur adipisicing elit.',
      },
      {
        label: 'Roadmap',
        url: '#',
        description:
          'Roadmap lorem ipsum dolor sit amet consecteteur adipisicing elit.',
      },
    ],
  },
  {
    label: 'Bridge',
    url: '#',
    links: [
      {
        label: 'Synapse Bridge',
        url: '#',
        description:
          'Smart routes & real-time competitive quotes on 20 supported chains.',
      },
      {
        label: 'On-chain swap',
        url: '#',
        description:
          'Swap lorem ipsum dolor sit amet consecteteur adipisicing elit.',
      },
      {
        label: 'Solana bridge',
        url: '#',
        description:
          'Solana lorem ipsum dolor sit amet consecteteur adipisicing elit.',
      },
    ],
  },
  {
    label: 'Co',
    after: 'nnect',
    afterXs: 'mmunity',
    url: '#',
    links: [
      {
        label: 'Discord',
        url: '#',
        description:
          'Discord lorem ipsum dolor sit amet consecteteur adipisicing elit.',
      },
      {
        label: 'Telegram',
        url: '#',
        description:
          'Telegram lorem ipsum dolor sit amet consecteteur adipisicing elit.',
      },
      {
        label: 'Twitter',
        url: '#',
        description:
          'Twitter lorem ipsum dolor sit amet consecteteur adipisicing elit.',
      },
      {
        label: 'Blog',
        url: '#',
        description:
          'Blog lorem ipsum dolor sit amet consecteteur adipisicing elit.',
      },
      {
        label: 'Forum',
        url: '#',
        description:
          'Forum lorem ipsum dolor sit amet consecteteur adipisicing elit.',
      },
    ],
  },
  {
    label: 'Dev',
    after: 's',
    afterXs: 'elopers',
    url: '#',
    links: [
      {
        label: 'Docs',
        url: '#',
        description:
          'Docs lorem ipsum dolor sit amet consecteteur adipisicing elit.',
      },
      {
        label: 'GitHub',
        url: '#',
        description:
          'GitHub lorem ipsum dolor sit amet consecteteur adipisicing elit.',
      },
      {
        label: 'Synapse CNS',
        url: '#',
        description:
          'Synapse CNS lorem ipsum dolor sit amet consecteteur adipisicing elit.',
      },
      {
        label: 'Interchain Network',
        url: '#',
        description:
          'Interchain Network lorem ipsum dolor sit amet consecteteur adipisicing elit.',
      },
    ],
  },
  {
    label: 'Explorer',
    url: '#',
  },
]

export default function Header() {
  return (
    <nav className="border-b border-zinc-300 dark:border-zinc-800">
      <ul className="flex flex-wrap text-sm xs:text-base justify-center">
        {sections.map((section) => (
          <li
            key={section.label}
            className="group relative sm:first:inline-block text-zinc-500 hover:text-white"
          >
            <a
              href={section.url}
              className={`px-3 xs:px-4 lg:px-5 py-2.5 hover:bg-zinc-50 hover:dark:bg-zinc-950 border-l border-zinc-300 dark:border-transparent hover:border-fuchsia-500 hover:dark:border-fuchsia-500 inline-block after:content-['${section.after}'] xs:after:content-['${section.afterXs}']`}
            >
              {section.label}
            </a>
            {section.links && (
              <div
                className="hidden group-hover:block absolute p-2 animate-slide-down origin-top w-max z-10"
                style={{ lineHeight: '100%' }}
              >
                <dl className="bg-zinc-50 dark:bg-zinc-950 rounded text-base -ml-2 border border-zinc-200 dark:border-zinc-800 shadow-sm grid grid-cols-[auto_auto]">
                  {section.links.map((link) => {
                    return (
                      <Fragment key={link.label}>
                        <dt className="col-start-1">
                          <a
                            href={link.url}
                            className={`px-4 py-3 block border border-transparent hover:border-fuchsia-500 rounded`}
                          >
                            {link.label}
                          </a>
                        </dt>
                        <dd className="w-60 col-start-2 row-start-1 row-span-6 px-4 py-3 border-l border-zinc-200 dark:border-zinc-800 cursor-pointer hidden [:hover_+_&]:block hover:block">
                          <header>{link.label}</header>
                          <p className="mt-1 font-light tracking-wider">
                            {link.description}
                          </p>
                        </dd>
                      </Fragment>
                    )
                  })}
                </dl>
              </div>
            )}
          </li>
        ))}
      </ul>
    </nav>
  )
}
