import { useState } from "react";
import { QuestionOption } from "../../../types/question";

type Props = {
  questionId: string;
  questionBody: string;
  optionList?: QuestionOption[];
  answerOptions?: number[];
};

const EditMuliAnswerCard = (props: Props) => {
  const [selectedOptions, setSelectedOptions] = useState(props.answerOptions);
  const [questionBody, setQuestionBody] = useState(props.questionBody);
  const [optionList, setOptionList] = useState(props.optionList);

  const updateOption = (optionId: number, optionBody: string) => {
    setOptionList(
      optionList?.map((option) => {
        if (option.optionId === optionId) {
          return { ...option, optionBody };
        }
        return option;
      })
    );
  };

  const addOption = () => {
    const newOptionId = optionList ? optionList.length + 1 : 1;
    const newOptionList = optionList ? optionList : [];
    setOptionList([
      ...newOptionList,
      {
        optionId: newOptionId,
        optionBody: "",
      },
    ]);

  };

  const deleteOption = (optionId: number) => {

    const newSelectedOptions = selectedOptions?.includes(
      optionId
    )
      ? selectedOptions?.filter(
          (optionIdItem) => optionIdItem !== optionId
        )
      : [...selectedOptions||[1]];
      console.log(newSelectedOptions)
    setSelectedOptions(newSelectedOptions);
    setOptionList(optionList?.filter((option) => option.optionId !== optionId));
  };

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
      <h1 className="truncate h-fit mb-1">Question {props.questionId}</h1>
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
                  checked={selectedOptions?.includes(option.optionId)}
                  onChange={() => {
                    const newSelectedOptions = selectedOptions?.includes(
                      option.optionId
                    )
                      ? selectedOptions?.filter(
                          (optionId) => optionId !== option.optionId
                        )
                      : [...selectedOptions|| [], option.optionId];

                    setSelectedOptions(newSelectedOptions);
                  }}
                />
                <input
                  className="w-full text-left mx-1 py-1"
                  value={option.optionBody}
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
