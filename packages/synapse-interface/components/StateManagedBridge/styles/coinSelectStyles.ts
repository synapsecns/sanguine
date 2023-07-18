import tailwindConfig from '../../../tailwind.config.js'

const {
  theme: {
    extend: { colors },
  },
} = tailwindConfig

export const coinSelectStyles = {
  control: (baseStyles) => ({
    ...baseStyles,
    backgroundColor: colors.bgLightest,
    border: 'none',
    height: '40px',
    width: '150px',
  }),
  menu: (provided) => ({
    ...provided,
    width: 300,
  }),
  input: (provided) => ({
    ...provided,
    color: 'white',
  }),
  singleValue: (provided) => ({
    ...provided,
    color: 'white',
  }),
  indicatorSeparator: (provided) => ({
    ...provided,
    display: 'none',
  }),
}
