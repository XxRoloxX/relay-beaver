import {
  ReactNode,
  createContext,
  useContext,
  useEffect,
  useState,
} from "react";

export interface AuthenticationInfo {
  email: string;
  expires: number;
}

export interface AuthenticationContext {
  authenticationInfo: AuthenticationInfo;
  setAuthenticationInfo: (value: AuthenticationInfo) => void;
}

export const AuthContext = createContext<AuthenticationContext>({
  authenticationInfo: { email: "", expires: 0 },
  setAuthenticationInfo: () => {},
});

export const AuthProvider = (props: {
  children: ReactNode;
  authenticationInfo: AuthenticationInfo;
}) => {
  const [authenticationInfo, setAuthenticationInfo] =
    useState<AuthenticationInfo>(props.authenticationInfo);

  useEffect(() => {
    setAuthenticationInfo(props.authenticationInfo);
  }, [props, setAuthenticationInfo]);

  return (
    <AuthContext.Provider value={{ authenticationInfo, setAuthenticationInfo }}>
      {props.children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  return useContext(AuthContext);
};
