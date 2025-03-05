import { createAsyncThunk } from "@reduxjs/toolkit";
import { createSlice } from "@reduxjs/toolkit";
import axios from "axios";
import { Router } from "next/router";

const initialState = {
  loading: false,
  data: [],
  error: false,
};

export const fetchActivityByYear = createAsyncThunk(
  "activity/fetchActivityByYear",
  async ({ year }, { rejectWithValue }) => {
    try {
      const response = await axios({
        method: "GET",
        url: `/api/v1/private/log/solution/byday/${year}`,
        headers: {
          accept: "application/json",
        },
      });
      if (response.status === 200) {
        return response.data;
      }
    } catch (error) {
      return rejectWithValue(response.message);
    }
  }
);

const adminActivitySlice = createSlice({
  name: "adminActivity",
  initialState: initialState,
  extraReducers: (builder) => {
    builder
      .addCase(fetchActivityByYear.pending, (state) => {
        state.loading = true;
        state.error = false;
      })
      .addCase(fetchActivityByYear.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
      })
      .addCase(fetchActivityByYear.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
        state.data = [];
      });
  },
});

export default adminActivitySlice.reducer;
