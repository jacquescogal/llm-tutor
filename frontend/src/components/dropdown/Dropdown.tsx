import React, { ReactNode } from "react";

type Props = {
  dropSymbol: React.ReactNode;
  dropName: string;
  items?: ReactNode[];
};

const Dropdown = (props: Props) => {
  return (
    <div className="dropdown dropdown-top dropdown-hover">
      <div
        tabIndex={0}
        role="button"
      >
        <button className="flex flex-row px-1">
          {props.dropName}{props.dropSymbol}
        </button>
      </div>
      <ul
        tabIndex={0}
        className="dropdown-content menu bg-base-100 rounded-box z-[1] w-52 p-2 shadow"
      >
        {props.items?.map((item, index) => (
          <li key={index}>
            <a>{item}</a>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Dropdown;
