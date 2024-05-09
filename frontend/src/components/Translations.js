import { Typography } from "@mui/material";
import { useTranslation } from "next-i18next";

const Translations = ({ text }) => {
  const { t } = useTranslation();

  return <>{t(text)}</>;
};

export default Translations;
