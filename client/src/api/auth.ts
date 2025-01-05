import { API_BASE_URL } from "."
import { ErrorCodes } from "../helpers/errors";
import { AuthCredentials, ErrorResponse, LoginResponse, RegisterResponse, Result, VerifyAuthResponse } from "../types";


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

    const token = data.token;
    const userId = data.userID;

    return {
      error: null,
      data: {
        token,
        userId,
      },
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

export async function verifyAuth(): Promise<Result<ErrorResponse, VerifyAuthResponse>> {
  try {
    const res = await fetch(`${API_BASE_URL}/auth/verify`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
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
    const userId = data.userID;

    return {
      error: null,
      data: {
        userId
      },
    }
  } catch (error) {
    console.error("Verify auth error:", error);

    return {
      data: null,
      error: {
        code: "database_error",
        message: ErrorCodes.database_error,
      }
    }
  }
}