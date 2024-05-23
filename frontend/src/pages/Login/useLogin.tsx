import { useEffect, useState } from "react";
import { useAuth } from "../../providers/AuthProvider/AuthProvider";
import { login } from "../../api/proxyApi";
import { useNavigate } from "react-router-dom";

const getCodeFromParams = () => {
  const queryParams = new URLSearchParams(window.location.search);
  return queryParams.get("code");
};

const useLogin = () => {
  const [code] = useState(getCodeFromParams());
  const { authenticationInfo, setAuthenticationInfo } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (code) {
      login(code)
        .then((authenticationInfo) => {
          setAuthenticationInfo(authenticationInfo);
          navigate("/auth/traffic");
        })
        .catch((error) => {
          console.error("Failed to login", error);
        });
    }
  }, [code, setAuthenticationInfo, navigate]);

  return authenticationInfo;
};

export default useLogin;
