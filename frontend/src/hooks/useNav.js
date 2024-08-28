import { useContext } from "react";
import { NavContext } from "src/context/NavContext";

export const useNav = () => useContext(NavContext);
