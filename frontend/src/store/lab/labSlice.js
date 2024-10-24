import { createAsyncThunk } from "@reduxjs/toolkit";
import { createSlice } from "@reduxjs/toolkit";
import axios from "axios";
import Router, { useRouter } from "next/router";

const initialState = {
  loading: false,
  data: [],
  error: false,
};

export const getLabByProgramingId = createAsyncThunk(
  "lab/getLabByProgramingId",
  async (data, { rejectWithValue }) => {
    try {
      const response = await axios({
        method: "GET",
        url: `/api/v1/private/lab/${data.labID}?programmingID=${data.programmingID}`,
        headers: {
          'accept': 'application/json',
          'Language': data.language,
        }
      });
      if (response.status === 200) {
        return response.data.data;
      }
    } catch (error) {
      if (error.response && error.response.status === 404) {
        Router.push("/404");
      }
      return rejectWithValue(response.message || error.message);
    }
  }
);

export const getLabsById = createAsyncThunk(
  "lab/getLabsById",
  async (data, { rejectWithValue }) => {
    try {
      console.log("data", data);
      const response = await axios({
        method: "GET",
        url: `/api/v1/private/labs/${data.programmingID}`,
        headers: {
          'accept': 'application/json',
          'Language': data.language,
        }
      });
      if (response.status === 200) {
        return response.data.data;
      }
    } catch (error) {
      if (error.response && error.response.status === 404) {
        Router.push("/404");
      }
      return rejectWithValue(error.response?.data?.message || error.message);
    }
  }
);

const labSlice = createSlice({
  name: "lab",
  initialState: initialState,
  extraReducers: (builder) => {
    builder
      .addCase(getLabByProgramingId.pending, (state) => {
        state.loading = true;
        state.error = false;
      })
      .addCase(getLabByProgramingId.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
      })
      .addCase(getLabByProgramingId.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      })
      .addCase(getLabsById.pending, (state) => {
        state.loading = true;
        state.error = false;
      })
      .addCase(getLabsById.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
      })
      .addCase(getLabsById.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      }
      );
  }
});


export default labSlice.reducer;