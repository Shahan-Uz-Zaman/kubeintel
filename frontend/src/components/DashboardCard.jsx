import "./DashboardCard.css";

function DashboardCard({ title, value, icon, color }) {
  return (
    <div
      className="dashboard-card"
      style={{
        borderLeft: `6px solid ${color}`,
      }}
    >
      <div className="card-icon">
        {icon}
      </div>

      <div className="card-content">
        <h4>{title}</h4>
        <h2>{value}</h2>
      </div>
    </div>
  );
}

export default DashboardCard;