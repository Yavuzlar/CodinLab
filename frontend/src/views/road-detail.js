import { useTheme } from "@mui/material/styles";
import CustomBreadcrumbs from "src/components/breadcrumbs";
import { useTranslation } from "react-i18next";
import { Box, Button, Card, CardContent, Grid, Typography, Stack } from "@mui/material";
import CircleIcon from '@mui/icons-material/Circle';
import LockIcon from "src/assets/icons/padlock.png"
import PathIcon from "src/assets/icons/icons8-path-100.png"
import CIcon from "src/assets/icons/c.png"
import DoneIcon from "src/assets/icons/icons8-done-100 (1).png"
import Image from "next/image";
import { CircularProgressStatistics } from "src/components/progress/CircularProgressStatistics";
import { useEffect, useState } from "react";
import LinearProgess from "src/components/progress/LinearProgess";
import { useRouter } from "next/router";
import { useDispatch, useSelector } from "react-redux";
import { getUserRoadProgressStats } from "src/store/statistics/statisticsSlice";


const RoadDetails = ({ language = "" }) => {

    const capitalizedLanguage = language.charAt(0).toUpperCase() + language.slice(1);

    const theme = useTheme();
    const { t } = useTranslation();
    const router = useRouter();

    const [isStarted, setIsStarted] = useState(false) // Set this to true if the user has started the road on useEffect()

    const title = "What is C?"
    const description = "C is a programming language created by Dennis Ritchie at Bell Laboratories in 1972. It is a popular language due to its foundational nature and close association with UNIX."
    const amountOfCompletedPaths = 1;

    const dispatch = useDispatch();
    const { statistics: stateStatistics } = useSelector(
      (state) => state
    );
  
    useEffect(() => {
        dispatch(getUserRoadProgressStats());
    }, [dispatch]);

    const handleStartRoad = () => {
        // Redirect to the first path of the road
        router.push(`/roads/${language}/1`)
    }

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

    const roads = [
        {
            title: "Basic C Syntax",
            description: "Covers the basic syntax and language structure of C, including variables, data types, operators, and the usage of basic expressions.",
            completed: false,
        },
        {
            title: "Control Structures",
            description: "Focuses on control structures in C, including conditional statements, loops (for, while, do-while), and decision structures (if-else, switch-case).",
            completed: false
        },
        {
            title: "Functions",
            description: "Covers the definition, invocation, parameter usage, and return values of functions in C. Emphasizes the importance of functions in modular programming and code reusability.",
            completed: false
        },
        {
            title: "Arrays and Pointers",
            description: "Discusses the declaration, access, and manipulation of arrays in C. Also covers the usage of pointers, memory management, and manipulation of memory addresses.",
            completed: false
        },
    ];

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
                                <Typography variant="h4"> {language} </Typography>
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
                            <Box sx={{display: "flex", flexDirection: "column", gap: 2}}>
                                {progresses.map((progress, index) => (
                                <Box sx={{display: "flex", gap: 2, alignItems: "center"}} key={index}>
                                    <CircleIcon sx={{color: progress.color}} />
                                    <Typography variant="body1">{progress.name}</Typography>
                                    <Typography variant="body1">%{progress.value}</Typography>
                                </Box>
                                ))}
                            </Box>
                        </CardContent>
                    </Card>
                </Grid>

            </Grid>
        </Box>

        {/* Road Paths */}
        {roads.map((road, index) => (
            <Box key={index}>
            <Box sx={{
            borderWidth: 6,
            borderColor: road.completed ? "#39CE19" : theme.palette.primary.dark,
            borderStyle: index % 2 === 0 ? "none none dashed dashed" : "none dashed dashed none",
            p: 3
        }}>
            <Box sx={{
                mt: 2,
                display: "flex",
                gap: 2,
                alignItems: "center",
                border: road.completed ? "3px solid #39CE19" : "none",
                borderRadius: 6,
                backgroundColor: road.completed ? "#fff" : theme.palette.primary.dark,
                p: 3,
            }}>
                {road.completed ? 
                <Image src={DoneIcon} alt="Done Icon" width={30} height={30} /> 
                : 
                <Image src={LockIcon} alt="Next Path Icon" width={30} height={30} />}

                <Typography variant="body1" fontWeight={600} color={!road.completed ? "#fff" : "#0A3B7A"}> {road.title} : </Typography>
                <Typography variant="body1" color={!road.completed ? "#fff" : "#0A3B7A"}> {road.description}</Typography>
            </Box>
        </Box>
            </Box>
        ))}
    </Box>
    );
}
 
export default RoadDetails;