import { useTheme } from "@mui/material/styles";
import CustomBreadcrumbs from "src/components/breadcrumbs";
import { useTranslation } from "react-i18next";
import i18next from "i18next";
import {
  Box,
  Button,
  Card,
  CardContent,
  Grid,
  Typography,
  Stack,
  CircularProgress,
} from "@mui/material";
import LockIcon from "src/assets/icons/padlock.png";
import PathIcon from "src/assets/icons/icons8-path-100.png";
import DoneIcon from "src/assets/icons/icons8-done-100 (1).png";
import NextPathIcon from "src/assets/icons/icons8-signpost-100.png";
import Image from "next/image";
import { CircularProgressStatistics } from "src/components/progress/CircularProgressStatistics";
import { useContext, useEffect, useState } from "react";
import LinearProgess from "src/components/progress/LinearProgess";
import { useRouter } from "next/router";
import { useDispatch, useSelector } from "react-redux";
import { fetchPaths, startRoad } from "src/store/paths/pathsSlice";
import { showToast } from "src/utils/showToast";
import Translations from "src/components/Translations";
import { AuthContext } from "src/context/AuthContext";

const RoadDetails = ({ language = "" }) => {
  const theme = useTheme();
  const { t, i18n } = useTranslation();
  const router = useRouter();
  const dispatch = useDispatch();
  const { paths } = useSelector((state) => state);

  const [isLoading, setIsLoading] = useState(true); // Loading state for fetching the paths
  const [error, setError] = useState(null); // Error state for fetching the paths

  const [programmingId, setProgrammingId] = useState(null);

  const [languageName, setLanguageName] = useState("");
  const [pathsDataContent, setPathsDataContent] = useState([]);
  const [roadIsStarted, setRoadIsStarted] = useState(false);
  const [totalPath, setTotalPath] = useState(0);
  const [amountOfInProgressPaths, setAmountOfInProgressPaths] = useState(0); // Amount of in progress paths
  const [amountOfCompletedPaths, setAmountOfCompletedPaths] = useState(0); // Amount of completed paths // Path icon path
  const [programmingIcon, setProgrammingIcon] = useState("images/c.png"); // Programming icon path
  const [title, setTitle] = useState(""); // Road title
  const [description, setDescription] = useState(""); // Road description
  const [isImageExists, setIsImageExists] = useState(false);
  const { containerLoading } = useContext(AuthContext);

  const breadcrums = [
    {
      path: "/roads",
      title: t("home.roads.title"),
      permission: "roads",
    },
    {
      path: `/roads/${language}`,
      title: languageName,
      permission: "roads",
    },
  ];

  const progresses = [
    {
      name: t("in_progress"), // String
      // value: stateStatistics.data?.data?.progress, // Number
      value: (amountOfInProgressPaths * 100) / totalPath,
      color: theme.palette.primary.light, // String
    },
    {
      name: t("completed"), // String
      // value: stateStatistics.data?.data?.completed, // Number
      value: (amountOfCompletedPaths * 100) / totalPath,
      color: theme.palette.primary.dark, // String
    },
  ];

  useEffect(() => {
    setProgrammingId(parseInt(language));
  }, [language]);

  useEffect(() => {
    // Fetch the paths of the road
    if (programmingId) {
      dispatch(
        fetchPaths({ programmingid: programmingId, language: i18n.language })
      );
    }
  }, [programmingId, i18n.language]);

  useEffect(() => {
    if (paths) {
      setIsLoading(paths.loading);
      setError(paths.error);

      if (paths.data.paths) {
        setProgrammingIcon(paths.data.iconPath);

        setTitle(paths.data.name);
        setDescription(paths.data.description);
        setRoadIsStarted(paths.data.roadIsStarted);

        const pathsData = paths.data.paths;

        // Amount of completed paths
        const completedPaths = pathsData.filter((path) => path.pathIsFinished);

        // Amount of in progress paths
        const inProgressPaths = pathsData.filter(
          (path) => !path.pathIsFinished && path.pathIsFinished
        );

        setTotalPath(pathsData.length);
        setAmountOfInProgressPaths(inProgressPaths.length);
        setAmountOfCompletedPaths(completedPaths.length);
        setPathsDataContent(pathsData);
        setLanguageName(paths.data.name);
      }
    }
  }, [paths, i18next.language]);

  useEffect(() => {
    if (paths.data?.isImageExists) {
      setIsImageExists(true);
    }
  }, [paths.data?.isImageExists]);

  const renderPathIcon = (path) => {
    if (path.pathIsFinished) {
      return DoneIcon;
    } else if (path.pathIsStarted && !path.pathIsFinished) {
      return NextPathIcon;
    } else {
      return LockIcon;
    }
  };

  const isImageExist = paths.data?.isImageExists;
  const handleStartRoad = () => {
    if (isImageExist) {
      dispatch(startRoad({ programmingID: programmingId }));
      router.push(`/roads/${language}/1`);
    } else {
      showToast("error", "Image not found");
    }
  };

  const handlePath = (path) => {
    if (!(path.pathIsStarted && !path.pathIsFinished) && !path.pathIsFinished) {
      return;
    }

    router.push(`/roads/${language}/${path.id}`);
  };

  return (
    <Box>
      {/* Breadcrumbs */}
      <CustomBreadcrumbs titles={breadcrums} />

      {/* Header Cards */}
      <Box sx={{ mt: 2 }}>
        <Grid container spacing={2}>
          {/* Road Description and button */}
          <Grid item xs={12} md={8}>
            <Card sx={{ height: "100%" }}>
              <CardContent
                sx={{
                  display: "flex",
                  justifyContent: "start",
                  alignItems: "start",
                  gap: 3,
                  p: 4,
                }}
              >
                <Image
                  src={`/${programmingIcon}`}
                  alt="Programming Language Icon"
                  width={80}
                  height={80}
                  style={{}}
                />
                {!roadIsStarted ? (
                  <>
                    <Box
                      sx={{
                        display: "flex",
                        flexDirection: "column",
                        gap: 2,
                      }}
                    >
                      <Typography variant="h4" fontWeight={600}>
                        {title}
                      </Typography>
                      <Typography variant="body1">{description}</Typography>
                      <Button
                        variant="contained"
                        disabled={!isImageExists}
                        sx={{
                          backgroundColor: "#fff",
                          color: theme.palette.primary.dark,
                          fontWeight: 600,
                          maxWidth: "9.37rem",
                          maxHeight: "3.12rem",
                          minWidth: "9.37rem",
                          minHeight: "3.12rem",
                          ":hover": {
                            bgcolor: theme.palette.primary.light,
                          },
                        }}
                        onClick={handleStartRoad}
                      >
                        {containerLoading ? (
                          <CircularProgress
                            size={24}
                            sx={{ position: "absolute" }}
                          />
                        ) : (
                          <Translations text={"roads.path.start_road"} />
                        )}
                      </Button>
                    </Box>
                  </>
                ) : (
                  <Box
                    sx={{
                      display: "flex",
                      flexDirection: "column",
                      width: "100%",
                      gap: 3,
                    }}
                  >
                    <Typography variant="h4">{title}</Typography>
                    <Typography variant="body1">{description}</Typography>
                    <LinearProgess
                      progress={(amountOfCompletedPaths * 100) / totalPath}
                    />
                    <Stack direction={"row"} spacing={1}>
                      <Image
                        src={PathIcon}
                        alt="Path Icon"
                        width={25}
                        height={25}
                      />
                      <Typography variant="body1">
                        {amountOfCompletedPaths}/{totalPath} Path
                      </Typography>
                    </Stack>
                  </Box>
                )}
              </CardContent>
            </Card>
          </Grid>

          {/* Circular Progresses */}
          <Grid item xs={12} md={4}>
            <Card sx={{ height: "100%" }}>
              <CardContent
                sx={{
                  display: "flex",
                  justifyContent: "space-around",
                  alignItems: "center",
                }}
              >
                <CircularProgressStatistics
                  progresses={progresses}
                  flexDirection={"column"}
                />
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </Box>

      {pathsDataContent.map((path, index) => (
        <Box key={index}>
          <Box
            sx={{
              borderWidth: 6,
              borderColor: path.pathIsFinished
                ? "#39CE19"
                : theme.palette.primary.dark,
              borderStyle:
                index % 2 === 0
                  ? "none none dashed dashed"
                  : "none dashed dashed none",
              p: 3,
            }}
          >
            <Box
              onClick={() => handlePath(path)}
              sx={{
                mt: 2,
                display: "flex",
                gap: 2,
                alignItems: "center",
                border: path.pathIsFinished ? "3px solid #39CE19" : "none",
                borderRadius: 6,
                backgroundColor: path.pathIsFinished
                  ? "#fff"
                  : theme.palette.primary.dark,
                p: 3,
                cursor: path.pathIsStarted ? "pointer" : "not-allowed",
                "&:hover": {
                  boxShadow: 5,
                },
              }}
            >
              <Image
                src={renderPathIcon(path)}
                alt="Done Icon"
                width={30}
                height={30}
              />

              <Typography
                variant="body1"
                fontWeight={600}
                color={!path.pathIsFinished ? "#fff" : "#0A3B7A"}
              >
                {path.language.title}:
              </Typography>
              <Typography
                variant="body1"
                color={!path.pathIsFinished ? "#fff" : "#0A3B7A"}
              >
                {path.language.description}
              </Typography>
            </Box>
          </Box>
        </Box>
      ))}
    </Box>
  );
};

export default RoadDetails;
