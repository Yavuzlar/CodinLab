import React from "react";
import roadsImage from "../../assets/3d/3d-casual-life-young-women-working-with-computer.png";
import { roads } from "src/data/home";
import {
  Box,
  Card,
  CardContent,
  Typography,
  useMediaQuery,
} from "@mui/material";
import Image from "next/image";

const Info = ({ title = "", description = "", image = null }) => {
  const _smd = useMediaQuery((theme) => theme.breakpoints.down("smd"));

  return (
    <Box
      sx={{
        display: "flex",
        position: "relative",
        width: "100%",
        height: "100%",
      }}
    >
      <Card
        sx={{
          width: "100%",
          minHeight: "200px",
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
                maxWidth: "calc(100% - 154px)",

                paddingTop: "13px",
                ...(_smd && { maxWidth: "60ch" }),
              }}
            >
              {description}
            </Typography>
          </Box>
        </CardContent>
      </Card>
      {image && (
        <Image
          style={{
            position: "absolute",
            right: "0.5rem",
            bottom: "1.5rem",
            opacity: _smd ? "0.25" : "1",
          }}
          src={image}
          alt="Welcome"
          width={154}
          height={152}
        />
      )}
    </Box>
  );
};

export default Info;
