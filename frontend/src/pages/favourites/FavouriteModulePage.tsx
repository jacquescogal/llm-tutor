import React, { useEffect } from "react";
import Select from "../../components/select/Select";
import ToggleCheck from "../../components/checkbox/ToggleCheck";
import { TiArrowSortedDown, TiArrowSortedUp } from "react-icons/ti";
import ModuleCard from "../../components/cards/ModuleCard";
import { CrumbHelperSetRootLink } from "../../store/helpers/crumbHelper";
import { GetModulesResponse, FullModule, getFavouriteModulesByUserId } from '../../api/moduleService';
import { sortByMap } from "../../utilities/constants";
import { OrderByDirection, OrderByField } from "../../types/enums";
import { unixToDateString } from "../../utilities/timeUtilities";

const FavouriteModulePage = () => {
  const [isAsc, setIsAsc] = React.useState(false);
  const [sortBy, setSortBy] = React.useState<OrderByField>(OrderByField.ORDER_BY_FIELD_ID);
  CrumbHelperSetRootLink({name: "Favourite Modules"});
  const [modules, setModules] = React.useState<FullModule[]>([]);
  

  const fetchModules = async () => {
    try{
      const order = isAsc ? OrderByDirection.ORDER_BY_DIRECTION_ASC : OrderByDirection.ORDER_BY_DIRECTION_DESC;
      const response: GetModulesResponse = await getFavouriteModulesByUserId({page_number: 1, order:order, sort_by:sortBy});
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
    <div className="flex flex-col justify-center content-center items-center h-fit">
      <h1>Favourite Modules</h1>
      <div className="flex flex-row w-96 justify-end items-center">
        <Select name="Order By" items={Object.keys(sortByMap)} setSelected={(val:string)=>{setSortBy(sortByMap[val]);}}/>
        <ToggleCheck
          isChecked={isAsc}
          checkName="Asc"
          checkIcon={<TiArrowSortedUp/>}
          unCheckName="Desc"
          unCheckIcon={<TiArrowSortedDown/>}
          onClick={() => setIsAsc(!isAsc)}
        />
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2 w-fit m-4">
        {modules && modules.map((m, index) => (
          <ModuleCard
            key={index}
            title={m.module.module_name}
            moduleId={String(m.module.module_id)}
            // createdBy={module.crea}
            createdAt={unixToDateString(m.module.created_time)}
            updatedAt={unixToDateString(m.module.updated_time)}
            isFavourite={m.is_favourite}
          />
        ))}
      </div>
    </div>
  );
};

export default FavouriteModulePage;
