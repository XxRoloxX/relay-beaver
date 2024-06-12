import useTraffic from "../../hooks/useTraffic";
import TrafficTable from "./components/TrafficTable/TrafficTable";
import "./Traffic.scss";

const Traffic = () => {
  useTraffic();

  return (
    <div className="traffic-page">
      <h1 className="traffic-page__title">Current traffic</h1>
      <TrafficTable />
    </div>
  );
};
export default Traffic;
