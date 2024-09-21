import {getAccessToken} from '~/config/supertokens/helpers';
import {ApiClient} from '~/api-client/ApiClient';
import {NextJsApiClient} from '~/api-client/NextjsApiClient';
import {GoBackendUrl} from '~/config';

export function useApiOnServer() {
  const accessTokenPayload = getAccessToken();
  return new ApiClient(GoBackendUrl, accessTokenPayload);
}

export function useNextJsApiOnServer() {
  const accessTokenPayload = getAccessToken();
  return new NextJsApiClient('http://localhost:3000', accessTokenPayload);
}