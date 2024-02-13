export const Appearance = () => {
  return (
    <>
            <h4>Dark mode</h4>
      <p>
        To override the default light theme, set <code>bgColor</code> to{' '}
        <code>'dark'</code>.
      </p>
      <pre>customTheme = &#123; bgColor: 'light' &#125;</pre>
      <h4>Auto-palette</h4>
      <p>
        Generate a palette based on your brand colors by setting bgColor to any
        hex, rgb, or hsl color string. Hex values must contain 6 characters.
      </p>
      <pre>
{`customTheme = {
  bgColor: '#000A14'
  bgColor: 'rgb(0 10 20)'
  bgColor: 'hsl(210deg 100% 4%)'
}`}
      </pre>
      <h4>Global Overrides</h4>
      <p>
        The following CSS variables can be added to your CustomTheme to override
        the generated values. Any valid CSS color string can be used, including
        var() aliases.
      </p>
      <pre>
{`const customTheme = {
  '--synapse-text': 'white',
  '--synapse-secondary': '#ffffffb3',
  '--synapse-root': '#16182e',
  '--synapse-surface': 'linear-gradient(90deg, #1e223de6, #262b47e6)',
  '--synapse-border': 'transparent',
}`}
      </pre>
      <h4>Object Overrides</h4>
      <p>
        Select and button elements can be specifically overriddden to introduce
        brand colors or custom styles.
      </p>
      <pre>
{`const customTheme = {
  '--synapse-focus': 'var(--synapse-secondary)',

  '--synapse-select-bg': 'var(--synapse-root)',
  '--synapse-select-text': 'var(--synapse-text)',
  '--synapse-select-border': 'var(--synapse-border)',

  '--synapse-button-bg': 'var(--synapse-surface)',
  '--synapse-button-text': 'var(--synapse-text)',
  '--synapse-button-border': 'var(--synapse-border)',

  '--synapse-progress': 'hsl(265deg 100% 65%)',
  '--synapse-progress-flash': 'hsl(215deg 100% 65%)',
  '--synapse-progress-success': 'hsl(120deg 100% 30%)',
  '--synapse-progress-error': 'hsl(15deg 100% 65%)',
}`}
      </pre>
    </>
  )
}
