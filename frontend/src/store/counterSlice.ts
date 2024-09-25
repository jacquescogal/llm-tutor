import { createSlice } from "@reduxjs/toolkit";

// Define a type for the slice state
interface CounterState {
    value: number
}

export const counterSlice = createSlice({
    name: 'counter',
    initialState: { value: 0 } as CounterState,
    reducers: {
        increment: state => {
            state.value += 1
        },
        decrement: state => {
            state.value -= 1
        },
        incrementByAmount: (state, action) => {
            state.value += action.payload
        }
    }
})

// export actions
export const { increment, decrement, incrementByAmount } = counterSlice.actions

// export reducer
export default counterSlice.reducer;