// ** Toolkit imports
import { configureStore } from '@reduxjs/toolkit'
// ** Reducers
import authSlice from './auth/index.js'
import userSlice from './user/userSlice'
import statisticsSlice from './statistics/statisticsSlice.js'
import languageSlice from './language/languageSlice' 
import pathsSlice from './paths/pathsSlice.js'
import adminSlice from './admin/adminSlice.js'
import pathSlice from "./path/pathSlice.js";
import logSlice from "./log/logSlice.js";

export const store = configureStore({
  reducer: {
    auth: authSlice,
    user : userSlice, 
    statistics: statisticsSlice, 
    language : languageSlice,
    paths : pathsSlice,
    admin: adminSlice,
    path: pathSlice,
    log: logSlice
  },
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware({
      serializableCheck: false
    })
})
