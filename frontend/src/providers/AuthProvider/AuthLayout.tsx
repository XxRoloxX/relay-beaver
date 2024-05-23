import { Await, useLoaderData, useOutlet } from "react-router-dom";
import { AuthProvider } from "./AuthProvider";
import { Suspense } from "react";

export const AuthLayout = () => {
  const outlet = useOutlet();
  const prop = useLoaderData() as { user: string };

  return (
    <Suspense fallback={<div>Loading...</div>}>
      <Await
        resolve={prop.user}
        children={(user) => {
          return <AuthProvider user={user}>{outlet}</AuthProvider>;
        }}
      />
    </Suspense>
  );
};
