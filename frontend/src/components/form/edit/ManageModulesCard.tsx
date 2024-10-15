import React, { useEffect, useState } from "react";
import {
  FullModule,
  getModulesBySubjectId,
  GetModulesResponse,
  getPublicModules,
  setUserModuleFavourite,
} from "../../../api/moduleService";
import { unixToDateString } from "../../../utilities/timeUtilities";
import { OrderByDirection, OrderByField } from "../../../types/enums";
import Select from "../../select/Select";
import { sortByMap } from "../../../utilities/constants";
import ToggleCheck from "../../checkbox/ToggleCheck";
import { TiArrowSortedDown, TiArrowSortedUp } from "react-icons/ti";
import { setSubjectModuleMapping } from '../../../api/subjectService';

type Props = {
  subjectId: number;
  retriggerFunc?: () => void;
};
const ManageModulesCard = (props: Props) => {
  const [isAsc, setIsAsc] = React.useState(false);

  const [sortBy, setSortBy] = React.useState<OrderByField>(
    OrderByField.ORDER_BY_FIELD_ID
  );
  const [ownedModules, setOwnedModules] = useState<FullModule[]>([]);
  const [otherModules, setOtherModules] = useState<FullModule[]>([]);
  const hasFetched = React.useRef(false);

  const fetchOwnedModules = async () => {
    try{
      const response: GetModulesResponse = await getModulesBySubjectId({subject_id: Number(props.subjectId), page_number:1 , sort_by: sortBy, order: isAsc ? OrderByDirection.ORDER_BY_DIRECTION_ASC : OrderByDirection.ORDER_BY_DIRECTION_DESC});
      setOwnedModules(response.modules)
      await fetchModules(response.modules)
    }catch (e) {
        setOwnedModules([]);
      await fetchModules(ownedModules);
      if (e instanceof Error) {
        console.log(e);
      } else {
        console.log(e);
      }
    }
  }

  const fetchModules = async (curModules: FullModule[]) => {
    try {
      const order = isAsc
        ? OrderByDirection.ORDER_BY_DIRECTION_ASC
        : OrderByDirection.ORDER_BY_DIRECTION_DESC;
      const response: GetModulesResponse = await getPublicModules({
        page_number: 1,
        order: order,
        sort_by: sortBy,
      });
      hasFetched.current = true;
      // filter out modules that exist in ownedModules
      if (curModules && curModules.length > 0) {
        const otherModules = response.modules.filter(
            (otherModule) => !curModules.some((ownedModule) => ownedModule.module.module_id === otherModule.module.module_id)
        );

      setOtherModules(otherModules);
    }else{
        setOwnedModules([]);
        setOtherModules(response.modules);
    }
    } catch (e) {
      if (e instanceof Error) {
        console.log(e);
      } else {
        console.log(e);
      }
    }
  };

  useEffect(() => {
    if (!hasFetched.current) {
        fetchOwnedModules();
    }
  }, []);

  const onSaveChange = async () =>{

    try{
        await setSubjectModuleMapping({subject_id: props.subjectId, module_ids: ownedModules.map((module) => module.module.module_id)});
        if (props.retriggerFunc) {
          props.retriggerFunc();
        }
    }
    catch(e){
        console.log(e)
    }

  }

  const onRemoveClick = async (moduleId: number) => {
    // remove from owned modules
    // add to other modules
    const newOwnedModules = ownedModules.filter((module) => module.module.module_id !== moduleId);
    const newOtherModules = [...otherModules, ownedModules.find((module) => module.module.module_id === moduleId)] as FullModule[];
    setOwnedModules(newOwnedModules);
    setOtherModules(newOtherModules);
    }

  const onAddClick = async (moduleId: number) => {
    // remove from other modules
    // add to owned modules
    const newOwnedModules = [...ownedModules, otherModules.find((module) => module.module.module_id === moduleId)] as FullModule[];
    const newOtherModules = otherModules.filter((module) => module.module.module_id !== moduleId);
    setOwnedModules(newOwnedModules);
    setOtherModules(newOtherModules);
    }

  return (
    <div
      className="
        h-fit w-92
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
      <h1 className="truncate h-fit mb-1">Memory {""}</h1>
      <div className="flex flex-col h-fit">
        <div
          className="flex flex-col w-full"
          onSubmit={() => {
            console.log("submit");
          }}
        >
          <span>Owned:</span>
          <div className="w-full bg-gray-200 shadow-inset m-2 p-2 h-52 overflow-y-scroll w-92">
            {ownedModules &&
              ownedModules.map((module, index) => (
                <ModuleCard
                className="w-full"
                  key={index}
                  title={module.module.module_name}
                  moduleId={String(module.module.module_id)}
                  createdAt={unixToDateString(module.module.created_time)}
                  updatedAt={unixToDateString(module.module.updated_time)}
                  isFavourite={module.is_favourite}
                  
                >
                <button
                  className=""
                  type="submit"
                  onClick={e => {e.preventDefault(); onRemoveClick(module.module.module_id)}}
                >
                  Remove
                </button></ModuleCard>
              ))}
          </div>
        </div>
        <div className="divider divider-horizontal" />
        <div className="flex flex-col w-full">
            <div className="flex flex-row"><span>Others:</span>
          <Select
            name="Order By"
            items={Object.keys(sortByMap)}
            setSelected={(val: string) => {
              setSortBy(sortByMap[val]);
            }}
          />
          <ToggleCheck
            isChecked={isAsc}
            checkName="Asc"
            checkIcon={<TiArrowSortedUp />}
            unCheckName="Desc"
            unCheckIcon={<TiArrowSortedDown />}
            onClick={() => setIsAsc(!isAsc)}
          /></div>
          
          <div className="w-full bg-gray-200 shadow-inset m-2 p-2 h-52 overflow-y-scroll">
            {otherModules &&
              otherModules.map((module, index) => (
                <ModuleCard
                className="w-full"
                  key={index}
                  title={module.module.module_name}
                  moduleId={String(module.module.module_id)}
                  createdAt={unixToDateString(module.module.created_time)}
                  updatedAt={unixToDateString(module.module.updated_time)}
                  isFavourite={module.is_favourite}
                >

          <button
            className=""
            type="submit"
            onClick={e => {e.preventDefault(); onAddClick(module.module.module_id)}}
          >
            Add
          </button>
                </ModuleCard>
              ))}
          </div>
        </div>
      </div>
      <button
        className="mb-1"
        // log the selected option
        onClick={(e) => {
          e.preventDefault();
          console.log("Update");
          onSaveChange();
        }}
      >
        Save Changes
      </button>
      <button className="mb-1" type="submit">
        Cancel
      </button>
    </div>
  );
};

export default ManageModulesCard;



type cardProps = {
    title: string;
    moduleId: string;
    createdBy?: string;
    createdAt: string;
    updatedAt: string;
    isFavourite: boolean;
    className?: string;
    children?: React.ReactNode;
  };
  
  const ModuleCard = (props: cardProps) => {
    const [isFavourite, setIsFavourite] = React.useState(props.isFavourite);
  
    const onFavouriteClick = async () => {
      try {
        console.log(isFavourite)
        const response = await setUserModuleFavourite({ module_id: Number(props.moduleId), is_favourite: !isFavourite });
        console.log(response)
        setIsFavourite(!isFavourite);
      } catch (e) {
        console.log(e);
      }
    }

    // const onAddClick = async () => {
    // }
    return (
      <div
        className=
        {
          props.className
          + " "
          +
          `
        h-fit w-96
        max-w-full
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
        animate-appear-bot
        `
        }
      >
        <div className="flex flex-row justify-between align-middle content-center text-center align-center  ">
          
          <h1 className="truncate h-fit mb-1">{props.title}</h1>
          <span className="text-gray-400 font-bold flex  items-center">
            #{props.moduleId.padStart(8, "0")}
          </span>
        </div>
        <div className="flex flex-row justify-between align-middle content-center text-center align-center ">
          <div className="w-full flex flex-row text-left text-xs justify-between select-none text-gray-500">
            {props.createdBy && <div className="flex flex-col justify-between  w-full">
              <span>Owned By: </span>
              <span>{props.createdBy}</span>
            </div>}
            <div className="flex flex-col justify-between  w-full">
              <span>Created At: </span>
              <span>{props.createdAt}</span>
            </div>
            <div className="flex flex-col justify-between  w-full">
              <span>Last Update: </span>
              <span>{props.updatedAt}</span>
            </div>
          </div>
          {
            props.children
          }
          <Heart
            isFavourite={isFavourite}
            onClick={() => {
              onFavouriteClick();
            }}/>
        </div>
      </div>
    );
  };
  
  const Heart = (props: { isFavourite: boolean; onClick: () => void }) => {
    return (
      <svg
        onClick={props.onClick}
        xmlns="http://www.w3.org/2000/svg"
        className={
          `h-7 w-7 cursor-pointer mask mask-heart flex-none ml-1 ` +
          `${props.isFavourite
            ? "scale-90 bg-red-500 "
            : `transition duration-50 hover:scale-90 bg-base-100 text-gray-500  flex justify-center items-center content-center group hover:bg-red-400`}`
        }
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          className={
            props.isFavourite
              ? "stroke-transparent"
              : `stroke-gray-500 group-hover:stroke-transparent`
          }
          strokeLinecap="round"
          strokeLinejoin="round"
          strokeWidth="2"
          d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"
        />
      </svg>
    );
  };
  