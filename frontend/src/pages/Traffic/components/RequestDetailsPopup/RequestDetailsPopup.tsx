import { ProxiedRequest } from "../../../../api/proxiedRequestApi";
import Popup from "../../../../components/Popup/Popup";
import { defaultDateFromUnixTimestamp } from "../../../../lib/date";
import "./RequestDetailsPopup.scss";

interface RequestDetailsPopupProps {
  request: ProxiedRequest | null;
  showDetails: boolean;
  setShowDetails: (arg: boolean) => void;
}

const RequestDetailsPopup = ({
  request,
  showDetails,
  setShowDetails,
}: RequestDetailsPopupProps) => {
  return (
    request && (
      <Popup isDisplayed={showDetails} setIsDisplayed={setShowDetails}>
        <div className="request-details">
          <h2 className="request-details__header">Request details</h2>
          <div className="request-details__content">
            <div className="request_details__infos">
              <h3 className="request-details__subheader">Basic info</h3>
              <div className="request-details__basic-info">
                <h4>ID: {request.id}</h4>
                <h4>Host: {request.getHost()}</h4>
                <h4>Protocol: {request.request.protocol}</h4>
                <h4>Path: {request.request.path}</h4>
                <h4>Method: {request.request.method}</h4>
                <h4>Target: {request.target}</h4>
                <h4>Response: {request.response.statusCode}</h4>
                <h4>Date: {defaultDateFromUnixTimestamp(request.startTime)}</h4>
                <h4>Duration: {request.endTime - request.startTime} ms</h4>
              </div>
            </div>
            <div className="request-details__bodies">
              <h3 className="request-details__subheader">Response body</h3>
              <textarea
                className="request-details__body"
                value={request.response.body}
                readOnly
              />
            </div>
          </div>
        </div>
      </Popup>
    )
  );
};
export default RequestDetailsPopup;
