import { useEffect, useState } from "react";
import { ProxiedRequest } from "../api/proxiedRequestApi";

// const mockTraffic = Array.from({ length: 1 }, () => new ProxiedRequest());

const useTraffic = () => {
  const [traffic, setTraffic] = useState<ProxiedRequest[]>([]);
  const [socket, setSocket] = useState<WebSocket | null>(null);

  useEffect(() => {
    const socket = new WebSocket(
      import.meta.env.VITE_BACKEND_WS_URL + "/client-events",
    );
    socket.onopen = () => {
      console.log("Socket connected");
    };
    socket.onmessage = (event) => {
      console.log("Message received: ", event.data);
      setTraffic((prev) => {
        try {
          const newRequest = ProxiedRequest.fromJson(event.data);
          return [newRequest, ...prev];
        } catch (e) {
          console.error("Failed to parse message: ", e);
          return prev;
        }
      });
    };
    socket.onclose = () => {
      console.log("Socket closed");
    };

    setSocket(socket);
    return () => {
      socket.close();
    };
  }, []);

  return { traffic, setTraffic, socket };
};

export default useTraffic;
