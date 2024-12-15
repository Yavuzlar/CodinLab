import { useTheme } from "@mui/material/styles";
import {
  Card,
  CardContent,
  Typography,
  Box,
  Button,
  useMediaQuery,
  Alert,
  Grid,
} from "@mui/material";
import RestartAltIcon from "@mui/icons-material/RestartAlt";
import Tooltip from "@mui/material/Tooltip";
import CodeEditor from "src/components/code-editor";
import Output from "src/components/output";
import { useState, useEffect, useRef } from "react";
import CustomBreadcrumbs from "src/components/breadcrumbs";
import DoneIcon from "src/assets/icons/icons8-done-100 (1).png";
import Image from "next/image";
import { useTranslation } from "react-i18next";
import { useDispatch, useSelector } from "react-redux";
import { fetchPathById } from "src/store/path/pathSlice";
import { useRouter } from "next/router";
import { useAuth } from "src/hooks/useAuth";
import axios from "axios";
import ModalRoad from "src/components/modal/ModalRoad";

const LanguageRoad = ({ language = "", pathId }) => {
  // Language to be displayed
  const _language = language.toUpperCase();

  const { t, i18n } = useTranslation();
  const router = useRouter();
  const theme = useTheme();
  const { sendHistory } = useAuth();

  const { path } = useSelector((state) => state);
  const editorRef = useRef(null);
  const dispatch = useDispatch();

  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const [output, setOutput] = useState({});

  const [programmingId, setProgrammingId] = useState(null);

  const [languageName, setLanguageName] = useState("");
  const [isStarted, setIsStarted] = useState(false);
  const [isFinished, setIsFinished] = useState(false);
  const [extension, setExtension] = useState("");
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [content, setContent] = useState("");
  const [note, setNote] = useState("");
  const [template, setTemplate] = useState("");
  const [fileExtension, setFileExtension] = useState("");
  const [monacoEditor, setMonacoEditor] = useState("");
  const [userCode, setUserCode] = useState("");

  const _mdmd = useMediaQuery((theme) => theme.breakpoints.down("mdmd"));

  // API data
  const apiData = {
    programmingId: programmingId,
    pathId: pathId,
    endPoint: "road",
  };

  // Breadcrumbs
  const breadcrums = [
    {
      path: "/roads",
      title: "Roads",
      permission: "roads",
    },
    {
      path: `/roads/${language}`,
      title: languageName,
      permission: "roads",
    },
    {
      path: `/roads`,
      title: title,
      permission: "roads",
    },
  ];

  // Params for the code editor
  const params = {
    height: "30rem",
    width: "100%",
  };

  // Set the programming id
  useEffect(() => {
    setProgrammingId(language);
  }, [language]);

  // Fetch the path by id
  useEffect(() => {
    if (programmingId && pathId) {
      dispatch(
        fetchPathById({
          language: i18n.language,
          programmingId: programmingId,
          pathId: pathId,
        })
      );
    }
  }, [programmingId, pathId, i18n.language]);

  // Set the path data
  useEffect(() => {
    if (path) {
      if (path.data.data) {
        const pathData = path.data.data[0].paths[0];
        setExtension(path?.data?.data[0]?.fileExtension);
        setIsStarted(pathData.pathIsStarted);
        setIsFinished(pathData.pathIsFinished);
        setTitle(pathData.language.title);
        setDescription(pathData.language.description);
        setContent(pathData.language.content);
        setNote(pathData.language.note);
        setTemplate(pathData.template);
        setFileExtension(path.data.data[0].fileExtension);
        setMonacoEditor(path.data.data[0].monacoEditor);
        setLanguageName(path.data.data[0].name);
        setUserCode(pathData.template);
      }

      setError(path.error);
      setLoading(path.loading);
    }
  }, [path]);

  // Set isFinished to true if the output is correct
  useEffect(() => {
    if (output.isCorrect) {
      setIsFinished(true);
    }
  }, [output]);

  // Handle Code Editor functions
  const handleRun = (outputData) => {
    setOutput(outputData?.data);
  };

  const handleStop = (outputData) => {
    setOutput(outputData);
  };

  const handleNextPath = () => {
    router.push(`/roads/${language}/${parseInt(pathId) + 1}`);
  };
  // End Handle Code Editor functions

  // Reset the path
  const handleReset = async () => {
    try {
      const response = await axios({
        method: "GET",
        url: `/api/v1/private/road/reset/${programmingId}/${pathId}`,
      });
      if (response.status === 200) {
        const apiTemplate = response.data?.data?.template || "";
        editorRef.current.setValue(apiTemplate);
      }
    } catch (error) {
      console.log("Reset response error", error);
    }
  };

  const handleBeforeUnload = (event) => {
    const labPathType = "Road";

    sendHistory(
      userCode,
      parseInt(programmingId),
      parseInt(pathId),
      labPathType
    );

    event?.preventDefault();
  };

  const handleChange = (outputData) => {
    setUserCode(outputData);
  };

  // Trigger sendHistory when the user leaves the page
  useEffect(() => {
    window.addEventListener("beforeunload", handleBeforeUnload);

    return () => {
      window.removeEventListener("beforeunload", handleBeforeUnload);
    };
  }, []);

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
            {title}
          </Typography>
          <Typography
            variant="body1"
            sx={{
              mt: "10px",
              mb: "40px",
              color: "lightgrey",
              whiteSpace: "pre-line",
            }}
          >
          <Typography
            variant="body1"
            sx={{
              mt: "10px",
              mb: "40px",
              color: "lightgrey",
              whiteSpace: "pre-line",
            }}
          >
            {content}
          </Typography>
          <ModalRoad buttonMessage={t("road.modal.button")} message={note} />
          {isFinished && (
            <Box sx={{ position: "absolute", right: "1rem", top: "1rem" }}>
              <Image src={DoneIcon} height={30} width={30} alt="done" />
            </Box>
          )}
          {!isFinished && (
            <Tooltip title={t("roads.path.restart.button")}>
              <Button
                variant="dark"
                sx={{
                  position: "absolute",
                  right: "1rem",
                  top: "1rem",
                  minWidth: "1rem",
                }}
                onClick={handleReset}
              >
                <RestartAltIcon />
              </Button>
            </Tooltip>
          )}

          <Button
            variant="light"
            sx={{
              position: "absolute",
              right: "1rem",
              bottom: "1rem",
              fontWeight: 700,
              fontFamily: "Outfit",
              textTransform: "capitalize",
              py: 1,
              px: 3,
            }}
            onClick={handleNextPath}
            disabled={!isFinished}
          >
            {t("roads.path.next_path")}
          </Button>
        </CardContent>
      </Card>
      {output && output.output && (
        <Alert
          severity={output.isCorrect ? "success" : "error"}
          variant="filled"
          sx={{
            color: theme.palette.common.white,
            marginBottom: "10px",
            borderRadius: "10px",
          }}
          sx={{
            color: theme.palette.common.white,
            marginBottom: "10px",
            borderRadius: "10px",
          }}
        >
          {output.isCorrect
            ? t("CODE_SUCCESS")
            : `${t("CODE_ALERT")
                .replace("$$$", output.expectedOutput)
                .replace("***", output.output)}`}
        </Alert>
      )}
      <Grid container spacing={2}>
        <Grid item xs={12} md={6}>
          <CodeEditor
            key={template}
            onRun={handleRun}
            onStop={handleStop}
            onChange={handleChange}
            leng={monacoEditor}
            title={`example.${extension}`}
            apiData={apiData}
            editorRef={editorRef}
            val={template}
            params={params}
          />
        </Grid>
        <Grid item xs={12} md={6}>
          <Output params={params} value={output} />
        </Grid>
      </Grid>
    </>
  );
};

export default LanguageRoad;
