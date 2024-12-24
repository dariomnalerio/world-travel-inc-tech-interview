import { describe, expect, it } from "vitest";
import { render, screen } from "@testing-library/react";

function TestComponent({ message = "Hello from Test Component" }) {
  return <div data-testid="test-component">{message}</div>;
}

describe("TestComponent", () => {
  it("renders with default message", () => {
    render(<TestComponent />);
    const element = screen.getByTestId("test-component");
    expect(element).toBeInTheDocument();
    expect(element.textContent).toBe("Hello from Test Component");
  });

  it("renders with custom message", () => {
    render(<TestComponent message="Custom message" />);
    const element = screen.getByTestId("test-component");
    expect(element).toBeInTheDocument();
    expect(element.textContent).toBe("Custom message");
  });
});
