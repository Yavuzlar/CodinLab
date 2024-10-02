import { createAsyncThunk } from "@reduxjs/toolkit";
import { createSlice } from "@reduxjs/toolkit";
import axios from "axios";

const initialState = {
  loading: false,
  data: [],
  error: false,
};


// this api for get stop code component (stop code component is the last component in the container)
export const getStop = createAsyncThunk(
    "code/getStop",
    async (_,{ rejectWithValue }) => {
        try {
        const item =  localStorage.getItem("containerId");
        const response = await axios({
            method: "GET",
            url: `/api/v1/private/common/stop/${item}`,
            headers: {
              'accept': 'application/json',
            }
        });
        
        console.log("stop");
        if (response.status === 200) {
          localStorage.removeItem('containerId');
            return response.data.data;
        }
        } catch (error) {
        return rejectWithValue(response.message || error.message);
        }
      }
);


const codeSlice = createSlice({
  name: "code",
  initialState: initialState,
  extraReducers: (builder) => {
    builder
    .addCase(getStop.pending, (state) => {
        state.loading = true;
        state.error = false;
    })
    .addCase(getStop.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
    })
    .addCase(getStop.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
    })
    }
});


export default codeSlice.reducer;