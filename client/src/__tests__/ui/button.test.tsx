import { render, screen } from "@testing-library/react";
import { Button } from "../../components/ui/button/button";

describe("Button", () => {
  it("renders a button", () => {
    render(<Button />);
    const button = screen.getByRole("button");
    expect(button).toBeInTheDocument();
  });

  it("passes children through to the <button>", () => {
    const text = "Some Text";
    render(<Button>{text}</Button>);
    expect(screen.getByText(text)).toBeInTheDocument();
  });

  it("passes props through to the <button>", () => {
    const className = "my-class";
    const dataTest = "my-data";
    render(<Button className={className} data-test={dataTest} />);
    const button = screen.getByRole("button");

    expect(button).toHaveClass(className);
    expect(button).toHaveAttribute("data-test", dataTest);
  });
});
