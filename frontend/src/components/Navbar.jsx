import "./Navbar.css";

function Navbar() {
  return (
    <header className="navbar">
      <div className="navbar-left">
        <h2>KubeIntel Dashboard</h2>
      </div>

      <div className="navbar-right">
        <span className="cluster-status">
          🟢 Cluster Connected
        </span>
      </div>
    </header>
  );
}

export default Navbar;