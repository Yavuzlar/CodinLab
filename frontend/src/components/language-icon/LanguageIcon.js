import { Box, Typography } from "@mui/material";
import Image from "next/image";
import React from "react";
import C from "../../assets/language/c.png";
import Cpp from "../../assets/language/cpp.png";
import Go from "../../assets/language/go.png";
import Js from "../../assets/language/javascript.png";
import Python from "../../assets/language/python.png";

const languageIcons = {
  C: C,
  "C++": Cpp,
  GO: Go,
  JavaScript: Js,
  Python: Python,
};

const LanguageIcon = ({ language }) => {
  const icon = languageIcons[language];

  return (
    <Box
      sx={{
        display: "flex",
        alignItems: "center",
        justifyContent: "start",
        gap: "0.5rem",
      }}
    >
      {icon ? (
        <Image src={icon} width={25} height={25} />
      ) : (
        <Box sx={{ width: 25, height: 25 }}>
          <Typography>-</Typography>
        </Box>
      )}
      <Box
        component="span"
        sx={{
          minWidth: "80px",
        }}
      >
        {language}
      </Box>
    </Box>
  );
};

export default LanguageIcon;
