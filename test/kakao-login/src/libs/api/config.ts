import client from "libs/api/client";

export type Config = {
  kakao: Kakao;
};
type Kakao = {
  restApiKey: string;
  redirectUri: string;
};

export async function GetConfig() {
  try {
    const { data } = await client.get<Config>(`/configs`);
    return data;
  } catch (error) {
    throw new Error("GetConfig() 에러");
  }
}
