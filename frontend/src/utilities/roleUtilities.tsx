import { ReactNode } from "react";
import { UserModuleRole } from "../types/enums";

export const userModuleRoleToSpan = (role: UserModuleRole): ReactNode => {
    switch (role) {
        case UserModuleRole.USER_MODULE_ROLE_OWNER:
            return (
                <span className="text-red-500 font-bold">Owner</span>
            );
        case UserModuleRole.USER_MODULE_ROLE_ADMIN:
            return (
                <span className="text-blue-500 font-bold">Admin</span>
            );
        case UserModuleRole.USER_MODULE_ROLE_EDITOR:
            return (
                <span className="text-green-500 font-bold">Editor</span>
            );
        case UserModuleRole.USER_MODULE_ROLE_VIEWER:
            return (
                <span className="text-gray-500 font-bold">Viewer</span>
            );
        default:
            return (
                <span className="text-black font-bold">Public</span>
            );
    }
}
