import { useTranslation } from "react-i18next";

const Translations = ({ text }) => {
  const { t } = useTranslation();

  return <>{t(text)}</>;
};

export default Translations;
