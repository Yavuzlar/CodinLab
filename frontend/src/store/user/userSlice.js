import { createAsyncThunk } from "@reduxjs/toolkit";
import { createSlice } from "@reduxjs/toolkit";
import axios from "axios";
import { showToast } from "src/utils/showToast";

const initialState = {
  //initial state
  loading: false,
  data: [],
  adminUserData: [],
  error: false,
};

export const fetchProfileUser = createAsyncThunk(
  "user/fetchProfileUser",
  async (_, { rejectWithValue }) => {
    try {
      const response = await axios({
        // This is for the api call
        url: "/api/v1/private/user/",
        method: "GET",
      });
      if (response.status === 200) {
        return response.data; //return the data
      }
    } catch (error) {
      return rejectWithValue(response.message); //return the error message
    }
  }
);

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

export const getAdminUser = createAsyncThunk(
  "user/getAdminUser",
  async (_, { rejectWithValue }) => {
    try {
      const response = await axios({
        url: "/api/v1/private/admin/user",
        method: "GET",
      });
      if (response.status === 200) {
        
        return response.data;
      }
    } catch (error) {
      return rejectWithValue(response.message);
    }
  }
);

const userSlice = createSlice({
  name: "user",
  initialState: initialState,
  extraReducers: (builder) => {
    builder
      .addCase(fetchProfileUser.pending, (state) => {
        //pending is loading time
        state.loading = true; //set loading to true
        state.error = null; //set error to null
      })
      .addCase(fetchProfileUser.fulfilled, (state, action) => {
        //fulfilled is success time
        state.data = action.payload; //set data to payload
        state.loading = false; //set loading to false
      })
      .addCase(fetchProfileUser.rejected, (state, action) => {
        //rejected is error time
        state.loading = false; //set loading
        state.error = action.payload; //set error to payload
      })
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
      })
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
      })
      .addCase(getAdminUser.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(getAdminUser.fulfilled, (state, action) => {
        state.adminUserData = action.payload;
        state.loading = false;
      })
      .addCase(getAdminUser.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      });
  },
});

export default userSlice.reducer;
