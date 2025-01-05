import { render, screen } from "@testing-library/react";
import { Tooltip } from "../../components/ui/tooltip/tooltip";

describe("Tooltip component", () => {
  it("renders children", () => {
    render(
      <Tooltip text="test">
        <button>button</button>
      </Tooltip>
    );
    const btn = screen.getByRole("button");
    expect(btn).toBeInTheDocument();
  });

  it("renders tooltip text", () => {
    render(
      <Tooltip text="test">
        <button>button</button>
      </Tooltip>
    );
    const tooltipText = screen.getByTestId("tooltipText");
    expect(tooltipText).toBeInTheDocument();
    expect(tooltipText).toHaveTextContent("test");
  });

  it("does not render tooltip text if text is empty", () => {
    render(
      <Tooltip text="">
        <button>button</button>
      </Tooltip>
    );
    const tooltipText = screen.queryByTestId("tooltipText");
    expect(tooltipText).not.toBeInTheDocument();
  });
});
