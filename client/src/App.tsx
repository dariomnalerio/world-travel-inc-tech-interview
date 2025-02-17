import React, { JSX, useEffect } from "react";
import Header from "./components/Layout/header";
import { useView } from "./hooks/use-view";
import { View } from "./types";
import Login from "./components/Auth/login/login";
import Landing from "./components/Landing/landing";
import Register from "./components/Auth/register/register";
import { useAuth } from "./hooks/use-auth";
import { verifyAuth } from "./api/auth";
import Profile from "./components/Profile/profile";
import { ErrorBoundary } from "react-error-boundary";
import Error from "./components/Error/error";

// function that switches between views
const renderView = (currentView: View): JSX.Element => {
  switch (currentView) {
    case "home":
      return <Landing />;
    case "login":
      return <Login />;
    case "register":
      return <Register />;
    case "profile":
      return <Profile />;
    default:
      return <Landing />;
  }
};

function App(): JSX.Element {
  const { currentView } = useView();
  const { userId, updateUserId } = useAuth();

  useEffect(() => {
    const verifyAuthentication = async () => {
      const { data, error } = await verifyAuth();
      if (error) {
        return;
      }
      updateUserId(data.userId);
    };

    if (!userId) {
      verifyAuthentication();
    }
  }, [updateUserId, userId]);
  return (
    <React.Fragment>
      <Header title="PawPics" />
      <ErrorBoundary fallback={<Error />}>
        {renderView(currentView)}
      </ErrorBoundary>
    </React.Fragment>
  );
}

export default App;
