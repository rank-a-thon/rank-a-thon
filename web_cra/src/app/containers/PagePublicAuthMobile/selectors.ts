import { createSelector } from '@reduxjs/toolkit';

import { RootState } from 'types';
import { initialState } from './slice';

const selectDomain = (state: RootState) =>
  state.pagePublicAuthMobile || initialState;

export const selectPagePublicAuthMobile = createSelector(
  [selectDomain],
  pagePublicAuthMobileState => pagePublicAuthMobileState,
);
