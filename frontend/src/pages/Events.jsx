import { useEffect, useState } from "react";
import { getEvents } from "../api/events";

function Events() {

    const [events, setEvents] = useState([]);
    const [loading, setLoading] = useState(true);
    const [search, setSearch] = useState("");

    const loadEvents = async () => {
        try {
            setLoading(true);
            const data = await getEvents("default");
            setEvents(data);
        } catch (err) {
            console.error(err);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        loadEvents();
    }, []);

    const filtered = events.filter((event) =>
        event.reason.toLowerCase().includes(search.toLowerCase()) ||
        event.object.toLowerCase().includes(search.toLowerCase()) ||
        event.message.toLowerCase().includes(search.toLowerCase())
    );

    return (
        <div className="container mt-4">

            <div className="d-flex justify-content-between align-items-center mb-3">
                <h2>Kubernetes Events</h2>

                <button
                    className="btn btn-primary"
                    onClick={loadEvents}
                >
                    Refresh
                </button>
            </div>

            <input
                className="form-control mb-3"
                placeholder="Search events..."
                value={search}
                onChange={(e) => setSearch(e.target.value)}
            />

            {loading ? (

                <div className="text-center">

                    <div className="spinner-border text-primary"/>

                </div>

            ) : (

                <table className="table table-striped table-hover">

                    <thead className="table-dark">

                    <tr>
                        <th>Type</th>
                        <th>Reason</th>
                        <th>Object</th>
                        <th>Message</th>
                        <th>Time</th>
                    </tr>

                    </thead>

                    <tbody>

                    {filtered.map((event, index) => (

                        <tr key={index}>

                            <td>

                                <span
                                    className={
                                        event.type === "Warning"
                                            ? "badge bg-danger"
                                            : "badge bg-success"
                                    }
                                >
                                    {event.type}
                                </span>

                            </td>

                            <td>{event.reason}</td>

                            <td>{event.object}</td>

                            <td>{event.message}</td>

                            <td>{new Date(event.time).toLocaleString()}</td>

                        </tr>

                    ))}

                    </tbody>

                </table>

            )}

        </div>
    );
}

export default Events;