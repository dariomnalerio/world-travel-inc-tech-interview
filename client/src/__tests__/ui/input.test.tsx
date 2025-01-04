import { render, screen } from "@testing-library/react";
import { Input } from "../../components/ui/input/input";

describe("Input", () => {
  it("renders an input", () => {
    render(<Input />);
    const input = screen.getByRole("textbox");
    expect(input).toBeInTheDocument();
  });

  it("passes props through to the <input>", () => {
    const className = "my-class";
    render(<Input className={className} />);
    const input = screen.getByRole("textbox");
    expect(input).toHaveClass(className);
  });

  it("renders custom props", () => {
    const dataTest = "my-data";
    render(<Input data-test={dataTest} />);
    const input = screen.getByRole("textbox");
    expect(input).toHaveAttribute("data-test", dataTest);
  });

  it("does not render children", () => {
    const text = "Some Text";
    const consoleErrorSpy = vi
      .spyOn(console, "error")
      .mockImplementation(() => {});

    expect(() => render(<Input>{text}</Input>)).toThrow();
    consoleErrorSpy.mockClear();
  });
});
