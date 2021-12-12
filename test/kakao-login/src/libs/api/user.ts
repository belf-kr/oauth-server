import axios from "axios";
import client from "libs/api/client";
import {
  getLocalStorageAccessToken,
  getLocalStorageRefreshToken,
  setLocalStorageAccessToken,
  setLocalStorageRefreshToken,
} from "libs/local-storage";

export type UserInfo = {
  id: number;
  email: string;
  name: string;
  avatarImage: string;
};

type JWTToken = {
  accessToken: string;
  refreshToken: string;
};

type JWTRefreshToken = {
  refreshToken: string;
};

export async function GetUserInfo(): Promise<UserInfo> {
  async function work() {
    const accessToken = getLocalStorageAccessToken();
    const { data } = await client.get<UserInfo>(`/users`, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    });
    return data;
  }
  try {
    return await work();
  } catch (error) {
    if (error instanceof Error) {
      if (axios.isAxiosError(error)) {
        switch (error.response?.status) {
          case 401:
            // 재시도: 리프레쉬 토큰으로 엑세스 토큰을 다시 발급
            await TokenRefresh();
            return await work();
        }
      }
    }
    throw new Error("GetUserInfo() 에러");
  }
}

export async function TokenRefresh(): Promise<void> {
  try {
    const refreshToken = getLocalStorageRefreshToken();
    if (refreshToken === null) {
      throw new Error("refreshToken이 없습니다.");
    }
    const body: JWTRefreshToken = {
      refreshToken: refreshToken,
    };
    const { data } = await client.post<JWTToken>(`/users/token/refresh`, body);
    setLocalStorageAccessToken(data.accessToken);
    setLocalStorageRefreshToken(data.refreshToken);
  } catch (error) {
    throw new Error("TokenRefresh() 에러");
  }
}
