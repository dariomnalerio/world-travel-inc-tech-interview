import { ErrorCodeKey, ErrorCodes } from "../helpers/errors";

/**
 * Transforms the error response from the server into a more readable format.
 * 
 *  Message is a string that describes the error, suitable for displaying to the user.
 */
export interface ErrorResponse {
  code: keyof typeof ErrorCodes;
  message: ErrorCodeKey;
}

/**
 * Represents the result of an operation that can either succeed or fail.
 *
 * @template K - The type of the error.
 * @template T - The type of the data.
 *
 * @property {T | null} data - The data returned by the operation if it succeeds, otherwise null.
 * @property {K | null} error - The error returned by the operation if it fails, otherwise null.
 */
export type Result<K, T> = | { data: T; error: null; } | { data: null; error: K; }

export interface AuthCredentials {
  email: string;
  password: string;
}

export interface RandomImageResponse {
  image_url: string
}

export interface LoginResponse {
  token?: string;
  user?: {
    id: string;
    email: string;
  }
}

export interface RegisterResponse {
  email: string;
  id: string;
  updated_at: string;
  created_at: string;
}

