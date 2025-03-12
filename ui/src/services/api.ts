import axios from "axios";

const API_BASE_URL = "http://localhost:8088";

export const api = axios.create({
  baseURL: API_BASE_URL,
});
