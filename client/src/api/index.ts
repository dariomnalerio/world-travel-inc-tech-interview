const env = import.meta.env.MODE
export const API_BASE_URL = env === "development" ? "/api" : "http://server:8080/api/v1"