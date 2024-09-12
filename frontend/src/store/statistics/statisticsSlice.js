import { createAsyncThunk } from "@reduxjs/toolkit";
import { createSlice } from "@reduxjs/toolkit";
import axios from "axios";

const initialState = {
  loading: false,
  data: [],
  advancementData: [],
  difficultyStatsData: [],
  labsProgressStatsData: [],
  error: false,
};

export const fetchAdvancement = createAsyncThunk(
  "advancementStatistics/fetchAdvancement",
  async (_, { rejectWithValue }) => {
    try {
      const response = await axios({
        url: "/api/v1/private/home/advancement",
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

export const getUserDevelopment = createAsyncThunk(
  "statistics/getUserDevelopment",
  async (_, { rejectWithValue }) => {
    try {
      const response = await axios({
        url: "/api/v1/private/home/development",
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

export const getDifficultyStatistics = createAsyncThunk(
  "statistics/getDifficultyStatistics",
  async (_, { rejectWithValue }) => {
    try {
      const response = await axios({
        url: "/api/v1/private/labs/difficulty/stats",
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

export const getLabsProgressStats = createAsyncThunk(
  "statistics/getLabsProgressStats",
  async (_, { rejectWithValue }) => {
    try {
      const response = await axios({
        url: "/api/v1/private/labs/progress/stats",
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
      state.advancementData = action.payload;
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
      })
      .addCase(getUserDevelopment.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(getUserDevelopment.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
      })
      .addCase(getUserDevelopment.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      })
      .addCase(getDifficultyStatistics.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(getDifficultyStatistics.fulfilled, (state, action) => {
        state.difficultyStatsData = action.payload;
        state.loading = false;
      })
      .addCase(getDifficultyStatistics.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      })
      .addCase(getLabsProgressStats.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(getLabsProgressStats.fulfilled, (state, action) => {
        state.labsProgressStatsData = action.payload;
        state.loading = false;
      })
      .addCase(getLabsProgressStats.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      });
  },
});

export default statisticsSlice.reducer;
