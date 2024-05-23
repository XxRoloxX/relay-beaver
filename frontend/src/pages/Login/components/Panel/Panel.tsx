import "./Panel.scss";
import googleLogo from "@/assets/google-logo.webp";

const LoginPanel = () => {
  return (
    <div className="login-panel">
      <div className="login-panel__header">Let's get started</div>
      <div className="login-panel__body">
        <div className="login-panel__body__subtitle">Sign in</div>
        <button className="login-panel__body__button">
          <img
            src={googleLogo}
            alt="Google logo"
            className="login-panel__body__button__logo"
            width="40"
          />
          <span className="login-panel__body__button__text">
            Sign in with Google{" "}
          </span>
        </button>
      </div>
    </div>
  );
};
export default LoginPanel;
