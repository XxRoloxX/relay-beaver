import {
  ReactNode,
  createContext,
  useContext,
  useState,
  // useState,
} from "react";

export const AuthContext = createContext<string>("");

export const AuthProvider = (props: { children: ReactNode; user: string }) => {
  const [user] = useState<string>(props.user);

  return (
    <AuthContext.Provider value={user}>{props.children}</AuthContext.Provider>
  );
};

export const useAuth = () => {
  const user = useContext(AuthContext);
  return user;
};
