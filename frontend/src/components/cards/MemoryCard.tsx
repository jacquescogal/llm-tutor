import React from "react"

const MemoryCard = (props: {
    title: string,
    body: string,
    buttons: React.ReactNode[],
}) => {
  return (
    <div className="
    h-64 w-96 
    flex flex-col 
    bg-base-100
    text-base-content
    relative 
    border border-2 border-black
    text-left
    p-2
    text-wrap break-normal
    overflow-hidden
    text-ellipsis
    ">
        <h1 className="truncate h-10">{props.body}</h1>
        <span className="text-wrap break-words h-36 overflow-scroll bg-slate-50 shadow-inner">{props.body}</span>
        <div className="absolute bottom-0 right-0 p-2 flex flex-row h-16">
        {props.buttons.map((button, index) => (
            <div key={index} >
                {button}
            </div>
        ))}
        </div>
    </div>
  )
}

export default MemoryCard