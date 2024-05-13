import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import axios from 'axios';

const initialState = {
    loading: false,
    values: null,
};

// Fetch data
export const fetchData = createAsyncThunk('change_this/fetchData', async (_, { rejectWithValue }) => {
    try {
        const response = await axios({
            method: 'PUT',
            url: `${process.env.NEXT_PUBLIC_BASE_URL}/`,
        })

        return response.data?.data
    } catch (error) {
        return rejectWithValue(error.response.data.invalid_params)
    }
})

const change_thisSlice = createSlice({
    name: 'change_this',
    initialState,
    reducers: {
    },
    extraReducers: (builder) => {
        builder.addCase(fetchData.pending, (state) => {
            state.loading = true
            // showToast('dismiss')
            // showToast('loading', "")
        })
        builder.addCase(fetchData.fulfilled, (state, action) => {
            state.loading = false
            state.values = action.payload
            // showToast('dismiss')
            // showToast('success', "")
        })
        builder.addCase(fetchData.rejected, (state, action) => {
            state.loading = false
            // showToast('dismiss')
            // showToast('error', "")
        })
    },
});

export const getLoading = (state) => state.change_this.loading;
export const getValues = (state) => state.change_this.values;

export default change_thisSlice.reducer;
