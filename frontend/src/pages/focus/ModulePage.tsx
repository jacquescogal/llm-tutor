import React, { useEffect } from "react";
import Select from "../../components/select/Select";
import ToggleCheck from "../../components/checkbox/ToggleCheck";
import { TiArrowSortedDown, TiArrowSortedUp } from "react-icons/ti";
import Dropdown from "../../components/dropdown/Dropdown";
import { IoMdMenu } from "react-icons/io";
import { useLocation } from "react-router-dom";
import { CrumbHelperPushLink } from "../../store/helpers/crumbHelper";
import DocumentCard from "../../components/cards/DocumentCard";
import ModuleHero from "../../components/hero/ModuleHero";
import {
  FullModule,
  getModuleById,
  GetModuleByIdResponse,
} from "../../api/moduleService";
import {
  DBDoc,
  getDocumentsByModuleId,
  GetDocumentsByModuleIdResponse,
} from "../../api/documentService";
import { OrderByDirection, OrderByField } from "../../types/enums";
import { unixToDateString } from "../../utilities/timeUtilities";
import { sortByMap } from "../../utilities/constants";

const ModulePage = () => {
  const [moduleHeroDetails, setModuleHeroDetails] =
    React.useState<FullModule | null>(null);
  const [documents, setDocuments] = React.useState<DBDoc[]>([]);
  const [isAsc, setIsAsc] = React.useState(false);
  const [sortBy, setSortBy] = React.useState<OrderByField>(
    OrderByField.ORDER_BY_FIELD_ID
  );
  const location = useLocation();
  const locationPath = location.pathname.split("/");
  while (
    locationPath.length > 0 &&
    locationPath[locationPath.length - 1] === ""
  ) {
    locationPath.pop();
  }
  const moduleId = locationPath.pop() || "";
  CrumbHelperPushLink({ name: `Module: #${moduleId.padStart(5, "0")}` });
  useEffect(() => {
    const fetchModule = async () => {
      try {
        const response: GetModuleByIdResponse = await getModuleById({
          module_id: Number(moduleId),
        });
        console.log(response);
        setModuleHeroDetails(response.module);
      } catch (e) {
        if (e instanceof Error) {
          console.log(e);
        } else {
          console.log(e);
        }
      }
    };

    fetchModule();
  }, []);

  const fetchDocuments = async () => {
    try {
      const response: GetDocumentsByModuleIdResponse =
        await getDocumentsByModuleId({
          module_id: Number(moduleId),
          page_number: 1,
          sort_by: sortBy,
          order: isAsc
            ? OrderByDirection.ORDER_BY_DIRECTION_ASC
            : OrderByDirection.ORDER_BY_DIRECTION_DESC,
        });
      setDocuments(response.docs);
    } catch (e) {
      if (e instanceof Error) {
        console.log(e);
      } else {
        console.log(e);
      }
    }
  };

  useEffect(() => {
    fetchDocuments();
  }, [isAsc, sortBy]);
  return (
    <div className="flex flex-col items-center relative z-0">
      <div className="sticky top-0  bg-base-100 w-full flex justify-center z-10">
        {moduleHeroDetails && (
          <ModuleHero
            title={moduleHeroDetails.module.module_name}
            description={moduleHeroDetails.module.module_description}
            moduleId={String(moduleHeroDetails.module.module_id)}
            // createdBy={moduleHeroDetails.}
            createdAt={unixToDateString(moduleHeroDetails.module.created_time)}
            updatedAt={unixToDateString(moduleHeroDetails.module.updated_time)}
            isFavourite={moduleHeroDetails.is_favourite}
          />
        )}
      </div>
      <div className="top-20 flex flex-row w-96 justify-end items-center my-1">
        {/* TODO: Add actions like
        1. test subject -> modal: test configuration -> test (test will be saved in personal history)
        2. speak to AI mentor -> modal: chat with AI mentor (AI mentor given subject context)

        add remove action to the modules
        add create module action below the modules
        */}
        <span className="w-96 text-start font-bold">Modules:</span>
        <Select name="Order By" items={Object.keys(sortBy)} setSelected={s=>{setSortBy(sortByMap[s])}} />
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
        {!documents || documents.length === 0 && (
          <span className="text-center">No modules found</span>
        )}
        {documents && documents.map((doc, index) => (
          <DocumentCard
            key={index}
            title={doc.doc_name}
            documentId={String(doc.doc_id)}
            // createdBy={doc.created_time
            createdAt={unixToDateString(doc.created_time)}
            updatedAt={unixToDateString(doc.updated_time)}
            // memoryCount={0}
            // questionCount={0}
          />
        ))}
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

export default ModulePage;
