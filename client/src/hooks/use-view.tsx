import { useState, useCallback } from "react";

type View = "login" | "register" | "home";

type UseViewProps = {
  initialView?: View;
};

const useView = ({ initialView = "home" }: UseViewProps) => {
  const [currentView, setCurrentView] = useState<View>(initialView);

  const changeView = useCallback((view: View) => setCurrentView(view), []);

  return { currentView, changeView };
};

export { useView };
