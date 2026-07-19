import { useEffect, useState } from "react";
import {
    getLogs,
    getPods,
    getNamespaces
} from "../api/logs";

function Logs() {

    const [namespaces, setNamespaces] = useState([]);
    const [pods, setPods] = useState([]);

    const [namespace, setNamespace] = useState("default");
    const [pod, setPod] = useState("");

    const [logs, setLogs] = useState("");

    const [search, setSearch] = useState("");

    const loadLogs = async () => {

        if (!pod) return;

        try {

            const data = await getLogs(namespace, pod);

            setLogs(data);

        } catch (err) {

            console.error(err);

        }

    };

    useEffect(() => {

        async function load() {

            const ns = await getNamespaces();

            setNamespaces(ns);

        }

        load();

    }, []);

    useEffect(() => {

        async function loadPods() {

            const p = await getPods(namespace);

            setPods(p);

        }

        loadPods();

    }, [namespace]);

    useEffect(() => {

        if (!pod) return;

        loadLogs();

        const timer = setInterval(loadLogs, 5000);

        return () => clearInterval(timer);

    }, [pod]);

    const filteredLogs = logs
        .split("\n")
        .filter(line =>
            line.toLowerCase()
                .includes(search.toLowerCase()))
        .join("\n");

    const copyLogs = () => {

        navigator.clipboard.writeText(logs);

        alert("Logs copied");

    };

    const downloadLogs = () => {

        const blob = new Blob([logs]);

        const url = URL.createObjectURL(blob);

        const a = document.createElement("a");

        a.href = url;

        a.download = `${pod}.log`;

        a.click();

    };
    console.log("namespaces:", namespaces);
    console.log("pods:", pods);
    console.log("selected pod:", pod);
    console.log("logs:", logs);
    return (

        <div className="container mt-4">

            <div className="row mb-3">

                <div className="col">

                    <select
                        className="form-select"
                        value={namespace}
                        onChange={(e)=>setNamespace(e.target.value)}
                    >

                        {
                            namespaces.map(ns=>(
                                <option key={ns.name} value={ns.name}>
                                    {ns.name}
                                </option>
                            ))
                        }

                    </select>

                </div>

                <div className="col">

                    <select
                        className="form-select"
                        value={pod}
                        onChange={(e)=>setPod(e.target.value)}
                    >

                        <option value="">
                            Select Pod
                        </option>

                        {

                        pods
                            .filter(p => p.namespace === namespace)
                            .map(p => (
                                <option key={p.name} value={p.name}>
                                    {p.name}
                                </option>
                            ))

                        }

                    </select>

                </div>

            </div>

            <div className="d-flex gap-2 mb-3">

                <button
                    className="btn btn-primary"
                    onClick={loadLogs}
                >
                    Refresh
                </button>

                <button
                    className="btn btn-success"
                    onClick={copyLogs}
                >
                    Copy
                </button>

                <button
                    className="btn btn-warning"
                    onClick={downloadLogs}
                >
                    Download
                </button>

            </div>

            <input
                className="form-control mb-3"
                placeholder="Search logs..."
                value={search}
                onChange={(e)=>setSearch(e.target.value)}
            />

            <pre
                style={{
                    background:"#111",
                    color:"#00ff66",
                    padding:"20px",
                    height:"600px",
                    overflow:"auto",
                    borderRadius:"10px",
                    fontSize:"13px"
                }}
            >
                {filteredLogs}
            </pre>

        </div>

    );

}

export default Logs;