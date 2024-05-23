import LoginBanner from "./components/Banner/Banner";
import LoginPanel from "./components/Panel/Panel";
import "./Login.scss";

const Login = () => {
  return (
    <div className="login">
      <LoginBanner />
      <LoginPanel />
    </div>
  );
};

export default Login;
