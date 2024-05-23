import {
  Route,
  createBrowserRouter,
  createRoutesFromElements,
} from "react-router-dom";
import Login from "./Login/Login";
import { ProtectedLayout } from "../providers/AuthProvider/ProtectedLayout";
import Traffic from "./Traffic/Traffic";
import { AuthLayout } from "../providers/AuthProvider/AuthLayout";

const getUserData = () => {
  return { user: window.localStorage.getItem("user") };
};

const router = createBrowserRouter(
  createRoutesFromElements(
    <Route element={<AuthLayout />} loader={() => getUserData()}>
      <Route path="/" element={<Login />} />
      <Route path="/auth" element={<ProtectedLayout />}>
        <Route path="traffic" element={<Traffic />} />
        <Route path="weather" element={<div>WeatherPage</div>} />
      </Route>
    </Route>,
  ),
);

export default router;
