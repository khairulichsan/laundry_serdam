// src/api.js
import axios from "axios";

// Set the base URL for your API
const API_URL = "http://127.0.0.1:8080"; // Replace with your backend's address

const api = axios.create({
  baseURL: API_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

export default api;
