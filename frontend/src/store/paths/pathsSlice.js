import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";
import Router from "next/router";

const initialState = {
  loading: false,
  data: [],
  error: null,
};

export const fetchPaths = createAsyncThunk(
  "road/fetchRoad",
  async (data, { rejectWithValue }) => {
    try {
      const response = await axios({
        method: "GET",
        url: `/api/v1/private/road/${data.programmingid}`,
        headers: {
          accept: 'application/json',
          Language: data.language,
        }
      });
      if (response.status === 200) {
        return response.data.data;
      }
    } catch (error) {
      return rejectWithValue(error.response?.data?.message || error.message);
    }
  }
);

export const startRoad = createAsyncThunk(
  "road/startRoad",
  async (data, { rejectWithValue }) => {
    try {
      const response = await axios({
        method: "POST",
        url: "/api/v1/private/road/start",
        data: data,
        headers: {
          "Content-Type": "application/json",
        },
      });
      return response.data;
    } catch (error) {
      return rejectWithValue(error.response?.data?.message || error.message);
    }
  }
);

const pathsSlice = createSlice({
  name: "paths",
  initialState,
  extraReducers: (builder) => {
    builder
      .addCase(fetchPaths.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchPaths.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
      })
      .addCase(fetchPaths.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      }) // Start road
      .addCase(startRoad.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(startRoad.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
      })
      .addCase(startRoad.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      })
  }
});

export default pathsSlice.reducer;