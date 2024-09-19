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
import { fetchPathById, resetPathById, sendAnswerById } from "src/store/path/pathSlice";

const LanguageRoad = ({ language = "", pathId }) => {
  const [output, setOutput] = useState(""); // we will store the output here

  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const [programmingId, setProgrammingId] = useState(null)

  const _language = language.toUpperCase();

  const { t, i18n } = useTranslation();

  const dispatch = useDispatch();
  const { path } = useSelector((state) => state);

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
      console.log("Path fetched ->", path);

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
  };

  const theme = useTheme();

  const title = "Basic " + _language + " Syntax";
  const description =
    'Line 1: #include <stdio.h> is a header file library that lets us work with input and output functions, such as printf() (used in line 4). Header files add functionality to C programs. Line 2: A blank line. C ignores white space. But we use it to make the code more readable. Line 3: Another thing that always appear in a C program is main(). This is called a function. Any code inside its curly brackets {} will be executed. Line 4: printf() is a function used to output/print text to the screen. In our example, it will output "Hello World!". Line 5: return 0 ends the main() function. Line 6: Do not forget to add the closing curly bracket } to actually end the main function.';

  const params = {
    // these are the parameters for the component settings.
    height: "50vh",
    width: "50vw",
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
      title: _language,
      permission: "roads",
    },

    {
      path: `/roads`,
      title: title,
      permission: "roads",
    },
  ];

  // API response example
  const data = {
    data: {
      difficulty: 0,
      id: 0,
      isFinished: true,
      isStarted: true,
      languages: [
        {
          content: "string",
          description: "string",
          lang: "string",
          note: "string",
          title: "string",
        },
      ],
      name: "string",
    },
    data_count: 0,
    errors: "string",
    message: "string",
    status_code: 0,
  };

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
          {data.data.isFinished && (
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
            disabled={!data.data.isFinished}
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
          defValue={"// deneme"}
          title={"deneme.c"}
        />
        <Output value={output} params={params} />
      </Box>
    </>
  );
};

export default LanguageRoad;
