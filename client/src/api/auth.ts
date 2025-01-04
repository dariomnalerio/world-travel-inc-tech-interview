import { API_BASE_URL } from "."
import { ErrorCodes } from "../helpers/errors";
import { AuthCredentials, ErrorResponse, LoginResponse, RegisterResponse, Result } from "../types";


export async function login({ email, password }: AuthCredentials): Promise<Result<ErrorResponse, LoginResponse>> {
  try {
    const res = await fetch(`${API_BASE_URL}/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    });


    const data = await res.json();

    if (!res.ok) {
      const errorCode = data.code as keyof typeof ErrorCodes;
      return {
        data: null,
        error: {
          code: errorCode ?? "database_error",
          message: ErrorCodes[errorCode] ?? ErrorCodes.database_error,
        }
      }
    }

    return {
      error: null,
      data,
    }

  } catch (error) {
    console.error("Login error:", error);

    return {
      data: null,
      error: {
        code: "database_error",
        message: ErrorCodes.database_error,
      }
    }
  }
};

export async function register({ email, password }: AuthCredentials): Promise<Result<ErrorResponse, RegisterResponse>> {
  try {
    const res = await fetch(`${API_BASE_URL}/auth/register`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    });

    const data = await res.json();

    if (!res.ok) {
      const errorCode = data.code as keyof typeof ErrorCodes;
      return {
        data: null,
        error: {
          code: errorCode ?? "database_error",
          message: ErrorCodes[errorCode] ?? ErrorCodes.database_error,
        }
      }
    }

    return {
      error: null,
      data,
    }

  } catch (error) {
    console.error("Register error:", error);

    return {
      data: null,
      error: {
        code: "database_error",
        message: ErrorCodes.database_error,
      }
    }
  }
}