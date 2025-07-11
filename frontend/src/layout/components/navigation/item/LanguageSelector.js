import { Button, Menu, MenuItem, Typography } from "@mui/material";
import { useState, useEffect } from "react";
import Image from "next/image";
import Turkish from "src/assets/flags/turkish.png";
import English from "src/assets/flags/english.png";
import { useTranslation } from "react-i18next";

const LanguageSelector = () => {
  const [anchorEl, setAnchorEl] = useState(null);
  const { i18n, t } = useTranslation();

  const open = Boolean(anchorEl);

  const currentLanguageCode = i18n.language;

  const handleClick = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleLanguageChange = (langCode) => {
    i18n.changeLanguage(langCode);
    handleClose();
  };

  useEffect(() => {
    const handleResize = () => {
      setAnchorEl(null);
    };
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, []);

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
        <Image
          src={currentLanguageCode === "tr" ? Turkish : English}
          alt={currentLanguageCode === "tr" ? "Turkish Flag" : "English Flag"}
          width={24}
          height={24}
        />
        <Typography
          sx={{
            fontWeight: 300,
            textTransform: "capitalize",
            fontFamily: "Outfit",
            textAlign: "center",
            ml: 0.5,
          }}
        >
          {currentLanguageCode === "tr" ? tr : en}
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
        <MenuItem onClick={() => handleLanguageChange("tr")}>{tr}</MenuItem>
        <MenuItem onClick={() => handleLanguageChange("en")}>{en}</MenuItem>
      </Menu>
    </>
  );
};

export default LanguageSelector;
