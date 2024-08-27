import * as React from "react";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";
import Menu from "@mui/material/Menu";
import MenuIcon from "@mui/icons-material/Menu";
import SettingsIcon from "@mui/icons-material/Settings";
import Container from "@mui/material/Container";
import CircleIcon from "@mui/icons-material/Circle";
import LogoutIcon from "@mui/icons-material/Logout";
import Button from "@mui/material/Button";
import themeConfig from "src/configs/themeConfig";
import navigation from "src/navigation";
import NavItem from "./navigation/item/NavItem";
import LanguageSelector from "./navigation/item/LanguageSelector";
import { useAuth } from "src/hooks/useAuth";
import { useEffect } from "react";

function ResponsiveAppBar() {
  const { logout } = useAuth();
  const [anchorElNav, setAnchorElNav] = React.useState(null);
  const [anchorElSettings, setAnchorElSettings] = React.useState(null);

  const handleOpenNavMenu = (event) => {
    setAnchorElNav(event.currentTarget);
  };

  const handleCloseNavMenu = () => {
    setAnchorElNav(null);
  };

  const handleOpenSettingsMenu = (event) => {
    setAnchorElSettings(event.currentTarget);
  };

  const handleCloseSettingsMenu = () => {
    setAnchorElSettings(null);
  };

  useEffect(() => {
    const handleResize = () => {
      setAnchorElNav(null);
      setAnchorElSettings(null);
    };
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);

  }, [ 
    setAnchorElNav,
    setAnchorElSettings,
  ]);

  const handleLogout = async () => {
    try {
      await logout();
    } catch (error) {
      console.log(error);
    }
  };

  const SettingsMenu = (
    <Menu
      anchorEl={anchorElSettings}
      anchorOrigin={{
        vertical: "bottom",
        horizontal: "right",
      }}
      keepMounted
      transformOrigin={{
        vertical: "top",
        horizontal: "right",
      }}
      open={Boolean(anchorElSettings)}
      onClose={handleCloseSettingsMenu}
      sx={{ "& .MuiMenu-paper": { backgroundColor: "#0A3B7A" } }}
    >
      <Box sx={{ display: "flex", flexDirection: "column" }}>
        <IconButton onClick={handleLogout} color="inherit"
       sx={{ color: "inherit" , borderRadius: "0" }}
         >
          <LogoutIcon />
        </IconButton>
        <LanguageSelector isMenu={true} />
      </Box>
    </Menu>
  );

  return (
    <AppBar
      sx={{ backgroundColor: "#0A3B7A", boxShadow: "none" }}
    >
      <Container maxWidth="lgPlus">
        <Toolbar disableGutters variant="dense">
          <CircleIcon
            sx={{
              display: { xs: "none", mdlg: "flex" },
              height: 40,
              width: 40,
              mr: 1,
            }}
          />
          <Typography
            variant="h3"
            noWrap
            component="a"
            href="#app-bar-with-responsive-menu"
            sx={{
              mr: 2,
              display: { xs: "none", mdlg: "flex" },
              fontWeight: 700,
              color: "inherit",
              textDecoration: "none",
            }}
          >
            {themeConfig.projectName}
          </Typography>

          <CircleIcon
            sx={{
              display: { xs: "flex", mdlg: "none" },
              height: 40,
              width: 40,
              mr: 1,
            }}
          />
          <Typography
            variant="h5"
            noWrap
            component="a"
            href="#app-bar-with-responsive-menu"
            sx={{
              mr: 2,
              display: { xs: "flex", mdlg: "none" },
              flexGrow: 1,
              fontWeight: 700,
              color: "inherit",
              textDecoration: "none",
            }}
          >
            {themeConfig.projectName}
          </Typography>

          <Box
            sx={{
              flexGrow: 0,
              display: { xs: "none", mdlg: "flex" },
              ml: "auto",
              gap: 2,
            }}
          >
            {navigation.map((item, index) => (
              <NavItem key={index} {...item} />
            ))}
            <IconButton onClick={handleOpenSettingsMenu} >
              <SettingsIcon />
            </IconButton>
            {SettingsMenu}
          </Box>

          <Box
            sx={{
              flexGrow: 0,
              display: { xs: "flex", mdlg: "none" },
            }}
          >
            <IconButton
              size="large"
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              onClick={handleOpenNavMenu}
              color="inherit"
            >
              <MenuIcon />
            </IconButton>
            <Menu
              id="menu-appbar"
              anchorEl={anchorElNav}
              anchorOrigin={{
                vertical: "bottom",
                horizontal: "left",
              }}
              keepMounted
              transformOrigin={{
                vertical: "top",
                horizontal: "left",
              }}
              open={Boolean(anchorElNav)}
              onClose={handleCloseNavMenu}
              sx={{
                display: { xs: "block", mdlg: "none" },
                mt: "1px",
                "& .MuiMenu-paper": { backgroundColor: "#0A3B7A" },
              }}
            >
              <Box sx={{ display: "flex", flexDirection: "column" }}>
                {navigation.map((item, index) => (
                  <NavItem key={index} {...item} />
                ))}
                <IconButton onClick={handleOpenSettingsMenu} color="inherit">
                  <SettingsIcon />
                </IconButton>
              </Box>
              {SettingsMenu}
            </Menu>
          </Box>
        </Toolbar>
      </Container>
    </AppBar>
  );
}

export default ResponsiveAppBar;
