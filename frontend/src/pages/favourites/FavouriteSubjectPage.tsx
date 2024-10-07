import React from "react";
import SubjectCard from "../../components/cards/SubjectCard";
import Select from "../../components/select/Select";
import ToggleCheck from "../../components/checkbox/ToggleCheck";
import { TiArrowSortedDown, TiArrowSortedUp } from "react-icons/ti";
import { CrumbHelperSetRootLink } from "../../store/helpers/crumbHelper";



const FavouriteSubjectPage = () => {
  const [isAsc, setIsAsc] = React.useState(false);
  CrumbHelperSetRootLink({name: "Favourite Subjects"});
  const subjects = [
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
      <h1>Favourite Subjects</h1>
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
        {subjects.map((subject, index) => (
          <SubjectCard
            key={index}
            title={subject.title}
            subjectId={subject.subjectId}
            createdBy={subject.createdBy}
            createdAt={subject.createdAt}
            updatedAt={subject.updatedAt}
          />
        ))}
        {subjects.map((subject, index) => (
          <SubjectCard
            key={index}
            title={subject.title}
            subjectId={subject.subjectId}
            createdBy={subject.createdBy}
            createdAt={subject.createdAt}
            updatedAt={subject.updatedAt}
          />
        ))}
      </div>
    </div>
  );
};

export default FavouriteSubjectPage;

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
