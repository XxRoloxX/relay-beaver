import axios from "axios";

const GOOGLE_CLIENT_ID = import.meta.env.GOOGLE_CLIENT_ID;
const GOOGLE_AUTH_URL = 

const googleAxios = axios.create({
  baseURL: "https://www.googleapis.com",
  headers: {
    "Content-Type": "application/json",
  },
});
