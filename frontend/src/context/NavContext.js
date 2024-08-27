import { useRouter } from "next/router";
import React, { createContext, useContext, useState } from "react";

const defaultProvider = {
  anchorElNav: null,
  handleOpenNavMenu: null,
  handleCloseNavMenu: null,
};

const NavContext = createContext(defaultProvider);

const NavProvider = ({ children }) => {
  const [anchorElNav, setAnchorElNav] = useState(null);
  const router = useRouter();

  const handleOpenNavMenu = (event) => {
    setAnchorElNav(event.currentTarget);
  };

  const handleCloseNavMenu = () => {
    setAnchorElNav(null);
  };
  const handleLogoClick = () => {
    router.push("/");
  };

  const handleChangePage = () => {
    setAnchorElNav(null);
    console.log("page changed");
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
    LogoClick: handleLogoClick,
    ChangePage: handleChangePage,
  };

  return <NavContext.Provider value={value}>{children}</NavContext.Provider>;
};

export { NavContext, NavProvider };
