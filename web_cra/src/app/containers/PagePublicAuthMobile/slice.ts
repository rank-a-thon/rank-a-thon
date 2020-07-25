import { PayloadAction } from '@reduxjs/toolkit';
import { createSlice } from 'utils/@reduxjs/toolkit';
import { ContainerState } from './types';

// The initial state of the PagePublicAuthMobile container
export const initialState: ContainerState = {};

const pagePublicAuthMobileSlice = createSlice({
  name: 'pagePublicAuthMobile',
  initialState,
  reducers: {
    someAction(state, action: PayloadAction<any>) {},
  },
});

export const { actions, reducer, name: sliceKey } = pagePublicAuthMobileSlice;
