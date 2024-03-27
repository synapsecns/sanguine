declare namespace JSX {
  interface IntrinsicElements {
    set: React.SVGProps<SVGElement> & { attributeName?: string; begin?: string }
  }
}
