// import i18n from "i18next";
// import { initReactI18next } from "react-i18next";
// import Backend from "i18next-http-backend";
// import I18nextBrowserLanguageDetector from "i18next-browser-languagedetector";
// import themeConfig from "./themeConfig";

// i18n

//   // Enables the i18next backend
//   .use(Backend)

//   // Enable automatic language detection
//   .use(I18nextBrowserLanguageDetector)

//   // Enables the hook initialization module
//   .use(initReactI18next)
//   .init({
//     lng: "en",
//     defaultLng: themeConfig.defaultLng,
//     backend: {
//       /* translation file path */
//       loadPath: "/locales/{{lng}}/common.json",
//     },
//     fallbackLng: themeConfig.defaultLng,
//     debug: false,
//     keySeparator: false,
//     react: {
//       useSuspense: false,
//     },
//     interpolation: {
//       escapeValue: false,
//       formatSeparator: ",",
//     },
//   });

// export default i18n;

// i18n.js

import i18n from "i18next";
import { initReactI18next } from "react-i18next";
import Backend from "i18next-http-backend";
import I18nextBrowserLanguageDetector from "i18next-browser-languagedetector";
import themeConfig from "./themeConfig";

i18n
  .use(Backend)
  .use(I18nextBrowserLanguageDetector)
  .use(initReactI18next)
  .init({
    fallbackLng: themeConfig.defaultLng,

    backend: {
      loadPath: "/locales/{{lng}}/common.json",
    },

    detection: {
      order: ["localStorage", "navigator", "cookie", "htmlTag"],

      caches: ["localStorage"],
    },

    debug: false,
    keySeparator: false,
    react: {
      useSuspense: false,
    },
    interpolation: {
      escapeValue: false,
      formatSeparator: ",",
    },
  });

export default i18n;
