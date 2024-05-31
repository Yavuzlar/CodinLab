import { useTheme } from "@emotion/react";
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
import Image from "next/image";
import { useState, useEffect } from "react";
import { registerValidation } from "src/configs/validation/registerSchema";
import CardImage from "src/assets/3d/3d-casual-life-windows-with-developer-code-symbols.png";
import GirlImage from "src/assets/3d/3d-casual-life-girl-holding-laptop-and-having-an-idea.png";

const Register = () => {
  const [formData, setFormData] = useState(null);
  const [errors, setErrors] = useState({});
  const [formSubmit, setFormSubmit] = useState(false);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async () => {
    setFormSubmit(true);
  };

  useEffect(() => {
    const fetchData = async () => {
      if (formData && formSubmit) {
        const errors = await registerValidation(formData);
        setErrors(errors);
      }
    };
    fetchData();
  }, [formData, formSubmit]);

  const theme = useTheme();
  const bgColor = theme.palette.primary.dark;

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
    <>
      <Box
        sx={{
          position: "relative",
        }}>
        <Box
          sx={{
            display: { xs: "none", mdlg: "block" },
            position: "absolute",
            top: "-6.5%",
            left: {
              mdlg: "%1",
              lg: "3%",
              lgPlus: "5%",
              lgXl: "9%",
              xl: "20%",
            }, // Adjust for overlap at top-left
            zIndex: 1, // Ensure image is above the form container
          }}>
          <Image src={CardImage} width={368} height={226} />
        </Box>
        <Box
          sx={{
            display: { xs: "none", mdlg: "block" },
            position: "absolute",
            top: "3%", // Adjust for overlap at top-right
            right: {
              mdlg: "-17%",
              lg: "-10%",
              lgPlus: "-5%",
              lgXl: "2%",
              xl: "14%",
            },
            zIndex: 1, // Ensure image is above the form container
          }}>
          <Image src={GirlImage} width={368} height={803} />
        </Box>
        <Container sx={{ display: "flex", justifyContent: "center", mt: "4%" }}>
          <Box
            component={"form"}
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
                <Grid container direction="column" gap={3}>
                  <TextField
                    name="fullname"
                    label="Fullname"
                    InputLabelProps={inputLabelStyle}
                    onChange={handleChange}
                    error={errors.fullname ? true : false}
                    helperText={errors.fullname}
                    fullWidth
                  />
                  <TextField
                    name="username"
                    label="Username"
                    InputLabelProps={inputLabelStyle}
                    onChange={handleChange}
                    error={errors.username ? true : false}
                    helperText={errors.username}
                    fullWidth
                  />
                  <TextField
                    name="email"
                    label="E-mail"
                    InputLabelProps={inputLabelStyle}
                    onChange={handleChange}
                    error={errors.email ? true : false}
                    helperText={errors.email}
                    fullWidth
                  />
                  <TextField
                    name="password"
                    label="Password"
                    InputLabelProps={inputLabelStyle}
                    onChange={handleChange}
                    error={errors.password ? true : false}
                    helperText={errors.password}
                    fullWidth
                  />
                  <FormControlLabel
                    control={
                      <Checkbox
                        sx={{ color: "#fff" }}
                        onChange={handleChange}
                        error={errors.Checkbox ? true : false}
                      />
                    }
                    label={
                      <Typography
                        fontWeight={300}
                        fontSize={18}
                        fontFamily={"Outfit"}>
                        I accept the{" "}
                        <Link
                          sx={{ textDecoration: "none", fontWeight: "600" }}
                          color={"#0A3B7A"}
                          href="#">
                          privacy policy & terms
                        </Link>
                      </Typography>
                    }
                  />
                  <Button
                    sx={{
                      bgcolor: bgColor,
                      font: "normal normal 18px/23px Outfit",
                      fontWeight: "600",
                      textTransform: "capitalize",
                      py: 2,
                    }}
                    onClick={handleSubmit}
                    fullWidth>
                    Sign Up
                  </Button>
                </Grid>
              </FormControl>
              <Divider sx={{ mt: 3 }}>or</Divider>
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
                  color={bgColor}
                  sx={{ fontWeight: "600", textDecoration: "none", ml: 1 }}>
                  Login
                </Link>
              </Typography>
            </Grid>
          </Box>
        </Container>
      </Box>
    </>
  );
};

export default Register;
