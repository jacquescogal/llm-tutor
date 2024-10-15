import React, { useEffect, useState } from "react";
import Select from "../../components/select/Select";
import ToggleCheck from "../../components/checkbox/ToggleCheck";
import { TiArrowSortedDown, TiArrowSortedUp } from "react-icons/ti";
import Dropdown from "../../components/dropdown/Dropdown";
import { IoMdMenu } from "react-icons/io";
import { useLocation } from "react-router-dom";
import { CrumbHelperPushLink } from "../../store/helpers/crumbHelper";
import DocumentHero from "../../components/hero/DocumentHero";
import TabGroup from "../../components/tab/TabGroup";
import DocMemoryCard from "../../components/cards/DocMemoryCard";
import DocQuestionCard from "../../components/cards/DocQuestionCard";
import { OrderByDirection, OrderByField, UploadStatus } from "../../types/enums";
import { DBMemory, getMemoriesByDocId, GetMemoriesByDocIdResponse } from "../../api/memoryService";
import { sortByMap } from "../../utilities/constants";
import { unixToDateString } from "../../utilities/timeUtilities";
import { getQuestionsByDocId, QuestionReturn } from "../../api/questionService";
import ChatCard from "../../components/form/ChatCard";
import { ID_TYPE } from "../../types/chat";
import ModalSpan from "../../components/modal/ModalSpan";
import { DBDoc, getDocumentById, GetDocumentByIdResponse } from '../../api/documentService';
import QuestionPage from "../questions/QuestionPage";

const DocumentPage = () => {
  const [docHeroDetails, setDocHeroDetails] =
    React.useState<DBDoc | null>(null);
  const tabs = ["Memories", "Questions"];
  const [isAsc, setIsAsc] = React.useState(false);
  const [sortBy, setSortBy] = React.useState<OrderByField>(
    OrderByField.ORDER_BY_FIELD_ID
  );
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
  locationPath.pop();
  const moduleId = locationPath.pop() || "";
  CrumbHelperPushLink({ name: `Document: #${documentId.padStart(5, "0")}`});
  const [memories, setMemories] = useState<DBMemory[]>([]);
  const [questions, setQuestions] = useState<QuestionReturn[]>([]);
  
  // export interface GetMemoriesByDocIdPayload {
  // page_number: number;
  
  // sort_by: OrderByField;
  // order: OrderByDirection;

  useEffect(() => {
    const fetchDocument = async () => {
      try {
        const response: GetDocumentByIdResponse = await getDocumentById(Number(moduleId),Number(documentId));
        console.log(response);
        setDocHeroDetails(response.doc);
      } catch (e) {
        if (e instanceof Error) {
          console.log(e);
        } else {
          console.log(e);
        }
      }
    };

    fetchDocument();
  }, []);

  useEffect(() => {
    const fetchDocuments = async () => {
      const moduleId = 1
      const payload = {
        page_number: 1,
        sort_by: sortBy,
        order: isAsc
          ? OrderByDirection.ORDER_BY_DIRECTION_ASC
          : OrderByDirection.ORDER_BY_DIRECTION_DESC,
      }
      try {
        const response: GetMemoriesByDocIdResponse = await getMemoriesByDocId(
          moduleId, Number(documentId), payload
      );
        console.log(response.memories);
        setMemories(response.memories);
      } catch (e) {
        if (e instanceof Error) {
          console.log(e);
        } else {
          console.log(e);
        }
      }
    };

    fetchDocuments();
  }, []);

  useEffect(() => {
    const fetchQuestions = async () => {
      const moduleId = 1
      const payload = {
        page_number: 1,
        sort_by: sortBy,
        order: isAsc
          ? OrderByDirection.ORDER_BY_DIRECTION_ASC
          : OrderByDirection.ORDER_BY_DIRECTION_DESC,
      }
      try {
        const response: QuestionReturn[] = await getQuestionsByDocId(
          moduleId, Number(documentId), payload
      );
        console.log(response);
        setQuestions(response);
      } catch (e) {
        if (e instanceof Error) {
          console.log(e);
        } else {
          console.log(e);
        }
      }
    };

    fetchQuestions();
  }, []);

  return (
    <div className="flex flex-col items-center relative z-0">
      <div className="sticky top-0  bg-base-100 w-full flex justify-center z-10">
        <DocumentHero
          title={docHeroDetails?.doc_name || "Loading..."}
          description={docHeroDetails?.doc_summary || "Loading..."}
          documentId={documentId}
          // createdBy="John Doe"
          status={docHeroDetails && docHeroDetails?.upload_status || UploadStatus.UPLOAD_STATUS_PROCESSING}
          createdAt={docHeroDetails && unixToDateString(docHeroDetails?.created_time) || "Loading..."}
          updatedAt={docHeroDetails && unixToDateString(docHeroDetails?.updated_time) || "Loading..."}
        />
      </div>
      <TabGroup chosenTab={chosenTab} tabs={tabs} setChosenTab={setChosenTab} />

      

      <div className="top-20 flex flex-row w-96 justify-end items-center my-1">
        {/* TODO: Add actions like
        1. test subject -> modal: test configuration -> test (test will be saved in personal history)
        2. speak to AI mentor -> modal: chat with AI mentor (AI mentor given subject context)

        add remove action to the modules
        add create module action below the modules
        */}
        <span className="w-96 text-start font-bold">{chosenTab}:</span>
        <Select name="Order By" items={Object.keys(sortByMap)} setSelected={s=>{setSortBy(sortByMap[s])}} />
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
            {!memories || memories.length === 0 && (
              <span className="text-center">No modules found</span>
            )}
            {memories && memories.map((memory, index) => (
              <DocMemoryCard
                key={index}
                title={memory.memory_title}
                content={memory.memory_content}
                memoryId={String(memory.memory_id)}

                // createdBy={u
                createdAt={unixToDateString(memory.created_time)}
                updatedAt={unixToDateString(memory.updated_time)}
              />
            ))}
          </>
        )}

        {chosenTab === tabs[1] && (
          <>
            {!questions || questions.length === 0 && (
              <span className="text-center">No questions found</span>
            )}
            {questions && questions.map((question, index) => (
              <DocQuestionCard
                key={index}
                {...question}
              />
            ))}
          </>
        )}
      </div>
      <div className="sticky bottom-0 w-full bg-base-100">
        <Dropdown
          dropSymbol={<IoMdMenu />}
          dropName="Actions"
          items={[<ModalSpan buttonName="Test Me" className="w-40 py-2 align-text-left text-left">
            <QuestionPage 
            />
          </ModalSpan>, 
            <ModalSpan buttonName="Speak to AI Mentor" className="w-40 py-2 align-text-left text-left">
            <ChatCard 
            id={Number(documentId)}
            id_type={ID_TYPE.DOCUMENT}
            />
          </ModalSpan>]}
        />
      </div>
    </div>
  );
};

export default DocumentPage;
