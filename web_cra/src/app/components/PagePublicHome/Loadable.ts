/**
 *
 * Asynchronously loads the component for PagePublicHome
 *
 */

import { lazyLoad } from 'utils/loadable';

export const PagePublicHome = lazyLoad(
  () => import('./index'),
  module => module.PagePublicHome,
);
