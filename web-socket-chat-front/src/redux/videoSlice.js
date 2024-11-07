// src/redux/videoSlice.js
import { createSlice } from "@reduxjs/toolkit";

const videoSlice = createSlice({
  name: "video",
  initialState: {
    peers: [],
    localStream: null,
  },
  reducers: {
    addPeer(state, action) {
      state.peers.push(action.payload);
    },
    removePeer(state, action) {
      state.peers = state.peers.filter(peer => peer.id !== action.payload.id);
    },
    setLocalStream(state, action) {
      state.localStream = action.payload;
    },
    clearPeers(state) {
      state.peers = [];
    },
  },
});

export const { addPeer, removePeer, setLocalStream, clearPeers } = videoSlice.actions;
export default videoSlice.reducer;
