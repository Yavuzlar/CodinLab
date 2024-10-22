import CustomBreadcrumbs from "src/components/breadcrumbs";
import CodeEditor from "src/components/code-editor";
import { useTranslation } from "react-i18next";
import { useTheme } from "@mui/material/styles";
import {
  Card,
  CardContent,
  Typography,
  Box,
  useMediaQuery,
  Button,
  Modal,
  Tooltip,
  Alert,
} from "@mui/material";
import TestTubeGreen from "src/assets/icons/icons8-test-tube-100-green.png";
import TestTubeOrange from "src/assets/icons/icons8-test-tube-100-orange.png";
import TestTubeRed from "src/assets/icons/icons8-test-tube-100-red.png";
import LightBulb from "src/assets/icons/light-bulb.png";
import Image from "next/image";
import { use, useEffect, useRef, useState } from "react";
import { getProgrammingId } from "src/data/programmingIds";
import { useDispatch, useSelector } from "react-redux";
import { getLabByProgramingId } from "src/store/lab/labSlice";
import { data } from "autoprefixer";
import { getStop } from "src/store/code/codeSlice";
import RestartAltIcon from "@mui/icons-material/RestartAlt";
import axios from "axios";

const renderDifficulty = (difficulty) => {
  const { t } = useTranslation();

  switch (difficulty) {
    case 1:
      return (
        <Box
          sx={{
            display: "flex",
            gap: "0.4rem",
            backgroundColor: "#BDEEAF",
            borderRadius: "0.7rem",
            py: 1,
            px: 2,
          }}
        >
          <Image src={TestTubeGreen} alt="easy" width={25} height={25} />
          <Typography
            variant="body1"
            color={"#39CE19"}
            fontWeight={500}
            sx={{ textTransform: "capitalize" }}
          >
            {t("labs.difficulty.easy")}

          </Typography>
        </Box>
      );
    case 2:
      return (
        <Box
          sx={{
            display: "flex",
            gap: "0.4rem",
            backgroundColor: "#F3C9A5",
            borderRadius: "0.7rem",
            py: 1,
            px: 2,
          }}
        >
          <Image src={TestTubeOrange} alt="easy" width={25} height={25} />
          <Typography
            variant="body1"
            color={"#F07C1C"}
            fontWeight={500}
            sx={{ textTransform: "capitalize" }}
          >
            {t("labs.difficulty.medium")}
          </Typography>
        </Box>
      );
    case 3:
      return (
        <Box
          sx={{
            display: "flex",
            gap: "0.4rem",
            backgroundColor: "#F3B3B3",
            borderRadius: "0.7rem",
            py: 1,
            px: 2,
          }}
        >
          <Image src={TestTubeRed} alt="easy" width={25} height={25} />
          <Typography
            variant="body1"
            color={"#E00404"}
            fontWeight={500}
            sx={{ textTransform: "capitalize" }}
          >
            {t("labs.difficulty.hard")}
          </Typography>
        </Box>
      );
  }
};

const LabQuestion = ({ language = "", questionId }) => {
  const { lab: labSlice } = useSelector((state) => state);

  const { t, i18n } = useTranslation();
  const theme = useTheme();
  const isMobile = useMediaQuery((theme) => theme.breakpoints.down("smd"));

  const [output, setOutput] = useState(""); // we will store the output here

  const [isSubmitted, setIsSubmitted] = useState(false);

  const [isCompleted, setIsCompleted] = useState(false);

  const [isFailed, setIsFailed] = useState(false);

  const [labData, setLabData] = useState({
    title: "",
    programmingName: "",
    difficulty: "",
    description: "",
    questionNote: "",
    expectedOutputNote: "",
    expectedOutput: "",
    hint: "",
    template: ""
  });

  const [open, setOpen] = useState(false);

  const handleOpen = () => setOpen(true);

  const handleClose = () => setOpen(false);

  const _language = language.toUpperCase();

  const programmingID = language;


  const apiData = {
    programmingId: programmingID,
    pathId: questionId,
    endPoint: 'lab'
  };

  const editorRef = useRef(null);
  const dispatch = useDispatch();

  useEffect(() => {

    if (labSlice.data) {
      setLabData({
        title: labSlice.data[0]?.language?.title,
        programmingName: labSlice.data[0]?.programmingName,
        difficulty: labSlice.data[0]?.difficulty,
        description: labSlice.data[0]?.language?.description,
        questionNote: labSlice.data[0]?.language?.note,
        expectedOutputNote: labSlice.data[0]?.language?.expectedOutputNote,
        expectedOutput: labSlice.data[0]?.language?.expectedOutput,
        hint: labSlice.data[0]?.language?.hint,
        template: labSlice.data[0]?.template
      });
    }
  }, [labSlice.data]);

  const handleRun = (outputData) => {
    console.log(outputData)
    console.log(outputData?.data)
    setOutput(outputData?.data);
    setIsSubmitted(true);

    if (true) {
      setIsCompleted(true);
    } else {
      setIsFailed(true);
    }
  };


  const handleStop = (outputData) => {
    // this api for get stop code component (stop code component is the last component in the container)
    // but not use it in this component
    // dispatch(getStop())

    setOutput(outputData);
    setIsSubmitted(false);
    setIsCompleted(false);
  };

  const handleReset = async () => {
    try {
      const response = await axios({
        method: "GET",
        url: `/api/v1/private/lab/reset/${programmingID}/${questionId}`,
      });

      if (response.status === 200) {
        const apiTemplate = response.data?.data?.template || "";
        // const prevData = labData;
        // setLabData({
        //   ...labData,
        //   template: apiTemplate,
        // });
        editorRef.current.setValue(apiTemplate);

      }
    } catch (error) {
      console.log("Reset response error", error);
    }
  };


  useEffect(() => {
    dispatch(
      getLabByProgramingId({
        labID: questionId,
        programmingID: programmingID,
        language: i18n.language,
      })
    );
  }, [language, questionId]);


  // Breadcrumbs
  const breadcrums = [
    {
      path: `/labs/${language}`,
      title: t("nav.labs"),
      permission: "labs",
    },
    {
      path: `/labs/${language}`,
      title: labData.programmingName, // edit the sand backend programming id and take language namea.
      permission: "labs",
    },

    {
      path: `/labs/${language}/${questionId}`,
      title: labData.title,
      permission: "roads",
    },
  ];

  const params = {
    // these are the parameters for the component settings.
    height: !isMobile ? "30vw" : "90vw",
    width: !isMobile ? "40vw" : "98vw",
  };

  return (
    <>
      {/* Breadcrumbs */}
      <CustomBreadcrumbs titles={breadcrums} />

      {/* Outer container */}
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          gap: "1rem",
          flexDirection: isMobile ? "column" : "row",
          mt: 2,
        }}
      >
        {/* Question Description Card */}
        <Card
          sx={{
            width: isMobile ? "100%" : "50%",
            position: "relative",
          }}
        >
          <CardContent
            sx={{
              display: "flex",
              flexDirection: "column",
              gap: "1rem",
              mt: isMobile ? "2.5rem" : "0",
            }}
          >
            {/* Question Title */}
            <Typography variant="h4" fontWeight={600} >
              {labData.title}
            </Typography>

            {/* Difficulty, Hint button*/}
            <Box
              sx={{
                display: "flex",
                gap: "1rem",
                position: "absolute",
                bottom: "1rem",
                right: "1rem",
              }}
            >
              {renderDifficulty(labData.difficulty)}

              <Button
                onClick={handleOpen}
                sx={{
                  display: "flex",
                  gap: "0.4rem",
                  backgroundColor: "#FDEDAE",
                  borderRadius: "0.7rem",
                  py: 1,
                  px: 2,
                }}
              >
                <Image src={LightBulb} alt="hint" width={25} height={25} />
                <Typography variant="body1" color={"#FFCA00"} fontWeight={500}>
                  {t("labs.question.hint")}
                </Typography>
              </Button>
              <Tooltip title={t("roads.path.restart.button")}>
                <Button
                  variant="dark"
                  sx={{

                  }}
                  onClick={() => handleReset()}
                >
                  <RestartAltIcon />
                </Button>
              </Tooltip>
            </Box>

            <Modal
              open={open}
              onClose={handleClose}
              aria-labelledby="modal-modal-title"
              aria-describedby="modal-modal-description"
            >
              <Box sx={
                {
                  position: 'absolute',
                  top: '50%',
                  left: '50%',
                  transform: 'translate(-50%, -50%)',
                  maxWidth: 500,
                  bgcolor: theme.palette.primary.main,
                  p: 4,
                  borderRadius: '15px',
                  textAlign: 'center'
                }
              }>
                <Typography
                  id="modal-modal-title"
                  variant="h6"
                  component="h2"
                >
                  {t("labs.question.hint")}
                </Typography>
                <Typography id="modal-modal-description" sx={{ mt: 2 }}>
                  {labData.hint}
                </Typography>
              </Box>
            </Modal>


            {/* Question Description */}
            <Typography variant="body1">{labData.description}</Typography>

            {output && output.output && (
              <Alert
                severity={output.isCorrect ? "success" : "error"}
                variant="filled"
                sx={{ color: theme.palette.common.white, marginBottom: "10px", borderRadius: "10px" }}
              >
                {output.isCorrect ? t("CODE_SUCCESS") : `${t("CODE_ALERT").replace("$$$", output.expectedOutput).replace("***", output.output)}`}
              </Alert>
            )}
            {/* Question Note */}
            <Box
              sx={{
                backgroundColor: theme.palette.primary.dark,
                borderRadius: "1rem",
                padding: "2rem",
              }}
            >
              <strong>Note:</strong>
              {<br />}
              {<br />}
              <Typography variant="body1">{labData.questionNote}</Typography>
            </Box>
          </CardContent>
        </Card>


        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            gap: "1rem",
          }}
        >
          {/* Editor */}
          <CodeEditor
            params={params}
            onRun={handleRun}
            onStop={handleStop}
            leng={language}
            // defValue={labData.template}
            title={labData.title}
            apiData={apiData}
            val={labData?.template}
            editorRef={editorRef}
          />

          {/* Compilation messages after submitting */}
          {isCompleted && (
            <Typography
              variant="body1"
              fontWeight={"700"}
              color={"#39CE19"}
              sx={{ ml: 2 }}
            >
              {t("labs.question.completed")}
            </Typography>
          )}

          {isFailed && (
            <Typography
              variant="body1"
              fontWeight={"700"}
              color={"#e00404"}
              sx={{ ml: 2 }}
            >

              {t("labs.question.failed")}
            </Typography>
          )}

          {/* Expected output card */}

          {/* Output */}
          {isSubmitted && (
            <Card
              sx={{
                width: "100%",
                backgroundColor: isSubmitted ? "#0A3B7A" : "",
              }}
            >
              <CardContent
                sx={{ display: "flex", gap: "1rem", flexDirection: "column" }}
              >
                <Box
                  sx={{
                    display: "flex",
                    justifyContent: "space-between",
                    alignItems: "center",
                  }}
                >
                  <Box sx={{ width: "21%" }}>
                    <Typography variant="body1" fontWeight={"bold"}>
                      {t("labs.question.output")}{" "}
                    </Typography>
                  </Box>
                  <Box
                    sx={{
                      width: "100%",
                      borderRadius: "0.6rem",
                      backgroundColor: isSubmitted ? "#C3FFD3" : "#DAF0FE",
                      px: 2,
                      py: 1,
                    }}
                  >
                    <Typography
                      variant="body1"
                      color={"black"}
                      sx={{ whiteSpace: 'pre-line' }}
                    >
                      {output?.output}
                    </Typography>
                  </Box>
                </Box>
              </CardContent>
            </Card>
          )}

        </Box>
      </Box >
    </>
  );
};

export default LabQuestion;