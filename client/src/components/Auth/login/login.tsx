import MainContentLayout from "../../Layout/main-content-layout";
import { Card } from "../../ui/card/card";
import styles from "../auth.module.css";
import Form from "./login-form";

const Login = () => {
  return (
    <MainContentLayout>
      <Card className={styles.card}>
        <Card.Title>Login</Card.Title>
        <Card.Content>
          <Form />
        </Card.Content>
      </Card>
    </MainContentLayout>
  );
};

export default Login;
