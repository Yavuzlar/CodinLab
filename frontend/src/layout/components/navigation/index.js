import List from "@mui/material/List";
import navigation from "src/navigation";
import NavItem from "./item/NavItem";
import { Stack } from "@mui/material";
import LanguageSelector from "./item/LanguageSelector";
import { useState } from "react";

const NavigationList = () => {
  const [anchorElNav, setAnchorElNav] = useState(null);
  const handleOpenNavMenu = (event) => {
    setAnchorElNav(event.currentTarget);
  };
  
  return (
    <Stack
      direction={"row"}
      spacing={10}
      sx={{ display: { xs: "none", mdlg: "block" } }}>
      {navigation.map((item, index) => (
        <NavItem key={index} {...item} />
      ))}
      <LanguageSelector />
    </Stack>
  );
};

export default NavigationList;
