import { createContext, use } from "react";
import { View } from "../types";

interface ViewContextProps {
  currentView: View;
  changeView: (view: View) => void;
}

export const ViewContext = createContext<ViewContextProps | undefined>(
  undefined
);

const useView = () => {
  const context = use(ViewContext);

  if (!context) {
    throw new Error("useView must be used within a ViewProvider");
  }

  return context;
};

export { useView };
