import {
  Route,
  createBrowserRouter,
  createRoutesFromElements,
  defer,
} from "react-router-dom";
import Login from "./Login/Login";
import { ProtectedLayout } from "../providers/AuthProvider/ProtectedLayout";
import Traffic from "./Traffic/Traffic";
import { AuthLayout } from "../providers/AuthProvider/AuthLayout";
import { getTokenInfo } from "../api/proxyApi";
import Stats from "./Stats/Stats";
import Config from "./Config/Config";

const getUserData = () => {
  return getTokenInfo();
};

const router = createBrowserRouter(
  createRoutesFromElements(
    <Route
      element={<AuthLayout />}
      loader={() => {
        return defer({
          userData: getUserData(),
        });
      }}
      errorElement={<Login />}
    >
      <Route path="/" element={<Login />} />
      <Route path="/" element={<ProtectedLayout />}>
        <Route path="traffic" element={<Traffic />} />
        <Route path="stats" element={<Stats />} />
        <Route path="config" element={<Config />} />
      </Route>
    </Route>,
  ),
);

export default router;
