import { createRef, forwardRef, useEffect } from "react";
import "./HostSelector.scss";

interface HostPopupProps {
  hosts: string[];
  setHost: (arg: string) => void;
  host: string | null;
}

export const HostPopup = forwardRef<HTMLDivElement, HostPopupProps>(
  ({ hosts, setHost, host }, ref) => {
    return (
      <div ref={ref} className="host-popup">
        <h2 className="host-popup__header">Select host</h2>
        <div className="host-popup__content">
          {hosts.map((hostEntry: string) => (
            <div
              key={hostEntry}
              className={`host-popup__entry${host === hostEntry ? "--selected" : ""}`}
              onClick={() => setHost(hostEntry)}
            >
              {hostEntry}
            </div>
          ))}
        </div>
      </div>
    );
  },
);

export const HostSelector = ({
  host,
  hosts,
  setHost,
  showHosts,
  setShowHosts,
}: {
  host: string | null;
  hosts: string[] | null;
  setHost: (arg: string) => void;
  showHosts: boolean;
  setShowHosts: (arg: boolean | ((arg: boolean) => boolean)) => void;
}) => {
  const popupRef = createRef<HTMLDivElement>();

  const handleShowHosts = () => {
    setShowHosts((previousState: boolean) => !previousState);
  };

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        popupRef.current &&
        !popupRef.current.contains(event.target as Node)
      ) {
        setShowHosts(false);
      }
    };
    const handleEscape = (event: KeyboardEvent) => {
      if (event.key === "Escape") {
        setShowHosts(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    document.addEventListener("keydown", handleEscape);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
      document.removeEventListener("keydown", handleEscape);
    };
  });

  useEffect(() => {
    if (host) {
      setShowHosts(false);
    }
  }, [host]);

  useEffect(() => {
    if (showHosts) {
      popupRef.current?.focus();
    }
  }, [showHosts, popupRef]);

  return (
    <>
      <div className="host-selector" onClick={handleShowHosts}>
        <h2>{host}</h2>
      </div>
      <>
        {showHosts && (
          <HostPopup
            host={host}
            ref={popupRef}
            hosts={hosts!}
            setHost={setHost}
          />
        )}
      </>
    </>
  );
};
