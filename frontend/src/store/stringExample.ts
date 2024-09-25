import { createSlice } from "@reduxjs/toolkit";

interface StringExampleState {
    value: string;
}

export const stringExampleSlice = createSlice({
    name: 'stringExample',
    initialState: { value: '' } as StringExampleState,
    reducers: {
        changeString: (state, action) => {
            state.value += action.payload
        }
    }
})

export const {changeString} = stringExampleSlice.actions
export default stringExampleSlice.reducer;