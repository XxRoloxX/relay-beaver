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

const getUserData = () => {
  return getTokenInfo();
};

const router = createBrowserRouter(
  createRoutesFromElements(
    <Route
      element={<AuthLayout />}
      loader={() => {
        const userDataPromise = getUserData();
        return defer({
          userData: userDataPromise,
        });
      }}
    >
      <Route path="/" element={<Login />} />
      <Route path="/" element={<ProtectedLayout />}>
        <Route path="traffic" element={<Traffic />} />
        <Route path="weather" element={<div>WeatherPage</div>} />
      </Route>
    </Route>,
  ),
);

export default router;
