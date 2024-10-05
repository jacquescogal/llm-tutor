import React from "react";

type Props = {
  username: string;
  password: string;
  setUsername: React.Dispatch<React.SetStateAction<string>>;
  setPassword: React.Dispatch<React.SetStateAction<string>>;
  error: string;
  setError: React.Dispatch<React.SetStateAction<string>>;
};

const LoginCard = (props: Props) => {
  return (
    <div
      className="
    h-64 w-96 
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
    "
    >
      <h1 className="truncate h-10">Login</h1>
      <form
        className="flex flex-col"
        onSubmit={(e) => {
          e.preventDefault();
          console.log("submit");
        }}
      >
        <label>Username</label>
        <input
          className="text-wrap break-words h-8 overflow-scroll bg-slate-50 shadow-inner p-1"
          onChange={(e) => props.setUsername(e.target.value)}
          value={props.username}
        />
        <label>Password</label>
        <input
          className="text-wrap break-words h-8 overflow-scroll bg-slate-50 shadow-inner p-1 mb-1"
          onChange={(e) => props.setPassword(e.target.value)}
          value={props.password}
          
        />
        {props.error && <span className="normal-case rounded-badge bg-error text-error-content truncate p-1">{props.error}</span>}
        <div className="absolute bottom-0 right-0 p-2 flex flex-row h-16 w-fit">
          <button className="mr-1" type="submit"
          onClick={()=>{props.setError(e=>e+'lol')}}>
            Login
          </button>
          <button type="submit">Register</button>
        </div>
      </form>
    </div>
  );
};

export default LoginCard;
