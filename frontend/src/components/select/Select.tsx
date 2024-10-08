import React from 'react'

type Props = {
    name: string;
    items: string[];
    setSelected: (value: string) => void;
}

const Select = (props: Props) => {
  return (
    <select className="select select-bordered select-sm w-fit "
    
    onChange={(e) => {props.setSelected(e.target.value);}}
    >
  <option disabled selected>{props.name}</option>
  {
    props.items.map((item, index) => (
      <option 
      key={index}>{item}</option>
    ))
  }
</select>
  )
}

export default Select