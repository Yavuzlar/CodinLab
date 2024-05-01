// ** Toolkit imports
import { configureStore } from '@reduxjs/toolkit'
// ** Reducers

export const store = configureStore({
  reducer: {
  },
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware({
      serializableCheck: false
    })
})
