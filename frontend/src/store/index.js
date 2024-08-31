// ** Toolkit imports
import { configureStore } from '@reduxjs/toolkit'
// ** Reducers
import authSlice from './auth/index.js'
import userSlice from './user/userSlice'
import statisticsSlice from './statistics/statisticsSlice.js'
import languageSlice from './language/languageSlice' // Import the language slice

export const store = configureStore({
  reducer: {
    auth: authSlice,
    user : userSlice, // Include the user reducer
    statistics: statisticsSlice, 
    language : languageSlice // Include the language reducer
  },
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware({
      serializableCheck: false
    })
})
