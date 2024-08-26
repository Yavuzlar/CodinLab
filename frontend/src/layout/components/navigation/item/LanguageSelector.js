import { Button, Menu, MenuItem, Typography } from "@mui/material";
import { useEffect, useState } from "react";
import Image from "next/image";
import Turkish from "src/assets/flags/turkish.png";
import English from "src/assets/flags/english.png";
import { useTranslation } from "react-i18next";


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

  const { i18n } = useTranslation();

  useEffect(() => {
    switch (language) {
      case "turkish":
        i18n.changeLanguage("tr");
        break;

      case "english":
        i18n.changeLanguage("en");

      default:
        i18n.changeLanguage("en");
        break;
    }
  }, [language]);

  useEffect(() => {
    const handleResize = () => {
      setAnchorEl(null);
    };
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, [setAnchorEl]);

  const { t } = useTranslation();
  const tr = t("languages.tr");
  const en = t("languages.en");

  return (
    <>
      <Button
        id="basic-button"
        aria-controls={open ? "basic-menu" : undefined}
        aria-haspopup="true"
        aria-expanded={open ? "true" : undefined}
        sx={{ backgroundColor: "transparent" }}
        onClick={handleClick}
      >
        {language === "turkish" && (
          <Image src={Turkish} width={24} height={24} />
        )}

        {language === "english" && (
          <Image src={English} width={24} height={24} />
        )}

        <Typography
          sx={{
            fontWeight: 300,
            textTransform: "capitalize",
            fontFamily: "Outfit",
            textAlign: "center",
            ml: 0.5,
          }}
        >
          {language === "turkish" && tr}

          {language === "english" && en}
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
        sx={{ mt: 1, "& .MuiMenu-paper": { backgroundColor: "#0A3B7A" } }}
      >
        <MenuItem
          onClick={(e) => {
            setLanguage("turkish");
            setAnchorEl(false);
          }}
        >
          {tr}
        </MenuItem>
        <MenuItem
          onClick={(e) => {
            setLanguage("english");
            setAnchorEl(false);
          }}
        >
          {en}
        </MenuItem>
      </Menu>
    </>
  );
};

export default LanguageSelector;
