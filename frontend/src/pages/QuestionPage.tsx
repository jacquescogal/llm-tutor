import React from "react";
import SingleAnswerCard from "../components/cards/questions/SingleAnswerCard";
import MultiAnswerCard from "../components/cards/questions/MultiAnswerCard";
import OpenAnswerCard from "../components/cards/questions/OpenAnswerCard";

enum QuestionType {
    singleAnswer = "singleAnswer",
    multiAnswer = "multiAnswer",
    openEnded = "openEnded"
}

interface SingleAnswerQuestion {
    questionBody: string;
    optionList: { optionId: number; optionBody: string }[];
    selectedOption: number;
}

interface MultiAnswerQuestion {
    questionBody: string;
    optionList: { optionId: number; optionBody: string }[];
    selectedOptions: number[];
}

interface OpenEndedQuestion {
    questionBody: string;
    answerAttempt: string;
}

interface QuestionPointer {
    type: QuestionType;
    subListIndex: number;
}

const QuestionPage = () => {
    const singleAnsQuestionList = React.useRef<SingleAnswerQuestion[]>([
        {
            questionBody: "What is the capital of France?",
            optionList: [
                { optionId: 1, optionBody: "Paris" },
                { optionId: 2, optionBody: "London" },
                { optionId: 3, optionBody: "Berlin" },
                { optionId: 4, optionBody: "Madrid" },
            ],
            selectedOption: 0,
        },
    ]);
    const multiAnsQuestionList = React.useRef<MultiAnswerQuestion[]>([
        {
            questionBody: "What are the countries in Europe?",
            optionList: [
                { optionId: 1, optionBody: "France" },
                { optionId: 2, optionBody: "UK" },
                { optionId: 3, optionBody: "Germany" },
                { optionId: 4, optionBody: "Singapore" },
            ],
            selectedOptions: [],
        },
    ]);

    const openEndedQuestionList = React.useRef<OpenEndedQuestion[]>([
        {
            questionBody: "What is the capital of France?",
            answerAttempt: "",
        },
    ]);
  const questionList = React.useRef<QuestionPointer[]>([
    {
        type: QuestionType.singleAnswer,
        subListIndex: 0
    },
    {
        type: QuestionType.multiAnswer,
        subListIndex: 0
    },
    {
        type: QuestionType.openEnded,
        subListIndex: 0
    }
  ]);
  return (
    <div>
      {questionList.current.map((question, index) => {
        return (
            <div key={index}>
            {
                question.type === QuestionType.singleAnswer &&
                <SingleAnswerCard
                key={index}
                questionId={index}
                questionBody={singleAnsQuestionList.current[question.subListIndex].questionBody}
                optionList={singleAnsQuestionList.current[question.subListIndex].optionList}
                selectedOption={singleAnsQuestionList.current[question.subListIndex].selectedOption}
                setSelectedOption={(optionId) => {
                    singleAnsQuestionList.current[question.subListIndex].selectedOption = optionId;
                }}
              />
            }
            {
                question.type === QuestionType.multiAnswer &&
                <MultiAnswerCard
                key={index}
                questionId={index}
                questionBody={multiAnsQuestionList.current[question.subListIndex].questionBody}
                optionList={multiAnsQuestionList.current[question.subListIndex].optionList}
                selectedOptions={multiAnsQuestionList.current[question.subListIndex].selectedOptions}
                setSelectedOptions={(optionIds) => {
                    multiAnsQuestionList.current[question.subListIndex].selectedOptions = optionIds;
                }}
              />
            }
            {
                question.type === QuestionType.openEnded &&
                <OpenAnswerCard
                key={index}
                questionId={index}
                questionBody={openEndedQuestionList.current[question.subListIndex].questionBody}
                answerAttempt={openEndedQuestionList.current[question.subListIndex].answerAttempt}
                setAnswerAttempt={(answerAttempt) => {
                    openEndedQuestionList.current[question.subListIndex].answerAttempt = answerAttempt;
                }}
              />
            }
            </div>
    )
}
      )}
    </div>
  );
};

export default QuestionPage;
