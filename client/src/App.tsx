import React, { JSX } from "react";
import Header from "./components/Layout/header";
import { useView } from "./hooks/use-view";
import { View } from "./types";
import Login from "./components/Auth/login/login";
import Landing from "./components/Landing/landing";
import Register from "./components/Auth/register/register";
import Footer from "./components/Layout/footer";

// function that switches between views
const renderView = (currentView: View): JSX.Element => {
  switch (currentView) {
    case "home":
      return <Landing />;
    case "login":
      return <Login />;
    case "register":
      return <Register />;
    default:
      return <Landing />;
  }
};

function App(): JSX.Element {
  const { currentView } = useView();
  return (
    <React.Fragment>
      <Header title="Placeholder" />
      {renderView(currentView)}
      <Footer />
    </React.Fragment>
  );
}

export default App;
