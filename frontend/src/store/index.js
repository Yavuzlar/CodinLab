// ** Toolkit imports
import { configureStore } from '@reduxjs/toolkit'
// ** Reducers
import authSlice from './auth/index.js'
import userSlice from './user/userSlice'
import languageSlice from './language/languageSlice' // Import the language slice

export const store = configureStore({
  reducer: {
    auth: authSlice,
    user : userSlice,
    language : languageSlice // Include the language reducer
  },
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware({
      serializableCheck: false
    })
})
