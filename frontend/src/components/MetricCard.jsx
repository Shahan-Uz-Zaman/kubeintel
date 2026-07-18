function MetricCard({ title, value, color }) {

    return (

        <div className="col-md-3 mb-3">

            <div
                className="card shadow"

                style={{
                    borderLeft: `6px solid ${color}`
                }}

            >

                <div className="card-body">

                    <h6>{title}</h6>

                    <h3>{value}</h3>

                </div>

            </div>

        </div>

    );

}

export default MetricCard;