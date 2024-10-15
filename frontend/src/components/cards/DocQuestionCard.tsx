import { QuestionReturn } from "../../api/questionService";
import EditQuestionCard from "../form/edit/EditQuestionCard";
import ModalButton from "../modal/ModalButton";
import { unixToDateString } from '../../utilities/timeUtilities';


const DocQuestionCard = (props: QuestionReturn) => {
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
        <h1 className="truncate h-fit mb-1">
          {props.question_title}
        </h1>
        <span className="text-gray-400 font-bold flex  items-center">
          #{String(props.question_id).padStart(8, "0")}
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
            <span>{unixToDateString(props.created_time)}</span>
          </div>
          <div className="flex flex-col justify-between  w-full">
            <span>Last Update: </span>
            <span>{unixToDateString(props.updated_time)}</span>
          </div>
        </div>
        <ModalButton buttonName="Edit">
            <EditQuestionCard {...props}/>
        </ModalButton>
      </div>
    </div>
  );
};

export default DocQuestionCard;
