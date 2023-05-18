
export function TableRow({ items, key }) {

  return (

    // @ts-expect-error TS(2304): Cannot find name 'tr'.
    <tr key={key} className="transition ease-out">

      {items.map((item) =>

        // @ts-expect-error TS(2304): Cannot find name 'td'.
        <td className="whitespace-nowrap px-2 py-2 text-sm  text-white">
          // @ts-expect-error TS(2552): Cannot find name 'item'. Did you mean 'items'?
          {item}
        </td>
      )
      }
    </tr>
  )
}

// TODO add animations to updated table
// import { Transition } from 'react-transition-group'
// export function TableRow({ items, key }) {

//   const defaultStyle = {
//     transition: `opacity 500ms ease-in-out`,
//     opacity: 0,
//   }

//   const transitionStyles = {
//     entering: { opacity: 1 },
//     entered: { opacity: 1 },
//     exiting: { opacity: 0 },
//     exited: { opacity: 0 },
//   };
//   let rowItems = (<>{items.map((item) => {
//     <td className="whitespace-nowrap px-2 py-2 text-sm  text-white" >
//       {item}
//     </td>
//   })}</>)
//   console.log(items, rowItems)
//   return (
//     <Transition
//       in={true}
//       appear={true}
//       timeout={{
//         appear: 500,
//         enter: 300,
//         exit: 500,
//       }}

//     >{state => (
//       <tr key={key} className="transition ease-out" style={{
//         ...defaultStyle,
//         ...transitionStyles[state]
//       }}>
//         {rowItems}


//       </tr>)}

//     </Transition>
//   )
// }
