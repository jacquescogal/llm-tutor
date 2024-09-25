import { configureStore } from '@reduxjs/toolkit'
import counterReducer from './counterSlice';
import stringReducer from './stringExample';

export const store = configureStore({
    reducer: {
        counter: counterReducer,
        hmmInteresting: stringReducer
    },
})


export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;