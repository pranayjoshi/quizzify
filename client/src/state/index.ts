import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  token: null,
  name: null
};

export const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    setToken: (state, action) => {
      state.token = action.payload.token;
      state.name = action.payload.name;
    },
    // setLogin: (state, action) => {
      
    //   state.user = action.payload.user;
    //   state.token = action.payload.token;
    // },
    // setLogout: (state) => {
    //   // console.log("awdada")
    //   state.user = null;
    //   state.token = null;
    // },
    // setPosts: (state, action) => {
    //   state.posts = action.payload.posts;
    // },
  },
});

export const { setToken, } =
  authSlice.actions;
export default authSlice.reducer;