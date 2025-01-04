/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from "react";
import type { ChangeEvent, FormEvent } from "react";
import type { ValidationError, Validator } from "../helpers/validators";

type FieldValues = Record<string, any>;

type FormState<T extends FieldValues> = {
  values: T;
  errors: ValidationError[];
  touched: Partial<Record<keyof T, boolean>>;
  isSubmitting: boolean;
  isValid: boolean;
};

type UseFormConfig<T extends FieldValues> = {
  validators?: Partial<Record<keyof T, Validator<any>>>;
};

type UseFormReturn<T extends FieldValues> = {
  values: T;
  errors: ValidationError[];
  touched: Partial<Record<keyof T, boolean>>;
  isSubmitting: boolean;
  isValid: boolean;
  getFieldError: (fieldName: keyof T) => string | undefined;
  handleChange: (e: ChangeEvent<HTMLInputElement>) => void;
  handleBlur: (e: ChangeEvent<HTMLInputElement>) => void;
  handleSubmit: (
    onSubmit: (values: T) => void
  ) => (e: FormEvent<HTMLFormElement>) => void;
  setFieldValue: (name: keyof T, value: T[keyof T]) => void;
  reset: () => void;
};

/**
 * Custom hook to manage form state and validation.
 *
 * @template T - The type of the form values.
 * @param {T} initialValues - The initial values of the form fields.
 * @param {UseFormConfig<T>} [config] - Optional configuration for the form, including validators.
 * @returns {UseFormReturn<T>} The form state and handlers.
 */
export function useForm<T extends FieldValues>(
  initialValues: T,
  config?: UseFormConfig<T>
): UseFormReturn<T> {
  const [formState, setFormState] = useState<FormState<T>>({
    values: initialValues,
    errors: [],
    touched: {},
    isSubmitting: false,
    isValid: true,
  });

  const validateField = (
    name: keyof T, // e.g "password"
    value: T[keyof T] // e.g "Password is required"
  ): ValidationError | null => {
    if (!config?.validators || !config.validators[name]) return null;

    const validator = config.validators[name];
    return validator(value);
  };

  // Validate all fields in the form
  const validateForm = (values: T): ValidationError[] => {
    if (!config?.validators) return [];

    const errors: ValidationError[] = [];

    Object.keys(config.validators).forEach((key) => {
      const error = validateField(key, values[key as keyof T]);
      if (error) {
        errors.push(error);
      }
    });

    return errors;
  };

  // Set and validate a field value
  const setFieldValue = (name: keyof T, value: T[keyof T]): void => {
    const error = validateField(name, value);

    setFormState((prev) => {
      const newErrors = prev.errors.filter((e) => e.field !== name);
      if (error) {
        newErrors.push(error);
      }

      return {
        ...prev,
        values: { ...prev.values, [name]: value },
        errors: newErrors,
        isValid: newErrors.length === 0,
      };
    });
  };

  const getFieldError = (name: keyof T): string | undefined => {
    return formState.errors.find((error) => error.field === name)?.message;
  };

  const handleChange = (e: ChangeEvent<HTMLInputElement>): void => {
    const { name, value } = e.target;
    setFieldValue(name as keyof T, value as unknown as T[keyof T]);
  };

  const handleBlur = (e: ChangeEvent<HTMLInputElement>): void => {
    const { name } = e.target;
    setFormState((prev) => ({
      ...prev,
      touched: { ...prev.touched, [name]: true },
    }));
  };

  // Handle form submission.
  // Calls the provided onSubmit function if the form is valid.
  const handleSubmit =
    (onSubmit: (values: T) => Promise<void> | void) => async (e: FormEvent) => {
      e.preventDefault();

      const errors = validateForm(formState.values);
      const isValid = errors.length === 0;

      setFormState((prev) => ({
        ...prev,
        errors,
        isValid,
        touched: Object.keys(prev.values).reduce(
          (acc, key) => ({ ...acc, [key]: true }),
          {}
        ),
      }));

      if (!isValid) return;

      setFormState((prev) => ({ ...prev, isSubmitting: true }));

      try {
        await onSubmit(formState.values);
      } finally {
        setFormState((prev) => ({ ...prev, isSubmitting: false }));
      }
    };

  const reset = (): void => {
    setFormState({
      values: initialValues,
      errors: [],
      touched: {},
      isSubmitting: false,
      isValid: true,
    });
  };

  return {
    ...formState,
    getFieldError,
    handleChange,
    handleBlur,
    handleSubmit,
    setFieldValue,
    reset,
  };
}
