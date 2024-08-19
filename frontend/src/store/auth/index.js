//Change Password, Update Profile

import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import axios from "axios";

const initialState = {
  data: [],
  loading: false,
  error: false,
};

//Profile Change API
export const changeProfile = createAsyncThunk(
  "auth/changeProfile",
  async (data, {rejectWithValue}) => {
    try {
      const response = await axios({
        method: "POST",
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
            method: "POST",
            url: "/api/v1/private/user/password",
            data: data,
            headers: {
            "Content-Type": "application/json",
            },
        });
    
        if (!response.ok || response.status !== 200) {
            return rejectWithValue(response.data);
        }
    
        return response.data;
        } catch (error) {
        return rejectWithValue(error.response.data);
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
        })
        .addCase(changeProfile.rejected, (state, action) => {
            state.error = action.payload;
            state.loading = false;
        });

        builder
        .addCase(changePassword.pending, (state) => {
            state.loading = true;
        })
        .addCase(changePassword.fulfilled, (state, action) => {
            state.data = action.payload;
            state.loading = false;
        })
        .addCase(changePassword.rejected, (state, action) => {
            state.error = action.payload;
            state.loading = false;
        });
    }

});

export default authSlice.reducer;
