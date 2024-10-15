import React from "react";
import {
  createSession,
  createUser,
} from "../../api/authService";
import { useNavigate } from "react-router-dom";

const LoginCard = () => {
  const [username, setUsername] = React.useState("admin");
  const [password, setPassword] = React.useState("admin");
  const [error, setError] = React.useState("");
  const nav = useNavigate();

  const onLoginClick = async () => {
    try {
      await createSession({ username, password });

      // Store username in local storage for 
      localStorage.setItem('username', username);
      nav("/dashboard/explore/subject");
    } catch (e) {
      if (e instanceof Error) {
        setError(e.message);
      } else {
        setError("An unknown error occurred");
      }
    }
  };

  const onRegisterClick = async () => {
    try {
      await createUser({ username, password });
      await onLoginClick();
    } catch (e) {
      if (e instanceof Error) {
        setError(e.message);
      } else {
        setError("An unknown error occurred");
      }
    }
  };

  return (
    <div
      className="
    h-fit w-96 
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
      <h1 className="truncate h-fit">Login</h1>
      <form
        className="flex flex-col"
        onSubmit={(e) => {
          e.preventDefault();
        }}
      >
        <label>Username</label>
        <input
          className="text-wrap break-words h-8 overflow-scroll bg-slate-50 shadow-inner p-1"
          onChange={(e) => setUsername(e.target.value)}
          value={username}
        />
        <label>Password</label>
        <input
          className="text-wrap break-words h-8 overflow-scroll bg-slate-50 shadow-inner p-1 mb-1"
          onChange={(e) => setPassword(e.target.value)}
          value={password}
          type="password"
        />
        {error && (
          <span className="normal-case rounded-badge bg-error text-error-content truncate p-1">
            {error}
          </span>
        )}
        <div className="p-2 flex flex-row h-16 w-fit justify-end w-full">
          <button
            className="mr-1"
            type="submit"
            onClick={(e) => {
              e.preventDefault();
              onLoginClick();
            }}
          >
            Login
          </button>
          <button
          onClick={(e) => {
            e.preventDefault();
            onRegisterClick();
          }}
          type="submit">Register</button>
        </div>
      </form>
    </div>
  );
};

export default LoginCard;
