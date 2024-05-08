import { Typography } from "@mui/material";
import { useTranslation } from "next-i18next";

const Translations = ({ text }) => {
  const { t } = useTranslation();

  return <Typography>{t(text)}</Typography>;
};

export default Translations;
