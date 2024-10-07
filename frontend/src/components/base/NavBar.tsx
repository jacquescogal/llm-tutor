import { useDispatch } from "react-redux";
import { toggleDrawer } from "../../store/drawerOpen";
import { useSelector } from "react-redux";
import { RootState } from "../../store/store";

const NavBar = () => {
  const dispatch = useDispatch();
  const breadCrumbList = useSelector((state: RootState) => state.breadCrumbList);
  return (
    <div className="navbar bg-base-100 shadow h-16">
      <div className="flex-none">
        <button
          className="btn btn-square btn-ghost"
          onClick={() => {
            dispatch(toggleDrawer());
          }}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            className="inline-block h-5 w-5 stroke-current"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="M4 6h16M4 12h16M4 18h16"
            ></path>
          </svg>
        </button>
      </div>
      <div className="flex-1">
        <div className="breadcrumbs text-sm">
        <ul>

        <li><a className="text-base-content font-boldtext-xl text-xl pr-0">Tutor-ai</a></li>
          {
            breadCrumbList.value.map((crumb, index) => (
              <li key={index}>
                <a href={crumb.path}>{crumb.name}</a>
              </li>
            ))
          }
        </ul>
      </div>
      </div>
      
      <div className="flex-none gap-2">
        
      </div>
    </div>
  );
};

export default NavBar;
