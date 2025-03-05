import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";

const initialState = {
  loading: false,
  data: [],
  weekData: [],
  error: false,
};

export const getLanguageUsageRates = createAsyncThunk(
  "log/getLanguageUsageRates",
  async (_, { rejectWithValue }) => {
    try {
      const response = await axios({
        method: "GET",
        url: `/api/v1/private/log/rates`,
      });
      if (response.status === 200) {
        return response.data;
      }
    } catch (error) {
      return rejectWithValue(response.message || error.message);
    }
  }
);

export const getSolitionWeek = createAsyncThunk(
  "log/getSolitionWeek",
  async (_, { rejectWithValue }) => {
    try {
      const response = await axios({
        method: "GET",
        url: `/api/v1/private/log/solution/week`,
      });
      if (response.status === 200) {
        return response.data;
      }
    } catch (error) {
      return rejectWithValue(response.message || error.message);
    }
  }
);

const logSlice = createSlice({
  name: "log",
  initialState: initialState,
  extraReducers: (builder) => {
    builder
      .addCase(getLanguageUsageRates.pending, (state) => {
        state.loading = true;
        state.error = false;
      })
      .addCase(getLanguageUsageRates.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
      })
      .addCase(getLanguageUsageRates.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      })
      .addCase(getSolitionWeek.pending, (state) => {
        state.loading = true;
        state.error = false;
      })
      .addCase(getSolitionWeek.fulfilled, (state, action) => {
        state.weekData = action.payload;
        state.loading = false;
      })
      .addCase(getSolitionWeek.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      });
  },
});

export default logSlice.reducer;
