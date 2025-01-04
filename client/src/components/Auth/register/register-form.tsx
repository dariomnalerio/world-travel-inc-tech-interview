import React, { JSX, useState } from "react";
import {
  combineValidators,
  createValidator,
  predicates,
} from "../../../helpers/validators";
import { Label } from "../../ui/label/label";
import { Input } from "../../ui/input/input";
import { LogIn } from "lucide-react";
import styles from "../auth.module.css";
import { useForm } from "../../../hooks/use-form";
import { useView } from "../../../hooks/use-view";
import { login, register } from "../../../api/auth";

const emailValidator = combineValidators(
  createValidator(predicates.required, "email", "Email is required"),
  createValidator(
    predicates.isEmail,
    "email",
    "Please enter a valid email address"
  )
);

const passwordValidator = combineValidators(
  createValidator(predicates.required, "password", "Password is required"),
  createValidator(
    predicates.minLength(8),
    "password",
    "Password must be at least 8 characters long"
  ),
  createValidator(
    predicates.maxLength(32),
    "password",
    "Password must be at most 32 characters long"
  ),
  createValidator(
    predicates.hasDigit,
    "password",
    "Password must contain a digit"
  ),
  createValidator(
    predicates.hasLowerCase,
    "password",
    "Password must contain a lowercase letter"
  ),
  createValidator(
    predicates.hasUpperCase,
    "password",
    "Password must contain an uppercase letter"
  ),
  createValidator(
    predicates.hasSpecialChar,
    "password",
    "Password must contain a special character"
  )
);

interface RegisterFormValues {
  email: string;
  password: string;
}

type FormProps = {
  customOnSubmit?: (formValues: RegisterFormValues) => void;
};

const Form = ({ customOnSubmit }: FormProps): JSX.Element => {
  const {
    getFieldError,
    handleBlur,
    handleChange,
    handleSubmit,
    isSubmitting,
    touched,
    values,
    reset,
  } = useForm<RegisterFormValues>(
    { email: "", password: "" },
    { validators: { email: emailValidator, password: passwordValidator } }
  );
  const [error, setError] = useState<string>("");
  const { changeView } = useView();

  const handleChangeView = () => changeView("login");

  const defaultOnSubmit = async (formValues: RegisterFormValues) => {
    setError("");
    const { email, password } = formValues;
    const { error } = await register({ email, password });
    if (error) {
      setError(error.message);
      return;
    }
    reset();
    const { error: loginError } = await login({ email, password });
    if (loginError) {
      setError(loginError.message);
      return;
    }
    changeView("home");
  };

  // this allows us to pass a custom submit handler, which is useful for testing and reusability
  const onSubmit = customOnSubmit ?? defaultOnSubmit;

  return (
    <React.Fragment>
      <form className={styles.form} onSubmit={handleSubmit(onSubmit)}>
        <div>
          <Label htmlFor="email">Email</Label>
          <Input
            type="email"
            name="email"
            id="email"
            placeholder="Please enter your email"
            autoComplete="email"
            disabled={isSubmitting}
            value={values.email}
            onChange={handleChange}
            onBlur={handleBlur}
          />
          <span className={styles.error}>
            {touched.email && getFieldError("email")}
          </span>
        </div>
        <div>
          <Label htmlFor="password">Password</Label>
          <Input
            type="password"
            name="password"
            id="password"
            placeholder="Please enter your password"
            disabled={isSubmitting}
            value={values.password}
            onChange={handleChange}
            onBlur={handleBlur}
          />
          <span className={styles.error}>
            {touched.password && getFieldError("password")}
          </span>
        </div>
        <div>
          <button className={styles.actionBtn} type="submit">
            <LogIn size={16} />
            <span>Register</span>
          </button>
          <span className={styles.error}>{error}</span>
        </div>
      </form>
      <p className={styles.authFooterText}>
        Already have an account?{" "}
        <button type="button" onClick={handleChangeView}>
          Login
        </button>
      </p>
    </React.Fragment>
  );
};

export default Form;
