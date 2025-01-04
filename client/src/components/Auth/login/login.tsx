import MainContentLayout from "../../Layout/main-content-layout";
import { Card } from "../../ui/card/card";
import styles from "./login.module.css";
import Form from "./form";

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
