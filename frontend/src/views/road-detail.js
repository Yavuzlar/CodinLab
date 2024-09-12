import { useTheme } from "@mui/material/styles";
import CustomBreadcrumbs from "src/components/breadcrumbs";
import { useTranslation } from "react-i18next";
import i18next from "i18next";
import { Box, Button, Card, CardContent, Grid, Typography, Stack } from "@mui/material";
import CircleIcon from '@mui/icons-material/Circle';
import LockIcon from "src/assets/icons/padlock.png"
import PathIcon from "src/assets/icons/icons8-path-100.png"
import CIcon from "src/assets/icons/c.png"
import DoneIcon from "src/assets/icons/icons8-done-100 (1).png"
import Image from "next/image";
import { CircularProgressStatistics } from "src/components/progress/CircularProgressStatistics";
import { useEffect, useEffect, useState } from "react";
import SkeletonLoader from "src/components/skeleton/SkeletonLoader.js";
import LinearProgess from "src/components/progress/LinearProgess";
import { useRouter } from "next/router";
import { useDispatch, useSelector } from "react-redux";
import { fetchPaths, startRoad } from "src/store/paths/pathsSlice";
import { getProgrammingId } from "src/data/programmingIds";
import { useDispatch, useSelector } from "react-redux";
import { getUserRoadProgressStats } from "src/store/statistics/statisticsSlice";


const RoadDetails = ({ language = "" }) => {

    const capitalizedLanguage = language.charAt(0).toUpperCase() + language.slice(1);
    
    const [programmingId, setProgrammingId] = useState(null)

    const theme = useTheme();
    const { t, i18n } = useTranslation();
    const router = useRouter();

    const dispatch = useDispatch();
    const { paths } = useSelector((state) => state);

    const [pathsDataContent, setPathsDataContent] = useState([])

    const [isStarted, setIsStarted] = useState(false) // Set this to true if the user has started the road on useEffect()

    const [amountOfInProgressPaths, setAmountOfInProgressPaths] = useState(0) // Amount of in progress paths

    const [amountOfCompletedPaths, setAmountOfCompletedPaths] = useState(0) // Amount of completed paths

    const [isLoading, setIsLoading] = useState(true) // Loading state for fetching the paths
    const [error, setError] = useState(null) // Error state for fetching the paths

    // TODO: Get the title and description from front-end side
    const title = "What is C?"
    const description = "C is a programming language created by Dennis Ritchie at Bell Laboratories in 1972. It is a popular language due to its foundational nature and close association with UNIX."

    const { statistics: stateStatistics } = useSelector(
      (state) => state
    );
  
    useEffect(() => {
        dispatch(getUserRoadProgressStats());
    }, [dispatch]);

    const handleStartRoad = () => {
        // Redirect to the first path of the road
        dispatch(startRoad({ programmingid: programmingId }))
        router.push(`/roads/${language}/1`)
    }

    useEffect(() => {
        console.log("Language useEffect: ", language)
        setProgrammingId(getProgrammingId[language])
    }, [language])

    useEffect(() => {
        // Fetch the paths of the road
        if (programmingId) {
            dispatch(fetchPaths({ programmingid: programmingId }))
        }
    }, [programmingId])

    useEffect(() => {
        
        if (paths) {
        console.log("Paths fetched: ", paths)
        
        setIsLoading(paths.loading)
        setError(paths.error)

        if (paths.data.paths) {

            console.log("Raw paths.data.paths: ", paths.data.paths)

            const pathsData = paths.data.paths
            
            console.log("Paths data: ", pathsData)

            // Amount of completed paths
            const completedPaths = pathsData.filter(path => path.isFinished)

            // Amount of in progress paths
            const inProgressPaths = pathsData.filter(path => !path.isFinished && path.isStarted)
            
            if (completedPaths.length > 0) {
                setIsStarted(true)
            }

            setAmountOfInProgressPaths(inProgressPaths.length)
            setAmountOfCompletedPaths(completedPaths.length)

            const pathContent = pathsData.map(item => ({
                ...item,
                languages: item.languages.filter(langItem => langItem.lang === i18n.language)
            }));

            console.log("Paths data content: ", pathContent)

            setPathsDataContent(pathContent)
        }
        }

    }, [paths, i18next.language])

    // Breadcrumbs
    const breadcrums = [
    {
        path: "/roads",
        title: t("home.roads.title"),
        permission: "roads"
        },
        {
        path: `/roads/${language}`,
        title: capitalizedLanguage,
        permission: "roads"
        },
    ]

    const progresses = [
  {
    name: "In progress", // String
    value: stateStatistics.data?.data?.progress, // Number
    color: "#8FDDFD" // String
  },
  {
    name: "Completed", // String
    value: stateStatistics.data?.data?.completed, // Number
    color: "#0A3B7A" // String
  }
    ]

    return (
        <Box>
            {/* Breadcrumbs */}
            <CustomBreadcrumbs titles={breadcrums} />
            
            {/* Header Cards */}
            <Box sx={{mt: 2}}>
                <Grid container spacing={2}>
                    {/* Road Description and button */}
                    <Grid item xs={12} sm={6} md={8}>
                        <Card sx={{ height: "100%" }}>
                            <CardContent sx={{display: "flex", justifyContent: "space-between", alignItems: "center", gap: 3, p: 4}}>
                                <Image src={CIcon} alt="C Icon" width={80} height={80} />
                                { !isStarted ? (
                                    <>
                                    <Box>
                                    <Typography variant="h4" fontWeight={600}>{title}</Typography>
                                    <Typography variant="body1">{description}</Typography>
                                </Box>
                                <Button
                                variant="contained"
                                sx={{
                                    backgroundColor: "#fff", 
                                    maxWidth: '9.37rem', 
                                    maxHeight: '3.12rem', 
                                    minWidth: '9.37rem', 
                                    minHeight: '3.12rem',
                                    ':hover': {
                                        bgcolor: theme.palette.primary.light,
                                        },
                                    }}
                                onClick={handleStartRoad}
                                    >
                                    <Typography 
                                    fontWeight={600}
                                    variant="body1" 
                                    color={theme.palette.primary.dark} 
                                    sx={{textTransform: "capitalize"}}> {t("roads.path.start_road")} </Typography>
                                </Button>
                                </>
                                ) :
                                <Box sx={{ display: "flex", flexDirection: "column", width: "100%", gap: 3 }}>
                                    <Typography variant="h4"> {capitalizedLanguage} </Typography>
                                    <LinearProgess progress={amountOfCompletedPaths} />
                                    <Stack direction={"row"} spacing={1}>
                                        <Image src={PathIcon} alt="Path Icon" width={25} height={25} />
                                        <Typography variant="body1">{amountOfCompletedPaths}/100 Path</Typography>
                                    </Stack>
                                </Box>
                                }
                            </CardContent>
                        </Card>
                    </Grid>

                    {/* Circular Progresses */}
                    <Grid item xs={12} sm={6} md={4}>    
                        <Card sx={{ height: "100%" }}>
                            <CardContent sx={{display: "flex", justifyContent: "space-around", alignItems: "center"}}>
                                <CircularProgressStatistics progresses={progresses} />
                            </CardContent>
                        </Card>
                    </Grid>

                </Grid>
            </Box>

            {/* Road Paths */}
            {pathsDataContent.map((path, index) => (
                <Box key={index}>
                <Box sx={{
                borderWidth: 6,
                borderColor: path.isFinished ? "#39CE19" : theme.palette.primary.dark,
                borderStyle: index % 2 === 0 ? "none none dashed dashed" : "none dashed dashed none",
                p: 3
            }}>
                <Box sx={{
                    mt: 2,
                    display: "flex",
                    gap: 2,
                    alignItems: "center",
                    border: path.isFinished ? "3px solid #39CE19" : "none",
                    borderRadius: 6,
                    backgroundColor: path.isFinished ? "#fff" : theme.palette.primary.dark,
                    p: 3,
                }}>
                    {path.isFinished ? 
                    <Image src={DoneIcon} alt="Done Icon" width={30} height={30} /> 
                    : 
                    <Image src={LockIcon} alt="Next Path Icon" width={30} height={30} />}

                    <Typography variant="body1" fontWeight={600} color={!path.isFinished ? "#fff" : "#0A3B7A"}> {path.languages[0].title}    
                    : </Typography>
                    <Typography variant="body1" color={!path.isFinished ? "#fff" : "#0A3B7A"}> {path.languages[0].description}</Typography>
                </Box>
            </Box>
                </Box>
            ))}
        </Box>
    );
}
 
export default RoadDetails;