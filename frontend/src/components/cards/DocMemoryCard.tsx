import EditMemoryCard from "../form/edit/EditMemoryCard";
import ModalButton from "../modal/ModalButton";

type Props = {
  title: string;
  content: string;
  memoryId: string;
  createdAt: string;
  updatedAt: string;
}

const DocMemoryCard = (props: Props) => {
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
          #{props.memoryId.padStart(8, "0")}
        </span>
      </div>
      <div className="flex flex-row justify-between align-middle content-center text-center align-center ">
        <div className="w-full flex flex-row text-left text-xs justify-between select-none text-gray-500">
          {/* <div className="flex flex-col justify-between  w-full">
            <span>Uploaded By: </span>
            <span>{props.createdBy}</span>
          </div> */}
          <div className="flex flex-col justify-between  w-full">
            <span>Uploaded On: </span>
            <span>{props.createdAt}</span>
          </div>
          <div className="flex flex-col justify-between  w-full">
            <span>Last Update: </span>
            <span>{props.updatedAt}</span>
          </div>
        </div>
        <ModalButton
        buttonName="Edit"
        >
          <EditMemoryCard
            memoryId={props.memoryId}
            memoryTitle={props.title}
            memoryContent={props.content}
          />
        </ModalButton>
      </div>
    </div>
  );
};

export default DocMemoryCard;
