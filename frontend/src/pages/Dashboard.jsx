import { useEffect, useState } from "react";

import DashboardCard from "../components/DashboardCard";
import { getDashboard } from "../services/api";

import "./Dashboard.css";

function Dashboard() {
  const [dashboard, setDashboard] = useState({
    clusterStatus: "Loading...",
    nodeCount: 0,
    namespaceCount: 0,
    runningPods: 0,
    failedPods: 0,
  });

  const [loading, setLoading] = useState(true);

  const [error, setError] = useState("");

  useEffect(() => {
    loadDashboard();
  }, []);

  async function loadDashboard() {
    try {
      setLoading(true);

      const response = await getDashboard();
      console.log("Dashboard API Response:", response);
      console.log("Dashboard Data:", response.data);

      setDashboard(response.data);

      setError("");
    } catch (err) {
      console.error(err);

      setError("Unable to connect to backend");

    } finally {
      setLoading(false);
    }
  }

  if (loading) {
    return (
      <div className="dashboard-loading">
        Loading Dashboard...
      </div>
    );
  }

  if (error) {
    return (
      <div className="dashboard-error">
        {error}
      </div>
    );
  }

  return (
    <div className="dashboard">

      <h1>Kubernetes Cluster Dashboard</h1>

      <div className="dashboard-grid">

        <DashboardCard
          title="Total Nodes"
          value={dashboard.nodeCount}
          icon="🖥️"
          color="#2563eb"
        />

        <DashboardCard
          title="Namespaces"
          value={dashboard.namespaceCount}
          icon="📂"
          color="#7c3aed"
        />

        <DashboardCard
          title="Running Pods"
          value={dashboard.runningPods}
          icon="📦"
          color="#22c55e"
        />

        <DashboardCard
          title="Failed Pods"
          value={dashboard.failedPods}
          icon="❌"
          color="#ef4444"
        />

        <DashboardCard
          title="Cluster Status"
          value={dashboard.clusterStatus}
          icon="☸️"
          color={
            dashboard.clusterStatus === "Healthy"
              ? "#22c55e"
              : "#f59e0b"
          }
        />

      </div>
    </div>
  );
}

export default Dashboard;