import React from "react";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";

import Dashboard from "./pages/Dashboard";
import System from "./pages/System";
import Gate from "./pages/Gate";
import Profile from "./pages/Profile";

function App() {
  return (
    <Router>
      <div style={{ display: "flex", height: "100vh" }}>
        {/* SIDEBAR */}
        <div style={{
          width: "220px",
          background: "#1e293b",
          color: "white",
          padding: "20px"
        }}>
          <h2>Gate App</h2>

          <nav style={{ display: "flex", flexDirection: "column", gap: "10px" }}>
            <Link to="/" style={linkStyle}>Dashboard</Link>
            <Link to="/system" style={linkStyle}>System</Link>
            <Link to="/gate" style={linkStyle}>Gate</Link>
            <Link to="/profile" style={linkStyle}>Profile</Link>
          </nav>
        </div>

        {/* CONTENT */}
        <div style={{ flex: 1, padding: "20px", background: "#f1f5f9" }}>
          <Routes>
            <Route path="/" element={<Dashboard />} />
            <Route path="/system" element={<System />} />
            <Route path="/gate" element={<Gate />} />
            <Route path="/profile" element={<Profile />} />
          </Routes>
        </div>
      </div>
    </Router>
  );
}

const linkStyle = {
  color: "white",
  textDecoration: "none",
  padding: "8px",
  borderRadius: "6px",
  background: "#334155"
};

export default App;