import React from "react";

type Props = {
  questionId: number;
  questionBody: string;
  answerAttempt: string;
  setAnswerAttempt: (answerAttempt: string) => void;
};

const OpenAnswerCard = (props: Props) => {
    const [answerAttempt, setAnswerAttempt] = React.useState(props.answerAttempt);
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
        <label>Attempt</label>
        <textarea
          className="text-wrap break-words h-32 overflow-scroll bg-slate-50 shadow-inner p-1 mb-1          
          "
          value={answerAttempt}
            onChange={(e) => setAnswerAttempt(e.target.value)}
        />
        <button className="" type="submit" 
        // log the selected option
        onClick={() => console.log("submit")}>
            Next
        </button>
      </form>
    </div>
  );
};

export default OpenAnswerCard;
