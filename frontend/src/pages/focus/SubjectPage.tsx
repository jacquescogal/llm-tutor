import React, { useEffect } from "react";
import SubjectHero from "../../components/hero/SubjectHero";
import Select from "../../components/select/Select";
import ToggleCheck from "../../components/checkbox/ToggleCheck";
import { TiArrowSortedDown, TiArrowSortedUp } from "react-icons/ti";
import Dropdown from "../../components/dropdown/Dropdown";
import { IoMdMenu } from "react-icons/io";
import { CrumbHelperPushLink} from "../../store/helpers/crumbHelper";
import { useLocation } from "react-router-dom";
import { FullSubject, getSubjectById, getSubjectByIdResponse } from "../../api/subjectService";
import { FullModule, getModulesBySubjectId, GetModulesResponse } from "../../api/moduleService";
import { unixToDateString } from '../../utilities/timeUtilities';
import { sortByMap } from "../../utilities/constants";
import { OrderByDirection, OrderByField } from "../../types/enums";
import ModuleCard from "../../components/cards/ModuleCard";

const SubjectPage = () => {
  const [subjectHeroDetails, setSubjectHeroDetails] = React.useState<FullSubject | null>(null);
  const [modules, setModules] = React.useState<FullModule[]>([]);
  const [isAsc, setIsAsc] = React.useState(false);
  const [sortBy, setSortBy] = React.useState<OrderByField>(OrderByField.ORDER_BY_FIELD_ID);
  const location = useLocation();
  const locationPath = location.pathname.split('/');
  while (locationPath.length>0 && locationPath[locationPath.length - 1] === "") {
    locationPath.pop();
  }
  const subjectId = locationPath.pop() || "";
  CrumbHelperPushLink({ name: `Subject: #${subjectId.padStart(5,'0')}` });
  useEffect(()=>{
    const fetchSubject = async () =>{
      try {
        const response: getSubjectByIdResponse = await getSubjectById({subject_id: Number(subjectId)});
        console.log(response);
        setSubjectHeroDetails(response.subject);
      } catch (e) {
        if (e instanceof Error) {
          console.log(e);
        } else {
          console.log(e);
        }
      }
    }

    fetchSubject();
  },[]);

  const fetchModules = async () => {
    try{
      console.log(subjectId)
      const response: GetModulesResponse = await getModulesBySubjectId({subject_id: Number(subjectId), page_number:1 , sort_by: sortBy, order: isAsc ? OrderByDirection.ORDER_BY_DIRECTION_ASC : OrderByDirection.ORDER_BY_DIRECTION_DESC});
      setModules(response.modules);
    }catch (e) {
      if (e instanceof Error) {
        console.log(e);
      } else {
        console.log(e);
      }
    }
  }

  useEffect(()=>{
    fetchModules();
  },[isAsc, sortBy]);


  return (
    <div className="flex flex-col items-center relative z-0">
      <div className="sticky top-0  bg-base-100 w-full flex justify-center z-10">
      <SubjectHero
        title={subjectHeroDetails?.subject.subject_name || "Loading..."}
        description={subjectHeroDetails?.subject.subject_description || "Loading..."}
        subjectId={subjectId}
        createdAt={subjectHeroDetails && unixToDateString(subjectHeroDetails?.subject.created_time) || "Loading..."}
        updatedAt= {subjectHeroDetails && unixToDateString(subjectHeroDetails?.subject.updated_time) || "Loading..."}
        isFavourite={subjectHeroDetails?.is_favourite || false}
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
        <Select name="Order By" items={Object.keys(sortByMap)} setSelected={(s)=>{setSortBy(sortByMap[s])}}/>
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
        {!modules || modules.length === 0 && <span className="text-center">No modules found</span>}
         {modules && modules.map((module, index) => (
          <ModuleCard
          key={index}
          title={module.module.module_name}
          moduleId={String(module.module.module_id)}
          createdAt={unixToDateString(module.module.created_time)}
          updatedAt={unixToDateString(module.module.updated_time)}
          isFavourite={module.is_favourite}
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

export default SubjectPage;
