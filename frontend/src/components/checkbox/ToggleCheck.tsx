import React from "react";
import { IconContext } from "react-icons";

type Props = {
  isChecked: boolean;
  checkName: string;
  checkIcon: React.ReactNode;
  unCheckName: string;
  unCheckIcon: React.ReactNode;
    onClick: () => void;
};

const ToggleCheck = (props: Props) => {
  return (
    <div className="flex flex-row items-center rounded-xl w-24 overflow-hidden h-fit px-1 mx-1 text-sm bg-base-100 border-2 border-gray-300 cursor-pointer select-none hover:bg-gray-100"
    onClick={props.onClick}>
        <IconContext.Provider value={{ className: "flex-none text-gray-400" }}>
      {props.isChecked ? props.checkIcon : props.unCheckIcon}
        </IconContext.Provider>
      <span className="font-bold">{props.isChecked ? props.checkName : props.unCheckName}</span>
    </div>
  );
};

export default ToggleCheck;
