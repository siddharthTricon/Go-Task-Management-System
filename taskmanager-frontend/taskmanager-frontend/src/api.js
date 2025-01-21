import axios from "axios";

const API = axios.create({
  baseURL: "http://localhost:8080", // Replace with your backend URL
});

// Add Authorization header for JWT
API.interceptors.request.use((req) => {
  const token = localStorage.getItem("token");
  if (token) {
    req.headers.Authorization = `Bearer ${token}`;
  }
  return req;
});

export const loginUser = (credentials) => API.post("/login", credentials);
export const registerUser = (userData) => API.post("/register", userData);
export const fetchTasks = () => API.get("/tasks");
export const createTask = (taskData) => API.post("/tasks", taskData);
export const deleteTask = (taskId) => API.delete(`/tasks/${taskId}`);
