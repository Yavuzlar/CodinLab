// ** Toolkit imports
import { configureStore } from '@reduxjs/toolkit'
// ** Reducers
import authSlice from './auth/index.js'
import userSlice from './user/userSlice'
import statisticsSlice from './statistics/statisticsSlice.js'
import languageSlice from './language/languageSlice' 

export const store = configureStore({
  reducer: {
    auth: authSlice,
    user : userSlice, 
    statistics: statisticsSlice, 
    language : languageSlice,
  },
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware({
      serializableCheck: false
    })
})
