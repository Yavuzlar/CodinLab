import CircleIcon from "@mui/icons-material/Circle";
import { Typography } from "@mui/material";
import themeConfig from "src/configs/themeConfig";

const Title = () => {
  const textStyle = {
    fontWeight: "bold",
    flexGrow: 1,
    ml: 1,
  };
  return (
    <>
      <CircleIcon sx={{ width: 40, height: 40 }} />
      <Typography variant="h4" sx={textStyle}>
        {themeConfig.templateName}
      </Typography>
    </>
  );
};

export default Title;
