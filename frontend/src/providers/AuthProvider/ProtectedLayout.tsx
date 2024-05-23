import { useNavigate, useOutlet } from "react-router-dom";
import { useAuth } from "./AuthProvider";
import { useEffect } from "react";
import Navbar from "../../components/Navbar/Navbar";
import "./ProtectedLayout.scss";

export const ProtectedLayout = () => {
  const {
    authenticationInfo: { expires },
  } = useAuth();
  const outlet = useOutlet();
  const navigate = useNavigate();

  useEffect(() => {
    if (!expires || expires < Date.now() / 1000) {
      navigate("/");
    }
  }, [expires, navigate]);

  return (
    <div className="protected-layout">
      <Navbar />
      {outlet}
    </div>
  );
};
