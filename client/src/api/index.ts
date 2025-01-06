const ENV = import.meta.env.MODE
export const API_BASE_URL = ENV != "production" ? "/api" : "/api/v1"