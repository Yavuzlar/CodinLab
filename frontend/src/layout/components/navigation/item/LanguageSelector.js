import { Button, Menu, MenuItem, Typography } from "@mui/material";
import { useState } from "react";
import Image from "next/image";
import Turkish from "src/assets/flags/turkish.png";
import English from "src/assets/flags/english.png";

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
        {language === "turkish" && (
          <Image src={Turkish} width={30} height={30} />
        )}

        {language === "english" && (
          <Image src={English} width={30} height={30} />
        )}

        <Typography
          sx={{
            fontWeight: 300,
            textTransform: "capitalize",
            opacity: 0.6,
            fontFamily: "Outfit",
            textAlign: "center",
            ml: 0.5,
            mb: 0.5,
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
        }}
        sx={{ mt: 1, "& .MuiMenu-paper": { backgroundColor: "#0A3B7A" } }}>
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
