import { useEffect, useState } from "react";

const useTraffic = () => {
  const [traffic, setTraffic] = useState([]);
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
      // setTraffic(JSON.parse(event.data));
    };

    setSocket(socket);
    return () => {
      socket.close();
    };
  }, []);

  return { traffic };
};

export default useTraffic;
