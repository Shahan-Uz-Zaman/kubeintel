import { Routes, Route, Navigate } from "react-router-dom";

import Dashboard from "./pages/Dashboard";
import Navbar from "./components/Navbar";
import Sidebar from "./components/Sidebar";
import Monitoring from "./pages/Monitoring";
import DeploymentManager from "./pages/DeploymentManager";

import "./App.css";

function App() {
  return (
    <div className="app-container">
      <Sidebar />

      <div className="main-content">
        <Navbar />

        <div className="page-content">
          <Routes>
              <Route
                  path="/"
                  element={<Navigate to="/dashboard" replace />}
              />

              <Route
                  path="/dashboard"
                  element={<Dashboard />}
              />

              <Route
                  path="/monitoring"
                  element={<Monitoring />}
              />

              <Route
                  path="/deployments"
                  element={<DeploymentManager />}
              />
          </Routes>
        </div>
      </div>
    </div>
  );
}

export default App;