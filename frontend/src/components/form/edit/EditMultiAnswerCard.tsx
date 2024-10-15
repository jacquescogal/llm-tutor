import { useState } from "react";
import { MCQQuestion, QuestionReturn, unserializeQuestion } from "../../../api/questionService";


const EditMuliAnswerCard = (props: QuestionReturn) => {
  const actQuestion:MCQQuestion = unserializeQuestion(props) as MCQQuestion;
  const [questionBody, setQuestionBody] = useState(props.question_title);
  const [optionList, setOptionList] = useState(actQuestion.choices.map((choice, index) => {
    return {
      optionId: index,
      choice: choice.choice,
      isCorrect: choice.isCorrect,
    };
  }));

  const updateOption = (optionId: number, optionBody: string) => {
    setOptionList(
      optionList?.map((option) => {
        if (option.optionId === optionId) {
          return {
            ...option,
            choice: optionBody,
          };
        }
        return option;
      })
    );
  }

  const addOption = () => {
    const newOptionId = optionList ? optionList.length + 1 : 1;
    const newOptionList = optionList ? optionList : [];
    setOptionList([
      ...newOptionList,
      {
        optionId: newOptionId,
        choice: "",
        isCorrect: false,
      },
    ]);
  }

  const deleteOption = (optionId: number) => {
    setOptionList(
      optionList?.filter((option) => option.optionId !== optionId)
    );
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
        <span>Select True Options:</span>
        <div className=" p-2 flex flex-col h-40 overflow-y-scroll bg-secondary">
          {optionList?.length === 0 && <span>No options available</span>}
          {optionList?.map((option, index) => (
            <div key={index} className="group">
              <label className="label cursor-pointer w-full h-8 flex flex-row items-center align-center content-center justify-center mb-1">
                <input
                  type="checkbox"
                  className="checkbox"
                  checked={option.isCorrect}
                  onChange={() => {
                    setOptionList(
                      optionList?.map((opt) => {
                        if (opt.optionId === option.optionId) {
                          return {
                            ...opt,
                            isCorrect: !opt.isCorrect,
                          };
                        }
                        return opt;
                      })
                    );
                  }}
                />
                <input
                  className="w-full text-left mx-1 py-1"
                  value={option.choice}
                  onChange={(e) => {
                    updateOption(option.optionId, e.target.value);
                  }}
                />
                <button
                  className="items-center content-center justify-center"
                  onClick={(e) => {
                    e.preventDefault();
                    deleteOption(option.optionId);
                  }}
                >
                  Delete
                </button>
              </label>
            </div>
          ))}
        </div>
        <button
          className="mb-1"
          onClick={(e) => {
            e.preventDefault();
            addOption();
          }}
        >
          Add Option
        </button>
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

export default EditMuliAnswerCard;
