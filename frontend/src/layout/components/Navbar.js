import * as React from "react";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";
import Menu from "@mui/material/Menu";
import MenuIcon from "@mui/icons-material/Menu";
import Container from "@mui/material/Container";
import CircleIcon from "@mui/icons-material/Circle";
import themeConfig from "src/configs/themeConfig";
import navigation from "src/navigation";
import NavItem from "./navigation/item/NavItem";
import LanguageSelector from "./navigation/item/LanguageSelector";
import { MenuItem } from "@mui/material";

function ResponsiveAppBar() {
  const [anchorElNav, setAnchorElNav] = React.useState(null);

  const handleOpenNavMenu = (event) => {
    setAnchorElNav(event.currentTarget);
  };

  const handleCloseNavMenu = () => {
    setAnchorElNav(null);
  };

  return (
    <AppBar
      // position="static" // removed due to incorrect appreance
      sx={{ backgroundColor: "#0A3B7A", boxShadow: "none" }}
    >
      <Container maxWidth="xl">
        <Toolbar disableGutters variant="dense" sx={{ mx: 7 }}>
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

          <Box sx={{ flexGrow: 0, display: { xs: "flex", mdlg: "none" } }}>
            <IconButton
              size="large"
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              s
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
              {navigation.map((item, index) => (
                <MenuItem>
                  <NavItem key={index} {...item} />
                </MenuItem>
              ))}
              <LanguageSelector isMenu={true} />
            </Menu>
          </Box>

          <Box
            sx={{
              flexGrow: 0,
              display: { xs: "none", mdlg: "flex" },
              ml: "auto",
              gap: 13,
            }}
          >
            {navigation.map((item, index) => (
              <NavItem key={index} {...item} />
            ))}
            <LanguageSelector />
          </Box>
        </Toolbar>
      </Container>
    </AppBar>
  );
}
export default ResponsiveAppBar;
