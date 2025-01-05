import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./styles/index.css";
import App from "./App.tsx";
import { ViewProvider } from "./contexts/view-provider.tsx";
import { AuthProvider } from "./contexts/auth-provider.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <AuthProvider>
      <ViewProvider initialView="home">
        <App />
      </ViewProvider>
    </AuthProvider>
  </StrictMode>
);
