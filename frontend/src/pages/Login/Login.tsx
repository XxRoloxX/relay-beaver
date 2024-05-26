import LoginBanner from "./components/Banner/Banner";
import LoginPanel from "./components/Panel/Panel";
import "./Login.scss";
import useLogin from "./useLogin";

const Login = () => {
  useLogin();

  return (
    <div className="login">
      <LoginBanner />
      <LoginPanel />
    </div>
  );
};

export default Login;
