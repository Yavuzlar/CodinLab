import {
    Box,
    Card,
    CardContent,
    Typography,
    useMediaQuery
  } from "@mui/material";
  
  const Activity = ({ title = "", description = ""}) => {
    const _smd = useMediaQuery((theme) => theme.breakpoints.down("smd"));
  
    return (
      <Box
        sx={{
          display: "flex",
          width: "100%",
          height: "100%",
        }}
      >
        <Card
          sx={{
            width: "100%",
            minHeight: "12.5rem",
            display: "flex",
            flexDirection: "column",
            justifyContent: "space-between",
          }}
        >
          <CardContent sx={{ height: "calc(100% - 3rem)" }}>
            <Box
              sx={{
                display: "flex",
                flexDirection: "column",
                height: "100%",
                marginBottom: "1rem",
              }}
            >
              <Typography variant="title" sx={{ fontWeight: "bold" }}>
                {title}
              </Typography>
              <Typography
                sx={{
                  maxWidth: "calc(100% - 9.625rem)",
  
                  paddingTop: "0.8rem",
                  ...(_smd && { maxWidth: "60ch" }),
                }}
              >
                {description}
              </Typography>
            </Box>
          </CardContent>
        </Card>
      </Box>
    );
  };
  
  export default Activity;
  