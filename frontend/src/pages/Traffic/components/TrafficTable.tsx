import useTraffic from "../../../hooks/useTraffic";
import { defaultDateFromUnixTimestamp } from "../../../lib/date";
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
            <td>{defaultDateFromUnixTimestamp(request.startTime)}</td>
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
