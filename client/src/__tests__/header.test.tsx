import { fireEvent, render, screen } from "@testing-library/react";
import Header from "../components/Header/header";
import * as useViewModule from "../hooks/use-view";

describe("Header", () => {
  it("renders the header", () => {
    render(<Header />);
    const header = screen.getByRole("banner");
    expect(header).toBeInTheDocument();
  });

  it("renders the logo", () => {
    render(<Header />);
    const logo = screen.getByTestId("logo");
    expect(logo).toBeInTheDocument();
  });

  it("renders the default title", () => {
    render(<Header />);
    const title = screen.getByText("Default Title");
    expect(title).toBeInTheDocument();
  });

  it("renders a custom title", () => {
    const title = "Custom Title";
    render(<Header title={title} />);
    const customTitle = screen.getByText(title);
    expect(customTitle).toBeInTheDocument();
  });

  it("renders the navigation buttons", () => {
    render(<Header />);
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

    render(<Header />);
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

    render(<Header />);
    const registerButton = screen.getByTestId("registerBtn");
    fireEvent.click(registerButton);

    expect(changeViewMock).toHaveBeenCalledWith("register");
  });
});
