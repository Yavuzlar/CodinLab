import { FormControl, Button, TextField } from "@mui/material";
import { loginValidation } from "@/configs/validation/loginScheme";
import { useState, useEffect } from "react";

const { default: BlankLayout } = require("@/layout/BlankLayout");

const Login = () => {
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
    <div>
      <FormControl>
        <TextField
          name="email"
          label="Email"
          variant="outlined"
          onChange={handleChange}
          error={errors.email ? true : false}
          helperText={errors.email}
        />
      </FormControl>
      <FormControl>
        <TextField
          name="password"
          label="Password"
          variant="outlined"
          type="password"
          onChange={handleChange}
          error={errors.password ? true : false}
          helperText={errors.password}
        />
      </FormControl>
      <Button onClick={handleSubmit}>Login</Button>
    </div>
  );
};

Login.guestGuard = true;
Login.getLayout = (page) => <BlankLayout>{page}</BlankLayout>;

export default Login;
