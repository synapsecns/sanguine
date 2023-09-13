export default function ProcessingIcon({ ...props }) {
  return (
    <svg
      height="2"
      viewBox="0 0 10 2"
      preserveAspectRatio="none"
      xmlns="http://www.w3.org/2000/svg"
      {...props}
    >
      <rect width="10" height="2" fill="var(--separator)">
        <animate
          attributeName="x"
          values="-10; 11"
          dur="1.6s"
          keyTimes="0; 1"
          calcMode="spline"
          keySplines=".5 0 .2 1"
          repeatCount="indefinite"
        />
      </rect>
    </svg>
  )
}
