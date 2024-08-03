import {
  Card,
  Typography,
  CardContent,
  useMediaQuery,
  Grid,
  Box,
  Button
} from "@mui/material";
import Image from "next/image";
import LinearProgess from "../progress/LinearProgess";
import labsIcon from "../../assets/icons/icons8-test-tube-100.png";

const LanguageProgress = ({ language }) => {
  const _md = useMediaQuery((theme) => theme.breakpoints.down("md"));

  return (
    <Card>
      <CardContent
        sx={{
          display: "flex",
          flexDirection: "row",
          justifyContent: "start",
          alignItems: "center",
          gap: "1.5rem",
        }}
      >
        <Image
          src={language.image}
          alt={language.name}
          width={60}
          height={60}
        />

        {
          language.progress == 0
            ?
            <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', width: '100%' }}>
              <Typography
                variant="h5"
                sx={{ fontWeight: 600, width: '120px' }}
              >
                {language.name}
              </Typography>

              <Box sx={{ display: 'flex', alignItems: 'center', gap: '1rem' }}>
                <Image src={labsIcon} alt={"labsIcon"} width={24} height={24} />

                <Typography variant="infoText">20 Lab</Typography>
              </Box>

              <Button variant="dark" sx={{ textTransform: 'none', width: '120px' }}>
                <Typography variant="infoText">Start</Typography>
              </Button>
            </Box>
            : <Grid container spacing={0}>
              <Grid item xs={12}>
                <Typography
                  variant="h5"
                  sx={{ fontWeight: 600 }}
                >
                  {language.name}
                </Typography>
              </Grid>

              <Grid item xs={12}>
                <LinearProgess progress={language.progress} />
              </Grid>

              <Grid item xs={12}>
                <Box sx={{ display: 'flex', alignItems: 'center', gap: '1rem' }}>
                  <Image src={labsIcon} alt={"labsIcon"} width={24} height={24} />

                  <Typography variant="infoText">20 / 40 Lab</Typography>
                </Box>
              </Grid>
            </Grid>
        }
      </CardContent>

    </Card>
  );
};

export default LanguageProgress;
