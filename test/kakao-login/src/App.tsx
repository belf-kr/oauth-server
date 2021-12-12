import AuthEnd from "pages/AuthEnd";
import Home from "pages/Home";
import NotFound from "pages/NotFound";
import { Route, Routes } from "react-router-dom";

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
