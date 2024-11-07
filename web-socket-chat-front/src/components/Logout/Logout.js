// src/components/Logout.js
import React from "react";
import { useDispatch } from "react-redux";
import { logout } from "../redux/authSlice";
import { logout as apiLogout } from "../../services/api";

function Logout() {
  const dispatch = useDispatch();

  const handleLogout = async () => {
    await apiLogout();
    dispatch(logout());
  };

  return <button onClick={handleLogout}>Logout</button>;
}

export default Logout;
