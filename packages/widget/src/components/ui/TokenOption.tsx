import { BridgeableToken } from 'types'

export const TokenOption = ({
  option,
  onSelect,
  selected,
  parsedBalance,
}: {
  option: BridgeableToken
  onSelect: (option: BridgeableToken) => void
  selected: BridgeableToken
  parsedBalance: string
}) => {
  return (
    <li
      data-test-id="token-option"
      className={`
        flex gap-4 items-center justify-between
        cursor-pointer rounded border border-solid
        hover:border-[--synapse-focus] active:opacity-40 group
        ${
          option?.symbol === selected?.symbol
            ? 'border-[--synapse-focus] hover:opacity-70'
            : 'border-transparent'
        }
      `}
      onClick={() => onSelect(option)}
    >
      <abbr
        title={option?.name}
        className="p-2.5 no-underline flex items-center"
      >
        {option?.imgUrl && (
          <img
            src={option?.imgUrl}
            alt={`${option?.symbol} token icon`}
            className="inline w-4 h-4 mr-2"
          />
        )}
        {option?.symbol}
      </abbr>
      <data
        value={parsedBalance}
        className={`
          text-sm p-2.5
          ${
            parsedBalance
              ? 'text-[--synapse-secondary]'
              : 'text-[--synapse-focus]'
          }
        `}
      >
        {parsedBalance ? (
          parsedBalance === '0.0' ? (
            '−'
          ) : (
            parsedBalance
          )
        ) : (
          <span className="opacity-0 text-sm text-[--synapse-secondary] group-hover:opacity-100">
            Receive
          </span>
        )}
      </data>
    </li>
  )
}
