import { fireEvent, render, screen } from "@testing-library/react";
import Header from "../components/Layout/header";
import * as useViewModule from "../hooks/use-view";
import { ViewProvider } from "../contexts/view-provider";
import { View } from "../types";

describe("Header", () => {
  const renderHeader = (initialView = "home", title?: string) => {
    render(
      <ViewProvider initialView={initialView as View}>
        <Header title={title} />
      </ViewProvider>
    );
  };

  it("renders the header", () => {
    renderHeader();
    const header = screen.getByRole("banner");
    expect(header).toBeInTheDocument();
  });

  it("renders the logo", () => {
    renderHeader();
    const logo = screen.getByTestId("logo");
    expect(logo).toBeInTheDocument();
  });

  it("renders the default title", () => {
    renderHeader();
    const title = screen.getByText("Default Title");
    expect(title).toBeInTheDocument();
  });

  it("renders a custom title", () => {
    const title = "Custom Title";
    renderHeader("home", title);
    const customTitle = screen.getByText(title);
    expect(customTitle).toBeInTheDocument();
  });

  it("renders the navigation buttons", () => {
    renderHeader();
    const loginButton = screen.getByTestId("loginBtn");
    const registerButton = screen.getByTestId("registerBtn");
    expect(loginButton).toBeInTheDocument();
    expect(registerButton).toBeInTheDocument();
  });

  it("calls changeView with 'login' when the login button is clicked", () => {
    const changeViewMock = vi.fn();
    vi.spyOn(useViewModule, "useView").mockReturnValue({
      currentView: "home",
      changeView: changeViewMock,
    });

    renderHeader();
    const loginButton = screen.getByTestId("loginBtn");
    fireEvent.click(loginButton);

    expect(changeViewMock).toHaveBeenCalledWith("login");
  });

  it("calls changeView with 'register' when the register button is clicked", () => {
    const changeViewMock = vi.fn();
    vi.spyOn(useViewModule, "useView").mockReturnValue({
      currentView: "home",
      changeView: changeViewMock,
    });

    renderHeader();
    const registerButton = screen.getByTestId("registerBtn");
    fireEvent.click(registerButton);

    expect(changeViewMock).toHaveBeenCalledWith("register");
  });
});
