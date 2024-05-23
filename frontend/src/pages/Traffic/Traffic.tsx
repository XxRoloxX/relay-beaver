import useTraffic from "../../hooks/useTraffic";

const Traffic = () => {
  const traffic = useTraffic();
  return (
    <div className="protected-route">
      <h1>Current traffic</h1>
    </div>
  );
};
export default Traffic;
