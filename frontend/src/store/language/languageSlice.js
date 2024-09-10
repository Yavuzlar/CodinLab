import { createAsyncThunk } from "@reduxjs/toolkit";
import { createSlice } from "@reduxjs/toolkit";
import axios from "axios";

const initialState = {
  loading: false,
  data: [],
  error: false,
};

export const getInventories = createAsyncThunk(
    "language/getInventories",
  async (_, { rejectWithValue }) => {
    try {
      const response = await axios({ 
        url: "/api/v1/private/home/inventories",
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

export const getUserLanguageLabStats = createAsyncThunk(
  "language/getUserLanguageLabStats",
  async (_, { rejectWithValue }) => {
    try {
      const response = await axios({
        url: "/api/v1/private/labs/general/stats",
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

export const getUserLanguageRoadStats = createAsyncThunk(
  "language/getUserLanguageRoadStats",
  async (_, { rejectWithValue }) => {
    try {
      const response = await axios({
        url: "/api/v1/private/road/general/stats",
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


const languageSlice = createSlice({
  name: "language",
  initialState: initialState,
  extraReducers: (builder) => {
    builder
    .addCase(getInventories.pending, (state) => {  
        state.loading = true; 
        state.error = null; 
      })
      .addCase(getInventories.fulfilled, (state, action) => { 
        state.data = action.payload;
        state.loading = false; 
      })
      .addCase(getInventories.rejected, (state, action) => { 
        state.loading = false; 
        state.error = action.payload; 
      }) 
      .addCase(getUserLanguageLabStats.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(getUserLanguageLabStats.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
      })
      .addCase(getUserLanguageLabStats.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      })
      .addCase(getUserLanguageRoadStats.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(getUserLanguageRoadStats.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
      })
      .addCase(getUserLanguageRoadStats.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      });
      
    }
});

export default languageSlice.reducer;
