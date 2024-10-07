import { createSlice } from "@reduxjs/toolkit";

interface BreadCrumb {
  name: string;
  path: string;
}

// Define a type for the slice state
interface BreadCrumbState {
  value: BreadCrumb[];
}

export const breadCrumbSlice = createSlice({
  name: "counter",
  initialState: {
    value: [
    ],
  } as BreadCrumbState,
  reducers: {
    pushLink: (state, crumb) => {
      state.value.push(crumb.payload);
    },
    popLink: (state) => {
      state.value.pop();
    },
    clearLinkList: (state) => {
        state.value = [];
    },
    setRootLink: (state, crumb) => {
        state.value = [crumb.payload];
    },
    setLinks: (state, crumbList) => {
      state.value = crumbList.payload;
    },
  },
});

// export actions
export const { pushLink, popLink, clearLinkList, setRootLink , setLinks} = breadCrumbSlice.actions;

// export reducer
export default breadCrumbSlice.reducer;
