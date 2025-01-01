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
import CofeeImage from "../assets/3d/casual-life-3d-green-coffee-cup.png";
import Translations from "src/components/Translations";
import visibilityOnIcon from "../assets/icons/icons8-eye-1.png";
import visibilityOffIcon from "../assets/icons/eye-hidden.png";
import { useTheme } from "@mui/material/styles";
import { useState, useEffect } from "react";
import MauseImage from "../assets/3d/casual-life-3d-cursor.png";
import { profileSettingsValidation } from "src/configs/validation/profileSettingsSchema";
import { changePasswordValidation } from "src/configs/validation/changePassSchema";
import { useDispatch, useSelector } from "react-redux";
import { fetchProfileUser } from "src/store/user/userSlice";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogTitle from "@mui/material/DialogTitle";
import { changePassword, changeProfile } from "src/store/user/userSlice";
import CustomBreadcrumbs from "src/components/breadcrumbs";

const settings = () => {
  const theme = useTheme();
  const dispatch = useDispatch();
  const { user: stateUser } = useSelector((state) => state);

  const [passwordSettingsData, setPasswordSettingsData] = useState();
  const [infoSettingsData, setInfoSettingsData] = useState({});

  const [showOldPassword, setShowOldPassword] = useState(false);
  const [showNewPassword, setShowNewPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [showPassword, setShowPassword] = useState(false);

  const [visibleUsernameLabel, setVisibleUsernameLabel] = useState(false);
  const [visibleGithubLabel, setVisibleGithubLabel] = useState(false);
  const [visibleNameLabe, setVisibleNameLabel] = useState(false);
  const [visibleSurnameLabel, setVisibleSurnameLabel] = useState(false);

  const [visibleOldPasswordLabel, setVisibleOldPasswordLabel] = useState(false);
  const [visibleNewPasswordLabel, setVisibleNewPasswordLabel] = useState(false);
  const [visibleConfirmPasswordLabel, setVisibleConfirmPasswordLabel] =
    useState(false);

  const [infoSettingsSubmitted, setInfoSettingsSubmitted] = useState(false);
  const [passwordSettingsSubmitted, setPasswordSettingsSubmitted] =
    useState(false);

  const [errorInfo, setErrorInfo] = useState();
  const [errorPassword, setErrorPassword] = useState({});

  const [openDialog, setOpenDialog] = useState(false);
  const [password, setPassword] = useState("");

  useEffect(() => {
    const validateInfoSettings = async () => {
      if (infoSettingsSubmitted) {
        const validationInfoErrors = await profileSettingsValidation(
          infoSettingsData
        );
        setErrorInfo(validationInfoErrors);
      }
    };
    validateInfoSettings();
  }, [infoSettingsData, infoSettingsSubmitted]);

  useEffect(() => {
    const validateInfoSettings = async () => {
      if (infoSettingsSubmitted) {
        const validationInfoErrors = await profileSettingsValidation(
          infoSettingsData
        );
        setErrorInfo(validationInfoErrors);
      }
    };
    validateInfoSettings();
  }, [infoSettingsData, infoSettingsSubmitted]);

  useEffect(() => {
    dispatch(fetchProfileUser());
  }, []);

  useEffect(() => {
    //this is for the api call
    dispatch(fetchProfileUser());
  }, []);

  useEffect(() => {
    //this is the  data for the user in api
    if (stateUser.data) {
      //this is checking if the data is available
      setInfoSettingsData({
        name: stateUser.data.data?.name,
        surname: stateUser.data.data?.surname,
        username: stateUser.data.data?.username,
        githubProfile: stateUser.data.data?.githubProfile,
      });
    }
  }, [stateUser.data]);

  const hanldeClickShowOldPassword = () => {
    setShowOldPassword(!showOldPassword);
  };

  const hanldeClickShowNewPassword = () => {
    setShowNewPassword(!showNewPassword);
  };

  const hanldeClickShowConfirmPassword = () => {
    setShowConfirmPassword(!showConfirmPassword);
  };

  const hanldeClickShowPassword = () => {
    setShowPassword(!showPassword);
  };

  const handleClose = () => {
    setOpenDialog(false);
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

    if (e.target.name === "githubProfile") {
      if (e.target.value.length > 0) {
        setVisibleGithubLabel(true);
      } else {
        setVisibleGithubLabel(false);
      }
    }

    if (e.target.name === "name") {
      if (e.target.value.length > 0) {
        setVisibleNameLabel(true);
      } else {
        setVisibleNameLabel(false);
      }
    }

    if (e.target.name === "surname") {
      if (e.target.value.length > 0) {
        setVisibleSurnameLabel(true);
      } else {
        setVisibleSurnameLabel(false);
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

  const handleSubmitInfoSettings = async (e) => {
    e.preventDefault();
    setInfoSettingsSubmitted(true);

    const validationInfoErrors = await profileSettingsValidation(
      infoSettingsData
    );
    setErrorInfo(validationInfoErrors);
    if (
      !validationInfoErrors ||
      Object.keys(validationInfoErrors).length === 0
    ) {
      setOpenDialog(true);
    }
  };

  const handleDialogSubmit = () => {
    const dataToSend = {
      ...infoSettingsData,
      password,
    };

    dispatch(changeProfile(dataToSend));

    setOpenDialog(false);
  };

  const handleSubmitPasswordSettings = async (e) => {
    e.preventDefault();
    setPasswordSettingsSubmitted(true);

    const validationPasswordErrors = await changePasswordValidation(
      passwordSettingsData
    );
    setErrorPassword(validationPasswordErrors);
    dispatch(changePassword(passwordSettingsData));

    if (validationPasswordErrors) {
      return;
    }
    try {
      // console.log("TryGönderilenVeri", passwordSettingsData);
    } catch (error) {
      console.log("Catch Error", error);
    }
  };
  const breacrumbs = [
    {
      path: "/settings",
      title: <Translations text={"settings"} />,
      permission: "settings",
    },
  ];

  return (
    <Box>
      <CustomBreadcrumbs titles={breacrumbs} />
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

                  <FormControl
                    fullWidth
                    sx={{
                      marginTop: 4,
                    }}
                  >
                    <Typography
                      sx={{
                        mb: "0.2rem",
                        color: (theme) =>
                          `${theme.palette.border.secondary} !important`,
                        fontWeight: "bold",
                        transform: visibleNameLabe
                          ? "translateY(0)"
                          : "translateY(20px)",
                        opacity: visibleNameLabe ? 1 : 0,
                        transition: "all 0.3s ease-in-out",
                      }}
                    >
                      <Translations text={"settings.change.name"} />
                    </Typography>
                    <TextField
                      id="outlined-basic"
                      placeholder="Jhon"
                      variant="outlined"
                      name="name"
                      value={infoSettingsData?.name}
                      error={errorInfo?.name ? true : false}
                      helperText={errorInfo?.name}
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
                      marginTop: 4,
                    }}
                  >
                    <Typography
                      sx={{
                        mb: "0.2rem",
                        color: (theme) =>
                          `${theme.palette.border.secondary} !important`,
                        fontWeight: "bold",
                        transform: visibleSurnameLabel
                          ? "translateY(0)"
                          : "translateY(20px)",
                        opacity: visibleSurnameLabel ? 1 : 0,
                        transition: "all 0.3s ease-in-out",
                      }}
                    >
                      <Translations text={"settings.change.surname"} />
                    </Typography>
                    <TextField
                      id="outlined-basic"
                      placeholder="Doe"
                      variant="outlined"
                      name="surname"
                      value={infoSettingsData?.surname}
                      onChange={handleInfoSettings}
                      error={errorInfo?.surname ? true : false}
                      helperText={errorInfo?.surname}
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
                    sx={{
                      marginTop: 4,
                    }}
                    fullWidth
                  >
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
                      value={infoSettingsData?.username}
                      onChange={handleInfoSettings}
                      error={errorInfo?.username ? true : false}
                      helperText={errorInfo?.username}
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
                    sx={{
                      marginTop: 4,
                    }}
                    fullWidth
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
                      name="githubProfile"
                      value={infoSettingsData?.githubProfile}
                      onChange={handleInfoSettings}
                      error={errorInfo?.github ? true : false}
                      helperText={errorInfo?.github}
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
                    onClick={handleSubmitInfoSettings}
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
                      name="password"
                      type={showOldPassword ? "text" : "password"}
                      onChange={handlePasswordSettings}
                      error={errorPassword?.oldPassword ? true : false}
                      helperText={errorPassword?.oldPassword}
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
                      marginTop: 4,
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
                      error={errorPassword?.newPassword ? true : false}
                      helperText={errorPassword?.newPassword}
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
                      marginTop: 4,
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
                      error={errorPassword?.confirmPassword ? true : false}
                      helperText={errorPassword?.confirmPassword}
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
                    onClick={handleSubmitPasswordSettings}
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

      <Dialog
        open={openDialog}
        onClose={handleClose}
        sx={{
          "& .MuiDialog-paper": {
            color: (theme) => `${theme.palette.text.primary} !important`,
            borderRadius: "16px",
            boxShadow: "0px 4px 4px rgba(0, 0, 0, 0.25)",
            width: "100%",
            maxWidth: "500px",
            padding: "20px",
          },
        }}
      >
        <DialogTitle>
          <Translations text={"dialog.title.password"} />
        </DialogTitle>
        <DialogContent>
          <DialogContentText
            sx={{
              color: (theme) => `${theme.palette.text.primary} !important`,
            }}
          >
            <Translations text={"dialog.content.password"} />
          </DialogContentText>
          <TextField
            sx={{
              marginTop: "20px",
              borderRadius: "15px",
              border: "2px solid #0A3B7A",

              "& .MuiInputBase-root": {
                backgroundColor: (theme) =>
                  `${theme.palette.background.default} !important`,
              },

              "& .MuiOutlinedInput-root": {
                borderRadius: "15px",
                "&.Mui-focused fieldset": {
                  borderColor: (theme) =>
                    `${theme.palette.primary.dark} !important`,
                },
              },

              "& .MuiInputBase-input": {
                color: "#0A3B7A",
                fontWeight: "bold",
                marginTop: "5px",
              },

              "& .MuiInputLabel-root": {
                color: (theme) => `${theme.palette.primary.dark} !important`,
                fontWeight: "bold",
              },

              "& .Mui-focused": {
                "& .MuiOutlinedInput-notchedOutline": {
                  borderColor: "#0A3B7A",
                },
              },
            }}
            autoFocus
            margin="dense"
            label={<Translations text={"settings.old.password"} />}
            type={showPassword ? "text" : "password"}
            variant="filled"
            fullWidth
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            InputProps={{
              endAdornment: (
                <IconButton onClick={hanldeClickShowPassword} edge="end">
                  <Image
                    style={{ zIndex: 99 }}
                    src={showPassword ? visibilityOnIcon : visibilityOffIcon}
                    alt={
                      showPassword ? "visibilityOnIcon" : "visibilityOffIcon"
                    }
                    width={30}
                    height={30}
                  />
                </IconButton>
              ),
            }}
          />
        </DialogContent>
        <DialogActions
          sx={{
            display: "flex",
            justifyContent: "space-between",
            padding: "20px",
          }}
        >
          <Button onClick={handleClose} variant="dark">
            <Translations text={"dialog.button.cancel"} />
          </Button>
          <Button onClick={handleDialogSubmit} variant="dark">
            <Translations text={"dialog.button.submit"} />
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
};

export default settings;
