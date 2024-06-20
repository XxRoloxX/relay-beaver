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
  const { authenticationInfo, setAuthenticationInfo, isTokenValid } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (code) {
      login(code)
        .then((authenticationInfo) => {
          setAuthenticationInfo(authenticationInfo);
          navigate("/config");
        })
        .catch((error) => {
          console.error("Failed to login", error);
        });
    }

    if (isTokenValid()) {
      navigate("/config");
    }
  }, [code, setAuthenticationInfo, navigate, isTokenValid]);

  return authenticationInfo;
};

export default useLogin;
