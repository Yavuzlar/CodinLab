import { AppBar, Toolbar } from "@mui/material";
import NavigationList from "./navigation";
import Title from "./navigation/item/Title";
import LanguageSelector from "./navigation/item/LanguageSelector";

const Navbar = () => {
  return (
    <AppBar
      sx={{
        backgroundColor: "#0A3B7A",
        position: "static",
        height: 52,
        width: "100%",
        display: "flex",
        justifyContent: "center",
        px: 12,
      }}>
      <Toolbar>
        <Title />
        <NavigationList />
      </Toolbar>
    </AppBar>
  );
};

export default Navbar;
