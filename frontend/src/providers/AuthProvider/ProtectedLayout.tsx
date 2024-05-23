import { useNavigate, useOutlet } from "react-router-dom";
import { useAuth } from "./AuthProvider";
import { useEffect } from "react";

export const ProtectedLayout = () => {
  const user = useAuth();
  const outlet = useOutlet();
  const navigate = useNavigate();

  useEffect(() => {
    if (!user) {
      console.log("Redirecting to login page");
      navigate("/");
    }
  }, [user, navigate]);

  return <>{outlet}</>;
};
