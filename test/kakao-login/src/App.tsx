import { Route, Routes } from "react-router-dom";
import Home from "pages/Home";
import AuthEnd from "pages/AuthEnd";
import NotFound from "pages/NotFound";

export const HomePath = "/";
export const AuthEndPath = "/auth-end";

export default function App() {
  return (
    <>
      <Routes>
        <Route path={HomePath} element={<Home />} />
        <Route path={AuthEndPath} element={<AuthEnd />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </>
  );
}
