import useTraffic from "../../../hooks/useTraffic";
import "./TrafficTable.scss";

const TrafficTable = () => {
  const { traffic } = useTraffic();

  return (
    <table className="table">
      <thead className="table__header">
        <tr className="table__row">
          <th>Date</th>
          <th>Destination</th>
          <th>Method</th>
          <th>Target</th>
          <th>Response</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        {traffic.map((request, index) => (
          <tr key={index} className="table__row">
            <td>{new Date(request.startTime * 1000).toLocaleString()}</td>
            <td>{request.getHost()}</td>
            <td>{request.request.method}</td>
            <td>{request.target}</td>
            <td>{request.response.statusCode}</td>
            <td>
              <button className="table__button">Delete</button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default TrafficTable;
