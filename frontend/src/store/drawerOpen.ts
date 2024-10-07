import { createSlice } from "@reduxjs/toolkit";

interface DrawerOpenState {
    value: boolean;
}

export const drawerOpenSlice = createSlice({
    name: 'drawerOpen',
    initialState: { value: false } as DrawerOpenState,
    reducers: {
        toggleDrawer: (state) => {
            state.value = !state.value;
        }
    }
})

export const {toggleDrawer} = drawerOpenSlice.actions
export default drawerOpenSlice.reducer;
