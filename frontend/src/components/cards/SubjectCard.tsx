import React from "react";
import { useNavigate } from "react-router-dom";
import { setUserSubjectFavourite } from "../../api/subjectService";

type Props = {
  title: string;
  subjectId: string;
  createdBy?: string;
  createdAt: string;
  updatedAt: string;
  isFavourite: boolean;
};

const SubjectCard = (props: Props) => {
  const [isFavourite, setIsFavourite] = React.useState(props.isFavourite);
  const nav = useNavigate();

  const onFavouriteClick = async () => {
    try {
      console.log(isFavourite)
      const response = await setUserSubjectFavourite({ subject_id: Number(props.subjectId), is_favourite: !isFavourite });
      console.log(response)
      setIsFavourite(!isFavourite);
    } catch (e) {
      console.log(e);
    }
  }
  return (
    <div
      className="
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
      "
    >
      <div className="flex flex-row justify-between align-middle content-center text-center align-center  ">
        
        <h1 className="truncate h-fit mb-1">{props.title}</h1>
        <span className="text-gray-400 font-bold flex  items-center">
          #{props.subjectId.padStart(8, "0")}
        </span>
      </div>
      <div className="flex flex-row justify-between align-middle content-center text-center align-center ">
        <div className="w-full flex flex-row text-left text-xs justify-between select-none text-gray-500">
          
          {props.createdBy && <div className="flex flex-col justify-between  w-full">
            <span>Owned By: </span>
            <span>{props.createdBy}</span>
          </div>
}
          <div className="flex flex-col justify-between  w-full">
            <span>Created At: </span>
            <span>{props.createdAt}</span>
          </div>
          <div className="flex flex-col justify-between  w-full">
            <span>Last Update: </span>
            <span>{props.updatedAt}</span>
          </div>
        </div>
        <button
          className=""
          type="submit"
          // log the selected option
          // TODO: send this to proper page
          onClick={() => nav(props.subjectId)}
        >
          Next
        </button>
        <Heart
          isFavourite={isFavourite}
          onClick={() => {
            onFavouriteClick();
          }}/>
      </div>
    </div>
  );
};

export default SubjectCard;

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
