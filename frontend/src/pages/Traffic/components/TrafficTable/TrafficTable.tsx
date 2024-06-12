import { useEffect, useRef, useState } from "react";
import useTraffic from "../../../../hooks/useTraffic";
import { defaultDateFromUnixTimestamp } from "../../../../lib/date";
import "./TrafficTable.scss";
import { ProxiedRequest } from "../../../../api/proxiedRequestApi";
import Shadow from "../../../../components/Shadow/Shadow";
import RequestDetailsPopup from "../RequestDetailsPopup/RequestDetailsPopup";

const TrafficTable = () => {
  const { traffic } = useTraffic();
  const [showDetails, setShowDetails] = useState(false);
  const [selectedRequest, setSelectedRequest] = useState<ProxiedRequest | null>(
    null,
  );

  const tableBottomRef = useRef<HTMLDivElement>(null);

  const openRequestDetails = (request: ProxiedRequest) => () => {
    setSelectedRequest(request);
    setShowDetails(true);
  };
  const handleRedo = (request: ProxiedRequest) => () => {
    request.redoRequest();
  };

  useEffect(() => {
    tableBottomRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [traffic]);

  return (
    <div className="table">
      <RequestDetailsPopup
        request={selectedRequest!}
        showDetails={showDetails}
        setShowDetails={setShowDetails}
      />
      {showDetails && <Shadow />}
      <table className="table__content">
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
                <button className="table__button" onClick={handleRedo(request)}>
                  Redo
                </button>
                <button
                  className="table__button"
                  onClick={openRequestDetails(request)}
                >
                  Details
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      <div ref={tableBottomRef} />
    </div>
  );
};

export default TrafficTable;
