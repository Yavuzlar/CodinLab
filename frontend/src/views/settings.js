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
  };

  const handlePasswordSettings = (e) => {
    setPasswordSettingsData({
      ...passwordSettingsData,
      [e.target.name]: e.target.value,
    });
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
                    General Settings
                  </Typography>

                  <FormControl
                    fullWidth
                    sx={{
                      marginTop: 4,
                      marginBottom: 4,
                    }}
                  >
                    <TextField
                      id="outlined-basic"
                      label={<Translations text={"admin.change.username"} />}
                      variant="outlined"
                      onChange={handleInfoSettings}
                      fullWidth
                      sx={{
                        height: "52px",
                        "& .MuiFormLabel-root": {
                          color: "#0A3B7A",
                          fontWeight: "bold",
                          backgroundColor: "transparent",
                        },
                        "& .MuiOutlinedInput-root": {
                          "& fieldset": {
                            color: "#0A3B7A",
                            fontWeight: "bold",
                          },
                          "&:hover fieldset": {
                            borderColor: "#0A3B7A",
                          },
                          "&.Mui-focused fieldset": {
                            borderColor: "#0A3B7A",
                          },
                        },
                        "& .MuiInputLabel-shrink": {
                          backgroundColor: "transparent",
                        },
                        "& .MuiInputBase-input": {
                          color: "#0A3B7A",
                          fontWeight: "bold",
                        },
                        marginTop: 2,
                      }}
                    />
                    <TextField
                      id="outlined-basic"
                      label={<Translations text={"admin.change.github"} />}
                      variant="outlined"
                      fullWidth
                      sx={{
                        height: "52px",
                        "& .MuiFormLabel-root": {
                          color: "#0A3B7A",
                          fontWeight: "bold",
                          backgroundColor: "transparent",
                        },
                        "& .MuiOutlinedInput-root": {
                          "& fieldset": {
                            color: "#0A3B7A",
                            fontWeight: "bold",
                          },
                          "&:hover fieldset": {
                            borderColor: "#0A3B7A",
                          },
                          "&.Mui-focused fieldset": {
                            borderColor: "#0A3B7A",
                          },
                        },
                        "& .MuiInputLabel-shrink": {
                          backgroundColor: "transparent",
                        },
                        "& .MuiInputBase-input": {
                          color: "#0A3B7A",
                          fontWeight: "bold",
                        },
                        marginTop: 2,
                      }}
                    />
                    <Button
                      variant="dark"
                      sx={{
                        marginTop: 2,
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
                        <Translations text={"admin.change.button"} />
                      </Typography>
                    </Button>
                  </FormControl>

                  <FormControl
                    fullWidth
                    sx={{
                      marginTop: 4,
                      marginBottom: 4,
                    }}
                  >
                    <TextField
                      id="outlined-basic"
                      label={<Translations text={"admin.old.password"} />}
                      variant="outlined"
                      type={showOldPassword ? "text" : "password"}
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
                        "& .MuiFormLabel-root": {
                          color: "#0A3B7A",
                          fontWeight: "bold",
                          backgroundColor: "transparent",
                        },
                        "& .MuiOutlinedInput-root": {
                          "& fieldset": {
                            color: "#0A3B7A",
                            fontWeight: "bold",
                          },
                          "&:hover fieldset": {
                            borderColor: "#0A3B7A",
                          },
                          "&.Mui-focused fieldset": {
                            borderColor: "#0A3B7A",
                          },
                        },
                        "& .MuiInputLabel-shrink": {
                          backgroundColor: "transparent",
                        },
                        "& .MuiInputBase-input": {
                          color: "#0A3B7A",
                          fontWeight: "bold",
                        },
                        marginTop: 2,
                      }}
                    />
                    <TextField
                      id="outlined-basic"
                      label={<Translations text={"admin.new.password"} />}
                      variant="outlined"
                      type={showNewPassword ? "text" : "password"}
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
                        "& .MuiFormLabel-root": {
                          color: "#0A3B7A",
                          fontWeight: "bold",
                          backgroundColor: "transparent",
                        },
                        "& .MuiOutlinedInput-root": {
                          "& fieldset": {
                            color: "#0A3B7A",
                            fontWeight: "bold",
                          },
                          "&:hover fieldset": {
                            borderColor: "#0A3B7A",
                          },
                          "&.Mui-focused fieldset": {
                            borderColor: "#0A3B7A",
                          },
                        },
                        "& .MuiInputLabel-shrink": {
                          backgroundColor: "transparent",
                        },
                        "& .MuiInputBase-input": {
                          color: "#0A3B7A",
                          fontWeight: "bold",
                        },
                        marginTop: 2,
                      }}
                    />
                    <TextField
                      id="outlined-basic"
                      label={<Translations text={"admin.confirm.password"} />}
                      variant="outlined"
                      type={showConfirmPassword ? "text" : "password"}
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
                        "& .MuiFormLabel-root": {
                          color: "#0A3B7A",
                          fontWeight: "bold",
                          backgroundColor: "transparent",
                        },
                        "& .MuiOutlinedInput-root": {
                          "& fieldset": {
                            color: "#0A3B7A",
                            fontWeight: "bold",
                          },
                          "&:hover fieldset": {
                            borderColor: "#0A3B7A",
                          },
                          "&.Mui-focused fieldset": {
                            borderColor: "#0A3B7A",
                          },
                        },
                        "& .MuiInputLabel-shrink": {
                          backgroundColor: "transparent",
                        },
                        "& .MuiInputBase-input": {
                          color: "#0A3B7A",
                          fontWeight: "bold",
                        },
                        marginTop: 2,
                      }}
                    />
                    <Button
                      variant="dark"
                      sx={{
                        marginTop: 2,
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
                        <Translations text={"admin.password.button"} />
                      </Typography>
                    </Button>
                  </FormControl>
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
