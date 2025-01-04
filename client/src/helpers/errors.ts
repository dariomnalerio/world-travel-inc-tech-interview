export const ErrorCodes = {
  // Authentication & User Account
  invalid_email: "Please enter a valid email address",
  email_already_exists: "An account with this email already exists",
  invalid_credentials: "Invalid email or password",
  user_not_found: "Invalid email or password",

  // Session Management
  invalid_token: "Your session has expired. Please sign in again",
  jwt_error: "Your session has expired. Please sign in again",

  // Image Operations
  empty_image_url: "Please provide an image",
  malformed_url: "The provided link is invalid",
  invalid_image_extension: "Unsupported image format. Please use JPG, PNG, or GIF",
  invalid_protocol: "Please use a secure (HTTPS) link",
  image_already_liked: "You've already liked this image",
  image_not_liked: "You haven't liked this image yet",

  // System Errors
  failed_hash: "An error occurred during signup. Please try again",
  database_error: "Something went wrong. Please try again later",
  external_api_error: "Something went wrong. Please try again later",
} as const

export type ErrorCodeKey = typeof ErrorCodes[keyof typeof ErrorCodes];
