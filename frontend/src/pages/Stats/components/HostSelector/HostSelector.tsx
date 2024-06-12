import "./HostSelector.scss";

interface HostSelectorProps {
  hosts: string[];
  setHost: (arg: string) => void;
  host: string | null;
}

export const HostSelector = ({ host, hosts, setHost }: HostSelectorProps) => {
  return (
    <div className="host-popup">
      <h2 className="host-popup__header">Select host</h2>
      <div className="host-popup__content">
        {hosts
          ? hosts.map((hostEntry: string) => (
            <div
              key={hostEntry}
              className={`host-popup__entry${host === hostEntry ? "--selected" : ""}`}
              onClick={() => setHost(hostEntry)}
            >
              {hostEntry}
            </div>
          ))
          : null}
      </div>
    </div>
  );
};
