export default function SynapseIcon({ width, height }) {
  return (
    <svg
      width={width}
      height={height}
      viewBox="0 0 48 48"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      className="w-8 h-8 md:w-10 md:h-10"
    >
      <defs>
        <linearGradient id="linear-gradient">
          <stop stop-color="#e54de5" />
          <stop offset="1" stop-color="#b580ff" />
        </linearGradient>
        <marker id="marker" viewBox="-1 -1 2 2">
          <circle r=".8" fill="url(#linear-gradient)" />
        </marker>
      </defs>
      <path
        d="M24,42 42,24 6,24 24,6"
        stroke="url(#linear-gradient)"
        stroke-width="5"
        stroke-linejoin="bevel"
        stroke-opacity=".5"
        marker-start="url(#marker)"
        marker-mid="url(#marker)"
        marker-end="url(#marker)"
      />
    </svg>
  )
}

export function SynapseAnchor() {
  return (
    <a
      href="#"
      className="text-2xl md:text-3xl font-medium flex gap-2 items-center"
    >
      <SynapseIcon width={40} height={40} />
      <span className="-mt-1.5">Synapse</span>
    </a>
  )
}
