import {getAccessToken} from '@/app/config/supertokens/helpers';
import {ApiClient} from '@/app/api-client/ApiClient';
import {NextJsApiClient} from '@/app/api-client/NextjsApiClient';

export function useApiOnServer() {
  const accessTokenPayload = getAccessToken();
  return new ApiClient('http://localhost:4000', accessTokenPayload);
}

export function useNextJsApiOnServer() {
  const accessTokenPayload = getAccessToken();
  return new NextJsApiClient('http://localhost:3000', accessTokenPayload);
}