import logo from "@/assets/relay-beaver.png";
import "./Banner.scss";

const LoginBanner = () => {
  return (
    <div className="login-banner">
      <img src={logo} alt="Relay Beaver" className="login-banner__logo" />
      <h1 className="login-banner__title">Relay Beaver</h1>
      <p className="login-banner__subtitle">Proxy your way to observability</p>
    </div>
  );
};

export default LoginBanner;
