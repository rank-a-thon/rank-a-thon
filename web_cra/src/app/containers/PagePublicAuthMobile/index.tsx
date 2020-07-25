/**
 *
 * PagePublicAuthMobile
 *
 */

import React from 'react';
import { Helmet } from 'react-helmet-async';
import { useSelector, useDispatch } from 'react-redux';
import styled from 'styled-components/macro';

import { useInjectReducer, useInjectSaga } from 'utils/redux-injectors';
import { reducer, sliceKey } from './slice';
import { selectPagePublicAuthMobile } from './selectors';
import { pagePublicAuthMobileSaga } from './saga';

interface Props {}

export function PagePublicAuthMobile(props: Props) {
  useInjectReducer({ key: sliceKey, reducer: reducer });
  useInjectSaga({ key: sliceKey, saga: pagePublicAuthMobileSaga });

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const pagePublicAuthMobile = useSelector(selectPagePublicAuthMobile);
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const dispatch = useDispatch();

  return (
    <>
      <Helmet>
        <title>PagePublicAuthMobile</title>
        <meta
          name="description"
          content="Description of PagePublicAuthMobile"
        />
      </Helmet>
      <Div></Div>
    </>
  );
}

const Div = styled.div``;
