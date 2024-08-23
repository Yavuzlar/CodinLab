import {
  FormControl,
  Button,
  TextField,
  Card,
  CardContent,
  Typography,
  Box,
  Checkbox,
  FormGroup,
  FormControlLabel,
  Divider,
  IconButton,
  InputLabel,
  FilledInput,
  InputAdornment,
  OutlinedInput,
  CardMedia,
  Grid,
  useMediaQuery,
} from "@mui/material";
import { useState, useEffect } from "react";
import Translations from "src/components/Translations";
import { loginValidation } from "src/configs/validation/loginSchema";
import rocketImg from "../assets/3d/3d-casual-life-space-white-starship.png";
import googleIcon from "../assets/icons/icons8-google-100.png";
import githubIcon from "../assets/icons/icons8-github-144.png";
import visibilityOnIcon from "../assets/icons/icons8-eye-1.png";
import visibilityOffIcon from "../assets/icons/icons8-eye-1.png";
import Image from "next/image";
import manImg from "../assets/3d/3d-casual-life-young-man-sitting-with-laptop-and-waving.png";
import { useTranslation } from "next-i18next";
import themeConfig from "src/configs/themeConfig";
import { useAuth } from "src/hooks/useAuth";
const { default: BlankLayout } = require("src/layout/BlankLayout");

const Login = () => {
  const [username, setUsername] = useState();
  const [password, setPassword] = useState();
  const [showPassword, setShowPassword] = useState(false);
  const [formData, setFormData] = useState();
  const [errors, setErrors] = useState({});
  const [formSubmit, setFormSubmit] = useState(false);
  const [visibleUsernameLabel, setVisibleUsernameLabel] = useState(false);
  const [visiblePasswordLabel, setVisiblePasswordLabel] = useState(false);
  const { login } = useAuth();


  const handleClickShowPassword = () => setShowPassword(!showPassword);
  // const handleMouseDownPassword = (event) => {
  //   event.preventDefault();
  // };
  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });

    if (e.target.name === "username") {
      setUsername(e.target.value);
      handleVisibleUsernameLabel(e.target.value);
    }

    if (e.target.name === "password") {
      setPassword(e.target.value);
      handleVisiblePasswordLabel(e.target.value);
    }
  };

  const handleVisibleUsernameLabel = (username) => {
    setVisibleUsernameLabel(username !== "");
  };

  const handleVisiblePasswordLabel = (password) => {
    setVisiblePasswordLabel(password !== "");
  };

  const handleSubmit = async () => {
    setFormSubmit(true); 
    const validationErrors = await loginValidation(formData);
    setErrors(validationErrors);

    if (Object.keys(validationErrors).length > 0) {
      console.log("Form has errors:", validationErrors);
      return;
    }
    // Call API
    try {
      console.log(formData);
      await login(formData);
    } catch (error) {
      console.log(error);
    }
  };

  useEffect(() => {
    if (formSubmit) {
      const validateForm = async () => {
        const validationErrors = await loginValidation(formData);
        setErrors(validationErrors);
        console.log(validationErrors, "errors");

      };
      validateForm();
    }
  }, [formData, formSubmit]);

  const { t } = useTranslation();
  const sm_down = useMediaQuery((theme) => theme.breakpoints.down("sm"));
  const sm_up = useMediaQuery((theme) => theme.breakpoints.up("sm"));
  const mdmd_down = useMediaQuery((theme) => theme.breakpoints.down("mdmd"));

  return (
    <Box
      sx={{
        display: "flex",
        width: "100%",
        alignItems: "center",
        justifyContent: "center",
        height: "100vh",
      }}
    >
      <Box
        sx={{
          width: "50.625rem",
          position: "relative",
          alignItems: "center",
          justifyContent: "center",
          height: "100vh",
          display: "flex",
          px: mdmd_down ? "5rem" : "10rem",
        }}
      >
        <Card
          sx={{
            width: "100%",
            height: "50rem",
          }}
        >
          <CardContent sx={{ height: "calc(100% - 3rem)" }}>
            <Grid
              container
              sx={{
                display: "flex",
                flexDirection: "column",
                justifyContent: "center",
                alignItems: "center",
                px: sm_down ? "1rem" : "3rem",
              }}
            >
              <Grid
                item
                xs={12}
                sx={{
                  display: "flex",
                  marginTop: sm_down ? "5rem" : "8.438rem",
                }}
              >
                <Box
                  sx={{
                    width: "2.5rem",
                    height: "2.5rem",
                    borderRadius: "50%",
                    background: "#FFFFFF",
                    opacity: 1,
                  }}
                />
                <Typography
                  sx={{
                    marginLeft: "1.563rem",
                    font: "normal normal bold 35px/44px Outfit;",
                  }}
                >
                  {themeConfig.projectName}
                </Typography>
              </Grid>
              <Grid item xs={12} sx={{ width: "100%", mt: "2.5rem" }}>
                <Typography
                  sx={{
                    // display: visibleUsernameLabel ? "block" : "none",
                    mb: "0.438rem",
                    ml: "1.438rem",
                    color: (theme) => `${theme.palette.border.secondary}`,
                    fontWeight: "bold",
                    transform: visibleUsernameLabel
                      ? "translateY(0)"
                      : "translateY(20px)",
                    opacity: visibleUsernameLabel ? 1 : 0,
                    transition: "all 0.3s ease-in-out",
                  }}
                >
                  <Translations text={"login.username"} />
                </Typography>
                <FormControl
                  sx={{
                    alignItems: "center",
                    textAlign: "center",
                    width: "100%",
                  }}
                >
                  <TextField
                    name="username"
                    placeholder={t("login.username")}
                    variant="outlined"
                    onChange={handleChange}
                    error={errors.username ? true : false}
                    helperText={errors.username}
                    InputProps={{ style: { color: "#0A3B7A" } }}
                    sx={{
                      "& .MuiFormLabel-root": {
                        color: "#0A3B7A",
                        fontWeight: "bold",
                      },
                      "& .MuiOutlinedInput-root": {
                        "& fieldset": {
                          color: "#0A3B7A",
                          fontWeight: "bold",
                        },
                        "&:hover fieldset": {
                          color: "#0A3B7A",
                          fontWeight: "bold",
                        },
                        "&.Mui-focused": {
                          color: "#0A3B7A",
                          fontWeight: "bold",
                        },
                      },

                      width: "100%",
                      height: "3.125rem",
                    }}
                  />
                </FormControl>
              </Grid>
              <Grid item xs={12} sx={{ width: "100%", mt: "2.5rem" }}>
                <Typography
                  sx={{
                    // display: visiblePasswordLabel ? "block" : "none",
                    mb: "0.438rem",
                    ml: "1.438rem",
                    color: "#0A3B7A",
                    fontWeight: "bold",
                    transform: visiblePasswordLabel
                      ? "translateY(0)"
                      : "translateY(20px)",
                    opacity: visiblePasswordLabel ? 1 : 0,
                    transition: "all 0.3s ease-in-out",
                  }}
                >
                  <Translations text={"login.password"} />
                </Typography>
                <FormControl sx={{ width: "100%" }}>
                  <TextField
                    name="password"
                    placeholder={t("login.password")}
                    variant="outlined"
                    type={showPassword ? "text" : "password"}
                    onChange={handleChange}
                    error={errors.password ? true : false}
                    helperText={errors.password}
                    InputProps={{
                      style: { color: "#0A3B7A" },
                      endAdornment: (
                        <InputAdornment position="end">
                          <IconButton
                            onClick={handleClickShowPassword}
                            edge="end"
                          >
                            <Image
                              style={{ zIndex: 99 }}
                              src={
                                showPassword
                                  ? visibilityOnIcon
                                  : visibilityOffIcon
                              }
                              alt={
                                showPassword
                                  ? "visibilityOnIcon"
                                  : "visibilityOffIcon"
                              }
                              width={30}
                              height={30}
                            />
                          </IconButton>
                        </InputAdornment>
                      ),
                    }}
                    sx={{
                      "& .MuiFormLabel-root": {
                        color: "#0A3B7A",
                        fontWeight: "bold",
                      },
                      "& .MuiOutlinedInput-root": {
                        "& fieldset": {
                          color: "#0A3B7A",
                          fontWeight: "bold",
                        },
                        "&:hover fieldset": {},
                        "&.Mui-focused": {
                          color: "#0A3B7A",
                          fontWeight: "bold",
                        },
                      },
                      zIndex: 99,
                      width: "100%",
                      height: "3.125rem",
                    }}
                  />
                </FormControl>
              </Grid>
              <Grid
                item
                xs={12}
                sx={{
                  width: "100%",
                  mt: "2.5rem",
                  display: "flex",
                  justifyContent: "space-between",
                  flexDirection: mdmd_down ? "column" : "row",
                  textAlign: "center",
                  alignItems: "center",
                }}
              >
                <FormControlLabel
                  control={
                    <Checkbox
                      sx={{
                        color: "#FFF",
                        "&.Mui-checked": {
                          color: "#0A3B7A",
                        },
                      }}
                    />
                  }
                  label={<Translations text={"login.remember.me"} />}
                />
                <Button href="#ForgotPassword" sx={{ color: "#0A3B7A" }}>
                  <Translations text={"login.forget.password"} />
                </Button>
              </Grid>
              <Grid item xs={12} sx={{ width: "100%", mt: "1.163rem" }}>
                <Button
                  sx={{
                    width: "100%",
                    height: "3.125rem",
                  }}
                  variant="dark"
                  onClick={handleSubmit}
                >
                  <Typography
                    sx={{
                      font: "normal normal bold 18px/23px Outfit",
                      letterSpacing: "0px",
                      opacity: 1,
                    }}
                  >
                    <Translations text={"login.login.button"} />
                  </Typography>
                </Button>
              </Grid>
              <Grid item xs={12} sx={{ width: "100%", mt: "1.563rem" }}>
                <Divider sx={{ width: "100%" }}>
                  <Translations text={"login.divider.or"} />
                </Divider>
              </Grid>
              <Grid
                item
                xs={12}
                sx={{
                  mt: "1.563rem",
                  width: "100%",
                  display: "flex",
                  justifyContent: "center",
                }}
              >
                <Button
                  variant="dark"
                  sx={{
                    width: "3.125rem",
                    height: "3.125rem",
                    borderRadius: "0.938rem",
                  }}
                >
                  <Image
                    src={googleIcon}
                    alt="google-icon"
                    width={40}
                    height={40}
                  />
                </Button>
                <Button
                  variant="dark"
                  sx={{
                    width: "3.125rem",
                    height: "3.125rem",
                    borderRadius: "0.938rem",
                    ml: "1.563rem",
                  }}
                >
                  <Image
                    src={githubIcon}
                    alt="github-icon"
                    width={50}
                    height={50}
                  />
                </Button>
              </Grid>
              <Grid
                item
                xs={12}
                sx={{
                  display: "flex",
                  mt: "1.563rem",
                  width: "100%",
                  display: "flex",
                  alignItems: "center",
                  justifyContent: "center",
                  flexDirection: sm_down ? "column" : "row",
                  mb: sm_down ? "1.563rem" : "",
                }}
              >
                <Typography
                  sx={{
                    font: "normal normal normal 18px/23px Outfit",
                    textAlign: "center",
                  }}
                >
                  <Translations text={"login.new.on.platform"} />
                </Typography>
                <Typography
                  sx={{
                    ml: sm_up ? "0.938rem" : 0,
                    mt: sm_down ? "0.938rem" : 0,
                    color: "#0A3B7A",
                    font: "normal normal 600 18px/23px Outfit",
                    textAlign: "center",
                  }}
                >
                  <Translations text={"login.create.new.account"} />
                </Typography>
              </Grid>
            </Grid>
          </CardContent>
        </Card>
        {mdmd_down ? (
          ""
        ) : (
          <Image
            src={rocketImg}
            alt="rocket-icon"
            style={{
              width: "12.438rem",
              height: "16rem",
              position: "absolute",
              top: 50,
              right: 69,
              transform: "rotate(-57deg)",
            }}
          />
        )}
        {mdmd_down ? (
          ""
        ) : (
          <Image
            src={manImg}
            alt="man-icon"
            style={{
              position: "absolute",
              bottom: 60,
              left: -50,
              width: sm_down ? "60px" : "",
              height: sm_down ? "60px" : "",
            }}
          />
        )}
      </Box>
    </Box>
  );
};

Login.guestGuard = true;
Login.getLayout = (page) => <BlankLayout>{page}</BlankLayout>;

export default Login;
