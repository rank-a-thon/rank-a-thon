import { PagePublicAuthMobileState } from 'app/containers/PagePublicAuthMobile/types';
// [IMPORT NEW CONTAINERSTATE ABOVE] < Needed for generating containers seamlessly

/* 
  Because the redux-injectors injects your reducers asynchronously somewhere in your code
  You have to declare them here manually
*/
export interface RootState {
  pagePublicAuthMobile?: PagePublicAuthMobileState;
  // [INSERT NEW REDUCER KEY ABOVE] < Needed for generating containers seamlessly
}
