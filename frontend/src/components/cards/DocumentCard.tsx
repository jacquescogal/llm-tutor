import React from "react";
import { useLocation, useNavigate } from "react-router-dom";

type Props = {
  title: string;
  documentId: string;
  createdBy?: string;
  createdAt: string;
  updatedAt: string;
  memoryCount?: number;
  questionCount?: number;
};

const DocumentCard = (props: Props) => {
  const nav = useNavigate();
  const location = useLocation();
  return (
    <div
      className="
      h-fit w-96
      max-w-full
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
      animate-appear-bot
      "
    >
      <div className="flex flex-row justify-between align-middle content-center text-center align-center  ">
        <h1 className="truncate h-fit mb-1">{props.title}</h1>
        <span className="text-gray-400 font-bold flex  items-center">
          #{props.documentId.padStart(8, "0")}
        </span>
      </div>
      {
        props.memoryCount && props.questionCount && (
      
      <div className="w-full flex flex-row text-left text-xs justify-between select-none text-gray-500 bg-info text-info-content p-1">
        <div className="flex flex-row  w-full">
          <span className="font-bold mr-1">Memory Count:</span>
          <span>{props.memoryCount}</span>
        </div>
        <div className="flex flex-row  w-full">
          <span className="font-bold mr-1">Question Count:</span>
          <span>{props.questionCount}</span>
        </div>
      </div>
        )}
      <div className="flex flex-row justify-between align-middle content-center text-center align-center ">
        <div className="w-full flex flex-row text-left text-xs justify-between select-none text-gray-500">
          {
            props.createdBy && (
          
          <div className="flex flex-col justify-between  w-full">
            <span>Uploaded By: </span>
            <span>{props.createdBy}</span>
          </div>
            )}
          <div className="flex flex-col justify-between  w-full">
            <span>Uploaded On: </span>
            <span>{props.createdAt}</span>
          </div>
          <div className="flex flex-col justify-between  w-full">
            <span>Last Update: </span>
            <span>{props.updatedAt}</span>
          </div>
        </div>
        <button
          className=""
          type="submit"
          // log the selected option
          // TODO: navigate to the next module
          onClick={() =>
            nav(
              (location.pathname.endsWith("document") ? "" : "document/") +
                props.documentId
            )
          }
        >
          Edit
        </button>
      </div>
    </div>
  );
};

export default DocumentCard;
