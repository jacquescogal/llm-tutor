import React from "react";
import { QuestionOption } from "../../../types/question";


type Props = {
  questionId: number;
  questionBody: string;
  optionList: QuestionOption[];
  selectedOption: number | undefined;
  setSelectedOption: (optionId: number) => void;
};

const SingleAnswerCard = (props: Props) => {
    const [selectedOption, setSelectedOption] = React.useState(props.selectedOption);
  return (
    <div
      className="
        h-fit w-96 
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
        "
    >
      <h1 className="truncate h-fit mb-1">Question {props.questionId+1}</h1>
      <span className="text-wrap break-words h-20 overflow-scroll bg-slate-50 ">
        {props.questionBody}
      </span>
      <form
        className="flex flex-col"
        onSubmit={(e) => {
          e.preventDefault();
          console.log("submit");
        }}
      >
        <div className=" p-2 flex flex-col h-fit">
          {props.optionList.map((option, index) => (
            <div key={index}>
              <label className="label cursor-pointer w-full mb-0 h-8 flex flex-row">
                <input
                  type="radio"
                  name="radioOption"
                  className="radio"
                  checked={selectedOption === option.optionId}
                  onChange={() => {props.setSelectedOption(option.optionId); setSelectedOption(option.optionId)}}
                />
                <span className="w-full text-left label-text ml-1 truncate">
                  {option.optionBody}
                </span>
              </label>
            </div>
          ))}
        </div>
        <button className="" type="submit" 
        // log the selected option
        onClick={() => console.log("submit")}>
            Next
        </button>
      </form>
    </div>
  );
};

export default SingleAnswerCard;
