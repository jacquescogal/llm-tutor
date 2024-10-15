import { configureStore } from '@reduxjs/toolkit';
import drawerOpenReducer from './drawerOpen';
import breadCrumbReducer from './breadcrumbSlice';
import isLoadingReducer from './loaderSlice';
import {
  persistStore,
  persistReducer,
  FLUSH,
  REHYDRATE,
  PAUSE,
  PERSIST,
  PURGE,
  REGISTER,
} from 'redux-persist';
import storage from 'redux-persist/lib/storage'; // Default is localStorage for web
import { combineReducers } from 'redux';

const rootReducer = combineReducers({
  drawerOpen: drawerOpenReducer,
  breadCrumbList: breadCrumbReducer,
  isLoading: isLoadingReducer,
});

const persistConfig = {
  key: 'root', // Key used to persist state
  storage, // Use localStorage
  whitelist: ['breadCrumbList'],
};

const persistedReducer = persistReducer(persistConfig, rootReducer);

export const store = configureStore({
  reducer: persistedReducer,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: {
        ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER],
      },
    }),
});

export const persistor = persistStore(store);
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
