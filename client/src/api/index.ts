const env = import.meta.env.MODE
const API_URL = import.meta.env.VITE_API_URL
export const API_BASE_URL = env !== "production" ? "/api" : API_URL