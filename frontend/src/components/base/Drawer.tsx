import { RootState } from "../../store/store";
import { useDispatch } from "react-redux";
import { useSelector } from "react-redux";
import { toggleDrawer } from "../../store/drawerOpen";
import { useNavigate } from "react-router-dom";
import { FaRegHeart } from "react-icons/fa";
import { IoSearch } from "react-icons/io5";
import { GrDocumentTest, GrUpload } from "react-icons/gr";
import ModalButton from "../modal/ModalButton";
import { RiLogoutBoxLine } from "react-icons/ri";
import { deleteSession } from "../../api/authService";
import CreateSubjectCard from "../form/create/CreateSubjectCard";
import CreateModuleCard from "../form/create/CreateModuleCard";

type Props = {
  children?: React.ReactNode;
};

// type SubDestinationList = {

type Destination = {
  name: React.ReactNode;
  url?: string;
  child?: Destination[];
};

const Drawer = (props: Props) => {
  const drawerOpen = useSelector((state: RootState) => state.drawerOpen.value);
  const dispatch = useDispatch();
  const nav = useNavigate();
  const logout = async () => {
    deleteSession()
      .then(() => nav("/login"))
      .catch((e) => console.log(e));
  };
  const MenuOption = (destination: Destination) => {
    return (
      <>
        {destination.child ? (
          <Parent {...destination} />
        ) : (
          <Leaf {...destination} />
        )}
      </>
    );
  };

  const Parent = ({ name, child }: Destination) => {
    return (
      <li>
        <details open>
          <summary>{name}</summary>
          <ul>
            {child?.map((destination, index) => (
              <>
                {destination.child ? (
                  <Parent key={index} {...destination} />
                ) : (
                  <Leaf key={index} {...destination} />
                )}
              </>
            ))}
          </ul>
        </details>
      </li>
    );
  };

  const Leaf = ({ name, url }: Destination) => {
    const urlString = url ? url : "";
    return (
      <li>
        <a onClick={() => nav(urlString)}>{name}</a>
      </li>
    );
  };
  // explore > subject | modules
  // favourites > subject | modules
  // owned > subject | modules
  // privileged > subject | modules

  const destinations: Destination[] = [
    {
      name: (
        <>
          <IoSearch />
          Explore
        </>
      ),
      child: [
        {
          name: "Subject",
          url: "explore/subject",
        },
        {
          name: "Module",
          url: "explore/module",
        },
      ],
    },
    {
      name: (
        <>
          <FaRegHeart />
          Favourites
        </>
      ),
      child: [
        {
          name: "Subject",
          url: "favourite/subject",
        },
        {
          name: "Module",
          url: "favourite/module",
        },
      ],
    },
    {
      name: (
        <>
          <GrDocumentTest />
          Tests
        </>
      ),
      child: [
        {
          name: "New",
          url: "favourite/module",
        },
        {
          name: "In Progress",
          url: "favourite/module",
        },
        {
          name: "Results",
          url: "favourite/module",
        },
      ],
    },
    {
      name: (
        <>
          <GrUpload />
          Created / Uploaded
        </>
      ),
      child: [
        {
          name: "Subject",
          url: "favourite/module",
        },
        {
          name: "Modules",
          url: "favourite/module",
        },
        {
          name: "Document",
          url: "favourite/subject",
        },
      ],
    },
  ];

  return (
    <div className="drawer">
      <input
        id="my-drawer"
        type="checkbox"
        className="drawer-toggle"
        checked={drawerOpen}
      />
      <div className="drawer-content overflow-y-scroll h-less-nav">
        {props.children}
      </div>
      <div className="drawer-side">
        <label
          htmlFor="my-drawer"
          className="drawer-overlay"
          onClick={() => dispatch(toggleDrawer())}
        ></label>
        <ul className="menu bg-base-200 text-base-content min-h-full w-64 p-4">
          {destinations.map((destination, index) => (
            <MenuOption key={index} {...destination} />
          ))}
          <div className="bg-gray-300 shadow shadow-inset h-fit w-full p-2">
          <span>Create</span>
          <div className="flex flex-col">
            <ModalButton className="mb-1" buttonName="Subject"><CreateSubjectCard/></ModalButton>
            <ModalButton buttonName="Modules"><CreateModuleCard/></ModalButton>
          </div>
          </div>
          <div className="divider" />
          <button onClick={() => logout()}>
            <div className="flex flex-row h-fit justify-center">
              {<RiLogoutBoxLine />}Logout
            </div>
          </button>
        </ul>
      </div>
    </div>
  );
};

export default Drawer;
