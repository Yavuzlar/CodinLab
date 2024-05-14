import { Image } from "@mui/icons-material";
import { Button, Menu, MenuItem, Typography } from "@mui/material";
import { useState } from "react";
import LanguageIcon from "@mui/icons-material/Language";

const LanguageSelector = () => {
  const [anchorEl, setAnchorEl] = useState(null);
  const open = Boolean(anchorEl);
  const handleClick = (event) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };
  const [language, setLanguage] = useState("turkish");
  return (
    <>
      <Button
        id="basic-button"
        aria-controls={open ? "basic-menu" : undefined}
        aria-haspopup="true"
        aria-expanded={open ? "true" : undefined}
        sx={{ backgroundColor: "transparent" }}
        onClick={handleClick}>
        <LanguageIcon sx={{ mr: 1 }} />
        <Typography
          sx={{
            textTransform: "capitalize",
            opacity: 0.6,
            fontFamily: "Outfit",
          }}>
          {language}
        </Typography>
      </Button>
      <Menu
        id="basic-menu"
        anchorEl={anchorEl}
        open={open}
        onClose={handleClose}
        MenuListProps={{
          "aria-labelledby": "basic-button",
        }}>
        <MenuItem
          onClick={(e) => {
            setLanguage("turkish");
            setAnchorEl(false);
          }}>
          Turkish
        </MenuItem>
        <MenuItem
          onClick={(e) => {
            setLanguage("english");
            setAnchorEl(false);
          }}>
          English
        </MenuItem>
      </Menu>
    </>
  );
};

export default LanguageSelector;
