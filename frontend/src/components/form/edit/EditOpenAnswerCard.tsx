import { useState } from "react";
import { QuestionReturn, TextInputQuestion, unserializeQuestion } from "../../../api/questionService";


const EditOpenAnswerCard = (props:QuestionReturn) => {
  const actQuestion:TextInputQuestion = unserializeQuestion(props) as TextInputQuestion;
  const [questionBody, setQuestionBody] = useState(props.question_title);
  const [answer, setAnswer] = useState(actQuestion.answer);

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
      <h1 className="truncate h-fit mb-1">Question {props.question_id}</h1>
      <textarea
        className="text-wrap break-words h-20 overflow-scroll bg-slate-50 shadow-inner p-1 mb-1          
          "
        onChange={(e) => setQuestionBody(e.target.value)}
        value={questionBody}
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
          onChange={(e) => setAnswer(e.target.value)}
          value={answer}
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

export default EditOpenAnswerCard;
