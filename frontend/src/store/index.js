// ** Toolkit imports
import { configureStore } from '@reduxjs/toolkit'
// ** Reducers
import userSlice from './user/userSlice'

export const store = configureStore({
  reducer: {
    user : userSlice // Include the user reducer
  },
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware({
      serializableCheck: false
    })
})
