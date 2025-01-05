import React, { JSX, useEffect } from "react";
import Header from "./components/Layout/header";
import { useView } from "./hooks/use-view";
import { View } from "./types";
import Login from "./components/Auth/login/login";
import Landing from "./components/Landing/landing";
import Footer from "./components/Layout/footer";
import Register from "./components/Auth/register/register";
import { useAuth } from "./hooks/use-auth";
import { verifyAuth } from "./api/auth";

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
      return <div>esto es el perfil</div>;
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
      <Header title="Placeholder" />
      {renderView(currentView)}
      <Footer />
    </React.Fragment>
  );
}

export default App;
