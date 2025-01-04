import React from "react";
import {
  createValidator,
  combineValidators,
  predicates,
} from "../../../helpers/validators";
import { useForm } from "../../../hooks/use-form";
import styles from "./login.module.css";
import { Input } from "../../ui/input/input";
import { Label } from "../../ui/label/label";
import { LogIn } from "lucide-react";

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
const Form = ({
  customOnSubmit,
}: {
  customOnSubmit?: (formValues: LoginFormValues) => void;
}) => {
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
    {
      email: "",
      password: "",
    },
    {
      validators: {
        email: emailValidator,
        password: passwordValidator,
      },
    }
  );

  const defaultOnSubmit = async (formValues: LoginFormValues) => {
    console.log(formValues);
    reset();
  };

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
            value={values.email}
            onChange={handleChange}
            onBlur={handleBlur}
          />
          <span className={styles.error}>
            {touched.email && getFieldError("email")}
          </span>
        </div>
        <div className={styles.item}>
          <Label htmlFor="password">Password</Label>
          <Input
            type="password"
            name="password"
            id="password"
            placeholder="Please enter your password"
            value={values.password}
            onChange={handleChange}
            onBlur={handleBlur}
          />

          <span className={styles.error}>
            {touched.password && getFieldError("password")}
          </span>
        </div>
        <button
          disabled={isSubmitting}
          className={styles.loginBtn}
          type="submit"
        >
          <LogIn size={16} />
          <span>Login</span>
        </button>
      </form>
      <p className={styles.noAccountText}>
        Don't have an account? <button type="button">Register</button>
      </p>
    </React.Fragment>
  );
};

export default Form;
