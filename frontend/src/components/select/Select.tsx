import React from 'react'

type Props = {
    name: string;
    items: string[];
}

const Select = (props: Props) => {
  return (
    <select className="select select-bordered select-sm w-fit ">
  <option disabled selected>{props.name}</option>
  {
    props.items.map((item, index) => (
      <option key={index}>{item}</option>
    ))
  }
</select>
  )
}

export default Select