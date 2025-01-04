import { useState, useCallback } from "react";
import { View } from "../types";
import { ViewContext } from "../hooks/use-view";

interface ViewProviderProps {
  children: React.ReactNode;
  initialView?: View;
}

const ViewProvider: React.FC<ViewProviderProps> = ({
  children,
  initialView = "home",
}) => {
  const [currentView, setCurrentView] = useState<View>(initialView);

  const changeView = useCallback((view: View) => setCurrentView(view), []);

  return (
    <ViewContext.Provider value={{ currentView, changeView }}>
      {children}
    </ViewContext.Provider>
  );
};

export { ViewProvider };
