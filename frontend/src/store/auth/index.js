import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import axios from "axios";
import { showToast } from "src/utils/showToast";

const initialState = {
  data: [],
  loading: false,
  error: null,
};

// Profile Change API
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
      return response.data;
    } catch (error) {
      return rejectWithValue(error.response.data.message || error.message);
    }
  }
);

// Change Password API
export const changePassword = createAsyncThunk(
  "changePassword",
  async (data, { rejectWithValue }) => {
    try {
      const response = await axios({
        method: "PUT",
        url: "/api/v1/private/user/password",
        data: data,
        headers: {
          "Content-Type": "application/json",
        },
      });
      return response.data;
    } catch (error) {
      return rejectWithValue(error.response.data.message || error.message);
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
        state.error = null;
      })
      .addCase(changeProfile.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
        showToast("dismiss");
        showToast("success", state.data.message);
      })
      .addCase(changeProfile.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
        showToast("dismiss");
        showToast("error", state.error);
      });

    builder
      .addCase(changePassword.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(changePassword.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
        showToast("dismiss");
        showToast("success", state.data.message);
      })
      .addCase(changePassword.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
        showToast("dismiss");
        showToast("error", state.error);
      });
  },
});

export default authSlice.reducer;
