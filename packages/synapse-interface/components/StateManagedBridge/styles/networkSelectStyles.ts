import tailwindConfig from '../../../tailwind.config.js'

const {
  theme: {
    extend: { colors },
  },
} = tailwindConfig

export const networkSelectStyles = {
  control: (baseStyles, state) => ({
    ...baseStyles,
    backgroundColor: 'none',
    border: 'none',
    width: state.hasValue ? '205px' : '150px',
  }),
  menu: (provided) => ({
    ...provided,
    width: 300,
    backgroundColor: colors.bgLight,
    fontWeight: '300',
  }),
  menuList: (provided) => ({
    ...provided,
    '::-webkit-scrollbar': {
      display: 'none',
    },
    msOverflowStyle: 'none',
    scrollbarWidth: 'none',
    overflow: 'auto',
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
  option: (provided, state) => {
    const backgroundColor =
      state.isSelected || state.isFocused
        ? colors.bgBase
        : provided.backgroundColor
    return {
      ...provided,
      backgroundColor,
      outline: 'none',
      ':active': {
        backgroundColor: state.isSelected ? colors.bgBase : colors.bgLighter,
      },
    }
  },
}
