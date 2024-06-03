import { useTheme } from "@emotion/react";
import {
  Circle,
  Google,
  GitHub,
  Visibility,
  VisibilityOff,
} from "@mui/icons-material";
import {
  Box,
  Checkbox,
  Container,
  FormControl,
  FormControlLabel,
  Grid,
  InputAdornment,
  Card,
  CardContent,
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
import Translations from "src/components/Translations";

const Register = () => {
  const [formData, setFormData] = useState(null);
  const [errors, setErrors] = useState({});
  const [formSubmit, setFormSubmit] = useState(false);

  const [showPassword, setShowPassword] = useState(false);
  const handleClickShowPassword = () => setShowPassword(!showPassword);

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
              xl: "12%",
              xxl: "16%",
              xxxl: "18%",
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
              xl: "4%",
              xxl: "8%",
              xxxl: "14%",
            },
            zIndex: 1, // Ensure image is above the form container
          }}>
          <Image src={GirlImage} width={368} height={803} />
        </Box>
        <Container sx={{ display: "flex", justifyContent: "center", mt: "4%" }}>
          <Card>
            <CardContent
              sx={{
                width: "812px",
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
                      label={<Translations text="register.fullname" />}
                      InputLabelProps={inputLabelStyle}
                      onChange={handleChange}
                      error={errors.fullname ? true : false}
                      helperText={errors.fullname}
                    />
                    <TextField
                      name="username"
                      label={<Translations text="register.username" />}
                      InputLabelProps={inputLabelStyle}
                      onChange={handleChange}
                      error={errors.username ? true : false}
                      helperText={errors.username}
                    />
                    <TextField
                      name="email"
                      label={<Translations text="register.email" />}
                      InputLabelProps={inputLabelStyle}
                      type="email"
                      onChange={handleChange}
                      error={errors.email ? true : false}
                      helperText={errors.email}
                    />
                    <TextField
                      name="password"
                      label={<Translations text="register.password" />}
                      InputLabelProps={inputLabelStyle}
                      type={showPassword ? "text" : "password"}
                      onChange={handleChange}
                      error={errors.password ? true : false}
                      helperText={errors.password}
                      InputProps={{
                        endAdornment: (
                          <InputAdornment position="end">
                            <IconButton
                              aria-label="toggle password visibility"
                              onClick={handleClickShowPassword}>
                              {showPassword ? (
                                <VisibilityOff />
                              ) : (
                                <Visibility />
                              )}
                            </IconButton>
                          </InputAdornment>
                        ),
                      }}
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
                          {<Translations text="register.accept" />}
                          <Link
                            sx={{ textDecoration: "none", fontWeight: "600" }}
                            color={"#0A3B7A"}
                            href="#">
                            {<Translations text="register.terms" />}
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
                      {<Translations text="register.signup" />}
                    </Button>
                  </Grid>
                </FormControl>
                <Divider sx={{ mt: 3 }}>
                  {" "}
                  {<Translations text={"register.or"} />}{" "}
                </Divider>
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
                  {<Translations text="register.already" />}
                  <Link
                    href="#"
                    color={bgColor}
                    sx={{ fontWeight: "600", textDecoration: "none", ml: 1 }}>
                    {<Translations text="register.login" />}
                  </Link>
                </Typography>
              </Grid>
            </CardContent>
          </Card>
        </Container>
      </Box>
    </>
  );
};

export default Register;
