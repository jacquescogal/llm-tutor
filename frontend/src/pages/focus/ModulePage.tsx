import React from "react";
import Select from "../../components/select/Select";
import ToggleCheck from "../../components/checkbox/ToggleCheck";
import { TiArrowSortedDown, TiArrowSortedUp } from "react-icons/ti";
import Dropdown from "../../components/dropdown/Dropdown";
import { IoMdMenu } from "react-icons/io";
import { useLocation } from "react-router-dom";
import { CrumbHelperPushLink } from "../../store/helpers/crumbHelper";
import DocumentCard from "../../components/cards/DocumentCard";
import ModuleHero from "../../components/hero/ModuleHero";

const ModulePage = () => {
  const [isAsc, setIsAsc] = React.useState(false);
  const location = useLocation();
  const locationPath = location.pathname.split('/');
  while (locationPath.length>0 && locationPath[locationPath.length - 1] === "") {
    locationPath.pop();
  }
  const moduleId = locationPath.pop() || "";
  CrumbHelperPushLink({ name: `Module:${moduleId}` });
  const modules = [
    {
      title: "Math",
      documentId: "123",
      createdBy: "John Doe",
      createdAt: "2021-09-27",
      updatedAt: "2021-09-27",
    },
    {
      title: "Math",
      documentId: "123",
      createdBy: "John Doe",
      createdAt: "2021-09-27",
      updatedAt: "2021-09-27",
    },
    {
      title: "Math",
      documentId: "123",
      createdBy: "John Doe",
      createdAt: "2021-09-27",
      updatedAt: "2021-09-27",
    },
    {
      title: "Math",
      documentId: "123",
      createdBy: "John Doe",
      createdAt: "2021-09-27",
      updatedAt: "2021-09-27",
    },
    {
      title: "Math",
      documentId: "123",
      createdBy: "John Doe",
      createdAt: "2021-09-27",
      updatedAt: "2021-09-27",
    },
    {
      title: "Math",
      documentId: "123",
      createdBy: "John Doe",
      createdAt: "2021-09-27",
      updatedAt: "2021-09-27",
    },{
      title: "Math",
      documentId: "123",
      createdBy: "John Doe",
      createdAt: "2021-09-27",
      updatedAt: "2021-09-27",
    },
    {
      title: "Math",
      documentId: "123",
      createdBy: "John Doe",
      createdAt: "2021-09-27",
      updatedAt: "2021-09-27",
    },
    {
      title: "Math",
      documentId: "123",
      createdBy: "John Doe",
      createdAt: "2021-09-27",
      updatedAt: "2021-09-27",
    },
  ];
  return (
    <div className="flex flex-col items-center relative z-0">
      <div className="sticky top-0  bg-base-100 w-full flex justify-center z-10">
      <ModuleHero
        title="Mathematics"
        description='Mathematics is the study of numbers, shapes and patterns. The word comes from the Greek word "μάθημα" (máthema), meaning "science, knowledge, or learning", and is sometimes shortened to maths (in England, Australia, Ireland, and New Zealand) or math (in the United States and Canada).Mathematics is the study of numbers, shapes and patterns. The word comes from the Greek word "μάθημα" (máthema), meaning "science, knowledge, or learning", and is sometimes shortened to maths (in England, Australia, Ireland, and New Zealand) or math (in the United States and Canada).Mathematics is the study of numbers, shapes and patterns. The word comes from the Greek word "μάθημα" (máthema), meaning "science, knowledge, or learning", and is sometimes shortened to maths (in England, Australia, Ireland, and New Zealand) or math (in the United States and Canada).Mathematics is the study of numbers, shapes and patterns. The word comes from the Greek word "μάθημα" (máthema), meaning "science, knowledge, or learning", and is sometimes shortened to maths (in England, Australia, Ireland, and New Zealand) or math (in the United States and Canada).Mathematics is the study of numbers, shapes and patterns. The word comes from the Greek word "μάθημα" (máthema), meaning "science, knowledge, or learning", and is sometimes shortened to maths (in England, Australia, Ireland, and New Zealand) or math (in the United States and Canada).Mathematics is the study of numbers, shapes and patterns. The word comes from the Greek word "μάθημα" (máthema), meaning "science, knowledge, or learning", and is sometimes shortened to maths (in England, Australia, Ireland, and New Zealand) or math (in the United States and Canada).Mathematics is the study of numbers, shapes and patterns. The word comes from the Greek word "μάθημα" (máthema), meaning "science, knowledge, or learning", and is sometimes shortened to maths (in England, Australia, Ireland, and New Zealand) or math (in the United States and Canada).Mathematics is the study of numbers, shapes and patterns. The word comes from the Greek word "μάθημα" (máthema), meaning "science, knowledge, or learning", and is sometimes shortened to maths (in England, Australia, Ireland, and New Zealand) or math (in the United States and Canada).'
        moduleId={moduleId}
        createdBy="John Doe"
        createdAt="2021-09-27"
        updatedAt="2021-09-27"
      />

</div>
      <div className="top-20 flex flex-row w-96 justify-end items-center my-1">

        {/* TODO: Add actions like
        1. test subject -> modal: test configuration -> test (test will be saved in personal history)
        2. speak to AI mentor -> modal: chat with AI mentor (AI mentor given subject context)

        add remove action to the modules
        add create module action below the modules
        */}
      <span className="w-96 text-start font-bold">Modules:</span>
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
        {modules.length === 0 && <span className="text-center">No modules found</span>}
         {modules.map((module, index) => (
          <DocumentCard
            key={index}
            title={module.title}
            documentId={module.documentId}
            createdBy={module.createdBy}
            createdAt={module.createdAt}
            updatedAt={module.updatedAt}
            memoryCount={0}
            questionCount={0}
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
