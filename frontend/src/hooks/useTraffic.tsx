import { useEffect, useState } from "react";

export class Request {
  public destination: string;
  public method: string;
  public target: string;
  public response: string;

  constructor(
    destination: string = "",
    method: string = "",
    target: string = "",
    response: string = "",
  ) {
    this.destination = destination;
    this.method = method;
    this.target = target;
    this.response = response;
  }

  public static fromJson(json: string): Request {
    const parsed = JSON.parse(json);
    return new Request(
      parsed.destination,
      parsed.method,
      parsed.target,
      parsed.response,
    );
  }
}

const useTraffic = () => {
  const [traffic, setTraffic] = useState<Request[]>([]);
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
      setTraffic((prev) => [...prev, Request.fromJson(event.data)]);
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
