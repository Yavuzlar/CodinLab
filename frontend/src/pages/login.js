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
} from "@mui/material";
// import { GoogleIcon, GitHubIcon } from "@mui/icons-material";
import GoogleIcon from "@mui/icons-material/Google";
import GitHubIcon from "@mui/icons-material/GitHub";
import { useState, useEffect } from "react";
import Translations from "src/components/Translations";
import { loginValidation } from "src/configs/validation/loginSchema";
import VisibilityIcon from "@mui/icons-material/Visibility";
import VisibilityOffIcon from "@mui/icons-material/VisibilityOff";
import rocketImg from "../assets/3d/3d-casual-life-space-white-starship.png";
import Image from "next/image";

const { default: BlankLayout } = require("src/layout/BlankLayout");

const Login = () => {
  const [email, setEmail] = useState();
  const [password, setPassword] = useState();
  const [showPassword, setShowPassword] = useState(false);
  const handleClickShowPassword = () => setShowPassword(!showPassword);
  // const handleMouseDownPassword = (event) => {
  //   event.preventDefault();
  // };
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
        const errors = await loginValidation(formData);
        setErrors(errors);
      }
    };
    fetchData();
  }, [formData, formSubmit]);

  return (
    <Box
      sx={{
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        height: "100vh",
        position: "relative",
      }}
    >
      <Image
        src={rocketImg}
        style={{
          width: "12.438rem",
          height: "16rem",

          position: "absolute",
          left: 1267,
          top: 62,
          transform: "rotate(-52deg)",
        }}
      />
      <Card
        sx={{
          width: "50.625rem",
          height: "50rem",
        }}
      >
        <CardContent
          sx={{
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            flexDirection: "column",
          }}
        >
          <Box sx={{ display: "flex", marginTop: "8.438rem" }}>
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
              CodinLab
            </Typography>
          </Box>
          <FormControl sx={{ mt: "3.125rem" }}>
            <TextField
              name="email"
              label="Email"
              variant="outlined"
              onChange={handleChange}
              error={errors.email ? true : false}
              helperText={errors.email}
              InputProps={{ style: { color: "#0A3B7A" } }}
              sx={{
                "& .MuiFormLabel-root": {
                  color: "#0A3B7A",
                  fontWeight: "bold",
                },
                "& .MuiOutlinedInput-root": {
                  "& fieldset": { color: "#0A3B7A", fontWeight: "bold" },
                  "&:hover fieldset": {},
                  "&.Mui-focused": {
                    color: "#0A3B7A",
                    fontWeight: "bold",
                  },
                },

                width: "40.313rem",
                height: "3.125rem",
              }}
            />
          </FormControl>
          <FormControl>
            <TextField
              name="password"
              label="Password"
              variant="outlined"
              type={showPassword ? "text" : "password"}
              onChange={handleChange}
              error={errors.password ? true : false}
              helperText={errors.password}
              InputProps={{
                style: { color: "#0A3B7A" },
                endAdornment: (
                  <InputAdornment position="end">
                    <IconButton onClick={handleClickShowPassword} edge="end">
                      {showPassword ? (
                        <VisibilityIcon />
                      ) : (
                        <VisibilityOffIcon />
                      )}
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
                mt: "1.563rem",
                width: "40.313rem",
                height: "3.125rem",
              }}
            />
            {/* <InputLabel htmlFor="outlined-adornment-password">
              Password
            </InputLabel>
            <OutlinedInput
              id="outlined-adornment-password"
              type={showPassword ? "text" : "password"}
              endAdornment={
                <InputAdornment position="end">
                  <IconButton
                    aria-label="toggle password visibility"
                    onClick={handleClickShowPassword}
                    onMouseDown={handleMouseDownPassword}
                    edge="end"
                  >
                    {showPassword ? <VisibilityOff /> : <Visibility />}
                  </IconButton>
                </InputAdornment>
              }
              label="Password"
            /> */}
          </FormControl>

          <Button
            sx={{ mt: "1.563rem", width: "40.313rem", height: "3.125rem" }}
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
          <Divider sx={{ mt: "1.563rem", width: "38.75rem" }}>
            <Translations text={"login.divider.or"} />
          </Divider>
          <Box sx={{ mt: "1.563rem" }}>
            <Button
              variant="dark"
              sx={{
                width: "3.125rem",
                height: "3.125rem",
                borderRadius: "0.938rem",
              }}
            >
              <GoogleIcon
                sx={{
                  width: "1.875rem",
                  height: "1.875rem",
                }}
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
              <GitHubIcon
                sx={{
                  width: "1.875rem",
                  height: "1.875rem",
                }}
              />
            </Button>
          </Box>
          <Box sx={{ display: "flex", mt: "1.563rem" }}>
            <Typography sx={{ font: "normal normal normal 18px/23px Outfit" }}>
              <Translations text={"login.new.on.platform"} />
            </Typography>
            <Typography
              sx={{
                ml: "0.938rem",
                color: "#0A3B7A",
                font: "normal normal 600 18px/23px Outfit",
              }}
            >
              <Translations text={"login.create.new.account"} />
            </Typography>
          </Box>
        </CardContent>
      </Card>
    </Box>
  );
};

Login.guestGuard = true;
Login.getLayout = (page) => <BlankLayout>{page}</BlankLayout>;

export default Login;
