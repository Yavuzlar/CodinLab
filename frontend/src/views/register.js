import { Circle, Google, GitHub } from "@mui/icons-material";
import {
  Box,
  Checkbox,
  Container,
  FormControl,
  FormControlLabel,
  Grid,
  TextField,
  Typography,
  Link,
  Button,
  Divider,
  Stack,
  IconButton,
} from "@mui/material";
import { useState, useEffect } from "react";

const Register = () => {
  const [formData, setFormData] = useState(null);
  const [errors, setErrors] = useState({});
  const [formSubmit, setFormSubmit] = useState(false);

  const iconSize = {
    width: 30,
    height: 30,
  };

  const iconBtnStyle = {
    bgcolor: "#0A3B7A",
    color: "#fff",
    borderRadius: 4,
  };

  const inputLabelStyle = {
    sx: {
      color: "#0A3B7A",
      font: "normal normal bold 18px/23px Outfit",
      ml: 1,
    },
  };

  return (
    <Container sx={{ display: "flex", justifyContent: "center" }}>
      <Box
        sx={{
          bgcolor: "#3894D0",
          borderRadius: "16px",
          width: "812px",
          height: "800px",
        }}>
        <Grid container direction="column" px={10}>
          <Box
            sx={{
              display: "flex",
              justifyContent: "center",
              alignItems: "center",
              gap: 1,
              my: 5,
            }}>
            <Circle sx={{ width: 40, height: 40, mr: 1 }} />
            <Typography
              textAlign="center"
              variant="body1"
              fontFamily="Outfit"
              fontWeight="600"
              fontSize="35px">
              CodinLab
            </Typography>
          </Box>
          <FormControl>
            <Grid container direction="column" gap={2}>
              <TextField
                name="fullname"
                label="Fullname"
                InputLabelProps={inputLabelStyle}
                fullWidth
              />
              <TextField
                name="username"
                label="Username"
                InputLabelProps={inputLabelStyle}
                fullWidth
              />
              <TextField
                name="email"
                label="E-mail"
                InputLabelProps={inputLabelStyle}
                fullWidth
              />
              <TextField
                name="password"
                label="Password"
                InputLabelProps={inputLabelStyle}
                fullWidth
              />
              <FormControlLabel
                control={<Checkbox sx={{ color: "#fff" }} />}
                label={
                  <Typography fontSize={18} fontFamily={"Outfit"}>
                    I accept the{" "}
                    <Link color={"#fffddd"} href="#">
                      privacy policy & terms
                    </Link>
                  </Typography>
                }
              />
              <Button
                sx={{
                  bgcolor: "#0A3B7A",
                  font: "normal normal 18px/23px Outfit",
                  fontWeight: "600",
                  textTransform: "capitalize",
                  py: 2,
                }}
                fullWidth>
                Sign Up
              </Button>
            </Grid>
          </FormControl>
          <Divider sx={{ mt: 2 }}>or</Divider>
          <Stack direction="row" justifyContent="center" gap={3} mt={3}>
            <IconButton variant="contained" sx={iconBtnStyle}>
              <Google sx={iconSize} />
            </IconButton>
            <IconButton variant="contained" sx={iconBtnStyle}>
              <GitHub sx={iconSize} />
            </IconButton>
          </Stack>
          <Typography
            variant="body1"
            textAlign={"center"}
            mt={4}
            fontFamily={"Outfit"}>
            Already have an account?{" "}
            <Link
              href="#"
              color="#0A3B7A"
              sx={{ fontWeight: "600", textDecoration: "none", ml: 1 }}>
              Login
            </Link>
          </Typography>
        </Grid>
      </Box>
    </Container>
  );
};

export default Register;
