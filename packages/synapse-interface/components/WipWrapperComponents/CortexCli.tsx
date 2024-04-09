import { useEffect } from 'react'

let prefersColorScheme, windowPrefersDark

export default function ({ setPrefersDark }) {
  useEffect(() => {
    prefersColorScheme = localStorage.getItem('prefers-color-scheme')
    windowPrefersDark = window.matchMedia('(prefers-color-scheme: dark)')
    windowPrefersDark.addEventListener('change', handleWindowPrefersDark)
    return () =>
      windowPrefersDark.removeEventListener('change', handleWindowPrefersDark)
  }, [])

  function selectPrefersDark(e) {
    switch (e.target.value) {
      case 'Dark mode':
        localStorage.setItem('prefers-color-scheme', 'dark')
        setPrefersDark(true)
        break
      case 'Light mode':
        localStorage.setItem('prefers-color-scheme', 'light')
        setPrefersDark(false)
        break
      default:
        localStorage.removeItem('prefers-color-scheme')
        setPrefersDark(windowPrefersDark.matches)
    }
  }

  function handleWindowPrefersDark(e) {
    !prefersColorScheme && setPrefersDark(e.matches)
  }

  return (
    <div className="bg-zinc-50 dark:bg-zinc-950 sticky bottom-0 mt-24">
      <article className="max-w-7xl mx-auto bg-zinc-50 dark:bg-zinc-950 text-sm flex">
        <input
          placeholder=">_"
          className="text-sm grow bg-transparent text-inherit h-8 border-none"
        ></input>
        <select
          className="h-8 py-0 text-sm bg-white dark:bg-black text-sm text-inherit cursor-pointer rounded border-zinc-200 dark:border-zinc-800 justify-self-end w-min hover:border-zinc-300 hover:dark:bg-zinc-950 hover:dark:border-zinc-700 col-end-4"
          onChange={selectPrefersDark}
        >
          <option
            defaultValue={prefersColorScheme === 'dark' ? 'true' : 'false'}
          >
            Dark mode
          </option>
          <option
            defaultValue={prefersColorScheme === 'light' ? 'true' : 'false'}
          >
            Light mode
          </option>
          <option defaultValue={!prefersColorScheme ? 'true' : 'false'}>
            System {windowPrefersDark?.matches ? 'dark' : 'light'}
          </option>
        </select>
      </article>
    </div>
  )
}
