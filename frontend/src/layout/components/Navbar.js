import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Menu from "@mui/material/Menu";
import MenuIcon from "@mui/icons-material/Menu";
import Container from "@mui/material/Container";
import themeConfig from "src/configs/themeConfig";
import navigation from "src/navigation";
import NavItem from "./navigation/item/NavItem";
import LanguageSelector from "./navigation/item/LanguageSelector";
import LogoutIcon from "@mui/icons-material/Logout";
import { Divider, IconButton } from "@mui/material";
import { useAuth } from "src/hooks/useAuth";
import { useTranslation } from "react-i18next";
import { useNav } from "src/hooks/useNav";
import Logo from "../../assets/logo/codinlab-logo-light.png";
import Image from "next/image";

function ResponsiveAppBar() {
  const { logout } = useAuth();
  const { anchorElNav, OpenNavMenu, CloseNavMenu, LogoClick } = useNav();

  const handleLogout = async () => {
    try {
      await logout();
    } catch (error) {
    }
  };

  const { t } = useTranslation();

  const logoutText = t("logout");

  return (
    <AppBar
      sx={{ backgroundColor: "#0A3B7A", boxShadow: "none", py: "0.5rem" }}
    >
      <Container maxWidth="lgPlus">
        <Toolbar disableGutters variant="dense">
          <Box
            sx={{
              display: "flex",
              alignItems: "center",
              cursor: "pointer",
              gap: "1rem",
            }}
            onClick={LogoClick}
          >
            <Box
              sx={{
                display: { xs: "none", mdlg: "flex" },
              }}
            >
              <Image
                src={Logo}
                alt="CodinLab-Logo"
                width={"auto"}
                height={64}
              />
            </Box>
            <Typography
              variant="h3"
              noWrap
              component="a"
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
          </Box>

          <Box
            onClick={LogoClick}
            sx={{
              display: "flex",
              alignItems: "center",
              flexGrow: 1,
              cursor: "pointer",
              gap: "0.5rem",
            }}
          >
            <Box
              sx={{
                display: { xs: "flex", mdlg: "none" },
              }}
            >
              <Image
                src={Logo}
                alt="CodinLab-Logo"
                width={"auto"}
                height={40}
              />
            </Box>

            <Typography
              variant="h5"
              noWrap
              component="a"
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
          </Box>

          <Box
            sx={{
              display: { xs: "flex", mdlg: "none" },
              alignItems: "center",
            }}
          >
            <LanguageSelector isMenu={true} />
            <IconButton onClick={OpenNavMenu}>
              <MenuIcon />
            </IconButton>
          </Box>

          <Box
            sx={{
              flexGrow: 0,
              display: { xs: "none", mdlg: "flex" },
              ml: "auto",
              gap: "1.5rem",
              alignItems: "center",
            }}
          >
            {navigation.map((item, index) => (
              <NavItem key={index} {...item} />
            ))}
            <LanguageSelector />
            <IconButton
              onClick={handleLogout}
              sx={{
                display: "flex",
                flexDirection: "row",
                alignItems: "center",
                gap: "0.5rem",
                "&:hover": {
                  backgroundColor: "rgba(255, 255, 255, 0.3)",
                  borderRadius: "5px",
                },
              }}
            >
              <LogoutIcon
                sx={{
                  width: 24,
                  height: 24,
                }}
              />
              <Typography
                sx={{
                  fontWeight: 300,
                  textTransform: "capitalize",
                  fontFamily: "Outfit",
                  textAlign: "center",
                }}
              >
                {logoutText}
              </Typography>
            </IconButton>
          </Box>

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
            onClose={CloseNavMenu}
            sx={{
              display: { xs: "block", mdlg: "none" },
              mt: "1px",
              "& .MuiMenu-paper": {
                backgroundColor: "#0A3B7A",
              },
            }}
          >
            <Box sx={{ display: "flex", flexDirection: "column", gap: 1.5 }}>
              {navigation.map((item, index) => (
                <NavItem
                  key={index}
                  {...item}
                  sx={{
                    "&:hover": {
                      borderRadius: "0.938rem",
                      backgroundColor: "rgba(255, 255, 255, 0.1)", // hover effect
                    },
                  }}
                />
              ))}
              <Divider sx={{ borderColor: "white" }} />
              <Box sx={{ display: "flex", justifyContent: "center" }}>
                <IconButton
                  onClick={handleLogout}
                  sx={{
                    display: "flex",
                    flexDirection: "row",
                    alignItems: "center",
                    gap: "0.5rem",
                    "&:hover": {
                      backgroundColor: "rgba(255, 255, 255, 0.3)",
                      borderRadius: "5px",
                    },
                  }}
                >
                  <LogoutIcon />
                  <Typography>Logout</Typography>
                </IconButton>
              </Box>
            </Box>
          </Menu>
        </Toolbar>
      </Container>
    </AppBar>
  );
}

export default ResponsiveAppBar;
