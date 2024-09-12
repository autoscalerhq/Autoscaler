import {ApiClient} from '~/api-client/ApiClient';
import {NextJsApiClient} from '~/api-client/NextjsApiClient';
import {GoBackendUrl} from '~/config';

export function useApiOnClient() {
  return new ApiClient(GoBackendUrl);
}

export function useNextJsApiOnClient() {
  return new NextJsApiClient('http://localhost:3000');
}