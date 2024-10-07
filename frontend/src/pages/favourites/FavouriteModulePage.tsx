import React from "react";
import Select from "../../components/select/Select";
import ToggleCheck from "../../components/checkbox/ToggleCheck";
import { TiArrowSortedDown, TiArrowSortedUp } from "react-icons/ti";
import ModuleCard from "../../components/cards/ModuleCard";
import { CrumbHelperSetRootLink } from "../../store/helpers/crumbHelper";

const FavouriteModulePage = () => {
  const [isAsc, setIsAsc] = React.useState(false);

  CrumbHelperSetRootLink({name: "Favourite Modules"});
  const modules = [
    {
      title: "Math",
      subjectId: "123",
      createdBy: "John Doe",
      createdAt: "2021-09-27",
      updatedAt: "2021-09-27",
    },
    {
      title: "Math",
      subjectId: "123",
      createdBy: "John Doe",
      createdAt: "2021-09-27",
      updatedAt: "2021-09-27",
    },
    {
      title: "Math",
      subjectId: "123",
      createdBy: "John Doe",
      createdAt: "2021-09-27",
      updatedAt: "2021-09-27",
    },
  ];
  return (
    <div className="flex flex-col justify-center content-center items-center h-fit">
      <h1>Favourite Modules</h1>
      <Search />
      <div className="flex flex-row w-96 justify-end items-center">
        <Select name="Order By" items={["Title", "Date"]} />
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
        {modules.map((module, index) => (
          <ModuleCard
            key={index}
            title={module.title}
            moduleId={module.subjectId}
            createdBy={module.createdBy}
            createdAt={module.createdAt}
            updatedAt={module.updatedAt}
          />
        ))}
      </div>
    </div>
  );
};

export default FavouriteModulePage;

const Search = () => {
  return (
    <div className="form-control flex flex-row my-2 w-96 justify-center">
      <input
        type="text"
        placeholder="Type here..."
        className="p-1 m-1 outline outline-slate-300 focus:outline-slate-500 rounded h-8"
      />
      <button className="p-2 m-1 h-8">Search</button>
    </div>
  );
};
