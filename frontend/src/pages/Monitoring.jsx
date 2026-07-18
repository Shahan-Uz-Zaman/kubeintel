import { useEffect, useState } from "react";
import {
  getClusterMetrics,
  getStorage,
  getNetwork
} from "../services/monitoringService";
import MetricCard from "../components/MetricCard";

function Monitoring() {

  const [cluster, setCluster] = useState(null);
  const [loading, setLoading] = useState(true);
  const [search,setSearch]=useState("");
  const [error, setError] = useState("");
  const [storage,setStorage]=useState(null);
  const [network,setNetwork]=useState(null);

  const loadMetrics = async () => {

    try {

      const response = await getClusterMetrics();

      setCluster(response.data);
      const storageRes = await getStorage();
      console.log(storageRes.data);
      setStorage(storageRes.data);

      const networkRes = await getNetwork();
      console.log(networkRes.data);
      setNetwork(networkRes.data);
      setError("");

    } catch (err) {

        console.error(err);

        setError(
            "Unable to connect to backend server."
        );

    } finally {

      setLoading(false);

    }
  };

  useEffect(() => {

    loadMetrics();

    const timer = setInterval(loadMetrics, 5000);

    return () => clearInterval(timer);

  }, []);

  if (loading) {
      return (
          <div
              className="d-flex justify-content-center align-items-center"
              style={{ height: "70vh" }}
          >
              <div className="text-center">

                  <div
                      className="spinner-border text-primary"
                      style={{ width: "4rem", height: "4rem" }}
                      role="status"
                  >
                      <span className="visually-hidden">
                          Loading...
                      </span>
                  </div>

                  <h3 className="mt-4">
                      Loading Cluster Metrics...
                  </h3>

              </div>
          </div>
      );
  }
  if (
      cluster &&
      cluster.nodes.length === 0 &&
      cluster.pods.length === 0
  ) {
      return (

          <div className="container mt-5">

              <div className="alert alert-warning">

                  <h4>No Monitoring Data Available</h4>

                  <p>

                      Kubernetes Metrics Server is running,
                      but no metrics are currently available.

                  </p>

              </div>

          </div>

      );
  }
  return (
    <div className="container mt-4">

      <h2>Resource Monitoring</h2>

      <hr />
      {
          error &&

          <div className="alert alert-danger">

              {error}

          </div>
      }
      <h3>Node Metrics</h3>
      <div className="row mb-4">

        <MetricCard

            title="Nodes"

            value={cluster.nodes.length}

            color="#0d6efd"

        />

        <MetricCard

            title="Pods"

            value={cluster.pods.length}

            color="#198754"

        />

        <MetricCard

            title="CPU"

            value={cluster.nodes[0]?.cpu}

            color="#dc3545"

        />

        <MetricCard

            title="Memory"

            value={cluster.nodes[0]?.memory}

            color="#ffc107"

        />
        <MetricCard

          title="Storage"

          value={`${storage?.availableGB.toFixed(2)} GB`}

          color="#6610f2"

        />

        <MetricCard

          title="Receive"

          value={`${network?.receive.toFixed(2)} MB/s`}

          color="#20c997"

        />

        <MetricCard

          title="Transmit"

          value={`${network?.transmit.toFixed(2)} MB/s`}

          color="#fd7e14"

        />

      </div>
      <div className="table-responsive">

        <table className="table table-bordered table-hover">


          <thead>

            <tr>

              <th>Name</th>
              <th>CPU</th>
              <th>Memory</th>

            </tr>

          </thead>

          <tbody>

            {cluster.nodes.map(node => (

              <tr key={node.name}>

                <td>{node.name}</td>

                <td>

                <div className="progress">

                <div
                className="progress-bar"

                style={{

                width:
                `${Math.min(node.cpuUsage/10,100)}%`

                }}

                >

                {node.cpu}

                </div>

                </div>
                </td>

                <td>

                <div className="progress">

                <div

                className="progress-bar bg-success"

                style={{

                width:
                `${Math.min(node.memoryUsage/40,100)}%`

                }}

                >

                {node.memory}

                </div>

                </div>

                </td>

              </tr>

            ))}

          </tbody>

        </table>

      </div>

      <div className="mb-3">

          <input
              type="text"
              className="form-control"
              placeholder="Search pod by name..."
              value={search}
              onChange={(e) => setSearch(e.target.value)}
          />

      </div>
      <h3>Storage</h3>

      <table className="table">

        <tbody>

          <tr>

            <td>Available Space</td>

            <td>{storage?.availableGB.toFixed(2)} GB</td>

          </tr>

        </tbody>

      </table>
      <h3>Network</h3>

      <table className="table">

        <tbody>

          <tr>

            <td>Receive</td>

            <td>{network?.receive.toFixed(2)} MB/s</td>

          </tr>

          <tr>

            <td>Transmit</td>

            <td>{network?.transmit.toFixed(2)} MB/s</td>

          </tr>

        </tbody>

      </table>
      <h3 className="mt-5">Pod Metrics</h3>
      <div className="table-responsive">

        <table className="table table-bordered table-hover">
        <thead>

          <tr>

            <th>Namespace</th>
            <th>Pod</th>
            <th>CPU</th>
            <th>Memory</th>

          </tr>

        </thead>

        <tbody>
          {cluster.pods

            .filter(p=>

            p.name

            .toLowerCase()

            .includes(search.toLowerCase())

            )

          .map((pod => (
            
            <tr key={pod.name}>

              <td>{pod.namespace}</td>

              <td>{pod.name}</td>

              <td>{pod.cpu}</td>

              <td>{pod.memory}</td>

            </tr>

          )))}

        </tbody>

      </table>

      </div>




    </div>

  );

}

export default Monitoring;