import { createAsyncThunk } from "@reduxjs/toolkit";
import { createSlice } from "@reduxjs/toolkit";
import axios from "axios";
import { showToast } from "src/utils/showToast";
import { getAdminUser } from "../user/userSlice";


const initialState = {
  loading: false,
  data: [],
  error: false,
};

export const fetchUserById = createAsyncThunk(
  "admin/fetchUserById",
  async (data, { rejectWithValue }) => {
    try {
      const response = await axios({
        method: "GET",
        url: `/api/v1/private/admin/user/${data}`,
      });
      if (response.status === 200) {
        return response.data.data;
      }
      
     
    } catch (error) {
      return rejectWithValue(error.response.data.message || error.message);
    }
  }
);

export const updateUserById = createAsyncThunk(
    "admin/updateUserById",
    async ({ data, userid }, { rejectWithValue,dispatch }) => {
      try {
        const response = await axios({
          method: "POST",
          url: `/api/v1/private/admin/user/${userid}`,
          data: data,
        });
        if (response.status === 200) {
          dispatch(getAdminUser());
          return response.data.data;
        }
      } catch (error) {
        return rejectWithValue(error.response.data.message || error.message);
      }
    }
  );

  export const deleteUserById = createAsyncThunk(
    "admin/deleteUserById",
    async (data, { rejectWithValue,dispatch }) => {
      try {
        const response = await axios({
          method: "DELETE",
          url: `/api/v1/private/admin/user/${data}`,
        });
        if (response.status === 200) {
          dispatch(getAdminUser());
          return response.data.data;
        }
      } catch (error) {
        return rejectWithValue(error.response.data.message || error.message);
      }
    }
  );

const adminSlice = createSlice({
  name: "admin",
  initialState: initialState,
  extraReducers: (builder) => {
    builder
      .addCase(fetchUserById.pending, (state) => {
        state.loading = true;
        state.error = false;
      })
      .addCase(fetchUserById.fulfilled, (state, action) => {
        state.data = action.payload;
        state.loading = false;
      })
      .addCase(fetchUserById.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
      })
      .addCase(updateUserById.pending, (state) => {
        state.loading = true;
        state.error = false;
      })
      .addCase(updateUserById.fulfilled, (state, action) => {
        state.loading = false;
        state.data = action.payload
        showToast("dismiss");
        showToast("success", action.payload?.message);
      })
      .addCase(updateUserById.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
        showToast("dismiss");
        showToast("error", state.error);
      })
      .addCase(deleteUserById.pending, (state) => {
        state.loading = true;
        state.error = false;
      })
      .addCase(deleteUserById.fulfilled, (state, action) => {
        state.loading = false;
        state.data = action.payload;
        showToast("dismiss");
        showToast("success", action.payload?.message);
      })
      .addCase(deleteUserById.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload;
        showToast("dismiss");
        showToast("error", state.error);
      });
  },
});

export default adminSlice.reducer;
