import React from "react";

type SearchProps = {
  search: string;
  setSearch: (value: string) => void;
  startSearch: () => void;
};

const SearchBar = (props: SearchProps) => {
  return (
    <div className="form-control flex flex-row my-2 w-96 justify-center">
      <input
        type="text"
        placeholder="Type here..."
        className="p-1 m-1 outline outline-slate-300 focus:outline-slate-500 rounded h-8"
        value={props.search}
        onChange={(e) => props.setSearch(e.target.value)}
      />
      <button className="p-2 m-1 h-8" onClick={props.startSearch}>
        Search
      </button>
    </div>
  );
};

export default SearchBar;
