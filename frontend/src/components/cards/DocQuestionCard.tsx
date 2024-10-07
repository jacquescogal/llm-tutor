import EditQuestionCard, { EditQuestionProps } from "../form/edit/EditQuestionCard";
import ModalButton from "../modal/ModalButton";

export type DocQuestionCardProp = {
  createdBy: string;
  createdAt: string;
  updatedAt: string;
  editQuestion: EditQuestionProps;
};

const DocQuestionCard = (props: DocQuestionCardProp) => {
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
          {props.editQuestion.questionBody}
        </h1>
        <span className="text-gray-400 font-bold flex  items-center">
          #{props.editQuestion.questionId.padStart(8, "0")}
        </span>
      </div>
      <div className="flex flex-row justify-between align-middle content-center text-center align-center ">
        <div className="w-full flex flex-row text-left text-xs justify-between select-none text-gray-500">
          <div className="flex flex-col justify-between  w-full">
            <span>Uploaded By: </span>
            <span>{props.createdBy}</span>
          </div>
          <div className="flex flex-col justify-between  w-full">
            <span>Uploaded On: </span>
            <span>{props.createdAt}</span>
          </div>
          <div className="flex flex-col justify-between  w-full">
            <span>Last Update: </span>
            <span>{props.updatedAt}</span>
          </div>
        </div>
        <ModalButton buttonName="Edit">
            <EditQuestionCard {...props.editQuestion} />
        </ModalButton>
      </div>
    </div>
  );
};

export default DocQuestionCard;
