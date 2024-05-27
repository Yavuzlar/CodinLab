import { Circle } from "@mui/icons-material";
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
} from "@mui/material";
import { useState, useEffect } from "react";

const Register = () => {
  const [formData, setFormData] = useState(null);
  const [errors, setErrors] = useState({});
  const [formSubmit, setFormSubmit] = useState(false);
  return (
    <Container>
      <Box
        sx={{
          bgcolor: "#3894D0",
          borderRadius: "16px",
          maxWidth: "812px",
        }}>
        <Grid container justifyContent="center" direction="column" px={10}>
          <Box
            sx={{
              display: "flex",
              justifyContent: "center",
              alignItems: "center",
              gap: 1,
              my: 5,
            }}>
            <Circle sx={{ width: 40, height: 40 }} />
            <Typography textAlign="center" variant="h4">
              CodinLab
            </Typography>
          </Box>
          <FormControl>
            <Grid
              container
              direction="column"
              justifyContent="center"
              alignItems="center"
              gap={2}>
              <TextField name="fullname" label="Fullname" fullWidth />
              <TextField name="username" label="Username" fullWidth />
              <TextField name="email" label="E-mail" fullWidth />
              <TextField name="password" label="Password" fullWidth />
              <FormControlLabel
                control={<Checkbox />}
                label={
                  <Typography>
                    I accept the{" "}
                    <Link color={"#fffddd"} href="#">
                      privacy policy & terms
                    </Link>
                  </Typography>
                }
              />
              <Button>Sign Up</Button>
            </Grid>
          </FormControl>
        </Grid>
      </Box>
    </Container>
  );
};

export default Register;
