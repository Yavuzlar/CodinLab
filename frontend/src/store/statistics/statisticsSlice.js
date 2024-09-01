import { createAsyncThunk } from "@reduxjs/toolkit";
import { createSlice } from "@reduxjs/toolkit";
import axios from "axios";

const initialState = {
  loading: false,
  data: [],
  error: false,
};

export const fetchAdvancement = createAsyncThunk(
  "statistics/fetchAdvancement",
  async (_, { rejectWithValue }) => {
    try {
      const response = await axios({
        url: "/api/v1/private/home/advancement",
        method: "GET",
      });
      if (response.status === 200) {
        return response.data.data; 
      }
    } catch (error) {
      return rejectWithValue(response.message); 
    }
  }
);

export const GetUserLevel = createAsyncThunk(
  "statistics/GetUserLevel",
  async (_, { rejectWithValue }) => {
    try {
      const response = await axios({
        url: "/api/v1/private/home/level",
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

const statisticsSlice = createSlice({
  name: "statistics",
  initialState: initialState,
  extraReducers: (builder) => {
    builder
      .addCase(fetchAdvancement.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchAdvancement.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
      })
      .addCase(fetchAdvancement.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      })
      .addCase(GetUserLevel.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(GetUserLevel.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
      })
      .addCase(GetUserLevel.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      });
  },
});

export default statisticsSlice.reducer;
