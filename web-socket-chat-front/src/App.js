// src/App.js
import React, { useState } from "react";
import Login from "./components/Login/Login";
import VideoCall from "./components/VideoCall/VideoCall";
import Logout from "./components/Logout/Logout";

function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  return (
    <div className="App">
      {!isAuthenticated ? (
        <Login onLogin={() => setIsAuthenticated(true)} />
      ) : (
        <>
          <Logout onLogout={() => setIsAuthenticated(false)} />
          <VideoCall />
        </>
      )}
    </div>
  );
}

export default App;
