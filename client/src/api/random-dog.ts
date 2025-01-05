import { API_BASE_URL } from "."
import { ErrorCodes } from "../helpers/errors";
import { ErrorResponse, RandomImageResponse, Result } from "../types";

export async function getRandomDog(): Promise<Result<ErrorResponse, RandomImageResponse>> {
  try {
    const res = await fetch(`${API_BASE_URL}/dog/random`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
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
      data
    }
  } catch (error) {
    console.error("Get random dog error:", error);

    return {
      data: null,
      error: {
        code: "database_error",
        message: ErrorCodes.database_error,
      }
    }
  }
}