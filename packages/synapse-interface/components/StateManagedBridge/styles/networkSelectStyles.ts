export const networkSelectStyles = {
  control: (baseStyles, state) => ({
    ...baseStyles,
    backgroundColor: 'none',
    border: 'none',
    width: state.hasValue ? '250px' : '150px',
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
