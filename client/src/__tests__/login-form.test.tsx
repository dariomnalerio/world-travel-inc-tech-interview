import { act, fireEvent, render, screen } from "@testing-library/react";
import Form from "../components/Auth/login/form";

describe("Login Form", () => {
  it("renders the form with email and password fields", () => {
    render(<Form />);
    expect(screen.getByLabelText(/email/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/password/i)).toBeInTheDocument();
  });

  it("does not show validation errors when fields are touched and left empty", async () => {
    render(<Form />);

    const emailInput = screen.getByLabelText(/email/i);

    await act(async () => {
      fireEvent.blur(emailInput);
    });

    expect(screen.queryByText(/please enter your email/i)).toBeNull();
  });

  it("shows an error for invalid email format", async () => {
    render(<Form />);

    const emailInput = screen.getByLabelText(/email/i);
    await act(async () => {
      fireEvent.change(emailInput, { target: { value: "invalid-email" } });
      fireEvent.blur(emailInput);
    });

    expect(
      await screen.findByText(/please enter a valid email address/i)
    ).toBeInTheDocument();
  });

  it("shows validation errors when fields are empty on submit", async () => {
    render(<Form />);
    const submitButton = screen.getByRole("button", { name: /login/i });

    await act(async () => {
      fireEvent.click(submitButton);
    });

    expect(await screen.findByText(/email is required/i)).toBeInTheDocument();
    expect(
      await screen.findByText(/password is required/i)
    ).toBeInTheDocument();
  });

  it("submits the form with valid data", async () => {
    const handleSubmit = vi.fn();
    render(<Form customOnSubmit={handleSubmit} />);

    const emailInput = screen.getByLabelText(/email/i);
    const passwordInput = screen.getByLabelText(/password/i);
    const submitButton = screen.getByRole("button", { name: /login/i });

    await act(async () => {
      fireEvent.change(emailInput, { target: { value: "test@example.com" } });
      fireEvent.change(passwordInput, { target: { value: "validPassword1!" } });
      fireEvent.click(submitButton);
    });

    expect(handleSubmit).toHaveBeenCalledWith({
      email: "test@example.com",
      password: "validPassword1!",
    });
  });

  it("resets the form after submission", async () => {
    render(<Form />);
    const emailInput = screen.getByLabelText(/email/i);
    const passwordInput = screen.getByLabelText(/password/i);
    const submitButton = screen.getByRole("button", { name: /login/i });

    await act(async () => {
      fireEvent.change(emailInput, {
        target: { value: "test@example.com" },
      });
      fireEvent.change(passwordInput, {
        target: { value: "validPassword1!" },
      });
      fireEvent.click(submitButton);
    });

    expect(emailInput).toHaveValue("");
    expect(passwordInput).toHaveValue("");
  });
});
