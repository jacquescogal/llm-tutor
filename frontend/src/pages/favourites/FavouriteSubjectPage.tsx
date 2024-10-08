import React, { useEffect } from "react";
import SubjectCard from "../../components/cards/SubjectCard";
import Select from "../../components/select/Select";
import ToggleCheck from "../../components/checkbox/ToggleCheck";
import { TiArrowSortedDown, TiArrowSortedUp } from "react-icons/ti";
import { CrumbHelperSetRootLink } from "../../store/helpers/crumbHelper";
import { FullSubject, getFavouriteSubjectsByUserId, getFavouriteSubjectsByUserIdResponse } from '../../api/subjectService';
import { unixToDateString } from "../../utilities/timeUtilities";
import { OrderByDirection, OrderByField } from "../../types/enums";
import { sortByMap } from "../../utilities/constants";



const FavouriteSubjectPage = () => {
  const [isAsc, setIsAsc] = React.useState(false);
  CrumbHelperSetRootLink({name: "Favourite Subjects"});
  const [subjects, setSubjects] = React.useState<FullSubject[]>([]);
  const [sortBy, setSortBy] = React.useState<OrderByField>(OrderByField.ORDER_BY_FIELD_ID);
  const fetchSubjects = async () => {
    try{
      const order = isAsc ? OrderByDirection.ORDER_BY_DIRECTION_ASC : OrderByDirection.ORDER_BY_DIRECTION_DESC;
      const response: getFavouriteSubjectsByUserIdResponse = await getFavouriteSubjectsByUserId({page_number: 1, order:order, sort_by:sortBy});
      setSubjects(response.subjects);
    }catch (e) {
    if (e instanceof Error) {
      // setError(e.message);
    } else {
      // setError("An unknown error occurred");
    }
  }
  }


  useEffect(()=>{
      fetchSubjects();
  },[isAsc, sortBy]);

  return (
    <div className="flex flex-col justify-center content-center items-center h-fit">
      <h1>Favourite Subjects</h1>
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
        {subjects && subjects.map((s, index) => (
          <SubjectCard
            key={index}
            title={s.subject.subject_name}
            subjectId={String(s.subject.subject_id)}
            createdAt={unixToDateString(s.subject.created_time)}
            updatedAt={unixToDateString(s.subject.updated_time)}
            isFavourite={s.is_favourite}
          />
        ))}
      </div>
    </div>
  );
};

export default FavouriteSubjectPage;

