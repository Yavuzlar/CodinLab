import CustomBreadcrumbs from "src/components/breadcrumbs";
import CodeEditor from "src/components/code-editor";
import { useTranslation } from "react-i18next";
import { useTheme } from "@mui/material/styles";
import { Card, CardContent, Typography, Box, useMediaQuery } from "@mui/material";
import TestTubeGreen from "src/assets/icons/icons8-test-tube-100-green.png";
import TestTubeOrange from "src/assets/icons/icons8-test-tube-100-orange.png";
import TestTubeRed from "src/assets/icons/icons8-test-tube-100-red.png";
import LightBulb from "src/assets/icons/light-bulb.png";
import Image from "next/image"
import { useState } from "react";

const LabQuestion = ({ language = "", questionId }) => {
    const { t } = useTranslation();
    const theme = useTheme();
    const isMobile = useMediaQuery((theme) => theme.breakpoints.down("smd"));

    const [output, setOutput] = useState(""); // we will store the output here

    const [isSubmitted, setIsSubmitted] = useState(false);

    const [isCompleted, setIsCompleted] = useState(false);

    const [isFailed, setIsFailed] = useState(false);

    const _language = language.toUpperCase();

    // Question data
    const title = "Perfect Number";
    const difficulty = "medium";
    const description = "Develop a C program to check if a given number is a perfect number. The program should prompt the user to enter a number and print whether the number is a perfect number or not."
    const questionNote = "• You can use a loop to find the divisors of a number. You may need to add a conditional statement to exclude the number itself from the sum of divisors. • You can determine if a number is perfect by checking if the sum of divisors is equal to the number."
    const expectedOutputNote = "The first number that the user will enter must be 28."
    const expectedOutput = "28 is a perfect number."

    const handleRun = (outputData) => { // this function will be called when the code is run
    setOutput(outputData);
    setIsSubmitted(true);

    // check if the output is correct
    if (true) {
      setIsCompleted(true);
    } else {
      setIsFailed(true);
    }
    };

  const handleStop = (outputData) => { // this function will be called when the code is stopped
    setOutput(outputData);
    setIsSubmitted(false);
    setIsCompleted(false);
    }

    const renderDifficulty = (difficulty) => { 
      switch (difficulty) {
        case "easy":
          return (
            <Box sx={{display: "flex", gap: "0.4rem", backgroundColor: "#BDEEAF", borderRadius: "0.7rem", py: 1, px: 2}}>
              <Image src={TestTubeGreen} alt="easy" width={25} height={25}/>
              <Typography variant="body1" color={"#39CE19"} fontWeight={500} sx={{ textTransform: "capitalize" }}> {difficulty} </Typography>
            </Box>
          )
        case "medium":
          return (
            <Box sx={{display: "flex", gap: "0.4rem", backgroundColor: "#F3C9A5", borderRadius: "0.7rem", py: 1, px: 2}}>
              <Image src={TestTubeOrange} alt="easy" width={25} height={25}/>
              <Typography variant="body1" color={"#F07C1C"} fontWeight={500} sx={{ textTransform: "capitalize" }}> {difficulty} </Typography>
            </Box>
          )
        case "hard":
          return (
          <Box sx={{display: "flex", gap: "0.4rem", backgroundColor: "#F3B3B3", borderRadius: "0.7rem", py: 1, px: 2}}>
              <Image src={TestTubeRed} alt="easy" width={25} height={25}/>
              <Typography variant="body1" color={"#E00404"} fontWeight={500} sx={{ textTransform: "capitalize" }} > {difficulty} </Typography>
            </Box>)
      }
    }

    // Breadcrumbs
    const breadcrums = [
      {
      path: "/labs",
      title: "Labs",
      permission: "labs"
    },
      {
        path: `/labs/${language}`,
        title: _language,
        permission: "labs"
      }, 

      {
        path: `/labs/${language}/${questionId}`,
        title: title,
        permission: "roads"
      }
    ]

  const params = { // these are the parameters for the component settings.
    height: !isMobile ? "30vw" : "90vw",
    width: !isMobile ? "40vw" : "98vw",
  };

  return (
    <>

    {/* Breadcrumbs */}
    <CustomBreadcrumbs titles={breadcrums} />

    {/* Outer container */}
    <Box sx={{
      display: "flex",
      justifyContent: "center",
      gap: "1rem",
      flexDirection: isMobile ? "column" : "row",
      mt: 2
    }}>

      {/* Question Description Card */}
      <Card sx={{
        width: isMobile ? "100%" : "50%",
        position: "relative"
      }}>
        <CardContent sx={{
        display: "flex",
        flexDirection: "column",
        gap: "1rem",
        mt: isMobile ? "2.5rem" : "0",
        }}>
          
          {/* Question Title */}
          <Typography variant="h4" fontWeight={600}> {title} </Typography>

          { /* Difficulty, Hint button*/ }
          <Box sx={{display: "flex", gap: "1rem", position: "absolute",top: "1rem", right: "1rem"}}>
             
            {renderDifficulty(difficulty)}
          

            <Box sx={{display: "flex", gap: "0.4rem", backgroundColor: "#FDEDAE", borderRadius: "0.7rem", py: 1, px: 2}} >
              <Image src={LightBulb} alt="hint" width={25} height={25} />
              <Typography variant="body1" color={"#FFCA00"} fontWeight={500}> { t("labs.question.hint") } </Typography>
            </Box>
          </Box>

          {/* Question Description */}
          <Typography variant="body1">
            { description }
          </Typography>

          {/* Question Note */}
          <Box sx={{backgroundColor: theme.palette.primary.dark, borderRadius: "1rem", padding: "2rem"}}>
            <strong>Note:</strong>{<br/>}{<br/>}
            <Typography variant="body1">
              { questionNote }
            </Typography>
          </Box>

          {/* Expected output note */}
          <Box sx={{border: "2px solid #0A3B7A", borderRadius: "0.6rem", backgroundColor: "#CCE5FB"}}>
              <Typography variant="body1" fontWeight={"bold"} color={"#0A3B7A"} sx={{p: "1rem"}}>
              { expectedOutputNote }
            </Typography>
          </Box>
        </CardContent>
      </Card>


      <Box sx={{
        display: "flex",
        flexDirection: "column",
        gap: "1rem"
      }}>
        {/* Editor */}
        <CodeEditor params={params} onRun={handleRun} onStop={handleStop} leng={language} defValue={"#include <stdio.h>"} title={"perfectNumber.c"} />
        
        { /* Compilation messages after submitting */ }
        {isCompleted && (
          <Typography variant="body1" fontWeight={"700"} color={"#39CE19"} sx={{ ml: 2 }}> { t("labs.question.completed") } </Typography>
        )}

        {isFailed && (
          <Typography variant="body1" fontWeight={"700"} color={"#e00404"} sx={{ ml: 2 }}> { t("labs.question.failed") } </Typography>
        )}

        

        {/* Expected output card */}
        <Card sx={{width: "100%", backgroundColor: isSubmitted ? "#0A3B7A" : ""}}>
          <CardContent sx={{ display: "flex", gap: "1rem", flexDirection:"column" }}>
            {/* Output */}
            {isSubmitted && (
              <Box sx={{ display: "flex", justifyContent: "space-between", alignItems: "center" }}>
                  <Box sx={{ width: "21%" }}>
                    <Typography variant="body1" fontWeight={"bold"}>{ t("labs.question.output") } </Typography>
                  </Box>
                <Box sx={{width:"100%", borderRadius: "0.6rem", backgroundColor: isSubmitted ? "#C3FFD3" : "#DAF0FE", px: 2, py: 1}}>
                  <Typography variant="body1" fontFamily={"Cascadia Code"} color={"black"}> "{output}" </Typography>
                </Box>
              </Box>
            )}
            <Box sx={{ display: "flex", gap: "1rem", alignItems: "center" }}>
              <Typography variant="body1" fontWeight={"bold"}>{ t("labs.question.expected") } </Typography>
            <Box sx={{ width:"100%", borderRadius: "0.6rem", backgroundColor: isSubmitted ? "#C3FFD3" : "#DAF0FE", px: 2, py: 1 }}>
            <Typography variant="body1" fontFamily={"Cascadia Code"} color={"black"}>"{ expectedOutput }"</Typography>
            </Box>
            </Box>
          </CardContent>
        </Card>

      </Box>

    </Box>
    </>
  )
}

export default LabQuestion;