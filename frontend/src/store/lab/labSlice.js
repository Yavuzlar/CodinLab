import { createAsyncThunk } from "@reduxjs/toolkit";
import { createSlice } from "@reduxjs/toolkit";
import axios from "axios";

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
        return rejectWithValue(response.message || error.message);
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
    }
});


export default labSlice.reducer;