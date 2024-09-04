import {ApiClient} from '@/app/api-client/ApiClient';
import {NextJsApiClient} from '@/app/api-client/NextjsApiClient';

export function useApiOnClient() {
  return new ApiClient('http://localhost:4000');
}

export function useNextJsApiOnClient() {
  return new NextJsApiClient('http://localhost:3000');
}