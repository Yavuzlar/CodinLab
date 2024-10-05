import { Box, Button, Card, Typography } from "@mui/material";
import TestTubeRed from "../../assets/icons/red.png";
import TestTubeOrange from "../../assets/icons/orange.png";
import TestTubeGreen from "../../assets/icons/green.png";
import Image from "next/image";
import Translations from "../Translations";
import { useDispatch, useSelector } from "react-redux";
import { useEffect } from "react";
import { getLabsById } from "src/store/lab/labSlice";
import { useTranslation } from "react-i18next";
import { useRouter } from "next/router";
const LabInfo = ({ programingId }) => {
  const dispatch = useDispatch();
  const { lab: stateLabs } = useSelector((state) => state);

  const router = useRouter();

  const { t, i18n } = useTranslation();

  useEffect(() => {
    dispatch(
      getLabsById({
        programmingID: programingId,
        language: i18n.language,
      })
    );
  }, [dispatch, programingId, i18n.language]);

  return (
    <Box
      sx={{
        display: "flex",
        gap: "1.5rem",
        flexWrap: "wrap",
        justifyContent: "center",
      }}
    >
      {stateLabs.data?.labs?.map((lab, index) => (
        <Card
          sx={{
            width: "375px",
            height: "170px",
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            flexDirection: "column",
            gap: "1rem",
            padding: "1rem",
            borderRadius: "16px",
          }}
          key={index}
        >
          <Box
            sx={{
              display: "flex",
              justifyContent: "center",
              alignItems: "center",
              flexDirection: "column",
            }}
          >
            <Image
              src={
                lab.difficulty === 1
                  ? TestTubeGreen
                  : lab.difficulty === 2
                  ? TestTubeOrange
                  : TestTubeRed
              }
              alt="difficulty"
              width={40}
              height={40}
            />
            {/* difficulty   */}
            <Typography
              variant="h5"
              sx={{
                fontFamily: "Outfit",
                fontWeight: "bold",
                fontSize: "25px",
                letterSpacing: "0",
                textAlign: "center",
                marginTop: "20px",
              }}
            >
              {lab.language.title}
            </Typography>

            <Button
              variant={lab.finished ? "dark" : "light"}
              sx={{
                borderRadius: "16px",
                marginTop: "20px",
                width: "140px",
                height: "40px",
                textTransform: "none",
                fontSize: "15px",
                fontWeight: "bold",
                fontFamily: "Outfit",
                letterSpacing: "0",
              }}
              onClick={() => {
                // this router sistem will be changed to the following cuse this is not true.
                // this is the try to solve the problem of the router.
                // router.push(`/labs/${programingId}/${lab.id}`);
              router.push(`/labs/${programingId}/${lab.id}`);
              }
              }
            >
              {lab.isFinished ? (
                <Translations text={"lab.button.review"} />
              ) : (
                <Translations text={"lab.button.solve"} />
              )}
            </Button>
          </Box>
        </Card>
      ))}
    </Box>
  );
};

export default LabInfo;
