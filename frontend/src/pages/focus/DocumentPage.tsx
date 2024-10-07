import React from "react";
import Select from "../../components/select/Select";
import ToggleCheck from "../../components/checkbox/ToggleCheck";
import { TiArrowSortedDown, TiArrowSortedUp } from "react-icons/ti";
import Dropdown from "../../components/dropdown/Dropdown";
import { IoMdMenu } from "react-icons/io";
import { useLocation } from "react-router-dom";
import { CrumbHelperPushLink } from "../../store/helpers/crumbHelper";
import DocumentHero from "../../components/hero/DocumentHero";
import TabGroup from "../../components/tab/TabGroup";
import ModalButton from "../../components/modal/ModalButton";
import DocMemoryCard from "../../components/cards/DocMemoryCard";
import EditQuestionCard, {
  EditQuestionProps,
} from "../../components/form/edit/EditQuestionCard";
import { QuestionType } from "../../types/question";
import DocQuestionCard, { DocQuestionCardProp } from "../../components/cards/DocQuestionCard";

const DocumentPage = () => {
  const tabs = ["Memories", "Questions"];
  const [isAsc, setIsAsc] = React.useState(false);
  const [chosenTab, setChosenTab] = React.useState(tabs[0]);
  const location = useLocation();
  const locationPath = location.pathname.split("/");
  while (
    locationPath.length > 0 &&
    locationPath[locationPath.length - 1] === ""
  ) {
    locationPath.pop();
  }
  const documentId = locationPath.pop() || "";
  CrumbHelperPushLink({ name: `Document:${documentId}` });
  const memories = [
    {
      title: "Math",
      documentId: "123",
      createdBy: "John Doe",
      createdAt: "2021-09-27",
      updatedAt: "2021-09-27",
      content: "hello",
    },
  ];

  const questions: DocQuestionCardProp[] = [
    {
      createdBy: "John Doe",
      createdAt: "2021-09-27",
      updatedAt: "2021-09-27",
      editQuestion: {
        questionType: QuestionType.singleAnswer,
        questionId: "123",
        questionBody: "What is 1 + 1?",
        optionList: [
          { optionId: 1, optionBody: "1" },
          { optionId: 2, optionBody: "2" },
          { optionId: 3, optionBody: "3" },
          { optionId: 4, optionBody: "4" },
        ],
        answer: "2",
        answerOption: 2,
        answerOptions: [],
      },
    },
    {
      createdBy: "John Doe",
      createdAt: "2021-09-27",
      updatedAt: "2021-09-27",
      editQuestion: {
        questionType: QuestionType.multiAnswer,
        questionId: "123",
        questionBody: "what countries are in Africa?",
        optionList: [
          { optionId: 1, optionBody: "Nigeria" },
          { optionId: 2, optionBody: "Ghana" },
          { optionId: 3, optionBody: "USA" },
          { optionId: 4, optionBody: "UK" },
        ],
        answer: "2",
        answerOption: 1,
        answerOptions: [1, 2],
      },
    },
    {
      createdBy: "John Doe",
      createdAt: "2021-09-27",
      updatedAt: "2021-09-27",
      editQuestion: {
        questionType: QuestionType.openEnded,
        questionId: "123",
        questionBody: "What is 1 + 1?",
        optionList: [],
        answer: "2",
        answerOption: 1,
        answerOptions: [],
      },
    },
  ];
  return (
    <div className="flex flex-col items-center relative z-0">
      <div className="sticky top-0  bg-base-100 w-full flex justify-center z-10">
        <DocumentHero
          title="Mathematics"
          description='Mathematics is the study of numbers, shapes and patterns. The word comes from the Greek word "μάθημα" (máthema), meaning "science, knowledge, or learning", and is sometimes shortened to maths (in England, Australia, Ireland, and New Zealand) or math (in the United States and Canada).Mathematics is the study of numbers, shapes and patterns. The word comes from the Greek word "μάθημα" (máthema), meaning "science, knowledge, or learning", and is sometimes shortened to maths (in England, Australia, Ireland, and New Zealand) or math (in the United States and Canada).Mathematics is the study of numbers, shapes and patterns. The word comes from the Greek word "μάθημα" (máthema), meaning "science, knowledge, or learning", and is sometimes shortened to maths (in England, Australia, Ireland, and New Zealand) or math (in the United States and Canada).Mathematics is the study of numbers, shapes and patterns. The word comes from the Greek word "μάθημα" (máthema), meaning "science, knowledge, or learning", and is sometimes shortened to maths (in England, Australia, Ireland, and New Zealand) or math (in the United States and Canada).Mathematics is the study of numbers, shapes and patterns. The word comes from the Greek word "μάθημα" (máthema), meaning "science, knowledge, or learning", and is sometimes shortened to maths (in England, Australia, Ireland, and New Zealand) or math (in the United States and Canada).Mathematics is the study of numbers, shapes and patterns. The word comes from the Greek word "μάθημα" (máthema), meaning "science, knowledge, or learning", and is sometimes shortened to maths (in England, Australia, Ireland, and New Zealand) or math (in the United States and Canada).Mathematics is the study of numbers, shapes and patterns. The word comes from the Greek word "μάθημα" (máthema), meaning "science, knowledge, or learning", and is sometimes shortened to maths (in England, Australia, Ireland, and New Zealand) or math (in the United States and Canada).Mathematics is the study of numbers, shapes and patterns. The word comes from the Greek word "μάθημα" (máthema), meaning "science, knowledge, or learning", and is sometimes shortened to maths (in England, Australia, Ireland, and New Zealand) or math (in the United States and Canada).'
          documentId={documentId}
          createdBy="John Doe"
          createdAt="2021-09-27"
          updatedAt="2021-09-27"
        />
      </div>
      <TabGroup chosenTab={chosenTab} tabs={tabs} setChosenTab={setChosenTab} />

      <ModalButton buttonName="Create Module">
        <div>
          <h1>Modal</h1>
          <p>Modal Content</p>
        </div>
      </ModalButton>

      <div className="top-20 flex flex-row w-96 justify-end items-center my-1">
        {/* TODO: Add actions like
        1. test subject -> modal: test configuration -> test (test will be saved in personal history)
        2. speak to AI mentor -> modal: chat with AI mentor (AI mentor given subject context)

        add remove action to the modules
        add create module action below the modules
        */}
        <span className="w-96 text-start font-bold">{chosenTab}:</span>
        <Select name="Order By" items={["Title", "Date"]} />
        <ToggleCheck
          isChecked={isAsc}
          checkName="Asc"
          checkIcon={<TiArrowSortedUp />}
          unCheckName="Desc"
          unCheckIcon={<TiArrowSortedDown />}
          onClick={() => setIsAsc(!isAsc)}
        />
      </div>
      <div className="grid grid-cols-1 gap-2 w-96 mx-4 mb-1 h-fit min-h-40 max-h-96 overflow-y-scroll border-t-2 border-b-2 border-gray-200 shadow-inset bg-secondary p-2 ">
        {chosenTab === tabs[0] && (
          <>
            {memories.length === 0 && (
              <span className="text-center">No modules found</span>
            )}
            {memories.map((memory, index) => (
              <DocMemoryCard
                key={index}
                title={memory.title}
                content={memory.content}
                memoryId={memory.documentId}
                createdBy={memory.createdBy}
                createdAt={memory.createdAt}
                updatedAt={memory.updatedAt}
              />
            ))}
          </>
        )}

        {chosenTab === tabs[1] && (
          <>
            {questions.length === 0 && (
              <span className="text-center">No questions found</span>
            )}
            {questions.map((question, index) => (
              <DocQuestionCard
                key={index}
                createdBy={question.createdBy}
                createdAt={question.createdAt}
                updatedAt={question.updatedAt}
                editQuestion={question.editQuestion}
              />
            ))}
          </>
        )}
      </div>
      <div className="sticky bottom-0 w-full bg-base-100">
        <Dropdown
          dropSymbol={<IoMdMenu />}
          dropName="Actions"
          items={["Test Subject", "Speak to AI Mentor"]}
        />
      </div>
    </div>
  );
};

export default DocumentPage;
