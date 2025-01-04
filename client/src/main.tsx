import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./styles/index.css";
import App from "./App.tsx";
import { ViewProvider } from "./contexts/view-provider.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ViewProvider initialView="home">
      <App />
    </ViewProvider>
  </StrictMode>
);
