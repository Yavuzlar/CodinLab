import { useTheme } from "@emotion/react";
import {
  Box,
  Container,
  FormControl,
  Grid,
  InputAdornment,
  Card,
  CardContent,
  TextField,
  Typography,
  Link,
  Button,
  IconButton,
  useMediaQuery,
} from "@mui/material";
import Image from "next/image";
import CodinLabLogo from "../assets/logo/main-vertical-square.png";
import visibilityOnIcon from "../assets/icons/icons8-eye-1.png";
import visibilityOffIcon from "../assets/icons/eye-hidden.png";
import LanguageSelector from "src/layout/components/navigation/item/LanguageSelector";
import { useState, useEffect } from "react";
import { registerValidation } from "src/configs/validation/registerSchema";
import CardImage from "src/assets/3d/3d-casual-life-windows-with-developer-code-symbols.png";
import GirlImage from "src/assets/3d/3d-casual-life-girl-holding-laptop-and-having-an-idea.png";
import themeConfig from "src/configs/themeConfig";
import { useTranslation } from "react-i18next";
import { useAuth } from "src/hooks/useAuth";
const { default: BlankLayout } = require("src/layout/BlankLayout");

const Register = () => {
  const [formData, setFormData] = useState();
  const [errors, setErrors] = useState({});
  const [formSubmit, setFormSubmit] = useState(false);

  const md_down = useMediaQuery((theme) => theme.breakpoints.down("md"));

  const [showPassword, setShowPassword] = useState(false);
  const handleClickShowPassword = () => setShowPassword(!showPassword);
  const { register } = useAuth();

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setFormSubmit(true);

    const validationErrors = await registerValidation(formData);
    setErrors(validationErrors);

    if (Object.keys(validationErrors).length > 0) {
      return;
    }
    // Call API
    try {
      await register(formData);
    } catch (error) {}

  };

  addEventListener("keydown", (event) => {
    if (event.key === "Enter") {
      handleSubmit(event);
    }
  });

  useEffect(() => {
    const validate = async () => {
      if (formSubmit) {
        const errors = await registerValidation(formData);
        setErrors(errors);
      }
    };
    validate();
  }, [formData, formSubmit]);

  const theme = useTheme();
  const bgColor = theme.palette.primary.dark;
  const { t } = useTranslation();

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

  const textFieldStyle = {
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
  };

  return (
    <Box
      sx={{
        position: "relative",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",

        height: "100vh",
      }}
    >
      <Box>
        {md_down ? (
          ""
        ) : (
          <Button sx={{ top: 0, right: 0, position: "absolute" }}>
            <LanguageSelector />
          </Button>
        )}
      </Box>
      <Container
        sx={{
          display: "flex",
          justifyContent: "center",
          position: "relative",
        }}
      >
        <Box
        sx={{
          display: { xs: "none", md: "block" },
          position: "absolute",
          top: "-6.5%",
          left: "-1rem",
          zIndex: 1,
        }}
      >
        <Image src={CardImage} width={368} height={226} alt="Cards" />
      </Box>
      <Box
        sx={{
          display: { xs: "none", md: "block" },
          position: "absolute",
          top: "3.4rem",
          right: "-7.8rem",
          zIndex: 1,
        }}
      >
        <Image
          src={GirlImage}
          width={368}
          height={803}
          priority
          alt="Girl holding laptop"
        />
      </Box>
        <Card
          sx={{
            width: { xs: "100%", sm: "auto" },
          }}
        >
          <CardContent
            sx={{
              width: {
                mdmd: "auto",
                md: "35rem",
                mdlg: "40rem",
                mdxl: "45rem",
                lg: "50.75rem",
              },
            }}
          >
            <Grid
              container
              direction="column"
              sx={{
                px: { xs: 4, sm: 6, md: 8, lg: 10, xl: 12, xxl: 14 },
                position: "relative",
              }}
            >
              {md_down ? (
                <Button sx={{ top: 0, right: 0, position: "absolute" }}>
                  <LanguageSelector />
                </Button>
              ) : (
                ""
              )}
              <Box
                sx={{
                  display: "flex",
                  justifyContent: "center",
                  alignItems: "center",
                  flexDirection: "column",
                  gap: 1,
                  my: 5,
                }}
              >
                <Image
                  src={CodinLabLogo}
                  alt="codinlab-logo"
                  width={200}
                  height={200}
                />
              </Box>
              <FormControl>
                <Grid container direction="column" gap={3}>
                  <TextField
                    name="name"
                    placeholder={t("register.name")}
                    variant="outlined"
                    InputLabelProps={inputLabelStyle}
                    onChange={handleChange}
                    error={errors.name ? true : false}
                    helperText={errors.name}
                    sx={textFieldStyle}
                  />
                  <TextField
                    name="surname"
                    placeholder={t("register.surname")}
                    InputLabelProps={inputLabelStyle}
                    onChange={handleChange}
                    error={errors.surname ? true : false}
                    helperText={errors.surname}
                    sx={textFieldStyle}
                  />
                  <TextField
                    name="username"
                    placeholder={t("register.username")}
                    InputLabelProps={inputLabelStyle}
                    onChange={handleChange}
                    error={errors.username ? true : false}
                    helperText={errors.username}
                    sx={textFieldStyle}
                  />
                  <TextField
                    name="githubProfile"
                    placeholder={t("register.githubProfile")}
                    InputLabelProps={inputLabelStyle}
                    onChange={handleChange}
                    error={errors.githubProfile ? true : false}
                    helperText={errors.githubProfile}
                    sx={textFieldStyle}
                  />
                  <TextField
                    name="password"
                    placeholder={t("register.password")}
                    InputLabelProps={inputLabelStyle}
                    type={showPassword ? "text" : "password"}
                    onChange={handleChange}
                    error={errors.password ? true : false}
                    helperText={errors.password}
                    sx={textFieldStyle}
                    InputProps={{
                      endAdornment: (
                        <InputAdornment>
                          <IconButton
                            sx={{ zIndex: 999 }}
                            aria-label="toggle password visibility"
                            onClick={handleClickShowPassword}
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
                  />
                  {/* CheckBox Start */}
                  {/* <FormControlLabel
                    control={
                      <Checkbox
                        name="checkbox"
                        onChange={handleChange}
                        sx={{
                          color: "#FFF",
                          "&.Mui-checked": {
                            color: "#0A3B7A",
                          },
                          "& .MuiSvgIcon-root": {
                            color: errors.checkbox ? "red" : "#FFF",
                          },
                        }}
                        error={errors.checkbox ? true : false}
                      />
                    }
                    label={
                      <Typography
                        fontWeight={300}
                        fontSize={18}
                        fontFamily={"Outfit"}
                      >
                        {t("register.accept")}
                        <Link
                          sx={{ textDecoration: "none", fontWeight: "600" }}
                          color={"#0A3B7A"}
                          href="#"
                        >
                          {t("register.terms")}
                        </Link>
                      </Typography>
                    }
                  /> */}
                  {/* CheckBox End */}
                  <Button
                    variant="dark"
                    sx={{
                      font: "normal normal 18px/23px Outfit",
                      fontWeight: "600",
                      textTransform: "capitalize",
                      py: 2,
                    }}
                    onClick={handleSubmit}
                    fullWidth
                  >
                    {t("register.signup")}
                  </Button>
                </Grid>
              </FormControl>
              {/* Divider and Google & GitHub Buttuns Start */}
              {/* <Divider sx={{ mt: 3 }}> {t("register.or")}</Divider> */}
              {/* <Stack direction="row" justifyContent="center" gap={3} mt={3}>
                <IconButton variant="contained" sx={iconBtnStyle}>
                  <Google sx={iconSize} />
                </IconButton>
                <IconButton variant="contained" sx={iconBtnStyle}>
                  <GitHub sx={iconSize} />
                </IconButton>
              </Stack> */}
              {/* Divider and Google & GitHub Buttuns End */}
              <Typography
                variant="body1"
                textAlign={"center"}
                mt={4}
                fontFamily={"Outfit"}
              >
                {t("register.already")}
                <Link
                  href="/login"
                  color={bgColor}
                  sx={{ fontWeight: "600", textDecoration: "none", ml: 1 }}
                >
                  {t("register.login")}
                </Link>
              </Typography>
            </Grid>
          </CardContent>
        </Card>
      </Container>
    </Box>
  );
};

Register.guestGuard = true;
Register.getLayout = (page) => <BlankLayout>{page}</BlankLayout>;

export default Register;
