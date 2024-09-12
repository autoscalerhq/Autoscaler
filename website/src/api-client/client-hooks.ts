import {ApiClient} from '~/api-client/ApiClient';
import {NextJsApiClient} from '~/api-client/NextjsApiClient';

export function useApiOnClient() {
  return new ApiClient('http://localhost:4000');
}

export function useNextJsApiOnClient() {
  return new NextJsApiClient('http://localhost:3000');
}