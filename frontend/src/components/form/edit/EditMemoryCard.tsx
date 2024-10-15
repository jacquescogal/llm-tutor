import React, { useState } from "react";

type Props = {
  memoryId: string;
  memoryTitle: string;
  memoryContent: string;
};

const EditMemoryCard = (props: Props) => {
  const [memoryTitle, setMemoryTitle] = useState(props.memoryTitle);
  const [memoryContent, setMemoryContent] = useState(props.memoryContent);
  // TODO: allow memory edit 
  return (
    <div
      className="
        h-fit w-full
        flex flex-col 
        bg-base-100
        text-base-content
        relative 
        text-left
        p-2
        text-wrap break-normal
        overflow-hidden
        text-ellipsis
        "
    >
      <h1 className="truncate h-fit mb-1">Memory {props.memoryId}</h1>
      <textarea
        className="text-wrap break-words h-20 overflow-scroll bg-slate-50 shadow-inner p-1 mb-1          
          "
        onChange={(e) => setMemoryTitle(e.target.value)}
        value={memoryTitle}
      />
      <div
        className="flex flex-col"
        onSubmit={() => {
          console.log("submit");
        }}
      >
        <span>Answer:</span>
        <textarea
          className="text-wrap break-words h-20 overflow-scroll bg-slate-50 shadow-inner p-1 mb-1          
          "
          onChange={(e) => setMemoryContent(e.target.value)}
          value={memoryContent}
        />

        <button
          className="mb-1"
          // log the selected option
          onClick={(e) => {
            e.preventDefault();
            console.log("Update");
          }}
        >
          Save Changes
        </button>
        <button
          className="mb-1"
          // log the selected option
        >
          Delete
        </button>
        <button className="mb-1" type="submit">
          Cancel
        </button>
      </div>
    </div>
  );
};

export default EditMemoryCard;
