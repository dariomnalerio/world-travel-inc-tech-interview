import MainContentLayout from "../../Layout/main-content-layout";
import { Card } from "../../ui/card/card";
import styles from "../auth.module.css";
import Form from "./register-form";

const Register = () => {
  return (
    <MainContentLayout>
      <Card className={styles.card}>
        <Card.Title>Register</Card.Title>
        <Card.Content>
          <Form />
        </Card.Content>
      </Card>
    </MainContentLayout>
  );
};

export default Register;
