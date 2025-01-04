import { ErrorCodeKey, ErrorCodes } from "../helpers/errors";

export interface AuthCredentials {
  email: string;
  password: string;
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

export interface ErrorResponse {
  code: keyof typeof ErrorCodes;
  message: ErrorCodeKey;
}

export type Result<K, T> = | { data: T; error: null; } | { data: null; error: K; }