interface Request {
  method: string;
  protocol: string;
  path: string;
  headers: Header[];
  body: string;
}

interface Header {
  key: string;
  value: string;
}

function isRequest(obj: unknown): obj is Request {
  if (typeof obj !== "object" || obj === null) return false;
  const castedObj = obj as Request;
  return (
    typeof castedObj.method === "string" &&
    typeof castedObj.protocol === "string" &&
    typeof castedObj.path === "string" &&
    Array.isArray(castedObj.headers) &&
    typeof castedObj.body === "string"
  );
}

function isResponse(obj: unknown): obj is Response {
  if (typeof obj !== "object" || obj === null) return false;
  const castedObj = obj as Response;
  return (
    typeof castedObj.statusCode === "number" &&
    typeof castedObj.protocol === "string" &&
    Array.isArray(castedObj.headers) &&
    typeof castedObj.body === "string"
  );
}

interface Response {
  statusCode: number;
  protocol: string;
  headers: Header[];
  body: string;
}

interface IProxiedRequest {
  id: string;
  request: Request;
  response: Response;
  target: string;
  startTime: number;
  endTime: number;
}

export class ProxiedRequest {
  id: string = "";
  request: Request = {
    method: "",
    protocol: "",
    path: "",
    headers: [],
    body: "",
  };
  response: Response = {
    statusCode: 0,
    protocol: "",
    headers: [],
    body: "",
  };
  target: string = "";
  startTime: number = 0;
  endTime: number = 0;

  public getHost() {
    for (const header of this.request.headers) {
      if (header.key.toLowerCase() === "host") {
        return header.value;
      }
    }

    return "";
  }

  public static fromIProxiedRequest(request: IProxiedRequest): ProxiedRequest {
    const result = new ProxiedRequest();
    result.id = request.id;
    result.request = request.request;
    result.response = request.response;
    result.target = request.target;
    result.startTime = request.startTime;
    result.endTime = request.endTime;
    return result;
  }

  public static fromJson(json: string): ProxiedRequest {
    try {
      return tryParseProxiedRequest(json);
    } catch (e) {
      console.error(`Couldn't parse ProxiesRequest from JSON: ${json}`);
      throw e;
    }
  }
}

export function isProxiedRequest(obj: unknown): obj is IProxiedRequest {
  if (typeof obj !== "object" || obj === null) return false;
  const castedObj = obj as ProxiedRequest;
  return (
    isRequest(castedObj.request) &&
    isResponse(castedObj.response) &&
    typeof castedObj.target === "string" &&
    typeof castedObj.startTime === "number" &&
    typeof castedObj.endTime === "number"
  );
}

export function tryParseRequest(json: string): Request {
  let rawObject;

  try {
    rawObject = JSON.parse(json);
  } catch (e) {
    console.error(`Couldn't parse Request from JSON: ${json}`);
    throw e;
  }

  if (typeof rawObject !== "object" || rawObject === null) {
    throw new Error("Request JSON is not an object");
  }
  const castedObject = rawObject as Request;

  if (typeof castedObject.method !== "string") {
    throw new Error("Request JSON has invalid method");
  }
  if (typeof castedObject.protocol !== "string") {
    throw new Error("Request JSON has invalid protocol");
  }
  if (typeof castedObject.path !== "string") {
    throw new Error("Request JSON has invalid path");
  }
  if (!Array.isArray(castedObject.headers)) {
    throw new Error("Request JSON has invalid headers");
  }
  if (typeof castedObject.body !== "string") {
    throw new Error("Request JSON has invalid body");
  }

  return castedObject;
}

export const tryParseResponse = (json: string): Response => {
  const rawObject = JSON.parse(json);
  if (typeof rawObject !== "object" || rawObject === null) {
    throw new Error("Response JSON is not an object");
  }
  const castedObject = rawObject as Response;

  if (typeof castedObject.statusCode !== "number") {
    throw new Error("Response JSON has invalid statusCode");
  }
  if (typeof castedObject.protocol !== "string") {
    throw new Error("Response JSON has invalid protocol");
  }
  if (!Array.isArray(castedObject.headers)) {
    throw new Error("Response JSON has invalid headers");
  }
  if (typeof castedObject.body !== "string") {
    throw new Error("Response JSON has invalid body");
  }

  return castedObject;
};

export function tryParseProxiedRequest(json: string): ProxiedRequest {
  if (typeof json !== "string") {
    throw new Error("ProxiedRequest JSON is not a string");
  }

  const rawObject = JSON.parse(json);

  if (typeof rawObject !== "object" || rawObject === null) {
    throw new Error("ProxiedRequest JSON is not an object");
  }
  const castedObject = rawObject as ProxiedRequest;

  tryParseRequest(JSON.stringify(castedObject.request));

  tryParseResponse(JSON.stringify(castedObject.response));

  if (typeof castedObject.target !== "string") {
    throw new Error("ProxiedRequest JSON has invalid target");
  }
  if (typeof castedObject.startTime !== "number") {
    throw new Error("ProxiedRequest JSON has invalid startTime");
  }
  if (typeof castedObject.endTime !== "number") {
    throw new Error("ProxiedRequest JSON has invalid endTime");
  }

  return ProxiedRequest.fromIProxiedRequest(castedObject);
}
