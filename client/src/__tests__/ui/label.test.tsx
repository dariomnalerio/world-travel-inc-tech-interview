import { render, screen } from "@testing-library/react";
import { Label } from "../../components/ui/label/label";

describe("Label", () => {
  it("renders a label", () => {
    render(<Label />);
    const label = screen.getByTestId("label");
    expect(label).toBeInTheDocument();
  });

  it("passes children through to the <label>", () => {
    const text = "Some Text";
    render(<Label>{text}</Label>);
    expect(screen.getByText(text)).toBeInTheDocument();
  });

  it("passes props through to the <label>", () => {
    const className = "my-class";
    render(<Label className={className} />);
    const label = screen.getByTestId("label");
    expect(label).toHaveClass(className);
  });

  it("passes custom props through to the <label>", () => {
    const dataTest = "my-data";
    render(<Label data-test={dataTest} />);
    const label = screen.getByTestId("label");
    expect(label).toHaveAttribute("data-test", dataTest);
  });

  it("associates the label with an input", () => {
    const inputId = "my-input";
    render(
      <>
        <Label htmlFor={inputId}>My Label</Label>
        <input id={inputId} />
      </>
    );
    const label = screen.getByText("My Label");
    const input = screen.getByLabelText("My Label");
    expect(label).toBeInTheDocument();
    expect(input).toBeInTheDocument();
  });
});
