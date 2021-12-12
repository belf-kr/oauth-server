import { HomePath } from "App";
import {
  setLocalStorageAccessToken,
  setLocalStorageRefreshToken,
} from "libs/local-storage";
import { useEffect } from "react";
import { useLocation, useNavigate } from "react-router-dom";

export default function AuthEnd() {
  const navigate = useNavigate();
  const queryString = new URLSearchParams(useLocation().search);

  const accessToken = queryString.get("accessToken");
  const refreshToken = queryString.get("refreshToken");

  useEffect(() => {
    navigate(HomePath, { replace: true });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  if (accessToken === null || refreshToken === null) {
    return (
      <>
        <h3>토큰이 정상적으로 발급되지 않았습니다.</h3>
      </>
    );
  }

  setLocalStorageAccessToken(accessToken);
  setLocalStorageRefreshToken(refreshToken);

  return (
    <>
      <h3>로그인 처리 중...</h3>
    </>
  );
}
