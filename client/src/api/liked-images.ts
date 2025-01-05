import { API_BASE_URL } from ".";
import { ErrorCodes } from "../helpers/errors";
import { ErrorResponse, LikeDogImageResponse, Result } from "../types";

export async function likeDogImage(userId: string, imageUrl: string): Promise<Result<ErrorResponse, LikeDogImageResponse>> {
  try {
    const res = await fetch(`${API_BASE_URL}/liked_images/${userId}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ imageURL: imageUrl }),
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

    return {
      error: null,
      data: {
        imageUrl,
        success: true,
      }
    }
  } catch (error) {
    console.error("Like dog image error:", error);

    return {
      data: null,
      error: {
        code: "database_error",
        message: ErrorCodes.database_error,
      }
    }
  }
}

export async function unlikeDogImage(userId: string, imageUrl: string): Promise<Result<ErrorResponse, LikeDogImageResponse>> {
  try {
    const res = await fetch(`${API_BASE_URL}/liked_images/${userId}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ imageURL: imageUrl }),
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

    return {
      error: null,
      data: {
        imageUrl,
        success: true,
      }
    }
  } catch (error) {
    console.error("Unlike dog image error:", error);

    return {
      data: null,
      error: {
        code: "database_error",
        message: ErrorCodes.database_error,
      }
    }
  }
}