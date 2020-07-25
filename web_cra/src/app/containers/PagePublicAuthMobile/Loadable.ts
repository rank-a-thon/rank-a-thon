/**
 *
 * Asynchronously loads the component for PagePublicAuthMobile
 *
 */

import { lazyLoad } from 'utils/loadable';

export const PagePublicAuthMobile = lazyLoad(
  () => import('./index'),
  module => module.PagePublicAuthMobile,
);
