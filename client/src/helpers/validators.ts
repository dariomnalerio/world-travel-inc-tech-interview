export interface ValidationError {
  field: string;
  message: string;
}

export interface ValidationResult<T> {
  isValid: boolean;
  errors: ValidationError[];
  values: Record<string, T>;
}

/**
 * A type representing a function that validates a value of type `T`.
 *
 * @template T - The type of the value to be validated.
 * @param value - The value to be validated.
 * @returns A `ValidationError` if the validation fails, or `null` if the validation succeeds.
 */
export type Validator<T> = (value: T) => ValidationError | null;

/**
 * Creates a validator function that checks if a value satisfies a given predicate.
 *
 * @template T - The type of the value to be validated.
 * @param predicate - A function that takes a value of type T and returns a boolean indicating whether the value is valid.
 * @param field - The name of the field being validated.
 * @param message - The error message to return if the value is invalid.
 * @returns A validator function that takes a value of type T and returns null if the value is valid, or an object with the field and message if the value is invalid.
 */
export function createValidator<T>(
  predicate: (value: T) => boolean,
  field: string,
  message: string
): Validator<T> {
  return (value: T) => (predicate(value) ? null : { field, message })
}

/**
 * Combines multiple validators into a single validator function.
 * 
 * @template T - The type of the value to be validated.
 * @param {...Validator<T>[]} validators - An array of validator functions to be combined.
 * @returns {Validator<T>} A single validator function that runs each provided validator in sequence.
 * If any validator returns an error, the combined validator will return that error immediately.
 * If all validators pass, the combined validator will return null.
 */
export function combineValidators<T>(
  ...validators: Validator<T>[]
): Validator<T> {
  return (value: T) => {
    for (const validator of validators) {
      const error = validator(value);
      if (error) return error;
    }
    return null
  }
}

/**
 * A collection of predicate functions for validating strings.
 */
export const predicates = {
  required: (value: string) => value !== undefined && value.trim() != "",
  minLength: (min: number) => (value: string) => value.length >= min,
  maxLength: (max: number) => (value: string) => value.length <= max,
  hasUpperCase: (value: string) => /[A-Z]/.test(value),
  hasLowerCase: (value: string) => /[a-z]/.test(value),
  hasDigit: (value: string) => /\d/.test(value),
  hasSpecialChar: (value: string) => /[#?!@$%^&*-]/.test(value),
  isEmail: (value: string) => /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value),
}