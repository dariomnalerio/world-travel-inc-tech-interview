import React, { JSX, useState } from "react";
import {
  createValidator,
  combineValidators,
  predicates,
} from "../../../helpers/validators";
import { useForm } from "../../../hooks/use-form";
import styles from "../auth.module.css";
import { Input } from "../../ui/input/input";
import { Label } from "../../ui/label/label";
import { LogIn } from "lucide-react";
import { login } from "../../../api/auth";
import { useView } from "../../../hooks/use-view";
import { useAuth } from "../../../hooks/use-auth";

const emailValidator = combineValidators(
  createValidator(predicates.required, "email", "Email is required"),
  createValidator(
    predicates.isEmail,
    "email",
    "Please enter a valid email address"
  )
);

const passwordValidator = createValidator(
  predicates.required,
  "password",
  "Password is required"
);

interface LoginFormValues {
  email: string;
  password: string;
}

type FormProps = {
  customOnSubmit?: (formValues: LoginFormValues) => void;
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
  } = useForm<LoginFormValues>(
    { email: "", password: "" },
    { validators: { email: emailValidator, password: passwordValidator } }
  );
  const [error, setError] = useState<string>("");
  const { changeView } = useView();
  const { updateUserId } = useAuth();

  const handleChangeView = () => changeView("register");

  const defaultOnSubmit = async (formValues: LoginFormValues) => {
    setError("");
    const { email, password } = formValues;
    const { error, data } = await login({ email, password });
    if (error) {
      setError(error.message);
      return;
    }
    updateUserId(data.userId);
    reset();
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
          <button
            disabled={isSubmitting}
            className={styles.actionBtn}
            type="submit"
          >
            <LogIn size={16} />
            Login
          </button>
          <span className={styles.error}>{error}</span>
        </div>
      </form>
      <p className={styles.authFooterText}>
        Don't have an account?{" "}
        <button type="button" onClick={handleChangeView}>
          Register
        </button>
      </p>
    </React.Fragment>
  );
};

export default Form;
