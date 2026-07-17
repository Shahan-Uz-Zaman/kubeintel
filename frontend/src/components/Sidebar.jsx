import { NavLink } from "react-router-dom";
import "./Sidebar.css";

const menuItems = [
  {
    name: "Dashboard",
    path: "/dashboard",
    icon: "🏠",
  },
  {
    name: "Nodes",
    path: "/nodes",
    icon: "🖥️",
  },
  {
    name: "Pods",
    path: "/pods",
    icon: "📦",
  },
  {
    name: "Deployments",
    path: "/deployments",
    icon: "🚀",
  },
  {
    name: "Monitoring",
    path: "/monitoring",
    icon: "📊",
  },
  {
    name: "Logs",
    path: "/logs",
    icon: "📜",
  },
  {
    name: "Events",
    path: "/events",
    icon: "📅",
  },
  {
    name: "Health",
    path: "/health",
    icon: "❤️",
  },
  {
    name: "Recommendations",
    path: "/recommendations",
    icon: "💡",
  },
  {
    name: "Settings",
    path: "/settings",
    icon: "⚙️",
  },
];

function Sidebar() {
  return (
    <aside className="sidebar">
      <div className="sidebar-header">
        <h2>KubeIntel</h2>
      </div>

      <nav className="sidebar-menu">
        {menuItems.map((item) => (
          <NavLink
            key={item.name}
            to={item.path}
            className={({ isActive }) =>
              isActive ? "menu-item active" : "menu-item"
            }
          >
            <span className="icon">{item.icon}</span>
            <span>{item.name}</span>
          </NavLink>
        ))}
      </nav>
    </aside>
  );
}

export default Sidebar;