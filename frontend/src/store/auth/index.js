//Change Password, Update Profile

import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import axios from "axios";
import { showToast } from "src/utils/showToast";

const initialState = {
  data: [],
  loading: false,
  error: false,
};

//Profile Change API
export const changeProfile = createAsyncThunk(
  "auth/changeProfile",
  async (data, { rejectWithValue }) => {
    try {
      const response = await axios({
        method: "PUT",
        url: "/api/v1/private/user/",
        data: data,
        headers: {
          "Content-Type": "application/json",
        },
      });
      console.log("deneme", response);
      if (!response.ok) {
        return rejectWithValue(response.data);
      }
      console.log("deneme", response);
      return response.data;
    } catch (error) {
      return rejectWithValue(error.response.data);
    }
  }
);

//Change Password API

export const changePassword = createAsyncThunk(
  "changePassword",
  async (data, rejectWithValue) => {
    try {
      const response = await axios({
        method: "PUT",
        url: "/api/v1/private/user/password",
        data: data,
        headers: {
          "Content-Type": "application/json",
        },
      });

      if (!response.ok || response.status !== 200) {
        router.push("/");
        return rejectWithValue(response.data);
      }
      return response.data;

    } catch (error) {
      console.log("deneme", response.message);


      return rejectWithValue(response.message);
    }
  }
);

const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(changeProfile.pending, (state) => {
        state.loading = true;
      })
      .addCase(changeProfile.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
        showToast("dismiss");
        showToast("succsess", state.message);
      })
      .addCase(changeProfile.rejected, (state, action) => {
        state.error = action.payload;
        state.loading = false;
        showToast("dismiss");
        showToast("error", state.message);
      });

    builder
      .addCase(changePassword.pending, (state) => {
        state.loading = true;
      })
      .addCase(changePassword.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
        showToast("dismiss");
        showToast("succsess", state.data);
      })
      .addCase(changePassword.rejected, (state, action) => {
        state.error = action.payload;
        // console.log("action", action);
        // console.log("state", state);
        // console.log("state.error", state.error);
        state.loading = false;
        showToast("dismiss");
        showToast("error", action.error.message);
      });
  },
});

export default authSlice.reducer;
