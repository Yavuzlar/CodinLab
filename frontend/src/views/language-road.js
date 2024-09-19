import { useTheme } from "@mui/material/styles";
import { Card, CardContent, Typography, Box, Button } from "@mui/material";
import CodeEditor from "src/components/code-editor";
import Output from "src/components/output";
import { useState, useEffect } from "react";
import CustomBreadcrumbs from "src/components/breadcrumbs";
import DoneIcon from "src/assets/icons/icons8-done-100 (1).png";
import Image from "next/image";
import { useTranslation } from "react-i18next";
import { getProgrammingId } from "src/data/programmingIds";
import { useDispatch, useSelector } from "react-redux";
import { fetchPathById, resetPathById } from "src/store/path/pathSlice";
import { useRouter } from "next/router";

const LanguageRoad = ({ language = "", pathId }) => {
  const [output, setOutput] = useState(""); // we will store the output here

  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const [programmingId, setProgrammingId] = useState(null)

  const _language = language.toUpperCase();

  const { t, i18n } = useTranslation();
  const router = useRouter();


  const dispatch = useDispatch();
  const { path } = useSelector((state) => state);

  const [isStarted, setIsStarted] = useState(false);
  const [isFinished, setIsFinished] = useState(false);

  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [content, setContent] = useState("");
  const [note, setNote] = useState("");
  const [template, setTemplate] = useState("");

  useEffect(() => {
    console.log("Language useEffect: ", language);
    setProgrammingId(getProgrammingId[language]);
  }, [language]);

  useEffect(() => {
    if (programmingId && pathId) {
      console.log("i18n language ->", i18n.language);
      dispatch(fetchPathById(
        {
          language: i18n.language,
          programmingId: programmingId,
          pathId: pathId,
        }
      ));
    }
  }, [programmingId, pathId, i18n.language]);

  useEffect(() => {
    if (path) {

      if (path.data.data) {

        const pathData = path.data.data[0].paths[0];

        setTitle(pathData.language.title);
        setDescription(pathData.language.description);
        setContent(pathData.language.content);
        setNote(pathData.language.note);
        setTemplate(pathData.template);
      }

      setError(path.error);
      setLoading(path.loading);


    }
  }, [path]);

  const handleRun = (outputData) => {
    // this function will be called when the code is run

    setOutput(outputData);
  };

  const handleStop = (outputData) => {
    // this function will be called when the code is stopped
    setOutput(outputData);
  };

  const handleNextPath = () => {
    // here we will add the next path api call
    router.push(`/roads/${language}/${pathId + 1}`);
  };

  const theme = useTheme();

  const params = {
    // these are the parameters for the component settings.
    height: "50vh",
    width: "50vw",
  };

  const apiData = {
    programmingId: programmingId,
    pathId: pathId,
    endPoint: "road"
  }

  // Breadcrumbs
  const breadcrums = [
    {
      path: "/roads",
      title: "Roads",
      permission: "roads",
    },
    {
      path: `/roads/${language}`,
      title: _language,
      permission: "roads",
    },

    {
      path: `/roads`,
      title: title,
      permission: "roads",
    },
  ];

  return (
    <>
      <CustomBreadcrumbs titles={breadcrums} />
      <Card
        sx={{
          position: "relative",
          backgroundColor: theme.palette.primary.dark,
          paddingY: 2,
          my: 2,
        }}
      >
        <CardContent>
          <Typography variant="h4" fontWeight={500}>
            {" "}
            {title}{" "}
          </Typography>
          <Typography variant="body1" sx={{ lineHeight: 2.5 }}>
            {" "}
            {description}{" "}
          </Typography>
          {isFinished && (
            <Box sx={{ position: "absolute", right: "1rem", top: "1rem" }}>
              <Image src={DoneIcon} height={30} width={30} alt="done" />
            </Box>
          )}
          <Button
            sx={{
              position: "absolute",
              right: "1rem",
              bottom: "1rem",
              backgroundColor: "#fff",
              color: theme.palette.primary.dark,
              fontWeight: 700,
              fontFamily: "Outfit",
              textTransform: "capitalize",
              py: 1,
              px: 3,
            }}
            onClick={handleNextPath}
            disabled={!isFinished}
          >
            {" "}
            {t("roads.path.next_path")}{" "}
          </Button>
        </CardContent>
      </Card>
      <Box sx={{ display: "flex", gap: 2 }}>
        <CodeEditor
          params={params}
          onRun={handleRun}
          onStop={handleStop}
          leng={language}
          defValue={template}
          title={"example.c"}
          apiData={apiData}
        />
        <Output value={output} params={params} />
      </Box>
    </>
  );
};

export default LanguageRoad;
