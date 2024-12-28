import {
  FormControl,
  Button,
  TextField,
  Card,
  CardContent,
  Typography,
  Box,
  IconButton,
  InputAdornment,
  Grid,
  useMediaQuery,
} from "@mui/material";
import { useState, useEffect } from "react";
import Translations from "src/components/Translations";
import { loginValidation } from "src/configs/validation/loginSchema";
import rocketImg from "../assets/3d/3d-casual-life-space-white-starship.png";
import visibilityOnIcon from "../assets/icons/icons8-eye-1.png";
import visibilityOffIcon from "../assets/icons/eye-hidden.png";
import CodinLabLogo from "../assets/logo/codinlab-logo-main.png";
import Image from "next/image";
import manImg from "../assets/3d/3d-casual-life-young-man-sitting-with-laptop-and-waving.png";
import { useTranslation } from "next-i18next";
import themeConfig from "src/configs/themeConfig";
import { useAuth } from "src/hooks/useAuth";
import LanguageSelector from "src/layout/components/navigation/item/LanguageSelector";
import { useRouter } from "next/router";

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
  const router = useRouter();
  // enter tuşu ile login olma
  // addEventListener("keydown", (event) => {
  //   if (event.key === "Enter") {
  //     handleSubmit();
  //   }
  // });

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
      await login(formData);
    } catch (error) {}
  };

  const goRegisterPage = () => {
    router.push("register");
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
  const md_down = useMediaQuery((theme) => theme.breakpoints.down("md"));
  const mdxl_down = useMediaQuery((theme) => theme.breakpoints.down("mdxl"));
  const lg_down = useMediaQuery((theme) => theme.breakpoints.down("lg"));

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
      {md_down ? (
        ""
      ) : (
        <Button sx={{ top: 10, right: 5, position: "absolute" }}>
          <LanguageSelector />
        </Button>
      )}
      <Box
        sx={{
          width: "100%",
          ...(mdxl_down ? { maxWidth: "35rem" } : { maxWidth: "45rem" }),
          position: "relative",
          alignItems: "center",
          justifyContent: "center",
          maxHeight: "100%",
          display: "flex",
        }}
      >
        <Card
          sx={{
            width: "100%",
          }}
        >
          <CardContent
            sx={{ height: "calc(100% - 3rem)", position: "relative" }}
          >
            <Grid
              container
              sx={{
                display: "flex",
                flexDirection: "column",
                justifyContent: "center",
                alignItems: "center",
                px: sm_down ? "1rem" : "3rem",
              }}
              spacing={4}
            >
              {md_down ? (
                <Button sx={{ top: 10, right: 5, position: "absolute" }}>
                  <LanguageSelector />
                </Button>
              ) : (
                ""
              )}

              <Grid
                item
                xs={12}
                sx={{
                  display: "flex",
                  marginTop: sm_down ? "2rem" : "4rem",
                  alignItems: "center",
                  alignContent: "center",
                  flexDirection: "column",
                }}
              >
                <Image
                  src={CodinLabLogo}
                  alt="codinlab-logo"
                  width={80}
                  height={120}
                />
                <Typography
                  sx={{
                    font: "normal normal bold 35px/44px Outfit;",
                  }}
                >
                  {themeConfig.projectName}
                </Typography>
              </Grid>
              <Grid item xs={12} sx={{ width: "100%" }}>
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
                    InputProps={{
                      style: { color: "#0A3B7A" },
                    }}
                    sx={{
                      "& .MuiFormLabel-root": {
                        color: "#0A3B7A",
                        // fontWeight: "bold",
                      },
                      "& .MuiOutlinedInput-root": {
                        "& input::placeholder": {
                          fontWeight: "bold",
                        },
                        // "& fieldset": {
                        //   color: "#0A3B7A",
                        //   // fontWeight: "bold",
                        // },
                        // "&:hover fieldset": {
                        //   color: "#0A3B7A",
                        //   // fontWeight: "bold",
                        // },
                        // "&.Mui-focused": {
                        //   color: "#0A3B7A",
                        //   // fontWeight: "bold",
                        // },
                      },

                      width: "100%",
                      height: "3.125rem",
                    }}
                  />
                </FormControl>
              </Grid>
              <Grid item xs={12} sx={{ width: "100%" }}>
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
                        "& input::placeholder": {
                          fontWeight: "bold",
                        },
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
                }}
              >
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
              {/*               
              <Grid item xs={12} sx={{ width: "100%" }}>
                <Divider sx={{ width: "100%" }}>
                  <Translations text={"login.divider.or"} />
                </Divider>
              </Grid>
              <Grid
                item
                xs={12}
                sx={{
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
              </Grid> */}
              <Grid
                item
                xs={12}
                sx={{
                  display: "flex",
                  width: "100%",
                  display: "flex",
                  alignItems: "center",
                  justifyContent: "center",
                  flexDirection: sm_down ? "column" : "row",
                  mb: sm_down ? "1.563rem" : "",
                  gap: 2,
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
                <Button onClick={goRegisterPage}>
                  <Typography
                    sx={{
                      color: "#0A3B7A !important",

                      font: "normal normal 600 18px/23px Outfit",
                      textAlign: "center",
                    }}
                  >
                    <Translations text={"login.create.new.account"} />
                  </Typography>
                </Button>
              </Grid>
            </Grid>
          </CardContent>
        </Card>
        {md_down ? (
          ""
        ) : (
          <Image
            src={rocketImg}
            alt="rocket-icon"
            style={{
              ...(lg_down ? { width: "10rem" } : { width: "12rem" }),
              ...(lg_down ? { height: "12rem" } : { height: "16rem" }),
              ...(lg_down ? { right: -75 } : { right: -88 }),
              position: "absolute",
              top: -64,
              transform: "scaleX(-1)",
            }}
          />
        )}
        {md_down ? (
          ""
        ) : (
          <Image
            src={manImg}
            alt="man-icon"
            style={{
              position: "absolute",
              bottom: -48,
              ...(lg_down ? { width: "18rem" } : { width: "" }),
              ...(lg_down ? { height: "20rem" } : { height: "" }),
              ...(lg_down ? { left: -210 } : { left: -264 }),
              marginLeft: "1rem",
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
