import { createAsyncThunk } from "@reduxjs/toolkit";
import { createSlice } from "@reduxjs/toolkit";
import axios from "axios";

const initialState = { //initial state
  loading: false,
  data: [],
  error: false,
};

export const fetchProfileUser = createAsyncThunk(
  "user/fetchProfileUser",
  async (_, { rejectWithValue }) => {
    try {
      const response = await axios({ //this is the for the api call
        url: "/api/v1/private/user/",
        method: "GET",
      });
      if (response.status === 200) {
        return response.data; //return the data
      }
    } catch (error) {
      return rejectWithValue(response.message); //return the error message
    }
  }
);

const userSlice = createSlice({
  name: "user",
  initialState: initialState,
  extraReducers: (builder) => {
    builder
    .addCase(fetchProfileUser.pending, (state) => {  //pending is loading time
      state.loading = true; //set loading to true
      state.error = null; //set error to null
    })
    .addCase(fetchProfileUser.fulfilled, (state, action) => { //fulfilled is success time
      state.data = action.payload; //set data to payload
      state.loading = false; //set loading to false
    })
    .addCase(fetchProfileUser.rejected, (state, action) => { //rejected is error time
      state.loading = false; //set loading
      state.error = action.payload; //set error to payload
    }) 
    }
});

export default userSlice.reducer;
