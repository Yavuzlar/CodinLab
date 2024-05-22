import { Box, Card, CardActions, CardContent, Typography } from "@mui/material";

const LevelStatistic = ({ levels, progress }) => {
  const height = progress + "%";

  return (
    <Card
      style={{
        width: "529px",
        height: "319px",
      }}
    >
      <CardContent
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "center",
          fontSize: "36px",
          fontWeight: "bold",
          textAlign: "center",
        }}
      >
        Level
      </CardContent>
      <CardActions>
        <CardContent
          style={{
            display: "flex",
            width: "91%",
            height: "100%",
            alignItems: "center",
            marginTop: "10px",
            justifyContent: "center",
            alignItems: "center",
          }}
        >
          <Box
            display="flex"
            justifyContent="center"
            alignItems="center"
            sx={{
              position: "relative",
              flexDirection: "row",
            }}
          >
            <Box
              display={"flex"}
              justifyContent="center"
              alignItems="center"
              width={"39px"}
              height={"116px"}
            >
              <Box
                display="flex"
                alignItems="center"
                sx={{ position: "relative", flexDirection: "row" }}
              >
                <Box
                  sx={{
                    display: "flex",
                    flexDirection: "row",
                    alignItems: "center",
                    justifyContent: "center",
                    width : "65px",
                    height: "100%",
                    marginRight: "27px",
                  }}
                >
                <Typography
                  sx={{
                    fontSize: "30px",
                    marginTop: "28px",
                    marginRight: "8px",
                  }}
                >
                  Level 
                  
                </Typography>
                  <Typography
                    sx={{
                      fontSize: "30px",
                      marginTop: "28px",
                    }}
                  >
                    {levels}
                  </Typography>
                </Box>
              </Box>
              <Box
                sx={{
                  display: "flex",
                  flexDirection: "row",
                  alignItems: "center",
                  justifyContent: "center",
                  width: "100%",
                  height: "100%",
                  borderRadius: "7px",
                  marginLeft: "80px",
                  marginRight: "60px",
                }}
              >
                <Typography
                  sx={{
                    fontSize: "200px",
                    color: "#0A3B7A",
                    fontFamily: "Nunito",
                  }}
                >
                  &#123;
                </Typography>
                <Box
                  sx={{
                    width: "100%",
                    height: "100%",
                    borderRadius: "7px",
                    backgroundColor: "#fff",
                    position: "absolute",
                    marginTop: "30px",
                  }}
                >
                  <Box
                    sx={{
                      width: "100%",
                      height: height,
                      bottom: "0",
                      borderRadius: "7px",
                      backgroundColor: "#8FDDFD",
                      position: "absolute",
                      maxHeight: "100% !important",
                    }}
                  >
                    {/* Mavi kutu */}
                  </Box>
                </Box>
                <Typography
                  sx={{
                    fontSize: "200px",
                    color: "#0A3B7A",
                    fontFamily: "Nunito",
                    marginLeft: "10px",
                  }}
                >
                  &#125;
                </Typography>
              </Box>

              <Typography
                sx={{ fontSize: "30px", marginTop: "28px", marginLeft: "20px", width : "65px" }}
              >
                %{progress}
              </Typography>
            </Box>
          </Box>
        </CardContent>
      </CardActions>
    </Card>
  );
};

export default LevelStatistic;
