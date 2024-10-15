import { createSlice } from "@reduxjs/toolkit";

interface IsLoadingSlice {
    loadingCount: number;
}

export const isLoadingSlice = createSlice({
    name: 'isLoading',
    initialState: { loadingCount: 0 } as IsLoadingSlice,
    reducers: {
        addLoad: (state) => {
            state.loadingCount++;
        },
        lessLoad: (state) => {
            state.loadingCount--;
        }

    }
})

export const {addLoad, lessLoad} = isLoadingSlice.actions
export default isLoadingSlice.reducer;
