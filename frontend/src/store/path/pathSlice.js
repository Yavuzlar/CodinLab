import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";

const initialState = {
  loading: true,
  data: [],
  error: null,
};

export const fetchPathById = createAsyncThunk(
  "path/fetchPathById",
  async (data, { rejectWithValue }) => {
    try {
      const response = await axios({
        method: "GET",
        url: `/api/v1/private/road/path/${data.programmingId}/${data.pathId}`,
        headers: {
          accept: "application/json",
          Language: data.language,
        },
      });
      if (response.status === 200) {
        return response.data;
      }
    } catch (error) {
      return rejectWithValue(error.response?.data?.message || error.message);
    }
  }
);

const pathSlice = createSlice({
  name: "path",
  initialState,
  extraReducers: (builder) => {
    builder
      .addCase(fetchPathById.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchPathById.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
      })
      .addCase(fetchPathById.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      });
  },
});

export default pathSlice.reducer;
