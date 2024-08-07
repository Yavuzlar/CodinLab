import {
    Box,
    Button,
    Card,
    CardContent,
    FormControl,
    Grid,
    IconButton,
    TextField,
    Typography,
  } from "@mui/material";
  import Image from "next/image";
  import ComputerImage from "../assets/3d/3d-casual-life-website-browser-window-in-laptop.png";
  import MauseImage from "../assets/3d/casual-life-3d-cursor.png";
  import CofeeImage from "../assets/3d/casual-life-3d-green-coffee-cup.png";
  import Translations from "src/components/Translations";
  import { useState } from "react";
  import visibilityOnIcon from "../assets/icons/icons8-eye-1.png";
  import visibilityOffIcon from "../assets/icons/eye-hidden.png";
  import { useTheme } from "@mui/material/styles";
  
  const settings = () => {
    const [infoSettingsData, setInfoSettingsData] = useState(null);
    const [passwordSettingsData, setPasswordSettingsData] = useState(null);
  
    const [showOldPassword, setShowOldPassword] = useState(false);
    const [showNewPassword, setShowNewPassword] = useState(false);
    const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  
    const [visibleUsernameLabel, setVisibleUsernameLabel] = useState(false);
    const [visibleGithubLabel, setVisibleGithubLabel] = useState(false);
  
    const [visibleOldPasswordLabel, setVisibleOldPasswordLabel] = useState(false);
    const [visibleNewPasswordLabel, setVisibleNewPasswordLabel] = useState(false);
    const [visibleConfirmPasswordLabel, setVisibleConfirmPasswordLabel] =
      useState(false);
  
    const [error, setError] = useState(null);
  
    const hanldeClickShowOldPassword = () => {
      setShowOldPassword(!showOldPassword);
    };
    const hanldeClickShowNewPassword = () => {
      setShowNewPassword(!showNewPassword);
    };
    const hanldeClickShowConfirmPassword = () => {
      setShowConfirmPassword(!showConfirmPassword);
    };
  
    const handleInfoSettings = (e) => {
      setInfoSettingsData({
        ...infoSettingsData,
        [e.target.name]: e.target.value,
      });
  
      if (e.target.name === "username") {
        if (e.target.value.length > 0) {
          setVisibleUsernameLabel(true);
        } else {
          setVisibleUsernameLabel(false);
        }
      }
  
      if (e.target.name === "github") {
        if (e.target.value.length > 0) {
          setVisibleGithubLabel(true);
        } else {
          setVisibleGithubLabel(false);
        }
      }
    };
  
    const handlePasswordSettings = (e) => {
      setPasswordSettingsData({
        ...passwordSettingsData,
        [e.target.name]: e.target.value,
      });
  
      if (e.target.name === "oldPassword") {
        if (e.target.value.length > 0) {
          setVisibleOldPasswordLabel(true);
        } else {
          setVisibleOldPasswordLabel(false);
        }
      }
  
      if (e.target.name === "newPassword") {
        if (e.target.value.length > 0) {
          setVisibleNewPasswordLabel(true);
        } else {
          setVisibleNewPasswordLabel(false);
        }
      }
  
      if (e.target.name === "confirmPassword") {
        if (e.target.value.length > 0) {
          setVisibleConfirmPasswordLabel(true);
        } else {
          setVisibleConfirmPasswordLabel(false);
        }
      }
    };
  
    const theme = useTheme();
  
   
  
    // when de settings validation is done
    // useEffect will be added here for the passwordSettingsData and infoSettingsData
  
    return (
      <div>
        <Box
          sx={{
            display: "flex",
            justifyContent: "flex-start",
            alignItems: "center",
            position: "relative",
            width: "100%",
          }}
        >
          <Box>
            <Box
              sx={{
                display: { xs: "none", mdlg: "block" },
                position: "absolute",
                top: "35.5%",
                right: {
                  mdlg: "-5%",
                  mdxl: "1%",
                  lg: "6%",
                  lgPlus: "6%",
                  lgXl: "6%",
                  xl: "6%",
                  xxl: "6%",
                },
                zIndex: 1,
              }}
            >
              <Image src={ComputerImage} width={450} height={256} />
            </Box>
            <Box
              sx={{
                display: { xs: "none", mdlg: "block" },
                position: "absolute",
                top: "16%",
                right: {
                  mdlg: "2%",
                  mdxl: "2%",
                  lg: "5%",
                  lgPlus: "5%",
                  lgXl: "5%",
                  xl: "5%",
                  xxl: "5%",
                },
                zIndex: 1,
              }}
            >
              <Image src={MauseImage} width={100} height={150} />
            </Box>
            <Box
              sx={{
                display: { xs: "none", mdlg: "block" },
                position: "absolute",
                top: "60%",
                right: {
                  mdlg: "-3%",
                  mdxl: "-3%",
                  lg: "4%",
                  lgPlus: "4%",
                  lgXl: "4%",
                  xl: "4%",
                  xxl: "4%",
                },
                zIndex: 1,
              }}
            >
              <Image src={CofeeImage} width={140} height={110} />
            </Box>
          </Box>
  
          <Box
            sx={{
              width: "100%",
              margin: "auto",
              display: {
                xs: "flex",
                sm: "flex",
                smd: "flex",
                mdmd: "flex",
                md: "flex",
                mdlg: "block",
              },
              justifyContent: {
                xs: "center",
                sm: "center",
                smd: "center",
                mdmd: "center",
                md: "center",
              },
              alignItems: {
                xs: "center",
                sm: "center",
                smd: "center",
                mdmd: "center",
                md: "center",
              },
            }}
          >
            <Card sx={{ width: "50%", padding: "40px" }}>
              <CardContent>
                <Grid container spacing={2}>
                  <Grid item xs={12} md={12}>
                    <Typography
                      variant="h5"
                      sx={{
                        fontWeight: "bold",
                        display: "flex",
                        justifyContent: "center",
                        alignItems: "center",
                        marginBottom: 2,
                      }}
                    >
                      <Translations text={"settings.page.title"} />
                    </Typography>
  
                    <FormControl fullWidth>
                      <Typography
                        sx={{
                          mb: "0.2rem",
                          color: (theme) =>
                            `${theme.palette.border.secondary} !important`,
                          fontWeight: "bold",
                          transform: visibleUsernameLabel
                            ? "translateY(0)"
                            : "translateY(20px)",
                          opacity: visibleUsernameLabel ? 1 : 0,
                          transition: "all 0.3s ease-in-out",
                        }}
                      >
                        <Translations text={"settings.change.username"} />
                      </Typography>
                      <TextField
                        id="outlined-basic"
                        placeholder="JhonDoe"
                        variant="outlined"
                        name="username"
                        onChange={handleInfoSettings}
                        fullWidth
                        sx={{
                          height: "52px",
                          "& .MuiOutlinedInput-root": {
                            "&.Mui-focused fieldset": {
                              borderColor: "#0A3B7A",
                            },
                          },
                          "& .MuiInputBase-input": {
                            color: "#0A3B7A",
                            fontWeight: "bold",
                          },
                        }}
                      />
                    </FormControl>
  
                    <FormControl
                      fullWidth
                      sx={{
                        marginTop: 3,
                      }}
                    >
                      <Typography
                        sx={{
                          mb: "0.2rem",
                          color: (theme) =>
                            `${theme.palette.border.secondary} !important`,
                          fontWeight: "bold",
                          transform: visibleGithubLabel
                            ? "translateY(0)"
                            : "translateY(20px)",
                          opacity: visibleGithubLabel ? 1 : 0,
                          transition: "all 0.3s ease-in-out",
                        }}
                      >
                        <Translations text={"settings.change.github"} />
                      </Typography>
                      <TextField
                        id="outlined-basic"
                        placeholder="jhondoe"
                        variant="outlined"
                        name="github"
                        onChange={handleInfoSettings}
                        fullWidth
                        sx={{
                          height: "52px",
                          "& .MuiOutlinedInput-root": {
                            "&.Mui-focused fieldset": {
                              borderColor: "#0A3B7A",
                            },
                          },
                          "& .MuiInputBase-input": {
                            color: "#0A3B7A",
                            fontWeight: "bold",
                          },
                        }}
                      />
                    </FormControl>
                    <Button
                      variant="dark"
                      sx={{
                        marginTop: 4,
                        width: "45%",
                        position: "relative",
                        left: "55%",
                        height: "52px",
                        textTransform: "none",
                      }}
                    >
                      <Typography
                        variant="infoText"
                        sx={{
                          color: `${theme.palette.common.white} !important`,
                        }}
                      >
                        <Translations text={"settings.change.button"} />
                      </Typography>
                    </Button>
  
                    <FormControl
                      fullWidth
                      sx={{
                        marginTop: 8,
                      }}
                    >
                      <Typography
                        sx={{
                          mb: "0.2rem",
                          color: (theme) =>
                            `${theme.palette.border.secondary} !important`,
                          fontWeight: "bold",
                          transform: visibleOldPasswordLabel
                            ? "translateY(0)"
                            : "translateY(20px)",
                          opacity: visibleOldPasswordLabel ? 1 : 0,
                          transition: "all 0.3s ease-in-out",
                        }}
                      >
                        <Translations text={"settings.old.password"} />
                      </Typography>
                      <TextField
                        id="outlined-basic"
                        placeholder="********"
                        variant="outlined"
                        name="oldPassword"
                        type={showOldPassword ? "text" : "password"}
                        onChange={handlePasswordSettings}
                        InputProps={{
                          endAdornment: (
                            <IconButton
                              onClick={hanldeClickShowOldPassword}
                              edge="end"
                            >
                              <Image
                                style={{ zIndex: 99 }}
                                src={
                                  showOldPassword
                                    ? visibilityOnIcon
                                    : visibilityOffIcon
                                }
                                alt={
                                  showOldPassword
                                    ? "visibilityOnIcon"
                                    : "visibilityOffIcon"
                                }
                                width={30}
                                height={30}
                              />
                            </IconButton>
                          ),
                        }}
                        fullWidth
                        sx={{
                          height: "52px",
                          "& .MuiOutlinedInput-root": {
                            "&.Mui-focused fieldset": {
                              borderColor: "#0A3B7A",
                            },
                          },
                          "& .MuiInputBase-input": {
                            color: "#0A3B7A",
                            fontWeight: "bold",
                          },
                        }}
                      />
                    </FormControl>
  
                    <FormControl
                      fullWidth
                      sx={{
                        marginTop: 3,
                      }}
                    >
                      <Typography
                        sx={{
                          mb: "0.2rem",
                          color: (theme) =>
                            `${theme.palette.border.secondary} !important`,
                          fontWeight: "bold",
                          transform: visibleNewPasswordLabel
                            ? "translateY(0)"
                            : "translateY(20px)",
                          opacity: visibleNewPasswordLabel ? 1 : 0,
                          transition: "all 0.3s ease-in-out",
                        }}
                      >
                        <Translations text={"settings.new.password"} />
                      </Typography>
                      <TextField
                        id="outlined-basic"
                        placeholder="********"
                        variant="outlined"
                        name="newPassword"
                        type={showNewPassword ? "text" : "password"}
                        onChange={handlePasswordSettings}
                        fullWidth
                        InputProps={{
                          endAdornment: (
                            <IconButton
                              onClick={hanldeClickShowNewPassword}
                              edge="end"
                            >
                              <Image
                                style={{ zIndex: 99 }}
                                src={
                                  showNewPassword
                                    ? visibilityOnIcon
                                    : visibilityOffIcon
                                }
                                alt={
                                  showNewPassword
                                    ? "visibilityOnIcon"
                                    : "visibilityOffIcon"
                                }
                                width={30}
                                height={30}
                              />
                            </IconButton>
                          ),
                        }}
                        sx={{
                          height: "52px",
                          "& .MuiOutlinedInput-root": {
                            "&.Mui-focused fieldset": {
                              borderColor: "#0A3B7A",
                            },
                          },
                          "& .MuiInputBase-input": {
                            color: "#0A3B7A",
                            fontWeight: "bold",
                          },
                        }}
                      />
                    </FormControl>
  
                    <FormControl
                      fullWidth
                      sx={{
                        marginTop: 3,
                      }}
                    >
                      <Typography
                        sx={{
                          mb: "0.2rem",
                          color: (theme) =>
                            `${theme.palette.border.secondary} !important`,
                          fontWeight: "bold",
                          transform: visibleConfirmPasswordLabel
                            ? "translateY(0)"
                            : "translateY(20px)",
                          opacity: visibleConfirmPasswordLabel ? 1 : 0,
                          transition: "all 0.3s ease-in-out",
                        }}
                      >
                        <Translations text={"settings.confirm.password"} />
                      </Typography>
                      <TextField
                        id="outlined-basic"
                        placeholder="********"
                        variant="outlined"
                        name="confirmPassword"
                        type={showConfirmPassword ? "text" : "password"}
                        onChange={handlePasswordSettings}
                        fullWidth
                        InputProps={{
                          endAdornment: (
                            <IconButton
                              onClick={hanldeClickShowConfirmPassword}
                              edge="end"
                            >
                              <Image
                                style={{ zIndex: 99 }}
                                src={
                                  showConfirmPassword
                                    ? visibilityOnIcon
                                    : visibilityOffIcon
                                }
                                alt={
                                  showOldPassword
                                    ? "visibilityOnIcon"
                                    : "visibilityOffIcon"
                                }
                                width={30}
                                height={30}
                              />
                            </IconButton>
                          ),
                        }}
                        sx={{
                          height: "52px",
                          "& .MuiOutlinedInput-root": {
                            "&.Mui-focused fieldset": {
                              borderColor: "#0A3B7A",
                            },
                          },
                          "& .MuiInputBase-input": {
                            color: "#0A3B7A",
                            fontWeight: "bold",
                          },
                        }}
                      />
                    </FormControl>
                    <Button
                      variant="dark"
                      sx={{
                        marginTop: 4,
                        width: "45%",
                        position: "relative",
                        left: "55%",
                        height: "52px",
                        textTransform: "none",
                      }}
                    >
                      <Typography
                        sx={{
                          color: `${theme.palette.common.white} !important`,
                        }}
                      >
                        <Translations text={"settings.password.button"} />
                      </Typography>
                    </Button>
                  </Grid>
                </Grid>
              </CardContent>
            </Card>
          </Box>
        </Box>
      </div>
    );
  };
  
  export default settings;
  