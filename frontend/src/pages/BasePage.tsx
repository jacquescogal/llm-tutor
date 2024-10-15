import { Outlet } from "react-router-dom";
import NavBar from "../components/base/NavBar";
import Drawer from "../components/base/Drawer";

const BasePage = () => {
  return (
    <div className="min-h-screen min-w-screen">
        <NavBar/>
        {/* Drawer */}
      <Drawer>
        <Outlet/>
      </Drawer>
    </div>
  );
};

export default BasePage;
