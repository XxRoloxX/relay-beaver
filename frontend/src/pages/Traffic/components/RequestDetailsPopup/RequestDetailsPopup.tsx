import { useState } from "react";
import { ProxiedRequest } from "../../../../api/proxiedRequestApi";
import Popup from "../../../../components/Popup/Popup";
import { defaultDateFromUnixTimestamp } from "../../../../lib/date";
import "./RequestDetailsPopup.scss";

interface RequestDetailsPopupProps {
  request: ProxiedRequest | null;
  showDetails: boolean;
  setShowDetails: (arg: boolean) => void;
}

enum Tab {
  ResponseBody = "Response body",
  RequestBody = "Request body",
  RequestHeaders = "Request headers",
  ResponseHeaders = "Response headers",
}

const TABS = [
  Tab.ResponseBody,
  Tab.RequestBody,
  Tab.RequestHeaders,
  Tab.ResponseHeaders,
];

const ReponseBody = ({ response }: { response: ProxiedRequest }) => {
  return (
    <>
      <textarea
        className="request-details__body"
        value={response.response.body}
        readOnly
      />
    </>
  );
};

const RequestBody = ({ request }: { request: ProxiedRequest }) => {
  return (
    <>
      <textarea
        className="request-details__body"
        value={request.request.body}
        readOnly
      />
    </>
  );
};

const RequestHeaders = ({ request }: { request: ProxiedRequest }) => {
  return (
    <>
      <textarea
        className="request-details__body"
        value={JSON.stringify(request.request.headers, null, 2)}
        readOnly
      />
    </>
  );
};

const ResponseHeaders = ({ response }: { response: ProxiedRequest }) => {
  return (
    <>
      <textarea
        className="request-details__body"
        value={JSON.stringify(response.response.headers, null, 2)}
        readOnly
      />
    </>
  );
};

const mapTabToComponent = (tab: Tab, request: ProxiedRequest) => {
  switch (tab) {
    case Tab.ResponseBody:
      return <ReponseBody response={request} />;
    case Tab.RequestBody:
      return <RequestBody request={request} />;
    case Tab.RequestHeaders:
      return <RequestHeaders request={request} />;
    case Tab.ResponseHeaders:
      return <ResponseHeaders response={request} />;
  }
};

const RequestDetailsPopup = ({
  request,
  showDetails,
  setShowDetails,
}: RequestDetailsPopupProps) => {
  const [selectedTab, setSelectedTab] = useState(Tab.ResponseBody);

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
              <div className="request-details__tabs">
                {TABS.map((tab) => (
                  <h3
                    key={tab}
                    className={`request-details__tab ${selectedTab === tab
                        ? "request-details__tab--selected"
                        : ""
                      }`}
                    onClick={() => setSelectedTab(tab)}
                  >
                    {tab}
                  </h3>
                ))}
              </div>
              {mapTabToComponent(selectedTab, request)}
            </div>
          </div>
        </div>
      </Popup>
    )
  );
};
export default RequestDetailsPopup;
