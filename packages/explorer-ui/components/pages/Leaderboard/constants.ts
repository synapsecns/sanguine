// may be able to rewrite this with some tailwind resolve config magic
export const directColors = {
  yellow: '#ecae0b',
  eth: '#5170ad',
  purple: 'rgb(168 85 247)',
  blue: 'rgb(59 130 246)',
  lime: 'rgb(132 204 22)',
  gray: 'rgb(107 114 128)',
  red: 'rgb(239 68 68)',
  cyan: 'rgb(6 182 212)',
  teal: 'rgb(20 184 166)',
  sky: 'rgb(2 132 199)',
  green: 'rgb(34 197 94)',
}

export const chartOptions = {
  indexAxis: 'y',
  responsive: true,
  plugins: {
    legend: {
      display: false,
    },
  },
  animation: {
    duration: 0,
  },
}
