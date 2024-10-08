import React, { useState } from "react";
import { createModule } from "../../../api/moduleService";


const CreateModuleCard = () => {
  const [moduleTitle, setModuleTitle] = useState("");
  const [moduleContent, setModuleContent] = useState("");
  const onCreateClick = async()=>{
    try {
      await createModule({ module_name: moduleTitle, module_description: moduleContent, is_public: true });
      console.log("Subject Created");
    } catch (e) {
      if (e instanceof Error) {
        console.log(e.message);
      } else {
        console.log("An unknown error occurred");
      }
    }
  }
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
      <h1 className="truncate h-fit mb-1">Create Module</h1>
      <span className="font-bold">Title:</span>
      <textarea
        className="text-wrap break-words h-20 overflow-scroll bg-slate-50 shadow-inner p-1 mb-1          
          "
        onChange={(e) => setModuleTitle(e.target.value)}
        value={moduleTitle}
      />
      <div
        className="flex flex-col"
        onSubmit={() => {
          console.log("submit");
        }}
      >
        <span className="font-bold">Description:</span>
        <textarea
          className="text-wrap break-words h-20 overflow-scroll bg-slate-50 shadow-inner p-1 mb-1          
          "
          onChange={(e) => setModuleContent(e.target.value)}
          value={moduleContent}
        />

        <button
          className="mb-2 p-6"
          // log the selected option
          onClick={(e) => {
            e.preventDefault();
            console.log("Update");
            onCreateClick();
          }}
        >
            Create Module
        </button>
        <button className="mb-1 p-4" type="submit">
          Cancel
        </button>
      </div>
    </div>
  );
};

export default CreateModuleCard;
