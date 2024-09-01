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
import Translations from "../Translations";

const LanguageProgress = ({ language, icon, map }) => {
  const _lg = useMediaQuery((theme) => theme.breakpoints.down("lg"));
  const _md = useMediaQuery((theme) => theme.breakpoints.down("md"));
  const _sm = useMediaQuery((theme) => theme.breakpoints.down("sm"));

  return (
    <Card>
      <CardContent
        sx={{
          display: "flex",
          flexDirection: "row",
          justifyContent: "start",
          alignItems: "center",
          gap: "1rem",
        }}
      >
        <img
          src={"api/v1/"+language.iconPath}
          alt={language.name}
          width={56}
          height={56}
        />

        {
          language.percentage == 0
            ? _lg && !_md
              ? <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', width: '100%' }}>
                <Box sx={{ display: 'flex', flexDirection: 'column', gap: '0.25rem' }}>
                  <Typography
                    variant="h5"
                    sx={{
                      width: '100%',
                      fontWeight: 600,
                      textOverflow: 'ellipsis',
                      overflow: 'hidden'
                    }}
                  >
                    {language.name}
                  </Typography>

                  <Box sx={{ display: 'flex', alignItems: 'center', gap: '0.5rem', minWidth: 'fit-content' }}>
                    <Image src={icon} alt={"icon"} width={20} height={20} />

                    <Typography variant="infoText2">{language.totalLabs}</Typography>
                  </Box>
                </Box>

                <Button variant="dark" sx={{ textTransform: 'none', minWidth: '80px' }}>
                  <Typography variant="infoText"><Translations text="start" /></Typography>
                </Button>
              </Box>
              : <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', width: '100%' }}>
                <Typography
                  variant="h5"
                  sx={{
                    maxWidth: 'calc(100% - 200px)',
                    width: '100%',
                    fontWeight: 600,
                    textOverflow: 'ellipsis',
                    overflow: 'hidden'
                  }}
                >
                  {language.name}
                </Typography>

                <Box sx={{ display: 'flex', gap: '1rem', width: '100%', justifyContent: 'end' }}>
                  <Box sx={{ display: 'flex', alignItems: 'center', gap: '0.5rem', minWidth: 'fit-content' }}>
                    <Image src={icon} alt={"icon"} width={24} height={24} />

                    <Typography variant="infoText">{map}</Typography>
                  </Box>

                  <Button variant="dark" sx={{ textTransform: 'none', minWidth: '80px' }}>
                    <Typography variant="infoText"><Translations text="start" /></Typography>
                  </Button>
                </Box>
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
                <LinearProgess progress={language.percentage} />
              </Grid>

              <Grid item xs={12}>
                <Box sx={{ display: 'flex', alignItems: 'center', gap: '0.5rem' }}>
                  <Image src={icon} alt={"icon"} width={24} height={24} />

                  <Typography variant="infoText">{map}</Typography>
                </Box>
              </Grid>
            </Grid>
        }
      </CardContent>
    </Card>
  );
};

export default LanguageProgress;
