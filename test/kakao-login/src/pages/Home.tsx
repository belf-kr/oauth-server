import axios from "axios";
import { Config, GetConfig } from "libs/api/config";
import { GetUserInfo, UserInfo } from "libs/api/user";
import {
  delLocalStorageAccessToken,
  delLocalStorageRefreshToken,
  getLocalStorageAccessToken,
} from "libs/local-storage";
import { useEffect, useState } from "react";

export default function Home() {
  const [config, setConfig] = useState<Config>();
  const [userInfo, setUserInfo] = useState<UserInfo>();
  const [error, setError] = useState<any>();

  const accessToken = getLocalStorageAccessToken();

  function handleKakaoLogin() {
    if (!config) {
      throw new Error("config 값이 없습니다");
    }

    const clientId = config.kakao.restApiKey;
    const redirectUri = config.kakao.redirectUri;
    window.location.href = `https://kauth.kakao.com/oauth/authorize?client_id=${clientId}&redirect_uri=${redirectUri}&response_type=code`;
  }

  useEffect(() => {
    (async () => {
      try {
        if (!accessToken) {
          return;
        }
        const res = await GetUserInfo();
        setUserInfo(res);
      } catch (error) {
        if (error instanceof Error) {
          if (axios.isAxiosError(error)) {
            switch (error.response?.status) {
              case 401:
                // refresh token도 만료된 상태입니다: 사용자가 다시 로그인해야되는 시점
                delLocalStorageAccessToken();
                delLocalStorageRefreshToken();
                window.location.reload();
                break;
            }
          }
          setError(error.message);
          return;
        }
        setError(error);
      }
    })();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  useEffect(() => {
    (async () => {
      try {
        const res = await GetConfig();
        setConfig(res);
      } catch (error) {
        if (error instanceof Error) {
          setError(error.message);
          return;
        }
        setError(error);
      }
    })();
  }, []);

  if (error) {
    return (
      <>
        <h3>에러 발생</h3>
        <span>{error}</span>
      </>
    );
  }

  if (!config) {
    return (
      <>
        <h3>config 조회 중</h3>
      </>
    );
  }

  return (
    <>
      <button onClick={handleKakaoLogin}>카카오 로그인</button>
      {userInfo ? (
        <>
          <h3>카카오 로그인 완료</h3>
          <img
            src={userInfo.avatarImage}
            alt="사용자 프로필 사진이 없는거 같습니다"
            style={{ width: "150px" }}
          />
          <h5>id: {userInfo.id}</h5>
          <h5>name: {userInfo.name}</h5>
          <h5>email: {userInfo.email}</h5>
        </>
      ) : (
        <>
          <h3>로그인되지 않았거나 token이 손상됨</h3>
        </>
      )}
    </>
  );
}
