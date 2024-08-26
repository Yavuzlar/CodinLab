import React, { createContext, useContext, useState } from "react";

const defaultProvider = {
    anchorElNav: null,
    handleOpenNavMenu: null,
    handleCloseNavMenu: null,
  };
  

const NavContext = createContext( defaultProvider);

const useNav = () => useContext(NavContext);

const NavProvider = ({ children }) => {
  const [anchorElNav, setAnchorElNav] = useState(null);

  const handleOpenNavMenu = (event) => {
    setAnchorElNav(event.currentTarget);
  };

  const handleCloseNavMenu = () => {
    setAnchorElNav(null);
  };

  React.useEffect(() => {
    const handleResize = () => {
      setAnchorElNav(null);
    };
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, [setAnchorElNav]);

  const value = {
    anchorElNav: anchorElNav,
  OpenNavMenu: handleOpenNavMenu, 
  CloseNavMenu: handleCloseNavMenu,
  };

  return <NavContext.Provider value={value}>{children}</NavContext.Provider>;
};

export { NavContext, NavProvider };
